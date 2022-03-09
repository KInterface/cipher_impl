package model

import (
	"crypto"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"

	"github.com/aead/ecdh"
)

type (
	publicKey []byte
)

var (
	p256 elliptic.Curve
	pv   crypto.PrivateKey
)

func init() {
	p256 = elliptic.P256()
	p, _, _, err := elliptic.GenerateKey(p256, rand.Reader)
	if err != nil {
		panic("ecdh: fail generate a key pair.")
	}
	pv = crypto.PrivateKey(p)
}

func KeyExchange() publicKey {
	key, ok := checkPrivateKey(pv)
	if !ok {
		panic("ecdh: unexpected type of private key")
	}
	N := p256.Params().N
	if new(big.Int).SetBytes(key).Cmp(N) >= 0 {
		panic("ecdh: private key cannot used with given curve")
	}
	x, y := p256.ScalarBaseMult(key)

	return elliptic.Marshal(p256, x, y)
}

func checkPrivateKey(typeToCheck interface{}) (key []byte, ok bool) {
	switch t := typeToCheck.(type) {
	case []byte:
		key = t
		ok = true
	case *[]byte:
		key = *t
		ok = true
	}
	return
}

func Compute(pk crypto.PublicKey) []byte {
	curve := ecdh.Generic(p256)
	return curve.ComputeSecret(pv, pk)
}

func UnmarshalSharedPubKey(k []byte) (x, y *big.Int) {
	return elliptic.Unmarshal(p256, k)
}
