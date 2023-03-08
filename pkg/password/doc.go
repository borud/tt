// Package password implements convenience types and functions for dealing with passwords.
// The default parameters are based on recommendations for interactive use and should be
// okay.  In its simplest form you just use the default parameters to hash passwords:
//
//	hash, err := HashWithParam(DefaultInteractiveParameters, "mySecretPassword")
//
// Then to verify the hash you can do
//
//	res, err := Verify(hash, "mySecretPassword")
package password
