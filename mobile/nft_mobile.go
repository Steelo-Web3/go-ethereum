// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package geth

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
)

// NftMetaData contains all meta data concerning the Nft contract.
var NftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenTypeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"authorizedMinter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"createTokenType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenTypeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenHolder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenTypeURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenTypes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NftABI is the input ABI used to generate the binding from.
// Deprecated: Use NftMetaData.ABI instead.
var NftABI = NftMetaData.ABI

// Nft is an auto generated Go binding around an Ethereum contract.
type Nft struct {
	NftCaller     // Read-only binding to the contract
	NftTransactor // Write-only binding to the contract
	NftFilterer   // Log filterer for contract events
}

// NftCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftCaller struct {
	contract *BoundContract // Generic contract wrapper for the low level calls
}

// NftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftTransactor struct {
	contract *BoundContract // Generic contract wrapper for the low level calls
}

// NftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftFilterer struct {
	contract *BoundContract // Generic contract wrapper for the low level calls
}

// NftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftSession struct {
	Contract     *Nft         // Generic contract binding to set the session for
	CallOpts     CallOpts     // Call options to use throughout this session
	TransactOpts TransactOpts // Transaction auth options to use throughout this session
}

// NftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftCallerSession struct {
	Contract *NftCaller // Generic contract caller binding to set the session for
	CallOpts CallOpts   // Call options to use throughout this session
}

// NftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftTransactorSession struct {
	Contract     *NftTransactor // Generic contract transactor binding to set the session for
	TransactOpts TransactOpts   // Transaction auth options to use throughout this session
}

// NftRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftRaw struct {
	Contract *Nft // Generic contract binding to access the raw methods on
}

// NftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftCallerRaw struct {
	Contract *NftCaller // Generic read-only contract binding to access the raw methods on
}

// NftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftTransactorRaw struct {
	Contract *NftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNft creates a new instance of Nft, bound to a specific deployed contract.
func NewNft(address Address, client *EthereumClient) (*Nft, error) {
	contract, err := bindNft(address.address, client.client, client.client, client.client)
	if err != nil {
		return nil, err
	}
	return &Nft{NftCaller: NftCaller{contract: contract}, NftTransactor: NftTransactor{contract: contract}, NftFilterer: NftFilterer{contract: contract}}, nil
}

// bindNft binds a generic wrapper to an already deployed contract.
func bindNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NftABI))
	if err != nil {
		return nil, err
	}
	return &BoundContract{
		contract: bind.NewBoundContract(address, parsed, caller, transactor, filterer),
	}, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nft *NftRaw) Call(opts *CallOpts, result *Interfaces, method string, params *Interfaces) error {
	return _Nft.Contract.NftCaller.contract.Call(opts, result, method, params)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nft *NftRaw) Transfer(opts *TransactOpts) (*Transaction, error) {
	return _Nft.Contract.NftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nft *NftRaw) Transact(opts *TransactOpts, method string, params *Interfaces) (*Transaction, error) {
	return _Nft.Contract.NftTransactor.contract.Transact(opts, method, params)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nft *NftCallerRaw) Call(opts *CallOpts, result *Interfaces, method string, params *Interfaces) error {
	return _Nft.Contract.contract.Call(opts, result, method, params)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nft *NftTransactorRaw) Transfer(opts *TransactOpts) (*Transaction, error) {
	return _Nft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nft *NftTransactorRaw) Transact(opts *TransactOpts, method string, params *Interfaces) (*Transaction, error) {
	return _Nft.Contract.contract.Transact(opts, method, params)
}

// TokenHolder is a free data retrieval call binding the contract method 0x862a4bf2.
//
// Solidity: function tokenHolder(uint256 ) view returns(address)
func (_Nft *NftCaller) TokenHolder(opts *CallOpts, arg0 *BigInt) (Address, error) {
	var out Interfaces
	err := _Nft.contract.Call(opts, &out, "tokenHolder", &Interfaces{objects: []interface{}{arg0}})

	if err != nil {
		return *new(Address), err
	}

	out0 := *abi.ConvertType(out.objects[0], new(Address)).(*Address)

	return out0, err

}

