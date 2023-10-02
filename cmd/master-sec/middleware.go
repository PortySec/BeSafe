package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

const alphanumericRegex = "^[a-zA-Z0-9@!#$%^&*?]+$"

func hashMeUsage() {
	fmt.Println(`Usage: porty hashme [OPTIONS] [ARGUMENTS]

Hashing utility to compute hash values for files or plaintext.

Options:
  -a ALGORITHM   Specify the hashing algorithm. Supported algorithms: sha1, sha256, sha512. Default: sha256.
  -f FILE        Specify the filename to hash.
  -h, --help     Display this help message.

Examples:
  porty hashme -a sha512 -f /path/to/file.txt     Compute the SHA-512 hash of a file.
  porty hashme -a sha1 "Hello, World!"            Compute the SHA-1 hash of a plaintext string.
  porty hashme "Hello, World!"                    Compute the SHA-256 hash of a plaintext string (using default algorithm).

Note: If both -f and plaintext are provided, only the file will be hashed.`)
	os.Exit(0)
}

func ValidArgsMiddleware(tool string, args []string) error {
	switch tool {
	case "passguard":
		return validatePassguardArgs(args)
	case "scanner":
		return validateScannerArgs(args)
	case "hashme":
		return validateHashMeArgs(args)
	default:
		return errors.New("unknown tool")
	}
}

func validatePassguardArgs(args []string) error {
	if len(args) != 3 {
		return errors.New("usage: passguard <password>")
	}
	if !isAlphanumeric(args[2]) {
		return errors.New("invalid password: must contain only alphanumeric characters")
	}
	return nil
}

func validateScannerArgs(args []string) error {
	if len(args) != 4 {
		return errors.New("usage: scanner <target> <port>")
	}
	if !isAlphanumeric(args[2]) {
		return errors.New("invalid host: must contain only alphanumeric characters")
	}
	return nil
}

func validateHashMeArgs(args []string) error {
	if len(args) <= 2 {
		hashMeUsage()
		return nil // or return an error if you want to halt execution after displaying usage
	}
	if len(args) >= 2 && args[2] == "-h" || args[2] == "--help" {
		hashMeUsage()
		return nil
	}

	return nil
}

func isAlphanumeric(input string) bool {
	match, _ := regexp.MatchString(alphanumericRegex, input)
	return match
}
