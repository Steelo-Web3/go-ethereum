// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package codegen

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

// NFTMetaData contains all meta data concerning the NFT contract.
var NFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenTypeID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintCount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"NewTokenType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"burnToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenTypeID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"authorizedMinter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"createTokenType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"getIssued\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"getWalletTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"issuerTokenTypes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenTypeID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"mintToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenHolder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenMetadataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenTokenType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenTypeAuthorizedMinters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenTypeMetadataURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenTypeMintCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenTypes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenID\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wallets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NFTABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTMetaData.ABI instead.
var NFTABI = NFTMetaData.ABI

// NFT is an auto generated Go binding around an Ethereum contract.
type NFT struct {
	NFTCaller     // Read-only binding to the contract
	NFTTransactor // Write-only binding to the contract
	NFTFilterer   // Log filterer for contract events
}

// NFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTSession struct {
	Contract     *NFT              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTCallerSession struct {
	Contract *NFTCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTTransactorSession struct {
	Contract     *NFTTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTRaw struct {
	Contract *NFT // Generic contract binding to access the raw methods on
}

// NFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTCallerRaw struct {
	Contract *NFTCaller // Generic read-only contract binding to access the raw methods on
}

// NFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTTransactorRaw struct {
	Contract *NFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFT creates a new instance of NFT, bound to a specific deployed contract.
func NewNFT(address common.Address, backend bind.ContractBackend) (*NFT, error) {
	contract, err := bindNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFT{NFTCaller: NFTCaller{contract: contract}, NFTTransactor: NFTTransactor{contract: contract}, NFTFilterer: NFTFilterer{contract: contract}}, nil
}

// NewNFTCaller creates a new read-only instance of NFT, bound to a specific deployed contract.
func NewNFTCaller(address common.Address, caller bind.ContractCaller) (*NFTCaller, error) {
	contract, err := bindNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTCaller{contract: contract}, nil
}

// NewNFTTransactor creates a new write-only instance of NFT, bound to a specific deployed contract.
func NewNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTTransactor, error) {
	contract, err := bindNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTTransactor{contract: contract}, nil
}

// NewNFTFilterer creates a new log filterer instance of NFT, bound to a specific deployed contract.
func NewNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTFilterer, error) {
	contract, err := bindNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTFilterer{contract: contract}, nil
}

// bindNFT binds a generic wrapper to an already deployed contract.
func bindNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFT.Contract.NFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.Contract.NFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT.Contract.NFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT.Contract.contract.Transact(opts, method, params...)
}

// GetIssued is a free data retrieval call binding the contract method 0xe895bc2f.
//
// Solidity: function getIssued(address issuer) view returns(uint256[])
func (_NFT *NFTCaller) GetIssued(opts *bind.CallOpts, issuer common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "getIssued", issuer)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetIssued is a free data retrieval call binding the contract method 0xe895bc2f.
//
// Solidity: function getIssued(address issuer) view returns(uint256[])
func (_NFT *NFTSession) GetIssued(issuer common.Address) ([]*big.Int, error) {
	return _NFT.Contract.GetIssued(&_NFT.CallOpts, issuer)
}

// GetIssued is a free data retrieval call binding the contract method 0xe895bc2f.
//
// Solidity: function getIssued(address issuer) view returns(uint256[])
func (_NFT *NFTCallerSession) GetIssued(issuer common.Address) ([]*big.Int, error) {
	return _NFT.Contract.GetIssued(&_NFT.CallOpts, issuer)
}

// GetWalletTokens is a free data retrieval call binding the contract method 0xfabc38be.
//
// Solidity: function getWalletTokens(address holder) view returns(uint256[])
func (_NFT *NFTCaller) GetWalletTokens(opts *bind.CallOpts, holder common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "getWalletTokens", holder)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetWalletTokens is a free data retrieval call binding the contract method 0xfabc38be.
//
// Solidity: function getWalletTokens(address holder) view returns(uint256[])
func (_NFT *NFTSession) GetWalletTokens(holder common.Address) ([]*big.Int, error) {
	return _NFT.Contract.GetWalletTokens(&_NFT.CallOpts, holder)
}

