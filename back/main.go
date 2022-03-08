package main

import (
	"crypto"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"KInterface/ciper_impl/server"

	"github.com/aead/ecdh"
)

func main() {
	server.Serve()
	p256 := elliptic.P256()
	p, ax, ay, err := elliptic.GenerateKey(p256, rand.Reader)
	if err != nil {
		panic("ecdh: fail generate a key pair.")
	}
	pv := crypto.PrivateKey(p)
	x, y := p256.ScalarMult(ax, ay, p)
	pb := crypto.PublicKey(ecdh.Point{X: x, Y: y})
	gc := ecdh.Generic(p256)
	s := gc.ComputeSecret(pv, pb)
	s2 := gc.ComputeSecret(pv, pb)
	fmt.Println(s, s2)
}
