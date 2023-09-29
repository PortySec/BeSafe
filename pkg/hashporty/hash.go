package hashporty

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

func HashPlainText(hashType string, hashInput string) (string, error) {
	var hashed string
	switch hashType {
	case "sha256":
		hashedHex := sha256.Sum256([]byte(hashInput))
		hashed = hex.EncodeToString(hashedHex[:])
		return hashed, nil
	case "sha512":
		hashedHex := sha512.Sum512([]byte(hashInput))
		hashed = hex.EncodeToString(hashedHex[:])
		return hashed, nil
	case "sha1":
		hashedHex := sha1.Sum([]byte(hashInput))
		hashed = hex.EncodeToString(hashedHex[:])
		return hashed, nil
	default:
		return "", errors.New("unknown algorithm, algorithm must be one of SHA1, SHA256, SHA512")
	}

}