// GetWalletTokens is a free data retrieval call binding the contract method 0xfabc38be.
//
// Solidity: function getWalletTokens(address holder) view returns(uint256[])
func (_NFT *NFTCallerSession) GetWalletTokens(holder common.Address) ([]*big.Int, error) {
	return _NFT.Contract.GetWalletTokens(&_NFT.CallOpts, holder)
}

// IssuerTokenTypes is a free data retrieval call binding the contract method 0x18130d4f.
//
// Solidity: function issuerTokenTypes(address ) view returns(uint256 count)
func (_NFT *NFTCaller) IssuerTokenTypes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "issuerTokenTypes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IssuerTokenTypes is a free data retrieval call binding the contract method 0x18130d4f.
//
// Solidity: function issuerTokenTypes(address ) view returns(uint256 count)
func (_NFT *NFTSession) IssuerTokenTypes(arg0 common.Address) (*big.Int, error) {
	return _NFT.Contract.IssuerTokenTypes(&_NFT.CallOpts, arg0)
}

// IssuerTokenTypes is a free data retrieval call binding the contract method 0x18130d4f.
//
// Solidity: function issuerTokenTypes(address ) view returns(uint256 count)
func (_NFT *NFTCallerSession) IssuerTokenTypes(arg0 common.Address) (*big.Int, error) {
	return _NFT.Contract.IssuerTokenTypes(&_NFT.CallOpts, arg0)
}

