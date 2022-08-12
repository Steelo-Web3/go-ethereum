package geth

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
)

// NftAskIterator is returned from FilterAsk and is used to iterate over the raw logs and unpacked data for Ask events raised by the Nft contract.
type NftAskIterator struct {
	Event *NftAsk // Event containing the contract specifics and raw log

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
func (it *NftAskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftAsk)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftAsk)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftAskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftAskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftAsk represents a Ask event raised by the Nft contract.
type NftAsk struct {
	AuctionID *BigInt
	Amount    *BigInt
	Asker     Address
	TokenID   *BigInt
	AskID     *BigInt
	Raw       Log // Blockchain specific contextual infos
}

// FilterAsk is a free log retrieval operation binding the contract event 0x3fc45f0f85dbedbaa464bda774a9af14ee790eb25717ec4cd933f94fcdfc11e4.
//
// Solidity: event Ask(uint256 indexed auctionID, uint256 amount, address indexed asker, uint256 indexed tokenID, uint256 askID)
func (_Nft *NFT) FilterAsk(
	opts *FilterOpts,
	auctionID *BigInt,
	asker *Address,
	tokenID *BigInt,
) (*NftAskIterator, error) {
	var auctionIDRule []interface{}
	for _, auctionIDItem := range []*big.Int{auctionID.bigint} {
		auctionIDRule = append(auctionIDRule, *auctionIDItem)
	}

	var askerRule []interface{}
	for _, askerItem := range []common.Address{asker.address} {
		askerRule = append(askerRule, askerItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range []*big.Int{tokenID.bigint} {
		tokenIDRule = append(tokenIDRule, tokenIDItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "Ask", auctionIDRule, askerRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &NftAskIterator{contract: _Nft.contract, event: "Ask", logs: logs, sub: sub}, nil
}

// WatchAsk is a free log subscription operation binding the contract event 0x3fc45f0f85dbedbaa464bda774a9af14ee790eb25717ec4cd933f94fcdfc11e4.
//
// Solidity: event Ask(uint256 indexed auctionID, uint256 amount, address indexed asker, uint256 indexed tokenID, uint256 askID)
func (_Nft *NFT) WatchAsk(
	opts *WatchOpts,
	sink chan<- *NftAsk,
	auctionID *BigInt,
	asker *Address,
	tokenID *BigInt,
) (*Subscription, error) {
	var auctionIDRule []interface{}
	for _, auctionIDItem := range []*big.Int{auctionID.bigint} {
		auctionIDRule = append(auctionIDRule, *auctionIDItem)
	}

	var askerRule []interface{}
	for _, askerItem := range []common.Address{asker.address} {
		askerRule = append(askerRule, askerItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range []*big.Int{tokenID.bigint} {
		tokenIDRule = append(tokenIDRule, *tokenIDItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts.opts, "Ask", auctionIDRule, askerRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftAsk)
				if err := _Nft.contract.UnpackLog(event, "Ask", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})
	return &Subscription{sub: subscription}, nil
}

// ParseAsk is a log parse operation binding the contract event 0x3fc45f0f85dbedbaa464bda774a9af14ee790eb25717ec4cd933f94fcdfc11e4.
//
// Solidity: event Ask(uint256 indexed auctionID, uint256 amount, address indexed asker, uint256 indexed tokenID, uint256 askID)
func (_Nft *NFT) ParseAsk(log Log) (*NftAsk, error) {
	event := new(NftAsk)
	if err := _Nft.contract.UnpackLog(event, "Ask", *log.log); err != nil {
		return nil, err
	}
	event.Raw = Log{log: log.log}
	return event, nil
}

// NftBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Nft contract.
type NftBidIterator struct {
	Event *NftBid // Event containing the contract specifics and raw log

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
func (it *NftBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftBid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBid represents a Bid event raised by the Nft contract.
type NftBid struct {
	AuctionID *BigInt
	Amount    *BigInt
	Bidder    Address
	BidID     *BigInt
	Raw       Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xf657a33a260f0292d243ac4d7832a2af14d735a16351622191c6cf01af297c4e.
//
// Solidity: event Bid(uint256 indexed auctionID, uint256 amount, address indexed bidder, uint256 indexed bidID)
func (_Nft *NFT) FilterBid(
	opts *FilterOpts,
	auctionID *BigInt,
	bidder *Address,
	bidID *BigInt,
) (*NftBidIterator, error) {
	var auctionIDRule []interface{}
	for _, auctionIDItem := range []*big.Int{auctionID.bigint} {
		auctionIDRule = append(auctionIDRule, *auctionIDItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range []common.Address{bidder.address} {
		bidderRule = append(bidderRule, bidderItem)
	}
	var bidIDRule []interface{}
	for _, bidIDItem := range []*big.Int{bidID.bigint} {
		bidIDRule = append(bidIDRule, *bidIDItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "Bid", auctionIDRule, bidderRule, bidIDRule)
	if err != nil {
		return nil, err
	}
	return &NftBidIterator{contract: _Nft.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xf657a33a260f0292d243ac4d7832a2af14d735a16351622191c6cf01af297c4e.
//
// Solidity: event Bid(uint256 indexed auctionID, uint256 amount, address indexed bidder, uint256 indexed bidID)
func (_Nft *NFT) WatchBid(
	opts *WatchOpts,
	sink chan<- *NftBid,
	auctionID *BigInt,
	bidder *Address,
	bidID *BigInt,
) (*Subscription, error) {
	var auctionIDRule []interface{}
	for _, auctionIDItem := range []*big.Int{auctionID.bigint} {
		auctionIDRule = append(auctionIDRule, *auctionIDItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range []common.Address{bidder.address} {
		bidderRule = append(bidderRule, bidderItem)
	}
	var bidIDRule []interface{}
	for _, bidIDItem := range []*big.Int{bidID.bigint} {
		bidIDRule = append(bidIDRule, *bidIDItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts.opts, "Bid", auctionIDRule, bidderRule, bidIDRule)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBid)
				if err := _Nft.contract.UnpackLog(event, "Bid", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})

	return &Subscription{sub: subscription}, nil
}

// ParseBid is a log parse operation binding the contract event 0xf657a33a260f0292d243ac4d7832a2af14d735a16351622191c6cf01af297c4e.
//
// Solidity: event Bid(uint256 indexed auctionID, uint256 amount, address indexed bidder, uint256 indexed bidID)
func (_Nft *NFT) ParseBid(log Log) (*NftBid, error) {
	event := new(NftBid)
	if err := _Nft.contract.UnpackLog(event, "Bid", *log.log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Nft contract.
type NftBurnIterator struct {
	Event *NftBurn // Event containing the contract specifics and raw log

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
func (it *NftBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftBurn represents a Burn event raised by the Nft contract.
type NftBurn struct {
	Authorizer Address
	TokenID    *BigInt
	Raw        Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed authorizer, uint256 indexed tokenID)
func (_Nft *NFT) FilterBurn(
	opts *FilterOpts,
	authorizer *Address,
	tokenID *BigInt,
) (*NftBurnIterator, error) {
	var authorizerRule []interface{}
	for _, authorizerItem := range []common.Address{authorizer.address} {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range []*big.Int{tokenID.bigint} {
		tokenIDRule = append(tokenIDRule, *tokenIDItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "Burn", authorizerRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &NftBurnIterator{contract: _Nft.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed authorizer, uint256 indexed tokenID)
func (_Nft *NFT) WatchBurn(
	opts *WatchOpts,
	sink chan<- *NftBurn,
	authorizer *Address,
	tokenID *BigInt,
) (*Subscription, error) {
	var authorizerRule []interface{}
	for _, authorizerItem := range []common.Address{authorizer.address} {
		authorizerRule = append(authorizerRule, authorizerItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range []*big.Int{tokenID.bigint} {
		tokenIDRule = append(tokenIDRule, *tokenIDItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts.opts, "Burn", authorizerRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftBurn)
				if err := _Nft.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})
	return &Subscription{sub: subscription}, nil
}

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed authorizer, uint256 indexed tokenID)
func (_Nft *NFT) ParseBurn(log Log) (*NftBurn, error) {
	event := new(NftBurn)
	if err := _Nft.contract.UnpackLog(event, "Burn", *log.log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftCancelAskIterator is returned from FilterCancelAsk and is used to iterate over the raw logs and unpacked data for CancelAsk events raised by the Nft contract.
type NftCancelAskIterator struct {
	Event *NftCancelAsk // Event containing the contract specifics and raw log

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
func (it *NftCancelAskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftCancelAsk)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftCancelAsk)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftCancelAskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftCancelAskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftCancelAsk represents a CancelAsk event raised by the Nft contract.
type NftCancelAsk struct {
	AskID     *BigInt
	AuctionID *BigInt
	Asker     Address
	TokenID   *BigInt
	Raw       Log // Blockchain specific contextual infos
}

// FilterCancelAsk is a free log retrieval operation binding the contract event 0x1707cbe59d8b6e8d9fa88391cef815dd7ccc42755e02fb7d4b590befbfb6af12.
//
// Solidity: event CancelAsk(uint256 askID, uint256 indexed auctionID, address indexed asker, uint256 indexed tokenID)
func (_Nft *NFT) FilterCancelAsk(
	opts *FilterOpts,
	auctionID []*BigInt,
	asker []Address,
	tokenID []*BigInt,
) (*NftCancelAskIterator, error) {

	var auctionIDRule []interface{}
	for _, auctionIDItem := range auctionID {
		auctionIDRule = append(auctionIDRule, auctionIDItem.bigint)
	}
	var askerRule []interface{}
	for _, askerItem := range asker {
		askerRule = append(askerRule, askerItem.address)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem.bigint)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "CancelAsk", auctionIDRule, askerRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &NftCancelAskIterator{contract: _Nft.contract, event: "CancelAsk", logs: logs, sub: sub}, nil
}

// WatchCancelAsk is a free log subscription operation binding the contract event 0x1707cbe59d8b6e8d9fa88391cef815dd7ccc42755e02fb7d4b590befbfb6af12.
//
// Solidity: event CancelAsk(uint256 askID, uint256 indexed auctionID, address indexed asker, uint256 indexed tokenID)
func (_Nft *NFT) WatchCancelAsk(
	opts *WatchOpts,
	sink chan<- *NftCancelAsk,
	auctionID []*BigInt,
	asker []Address,
	tokenID []*BigInt,
) (*Subscription, error) {

	var auctionIDRule []interface{}
	for _, auctionIDItem := range auctionID {
		auctionIDRule = append(auctionIDRule, auctionIDItem.bigint)
	}
	var askerRule []interface{}
	for _, askerItem := range asker {
		askerRule = append(askerRule, askerItem.address)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem.bigint)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts.opts, "CancelAsk", auctionIDRule, askerRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftCancelAsk)
				if err := _Nft.contract.UnpackLog(event, "CancelAsk", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})
	return &Subscription{sub: subscription}, nil
}

// ParseCancelAsk is a log parse operation binding the contract event 0x1707cbe59d8b6e8d9fa88391cef815dd7ccc42755e02fb7d4b590befbfb6af12.
//
// Solidity: event CancelAsk(uint256 askID, uint256 indexed auctionID, address indexed asker, uint256 indexed tokenID)
func (_Nft *NFT) ParseCancelAsk(log Log) (*NftCancelAsk, error) {
	event := new(NftCancelAsk)
	if err := _Nft.contract.UnpackLog(event, "CancelAsk", *log.log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftCancelBidIterator is returned from FilterCancelBid and is used to iterate over the raw logs and unpacked data for CancelBid events raised by the Nft contract.
type NftCancelBidIterator struct {
	Event *NftCancelBid // Event containing the contract specifics and raw log

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
func (it *NftCancelBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftCancelBid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftCancelBid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftCancelBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftCancelBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftCancelBid represents a CancelBid event raised by the Nft contract.
type NftCancelBid struct {
	BidID     *BigInt
	AuctionID *BigInt
	Bidder    Address
	Amount    *BigInt
	Raw       Log // Blockchain specific contextual infos
}

// FilterCancelBid is a free log retrieval operation binding the contract event 0x44aad3fc20c7afe11491ebc440f443a1831651c1daec48b6ace99acdfba88c44.
//
// Solidity: event CancelBid(uint256 bidID, uint256 indexed auctionID, address indexed bidder, uint256 amount)
func (_Nft *NFT) FilterCancelBid(
	opts *FilterOpts,
	auctionID []*BigInt,
	bidder []Address,
) (*NftCancelBidIterator, error) {

	var auctionIDRule []interface{}
	for _, auctionIDItem := range auctionID {
		auctionIDRule = append(auctionIDRule, auctionIDItem.bigint)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem.address)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "CancelBid", auctionIDRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &NftCancelBidIterator{contract: _Nft.contract, event: "CancelBid", logs: logs, sub: sub}, nil
}

// WatchCancelBid is a free log subscription operation binding the contract event 0x44aad3fc20c7afe11491ebc440f443a1831651c1daec48b6ace99acdfba88c44.
//
// Solidity: event CancelBid(uint256 bidID, uint256 indexed auctionID, address indexed bidder, uint256 amount)
func (_Nft *NFT) WatchCancelBid(
	opts *WatchOpts,
	sink chan<- *NftCancelBid,
	auctionID []*BigInt,
	bidder []Address,
) (*Subscription, error) {

	var auctionIDRule []interface{}
	for _, auctionIDItem := range auctionID {
		auctionIDRule = append(auctionIDRule, auctionIDItem.bigint)
	}
	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem.address)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts.opts, "CancelBid", auctionIDRule, bidderRule)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftCancelBid)
				if err := _Nft.contract.UnpackLog(event, "CancelBid", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})
	return &Subscription{sub: subscription}, nil
}

// ParseCancelBid is a log parse operation binding the contract event 0x44aad3fc20c7afe11491ebc440f443a1831651c1daec48b6ace99acdfba88c44.
//
// Solidity: event CancelBid(uint256 bidID, uint256 indexed auctionID, address indexed bidder, uint256 amount)
func (_Nft *NFT) ParseCancelBid(log Log) (*NftCancelBid, error) {
	event := new(NftCancelBid)
	if err := _Nft.contract.UnpackLog(event, "CancelBid", *log.log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftCloseAuctionIterator is returned from FilterCloseAuction and is used to iterate over the raw logs and unpacked data for CloseAuction events raised by the Nft contract.
type NftCloseAuctionIterator struct {
	Event *NftCloseAuction // Event containing the contract specifics and raw log

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
func (it *NftCloseAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftCloseAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftCloseAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftCloseAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftCloseAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftCloseAuction represents a CloseAuction event raised by the Nft contract.
type NftCloseAuction struct {
	TokenTypeID *BigInt
	AuctionID   *BigInt
	Raw         Log // Blockchain specific contextual infos
}

// FilterCloseAuction is a free log retrieval operation binding the contract event 0xd4c37cf7a2dd87bf216584303aca56ea151859875ce77d12a2d35f240b39d748.
//
// Solidity: event CloseAuction(uint256 indexed tokenTypeID, uint256 auctionID)
func (_Nft *NFT) FilterCloseAuction(opts *FilterOpts, tokenTypeID []*BigInt) (*NftCloseAuctionIterator, error) {

	var tokenTypeIDRule []interface{}
	for _, tokenTypeIDItem := range tokenTypeID {
		tokenTypeIDRule = append(tokenTypeIDRule, tokenTypeIDItem.bigint)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "CloseAuction", tokenTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &NftCloseAuctionIterator{contract: _Nft.contract, event: "CloseAuction", logs: logs, sub: sub}, nil
}

// WatchCloseAuction is a free log subscription operation binding the contract event 0xd4c37cf7a2dd87bf216584303aca56ea151859875ce77d12a2d35f240b39d748.
//
// Solidity: event CloseAuction(uint256 indexed tokenTypeID, uint256 auctionID)
func (_Nft *NFT) WatchCloseAuction(
	opts *WatchOpts,
	sink chan<- *NftCloseAuction,
	tokenTypeID []*BigInt,
) (*Subscription, error) {

	var tokenTypeIDRule []interface{}
	for _, tokenTypeIDItem := range tokenTypeID {
		tokenTypeIDRule = append(tokenTypeIDRule, tokenTypeIDItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts.opts, "CloseAuction", tokenTypeIDRule)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftCloseAuction)
				if err := _Nft.contract.UnpackLog(event, "CloseAuction", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})
	return &Subscription{sub: subscription}, nil
}

// ParseCloseAuction is a log parse operation binding the contract event 0xd4c37cf7a2dd87bf216584303aca56ea151859875ce77d12a2d35f240b39d748.
//
// Solidity: event CloseAuction(uint256 indexed tokenTypeID, uint256 auctionID)
func (_Nft *NFT) ParseCloseAuction(log Log) (*NftCloseAuction, error) {
	event := new(NftCloseAuction)
	if err := _Nft.contract.UnpackLog(event, "CloseAuction", *log.log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftIssuerTakeRatePayoutIterator is returned from FilterIssuerTakeRatePayout and is used to iterate over the raw logs and unpacked data for IssuerTakeRatePayout events raised by the Nft contract.
type NftIssuerTakeRatePayoutIterator struct {
	Event *NftIssuerTakeRatePayout // Event containing the contract specifics and raw log

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
func (it *NftIssuerTakeRatePayoutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftIssuerTakeRatePayout)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftIssuerTakeRatePayout)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftIssuerTakeRatePayoutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftIssuerTakeRatePayoutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftIssuerTakeRatePayout represents a IssuerTakeRatePayout event raised by the Nft contract.
type NftIssuerTakeRatePayout struct {
	Issuer      Address
	TokenTypeID *BigInt
	AuctionID   *BigInt
	Raw         Log // Blockchain specific contextual infos
}

// FilterIssuerTakeRatePayout is a free log retrieval operation binding the contract event 0xbe6dc3ff42a1ff7df77852c377f9c53dc38edbc198ed92a4986516d8bf6f8e0a.
//
// Solidity: event IssuerTakeRatePayout(address indexed issuer, uint256 indexed tokenTypeID, uint256 indexed auctionID)
func (_Nft *NFT) FilterIssuerTakeRatePayout(
	opts *FilterOpts,
	issuer []Address,
	tokenTypeID []*BigInt,
	auctionID []*BigInt,
) (*NftIssuerTakeRatePayoutIterator, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem.address)
	}
	var tokenTypeIDRule []interface{}
	for _, tokenTypeIDItem := range tokenTypeID {
		tokenTypeIDRule = append(tokenTypeIDRule, tokenTypeIDItem.bigint)
	}
	var auctionIDRule []interface{}
	for _, auctionIDItem := range auctionID {
		auctionIDRule = append(auctionIDRule, auctionIDItem.bigint)
	}

	logs, sub, err := _Nft.contract.FilterLogs(
		opts.opts,
		"IssuerTakeRatePayout",
		issuerRule,
		tokenTypeIDRule,
		auctionIDRule,
	)
	if err != nil {
		return nil, err
	}
	return &NftIssuerTakeRatePayoutIterator{
		contract: _Nft.contract,
		event:    "IssuerTakeRatePayout",
		logs:     logs,
		sub:      sub,
	}, nil
}

// WatchIssuerTakeRatePayout is a free log subscription operation binding the contract event 0xbe6dc3ff42a1ff7df77852c377f9c53dc38edbc198ed92a4986516d8bf6f8e0a.
//
// Solidity: event IssuerTakeRatePayout(address indexed issuer, uint256 indexed tokenTypeID, uint256 indexed auctionID)
func (_Nft *NFT) WatchIssuerTakeRatePayout(
	opts *WatchOpts,
	sink chan<- *NftIssuerTakeRatePayout,
	issuer []Address,
	tokenTypeID []*BigInt,
	auctionID []*BigInt,
) (*Subscription, error) {

	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem.address)
	}
	var tokenTypeIDRule []interface{}
	for _, tokenTypeIDItem := range tokenTypeID {
		tokenTypeIDRule = append(tokenTypeIDRule, tokenTypeIDItem.bigint)
	}
	var auctionIDRule []interface{}
	for _, auctionIDItem := range auctionID {
		auctionIDRule = append(auctionIDRule, auctionIDItem.bigint)
	}

	logs, sub, err := _Nft.contract.WatchLogs(
		opts.opts,
		"IssuerTakeRatePayout",
		issuerRule,
		tokenTypeIDRule,
		auctionIDRule,
	)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftIssuerTakeRatePayout)
				if err := _Nft.contract.UnpackLog(event, "IssuerTakeRatePayout", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})
	return &Subscription{sub: subscription}, nil
}

// ParseIssuerTakeRatePayout is a log parse operation binding the contract event 0xbe6dc3ff42a1ff7df77852c377f9c53dc38edbc198ed92a4986516d8bf6f8e0a.
//
// Solidity: event IssuerTakeRatePayout(address indexed issuer, uint256 indexed tokenTypeID, uint256 indexed auctionID)
func (_Nft *NFT) ParseIssuerTakeRatePayout(log Log) (*NftIssuerTakeRatePayout, error) {
	event := new(NftIssuerTakeRatePayout)
	if err := _Nft.contract.UnpackLog(event, "IssuerTakeRatePayout", *log.log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Nft contract.
type NftMintIterator struct {
	Event *NftMint // Event containing the contract specifics and raw log

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
func (it *NftMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftMint represents a Mint event raised by the Nft contract.
type NftMint struct {
	To          Address
	TokenTypeID *BigInt
	TokenID     *BigInt
	Uri         string
	MintCount   *BigInt
	Raw         Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x397c4cd4ff46491382eefda34299d046406c58377a808e6eb8a0f0e892ec8bd1.
//
// Solidity: event Mint(address indexed to, uint256 indexed tokenTypeID, uint256 indexed tokenID, string uri, uint256 mintCount)
func (_Nft *NFT) FilterMint(
	opts *FilterOpts,
	to []Address,
	tokenTypeID []*BigInt,
	tokenID []*BigInt,
) (*NftMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem.address)
	}
	var tokenTypeIDRule []interface{}
	for _, tokenTypeIDItem := range tokenTypeID {
		tokenTypeIDRule = append(tokenTypeIDRule, tokenTypeIDItem.bigint)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem.bigint)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "Mint", toRule, tokenTypeIDRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	return &NftMintIterator{contract: _Nft.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x397c4cd4ff46491382eefda34299d046406c58377a808e6eb8a0f0e892ec8bd1.
//
// Solidity: event Mint(address indexed to, uint256 indexed tokenTypeID, uint256 indexed tokenID, string uri, uint256 mintCount)
func (_Nft *NFT) WatchMint(
	opts *WatchOpts,
	sink chan<- *NftMint,
	to []Address,
	tokenTypeID []*BigInt,
	tokenID []*BigInt,
) (*Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem.address)
	}
	var tokenTypeIDRule []interface{}
	for _, tokenTypeIDItem := range tokenTypeID {
		tokenTypeIDRule = append(tokenTypeIDRule, tokenTypeIDItem.bigint)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem.bigint)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts.opts, "Mint", toRule, tokenTypeIDRule, tokenIDRule)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftMint)
				if err := _Nft.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})
	return &Subscription{sub: subscription}, nil
}

// ParseMint is a log parse operation binding the contract event 0x397c4cd4ff46491382eefda34299d046406c58377a808e6eb8a0f0e892ec8bd1.
//
// Solidity: event Mint(address indexed to, uint256 indexed tokenTypeID, uint256 indexed tokenID, string uri, uint256 mintCount)
func (_Nft *NFT) ParseMint(log Log) (*NftMint, error) {
	event := new(NftMint)
	if err := _Nft.contract.UnpackLog(event, "Mint", *log.log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftNewTokenTypeIterator is returned from FilterNewTokenType and is used to iterate over the raw logs and unpacked data for NewTokenType events raised by the Nft contract.
type NftNewTokenTypeIterator struct {
	Event *NftNewTokenType // Event containing the contract specifics and raw log

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
func (it *NftNewTokenTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftNewTokenType)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftNewTokenType)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftNewTokenTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftNewTokenTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftNewTokenType represents a NewTokenType event raised by the Nft contract.
type NftNewTokenType struct {
	Issuer      Address
	Id          *BigInt
	OrgID       string
	MetadataURI string
	Raw         Log // Blockchain specific contextual infos
}

// FilterNewTokenType is a free log retrieval operation binding the contract event 0xc283822d2ec6cbcc46ff300f3abd709e5516aba2da7a04cd5010e131e4e2a4ab.
//
// Solidity: event NewTokenType(address indexed issuer, uint256 indexed id, string orgID, string metadataURI)
func (_Nft *NFT) FilterNewTokenType(
	opts *FilterOpts,
	issuer *Address,
	id *BigInt,
) (*NftNewTokenTypeIterator, error) {
	var issuerRule []interface{}
	for _, issuerItem := range []common.Address{issuer.address} {
		issuerRule = append(issuerRule, issuerItem)
	}
	var idRule []interface{}
	for _, idItem := range []*big.Int{id.bigint} {
		idRule = append(idRule, *idItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "NewTokenType", issuerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &NftNewTokenTypeIterator{contract: _Nft.contract, event: "NewTokenType", logs: logs, sub: sub}, nil
}

// WatchNewTokenType is a free log subscription operation binding the contract event 0xc283822d2ec6cbcc46ff300f3abd709e5516aba2da7a04cd5010e131e4e2a4ab.
//
// Solidity: event NewTokenType(address indexed issuer, uint256 indexed id, string orgID, string metadataURI)
func (_Nft *NFT) WatchNewTokenType(
	opts *WatchOpts,
	sink chan<- *NftNewTokenType,
	issuer *Address,
	id *BigInt,
) (*Subscription, error) {
	var issuerRule []interface{}
	for _, issuerItem := range []common.Address{issuer.address} {
		issuerRule = append(issuerRule, issuerItem)
	}
	var idRule []interface{}
	for _, idItem := range []*big.Int{id.bigint} {
		idRule = append(idRule, *idItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts.opts, "NewTokenType", issuerRule, idRule)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftNewTokenType)
				if err := _Nft.contract.UnpackLog(event, "NewTokenType", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})
	return &Subscription{sub: subscription}, nil
}

// ParseNewTokenType is a log parse operation binding the contract event 0xc283822d2ec6cbcc46ff300f3abd709e5516aba2da7a04cd5010e131e4e2a4ab.
//
// Solidity: event NewTokenType(address indexed issuer, uint256 indexed id, string orgID, string metadataURI)
func (_Nft *NFT) ParseNewTokenType(log Log) (*NftNewTokenType, error) {
	event := new(NftNewTokenType)
	if err := _Nft.contract.UnpackLog(event, "NewTokenType", *log.log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftOpenAuctionIterator is returned from FilterOpenAuction and is used to iterate over the raw logs and unpacked data for OpenAuction events raised by the Nft contract.
type NftOpenAuctionIterator struct {
	Event *NftOpenAuction // Event containing the contract specifics and raw log

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
func (it *NftOpenAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftOpenAuction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftOpenAuction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftOpenAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftOpenAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftOpenAuction represents a OpenAuction event raised by the Nft contract.
type NftOpenAuction struct {
	TokenTypeID *BigInt
	AuctionID   *BigInt
	Raw         Log // Blockchain specific contextual infos
}

// FilterOpenAuction is a free log retrieval operation binding the contract event 0xe5289abc23e6bab8c815c19f123d34b2ad142e1a05fb921a919f8f02ab88b052.
//
// Solidity: event OpenAuction(uint256 indexed tokenTypeID, uint256 indexed auctionID)
func (_Nft *NFT) FilterOpenAuction(
	opts *FilterOpts,
	tokenTypeID []*BigInt,
	auctionID []*BigInt,
) (*NftOpenAuctionIterator, error) {

	var tokenTypeIDRule []interface{}
	for _, tokenTypeIDItem := range tokenTypeID {
		tokenTypeIDRule = append(tokenTypeIDRule, tokenTypeIDItem.bigint)
	}
	var auctionIDRule []interface{}
	for _, auctionIDItem := range auctionID {
		auctionIDRule = append(auctionIDRule, auctionIDItem.bigint)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "OpenAuction", tokenTypeIDRule, auctionIDRule)
	if err != nil {
		return nil, err
	}
	return &NftOpenAuctionIterator{contract: _Nft.contract, event: "OpenAuction", logs: logs, sub: sub}, nil
}

// WatchOpenAuction is a free log subscription operation binding the contract event 0xe5289abc23e6bab8c815c19f123d34b2ad142e1a05fb921a919f8f02ab88b052.
//
// Solidity: event OpenAuction(uint256 indexed tokenTypeID, uint256 indexed auctionID)
func (_Nft *NFT) WatchOpenAuction(
	opts *WatchOpts,
	sink chan<- *NftOpenAuction,
	tokenTypeID []*BigInt,
	auctionID []*BigInt,
) (*Subscription, error) {

	var tokenTypeIDRule []interface{}
	for _, tokenTypeIDItem := range tokenTypeID {
		tokenTypeIDRule = append(tokenTypeIDRule, tokenTypeIDItem.bigint)
	}
	var auctionIDRule []interface{}
	for _, auctionIDItem := range auctionID {
		auctionIDRule = append(auctionIDRule, auctionIDItem.bigint)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts.opts, "OpenAuction", tokenTypeIDRule, auctionIDRule)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftOpenAuction)
				if err := _Nft.contract.UnpackLog(event, "OpenAuction", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})
	return &Subscription{sub: subscription}, nil
}

// ParseOpenAuction is a log parse operation binding the contract event 0xe5289abc23e6bab8c815c19f123d34b2ad142e1a05fb921a919f8f02ab88b052.
//
// Solidity: event OpenAuction(uint256 indexed tokenTypeID, uint256 indexed auctionID)
func (_Nft *NFT) ParseOpenAuction(log Log) (*NftOpenAuction, error) {
	event := new(NftOpenAuction)
	if err := _Nft.contract.UnpackLog(event, "OpenAuction", *log.log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftOrderExecutionIterator is returned from FilterOrderExecution and is used to iterate over the raw logs and unpacked data for OrderExecution events raised by the Nft contract.
type NftOrderExecutionIterator struct {
	Event *NftOrderExecution // Event containing the contract specifics and raw log

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
func (it *NftOrderExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftOrderExecution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftOrderExecution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftOrderExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftOrderExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftOrderExecution represents a OrderExecution event raised by the Nft contract.
type NftOrderExecution struct {
	AuctionID    *BigInt
	Price        *BigInt
	IssuerPayout *BigInt
	TokenID      *BigInt
	Issuer       Address
	AskID        *BigInt
	BidID        *BigInt
	Raw          Log // Blockchain specific contextual infos
}

// FilterOrderExecution is a free log retrieval operation binding the contract event 0xc43c5335fefedeee6d99290a4a9f99220ddc76f381bced0df3b950739a6e5988.
//
// Solidity: event OrderExecution(uint256 indexed auctionID, uint256 price, uint256 issuerPayout, uint256 indexed tokenID, address indexed issuer, uint256 askID, uint256 bidID)
func (_Nft *NFT) FilterOrderExecution(
	opts *FilterOpts,
	auctionID []*BigInt,
	tokenID []*BigInt,
	issuer []Address,
) (*NftOrderExecutionIterator, error) {

	var auctionIDRule []interface{}
	for _, auctionIDItem := range auctionID {
		auctionIDRule = append(auctionIDRule, auctionIDItem.bigint)
	}

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem.bigint)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem.address)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "OrderExecution", auctionIDRule, tokenIDRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &NftOrderExecutionIterator{contract: _Nft.contract, event: "OrderExecution", logs: logs, sub: sub}, nil
}

// WatchOrderExecution is a free log subscription operation binding the contract event 0xc43c5335fefedeee6d99290a4a9f99220ddc76f381bced0df3b950739a6e5988.
//
// Solidity: event OrderExecution(uint256 indexed auctionID, uint256 price, uint256 issuerPayout, uint256 indexed tokenID, address indexed issuer, uint256 askID, uint256 bidID)
func (_Nft *NFT) WatchOrderExecution(
	opts *WatchOpts,
	sink chan<- *NftOrderExecution,
	auctionID []*BigInt,
	tokenID []*BigInt,
	issuer []Address,
) (*Subscription, error) {

	var auctionIDRule []interface{}
	for _, auctionIDItem := range auctionID {
		auctionIDRule = append(auctionIDRule, auctionIDItem.bigint)
	}

	var tokenIDRule []interface{}
	for _, tokenIDItem := range tokenID {
		tokenIDRule = append(tokenIDRule, tokenIDItem.bigint)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem.address)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts.opts, "OrderExecution", auctionIDRule, tokenIDRule, issuerRule)
	if err != nil {
		return nil, err
	}
	subscription := event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftOrderExecution)
				if err := _Nft.contract.UnpackLog(event, "OrderExecution", log); err != nil {
					return err
				}
				event.Raw = Log{log: &log}

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
	})

	return &Subscription{sub: subscription}, nil
}

// ParseOrderExecution is a log parse operation binding the contract event 0xc43c5335fefedeee6d99290a4a9f99220ddc76f381bced0df3b950739a6e5988.
//
// Solidity: event OrderExecution(uint256 indexed auctionID, uint256 price, uint256 issuerPayout, uint256 indexed tokenID, address indexed issuer, uint256 askID, uint256 bidID)
func (_Nft *NFT) ParseOrderExecution(log Log) (*NftOrderExecution, error) {
	event := new(NftOrderExecution)
	if err := _Nft.contract.UnpackLog(event, "OrderExecution", *log.log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTransferTokenIterator is returned from FilterTransferToken and is used to iterate over the raw logs and unpacked data for TransferToken events raised by the Nft contract.
type NftTransferTokenIterator struct {
	Event *NftTransferToken // Event containing the contract specifics and raw log

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
func (it *NftTransferTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTransferToken)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = Log{log: &log}
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NftTransferToken)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = Log{log: &log}
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NftTransferTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTransferTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTransferToken represents a TransferToken event raised by the Nft contract.
type NftTransferToken struct {
	To          Address
	TokenID     *BigInt
	TokenTypeID *BigInt
	Raw         Log // Blockchain specific contextual infos
}

// FilterTransferToken is a free log retrieval operation binding the contract event 0x2c828e46ad986d6729a875018a0fcedfeab620472ea2dd717ed2e519524ab25b.
//
// Solidity: event TransferToken(address indexed to, uint256 indexed tokenID, uint256 indexed tokenTypeID)
func (_Nft *NFT) FilterTransferToken(
	opts *FilterOpts,
	to *Address,
	tokenID *BigInt,
	tokenTypeID *BigInt,
) (*NftTransferTokenIterator, error) {

	var toRule []interface{}
	for _, toItem := range []common.Address{to.address} {
		toRule = append(toRule, toItem)
	}
	var tokenIDRule []interface{}
	for _, tokenIDItem := range []*big.Int{tokenID.bigint} {
		tokenIDRule = append(tokenIDRule, *tokenIDItem)
	}
	var tokenTypeIDRule []interface{}
	for _, tokenTypeIDItem := range []*big.Int{tokenTypeID.bigint} {
		tokenTypeIDRule = append(tokenTypeIDRule, *tokenTypeIDItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts.opts, "TransferToken", toRule, tokenIDRule, tokenTypeIDRule)
	if err != nil {
		return nil, err
	}
	return &NftTransferTokenIterator{contract: _Nft.contract, event: "TransferToken", logs: logs, sub: sub}, nil
}

// ParseTransferToken is a log parse operation binding the contract event 0x2c828e46ad986d6729a875018a0fcedfeab620472ea2dd717ed2e519524ab25b.
//
// Solidity: event TransferToken(address indexed to, uint256 indexed tokenID, uint256 indexed tokenTypeID)
func (_Nft *NFT) ParseTransferToken(log *Log) (*NftTransferToken, error) {
	event := new(NftTransferToken)
	if err := _Nft.contract.UnpackLog(event, "TransferToken", *log.log); err != nil {
		return nil, err
	}
	event.Raw = *log
	return event, nil
}
