package logging

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashEvent(text string) string {
	hash := sha256.Sum256([]byte(text))
	return hex.EncodeToString(hash[:])
}
