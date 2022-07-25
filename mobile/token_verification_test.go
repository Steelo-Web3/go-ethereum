package geth

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func TestTokenVerification(t *testing.T) {
	ks := keystore.NewPlaintextKeyStore(".")
	acc, err := ks.NewAccount("test")

	err = ks.Unlock(acc, "test")
	assert.Nil(t, err, "account creation err")

	msg := "0x12345"
	data := []byte(msg)
	hashMsg := TextHash(data)

	sig, err := ks.SignHash(acc, hashMsg)
	assert.Nil(t, err, "sign hash err")

	recAddr, err := EcRecover(data, sig)
	assert.Nil(t, err, "ec recover err")

	assert.Equal(t, recAddr.address, acc.Address)

	notSignerAccount, err := ks.NewAccount("fail test")
	err = ks.Unlock(notSignerAccount, "fail test")
	assert.Nil(t, err, "account creation err")

	assert.NotEqual(t, recAddr.address, notSignerAccount.Address)

}
