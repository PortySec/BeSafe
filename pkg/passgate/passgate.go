package passgate

import "strings"

func IsValid(password string) bool {
	validUppers := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	validLowers := "abcdefghijklmnopqrstuvwxyz"
	validDigits := "0123456789"
	validChars := "!@#$%^&*()_+/*-+."

	if len(password) < 12 {
		return false
	}
	containsUpper := false
	containsLower := false
	containsDigit := false
	containsCharacter := false
	for _, element := range password {
		switch {
		case strings.Contains(validUppers, string(element)):
			containsUpper = true
		case strings.Contains(validLowers, string(element)):
			containsLower = true
		case strings.Contains(validDigits, string(element)):
			containsDigit = true
		case strings.Contains(validChars, string(element)):
			containsCharacter = true
		}
	}
	return containsCharacter && containsDigit && containsLower && containsUpper

}
