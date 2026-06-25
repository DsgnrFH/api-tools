package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashString returns the SHA-256 hash of a string input
func HashString(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
