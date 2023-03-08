package auth

import (
	"time"
)

// Claims is a minimalistic Claims structure to try to insulate us from the underlying
// JWT implementation.
type Claims struct {
	JTI string
	ISS string
	SUB string
	EXP time.Time
	NBF time.Time
	IAT time.Time
}
