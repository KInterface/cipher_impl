package model

import (
	"fmt"

	"golang.org/x/crypto/chacha20poly1305"
)

type (
	nonce            []byte
	message          []byte
	encryptedMessage []byte
	key              []byte
	encryptedData    struct {
		nonce   nonce
		message []byte
		key     key
	}
	EncryptedDataBuilder struct {
		actions []func(*encryptedData)
	}
)

func (b *EncryptedDataBuilder) Nonce(n nonce) *EncryptedDataBuilder {
	b.actions = append(b.actions, func(e *encryptedData) {
		e.nonce = n
	})
	return b
}

func (b *EncryptedDataBuilder) Message(m encryptedMessage) *EncryptedDataBuilder {
	b.actions = append(b.actions, func(e *encryptedData) {
		e.message = encryptedMessage(m)
	})
	return b
}

func (b *EncryptedDataBuilder) Key(k []byte) *EncryptedDataBuilder {
	b.actions = append(b.actions, func(e *encryptedData) {
		e.key = key(k)
	})
	return b
}

type Cipher interface {
	Encrypt(message string)
	Decrypt() (string, error)
}

func (e *encryptedData) Encrypt(m string) {
	aead, err := chacha20poly1305.NewX(e.key)
	if err != nil {
		panic("Could't launch the programme correctly.")
	}
	e.message = aead.Seal(e.nonce, e.nonce, message(m), nil)
}

func (e *encryptedData) Decrypt() (string, error) {
	aead, err := chacha20poly1305.NewX(e.key)
	if err != nil {
		panic("Could't launch the programme correctly.")
	}
	d, err := aead.Open(nil, e.nonce, e.message, nil)
	if err != nil {
		return "", fmt.Errorf("failed decrypting %v,%w", e, err)
	}
	e.message = d
	return string(e.message), nil
}

func (b *EncryptedDataBuilder) Build() Cipher {
	e := encryptedData{}
	for _, f := range b.actions {
		f(&e)
	}
	return &e
}
