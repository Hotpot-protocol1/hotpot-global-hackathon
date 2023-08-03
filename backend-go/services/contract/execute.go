package contractservice

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/Hotpot-protocol1/hotpot-global/hotpot"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Operator struct {
	privateKey string
	baseURL    string
	apiKey     string
}

func InitializeOperator(privateKey, baseURL, apiKey string) Operator {
	return Operator{privateKey: privateKey, baseURL: baseURL, apiKey: apiKey}
}

func (o *Operator) Execute(buyer, seller string, amount int64) error {
	client, err := ethclient.Dial(o.baseURL + o.apiKey)
	if err != nil {
		return fmt.Errorf("Dial error: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(o.privateKey)
	if err != nil {
		return fmt.Errorf("Hex error: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("Nonce error: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("Gas error: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		return fmt.Errorf("Chain error: %v", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(2500000000000000) // in wei
	auth.GasLimit = uint64(300000)            // in units
	auth.GasPrice = gasPrice

	amountInWei := big.NewInt(amount)
	buyerAddr := common.HexToAddress(buyer)
	sellerAddr := common.HexToAddress(seller)
	pendingAmount := big.NewInt(0)

	address := common.HexToAddress("0x8Dc5DfCED235d297d41FA932dBbF19F7aC20D990")
	instance, err := hotpot.NewHotpot(address, client)
	if err != nil {
		return fmt.Errorf("Instance error: %v", err)
	}

	tx, err := instance.ExecuteTrade(auth, amountInWei, buyerAddr, sellerAddr, pendingAmount, pendingAmount)
	if err != nil {
		return fmt.Errorf("Execute error: %v", err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	return nil
}
