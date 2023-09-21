package main

import (
	"fmt"
	"os"

	"github.com/PortySec/Besafe/pkg/scan"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: scanner <target> <port>")
		return
	}
	if scan.Scanner(os.Args[1], os.Args[2]) {
		fmt.Println("Connection has been Successfully established ")
		return
	} else {
		fmt.Printf("Error connecting %v on port %v,", os.Args[1], os.Args[2])
		return
	}

}
