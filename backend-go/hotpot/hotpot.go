// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hotpot

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IHotpotInitializeParams is an auto generated low-level Go binding around an user-defined struct.
type IHotpotInitializeParams struct {
	PotLimit         *big.Int
	RaffleTicketCost *big.Int
	ClaimWindow      *big.Int
	NumberOfWinners  uint16
	Fee              uint16
	TradeFee         uint16
	Marketplace      common.Address
	Operator         common.Address
}

// HotpotMetaData contains all meta data concerning the Hotpot contract.
var HotpotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vrfV2Wrapper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_buyerTicketIdStart\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_buyerTicketIdEnd\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_sellerTicketIdStart\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"_sellerTicketIdEnd\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_buyerPendingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_sellerPendingAmount\",\"type\":\"uint256\"}],\"name\":\"GenerateRaffleTickets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newMarketplace\",\"type\":\"address\"}],\"name\":\"MarketplaceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_nOfWinners\",\"type\":\"uint16\"}],\"name\":\"NumberOfWinnersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"OperatorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128[]\",\"name\":\"_newPrizeAmounts\",\"type\":\"uint128[]\"}],\"name\":\"PrizeAmountsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fromTicketId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"toTicketId\",\"type\":\"uint32\"}],\"name\":\"RandomWordRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"potId\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomWord\",\"type\":\"uint256\"}],\"name\":\"RandomnessFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_winners\",\"type\":\"address[]\"}],\"name\":\"WinnersAssigned\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"chainlinkRequests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"fullfilled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"randomWord\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimablePrizes\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deadline\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPotId\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPotSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_winners\",\"type\":\"address[]\"}],\"name\":\"executeRaffle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountInWei\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_buyerPendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sellerPendingAmount\",\"type\":\"uint256\"}],\"name\":\"executeTrade\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_potId\",\"type\":\"uint16\"}],\"name\":\"getWinningTicketIds\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"potLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"raffleTicketCost\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"claimWindow\",\"type\":\"uint128\"},{\"internalType\":\"uint16\",\"name\":\"numberOfWinners\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"fee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"tradeFee\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"marketplace\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structIHotpot.InitializeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRaffleTicketId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextPotTicketIdStart\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfWinners\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potTicketIdEnd\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potTicketIdStart\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"name\":\"prizeAmounts\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleTicketCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requestIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMarketplace\",\"type\":\"address\"}],\"name\":\"setMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPotLimit\",\"type\":\"uint256\"}],\"name\":\"setPotLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newRaffleTicketCost\",\"type\":\"uint256\"}],\"name\":\"setRaffleTicketCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newTradeFee\",\"type\":\"uint16\"}],\"name\":\"setTradeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tradeFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_nOfWinners\",\"type\":\"uint16\"}],\"name\":\"updateNumberOfWinners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128[]\",\"name\":\"_newPrizeAmounts\",\"type\":\"uint128[]\"}],\"name\":\"updatePrizeAmounts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winningTicketIds\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// HotpotABI is the input ABI used to generate the binding from.
// Deprecated: Use HotpotMetaData.ABI instead.
var HotpotABI = HotpotMetaData.ABI

// Hotpot is an auto generated Go binding around an Ethereum contract.
type Hotpot struct {
	HotpotCaller     // Read-only binding to the contract
	HotpotTransactor // Write-only binding to the contract
	HotpotFilterer   // Log filterer for contract events
}

