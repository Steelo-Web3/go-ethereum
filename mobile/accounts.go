// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains all the wrappers from the accounts package to support client side key
// management on mobile platforms.

package geth

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const (
	// StandardScryptN is the N parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	StandardScryptN = int(keystore.StandardScryptN)

	// StandardScryptP is the P parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	StandardScryptP = int(keystore.StandardScryptP)

	// LightScryptN is the N parameter of Scrypt encryption algorithm, using 4MB
	// memory and taking approximately 100ms CPU time on a modern processor.
	LightScryptN = int(keystore.LightScryptN)

	// LightScryptP is the P parameter of Scrypt encryption algorithm, using 4MB
	// memory and taking approximately 100ms CPU time on a modern processor.
	LightScryptP = int(keystore.LightScryptP)

	AccountStatus_Unlocked = int(0)
	AccountStatus_Locked   = int(1)
	AccountStatus_NotFound = int(2)
)

// Account represents a stored key.
type Account struct{ account accounts.Account }

func NewAccount(address Address, url URL) Account {
	return Account{
		accounts.Account{
			Address: address.address,
			URL:     url.url,
		},
	}
}

// Accounts represents a slice of accounts.
type Accounts struct{ accounts []accounts.Account }

// Size returns the number of accounts in the slice.
func (a *Accounts) Size() int {
	return len(a.accounts)
}

// Get returns the account at the given index from the slice.
func (a *Accounts) Get(index int) (account *Account, _ error) {
	if index < 0 || index >= len(a.accounts) {
		return nil, errors.New("index out of bounds")
	}
	return &Account{a.accounts[index]}, nil
}

// Set sets the account at the given index in the slice.
func (a *Accounts) Set(index int, account *Account) error {
	if index < 0 || index >= len(a.accounts) {
		return errors.New("index out of bounds")
	}
	a.accounts[index] = account.account
	return nil
}

// GetAddress retrieves the address associated with the account.
func (a *Account) GetAddress() *Address {
	return &Address{a.account.Address}
}

// GetURL retrieves the canonical URL of the account.
func (a *Account) GetURL() string {
	return a.account.URL.String()
}
