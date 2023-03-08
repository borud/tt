package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// UserCredentials contains a JWT token and implements the methods needed
// to use this as a authorization token in gRPC.
type UserCredentials struct {
	Token             string `json:"token"`
	TransportSecurity bool   `json:"transportSecurity"`
}

const (
	// AuthFileName is the file where we store the authetication credentials. This will be
	// on your home directory.
	AuthFileName = ".ttauth"
)

// Errors
var (
	ErrEmptyToken = errors.New("token is empty")
)

// GetRequestMetadata implements interface used by gRPC
func (t UserCredentials) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.Token,
	}, nil
}

// RequireTransportSecurity implements interface used by gRPC
func (t UserCredentials) RequireTransportSecurity() bool {
	return t.TransportSecurity
}

// SaveToConfigFile saves the auth token to the user's home directory.  Convenience function.
func (t UserCredentials) SaveToConfigFile() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	p := path.Join(home, AuthFileName)

	fmt.Printf("saving credentials to %s\n", p)

	return os.WriteFile(p, jsonData, 0600)
}

// LoadCredentialsFromConfigFile loads an auth token from the default file in the user's home directory.
// Convenience function.
func LoadCredentialsFromConfigFile() (UserCredentials, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return UserCredentials{}, err
	}

	data, err := os.ReadFile(path.Join(home, AuthFileName))
	if err != nil {
		return UserCredentials{}, err
	}

	var token UserCredentials
	return token, json.Unmarshal(data, &token)
}

// Username extracts username from credentials.
func (t UserCredentials) Username() (string, string, error) {
	parts := strings.Split(t.Token, ".")
	if len(parts) != 3 {
		return "", "", fmt.Errorf("wrong number of parts in jwt token (%d)", len(parts))
	}

	//lint:ignore SA1019 golang-jwt is useless shit so we should stop using it anyway
	seg, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		return "", "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(seg, &result)
	if err != nil {
		return "", "", err
	}

	su := result["sub"]
	ro := result["role"]

	user, ok := su.(string)
	if !ok {
		return "", "", fmt.Errorf("unable to interpret [%v] as user", su)
	}

	role, ok := ro.(string)
	if !ok {
		return "", "", fmt.Errorf("unable to interpret [%v] as role", su)
	}

	return user, role, nil
}