// TokenHolder is a free data retrieval call binding the contract method 0x862a4bf2.
//
// Solidity: function tokenHolder(uint256 ) view returns(address)
func (_Nft *NftSession) TokenHolder(arg0 *BigInt) (Address, error) {
	return _Nft.Contract.TokenHolder(&_Nft.CallOpts, arg0)
}

// TokenHolder is a free data retrieval call binding the contract method 0x862a4bf2.
//
// Solidity: function tokenHolder(uint256 ) view returns(address)
func (_Nft *NftCallerSession) TokenHolder(arg0 *BigInt) (Address, error) {
	return _Nft.Contract.TokenHolder(&_Nft.CallOpts, arg0)
}

// TokenTypeURI is a free data retrieval call binding the contract method 0x885a72a9.
//
// Solidity: function tokenTypeURI(uint256 ) view returns(string)
func (_Nft *NftCaller) TokenTypeURI(opts *CallOpts, arg0 *BigInt) (string, error) {
	var out Interfaces
	err := _Nft.contract.Call(opts, &out, "tokenTypeURI", &Interfaces{objects: []interface{}{arg0}})

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out.objects[0], new(string)).(*string)

	return out0, err

}

// TokenTypeURI is a free data retrieval call binding the contract method 0x885a72a9.
//
// Solidity: function tokenTypeURI(uint256 ) view returns(string)
func (_Nft *NftSession) TokenTypeURI(arg0 *BigInt) (string, error) {
	return _Nft.Contract.TokenTypeURI(&_Nft.CallOpts, arg0)
}

// TokenTypeURI is a free data retrieval call binding the contract method 0x885a72a9.
//
// Solidity: function tokenTypeURI(uint256 ) view returns(string)
func (_Nft *NftCallerSession) TokenTypeURI(arg0 *BigInt) (string, error) {
	return _Nft.Contract.TokenTypeURI(&_Nft.CallOpts, arg0)
}

// TokenTypes is a free data retrieval call binding the contract method 0x33f6832a.
//
// Solidity: function tokenTypes(uint256 ) view returns(bool)
func (_Nft *NftCaller) TokenTypes(opts *CallOpts, arg0 *BigInt) (bool, error) {
	var out Interfaces
	err := _Nft.contract.Call(opts, &out, "tokenTypes", &Interfaces{objects: []interface{}{arg0}})

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out.objects[0], new(bool)).(*bool)

	return out0, err

}

// TokenTypes is a free data retrieval call binding the contract method 0x33f6832a.
//
// Solidity: function tokenTypes(uint256 ) view returns(bool)
func (_Nft *NftSession) TokenTypes(arg0 *BigInt) (bool, error) {
	return _Nft.Contract.TokenTypes(&_Nft.CallOpts, arg0)
}

