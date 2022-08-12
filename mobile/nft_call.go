package geth

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_Nft *NFT) Owner(opts *CallOpts) (*Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "_owner")

	if err != nil {
		return *new(*Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return &Address{address: out0}, err
}

// AddressAsks is a free data retrieval call binding the contract method 0xf4f86d5d.
//
// Solidity: function addressAsks(address ) view returns(address asker)
func (_Nft *NFT) AddressAsks(opts *CallOpts, arg0 *Address) (*Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "addressAsks", arg0.address)

	if err != nil {
		return *new(*Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return &Address{address: out0}, err
}

// AskAmount is a free data retrieval call binding the contract method 0xe3bc3864.
//
// Solidity: function askAmount(uint256 ) view returns(uint256)
func (_Nft *NFT) AskAmount(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "askAmount", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// AskAuction is a free data retrieval call binding the contract method 0x80420938.
//
// Solidity: function askAuction(uint256 ) view returns(uint256)
func (_Nft *NFT) AskAuction(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "askAuction", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// AskOwner is a free data retrieval call binding the contract method 0xc7303945.
//
// Solidity: function askOwner(uint256 ) view returns(address)
func (_Nft *NFT) AskOwner(opts *CallOpts, arg0 *BigInt) (*Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "askOwner", arg0.bigint)

	if err != nil {
		return *new(*Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return &Address{address: out0}, err

}

// AskToken is a free data retrieval call binding the contract method 0xa8726716.
//
// Solidity: function askToken(uint256 ) view returns(uint256)
func (_Nft *NFT) AskToken(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "askToken", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// AuctionPrice is a free data retrieval call binding the contract method 0x6980fd6a.
//
// Solidity: function auctionPrice(uint256 ) view returns(uint256)
func (_Nft *NFT) AuctionPrice(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "auctionPrice", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// AuctionTokenType is a free data retrieval call binding the contract method 0x151c7e67.
//
// Solidity: function auctionTokenType(uint256 ) view returns(uint256)
func (_Nft *NFT) AuctionTokenType(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "auctionTokenType", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// BidAmount is a free data retrieval call binding the contract method 0x48d9be16.
//
// Solidity: function bidAmount(uint256 ) view returns(uint256)
func (_Nft *NFT) BidAmount(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "bidAmount", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// BidAuction is a free data retrieval call binding the contract method 0x64a7d7c7.
//
// Solidity: function bidAuction(uint256 ) view returns(uint256)
func (_Nft *NFT) BidAuction(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "bidAuction", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// BidOwner is a free data retrieval call binding the contract method 0xa3615bc0.
//
// Solidity: function bidOwner(uint256 ) view returns(address)
func (_Nft *NFT) BidOwner(opts *CallOpts, arg0 *BigInt) (*Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "bidOwner", arg0.bigint)

	if err != nil {
		return *new(*Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return &Address{address: out0}, err
}

// GetAddressAsks is a free data retrieval call binding the contract method 0xe46117d4.
//
// Solidity: function getAddressAsks(address asker) view returns(uint256[])
func (_Nft *NFT) GetAddressAsks(opts *CallOpts, asker *Address) (*BigInts, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "getAddressAsks", asker.address)

	if err != nil {
		return *new(*BigInts), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return &BigInts{bigints: out0}, err
}

// GetAddressBids is a free data retrieval call binding the contract method 0xd9ea9549.
//
// Solidity: function getAddressBids(address bidder) view returns(uint256[])
func (_Nft *NFT) GetAddressBids(opts *CallOpts, bidder *Address) (*BigInts, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "getAddressBids", bidder.address)

	if err != nil {
		return *new(*BigInts), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return &BigInts{bigints: out0}, err
}

// GetAuctionAsks is a free data retrieval call binding the contract method 0xf6b712b6.
//
// Solidity: function getAuctionAsks(uint256 auctionID) view returns(uint256[])
func (_Nft *NFT) GetAuctionAsks(opts *CallOpts, auctionID *BigInt) (*BigInts, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "getAuctionAsks", auctionID.bigint)

	if err != nil {
		return *new(*BigInts), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return &BigInts{bigints: out0}, err
}

// GetAuctionBids is a free data retrieval call binding the contract method 0x70b4768e.
//
// Solidity: function getAuctionBids(uint256 auctionID) view returns(uint256[])
func (_Nft *NFT) GetAuctionBids(opts *CallOpts, auctionID *BigInt) (*BigInts, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "getAuctionBids", auctionID.bigint)

	if err != nil {
		return *new(*BigInts), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return &BigInts{bigints: out0}, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Nft *NFT) GetBalance(opts *CallOpts) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "getBalance")

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// GetHighestBid is a free data retrieval call binding the contract method 0x8f288644.
//
// Solidity: function getHighestBid(uint256 auctionID) view returns(uint256)
func (_Nft *NFT) GetHighestBid(opts *CallOpts, auctionID *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "getHighestBid", auctionID.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// GetIssued is a free data retrieval call binding the contract method 0xe895bc2f.
//
// Solidity: function getIssued(address issuer) view returns(uint256[])
func (_Nft *NFT) GetIssued(opts *CallOpts, issuer *Address) (*BigInts, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "getIssued", issuer.address)

	if err != nil {
		return *new(*BigInts), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return &BigInts{bigints: out0}, err
}

// GetLowestAsk is a free data retrieval call binding the contract method 0x793052b8.
//
// Solidity: function getLowestAsk(uint256 auctionID) view returns(uint256)
func (_Nft *NFT) GetLowestAsk(opts *CallOpts, auctionID *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "getLowestAsk", auctionID.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// GetWalletTokens is a free data retrieval call binding the contract method 0xfabc38be.
//
// Solidity: function getWalletTokens(address owner) view returns(uint256[])
func (_Nft *NFT) GetWalletTokens(opts *CallOpts, owner *Address) (*BigInts, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "getWalletTokens", owner.address)

	if err != nil {
		return *new(*BigInts), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return &BigInts{bigints: out0}, err
}

// TokenHolder is a free data retrieval call binding the contract method 0x862a4bf2.
//
// Solidity: function tokenHolder(uint256 ) view returns(address)
func (_Nft *NFT) TokenHolder(opts *CallOpts, arg0 *BigInt) (*Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "tokenHolder", arg0.bigint)

	if err != nil {
		return *new(*Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return &Address{address: out0}, err
}

// TokenMetadataURI is a free data retrieval call binding the contract method 0x157c3df9.
//
// Solidity: function tokenMetadataURI(uint256 ) view returns(string)
func (_Nft *NFT) TokenMetadataURI(opts *CallOpts, arg0 *BigInt) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "tokenMetadataURI", arg0.bigint)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenTokenType is a free data retrieval call binding the contract method 0xfd2da784.
//
// Solidity: function tokenTokenType(uint256 ) view returns(uint256)
func (_Nft *NFT) TokenTokenType(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "tokenTokenType", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// TokenTypeAuthorizedMinters is a free data retrieval call binding the contract method 0xca0cdc24.
//
// Solidity: function tokenTypeAuthorizedMinters(uint256 , address ) view returns(bool)
func (_Nft *NFT) TokenTypeAuthorizedMinters(opts *CallOpts, arg0 *BigInt, arg1 *Address) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "tokenTypeAuthorizedMinters", arg0.bigint, arg1.address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// TokenTypeIssuer is a free data retrieval call binding the contract method 0xb281b5b6.
//
// Solidity: function tokenTypeIssuer(uint256 ) view returns(address)
func (_Nft *NFT) TokenTypeIssuer(opts *CallOpts, arg0 *BigInt) (*Address, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "tokenTypeIssuer", arg0.bigint)

	if err != nil {
		return *new(*Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return &Address{address: out0}, err
}

// TokenTypeIssuerTransferTakeRate is a free data retrieval call binding the contract method 0xe36c9b34.
//
// Solidity: function tokenTypeIssuerTransferTakeRate(uint256 ) view returns(uint256)
func (_Nft *NFT) TokenTypeIssuerTransferTakeRate(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "tokenTypeIssuerTransferTakeRate", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// TokenTypeMetadataURI is a free data retrieval call binding the contract method 0xa73d4ec3.
//
// Solidity: function tokenTypeMetadataURI(uint256 ) view returns(string)
func (_Nft *NFT) TokenTypeMetadataURI(opts *CallOpts, arg0 *BigInt) (string, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "tokenTypeMetadataURI", arg0.bigint)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

// TokenTypeMintCount is a free data retrieval call binding the contract method 0x2d4d5da3.
//
// Solidity: function tokenTypeMintCount(uint256 ) view returns(uint256)
func (_Nft *NFT) TokenTypeMintCount(opts *CallOpts, arg0 *BigInt) (*BigInt, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "tokenTypeMintCount", arg0.bigint)

	if err != nil {
		return *new(*BigInt), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return &BigInt{bigint: out0}, err
}

// TokenTypeTransferable is a free data retrieval call binding the contract method 0x7f555e52.
//
// Solidity: function tokenTypeTransferable(uint256 ) view returns(bool)
func (_Nft *NFT) TokenTypeTransferable(opts *CallOpts, arg0 *BigInt) (bool, error) {
	var out []interface{}
	err := _Nft.contract.Call(&opts.opts, &out, "tokenTypeTransferable", arg0.bigint)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}
