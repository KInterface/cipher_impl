package model

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	// ErrInvalidHash in returned by ComparePasswordAndHash if the provided
	// hash isn't in the expected format.
	ErrInvalidHash = errors.New("argon2id: hash is not in the correct format")

	// ErrIncompatibleVariant is returned by ComparePasswordAndHash if the
	// provided hash was created using a unsupported variant of Argon2.
	// Currently only argon2id is supported by this package.
	ErrIncompatibleVariant = errors.New("argon2id: incompatible variant of argon2")

	// ErrIncompatibleVersion is returned by ComparePasswordAndHash if the
	// provided hash was created using a different version of Argon2.
	ErrIncompatibleVersion = errors.New("argon2id: incompatible version of argon2")
)

type Params struct {
	// The amount of memory used by the algorithm (in kibibytes).
	Memory uint32

	// The number of iterations over the memory.
	Iterations uint32

	// The number of threads (or lanes) used by the algorithm.
	// Recommended value is between 1 and runtime.NumCPU().
	Parallelism uint8

	// Length of the random salt. 16 bytes is recommended for password hashing.
	SaltLength uint32

	// Length of the generated key. 16 bytes or more is recommended.
	KeyLength uint32
}

func ComparePasswordAndHash(password []byte, hash string) (match bool, err error) {
	match, _, err = CheckHash(password, hash)
	return match, err
}

// CheckHash is like ComparePasswordAndHash, except it also returns the params that the hash was
// created with. This can be useful if you want to update your hash params over time (which you
// should).
func CheckHash(password []byte, hash string) (match bool, key []byte, err error) {
	params, salt, key, err := DecodeHash(hash)
	if err != nil {
		return false, nil, err
	}

	otherKey := argon2.IDKey(password, salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)
	keyLen := int32(len(key))
	otherKeyLen := int32(len(otherKey))
	if subtle.ConstantTimeEq(keyLen, otherKeyLen) == 0 {
		return false, []byte{}, nil
	}
	if subtle.ConstantTimeCompare(key, otherKey) == 1 {
		return true, key, nil
	}
	return false, []byte{}, nil
}

// DecodeHash expects a hash created from this package, and parses it to return the params used to
// create it, as well as the salt and key (password hash).
func DecodeHash(hash string) (params *Params, salt, key []byte, err error) {
	vals := strings.Split(hash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	if vals[1] != "argon2id" {
		return nil, nil, nil, ErrIncompatibleVariant
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	params = &Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	params.SaltLength = uint32(len(salt))

	key, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	params.KeyLength = uint32(len(key))

	return params, salt, key, nil
}
