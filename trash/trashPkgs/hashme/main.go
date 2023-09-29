// package main

// import (
// 	"fmt"
// 	"os"

// 	hashporty "github.com/PortySec/Besafe/pkg/hash"
// )

// func main() {
// 	// fmt.Println(args)
// 	// os.Exit(0)
// 	if len(os.Args) != 3 {
// 		fmt.Println("Usage: hashme <sha1/sha256/sha512> <plaintext>")
// 		return
// 	}
// 	algorithm := os.Args[1]
// 	plaintext := os.Args[2]

// 	result, err := hashporty.HashPlain(algorithm, plaintext)

// 	if err != nil {
// 		fmt.Println("error in hashgraphy: ", err)
// 	}
// 	fmt.Println(result)
// }
