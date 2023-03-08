package password

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/argon2"
)

func TestHash(t *testing.T) {
	t.Parallel()

	hash, err := HashWithParam(DefaultInteractiveParameters, "secret")
	assert.Nil(t, err)
	assert.NotEmpty(t, hash)

	res, err := Verify("secret", hash)
	assert.Nil(t, err)
	assert.True(t, res)
}

func TestInteractive(t *testing.T) {
	hash, err := Hash("secret")
	assert.NoError(t, err)
	fmt.Printf("hash: [%s]", hash)
}

func TestHashWithParams(t *testing.T) {
	t.Parallel()

	// Hash the password
	hash, err := HashWithParam(DefaultInteractiveParameters, "secret")
	assert.Nil(t, err)
	assert.NotEmpty(t, hash)

	// Make sure we can derive the parameters from the hash
	params, err := DecodeParamsFromHash(hash)
	assert.Nil(t, err)
	assert.Equal(t, params.Memory, DefaultInteractiveParameters.Memory)
	assert.Equal(t, params.Time, DefaultInteractiveParameters.Time)
	assert.Equal(t, params.Threads, DefaultInteractiveParameters.Threads)
	assert.Equal(t, params.Version, argon2.Version)

	// Extract salt and hash using non-regex method
	parts := strings.Split(string(hash), "$")
	assert.Len(t, parts, 6)

	decodedSalt, err := base64.RawStdEncoding.DecodeString(parts[4])
	assert.Nil(t, err)

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	assert.Nil(t, err)

	// Make sure we get the salt and hash
	assert.Equal(t, decodedSalt, params.Salt)
	assert.Equal(t, decodedHash, params.Hash)

	// Verify password matches
	eq, err := Verify("secret", hash)
	assert.Nil(t, err)
	assert.True(t, eq)

	// Verify password match failure
	eq, err = Verify("wrong", hash)
	assert.Nil(t, err)
	assert.False(t, eq)

	// Nil hash
	eq, err = Verify("secret", nil)
	assert.Error(t, err)
	assert.False(t, eq)

	// Empty hash
	eq, err = Verify("secret", []byte{})
	assert.Error(t, err)
	assert.False(t, eq)
}

func BenchmarkBaseline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HashWithParam(DefaultInteractiveParameters, "some password that is of a bit of length")
	}
}

func BenchmarkHashTime2(b *testing.B) {
	paramTime2 := DefaultInteractiveParameters
	paramTime2.Time = 2
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashWithParam(paramTime2, "some password that is of a bit of length")
	}
}

func BenchmarkHashTime3(b *testing.B) {
	paramTime3 := DefaultInteractiveParameters
	paramTime3.Time = 3
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashWithParam(paramTime3, "some password that is of a bit of length")
	}
}

func BenchmarkHashTime4(b *testing.B) {
	paramTime4 := DefaultInteractiveParameters
	paramTime4.Time = 4
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HashWithParam(paramTime4, "some password that is of a bit of length")
	}
}
