package main

import (
	"errors"
	"regexp"
)

func ValidArgsMiddleware(tool string, args []string) error {
	switch tool {
	case "passguard":
		if len(args) != 3 {
			return errors.New("usage: passguard <password>")
		}
		if !isAlphanumeric(args[2]) {
			return errors.New("invalid password: must contain only alphanumeric characters")
		}
		return nil
	case "scanner":
		if len(args) != 4 {
			return errors.New("usage: scanner <target> <port>")
		}
		if !isAlphanumeric(args[2]) {
			return errors.New("invalid host: must contain only alphanumeric characters")
		}
		return nil
	case "hashme":
		if len(args) != 4 {
			return errors.New("usage: hashme <sha1/sha256/sha512> <plaintext>")
		}
		if !isAlphanumeric(args[3]) {
			return errors.New("invalid host: must contain only alphanumeric characters")
		}
		return nil
	default:
		printUsage()
	}
	return nil

}

func isAlphanumeric(input string) bool {
	// Use a regular expression to check for alphanumeric characters
	// You can customize this regex based on your requirements
	regex := "^[a-zA-Z0-9@!#$%^&*?]+$"
	match, _ := regexp.MatchString(regex, input)
	return match
}
