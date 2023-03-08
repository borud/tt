package util

import (
	"crypto/rand"
	"math/big"
)

const (
	// DefaultDataSourceKeyLenBits is the key length for generating datasource keys.
	DefaultDataSourceKeyLenBits = 512
)

// RandomString generates a string of random characters representing N bits of randomness.
func RandomString(numBits uint) (string, error) {
	n, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), numBits))
	if err != nil {
		return "", err
	}
	return n.Text(36), nil
}
