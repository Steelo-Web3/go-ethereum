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

// NftABI is the input ABI used to generate the binding from
var NftABI = NftMetaData.ABI

// NFT is an auto generated Go binding around an Ethereum contract.
type NFT struct {
	contract *bind.BoundContract
}

// NFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTSession struct {
	contract     *NFT         // Generic contract binding to set the session for
	callOpts     CallOpts     // Call options to use throughout this session
	transactOpts TransactOpts // Transaction auth options to use throughout this session
}

// NewNFT creates a new instance of NFT, bound to a specific deployed contract.
func NewNFT(address *Address, client *EthereumClient) (*NFT, error) {
	contract, err := bindNFT(address.address, client.client, client.client, client.client)
	if err != nil {
		return nil, err
	}
	return &NFT{contract: contract}, nil
}

// bindNFT binds a generic wrapper to an already deployed contract.
func bindNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFT) Call(opts *CallOpts, result *Interfaces, method string, params *Interfaces) error {
	return _NFT.contract.Call(&opts.opts, &result.objects, method, params)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFT) Transfer(opts *TransactOpts) (*Transaction, error) {
	tx, err := _NFT.contract.Transfer(&opts.opts)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		tx: tx,
	}, nil
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFT) Transact(opts *TransactOpts, method string, params *Interfaces) (*Transaction, error) {
	tx, err := _NFT.contract.Transact(&opts.opts, method, params.objects)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		tx: tx,
	}, nil
}

// GetIssued is a free data retrieval call binding the contract method 0xe895bc2f.
//
// Solidity: function getIssued(address issuer) view returns(uint256[])
func (_NFT *NFT) GetIssued(opts *CallOpts, issuer *Address) (*BigInts, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "getIssued", issuer.address)

	if err != nil {
		return *new(*BigInts), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return &BigInts{bigints: out0}, nil

}

// GetWalletTokens is a free data retrieval call binding the contract method 0xfabc38be.
//
// Solidity: function getWalletTokens(address holder) view returns(uint256[])
func (_NFT *NFT) GetWalletTokens(opts *CallOpts, holder *Address) (*BigInts, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "getWalletTokens", holder.address)

	if err != nil {
		return *new(*BigInts), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return &BigInts{bigints: out0}, nil

}

// IssuerTokenTypes is a free data retrieval call binding the contract method 0x18130d4f.
//
// Solidity: function issuerTokenTypes(address ) view returns(uint256 count)
func (_NFT *NFT) IssuerTokenTypes(opts *CallOpts, arg0 *Address) (*BigInt, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "issuerTokenTypes", arg0.address)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, nil

}

// Wallets is a free data retrieval call binding the contract method 0x89b08f11.
//
// Solidity: function wallets(address ) view returns(uint256 count)
func (_NFT *NFT) Wallets(opts *CallOpts, arg0 Address) (*BigInt, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "wallets", arg0.address)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, nil
}

// TokenHolder is a free data retrieval call binding the contract method 0x862a4bf2.
//
// Solidity: function tokenHolder(uint256 ) view returns(address)
func (_NFT *NFT) TokenHolder(opts *CallOpts, arg0 *BigInt) (*Address, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "tokenHolder", arg0.bigint)

	if err != nil {
		return *new(*Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return &Address{address: out0}, err

}

// TokenTokenType is a free data retrieval call binding the contract method 0xfd2da784.
//
// Solidity: function tokenTokenType(uint256 ) view returns(uint256)
func (_NFT *NFT) TokenTokenType(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "tokenTokenType", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err

}

// TokenTypeAuthorizedMinters is a free data retrieval call binding the contract method 0xca0cdc24.
//
// Solidity: function tokenTypeAuthorizedMinters(uint256 , address ) view returns(bool)
func (_NFT *NFT) TokenTypeAuthorizedMinters(opts *CallOpts, arg0 *BigInt, arg1 *Address) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "tokenTypeAuthorizedMinters", arg0.bigint, arg1.address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenTypeURI is a free data retrieval call binding the contract method 0x885a72a9.
//
// Solidity: function tokenTypeURI(uint256 ) view returns(string)
func (_NFT *NFT) TokenTypeURI(opts *CallOpts, arg0 *BigInt) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "tokenTypeURI", arg0.bigint)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenTypes is a free data retrieval call binding the contract method 0x33f6832a.
//
// Solidity: function tokenTypes(uint256 ) view returns(bool)
func (_NFT *NFT) TokenTypes(opts *CallOpts, arg0 *BigInt) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "tokenTypes", arg0.bigint)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 ) view returns(string)
func (_NFT *NFT) TokenURI(opts *CallOpts, arg0 *BigInt) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "tokenURI", arg0.bigint)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenTypeMintCount is a free data retrieval call binding the contract method 0x2d4d5da3.
//
// Solidity: function tokenTypeMintCount(uint256 ) view returns(uint256)
func (_NFT *NFT) TokenTypeMintCount(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "tokenTypeMintCount", arg0)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{
		bigint: out0,
	}, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(bool)
func (_NFT *NFT) Tokens(opts *CallOpts, arg0 *BigInt) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(&opts.opts, &out, "tokens", arg0.bigint)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreateTokenType is a paid mutator transaction binding the contract method 0x9799e525.
//
// Solidity: function createTokenType(uint256 tokenTypeID, address authorizedMinter, string uri) returns()
func (_NFT *NFT) CreateTokenType(opts *TransactOpts, tokenTypeID *BigInt, authorizedMinter *Address, uri string) (*Transaction, error) {
	tx, err := _NFT.contract.Transact(&opts.opts, "createTokenType", tokenTypeID.bigint, authorizedMinter.address, uri)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx: tx}, nil
}

// MintToken is a paid mutator transaction binding the contract method 0x87c6649c.
//
// Solidity: function mintToken(address to, uint256 tokenTypeID, uint256 tokenID, string uri) returns()
func (_NFT *NFT) MintToken(opts *TransactOpts, to *Address, tokenTypeID *BigInt, tokenID *BigInt, uri string) (*Transaction, error) {
	tx, err := _NFT.contract.Transact(&opts.opts, "mintToken", to.address, tokenTypeID.bigint, tokenID.bigint, uri)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx: tx}, nil
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 tokenID) returns()
func (_NFT *NFT) TransferToken(opts *TransactOpts, to *Address, tokenID *BigInt) (*Transaction, error) {
	tx, err := _NFT.contract.Transact(&opts.opts, "transferToken", to.address, tokenID.bigint)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx: tx}, nil
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 tokenID) returns()
func (_NFT *NFT) BurnToken(opts *TransactOpts, tokenID *BigInt) (*Transaction, error) {
	tx, err := _NFT.contract.Transact(&opts.opts, "burnToken", tokenID.bigint)
	if err != nil {
		return nil, err
	}
	return &Transaction{tx: tx}, nil
}
