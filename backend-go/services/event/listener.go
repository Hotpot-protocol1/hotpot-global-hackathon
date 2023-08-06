package eventservice

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/Hotpot-protocol1/hotpot-global/db"
	"github.com/Hotpot-protocol1/hotpot-global/db/models"
	"github.com/Hotpot-protocol1/hotpot-global/hotpot"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type Infura struct {
	baseURL      string
	apiKey       string
	proxyAddress string
	potID        uint16
}

func InitializeInfura(proxyAddress, rpcBaseURL, rpcApiKey string) Infura {
	return Infura{proxyAddress: proxyAddress, baseURL: rpcBaseURL, apiKey: rpcApiKey, potID: 1}
}

func (i *Infura) Start(userDBHandle db.UserTickets, log *logrus.Entry) {
	go func() {
		defer RecoverFromPanic(log)

		// for {
		err := i.listen(userDBHandle)
		if err != nil {
			log.WithError(err).Error("Listening to WS failed")
		}
		// }
	}()
}

func (i *Infura) listen(userDBHandle db.UserTickets) error {
	dialURL := i.baseURL + i.apiKey
	client, err := ethclient.Dial(dialURL)
	if err != nil {
		return fmt.Errorf("dial problem: %v", err)
	}

	addr := common.HexToAddress(i.proxyAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("logs problem: %v", err)
	}
	// logs := make(chan types.Log)

	marketplaceAbiFile, err := os.Open("config/hotpot-abi.json")
	if err != nil {
		return fmt.Errorf("open file problem: %v", err)
	}

	contractAbi, err := abi.JSON(marketplaceAbiFile)
	if err != nil {
		return fmt.Errorf("abi json problem: %v", err)
	}

	instance, err := hotpot.NewHotpot(addr, client)
	if err != nil {
		return fmt.Errorf("instance error: %v", err)
	}

	logGenerateTicketsSig := []byte("GenerateRaffleTickets(address,address,uint32,uint32,uint32,uint32,uint256,uint256)")
	logRandomWordReqSig := []byte("RandomWordRequested(uint256,uint32,uint32)")
	logRandomWordFulSig := []byte("RandomnessFulfilled(uint16,uint256)")
	logGenerateTicketsSigHash := crypto.Keccak256Hash(logGenerateTicketsSig)
	logRandomWordReqSigHash := crypto.Keccak256Hash(logRandomWordReqSig)
	logRandomWordFulSigHash := crypto.Keccak256Hash(logRandomWordFulSig)

	// sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	// if err != nil {
	// 	return fmt.Errorf("subscribe to logs problem: %v", err)
	// }

	// for {
	// 	select {
	// 	case err := <-sub.Err():
	// 		return fmt.Errorf("log sub problem %v", err)
	// 	case vLog := <-logs:
	// 		fmt.Println("TX HASH: ", vLog.TxHash.Hex())
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case logGenerateTicketsSigHash.Hex():
			fmt.Println("GENERATE TICKETS")
			err = i.tryRaffleTicketsCatch(instance, userDBHandle, contractAbi, vLog)
			if err != nil {
				fmt.Printf("unpack raffle tickets catch problem: %v \n", err)
			}
		case logRandomWordReqSigHash.Hex():
			fmt.Println("RANDOM WORD REQ")
			err = i.tryRandomWordRequestedCatch(instance, contractAbi, vLog)
			if err != nil {
				fmt.Printf("unpack random word requested catch problem: %v \n", err)
			}
		case logRandomWordFulSigHash.Hex():
			fmt.Println("RANDOM WORD FUL")
			err = i.tryRandomWordFulfilledCatch(instance, userDBHandle, contractAbi, vLog)
			if err != nil {
				fmt.Printf("unpack random word fulfilled catch problem: %v \n", err)
			}
		}
	}

	return nil
}

// }

func (infura *Infura) tryRandomWordRequestedCatch(hotpot *hotpot.Hotpot, contractAbi abi.ABI, vLog types.Log) error {
	_, err := contractAbi.Unpack("RandomWordRequested", vLog.Data)
	if err != nil {
		return err
	}

	infura.potID += 1
	fmt.Println("Pot incremented by 1, now ", infura.potID)

	return nil
}

