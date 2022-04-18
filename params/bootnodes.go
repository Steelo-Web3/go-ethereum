// Copyright 2015 The go-ethereum Authors
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

package params

import "github.com/ethereum/go-ethereum/common"

var Bootnodes = []string{
	"enode://b80e9256d64411f70191523fb10206b0624002117dd70f198ecac730ac2c7e83f3af53255ceefaae2a911c0be63bd6f1b72dd042530b7ad30c925235c3dc8857@192.168.1.217:30305",
	"enode://2c5b89bd02813e6b5dd95481c2dc468b89524ebd9f57ceba3cc534a1ccee4e09071484e422177237a498d6762baadb597c380b36ad275b3e7f0eb78af5fa6800@192.168.1.217:30303",
}

// TODO setup DNS network
const dnsPrefix = ""
const net = ""

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
