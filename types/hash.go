package types

import (
	"crypto/rand"
	"fmt"
)

type Hash [32]uint8

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		msg := fmt.Sprintf("Invalid hash length: %d", len(b))
		panic(msg)
	}

	var h Hash
	copy(h[:], b)
	return h
}

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}