func (infura *Infura) tryRandomWordFulfilledCatch(hotpot *hotpot.Hotpot, userDBHandle db.UserTickets, contractAbi abi.ABI, vLog types.Log) error {
	m := make(map[string]interface{})
	err := contractAbi.UnpackIntoMap(m, "RandomnessFulfilled", vLog.Data)
	if err != nil {
		return err
	}

	fmt.Println("MAP ", m, " HEX vLOG ", vLog.Topics[1].Hex())
	potID64, err := strconv.ParseUint(strings.Replace(vLog.Topics[1].Hex(), "0x", "", -1), 16, 64)
	if err != nil {
		return err
	}

	potID := uint16(potID64 - 1)
	if potID < 2 {
		fmt.Println("There's an error from contract returning number less than 2 ", potID)
	}

	err = userDBHandle.SetPotRaffleTimestamp(potID)
	if err != nil {
		return err
	}

	winningTicketIds, err := hotpot.GetWinningTicketIds(nil, potID)
	if err != nil {
		return err
	}

	for _, id := range winningTicketIds {
		fmt.Println("Setting winner for ", potID, " ID ", id)
		err = userDBHandle.SetWinnerForPot(potID, id)
		if err != nil {
			fmt.Println("Error while setting winner for pot ", potID, " and ticket ", id, " error ", err)
		}
	}

	return nil
}

func (infura *Infura) tryRaffleTicketsCatch(hotpot *hotpot.Hotpot, userDBHandle db.UserTickets, contractAbi abi.ABI, vLog types.Log) error {
	event := struct {
		Buyer                  string
		Seller                 string
		BuyerTicketIdStart     uint32   `abi:"_buyerTicketIdStart"`
		BuyerTicketIdEnd       uint32   `abi:"_buyerTicketIdEnd"`
		SellerTicketIdStart    uint32   `abi:"_sellerTicketIdStart"`
		SellerTicketIdEnd      uint32   `abi:"_sellerTicketIdEnd"`
		NewBuyerPendingAmount  *big.Int `abi:"_buyerPendingAmount"`
		NewSellerPendingAmount *big.Int `abi:"_sellerPendingAmount"`
	}{}

	err := contractAbi.UnpackIntoInterface(&event, "GenerateRaffleTickets", vLog.Data)
	if err != nil {
		return err
	}

	if infura.potID == 0 {
		infura.potID, err = hotpot.CurrentPotId(nil)
		if err != nil {
			return err
		}
	}

	event.Buyer = vLog.Topics[1].Hex()
	event.Seller = vLog.Topics[2].Hex()

	event.Buyer = "0x" + strings.TrimLeft(event.Buyer[2:], "0")
	event.Seller = "0x" + strings.TrimLeft(event.Seller[2:], "0")

	fmt.Println("New GENERATE RAFFLE TICKETS EVENT")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Buyer:", event.Buyer)
	fmt.Println("Seller:", event.Seller)
	fmt.Println("BuyerTicketIdStart:", event.BuyerTicketIdStart)
	fmt.Println("BuyerTicketIdEnd:", event.BuyerTicketIdEnd)
	fmt.Println("SellerTicketIdStart:", event.SellerTicketIdStart)
	fmt.Println("SellerTicketIdEnd:", event.SellerTicketIdEnd)
	fmt.Println("Pot ID:", infura.potID)
	fmt.Println("-----------------------------------------------------------")

	if event.BuyerTicketIdStart > 0 && event.BuyerTicketIdEnd > 0 {
		for i := event.BuyerTicketIdStart; i <= event.BuyerTicketIdEnd; i++ {
			err = userDBHandle.Insert(models.UserTickets{WalletAddress: event.Buyer, TicketID: i, PotID: infura.potID})
			if err != nil {
				fmt.Printf("insert buyer tickets problem: %v \n", err)
			}
		}
	}

	if event.SellerTicketIdStart > 0 && event.SellerTicketIdEnd > 0 {
		for i := event.SellerTicketIdStart; i <= event.SellerTicketIdEnd; i++ {
			err = userDBHandle.Insert(models.UserTickets{WalletAddress: event.Seller, TicketID: i, PotID: infura.potID})
			if err != nil {
				fmt.Printf("insert seller tickets problem: %v \n", err)
			}
		}
	}

	return nil
}

// RecoverFromPanic recovers from eventual runtime panic while handling
func RecoverFromPanic(log *logrus.Entry) {
	if r := recover(); r != nil {
		stackTrace := string(debug.Stack())

		err, ok := r.(error)
		if !ok {
			err = fmt.Errorf("[error type: %T] %v", r, r)
		}

		log.Errorf("Panic while processing request: %v\n%v", err, stackTrace)
	}
}