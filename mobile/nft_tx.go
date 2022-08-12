package geth

// Ask is a paid mutator transaction binding the contract method 0x72db88f5.
//
// Solidity: function ask(uint256 askID, uint256 auctionID, uint256 tokenID, uint256 amount) returns()
func (_Nft *NFT) Ask(
	opts *TransactOpts,
	askID *BigInt,
	auctionID *BigInt,
	tokenID *BigInt,
	amount *BigInt,
) (*Transaction, error) {
	transaction, err := _Nft.contract.Transact(
		&opts.opts,
		"ask",
		askID.bigint,
		auctionID.bigint,
		tokenID.bigint,
		amount.bigint,
	)
	if err != nil {
		return *new(*Transaction), err
	}
	return &Transaction{tx: transaction}, err
}

// Bid is a paid mutator transaction binding the contract method 0x2ac9bf09.
//
// Solidity: function bid(uint256 bidID, uint256 auctionID, uint256 amount) payable returns()
func (_Nft *NFT) Bid(opts *TransactOpts, bidID *BigInt, auctionID *BigInt, amount *BigInt) (*Transaction, error) {
	transaction, err := _Nft.contract.Transact(&opts.opts, "bid", bidID.bigint, auctionID.bigint, amount.bigint)
	if err != nil {
		return *new(*Transaction), err
	}
	return &Transaction{tx: transaction}, err
}

// BurnToken is a paid mutator transaction binding the contract method 0x7b47ec1a.
//
// Solidity: function burnToken(uint256 tokenID) returns()
func (_Nft *NFT) BurnToken(opts *TransactOpts, tokenID *BigInt) (*Transaction, error) {
	transaction, err := _Nft.contract.Transact(&opts.opts, "burnToken", tokenID.bigint)
	if err != nil {
		return *new(*Transaction), err
	}
	return &Transaction{tx: transaction}, err
}

// CancelAsk is a paid mutator transaction binding the contract method 0x1628e0f5.
//
// Solidity: function cancelAsk(uint256 askID) returns()
func (_Nft *NFT) CancelAsk(opts *TransactOpts, askID *BigInt) (*Transaction, error) {
	transaction, err := _Nft.contract.Transact(&opts.opts, "cancelAsk", askID.bigint)

	if err != nil {
		return *new(*Transaction), err
	}
	return &Transaction{tx: transaction}, err
}

// CancelBid is a paid mutator transaction binding the contract method 0x9703ef35.
//
// Solidity: function cancelBid(uint256 bidID) returns()
func (_Nft *NFT) CancelBid(opts *TransactOpts, bidID *BigInt) (*Transaction, error) {
	transaction, err := _Nft.contract.Transact(&opts.opts, "cancelBid", bidID.bigint)

	if err != nil {
		return *new(*Transaction), err
	}
	return &Transaction{tx: transaction}, err
}

// CloseAuction is a paid mutator transaction binding the contract method 0x236ed8f3.
//
// Solidity: function closeAuction(uint256 auctionID) returns()
func (_Nft *NFT) CloseAuction(opts *TransactOpts, auctionID *BigInt) (*Transaction, error) {
	transaction, err := _Nft.contract.Transact(&opts.opts, "closeAuction", auctionID.bigint)

	if err != nil {
		return *new(*Transaction), err
	}
	return &Transaction{tx: transaction}, err
}

// CreateTokenType is a paid mutator transaction binding the contract method 0x86c5c007.
//
func (_Nft *NFT) CreateTokenType(
	opts *TransactOpts,
	tokenTypeID *BigInt,
	metadataURI string,
	transferTakeRate *BigInt,
	auctionID *BigInt,
	orgID string,
) (*Transaction, error) {
	transaction, err := _Nft.contract.Transact(
		&opts.opts,
		"createTokenType",
		tokenTypeID.bigint,
		metadataURI,
		transferTakeRate.bigint,
		auctionID.bigint,
		orgID,
	)

	if err != nil {
		return *new(*Transaction), err
	}
	return &Transaction{tx: transaction}, err
}

// MintToken is a paid mutator transaction binding the contract method 0x87c6649c.
//
// Solidity: function mintToken(address to, uint256 tokenTypeID, uint256 tokenID, string metadataURI) returns()
func (_Nft *NFT) MintToken(
	opts *TransactOpts,
	to *Address,
	tokenTypeID *BigInt,
	tokenID *BigInt,
	metadataURI string,
) (*Transaction, error) {
	transaction, err := _Nft.contract.Transact(
		&opts.opts,
		"mintToken",
		to.address,
		tokenTypeID.bigint,
		tokenID.bigint,
		metadataURI,
	)

	if err != nil {
		return *new(*Transaction), err
	}
	return &Transaction{tx: transaction}, err
}

// TransferToken is a paid mutator transaction binding the contract method 0x1072cbea.
//
// Solidity: function transferToken(address to, uint256 tokenID) returns()
func (_Nft *NFT) TransferToken(opts *TransactOpts, to *Address, tokenID *BigInt) (*Transaction, error) {
	transaction, err := _Nft.contract.Transact(&opts.opts, "transferToken", to.address, tokenID.bigint)

	if err != nil {
		return *new(*Transaction), err
	}
	return &Transaction{tx: transaction}, err
}
