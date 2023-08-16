package eventservice

import (
	"context"
	"crypto/ecdsa"
	"errors"
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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type Infura struct {
	wsBaseURL    string
	httpsBaseURL string
	apiKey       string
	proxyAddress string
	potID        uint16
	privateKey   string
}

const (
	ChainMainnet = iota
	ChainSepolia
	ChainXDC
	ChainGoerli
)

func getChainIDForChain(chain int) int64 {
	switch chain {
	case ChainMainnet:
		return 1
	case ChainSepolia:
		return 11155111
	case ChainXDC:
		return 50
	case ChainGoerli:
		return 5
	}

	return 1
}

func InitializeInfura(proxyAddress, rpcHttpsBaseURL, rpcWsBaseURL, rpcApiKey, privateKey string) *Infura {
	return &Infura{proxyAddress: proxyAddress, httpsBaseURL: rpcHttpsBaseURL, wsBaseURL: rpcWsBaseURL, apiKey: rpcApiKey, potID: 1, privateKey: privateKey}
}

func (i *Infura) TestFulfilled(userDBHandle db.UserTickets) error {
	dialURL := i.wsBaseURL + i.apiKey
	client, err := ethclient.Dial(dialURL)
	if err != nil {
		return fmt.Errorf("dial problem: %v", err)
	}

	addr := common.HexToAddress(i.proxyAddress)
	instance, err := hotpot.NewHotpot(addr, client)
	if err != nil {
		return fmt.Errorf("instance error: %v", err)
	}

	marketplaceAbiFile, err := os.Open("config/hotpot-abi.json")
	if err != nil {
		return fmt.Errorf("open file problem: %v", err)
	}

	contractAbi, err := abi.JSON(marketplaceAbiFile)
	if err != nil {
		return fmt.Errorf("abi json problem: %v", err)
	}

	potID := uint16(1)
	winners, err := userDBHandle.GetWinnersForPot(ChainGoerli, potID)
	if err != nil {
		fmt.Println("Error while setting winner for pot ", potID, " and chain ", ChainGoerli, " error ", err)
		return err
	}

	limit, err := instance.PotLimit(nil)
	if err != nil {
		return err
	}

	numOfWinners, err := instance.NumberOfWinners(nil)
	if err != nil {
		return err
	}

	amountPerWin := limit.Int64() / int64(numOfWinners)

	winnerMap := make(map[string]int)
	winnerAddresses := make([]common.Address, 0)
	winnerAmounts := make([]*big.Int, 0)
	for _, win := range winners {
		if i, ok := winnerMap[win.WalletAddress]; ok {
			winnerAmounts[i].Add(winnerAmounts[i], big.NewInt(amountPerWin))
		}

		winnerAddresses = append(winnerAddresses, common.HexToAddress(win.WalletAddress))
		winnerAmounts = append(winnerAmounts, big.NewInt(amountPerWin))
		winnerMap[win.WalletAddress] = len(winnerAddresses) - 1
	}

	return executeRaffle(ChainGoerli, i.httpsBaseURL+i.apiKey, i.privateKey, i.proxyAddress, winnerAddresses, winnerAmounts, contractAbi)
}

func (i *Infura) Start(userDBHandle db.UserTickets, log *logrus.Entry) {
	go func() {
		defer RecoverFromPanic(log)

		err := i.init(userDBHandle)
		if err != nil {
			log.WithError(err).Error("Init to WS failed")
		}

		for {
			err := i.listen(userDBHandle)
			if err != nil {
				log.WithError(err).Error("Listening to WS failed")
			}
		}
	}()
}

func (i *Infura) init(userDBHandle db.UserTickets) error {
	dialURL := i.wsBaseURL + i.apiKey
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

	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case logGenerateTicketsSigHash.Hex():
			fmt.Println("GENERATE TICKETS")
			err = i.tryRaffleTicketsCatch(ChainGoerli, instance, userDBHandle, contractAbi, vLog)
			if err != nil {
				fmt.Printf("unpack raffle tickets catch problem: %v \n", err)
			}
		case logRandomWordReqSigHash.Hex():
			fmt.Println("RANDOM WORD REQ")
			err = i.tryRandomWordRequestedCatch(ChainGoerli, instance, contractAbi, vLog)
			if err != nil {
				fmt.Printf("unpack random word requested catch problem: %v \n", err)
			}
		case logRandomWordFulSigHash.Hex():
			fmt.Println("RANDOM WORD FUL")
			err = i.tryRandomWordFulfilledCatch(ChainGoerli, instance, userDBHandle, contractAbi, vLog)
			if err != nil {
				fmt.Printf("unpack random word fulfilled catch problem: %v \n", err)
			}
		}
	}

	return nil
}