// TokenHolder is a free data retrieval call binding the contract method 0x862a4bf2.
//
// Solidity: function tokenHolder(uint256 ) view returns(address)
func (_NFT *NFTCaller) TokenHolder(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenHolder", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenHolder is a free data retrieval call binding the contract method 0x862a4bf2.
//
// Solidity: function tokenHolder(uint256 ) view returns(address)
func (_NFT *NFTSession) TokenHolder(arg0 *big.Int) (common.Address, error) {
	return _NFT.Contract.TokenHolder(&_NFT.CallOpts, arg0)
}

// TokenHolder is a free data retrieval call binding the contract method 0x862a4bf2.
//
// Solidity: function tokenHolder(uint256 ) view returns(address)
func (_NFT *NFTCallerSession) TokenHolder(arg0 *big.Int) (common.Address, error) {
	return _NFT.Contract.TokenHolder(&_NFT.CallOpts, arg0)
}

// TokenMetadataURI is a free data retrieval call binding the contract method 0x157c3df9.
//
// Solidity: function tokenMetadataURI(uint256 ) view returns(string)
func (_NFT *NFTCaller) TokenMetadataURI(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenMetadataURI", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenMetadataURI is a free data retrieval call binding the contract method 0x157c3df9.
//
// Solidity: function tokenMetadataURI(uint256 ) view returns(string)
func (_NFT *NFTSession) TokenMetadataURI(arg0 *big.Int) (string, error) {
	return _NFT.Contract.TokenMetadataURI(&_NFT.CallOpts, arg0)
}

// TokenMetadataURI is a free data retrieval call binding the contract method 0x157c3df9.
//
// Solidity: function tokenMetadataURI(uint256 ) view returns(string)
func (_NFT *NFTCallerSession) TokenMetadataURI(arg0 *big.Int) (string, error) {
	return _NFT.Contract.TokenMetadataURI(&_NFT.CallOpts, arg0)
}

// TokenTokenType is a free data retrieval call binding the contract method 0xfd2da784.
//
// Solidity: function tokenTokenType(uint256 ) view returns(uint256)
func (_NFT *NFTCaller) TokenTokenType(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenTokenType", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenTokenType is a free data retrieval call binding the contract method 0xfd2da784.
//
// Solidity: function tokenTokenType(uint256 ) view returns(uint256)
func (_NFT *NFTSession) TokenTokenType(arg0 *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenTokenType(&_NFT.CallOpts, arg0)
}

// TokenTokenType is a free data retrieval call binding the contract method 0xfd2da784.
//
// Solidity: function tokenTokenType(uint256 ) view returns(uint256)
func (_NFT *NFTCallerSession) TokenTokenType(arg0 *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenTokenType(&_NFT.CallOpts, arg0)
}

// TokenTypeAuthorizedMinters is a free data retrieval call binding the contract method 0xca0cdc24.
//
// Solidity: function tokenTypeAuthorizedMinters(uint256 , address ) view returns(bool)
func (_NFT *NFTCaller) TokenTypeAuthorizedMinters(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenTypeAuthorizedMinters", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenTypeAuthorizedMinters is a free data retrieval call binding the contract method 0xca0cdc24.
//
// Solidity: function tokenTypeAuthorizedMinters(uint256 , address ) view returns(bool)
func (_NFT *NFTSession) TokenTypeAuthorizedMinters(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _NFT.Contract.TokenTypeAuthorizedMinters(&_NFT.CallOpts, arg0, arg1)
}

// TokenTypeAuthorizedMinters is a free data retrieval call binding the contract method 0xca0cdc24.
//
// Solidity: function tokenTypeAuthorizedMinters(uint256 , address ) view returns(bool)
func (_NFT *NFTCallerSession) TokenTypeAuthorizedMinters(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _NFT.Contract.TokenTypeAuthorizedMinters(&_NFT.CallOpts, arg0, arg1)
}

// TokenTypeMetadataURI is a free data retrieval call binding the contract method 0xa73d4ec3.
//
// Solidity: function tokenTypeMetadataURI(uint256 ) view returns(string)
func (_NFT *NFTCaller) TokenTypeMetadataURI(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenTypeMetadataURI", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenTypeMetadataURI is a free data retrieval call binding the contract method 0xa73d4ec3.
//
// Solidity: function tokenTypeMetadataURI(uint256 ) view returns(string)
func (_NFT *NFTSession) TokenTypeMetadataURI(arg0 *big.Int) (string, error) {
	return _NFT.Contract.TokenTypeMetadataURI(&_NFT.CallOpts, arg0)
}

// TokenTypeMetadataURI is a free data retrieval call binding the contract method 0xa73d4ec3.
//
// Solidity: function tokenTypeMetadataURI(uint256 ) view returns(string)
func (_NFT *NFTCallerSession) TokenTypeMetadataURI(arg0 *big.Int) (string, error) {
	return _NFT.Contract.TokenTypeMetadataURI(&_NFT.CallOpts, arg0)
}

// TokenTypeMintCount is a free data retrieval call binding the contract method 0x2d4d5da3.
//
// Solidity: function tokenTypeMintCount(uint256 ) view returns(uint256)
func (_NFT *NFTCaller) TokenTypeMintCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenTypeMintCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenTypeMintCount is a free data retrieval call binding the contract method 0x2d4d5da3.
//
// Solidity: function tokenTypeMintCount(uint256 ) view returns(uint256)
func (_NFT *NFTSession) TokenTypeMintCount(arg0 *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenTypeMintCount(&_NFT.CallOpts, arg0)
}

// TokenTypeMintCount is a free data retrieval call binding the contract method 0x2d4d5da3.
//
// Solidity: function tokenTypeMintCount(uint256 ) view returns(uint256)
func (_NFT *NFTCallerSession) TokenTypeMintCount(arg0 *big.Int) (*big.Int, error) {
	return _NFT.Contract.TokenTypeMintCount(&_NFT.CallOpts, arg0)
}

// TokenTypes is a free data retrieval call binding the contract method 0x33f6832a.
//
// Solidity: function tokenTypes(uint256 ) view returns(bool)
func (_NFT *NFTCaller) TokenTypes(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenTypes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenTypes is a free data retrieval call binding the contract method 0x33f6832a.
//
// Solidity: function tokenTypes(uint256 ) view returns(bool)
func (_NFT *NFTSession) TokenTypes(arg0 *big.Int) (bool, error) {
	return _NFT.Contract.TokenTypes(&_NFT.CallOpts, arg0)
}

// TokenTypes is a free data retrieval call binding the contract method 0x33f6832a.
//
// Solidity: function tokenTypes(uint256 ) view returns(bool)
func (_NFT *NFTCallerSession) TokenTypes(arg0 *big.Int) (bool, error) {
	return _NFT.Contract.TokenTypes(&_NFT.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(bool)
func (_NFT *NFTCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(bool)
func (_NFT *NFTSession) Tokens(arg0 *big.Int) (bool, error) {
	return _NFT.Contract.Tokens(&_NFT.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(bool)
func (_NFT *NFTCallerSession) Tokens(arg0 *big.Int) (bool, error) {
	return _NFT.Contract.Tokens(&_NFT.CallOpts, arg0)
}

// Wallets is a free data retrieval call binding the contract method 0x89b08f11.
//
// Solidity: function wallets(address ) view returns(uint256 count)
func (_NFT *NFTCaller) Wallets(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "wallets", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Wallets is a free data retrieval call binding the contract method 0x89b08f11.
//
// Solidity: function wallets(address ) view returns(uint256 count)
func (_NFT *NFTSession) Wallets(arg0 common.Address) (*big.Int, error) {
	return _NFT.Contract.Wallets(&_NFT.CallOpts, arg0)
}

// Wallets is a free data retrieval call binding the contract method 0x89b08f11.
//
// Solidity: function wallets(address ) view returns(uint256 count)
func (_NFT *NFTCallerSession) Wallets(arg0 common.Address) (*big.Int, error) {
	return _NFT.Contract.Wallets(&_NFT.CallOpts, arg0)
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 tokenID) returns()
func (_NFT *NFTTransactor) BurnToken(opts *bind.TransactOpts, tokenID *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "burnToken", tokenID)
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 tokenID) returns()
func (_NFT *NFTSession) BurnToken(tokenID *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.BurnToken(&_NFT.TransactOpts, tokenID)
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 tokenID) returns()
func (_NFT *NFTTransactorSession) BurnToken(tokenID *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.BurnToken(&_NFT.TransactOpts, tokenID)
}

// CreateTokenType is a paid mutator transaction binding the contract method 0x9799e525.
//
// Solidity: function createTokenType(uint256 tokenTypeID, address authorizedMinter, string metadataURI) returns()
func (_NFT *NFTTransactor) CreateTokenType(opts *bind.TransactOpts, tokenTypeID *big.Int, authorizedMinter common.Address, metadataURI string) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "createTokenType", tokenTypeID, authorizedMinter, metadataURI)
}

// CreateTokenType is a paid mutator transaction binding the contract method 0x9799e525.
//
// Solidity: function createTokenType(uint256 tokenTypeID, address authorizedMinter, string metadataURI) returns()
func (_NFT *NFTSession) CreateTokenType(tokenTypeID *big.Int, authorizedMinter common.Address, metadataURI string) (*types.Transaction, error) {
	return _NFT.Contract.CreateTokenType(&_NFT.TransactOpts, tokenTypeID, authorizedMinter, metadataURI)
}

// CreateTokenType is a paid mutator transaction binding the contract method 0x9799e525.
//
// Solidity: function createTokenType(uint256 tokenTypeID, address authorizedMinter, string metadataURI) returns()
func (_NFT *NFTTransactorSession) CreateTokenType(tokenTypeID *big.Int, authorizedMinter common.Address, metadataURI string) (*types.Transaction, error) {
	return _NFT.Contract.CreateTokenType(&_NFT.TransactOpts, tokenTypeID, authorizedMinter, metadataURI)
}

// MintToken is a paid mutator transaction binding the contract method 0x87c6649c.
//
// Solidity: function mintToken(address to, uint256 tokenTypeID, uint256 tokenID, string metadataURI) returns()
func (_NFT *NFTTransactor) MintToken(opts *bind.TransactOpts, to common.Address, tokenTypeID *big.Int, tokenID *big.Int, metadataURI string) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "mintToken", to, tokenTypeID, tokenID, metadataURI)
}

// MintToken is a paid mutator transaction binding the contract method 0x87c6649c.
//
// Solidity: function mintToken(address to, uint256 tokenTypeID, uint256 tokenID, string metadataURI) returns()
func (_NFT *NFTSession) MintToken(to common.Address, tokenTypeID *big.Int, tokenID *big.Int, metadataURI string) (*types.Transaction, error) {
	return _NFT.Contract.MintToken(&_NFT.TransactOpts, to, tokenTypeID, tokenID, metadataURI)
}

// MintToken is a paid mutator transaction binding the contract method 0x87c6649c.
//
// Solidity: function mintToken(address to, uint256 tokenTypeID, uint256 tokenID, string metadataURI) returns()
func (_NFT *NFTTransactorSession) MintToken(to common.Address, tokenTypeID *big.Int, tokenID *big.Int, metadataURI string) (*types.Transaction, error) {
	return _NFT.Contract.MintToken(&_NFT.TransactOpts, to, tokenTypeID, tokenID, metadataURI)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 tokenID) returns()
func (_NFT *NFTTransactor) TransferToken(opts *bind.TransactOpts, to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "transferToken", to, tokenID)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 tokenID) returns()
func (_NFT *NFTSession) TransferToken(to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.TransferToken(&_NFT.TransactOpts, to, tokenID)
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 tokenID) returns()
func (_NFT *NFTTransactorSession) TransferToken(to common.Address, tokenID *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.TransferToken(&_NFT.TransactOpts, to, tokenID)
}

// NFTBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the NFT contract.
type NFTBurnIterator struct {
	Event *NFTBurn // Event containing the contract specifics and raw log

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
func (it *NFTBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTBurn)
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
		it.Event = new(NFTBurn)
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
func (it *NFTBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTBurn represents a Burn event raised by the NFT contract.
type NFTBurn struct {
	Authorizer common.Address
	TokenID    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address authorizer, uint256 tokenID)
func (_NFT *NFTFilterer) FilterBurn(opts *bind.FilterOpts) (*NFTBurnIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return &NFTBurnIterator{contract: _NFT.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address authorizer, uint256 tokenID)
func (_NFT *NFTFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *NFTBurn) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Burn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTBurn)
				if err := _NFT.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address authorizer, uint256 tokenID)
func (_NFT *NFTFilterer) ParseBurn(log types.Log) (*NFTBurn, error) {
	event := new(NFTBurn)
	if err := _NFT.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the NFT contract.
type NFTMintIterator struct {
	Event *NFTMint // Event containing the contract specifics and raw log

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
func (it *NFTMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTMint)
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
		it.Event = new(NFTMint)
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
func (it *NFTMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTMint represents a Mint event raised by the NFT contract.
type NFTMint struct {
	To          common.Address
	TokenTypeID *big.Int
	TokenID     *big.Int
	Uri         string
	MintCount   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x397c4cd4ff46491382eefda34299d046406c58377a808e6eb8a0f0e892ec8bd1.
//
// Solidity: event Mint(address to, uint256 tokenTypeID, uint256 tokenID, string uri, uint256 mintCount)
func (_NFT *NFTFilterer) FilterMint(opts *bind.FilterOpts) (*NFTMintIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &NFTMintIterator{contract: _NFT.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x397c4cd4ff46491382eefda34299d046406c58377a808e6eb8a0f0e892ec8bd1.
//
// Solidity: event Mint(address to, uint256 tokenTypeID, uint256 tokenID, string uri, uint256 mintCount)
func (_NFT *NFTFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *NFTMint) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTMint)
				if err := _NFT.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x397c4cd4ff46491382eefda34299d046406c58377a808e6eb8a0f0e892ec8bd1.
//
// Solidity: event Mint(address to, uint256 tokenTypeID, uint256 tokenID, string uri, uint256 mintCount)
func (_NFT *NFTFilterer) ParseMint(log types.Log) (*NFTMint, error) {
	event := new(NFTMint)
	if err := _NFT.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTNewTokenTypeIterator is returned from FilterNewTokenType and is used to iterate over the raw logs and unpacked data for NewTokenType events raised by the NFT contract.
type NFTNewTokenTypeIterator struct {
	Event *NFTNewTokenType // Event containing the contract specifics and raw log

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
func (it *NFTNewTokenTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTNewTokenType)
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
		it.Event = new(NFTNewTokenType)
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
func (it *NFTNewTokenTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTNewTokenTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTNewTokenType represents a NewTokenType event raised by the NFT contract.
type NFTNewTokenType struct {
	Issuer common.Address
	Id     *big.Int
	Uri    string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewTokenType is a free log retrieval operation binding the contract event 0x8d201fa1da39f4c36d5f05fc422ac3a4d64121f878587926933a1ff90fc33ebe.
//
// Solidity: event NewTokenType(address issuer, uint256 id, string uri)
func (_NFT *NFTFilterer) FilterNewTokenType(opts *bind.FilterOpts) (*NFTNewTokenTypeIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "NewTokenType")
	if err != nil {
		return nil, err
	}
	return &NFTNewTokenTypeIterator{contract: _NFT.contract, event: "NewTokenType", logs: logs, sub: sub}, nil
}

// WatchNewTokenType is a free log subscription operation binding the contract event 0x8d201fa1da39f4c36d5f05fc422ac3a4d64121f878587926933a1ff90fc33ebe.
//
// Solidity: event NewTokenType(address issuer, uint256 id, string uri)
func (_NFT *NFTFilterer) WatchNewTokenType(opts *bind.WatchOpts, sink chan<- *NFTNewTokenType) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "NewTokenType")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTNewTokenType)
				if err := _NFT.contract.UnpackLog(event, "NewTokenType", log); err != nil {
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

// ParseNewTokenType is a log parse operation binding the contract event 0x8d201fa1da39f4c36d5f05fc422ac3a4d64121f878587926933a1ff90fc33ebe.
//
// Solidity: event NewTokenType(address issuer, uint256 id, string uri)
func (_NFT *NFTFilterer) ParseNewTokenType(log types.Log) (*NFTNewTokenType, error) {
	event := new(NFTNewTokenType)
	if err := _NFT.contract.UnpackLog(event, "NewTokenType", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NFT contract.
type NFTTransferIterator struct {
	Event *NFTTransfer // Event containing the contract specifics and raw log

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
func (it *NFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTTransfer)
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
		it.Event = new(NFTTransfer)
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
func (it *NFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTTransfer represents a Transfer event raised by the NFT contract.
type NFTTransfer struct {
	To      common.Address
	TokenID *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address to, uint256 tokenID)
func (_NFT *NFTFilterer) FilterTransfer(opts *bind.FilterOpts) (*NFTTransferIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &NFTTransferIterator{contract: _NFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address to, uint256 tokenID)
func (_NFT *NFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NFTTransfer) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTTransfer)
				if err := _NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x69ca02dd4edd7bf0a4abb9ed3b7af3f14778db5d61921c7dc7cd545266326de2.
//
// Solidity: event Transfer(address to, uint256 tokenID)
func (_NFT *NFTFilterer) ParseTransfer(log types.Log) (*NFTTransfer, error) {
	event := new(NFTTransfer)
	if err := _NFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
