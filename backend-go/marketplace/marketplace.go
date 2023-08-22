// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

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

// IOrderFulfillerOfferItem is an auto generated low-level Go binding around an user-defined struct.
type IOrderFulfillerOfferItem struct {
	OfferToken   common.Address
	OfferTokenId *big.Int
	OfferAmount  *big.Int
	EndTime      *big.Int
}

// IOrderFulfillerOrderParameters is an auto generated low-level Go binding around an user-defined struct.
type IOrderFulfillerOrderParameters struct {
	Offerer                 common.Address
	OfferItem               IOrderFulfillerOfferItem
	Royalty                 IOrderFulfillerRoyaltyData
	PendingAmountsData      IOrderFulfillerPendingAmountData
	Salt                    *big.Int
	OrderSignature          []byte
	PendingAmountsSignature []byte
}

// IOrderFulfillerPendingAmountData is an auto generated low-level Go binding around an user-defined struct.
type IOrderFulfillerPendingAmountData struct {
	OffererPendingAmount *big.Int
	BuyerPendingAmount   *big.Int
	OrderHash            [32]byte
}

// IOrderFulfillerPureOrder is an auto generated low-level Go binding around an user-defined struct.
type IOrderFulfillerPureOrder struct {
	Offerer   common.Address
	OfferItem IOrderFulfillerOfferItem
	Royalty   IOrderFulfillerRoyaltyData
	Salt      *big.Int
}

