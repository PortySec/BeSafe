// package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/PortySec/Besafe/pkg/passgate"
// )

// func main() {
// 	// fmt.Println(os.Args)
// 	if len(os.Args) != 2 {
// 		fmt.Println("Usage: passguard <password>")
// 		return
// 	}
// 	inputPass := os.Args[1]

// 	if passgate.IsValid(inputPass) {
// 		fmt.Println("Password Is Strong Enough")
// 	} else {
// 		fmt.Println("Password must Contain (A-Z,a-z,0-9,chars) !")
// 	}
// }
