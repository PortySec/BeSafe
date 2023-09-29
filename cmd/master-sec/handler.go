package main

import (
	"fmt"

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

func HashmeHandler(args []string) {
	algorithm := args[2]
	plaintext := args[3]
	result, err := hashporty.HashPlainText(algorithm, plaintext)
	if err != nil {
		fmt.Print("error in hashgraphy: ", err)
	}
	fmt.Println(result)
}