// HotpotCaller is an auto generated read-only Go binding around an Ethereum contract.
type HotpotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HotpotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HotpotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HotpotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HotpotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HotpotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HotpotSession struct {
	Contract     *Hotpot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HotpotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HotpotCallerSession struct {
	Contract *HotpotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HotpotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HotpotTransactorSession struct {
	Contract     *HotpotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HotpotRaw is an auto generated low-level Go binding around an Ethereum contract.
type HotpotRaw struct {
	Contract *Hotpot // Generic contract binding to access the raw methods on
}

// HotpotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HotpotCallerRaw struct {
	Contract *HotpotCaller // Generic read-only contract binding to access the raw methods on
}

// HotpotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HotpotTransactorRaw struct {
	Contract *HotpotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHotpot creates a new instance of Hotpot, bound to a specific deployed contract.
func NewHotpot(address common.Address, backend bind.ContractBackend) (*Hotpot, error) {
	contract, err := bindHotpot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hotpot{HotpotCaller: HotpotCaller{contract: contract}, HotpotTransactor: HotpotTransactor{contract: contract}, HotpotFilterer: HotpotFilterer{contract: contract}}, nil
}

// NewHotpotCaller creates a new read-only instance of Hotpot, bound to a specific deployed contract.
func NewHotpotCaller(address common.Address, caller bind.ContractCaller) (*HotpotCaller, error) {
	contract, err := bindHotpot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HotpotCaller{contract: contract}, nil
}

// NewHotpotTransactor creates a new write-only instance of Hotpot, bound to a specific deployed contract.
func NewHotpotTransactor(address common.Address, transactor bind.ContractTransactor) (*HotpotTransactor, error) {
	contract, err := bindHotpot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HotpotTransactor{contract: contract}, nil
}

// NewHotpotFilterer creates a new log filterer instance of Hotpot, bound to a specific deployed contract.
func NewHotpotFilterer(address common.Address, filterer bind.ContractFilterer) (*HotpotFilterer, error) {
	contract, err := bindHotpot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HotpotFilterer{contract: contract}, nil
}

// bindHotpot binds a generic wrapper to an already deployed contract.
func bindHotpot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HotpotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hotpot *HotpotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hotpot.Contract.HotpotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hotpot *HotpotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hotpot.Contract.HotpotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hotpot *HotpotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hotpot.Contract.HotpotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hotpot *HotpotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hotpot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hotpot *HotpotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hotpot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hotpot *HotpotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hotpot.Contract.contract.Transact(opts, method, params...)
}

// CanClaim is a free data retrieval call binding the contract method 0xbf3506c1.
//
// Solidity: function canClaim(address user) view returns(bool)
func (_Hotpot *HotpotCaller) CanClaim(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "canClaim", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaim is a free data retrieval call binding the contract method 0xbf3506c1.
//
// Solidity: function canClaim(address user) view returns(bool)
func (_Hotpot *HotpotSession) CanClaim(user common.Address) (bool, error) {
	return _Hotpot.Contract.CanClaim(&_Hotpot.CallOpts, user)
}

// CanClaim is a free data retrieval call binding the contract method 0xbf3506c1.
//
// Solidity: function canClaim(address user) view returns(bool)
func (_Hotpot *HotpotCallerSession) CanClaim(user common.Address) (bool, error) {
	return _Hotpot.Contract.CanClaim(&_Hotpot.CallOpts, user)
}

// ChainlinkRequests is a free data retrieval call binding the contract method 0x3f663fb0.
//
// Solidity: function chainlinkRequests(uint256 ) view returns(bool fullfilled, bool exists, uint256 randomWord)
func (_Hotpot *HotpotCaller) ChainlinkRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Fullfilled bool
	Exists     bool
	RandomWord *big.Int
}, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "chainlinkRequests", arg0)

	outstruct := new(struct {
		Fullfilled bool
		Exists     bool
		RandomWord *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fullfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Exists = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.RandomWord = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ChainlinkRequests is a free data retrieval call binding the contract method 0x3f663fb0.
//
// Solidity: function chainlinkRequests(uint256 ) view returns(bool fullfilled, bool exists, uint256 randomWord)
func (_Hotpot *HotpotSession) ChainlinkRequests(arg0 *big.Int) (struct {
	Fullfilled bool
	Exists     bool
	RandomWord *big.Int
}, error) {
	return _Hotpot.Contract.ChainlinkRequests(&_Hotpot.CallOpts, arg0)
}

// ChainlinkRequests is a free data retrieval call binding the contract method 0x3f663fb0.
//
// Solidity: function chainlinkRequests(uint256 ) view returns(bool fullfilled, bool exists, uint256 randomWord)
func (_Hotpot *HotpotCallerSession) ChainlinkRequests(arg0 *big.Int) (struct {
	Fullfilled bool
	Exists     bool
	RandomWord *big.Int
}, error) {
	return _Hotpot.Contract.ChainlinkRequests(&_Hotpot.CallOpts, arg0)
}

// ClaimablePrizes is a free data retrieval call binding the contract method 0xbef5d30c.
//
// Solidity: function claimablePrizes(address ) view returns(uint128 amount, uint128 deadline)
func (_Hotpot *HotpotCaller) ClaimablePrizes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Amount   *big.Int
	Deadline *big.Int
}, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "claimablePrizes", arg0)

	outstruct := new(struct {
		Amount   *big.Int
		Deadline *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ClaimablePrizes is a free data retrieval call binding the contract method 0xbef5d30c.
//
// Solidity: function claimablePrizes(address ) view returns(uint128 amount, uint128 deadline)
func (_Hotpot *HotpotSession) ClaimablePrizes(arg0 common.Address) (struct {
	Amount   *big.Int
	Deadline *big.Int
}, error) {
	return _Hotpot.Contract.ClaimablePrizes(&_Hotpot.CallOpts, arg0)
}

// ClaimablePrizes is a free data retrieval call binding the contract method 0xbef5d30c.
//
// Solidity: function claimablePrizes(address ) view returns(uint128 amount, uint128 deadline)
func (_Hotpot *HotpotCallerSession) ClaimablePrizes(arg0 common.Address) (struct {
	Amount   *big.Int
	Deadline *big.Int
}, error) {
	return _Hotpot.Contract.ClaimablePrizes(&_Hotpot.CallOpts, arg0)
}

// CurrentPotId is a free data retrieval call binding the contract method 0x7b1ff07f.
//
// Solidity: function currentPotId() view returns(uint16)
func (_Hotpot *HotpotCaller) CurrentPotId(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "currentPotId")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// CurrentPotId is a free data retrieval call binding the contract method 0x7b1ff07f.
//
// Solidity: function currentPotId() view returns(uint16)
func (_Hotpot *HotpotSession) CurrentPotId() (uint16, error) {
	return _Hotpot.Contract.CurrentPotId(&_Hotpot.CallOpts)
}

// CurrentPotId is a free data retrieval call binding the contract method 0x7b1ff07f.
//
// Solidity: function currentPotId() view returns(uint16)
func (_Hotpot *HotpotCallerSession) CurrentPotId() (uint16, error) {
	return _Hotpot.Contract.CurrentPotId(&_Hotpot.CallOpts)
}

// CurrentPotSize is a free data retrieval call binding the contract method 0x9476554d.
//
// Solidity: function currentPotSize() view returns(uint256)
func (_Hotpot *HotpotCaller) CurrentPotSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "currentPotSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPotSize is a free data retrieval call binding the contract method 0x9476554d.
//
// Solidity: function currentPotSize() view returns(uint256)
func (_Hotpot *HotpotSession) CurrentPotSize() (*big.Int, error) {
	return _Hotpot.Contract.CurrentPotSize(&_Hotpot.CallOpts)
}

// CurrentPotSize is a free data retrieval call binding the contract method 0x9476554d.
//
// Solidity: function currentPotSize() view returns(uint256)
func (_Hotpot *HotpotCallerSession) CurrentPotSize() (*big.Int, error) {
	return _Hotpot.Contract.CurrentPotSize(&_Hotpot.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint16)
func (_Hotpot *HotpotCaller) Fee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint16)
func (_Hotpot *HotpotSession) Fee() (uint16, error) {
	return _Hotpot.Contract.Fee(&_Hotpot.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint16)
func (_Hotpot *HotpotCallerSession) Fee() (uint16, error) {
	return _Hotpot.Contract.Fee(&_Hotpot.CallOpts)
}

// GetWinningTicketIds is a free data retrieval call binding the contract method 0x69c2d6cf.
//
// Solidity: function getWinningTicketIds(uint16 _potId) view returns(uint32[])
func (_Hotpot *HotpotCaller) GetWinningTicketIds(opts *bind.CallOpts, _potId uint16) ([]uint32, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "getWinningTicketIds", _potId)

	if err != nil {
		return *new([]uint32), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)

	return out0, err

}

// GetWinningTicketIds is a free data retrieval call binding the contract method 0x69c2d6cf.
//
// Solidity: function getWinningTicketIds(uint16 _potId) view returns(uint32[])
func (_Hotpot *HotpotSession) GetWinningTicketIds(_potId uint16) ([]uint32, error) {
	return _Hotpot.Contract.GetWinningTicketIds(&_Hotpot.CallOpts, _potId)
}

// GetWinningTicketIds is a free data retrieval call binding the contract method 0x69c2d6cf.
//
// Solidity: function getWinningTicketIds(uint16 _potId) view returns(uint32[])
func (_Hotpot *HotpotCallerSession) GetWinningTicketIds(_potId uint16) ([]uint32, error) {
	return _Hotpot.Contract.GetWinningTicketIds(&_Hotpot.CallOpts, _potId)
}

// LastRaffleTicketId is a free data retrieval call binding the contract method 0xed62b164.
//
// Solidity: function lastRaffleTicketId() view returns(uint32)
func (_Hotpot *HotpotCaller) LastRaffleTicketId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "lastRaffleTicketId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LastRaffleTicketId is a free data retrieval call binding the contract method 0xed62b164.
//
// Solidity: function lastRaffleTicketId() view returns(uint32)
func (_Hotpot *HotpotSession) LastRaffleTicketId() (uint32, error) {
	return _Hotpot.Contract.LastRaffleTicketId(&_Hotpot.CallOpts)
}

// LastRaffleTicketId is a free data retrieval call binding the contract method 0xed62b164.
//
// Solidity: function lastRaffleTicketId() view returns(uint32)
func (_Hotpot *HotpotCallerSession) LastRaffleTicketId() (uint32, error) {
	return _Hotpot.Contract.LastRaffleTicketId(&_Hotpot.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_Hotpot *HotpotCaller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_Hotpot *HotpotSession) LastRequestId() (*big.Int, error) {
	return _Hotpot.Contract.LastRequestId(&_Hotpot.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_Hotpot *HotpotCallerSession) LastRequestId() (*big.Int, error) {
	return _Hotpot.Contract.LastRequestId(&_Hotpot.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_Hotpot *HotpotCaller) Marketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "marketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_Hotpot *HotpotSession) Marketplace() (common.Address, error) {
	return _Hotpot.Contract.Marketplace(&_Hotpot.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_Hotpot *HotpotCallerSession) Marketplace() (common.Address, error) {
	return _Hotpot.Contract.Marketplace(&_Hotpot.CallOpts)
}

// NextPotTicketIdStart is a free data retrieval call binding the contract method 0x82d6f929.
//
// Solidity: function nextPotTicketIdStart() view returns(uint32)
func (_Hotpot *HotpotCaller) NextPotTicketIdStart(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "nextPotTicketIdStart")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NextPotTicketIdStart is a free data retrieval call binding the contract method 0x82d6f929.
//
// Solidity: function nextPotTicketIdStart() view returns(uint32)
func (_Hotpot *HotpotSession) NextPotTicketIdStart() (uint32, error) {
	return _Hotpot.Contract.NextPotTicketIdStart(&_Hotpot.CallOpts)
}

// NextPotTicketIdStart is a free data retrieval call binding the contract method 0x82d6f929.
//
// Solidity: function nextPotTicketIdStart() view returns(uint32)
func (_Hotpot *HotpotCallerSession) NextPotTicketIdStart() (uint32, error) {
	return _Hotpot.Contract.NextPotTicketIdStart(&_Hotpot.CallOpts)
}

// NumberOfWinners is a free data retrieval call binding the contract method 0x8acfaca9.
//
// Solidity: function numberOfWinners() view returns(uint16)
func (_Hotpot *HotpotCaller) NumberOfWinners(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "numberOfWinners")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// NumberOfWinners is a free data retrieval call binding the contract method 0x8acfaca9.
//
// Solidity: function numberOfWinners() view returns(uint16)
func (_Hotpot *HotpotSession) NumberOfWinners() (uint16, error) {
	return _Hotpot.Contract.NumberOfWinners(&_Hotpot.CallOpts)
}

// NumberOfWinners is a free data retrieval call binding the contract method 0x8acfaca9.
//
// Solidity: function numberOfWinners() view returns(uint16)
func (_Hotpot *HotpotCallerSession) NumberOfWinners() (uint16, error) {
	return _Hotpot.Contract.NumberOfWinners(&_Hotpot.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Hotpot *HotpotCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Hotpot *HotpotSession) Operator() (common.Address, error) {
	return _Hotpot.Contract.Operator(&_Hotpot.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Hotpot *HotpotCallerSession) Operator() (common.Address, error) {
	return _Hotpot.Contract.Operator(&_Hotpot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hotpot *HotpotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hotpot *HotpotSession) Owner() (common.Address, error) {
	return _Hotpot.Contract.Owner(&_Hotpot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hotpot *HotpotCallerSession) Owner() (common.Address, error) {
	return _Hotpot.Contract.Owner(&_Hotpot.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Hotpot *HotpotCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Hotpot *HotpotSession) Paused() (bool, error) {
	return _Hotpot.Contract.Paused(&_Hotpot.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Hotpot *HotpotCallerSession) Paused() (bool, error) {
	return _Hotpot.Contract.Paused(&_Hotpot.CallOpts)
}

// PotLimit is a free data retrieval call binding the contract method 0xd8232f29.
//
// Solidity: function potLimit() view returns(uint256)
func (_Hotpot *HotpotCaller) PotLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "potLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PotLimit is a free data retrieval call binding the contract method 0xd8232f29.
//
// Solidity: function potLimit() view returns(uint256)
func (_Hotpot *HotpotSession) PotLimit() (*big.Int, error) {
	return _Hotpot.Contract.PotLimit(&_Hotpot.CallOpts)
}

// PotLimit is a free data retrieval call binding the contract method 0xd8232f29.
//
// Solidity: function potLimit() view returns(uint256)
func (_Hotpot *HotpotCallerSession) PotLimit() (*big.Int, error) {
	return _Hotpot.Contract.PotLimit(&_Hotpot.CallOpts)
}

// PotTicketIdEnd is a free data retrieval call binding the contract method 0xd785025b.
//
// Solidity: function potTicketIdEnd() view returns(uint32)
func (_Hotpot *HotpotCaller) PotTicketIdEnd(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "potTicketIdEnd")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PotTicketIdEnd is a free data retrieval call binding the contract method 0xd785025b.
//
// Solidity: function potTicketIdEnd() view returns(uint32)
func (_Hotpot *HotpotSession) PotTicketIdEnd() (uint32, error) {
	return _Hotpot.Contract.PotTicketIdEnd(&_Hotpot.CallOpts)
}

// PotTicketIdEnd is a free data retrieval call binding the contract method 0xd785025b.
//
// Solidity: function potTicketIdEnd() view returns(uint32)
func (_Hotpot *HotpotCallerSession) PotTicketIdEnd() (uint32, error) {
	return _Hotpot.Contract.PotTicketIdEnd(&_Hotpot.CallOpts)
}

// PotTicketIdStart is a free data retrieval call binding the contract method 0x1243d370.
//
// Solidity: function potTicketIdStart() view returns(uint32)
func (_Hotpot *HotpotCaller) PotTicketIdStart(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "potTicketIdStart")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// PotTicketIdStart is a free data retrieval call binding the contract method 0x1243d370.
//
// Solidity: function potTicketIdStart() view returns(uint32)
func (_Hotpot *HotpotSession) PotTicketIdStart() (uint32, error) {
	return _Hotpot.Contract.PotTicketIdStart(&_Hotpot.CallOpts)
}

// PotTicketIdStart is a free data retrieval call binding the contract method 0x1243d370.
//
// Solidity: function potTicketIdStart() view returns(uint32)
func (_Hotpot *HotpotCallerSession) PotTicketIdStart() (uint32, error) {
	return _Hotpot.Contract.PotTicketIdStart(&_Hotpot.CallOpts)
}

// PrizeAmounts is a free data retrieval call binding the contract method 0x21cb5b99.
//
// Solidity: function prizeAmounts(uint16 ) view returns(uint128)
func (_Hotpot *HotpotCaller) PrizeAmounts(opts *bind.CallOpts, arg0 uint16) (*big.Int, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "prizeAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizeAmounts is a free data retrieval call binding the contract method 0x21cb5b99.
//
// Solidity: function prizeAmounts(uint16 ) view returns(uint128)
func (_Hotpot *HotpotSession) PrizeAmounts(arg0 uint16) (*big.Int, error) {
	return _Hotpot.Contract.PrizeAmounts(&_Hotpot.CallOpts, arg0)
}

// PrizeAmounts is a free data retrieval call binding the contract method 0x21cb5b99.
//
// Solidity: function prizeAmounts(uint16 ) view returns(uint128)
func (_Hotpot *HotpotCallerSession) PrizeAmounts(arg0 uint16) (*big.Int, error) {
	return _Hotpot.Contract.PrizeAmounts(&_Hotpot.CallOpts, arg0)
}

// RaffleTicketCost is a free data retrieval call binding the contract method 0x898b7dc7.
//
// Solidity: function raffleTicketCost() view returns(uint256)
func (_Hotpot *HotpotCaller) RaffleTicketCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "raffleTicketCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaffleTicketCost is a free data retrieval call binding the contract method 0x898b7dc7.
//
// Solidity: function raffleTicketCost() view returns(uint256)
func (_Hotpot *HotpotSession) RaffleTicketCost() (*big.Int, error) {
	return _Hotpot.Contract.RaffleTicketCost(&_Hotpot.CallOpts)
}

// RaffleTicketCost is a free data retrieval call binding the contract method 0x898b7dc7.
//
// Solidity: function raffleTicketCost() view returns(uint256)
func (_Hotpot *HotpotCallerSession) RaffleTicketCost() (*big.Int, error) {
	return _Hotpot.Contract.RaffleTicketCost(&_Hotpot.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_Hotpot *HotpotCaller) RequestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "requestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_Hotpot *HotpotSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _Hotpot.Contract.RequestIds(&_Hotpot.CallOpts, arg0)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_Hotpot *HotpotCallerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _Hotpot.Contract.RequestIds(&_Hotpot.CallOpts, arg0)
}

// TradeFee is a free data retrieval call binding the contract method 0x24bcdfbd.
//
// Solidity: function tradeFee() view returns(uint16)
func (_Hotpot *HotpotCaller) TradeFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "tradeFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TradeFee is a free data retrieval call binding the contract method 0x24bcdfbd.
//
// Solidity: function tradeFee() view returns(uint16)
func (_Hotpot *HotpotSession) TradeFee() (uint16, error) {
	return _Hotpot.Contract.TradeFee(&_Hotpot.CallOpts)
}

// TradeFee is a free data retrieval call binding the contract method 0x24bcdfbd.
//
// Solidity: function tradeFee() view returns(uint16)
func (_Hotpot *HotpotCallerSession) TradeFee() (uint16, error) {
	return _Hotpot.Contract.TradeFee(&_Hotpot.CallOpts)
}

// WinningTicketIds is a free data retrieval call binding the contract method 0xda5b6e2f.
//
// Solidity: function winningTicketIds(uint16 , uint256 ) view returns(uint32)
func (_Hotpot *HotpotCaller) WinningTicketIds(opts *bind.CallOpts, arg0 uint16, arg1 *big.Int) (uint32, error) {
	var out []interface{}
	err := _Hotpot.contract.Call(opts, &out, "winningTicketIds", arg0, arg1)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// WinningTicketIds is a free data retrieval call binding the contract method 0xda5b6e2f.
//
// Solidity: function winningTicketIds(uint16 , uint256 ) view returns(uint32)
func (_Hotpot *HotpotSession) WinningTicketIds(arg0 uint16, arg1 *big.Int) (uint32, error) {
	return _Hotpot.Contract.WinningTicketIds(&_Hotpot.CallOpts, arg0, arg1)
}

// WinningTicketIds is a free data retrieval call binding the contract method 0xda5b6e2f.
//
// Solidity: function winningTicketIds(uint16 , uint256 ) view returns(uint32)
func (_Hotpot *HotpotCallerSession) WinningTicketIds(arg0 uint16, arg1 *big.Int) (uint32, error) {
	return _Hotpot.Contract.WinningTicketIds(&_Hotpot.CallOpts, arg0, arg1)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Hotpot *HotpotTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Hotpot *HotpotSession) Claim() (*types.Transaction, error) {
	return _Hotpot.Contract.Claim(&_Hotpot.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Hotpot *HotpotTransactorSession) Claim() (*types.Transaction, error) {
	return _Hotpot.Contract.Claim(&_Hotpot.TransactOpts)
}

// ExecuteRaffle is a paid mutator transaction binding the contract method 0xce4b7633.
//
// Solidity: function executeRaffle(address[] _winners) returns()
func (_Hotpot *HotpotTransactor) ExecuteRaffle(opts *bind.TransactOpts, _winners []common.Address) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "executeRaffle", _winners)
}

// ExecuteRaffle is a paid mutator transaction binding the contract method 0xce4b7633.
//
// Solidity: function executeRaffle(address[] _winners) returns()
func (_Hotpot *HotpotSession) ExecuteRaffle(_winners []common.Address) (*types.Transaction, error) {
	return _Hotpot.Contract.ExecuteRaffle(&_Hotpot.TransactOpts, _winners)
}

// ExecuteRaffle is a paid mutator transaction binding the contract method 0xce4b7633.
//
// Solidity: function executeRaffle(address[] _winners) returns()
func (_Hotpot *HotpotTransactorSession) ExecuteRaffle(_winners []common.Address) (*types.Transaction, error) {
	return _Hotpot.Contract.ExecuteRaffle(&_Hotpot.TransactOpts, _winners)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xa449781b.
//
// Solidity: function executeTrade(uint256 _amountInWei, address _buyer, address _seller, uint256 _buyerPendingAmount, uint256 _sellerPendingAmount) payable returns()
func (_Hotpot *HotpotTransactor) ExecuteTrade(opts *bind.TransactOpts, _amountInWei *big.Int, _buyer common.Address, _seller common.Address, _buyerPendingAmount *big.Int, _sellerPendingAmount *big.Int) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "executeTrade", _amountInWei, _buyer, _seller, _buyerPendingAmount, _sellerPendingAmount)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xa449781b.
//
// Solidity: function executeTrade(uint256 _amountInWei, address _buyer, address _seller, uint256 _buyerPendingAmount, uint256 _sellerPendingAmount) payable returns()
func (_Hotpot *HotpotSession) ExecuteTrade(_amountInWei *big.Int, _buyer common.Address, _seller common.Address, _buyerPendingAmount *big.Int, _sellerPendingAmount *big.Int) (*types.Transaction, error) {
	return _Hotpot.Contract.ExecuteTrade(&_Hotpot.TransactOpts, _amountInWei, _buyer, _seller, _buyerPendingAmount, _sellerPendingAmount)
}

// ExecuteTrade is a paid mutator transaction binding the contract method 0xa449781b.
//
// Solidity: function executeTrade(uint256 _amountInWei, address _buyer, address _seller, uint256 _buyerPendingAmount, uint256 _sellerPendingAmount) payable returns()
func (_Hotpot *HotpotTransactorSession) ExecuteTrade(_amountInWei *big.Int, _buyer common.Address, _seller common.Address, _buyerPendingAmount *big.Int, _sellerPendingAmount *big.Int) (*types.Transaction, error) {
	return _Hotpot.Contract.ExecuteTrade(&_Hotpot.TransactOpts, _amountInWei, _buyer, _seller, _buyerPendingAmount, _sellerPendingAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf5ec2818.
//
// Solidity: function initialize(address _owner, (uint256,uint256,uint128,uint16,uint16,uint16,address,address) params) returns()
func (_Hotpot *HotpotTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, params IHotpotInitializeParams) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "initialize", _owner, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xf5ec2818.
//
// Solidity: function initialize(address _owner, (uint256,uint256,uint128,uint16,uint16,uint16,address,address) params) returns()
func (_Hotpot *HotpotSession) Initialize(_owner common.Address, params IHotpotInitializeParams) (*types.Transaction, error) {
	return _Hotpot.Contract.Initialize(&_Hotpot.TransactOpts, _owner, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xf5ec2818.
//
// Solidity: function initialize(address _owner, (uint256,uint256,uint128,uint16,uint16,uint16,address,address) params) returns()
func (_Hotpot *HotpotTransactorSession) Initialize(_owner common.Address, params IHotpotInitializeParams) (*types.Transaction, error) {
	return _Hotpot.Contract.Initialize(&_Hotpot.TransactOpts, _owner, params)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_Hotpot *HotpotTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "rawFulfillRandomWords", _requestId, _randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_Hotpot *HotpotSession) RawFulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _Hotpot.Contract.RawFulfillRandomWords(&_Hotpot.TransactOpts, _requestId, _randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_Hotpot *HotpotTransactorSession) RawFulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _Hotpot.Contract.RawFulfillRandomWords(&_Hotpot.TransactOpts, _requestId, _randomWords)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hotpot *HotpotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hotpot *HotpotSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hotpot.Contract.RenounceOwnership(&_Hotpot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hotpot *HotpotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hotpot.Contract.RenounceOwnership(&_Hotpot.TransactOpts)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address _newMarketplace) returns()
func (_Hotpot *HotpotTransactor) SetMarketplace(opts *bind.TransactOpts, _newMarketplace common.Address) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "setMarketplace", _newMarketplace)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address _newMarketplace) returns()
func (_Hotpot *HotpotSession) SetMarketplace(_newMarketplace common.Address) (*types.Transaction, error) {
	return _Hotpot.Contract.SetMarketplace(&_Hotpot.TransactOpts, _newMarketplace)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address _newMarketplace) returns()
func (_Hotpot *HotpotTransactorSession) SetMarketplace(_newMarketplace common.Address) (*types.Transaction, error) {
	return _Hotpot.Contract.SetMarketplace(&_Hotpot.TransactOpts, _newMarketplace)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_Hotpot *HotpotTransactor) SetOperator(opts *bind.TransactOpts, _newOperator common.Address) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "setOperator", _newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_Hotpot *HotpotSession) SetOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _Hotpot.Contract.SetOperator(&_Hotpot.TransactOpts, _newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_Hotpot *HotpotTransactorSession) SetOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _Hotpot.Contract.SetOperator(&_Hotpot.TransactOpts, _newOperator)
}

// SetPotLimit is a paid mutator transaction binding the contract method 0x03a6fd98.
//
// Solidity: function setPotLimit(uint256 _newPotLimit) returns()
func (_Hotpot *HotpotTransactor) SetPotLimit(opts *bind.TransactOpts, _newPotLimit *big.Int) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "setPotLimit", _newPotLimit)
}

// SetPotLimit is a paid mutator transaction binding the contract method 0x03a6fd98.
//
// Solidity: function setPotLimit(uint256 _newPotLimit) returns()
func (_Hotpot *HotpotSession) SetPotLimit(_newPotLimit *big.Int) (*types.Transaction, error) {
	return _Hotpot.Contract.SetPotLimit(&_Hotpot.TransactOpts, _newPotLimit)
}

// SetPotLimit is a paid mutator transaction binding the contract method 0x03a6fd98.
//
// Solidity: function setPotLimit(uint256 _newPotLimit) returns()
func (_Hotpot *HotpotTransactorSession) SetPotLimit(_newPotLimit *big.Int) (*types.Transaction, error) {
	return _Hotpot.Contract.SetPotLimit(&_Hotpot.TransactOpts, _newPotLimit)
}

// SetRaffleTicketCost is a paid mutator transaction binding the contract method 0x122e0d22.
//
// Solidity: function setRaffleTicketCost(uint256 _newRaffleTicketCost) returns()
func (_Hotpot *HotpotTransactor) SetRaffleTicketCost(opts *bind.TransactOpts, _newRaffleTicketCost *big.Int) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "setRaffleTicketCost", _newRaffleTicketCost)
}

// SetRaffleTicketCost is a paid mutator transaction binding the contract method 0x122e0d22.
//
// Solidity: function setRaffleTicketCost(uint256 _newRaffleTicketCost) returns()
func (_Hotpot *HotpotSession) SetRaffleTicketCost(_newRaffleTicketCost *big.Int) (*types.Transaction, error) {
	return _Hotpot.Contract.SetRaffleTicketCost(&_Hotpot.TransactOpts, _newRaffleTicketCost)
}

// SetRaffleTicketCost is a paid mutator transaction binding the contract method 0x122e0d22.
//
// Solidity: function setRaffleTicketCost(uint256 _newRaffleTicketCost) returns()
func (_Hotpot *HotpotTransactorSession) SetRaffleTicketCost(_newRaffleTicketCost *big.Int) (*types.Transaction, error) {
	return _Hotpot.Contract.SetRaffleTicketCost(&_Hotpot.TransactOpts, _newRaffleTicketCost)
}

// SetTradeFee is a paid mutator transaction binding the contract method 0x9fdccfb8.
//
// Solidity: function setTradeFee(uint16 _newTradeFee) returns()
func (_Hotpot *HotpotTransactor) SetTradeFee(opts *bind.TransactOpts, _newTradeFee uint16) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "setTradeFee", _newTradeFee)
}

// SetTradeFee is a paid mutator transaction binding the contract method 0x9fdccfb8.
//
// Solidity: function setTradeFee(uint16 _newTradeFee) returns()
func (_Hotpot *HotpotSession) SetTradeFee(_newTradeFee uint16) (*types.Transaction, error) {
	return _Hotpot.Contract.SetTradeFee(&_Hotpot.TransactOpts, _newTradeFee)
}

// SetTradeFee is a paid mutator transaction binding the contract method 0x9fdccfb8.
//
// Solidity: function setTradeFee(uint16 _newTradeFee) returns()
func (_Hotpot *HotpotTransactorSession) SetTradeFee(_newTradeFee uint16) (*types.Transaction, error) {
	return _Hotpot.Contract.SetTradeFee(&_Hotpot.TransactOpts, _newTradeFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hotpot *HotpotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hotpot *HotpotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hotpot.Contract.TransferOwnership(&_Hotpot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hotpot *HotpotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hotpot.Contract.TransferOwnership(&_Hotpot.TransactOpts, newOwner)
}

// UpdateNumberOfWinners is a paid mutator transaction binding the contract method 0x94d09761.
//
// Solidity: function updateNumberOfWinners(uint16 _nOfWinners) returns()
func (_Hotpot *HotpotTransactor) UpdateNumberOfWinners(opts *bind.TransactOpts, _nOfWinners uint16) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "updateNumberOfWinners", _nOfWinners)
}

// UpdateNumberOfWinners is a paid mutator transaction binding the contract method 0x94d09761.
//
// Solidity: function updateNumberOfWinners(uint16 _nOfWinners) returns()
func (_Hotpot *HotpotSession) UpdateNumberOfWinners(_nOfWinners uint16) (*types.Transaction, error) {
	return _Hotpot.Contract.UpdateNumberOfWinners(&_Hotpot.TransactOpts, _nOfWinners)
}

// UpdateNumberOfWinners is a paid mutator transaction binding the contract method 0x94d09761.
//
// Solidity: function updateNumberOfWinners(uint16 _nOfWinners) returns()
func (_Hotpot *HotpotTransactorSession) UpdateNumberOfWinners(_nOfWinners uint16) (*types.Transaction, error) {
	return _Hotpot.Contract.UpdateNumberOfWinners(&_Hotpot.TransactOpts, _nOfWinners)
}

// UpdatePrizeAmounts is a paid mutator transaction binding the contract method 0x550da497.
//
// Solidity: function updatePrizeAmounts(uint128[] _newPrizeAmounts) returns()
func (_Hotpot *HotpotTransactor) UpdatePrizeAmounts(opts *bind.TransactOpts, _newPrizeAmounts []*big.Int) (*types.Transaction, error) {
	return _Hotpot.contract.Transact(opts, "updatePrizeAmounts", _newPrizeAmounts)
}

// UpdatePrizeAmounts is a paid mutator transaction binding the contract method 0x550da497.
//
// Solidity: function updatePrizeAmounts(uint128[] _newPrizeAmounts) returns()
func (_Hotpot *HotpotSession) UpdatePrizeAmounts(_newPrizeAmounts []*big.Int) (*types.Transaction, error) {
	return _Hotpot.Contract.UpdatePrizeAmounts(&_Hotpot.TransactOpts, _newPrizeAmounts)
}

// UpdatePrizeAmounts is a paid mutator transaction binding the contract method 0x550da497.
//
// Solidity: function updatePrizeAmounts(uint128[] _newPrizeAmounts) returns()
func (_Hotpot *HotpotTransactorSession) UpdatePrizeAmounts(_newPrizeAmounts []*big.Int) (*types.Transaction, error) {
	return _Hotpot.Contract.UpdatePrizeAmounts(&_Hotpot.TransactOpts, _newPrizeAmounts)
}

// HotpotClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Hotpot contract.
type HotpotClaimIterator struct {
	Event *HotpotClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotClaim represents a Claim event raised by the Hotpot contract.
type HotpotClaim struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_Hotpot *HotpotFilterer) FilterClaim(opts *bind.FilterOpts, user []common.Address) (*HotpotClaimIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return &HotpotClaimIterator{contract: _Hotpot.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_Hotpot *HotpotFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *HotpotClaim, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotClaim)
				if err := _Hotpot.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed user, uint256 amount)
func (_Hotpot *HotpotFilterer) ParseClaim(log types.Log) (*HotpotClaim, error) {
	event := new(HotpotClaim)
	if err := _Hotpot.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotGenerateRaffleTicketsIterator is returned from FilterGenerateRaffleTickets and is used to iterate over the raw logs and unpacked data for GenerateRaffleTickets events raised by the Hotpot contract.
type HotpotGenerateRaffleTicketsIterator struct {
	Event *HotpotGenerateRaffleTickets // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotGenerateRaffleTicketsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotGenerateRaffleTickets)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotGenerateRaffleTickets)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotGenerateRaffleTicketsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotGenerateRaffleTicketsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotGenerateRaffleTickets represents a GenerateRaffleTickets event raised by the Hotpot contract.
type HotpotGenerateRaffleTickets struct {
	Buyer               common.Address
	Seller              common.Address
	BuyerTicketIdStart  uint32
	BuyerTicketIdEnd    uint32
	SellerTicketIdStart uint32
	SellerTicketIdEnd   uint32
	BuyerPendingAmount  *big.Int
	SellerPendingAmount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterGenerateRaffleTickets is a free log retrieval operation binding the contract event 0x9f36f71097b819d5a174fcd2024f3a286ba1ab3a530d8f992da43f26caa35145.
//
// Solidity: event GenerateRaffleTickets(address indexed _buyer, address indexed _seller, uint32 _buyerTicketIdStart, uint32 _buyerTicketIdEnd, uint32 _sellerTicketIdStart, uint32 _sellerTicketIdEnd, uint256 _buyerPendingAmount, uint256 _sellerPendingAmount)
func (_Hotpot *HotpotFilterer) FilterGenerateRaffleTickets(opts *bind.FilterOpts, _buyer []common.Address, _seller []common.Address) (*HotpotGenerateRaffleTicketsIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "GenerateRaffleTickets", _buyerRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return &HotpotGenerateRaffleTicketsIterator{contract: _Hotpot.contract, event: "GenerateRaffleTickets", logs: logs, sub: sub}, nil
}

// WatchGenerateRaffleTickets is a free log subscription operation binding the contract event 0x9f36f71097b819d5a174fcd2024f3a286ba1ab3a530d8f992da43f26caa35145.
//
// Solidity: event GenerateRaffleTickets(address indexed _buyer, address indexed _seller, uint32 _buyerTicketIdStart, uint32 _buyerTicketIdEnd, uint32 _sellerTicketIdStart, uint32 _sellerTicketIdEnd, uint256 _buyerPendingAmount, uint256 _sellerPendingAmount)
func (_Hotpot *HotpotFilterer) WatchGenerateRaffleTickets(opts *bind.WatchOpts, sink chan<- *HotpotGenerateRaffleTickets, _buyer []common.Address, _seller []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "GenerateRaffleTickets", _buyerRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotGenerateRaffleTickets)
				if err := _Hotpot.contract.UnpackLog(event, "GenerateRaffleTickets", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGenerateRaffleTickets is a log parse operation binding the contract event 0x9f36f71097b819d5a174fcd2024f3a286ba1ab3a530d8f992da43f26caa35145.
//
// Solidity: event GenerateRaffleTickets(address indexed _buyer, address indexed _seller, uint32 _buyerTicketIdStart, uint32 _buyerTicketIdEnd, uint32 _sellerTicketIdStart, uint32 _sellerTicketIdEnd, uint256 _buyerPendingAmount, uint256 _sellerPendingAmount)
func (_Hotpot *HotpotFilterer) ParseGenerateRaffleTickets(log types.Log) (*HotpotGenerateRaffleTickets, error) {
	event := new(HotpotGenerateRaffleTickets)
	if err := _Hotpot.contract.UnpackLog(event, "GenerateRaffleTickets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Hotpot contract.
type HotpotInitializedIterator struct {
	Event *HotpotInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotInitialized represents a Initialized event raised by the Hotpot contract.
type HotpotInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Hotpot *HotpotFilterer) FilterInitialized(opts *bind.FilterOpts) (*HotpotInitializedIterator, error) {

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &HotpotInitializedIterator{contract: _Hotpot.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Hotpot *HotpotFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *HotpotInitialized) (event.Subscription, error) {

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotInitialized)
				if err := _Hotpot.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Hotpot *HotpotFilterer) ParseInitialized(log types.Log) (*HotpotInitialized, error) {
	event := new(HotpotInitialized)
	if err := _Hotpot.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotMarketplaceUpdatedIterator is returned from FilterMarketplaceUpdated and is used to iterate over the raw logs and unpacked data for MarketplaceUpdated events raised by the Hotpot contract.
type HotpotMarketplaceUpdatedIterator struct {
	Event *HotpotMarketplaceUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotMarketplaceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotMarketplaceUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotMarketplaceUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotMarketplaceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotMarketplaceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotMarketplaceUpdated represents a MarketplaceUpdated event raised by the Hotpot contract.
type HotpotMarketplaceUpdated struct {
	NewMarketplace common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMarketplaceUpdated is a free log retrieval operation binding the contract event 0x210690abd7fd6cdbb8f2beb202b2a253d58d7a0813b2175c4172c14c0c1af6dc.
//
// Solidity: event MarketplaceUpdated(address _newMarketplace)
func (_Hotpot *HotpotFilterer) FilterMarketplaceUpdated(opts *bind.FilterOpts) (*HotpotMarketplaceUpdatedIterator, error) {

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "MarketplaceUpdated")
	if err != nil {
		return nil, err
	}
	return &HotpotMarketplaceUpdatedIterator{contract: _Hotpot.contract, event: "MarketplaceUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketplaceUpdated is a free log subscription operation binding the contract event 0x210690abd7fd6cdbb8f2beb202b2a253d58d7a0813b2175c4172c14c0c1af6dc.
//
// Solidity: event MarketplaceUpdated(address _newMarketplace)
func (_Hotpot *HotpotFilterer) WatchMarketplaceUpdated(opts *bind.WatchOpts, sink chan<- *HotpotMarketplaceUpdated) (event.Subscription, error) {

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "MarketplaceUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotMarketplaceUpdated)
				if err := _Hotpot.contract.UnpackLog(event, "MarketplaceUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketplaceUpdated is a log parse operation binding the contract event 0x210690abd7fd6cdbb8f2beb202b2a253d58d7a0813b2175c4172c14c0c1af6dc.
//
// Solidity: event MarketplaceUpdated(address _newMarketplace)
func (_Hotpot *HotpotFilterer) ParseMarketplaceUpdated(log types.Log) (*HotpotMarketplaceUpdated, error) {
	event := new(HotpotMarketplaceUpdated)
	if err := _Hotpot.contract.UnpackLog(event, "MarketplaceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotNumberOfWinnersUpdatedIterator is returned from FilterNumberOfWinnersUpdated and is used to iterate over the raw logs and unpacked data for NumberOfWinnersUpdated events raised by the Hotpot contract.
type HotpotNumberOfWinnersUpdatedIterator struct {
	Event *HotpotNumberOfWinnersUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotNumberOfWinnersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotNumberOfWinnersUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotNumberOfWinnersUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotNumberOfWinnersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotNumberOfWinnersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotNumberOfWinnersUpdated represents a NumberOfWinnersUpdated event raised by the Hotpot contract.
type HotpotNumberOfWinnersUpdated struct {
	NOfWinners uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNumberOfWinnersUpdated is a free log retrieval operation binding the contract event 0xaf655bb590a9a783519c5319a7f42e07720e729de4a0e85c3f92165bc2ffef83.
//
// Solidity: event NumberOfWinnersUpdated(uint16 _nOfWinners)
func (_Hotpot *HotpotFilterer) FilterNumberOfWinnersUpdated(opts *bind.FilterOpts) (*HotpotNumberOfWinnersUpdatedIterator, error) {

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "NumberOfWinnersUpdated")
	if err != nil {
		return nil, err
	}
	return &HotpotNumberOfWinnersUpdatedIterator{contract: _Hotpot.contract, event: "NumberOfWinnersUpdated", logs: logs, sub: sub}, nil
}

// WatchNumberOfWinnersUpdated is a free log subscription operation binding the contract event 0xaf655bb590a9a783519c5319a7f42e07720e729de4a0e85c3f92165bc2ffef83.
//
// Solidity: event NumberOfWinnersUpdated(uint16 _nOfWinners)
func (_Hotpot *HotpotFilterer) WatchNumberOfWinnersUpdated(opts *bind.WatchOpts, sink chan<- *HotpotNumberOfWinnersUpdated) (event.Subscription, error) {

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "NumberOfWinnersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotNumberOfWinnersUpdated)
				if err := _Hotpot.contract.UnpackLog(event, "NumberOfWinnersUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNumberOfWinnersUpdated is a log parse operation binding the contract event 0xaf655bb590a9a783519c5319a7f42e07720e729de4a0e85c3f92165bc2ffef83.
//
// Solidity: event NumberOfWinnersUpdated(uint16 _nOfWinners)
func (_Hotpot *HotpotFilterer) ParseNumberOfWinnersUpdated(log types.Log) (*HotpotNumberOfWinnersUpdated, error) {
	event := new(HotpotNumberOfWinnersUpdated)
	if err := _Hotpot.contract.UnpackLog(event, "NumberOfWinnersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotOperatorUpdatedIterator is returned from FilterOperatorUpdated and is used to iterate over the raw logs and unpacked data for OperatorUpdated events raised by the Hotpot contract.
type HotpotOperatorUpdatedIterator struct {
	Event *HotpotOperatorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotOperatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotOperatorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotOperatorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotOperatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotOperatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotOperatorUpdated represents a OperatorUpdated event raised by the Hotpot contract.
type HotpotOperatorUpdated struct {
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorUpdated is a free log retrieval operation binding the contract event 0xb3b3f5f64ab192e4b5fefde1f51ce9733bbdcf831951543b325aebd49cc27ec4.
//
// Solidity: event OperatorUpdated(address _newOperator)
func (_Hotpot *HotpotFilterer) FilterOperatorUpdated(opts *bind.FilterOpts) (*HotpotOperatorUpdatedIterator, error) {

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "OperatorUpdated")
	if err != nil {
		return nil, err
	}
	return &HotpotOperatorUpdatedIterator{contract: _Hotpot.contract, event: "OperatorUpdated", logs: logs, sub: sub}, nil
}

// WatchOperatorUpdated is a free log subscription operation binding the contract event 0xb3b3f5f64ab192e4b5fefde1f51ce9733bbdcf831951543b325aebd49cc27ec4.
//
// Solidity: event OperatorUpdated(address _newOperator)
func (_Hotpot *HotpotFilterer) WatchOperatorUpdated(opts *bind.WatchOpts, sink chan<- *HotpotOperatorUpdated) (event.Subscription, error) {

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "OperatorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotOperatorUpdated)
				if err := _Hotpot.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOperatorUpdated is a log parse operation binding the contract event 0xb3b3f5f64ab192e4b5fefde1f51ce9733bbdcf831951543b325aebd49cc27ec4.
//
// Solidity: event OperatorUpdated(address _newOperator)
func (_Hotpot *HotpotFilterer) ParseOperatorUpdated(log types.Log) (*HotpotOperatorUpdated, error) {
	event := new(HotpotOperatorUpdated)
	if err := _Hotpot.contract.UnpackLog(event, "OperatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hotpot contract.
type HotpotOwnershipTransferredIterator struct {
	Event *HotpotOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotOwnershipTransferred represents a OwnershipTransferred event raised by the Hotpot contract.
type HotpotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hotpot *HotpotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HotpotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HotpotOwnershipTransferredIterator{contract: _Hotpot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hotpot *HotpotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HotpotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotOwnershipTransferred)
				if err := _Hotpot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hotpot *HotpotFilterer) ParseOwnershipTransferred(log types.Log) (*HotpotOwnershipTransferred, error) {
	event := new(HotpotOwnershipTransferred)
	if err := _Hotpot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Hotpot contract.
type HotpotPausedIterator struct {
	Event *HotpotPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotPaused represents a Paused event raised by the Hotpot contract.
type HotpotPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Hotpot *HotpotFilterer) FilterPaused(opts *bind.FilterOpts) (*HotpotPausedIterator, error) {

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HotpotPausedIterator{contract: _Hotpot.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Hotpot *HotpotFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HotpotPaused) (event.Subscription, error) {

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotPaused)
				if err := _Hotpot.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Hotpot *HotpotFilterer) ParsePaused(log types.Log) (*HotpotPaused, error) {
	event := new(HotpotPaused)
	if err := _Hotpot.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotPrizeAmountsUpdatedIterator is returned from FilterPrizeAmountsUpdated and is used to iterate over the raw logs and unpacked data for PrizeAmountsUpdated events raised by the Hotpot contract.
type HotpotPrizeAmountsUpdatedIterator struct {
	Event *HotpotPrizeAmountsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotPrizeAmountsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotPrizeAmountsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotPrizeAmountsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotPrizeAmountsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotPrizeAmountsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotPrizeAmountsUpdated represents a PrizeAmountsUpdated event raised by the Hotpot contract.
type HotpotPrizeAmountsUpdated struct {
	NewPrizeAmounts []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPrizeAmountsUpdated is a free log retrieval operation binding the contract event 0x7b448bf1cdbff04f2b0fa7bda6c864f8670ce866f64332d09288e209bb9aa4c8.
//
// Solidity: event PrizeAmountsUpdated(uint128[] _newPrizeAmounts)
func (_Hotpot *HotpotFilterer) FilterPrizeAmountsUpdated(opts *bind.FilterOpts) (*HotpotPrizeAmountsUpdatedIterator, error) {

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "PrizeAmountsUpdated")
	if err != nil {
		return nil, err
	}
	return &HotpotPrizeAmountsUpdatedIterator{contract: _Hotpot.contract, event: "PrizeAmountsUpdated", logs: logs, sub: sub}, nil
}

// WatchPrizeAmountsUpdated is a free log subscription operation binding the contract event 0x7b448bf1cdbff04f2b0fa7bda6c864f8670ce866f64332d09288e209bb9aa4c8.
//
// Solidity: event PrizeAmountsUpdated(uint128[] _newPrizeAmounts)
func (_Hotpot *HotpotFilterer) WatchPrizeAmountsUpdated(opts *bind.WatchOpts, sink chan<- *HotpotPrizeAmountsUpdated) (event.Subscription, error) {

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "PrizeAmountsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotPrizeAmountsUpdated)
				if err := _Hotpot.contract.UnpackLog(event, "PrizeAmountsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePrizeAmountsUpdated is a log parse operation binding the contract event 0x7b448bf1cdbff04f2b0fa7bda6c864f8670ce866f64332d09288e209bb9aa4c8.
//
// Solidity: event PrizeAmountsUpdated(uint128[] _newPrizeAmounts)
func (_Hotpot *HotpotFilterer) ParsePrizeAmountsUpdated(log types.Log) (*HotpotPrizeAmountsUpdated, error) {
	event := new(HotpotPrizeAmountsUpdated)
	if err := _Hotpot.contract.UnpackLog(event, "PrizeAmountsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotRandomWordRequestedIterator is returned from FilterRandomWordRequested and is used to iterate over the raw logs and unpacked data for RandomWordRequested events raised by the Hotpot contract.
type HotpotRandomWordRequestedIterator struct {
	Event *HotpotRandomWordRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotRandomWordRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotRandomWordRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotRandomWordRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotRandomWordRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotRandomWordRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotRandomWordRequested represents a RandomWordRequested event raised by the Hotpot contract.
type HotpotRandomWordRequested struct {
	RequestId    *big.Int
	FromTicketId uint32
	ToTicketId   uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRandomWordRequested is a free log retrieval operation binding the contract event 0xdca7460aac1cfbfb8c91e4c853980bb7cd6fc8667e491139a8719fa8af124882.
//
// Solidity: event RandomWordRequested(uint256 requestId, uint32 fromTicketId, uint32 toTicketId)
func (_Hotpot *HotpotFilterer) FilterRandomWordRequested(opts *bind.FilterOpts) (*HotpotRandomWordRequestedIterator, error) {

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "RandomWordRequested")
	if err != nil {
		return nil, err
	}
	return &HotpotRandomWordRequestedIterator{contract: _Hotpot.contract, event: "RandomWordRequested", logs: logs, sub: sub}, nil
}

// WatchRandomWordRequested is a free log subscription operation binding the contract event 0xdca7460aac1cfbfb8c91e4c853980bb7cd6fc8667e491139a8719fa8af124882.
//
// Solidity: event RandomWordRequested(uint256 requestId, uint32 fromTicketId, uint32 toTicketId)
func (_Hotpot *HotpotFilterer) WatchRandomWordRequested(opts *bind.WatchOpts, sink chan<- *HotpotRandomWordRequested) (event.Subscription, error) {

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "RandomWordRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotRandomWordRequested)
				if err := _Hotpot.contract.UnpackLog(event, "RandomWordRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomWordRequested is a log parse operation binding the contract event 0xdca7460aac1cfbfb8c91e4c853980bb7cd6fc8667e491139a8719fa8af124882.
//
// Solidity: event RandomWordRequested(uint256 requestId, uint32 fromTicketId, uint32 toTicketId)
func (_Hotpot *HotpotFilterer) ParseRandomWordRequested(log types.Log) (*HotpotRandomWordRequested, error) {
	event := new(HotpotRandomWordRequested)
	if err := _Hotpot.contract.UnpackLog(event, "RandomWordRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotRandomnessFulfilledIterator is returned from FilterRandomnessFulfilled and is used to iterate over the raw logs and unpacked data for RandomnessFulfilled events raised by the Hotpot contract.
type HotpotRandomnessFulfilledIterator struct {
	Event *HotpotRandomnessFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotRandomnessFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotRandomnessFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotRandomnessFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotRandomnessFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotRandomnessFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotRandomnessFulfilled represents a RandomnessFulfilled event raised by the Hotpot contract.
type HotpotRandomnessFulfilled struct {
	PotId      uint16
	RandomWord *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRandomnessFulfilled is a free log retrieval operation binding the contract event 0xbca09660beb78b208cca8d25782c2c0e5b09282d7002e566ff7a1f5fe19c4742.
//
// Solidity: event RandomnessFulfilled(uint16 indexed potId, uint256 randomWord)
func (_Hotpot *HotpotFilterer) FilterRandomnessFulfilled(opts *bind.FilterOpts, potId []uint16) (*HotpotRandomnessFulfilledIterator, error) {

	var potIdRule []interface{}
	for _, potIdItem := range potId {
		potIdRule = append(potIdRule, potIdItem)
	}

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "RandomnessFulfilled", potIdRule)
	if err != nil {
		return nil, err
	}
	return &HotpotRandomnessFulfilledIterator{contract: _Hotpot.contract, event: "RandomnessFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomnessFulfilled is a free log subscription operation binding the contract event 0xbca09660beb78b208cca8d25782c2c0e5b09282d7002e566ff7a1f5fe19c4742.
//
// Solidity: event RandomnessFulfilled(uint16 indexed potId, uint256 randomWord)
func (_Hotpot *HotpotFilterer) WatchRandomnessFulfilled(opts *bind.WatchOpts, sink chan<- *HotpotRandomnessFulfilled, potId []uint16) (event.Subscription, error) {

	var potIdRule []interface{}
	for _, potIdItem := range potId {
		potIdRule = append(potIdRule, potIdItem)
	}

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "RandomnessFulfilled", potIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotRandomnessFulfilled)
				if err := _Hotpot.contract.UnpackLog(event, "RandomnessFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomnessFulfilled is a log parse operation binding the contract event 0xbca09660beb78b208cca8d25782c2c0e5b09282d7002e566ff7a1f5fe19c4742.
//
// Solidity: event RandomnessFulfilled(uint16 indexed potId, uint256 randomWord)
func (_Hotpot *HotpotFilterer) ParseRandomnessFulfilled(log types.Log) (*HotpotRandomnessFulfilled, error) {
	event := new(HotpotRandomnessFulfilled)
	if err := _Hotpot.contract.UnpackLog(event, "RandomnessFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Hotpot contract.
type HotpotUnpausedIterator struct {
	Event *HotpotUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotUnpaused represents a Unpaused event raised by the Hotpot contract.
type HotpotUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Hotpot *HotpotFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HotpotUnpausedIterator, error) {

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HotpotUnpausedIterator{contract: _Hotpot.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Hotpot *HotpotFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HotpotUnpaused) (event.Subscription, error) {

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotUnpaused)
				if err := _Hotpot.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Hotpot *HotpotFilterer) ParseUnpaused(log types.Log) (*HotpotUnpaused, error) {
	event := new(HotpotUnpaused)
	if err := _Hotpot.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HotpotWinnersAssignedIterator is returned from FilterWinnersAssigned and is used to iterate over the raw logs and unpacked data for WinnersAssigned events raised by the Hotpot contract.
type HotpotWinnersAssignedIterator struct {
	Event *HotpotWinnersAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HotpotWinnersAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HotpotWinnersAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HotpotWinnersAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HotpotWinnersAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HotpotWinnersAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HotpotWinnersAssigned represents a WinnersAssigned event raised by the Hotpot contract.
type HotpotWinnersAssigned struct {
	Winners []common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWinnersAssigned is a free log retrieval operation binding the contract event 0x6e4b5e113e3c80149f0005be76d3eeafe668dba9e5de7fa8bb935d283bccd2a7.
//
// Solidity: event WinnersAssigned(address[] _winners)
func (_Hotpot *HotpotFilterer) FilterWinnersAssigned(opts *bind.FilterOpts) (*HotpotWinnersAssignedIterator, error) {

	logs, sub, err := _Hotpot.contract.FilterLogs(opts, "WinnersAssigned")
	if err != nil {
		return nil, err
	}
	return &HotpotWinnersAssignedIterator{contract: _Hotpot.contract, event: "WinnersAssigned", logs: logs, sub: sub}, nil
}

// WatchWinnersAssigned is a free log subscription operation binding the contract event 0x6e4b5e113e3c80149f0005be76d3eeafe668dba9e5de7fa8bb935d283bccd2a7.
//
// Solidity: event WinnersAssigned(address[] _winners)
func (_Hotpot *HotpotFilterer) WatchWinnersAssigned(opts *bind.WatchOpts, sink chan<- *HotpotWinnersAssigned) (event.Subscription, error) {

	logs, sub, err := _Hotpot.contract.WatchLogs(opts, "WinnersAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HotpotWinnersAssigned)
				if err := _Hotpot.contract.UnpackLog(event, "WinnersAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWinnersAssigned is a log parse operation binding the contract event 0x6e4b5e113e3c80149f0005be76d3eeafe668dba9e5de7fa8bb935d283bccd2a7.
//
// Solidity: event WinnersAssigned(address[] _winners)
func (_Hotpot *HotpotFilterer) ParseWinnersAssigned(log types.Log) (*HotpotWinnersAssigned, error) {
	event := new(HotpotWinnersAssigned)
	if err := _Hotpot.contract.UnpackLog(event, "WinnersAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