// IOrderFulfillerRoyaltyData is an auto generated low-level Go binding around an user-defined struct.
type IOrderFulfillerRoyaltyData struct {
	RoyaltyPercent   *big.Int
	RoyaltyRecipient common.Address
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"OperatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_raffleAddress\",\"type\":\"address\"}],\"name\":\"RaffleAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_newTradeFee\",\"type\":\"uint16\"}],\"name\":\"RaffleTradeFeeChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderFulfiller.OfferItem\",\"name\":\"offerItem\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"royaltyPercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"internalType\":\"structIOrderFulfiller.RoyaltyData\",\"name\":\"royalty\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderFulfiller.PureOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"offerer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"offerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offerTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offerAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"internalType\":\"structIOrderFulfiller.OfferItem\",\"name\":\"offerItem\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"royaltyPercent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"}],\"internalType\":\"structIOrderFulfiller.RoyaltyData\",\"name\":\"royalty\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"offererPendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerPendingAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"internalType\":\"structIOrderFulfiller.PendingAmountData\",\"name\":\"pendingAmountsData\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"orderSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"pendingAmountsSignature\",\"type\":\"bytes\"}],\"internalType\":\"structIOrderFulfiller.OrderParameters\",\"name\":\"parameters\",\"type\":\"tuple\"}],\"name\":\"fulfillOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_raffleTradeFee\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"raffleTradeFee\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_raffleAddress\",\"type\":\"address\"}],\"name\":\"setRaffleAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_newTradeFee\",\"type\":\"uint16\"}],\"name\":\"setRaffleTradeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Marketplace *MarketplaceSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Marketplace.Contract.DOMAINSEPARATOR(&_Marketplace.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Marketplace.Contract.DOMAINSEPARATOR(&_Marketplace.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Marketplace *MarketplaceCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Marketplace *MarketplaceSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Marketplace.Contract.Eip712Domain(&_Marketplace.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Marketplace *MarketplaceCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Marketplace.Contract.Eip712Domain(&_Marketplace.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Marketplace *MarketplaceCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Marketplace *MarketplaceSession) Operator() (common.Address, error) {
	return _Marketplace.Contract.Operator(&_Marketplace.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Operator() (common.Address, error) {
	return _Marketplace.Contract.Operator(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// RaffleContract is a free data retrieval call binding the contract method 0x5d7916ed.
//
// Solidity: function raffleContract() view returns(address)
func (_Marketplace *MarketplaceCaller) RaffleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "raffleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RaffleContract is a free data retrieval call binding the contract method 0x5d7916ed.
//
// Solidity: function raffleContract() view returns(address)
func (_Marketplace *MarketplaceSession) RaffleContract() (common.Address, error) {
	return _Marketplace.Contract.RaffleContract(&_Marketplace.CallOpts)
}

// RaffleContract is a free data retrieval call binding the contract method 0x5d7916ed.
//
// Solidity: function raffleContract() view returns(address)
func (_Marketplace *MarketplaceCallerSession) RaffleContract() (common.Address, error) {
	return _Marketplace.Contract.RaffleContract(&_Marketplace.CallOpts)
}

// RaffleTradeFee is a free data retrieval call binding the contract method 0x7dd5a2af.
//
// Solidity: function raffleTradeFee() view returns(uint16)
func (_Marketplace *MarketplaceCaller) RaffleTradeFee(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "raffleTradeFee")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// RaffleTradeFee is a free data retrieval call binding the contract method 0x7dd5a2af.
//
// Solidity: function raffleTradeFee() view returns(uint16)
func (_Marketplace *MarketplaceSession) RaffleTradeFee() (uint16, error) {
	return _Marketplace.Contract.RaffleTradeFee(&_Marketplace.CallOpts)
}

// RaffleTradeFee is a free data retrieval call binding the contract method 0x7dd5a2af.
//
// Solidity: function raffleTradeFee() view returns(uint16)
func (_Marketplace *MarketplaceCallerSession) RaffleTradeFee() (uint16, error) {
	return _Marketplace.Contract.RaffleTradeFee(&_Marketplace.CallOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x200916ff.
//
// Solidity: function cancelOrder((address,(address,uint256,uint256,uint256),(uint256,address),uint256) order) returns()
func (_Marketplace *MarketplaceTransactor) CancelOrder(opts *bind.TransactOpts, order IOrderFulfillerPureOrder) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x200916ff.
//
// Solidity: function cancelOrder((address,(address,uint256,uint256,uint256),(uint256,address),uint256) order) returns()
func (_Marketplace *MarketplaceSession) CancelOrder(order IOrderFulfillerPureOrder) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelOrder(&_Marketplace.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x200916ff.
//
// Solidity: function cancelOrder((address,(address,uint256,uint256,uint256),(uint256,address),uint256) order) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelOrder(order IOrderFulfillerPureOrder) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelOrder(&_Marketplace.TransactOpts, order)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xbd2cd3dc.
//
// Solidity: function fulfillOrder((address,(address,uint256,uint256,uint256),(uint256,address),(uint256,uint256,bytes32),uint256,bytes,bytes) parameters) payable returns()
func (_Marketplace *MarketplaceTransactor) FulfillOrder(opts *bind.TransactOpts, parameters IOrderFulfillerOrderParameters) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "fulfillOrder", parameters)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xbd2cd3dc.
//
// Solidity: function fulfillOrder((address,(address,uint256,uint256,uint256),(uint256,address),(uint256,uint256,bytes32),uint256,bytes,bytes) parameters) payable returns()
func (_Marketplace *MarketplaceSession) FulfillOrder(parameters IOrderFulfillerOrderParameters) (*types.Transaction, error) {
	return _Marketplace.Contract.FulfillOrder(&_Marketplace.TransactOpts, parameters)
}

// FulfillOrder is a paid mutator transaction binding the contract method 0xbd2cd3dc.
//
// Solidity: function fulfillOrder((address,(address,uint256,uint256,uint256),(uint256,address),(uint256,uint256,bytes32),uint256,bytes,bytes) parameters) payable returns()
func (_Marketplace *MarketplaceTransactorSession) FulfillOrder(parameters IOrderFulfillerOrderParameters) (*types.Transaction, error) {
	return _Marketplace.Contract.FulfillOrder(&_Marketplace.TransactOpts, parameters)
}

// Initialize is a paid mutator transaction binding the contract method 0xe0dbcde5.
//
// Solidity: function initialize(uint16 _raffleTradeFee, address _operator) returns()
func (_Marketplace *MarketplaceTransactor) Initialize(opts *bind.TransactOpts, _raffleTradeFee uint16, _operator common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "initialize", _raffleTradeFee, _operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xe0dbcde5.
//
// Solidity: function initialize(uint16 _raffleTradeFee, address _operator) returns()
func (_Marketplace *MarketplaceSession) Initialize(_raffleTradeFee uint16, _operator common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Initialize(&_Marketplace.TransactOpts, _raffleTradeFee, _operator)
}

// Initialize is a paid mutator transaction binding the contract method 0xe0dbcde5.
//
// Solidity: function initialize(uint16 _raffleTradeFee, address _operator) returns()
func (_Marketplace *MarketplaceTransactorSession) Initialize(_raffleTradeFee uint16, _operator common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Initialize(&_Marketplace.TransactOpts, _raffleTradeFee, _operator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_Marketplace *MarketplaceTransactor) SetOperator(opts *bind.TransactOpts, _newOperator common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setOperator", _newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_Marketplace *MarketplaceSession) SetOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetOperator(&_Marketplace.TransactOpts, _newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_Marketplace *MarketplaceTransactorSession) SetOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetOperator(&_Marketplace.TransactOpts, _newOperator)
}

// SetRaffleAddress is a paid mutator transaction binding the contract method 0x97709ce7.
//
// Solidity: function setRaffleAddress(address _raffleAddress) returns()
func (_Marketplace *MarketplaceTransactor) SetRaffleAddress(opts *bind.TransactOpts, _raffleAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setRaffleAddress", _raffleAddress)
}

// SetRaffleAddress is a paid mutator transaction binding the contract method 0x97709ce7.
//
// Solidity: function setRaffleAddress(address _raffleAddress) returns()
func (_Marketplace *MarketplaceSession) SetRaffleAddress(_raffleAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRaffleAddress(&_Marketplace.TransactOpts, _raffleAddress)
}

// SetRaffleAddress is a paid mutator transaction binding the contract method 0x97709ce7.
//
// Solidity: function setRaffleAddress(address _raffleAddress) returns()
func (_Marketplace *MarketplaceTransactorSession) SetRaffleAddress(_raffleAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRaffleAddress(&_Marketplace.TransactOpts, _raffleAddress)
}

// SetRaffleTradeFee is a paid mutator transaction binding the contract method 0x01111b36.
//
// Solidity: function setRaffleTradeFee(uint16 _newTradeFee) returns()
func (_Marketplace *MarketplaceTransactor) SetRaffleTradeFee(opts *bind.TransactOpts, _newTradeFee uint16) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setRaffleTradeFee", _newTradeFee)
}

// SetRaffleTradeFee is a paid mutator transaction binding the contract method 0x01111b36.
//
// Solidity: function setRaffleTradeFee(uint16 _newTradeFee) returns()
func (_Marketplace *MarketplaceSession) SetRaffleTradeFee(_newTradeFee uint16) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRaffleTradeFee(&_Marketplace.TransactOpts, _newTradeFee)
}

// SetRaffleTradeFee is a paid mutator transaction binding the contract method 0x01111b36.
//
// Solidity: function setRaffleTradeFee(uint16 _newTradeFee) returns()
func (_Marketplace *MarketplaceTransactorSession) SetRaffleTradeFee(_newTradeFee uint16) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRaffleTradeFee(&_Marketplace.TransactOpts, _newTradeFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// MarketplaceEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Marketplace contract.
type MarketplaceEIP712DomainChangedIterator struct {
	Event *MarketplaceEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *MarketplaceEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceEIP712DomainChanged)
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
		it.Event = new(MarketplaceEIP712DomainChanged)
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
func (it *MarketplaceEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceEIP712DomainChanged represents a EIP712DomainChanged event raised by the Marketplace contract.
type MarketplaceEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Marketplace *MarketplaceFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*MarketplaceEIP712DomainChangedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &MarketplaceEIP712DomainChangedIterator{contract: _Marketplace.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Marketplace *MarketplaceFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *MarketplaceEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceEIP712DomainChanged)
				if err := _Marketplace.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Marketplace *MarketplaceFilterer) ParseEIP712DomainChanged(log types.Log) (*MarketplaceEIP712DomainChanged, error) {
	event := new(MarketplaceEIP712DomainChanged)
	if err := _Marketplace.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Marketplace contract.
type MarketplaceInitializedIterator struct {
	Event *MarketplaceInitialized // Event containing the contract specifics and raw log

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
func (it *MarketplaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceInitialized)
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
		it.Event = new(MarketplaceInitialized)
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
func (it *MarketplaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceInitialized represents a Initialized event raised by the Marketplace contract.
type MarketplaceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Marketplace *MarketplaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*MarketplaceInitializedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MarketplaceInitializedIterator{contract: _Marketplace.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Marketplace *MarketplaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MarketplaceInitialized) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceInitialized)
				if err := _Marketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseInitialized(log types.Log) (*MarketplaceInitialized, error) {
	event := new(MarketplaceInitialized)
	if err := _Marketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceOperatorChangedIterator is returned from FilterOperatorChanged and is used to iterate over the raw logs and unpacked data for OperatorChanged events raised by the Marketplace contract.
type MarketplaceOperatorChangedIterator struct {
	Event *MarketplaceOperatorChanged // Event containing the contract specifics and raw log

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
func (it *MarketplaceOperatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceOperatorChanged)
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
		it.Event = new(MarketplaceOperatorChanged)
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
func (it *MarketplaceOperatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceOperatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceOperatorChanged represents a OperatorChanged event raised by the Marketplace contract.
type MarketplaceOperatorChanged struct {
	NewOperator common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOperatorChanged is a free log retrieval operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_Marketplace *MarketplaceFilterer) FilterOperatorChanged(opts *bind.FilterOpts) (*MarketplaceOperatorChangedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return &MarketplaceOperatorChangedIterator{contract: _Marketplace.contract, event: "OperatorChanged", logs: logs, sub: sub}, nil
}

// WatchOperatorChanged is a free log subscription operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_Marketplace *MarketplaceFilterer) WatchOperatorChanged(opts *bind.WatchOpts, sink chan<- *MarketplaceOperatorChanged) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "OperatorChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceOperatorChanged)
				if err := _Marketplace.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
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

// ParseOperatorChanged is a log parse operation binding the contract event 0x4721129e0e676ed6a92909bb24e853ccdd63ad72280cc2e974e38e480e0e6e54.
//
// Solidity: event OperatorChanged(address _newOperator)
func (_Marketplace *MarketplaceFilterer) ParseOperatorChanged(log types.Log) (*MarketplaceOperatorChanged, error) {
	event := new(MarketplaceOperatorChanged)
	if err := _Marketplace.contract.UnpackLog(event, "OperatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Marketplace contract.
type MarketplaceOrderCancelledIterator struct {
	Event *MarketplaceOrderCancelled // Event containing the contract specifics and raw log

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
func (it *MarketplaceOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceOrderCancelled)
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
		it.Event = new(MarketplaceOrderCancelled)
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
func (it *MarketplaceOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceOrderCancelled represents a OrderCancelled event raised by the Marketplace contract.
type MarketplaceOrderCancelled struct {
	Offerer    common.Address
	OfferToken common.Address
	TokenId    *big.Int
	OrderHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x151f2e6a40b8eb99b522d83ee781be0fda3a72f094311e92c5108e7ea7f83ec3.
//
// Solidity: event OrderCancelled(address offerer, address offerToken, uint256 tokenId, bytes32 orderHash)
func (_Marketplace *MarketplaceFilterer) FilterOrderCancelled(opts *bind.FilterOpts) (*MarketplaceOrderCancelledIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &MarketplaceOrderCancelledIterator{contract: _Marketplace.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x151f2e6a40b8eb99b522d83ee781be0fda3a72f094311e92c5108e7ea7f83ec3.
//
// Solidity: event OrderCancelled(address offerer, address offerToken, uint256 tokenId, bytes32 orderHash)
func (_Marketplace *MarketplaceFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *MarketplaceOrderCancelled) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceOrderCancelled)
				if err := _Marketplace.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x151f2e6a40b8eb99b522d83ee781be0fda3a72f094311e92c5108e7ea7f83ec3.
//
// Solidity: event OrderCancelled(address offerer, address offerToken, uint256 tokenId, bytes32 orderHash)
func (_Marketplace *MarketplaceFilterer) ParseOrderCancelled(log types.Log) (*MarketplaceOrderCancelled, error) {
	event := new(MarketplaceOrderCancelled)
	if err := _Marketplace.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceOrderFulfilledIterator is returned from FilterOrderFulfilled and is used to iterate over the raw logs and unpacked data for OrderFulfilled events raised by the Marketplace contract.
type MarketplaceOrderFulfilledIterator struct {
	Event *MarketplaceOrderFulfilled // Event containing the contract specifics and raw log

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
func (it *MarketplaceOrderFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceOrderFulfilled)
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
		it.Event = new(MarketplaceOrderFulfilled)
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
func (it *MarketplaceOrderFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceOrderFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceOrderFulfilled represents a OrderFulfilled event raised by the Marketplace contract.
type MarketplaceOrderFulfilled struct {
	Offerer     common.Address
	Buyer       common.Address
	OfferToken  common.Address
	TokenId     *big.Int
	TradeAmount *big.Int
	OrderHash   [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOrderFulfilled is a free log retrieval operation binding the contract event 0x1bbc142e8dd96aee2d7b1e4ceab7afdb9ccd9774e0bfc9215a26cf51f651f018.
//
// Solidity: event OrderFulfilled(address offerer, address buyer, address offerToken, uint256 tokenId, uint256 tradeAmount, bytes32 orderHash)
func (_Marketplace *MarketplaceFilterer) FilterOrderFulfilled(opts *bind.FilterOpts) (*MarketplaceOrderFulfilledIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return &MarketplaceOrderFulfilledIterator{contract: _Marketplace.contract, event: "OrderFulfilled", logs: logs, sub: sub}, nil
}

// WatchOrderFulfilled is a free log subscription operation binding the contract event 0x1bbc142e8dd96aee2d7b1e4ceab7afdb9ccd9774e0bfc9215a26cf51f651f018.
//
// Solidity: event OrderFulfilled(address offerer, address buyer, address offerToken, uint256 tokenId, uint256 tradeAmount, bytes32 orderHash)
func (_Marketplace *MarketplaceFilterer) WatchOrderFulfilled(opts *bind.WatchOpts, sink chan<- *MarketplaceOrderFulfilled) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "OrderFulfilled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceOrderFulfilled)
				if err := _Marketplace.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
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

// ParseOrderFulfilled is a log parse operation binding the contract event 0x1bbc142e8dd96aee2d7b1e4ceab7afdb9ccd9774e0bfc9215a26cf51f651f018.
//
// Solidity: event OrderFulfilled(address offerer, address buyer, address offerToken, uint256 tokenId, uint256 tradeAmount, bytes32 orderHash)
func (_Marketplace *MarketplaceFilterer) ParseOrderFulfilled(log types.Log) (*MarketplaceOrderFulfilled, error) {
	event := new(MarketplaceOrderFulfilled)
	if err := _Marketplace.contract.UnpackLog(event, "OrderFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Marketplace contract.
type MarketplaceOwnershipTransferredIterator struct {
	Event *MarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceOwnershipTransferred)
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
		it.Event = new(MarketplaceOwnershipTransferred)
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
func (it *MarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the Marketplace contract.
type MarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceOwnershipTransferredIterator{contract: _Marketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceOwnershipTransferred)
				if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*MarketplaceOwnershipTransferred, error) {
	event := new(MarketplaceOwnershipTransferred)
	if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRaffleAddressSetIterator is returned from FilterRaffleAddressSet and is used to iterate over the raw logs and unpacked data for RaffleAddressSet events raised by the Marketplace contract.
type MarketplaceRaffleAddressSetIterator struct {
	Event *MarketplaceRaffleAddressSet // Event containing the contract specifics and raw log

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
func (it *MarketplaceRaffleAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRaffleAddressSet)
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
		it.Event = new(MarketplaceRaffleAddressSet)
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
func (it *MarketplaceRaffleAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRaffleAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRaffleAddressSet represents a RaffleAddressSet event raised by the Marketplace contract.
type MarketplaceRaffleAddressSet struct {
	RaffleAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRaffleAddressSet is a free log retrieval operation binding the contract event 0xa9b95bfc7516964dba0c75460ee04dc0fe53113a2e9aea8fedca3378bdd31fb9.
//
// Solidity: event RaffleAddressSet(address _raffleAddress)
func (_Marketplace *MarketplaceFilterer) FilterRaffleAddressSet(opts *bind.FilterOpts) (*MarketplaceRaffleAddressSetIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RaffleAddressSet")
	if err != nil {
		return nil, err
	}
	return &MarketplaceRaffleAddressSetIterator{contract: _Marketplace.contract, event: "RaffleAddressSet", logs: logs, sub: sub}, nil
}

// WatchRaffleAddressSet is a free log subscription operation binding the contract event 0xa9b95bfc7516964dba0c75460ee04dc0fe53113a2e9aea8fedca3378bdd31fb9.
//
// Solidity: event RaffleAddressSet(address _raffleAddress)
func (_Marketplace *MarketplaceFilterer) WatchRaffleAddressSet(opts *bind.WatchOpts, sink chan<- *MarketplaceRaffleAddressSet) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RaffleAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRaffleAddressSet)
				if err := _Marketplace.contract.UnpackLog(event, "RaffleAddressSet", log); err != nil {
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

// ParseRaffleAddressSet is a log parse operation binding the contract event 0xa9b95bfc7516964dba0c75460ee04dc0fe53113a2e9aea8fedca3378bdd31fb9.
//
// Solidity: event RaffleAddressSet(address _raffleAddress)
func (_Marketplace *MarketplaceFilterer) ParseRaffleAddressSet(log types.Log) (*MarketplaceRaffleAddressSet, error) {
	event := new(MarketplaceRaffleAddressSet)
	if err := _Marketplace.contract.UnpackLog(event, "RaffleAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRaffleTradeFeeChangedIterator is returned from FilterRaffleTradeFeeChanged and is used to iterate over the raw logs and unpacked data for RaffleTradeFeeChanged events raised by the Marketplace contract.
type MarketplaceRaffleTradeFeeChangedIterator struct {
	Event *MarketplaceRaffleTradeFeeChanged // Event containing the contract specifics and raw log

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
func (it *MarketplaceRaffleTradeFeeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRaffleTradeFeeChanged)
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
		it.Event = new(MarketplaceRaffleTradeFeeChanged)
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
func (it *MarketplaceRaffleTradeFeeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRaffleTradeFeeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRaffleTradeFeeChanged represents a RaffleTradeFeeChanged event raised by the Marketplace contract.
type MarketplaceRaffleTradeFeeChanged struct {
	NewTradeFee uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRaffleTradeFeeChanged is a free log retrieval operation binding the contract event 0x229d146fba0de2d2d6b0ae22fa6469aae00e699998fb08a692f3bd8105262494.
//
// Solidity: event RaffleTradeFeeChanged(uint16 _newTradeFee)
func (_Marketplace *MarketplaceFilterer) FilterRaffleTradeFeeChanged(opts *bind.FilterOpts) (*MarketplaceRaffleTradeFeeChangedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RaffleTradeFeeChanged")
	if err != nil {
		return nil, err
	}
	return &MarketplaceRaffleTradeFeeChangedIterator{contract: _Marketplace.contract, event: "RaffleTradeFeeChanged", logs: logs, sub: sub}, nil
}

// WatchRaffleTradeFeeChanged is a free log subscription operation binding the contract event 0x229d146fba0de2d2d6b0ae22fa6469aae00e699998fb08a692f3bd8105262494.
//
// Solidity: event RaffleTradeFeeChanged(uint16 _newTradeFee)
func (_Marketplace *MarketplaceFilterer) WatchRaffleTradeFeeChanged(opts *bind.WatchOpts, sink chan<- *MarketplaceRaffleTradeFeeChanged) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RaffleTradeFeeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRaffleTradeFeeChanged)
				if err := _Marketplace.contract.UnpackLog(event, "RaffleTradeFeeChanged", log); err != nil {
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

// ParseRaffleTradeFeeChanged is a log parse operation binding the contract event 0x229d146fba0de2d2d6b0ae22fa6469aae00e699998fb08a692f3bd8105262494.
//
// Solidity: event RaffleTradeFeeChanged(uint16 _newTradeFee)
func (_Marketplace *MarketplaceFilterer) ParseRaffleTradeFeeChanged(log types.Log) (*MarketplaceRaffleTradeFeeChanged, error) {
	event := new(MarketplaceRaffleTradeFeeChanged)
	if err := _Marketplace.contract.UnpackLog(event, "RaffleTradeFeeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