// TokenTypes is a free data retrieval call binding the contract method 0x33f6832a.
//
// Solidity: function tokenTypes(uint256 ) view returns(bool)
func (_Nft *NftCallerSession) TokenTypes(arg0 *BigInt) (bool, error) {
	return _Nft.Contract.TokenTypes(&_Nft.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) view returns(string)
func (_Nft *NftCaller) TokenURI(opts *CallOpts, arg0 *BigInt) (string, error) {
	var out Interfaces
	err := _Nft.contract.Call(opts, &out, "tokenURI", &Interfaces{objects: []interface{}{arg0}})

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out.objects[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) view returns(string)
func (_Nft *NftSession) TokenURI(arg0 *BigInt) (string, error) {
	return _Nft.Contract.TokenURI(&_Nft.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) view returns(string)
func (_Nft *NftCallerSession) TokenURI(arg0 *BigInt) (string, error) {
	return _Nft.Contract.TokenURI(&_Nft.CallOpts, arg0)
}

// CreateTokenType is a paid mutator transaction binding the contract method 0x9799e525.
//
// Solidity: function createTokenType(uint256 tokenTypeID, address authorizedMinter, string uri) returns()
func (_Nft *NftTransactor) CreateTokenType(opts *TransactOpts, tokenTypeID *BigInt, authorizedMinter Address, uri string) (*Transaction, error) {
	return _Nft.contract.Transact(opts, "createTokenType", &Interfaces{objects: []interface{}{tokenTypeID, authorizedMinter, uri}})
}

// CreateTokenType is a paid mutator transaction binding the contract method 0x9799e525.
//
// Solidity: function createTokenType(uint256 tokenTypeID, address authorizedMinter, string uri) returns()
func (_Nft *NftSession) CreateTokenType(tokenTypeID *BigInt, authorizedMinter Address, uri string) (*Transaction, error) {
	return _Nft.Contract.CreateTokenType(&_Nft.TransactOpts, tokenTypeID, authorizedMinter, uri)
}

// CreateTokenType is a paid mutator transaction binding the contract method 0x9799e525.
//
// Solidity: function createTokenType(uint256 tokenTypeID, address authorizedMinter, string uri) returns()
func (_Nft *NftTransactorSession) CreateTokenType(tokenTypeID *BigInt, authorizedMinter Address, uri string) (*Transaction, error) {
	return _Nft.Contract.CreateTokenType(&_Nft.TransactOpts, tokenTypeID, authorizedMinter, uri)
}

// MintToken is a paid mutator transaction binding the contract method 0x87c6649c.
//
// Solidity: function mintToken(address to, uint256 tokenTypeID, uint256 tokenID, string uri) returns()
func (_Nft *NftTransactor) MintToken(opts *TransactOpts, to Address, tokenTypeID *BigInt, tokenID *BigInt, uri string) (*Transaction, error) {
	return _Nft.contract.Transact(opts, "mintToken", &Interfaces{objects: []interface{}{to.address, tokenTypeID.bigint, tokenID.bigint, uri}})
}

// MintToken is a paid mutator transaction binding the contract method 0x87c6649c.
//
// Solidity: function mintToken(address to, uint256 tokenTypeID, uint256 tokenID, string uri) returns()
func (_Nft *NftSession) MintToken(to Address, tokenTypeID *BigInt, tokenID *BigInt, uri string) (*Transaction, error) {
	return _Nft.Contract.MintToken(&_Nft.TransactOpts, to, tokenTypeID, tokenID, uri)
}

// MintToken is a paid mutator transaction binding the contract method 0x87c6649c.
//
// Solidity: function mintToken(address to, uint256 tokenTypeID, uint256 tokenID, string uri) returns()
func (_Nft *NftTransactorSession) MintToken(to Address, tokenTypeID *BigInt, tokenID *BigInt, uri string) (*Transaction, error) {
	return _Nft.Contract.MintToken(&_Nft.TransactOpts, to, tokenTypeID, tokenID, uri)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 tokenID) returns()
func (_Nft *NftTransactor) TransferToken(opts *TransactOpts, to Address, tokenID *BigInt) (*Transaction, error) {
	return _Nft.contract.Transact(opts, "transferToken", &Interfaces{objects: []interface{}{to.address, tokenID.bigint}})
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 tokenID) returns()
func (_Nft *NftSession) TransferToken(to Address, tokenID *BigInt) (*Transaction, error) {
	return _Nft.Contract.TransferToken(&_Nft.TransactOpts, to, tokenID)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 tokenID) returns()
func (_Nft *NftTransactorSession) TransferToken(to Address, tokenID *BigInt) (*Transaction, error) {
	return _Nft.Contract.TransferToken(&_Nft.TransactOpts, to, tokenID)
}
