package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func ToHash(b []byte) string {
	hash := sha256.Sum256(b)
	return hex.EncodeToString(hash[:])
}