func (i *Infura) listen(userDBHandle db.UserTickets) error {
	dialURL := i.wsBaseURL + i.apiKey
	client, err := ethclient.Dial(dialURL)
	if err != nil {
		return fmt.Errorf("dial problem: %v", err)
	}

	addr := common.HexToAddress(i.proxyAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}

	// logs, err := client.FilterLogs(context.Background(), query)
	// if err != nil {
	// 	return fmt.Errorf("logs problem: %v", err)
	// }
	logs := make(chan types.Log)

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

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return fmt.Errorf("subscribe to logs problem: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			return fmt.Errorf("log sub problem %v", err)
		case vLog := <-logs:
			fmt.Println("TX HASH: ", vLog.TxHash.Hex())
			// for _, vLog := range logs {
			switch vLog.Topics[0].Hex() {
			case logGenerateTicketsSigHash.Hex():
				fmt.Println("GENERATE TICKETS")
				err = i.tryRaffleTicketsCatch(ChainGoerli, instance, userDBHandle, contractAbi, vLog)
				if err != nil {
					fmt.Printf("unpack raffle tickets catch problem: %v \n", err)
				}
			case logRandomWordReqSigHash.Hex():
				fmt.Println("RANDOM WORD REQ")
				err = i.tryRandomWordRequestedCatch(ChainGoerli, instance, contractAbi, vLog)
				if err != nil {
					fmt.Printf("unpack random word requested catch problem: %v \n", err)
				}
			case logRandomWordFulSigHash.Hex():
				fmt.Println("RANDOM WORD FUL")
				err = i.tryRandomWordFulfilledCatch(ChainGoerli, instance, userDBHandle, contractAbi, vLog)
				if err != nil {
					fmt.Printf("unpack random word fulfilled catch problem: %v \n", err)
				}
			}
		}
	}

	// return nil
}

// }

func (infura *Infura) tryRandomWordRequestedCatch(chain int, hotpot *hotpot.Hotpot, contractAbi abi.ABI, vLog types.Log) error {
	_, err := contractAbi.Unpack("RandomWordRequested", vLog.Data)
	if err != nil {
		return err
	}

	infura.potID += 1
	fmt.Println("Pot incremented by 1, now ", infura.potID)

	return nil
}

func (i *Infura) GetCurrentPot() (uint16, error) {
	dialURL := i.wsBaseURL + i.apiKey
	client, err := ethclient.Dial(dialURL)
	if err != nil {
		return 0, fmt.Errorf("dial problem: %v", err)
	}

	addr := common.HexToAddress(i.proxyAddress)

	instance, err := hotpot.NewHotpot(addr, client)
	if err != nil {
		return 0, fmt.Errorf("instance error: %v", err)
	}

	potID, err := instance.CurrentPotId(nil)
	if err != nil {
		return 0, err
	}

	return potID, nil
}

