package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	// "github.com/PortySec/Besafe/pkg/hashporty"
	"github.com/PortySec/Besafe/pkg/hashporty"
	"github.com/PortySec/Besafe/pkg/passgateporty"
	"github.com/PortySec/Besafe/pkg/scanporty"
)

func PassGuardHandler(args []string) {
	inputPass := args[2]
	guardResult := passgateporty.IsValid(inputPass)
	if guardResult {
		fmt.Println("Password Is Strong Enough")
	} else {
		fmt.Println("Password must Contain [A-Z,a-z,0-9,chars] and have a length of at least 12")
	}
}

func ScannerHandler(args []string) {
	targetHost := args[2]
	targetPort := args[3]
	scanResult := scanporty.Scanner(targetHost, targetPort)
	if scanResult {
		fmt.Println("Connection has been Successfully established ")
		return
	} else {
		fmt.Printf("Error connecting %v on port %v,", targetHost, targetPort)
		return
	}
}

func HashmeHandler(args []string) error {
	hashmeCmd := flag.NewFlagSet("hashme", flag.ExitOnError)
	algorithmPtr := hashmeCmd.String("a", "sha256", "Hashing algorithm (e.g., sha256,sha512,sha1)")
	filePtr := hashmeCmd.String("f", "", "Filename to hash")

	hashmeCmd.Parse(args[2:])

	if *filePtr != "" {
		return hashFile(*filePtr, *algorithmPtr)
	}

	plaintext := hashmeCmd.Arg(0)
	if plaintext == "" {
		return errors.New("no plaintext provided")
	}
	return hashText(plaintext, *algorithmPtr)
}

func hashFile(filePath string, algorithm string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	result, err := hashporty.HashChecksum(algorithm, f)
	if err != nil {
		return fmt.Errorf("error in hashporty: %v", err)
	}

	hashedValue := fmt.Sprintf("%x", result.Sum(nil))
	fmt.Println(hashedValue)
	return nil
}

func hashText(plaintext string, algorithm string) error {
	result, err := hashporty.HashPlainText(algorithm, plaintext)
	if err != nil {
		return fmt.Errorf("error in hashporty: %v", err)
	}
	fmt.Println(result)
	return nil
}
