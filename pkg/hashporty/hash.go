package hashporty

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"hash"
	"io"
	"log"
	"os"
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

func HashChecksum(algorithm string, fileContent *os.File) (hash.Hash, error) {
	switch algorithm {
	case "sha256":
		h := sha256.New()
		if _, err := io.Copy(h, fileContent); err != nil {
			log.Fatal(err)
		}
		return h, nil
	case "sha512":
		h := sha512.New()
		if _, err := io.Copy(h, fileContent); err != nil {
			log.Fatal(err)
		}
		return h, nil
	case "sha1":
		h := sha1.New()
		if _, err := io.Copy(h, fileContent); err != nil {
			log.Fatal(err)
		}
		return h, nil
	default:
		return nil, errors.New("unknown algorithm, algorithm must be one of SHA1, SHA256, SHA512")
	}

}