func (infura *Infura) tryRandomWordFulfilledCatch(chain int, hotpot *hotpot.Hotpot, userDBHandle db.UserTickets, contractAbi abi.ABI, vLog types.Log) error {
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

	err = userDBHandle.SetPotRaffleTimestamp(chain, potID)
	if err != nil {
		return err
	}

	winningTicketIds, err := hotpot.GetWinningTicketIds(nil, potID)
	if err != nil {
		return err
	}

	for _, id := range winningTicketIds {
		fmt.Println("Setting winner for ", potID, " ID ", id)
		err = userDBHandle.SetWinnerForPot(chain, potID, id)
		if err != nil {
			fmt.Println("Error while setting winner for pot ", potID, " and ticket ", id, " error ", err)
		}
	}

	winners, err := userDBHandle.GetWinnersForPot(chain, potID)
	if err != nil {
		fmt.Println("Error while setting winner for pot ", potID, " and chain ", chain, " error ", err)
		return err
	}

	limit, err := hotpot.PotLimit(nil)
	if err != nil {
		return err
	}

	numOfWinners, err := hotpot.NumberOfWinners(nil)
	if err != nil {
		return err
	}

	amountPerWin := limit.Int64() / int64(numOfWinners)

	winnerMap := make(map[string]int)
	winnerAddresses := make([]common.Address, 0)
	winnerAmounts := make([]*big.Int, 0)
	for _, win := range winners {
		if i, ok := winnerMap[win.WalletAddress]; ok {
			winnerAmounts[i].Add(winnerAmounts[i], big.NewInt(amountPerWin))
		}

		winnerAddresses = append(winnerAddresses, common.HexToAddress(win.WalletAddress))
		winnerAmounts = append(winnerAmounts, big.NewInt(amountPerWin))
		winnerMap[win.WalletAddress] = len(winnerAddresses) - 1
	}

	return executeRaffle(chain, infura.httpsBaseURL+infura.apiKey, infura.privateKey, infura.proxyAddress, winnerAddresses, winnerAmounts, contractAbi)
}

func executeRaffle(chain int, dialString, pvtKey, proxyAddress string, winningAddresses []common.Address, amounts []*big.Int, contractAbi abi.ABI) error {
	fmt.Println("Addresses: ", winningAddresses)
	fmt.Println("Amounts: ", amounts)
	client, err := ethclient.Dial(dialString)
	if err != nil {
		return fmt.Errorf("dial error: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(pvtKey)
	if err != nil {
		return fmt.Errorf("hex error: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("nonce error: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("gas error: %v", err)
	}

	data, err := contractAbi.Pack("executeRaffle", winningAddresses, amounts)
	if err != nil {
		return fmt.Errorf("pack error: %v", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(getChainIDForChain(chain)))
	if err != nil {
		return fmt.Errorf("chain error: %v", err)
	}

	address := common.HexToAddress(proxyAddress)
	gas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		Data:  data,
		Value: big.NewInt(0),
		To:    &address,
		From:  fromAddress,
	})
	if err != nil {
		return fmt.Errorf("gas error: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = gas + 10000 // in units
	auth.GasPrice = gasPrice

	instance, err := hotpot.NewHotpot(address, client)
	if err != nil {
		return fmt.Errorf("instance error: %v", err)
	}

	tx, err := instance.ExecuteRaffle(auth, winningAddresses, amounts)
	if err != nil {
		return fmt.Errorf("execute error: %v", err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	return nil
}

func (infura *Infura) tryRaffleTicketsCatch(chain int, hotpot *hotpot.Hotpot, userDBHandle db.UserTickets, contractAbi abi.ABI, vLog types.Log) error {
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
			err = userDBHandle.Insert(chain, models.UserTickets{WalletAddress: event.Buyer, TicketID: i, PotID: infura.potID})
			if err != nil {
				fmt.Printf("insert buyer tickets problem: %v \n", err)
			}
		}
	}

	if event.SellerTicketIdStart > 0 && event.SellerTicketIdEnd > 0 {
		for i := event.SellerTicketIdStart; i <= event.SellerTicketIdEnd; i++ {
			err = userDBHandle.Insert(chain, models.UserTickets{WalletAddress: event.Seller, TicketID: i, PotID: infura.potID})
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
