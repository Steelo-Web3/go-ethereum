package geth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"

	"golang.org/x/crypto/sha3"
)

// TextHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func TextHash(data []byte) []byte {
	hash, _ := textAndHash(data)
	return hash
}

// TextAndHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func textAndHash(data []byte) ([]byte, string) {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), string(data))
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(msg))
	return hasher.Sum(nil), msg
}

func MsgFmt(data []byte) string {
	return fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), string(data))
}

// Returns the address for the Account that was used to create the signature.
//
// Note, this function is compatible with eth_sign and personal_sign. As such it recovers
// the address of:
// hash = keccak256("\x19${byteVersion}Ethereum Signed Message:\n${message length}${message}")
// addr = ecrecover(hash, signature)
//
// Note, the signature must conform to the secp256k1 curve R, S and V values, where
// the V value must be be 27 or 28 for legacy reasons.
//
// https://github.com/ethereum/go-ethereum/wiki/Management-APIs#personal_ecRecover
func EcRecover(data []byte, sig []byte) (addr *Address, _ error) {
	if len(sig) != 65 {
		return &Address{}, fmt.Errorf("signature must be 65 bytes long")
	}
	if sig[64] != 27 && sig[64] != 28 {
		return &Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27 // Transform yellow paper V from 27/28 to 0/1
	hash := TextHash(data)
	rpk, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return &Address{}, err
	}
	commonAddress := crypto.PubkeyToAddress(*rpk)
	return &Address{address: commonAddress}, nil
}
