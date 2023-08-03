package eventservice

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"runtime/debug"

	"github.com/Hotpot-protocol1/hotpot-global/db"
	"github.com/Hotpot-protocol1/hotpot-global/db/models"
	"github.com/Hotpot-protocol1/hotpot-global/hotpot"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type Infura struct {
	baseURL      string
	apiKey       string
	proxyAddress string
}

func InitializeInfura(proxyAddress, rpcBaseURL, rpcApiKey string) Infura {
	return Infura{proxyAddress: proxyAddress, baseURL: rpcBaseURL, apiKey: rpcApiKey}
}

func (i *Infura) Start(userDBHandle db.User, log *logrus.Entry) {
	go func() {
		defer RecoverFromPanic(log)

		err := i.listen(userDBHandle)
		if err != nil {
			log.WithError(err).Error("Listening to WS failed")
		}
	}()
}

func (i *Infura) listen(userDBHandle db.User) error {
	dialURL := i.baseURL + i.apiKey
	client, err := ethclient.Dial(dialURL)
	if err != nil {
		return fmt.Errorf("dial problem: %v", err)
	}

	addr := common.HexToAddress(i.proxyAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}

	// logs := make(chan types.Log)
	// sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for {
	// 	select {
	// 	case err := <-sub.Err():
	// 		return err
	// 	case vLog := <-logs:
	// 		fmt.Println(vLog) // pointer to event log
	// 	}
	// }

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

	fmt.Println(logs)

	for _, vLog := range logs {
		fmt.Println("TX HASH: ", vLog.TxHash.Hex())

		err = tryRaffleTicketsCatch(instance, userDBHandle, contractAbi, vLog)
		if err != nil {
			fmt.Printf("unpack raffle tickets catch problem: %v", err)
		}

	}

	return nil
}

// func tryOfferedCatch() {
// event := struct {
// 	ItemId  *big.Int
// 	Nft     string
// 	TokenId *big.Int
// 	Price   *big.Int
// 	Seller  string
// }{}

// err := contractAbi.UnpackIntoInterface(&event, "Offered", vLog.Data)
// if err == nil {
// 	event.Nft = vLog.Topics[1].Hex()
// 	event.Seller = vLog.Topics[2].Hex()

// 	fmt.Println("New OFFERED EVENT")
// 	fmt.Println("-----------------------------------------------------------")
// 	fmt.Println("Item ID:", event.ItemId)
// 	fmt.Println("NFT address:", event.Nft)
// 	fmt.Println("NFT token ID:", event.TokenId)
// 	fmt.Println("Price:", event.Price)
// 	fmt.Println("Seller address:", event.Seller)
// 	fmt.Println("-----------------------------------------------------------")
// 	fmt.Printf("unpack offered problem: %v", err)
// 	continue
// }
// }

func tryRaffleTicketsCatch(hotpot *hotpot.Hotpot, userDBHandle db.User, contractAbi abi.ABI, vLog types.Log) error {
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

	potID, err := hotpot.CurrentPotId(nil)
	if err != nil {
		return err
	}

	event.Buyer = vLog.Topics[1].Hex()
	event.Seller = vLog.Topics[2].Hex()

	fmt.Println("New GENERATE RAFFLE TICKETS EVENT")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Buyer:", event.Buyer)
	fmt.Println("Seller:", event.Seller)
	fmt.Println("BuyerTicketIdStart:", event.BuyerTicketIdStart)
	fmt.Println("BuyerTicketIdEnd:", event.BuyerTicketIdEnd)
	fmt.Println("SellerTicketIdStart:", event.SellerTicketIdStart)
	fmt.Println("SellerTicketIdEnd:", event.SellerTicketIdEnd)
	fmt.Println("Pot ID:", potID)
	fmt.Println("-----------------------------------------------------------")

	for i := event.BuyerTicketIdStart; i <= event.BuyerTicketIdEnd; i++ {
		err = userDBHandle.Insert(models.User{WalletAddress: event.Buyer, TicketID: i, PotID: potID})
		if err != nil {
			fmt.Printf("insert buyer tickets problem: %v", err)
		}
	}

	for i := event.SellerTicketIdStart; i <= event.SellerTicketIdEnd; i++ {
		err = userDBHandle.Insert(models.User{WalletAddress: event.Seller, TicketID: i, PotID: potID})
		if err != nil {
			fmt.Printf("insert seller tickets problem: %v", err)
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
