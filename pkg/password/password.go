package password

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"golang.org/x/crypto/argon2"
)

// Params represents the parameters used for password hashing.
type Params struct {
	Time    uint32
	Memory  uint32
	Threads uint8
	KeyLen  uint32
}

// ParsedParams represents the parameters extracted from the serialized hash.
type ParsedParams struct {
	Version int
	Time    uint32
	Memory  uint32
	Threads uint8
	Salt    []byte
	Hash    []byte
}

const (
	// saltLenInBytes is the number of bytes we use for the salt. 128 bits is the
	// recommended length, 64 bits is said to be tenable by the authors of argon2.
	// So we choose 128 bits to be on the safe side.  We make this a non-configurable
	// option to reduce the amount of rope provided to the developer.
	saltLenInBytes = 16
)

var (
	// DefaultInteractiveParameters are the recommended parameters for interactive use.
	// We might want to tweak those a bit.
	DefaultInteractiveParameters = Params{
		Time:    uint32(1),
		Memory:  uint32(64 * 1024),
		Threads: uint8(4),
		KeyLen:  uint32(32),
	}

	// argon2HashRegexp is a regexp that matches the argon2 serial format
	argon2HashRegexp = regexp.MustCompile(`[$]argon2(?:id|i)[$]v=(\d{1,3})[$]m=(\d{3,20}),t=(\d{1,4}),p=(\d{1,2})\$([A-Za-z0-9+/=]+)\$([A-Za-z0-9+/=]*)$`)
)

// HashWithParam hashes a password with a given parameter set.  The returned
// byte slice is a serialized hash including the parameters.
//
// Returns an error if rand.Read fails - which it shouldn't.
func HashWithParam(p Params, password string) ([]byte, error) {
	salt := make([]byte, saltLenInBytes)
	_, err := rand.Read(salt)
	if err != nil {
		return []byte{}, err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Time, p.Memory, p.Threads, p.KeyLen)

	return []byte(fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		p.Memory,
		p.Time,
		p.Threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))), nil
}

// Hash a password with the default parameters.  The returned byte slice
// is a serialized hash including the parameters.
//
// You should prefer to use this unless you have given the parameters some
// thought since we will try to maintain the DefaultInteractiveParameters
// and ensure they have appropriate values.
func Hash(password string) ([]byte, error) {
	return HashWithParam(DefaultInteractiveParameters, password)
}

// Verify that password matches (serialized) hash.
func Verify(password string, hash []byte) (bool, error) {
	dp, err := DecodeParamsFromHash(hash)
	if err != nil {
		return false, err
	}

	computedHash := argon2.IDKey([]byte(password), dp.Salt, dp.Time, dp.Memory, dp.Threads, uint32(len(dp.Hash)))

	return bytes.Equal(computedHash, dp.Hash), nil
}

// DecodeParamsFromHash decodes the passwod hashing parameters from the
// serialized password hash.
func DecodeParamsFromHash(hash []byte) (*ParsedParams, error) {
	m := argon2HashRegexp.FindSubmatch(hash)

	if len(m) != 7 {
		return nil, errors.New("parse error: wrong number of elements")
	}

	version, err := strconv.ParseInt(string(m[1]), 10, 0)
	if err != nil {
		return nil, fmt.Errorf("error parsing version: %w", err)
	}

	memory, err := strconv.ParseUint(string(m[2]), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error parsing memory: %w", err)
	}

	time, err := strconv.ParseUint(string(m[3]), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("error parsing time: %w", err)
	}

	threads, err := strconv.ParseUint(string(m[4]), 10, 8)
	if err != nil {
		return nil, fmt.Errorf("error parsing threads: %w", err)
	}

	decodedSalt, err := base64.RawStdEncoding.DecodeString(string(m[5]))
	if err != nil {
		return nil, fmt.Errorf("error decoding salt: %w", err)
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(string(m[6]))
	if err != nil {
		return nil, fmt.Errorf("error decoding hash: %w", err)
	}

	return &ParsedParams{
		Time:    uint32(time),
		Memory:  uint32(memory),
		Threads: uint8(threads),
		Version: int(version),
		Salt:    decodedSalt,
		Hash:    decodedHash,
	}, nil
}
