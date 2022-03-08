package server

import (
	"crypto"
	"encoding/json"
	"fmt"
	"net/http"

	"KInterface/ciper_impl/model"

	"github.com/aead/ecdh"
)

func Serve() {
	http.HandleFunc("/", handler)
	fmt.Println("Server start listening...")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		panic(err)
	}
}

type Body struct {
	Data      string
	Nonce     []byte
	Encrypted []byte
	SharedKey []byte
	Hash      string
}
type (
	token    []byte
	Response struct {
		Decoded string
		Encoded []byte
		Token   token
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method == "GET" {

		w.WriteHeader(http.StatusOK)
		pk := model.KeyExchange()
		fmt.Println(pk)
		w.Write(pk)
		return
	} else if r.Method == "POST" {
		var b Body
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		_, err := json.Marshal(b)
		if err != nil {
			panic(err)
		}
		e := model.EncryptedDataBuilder{}

		n, m, k, h := b.Nonce, b.Encrypted, b.SharedKey, b.Hash
		fmt.Println(h)
		x, y := model.A(k)
		pb := crypto.PublicKey(ecdh.Point{X: x, Y: y})
		s := model.Compute(pb)
		fmt.Println("here")

		_, ke, _ := model.CheckHash(s, h)
		fmt.Println(h, []byte(h))
		c := e.Nonce(n).Message(m).Key(ke).Build()
		txt, err := c.Decrypt()
		if err != nil {
			panic(err)
		}
		r, err := json.Marshal(Response{Decoded: txt, Encoded: m, Token: token(n)})
		if err != nil {
			panic(err)
		}
		fmt.Println(txt)
		w.Write(r)
		return
	}
}
