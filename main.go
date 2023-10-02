package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Check if at least one argument is provided

	f, err := os.Open("./bin/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", h.Sum(nil))

	// if len(os.Args) < 2 {
	// 	fmt.Println("No command provided")
	// 	return
	// }

	// // Define a string flag with a default value and a short description.
	// hashmeCmd := flag.NewFlagSet("hashme", flag.ExitOnError)
	// algorithmPtr := hashmeCmd.String("a", "", "Hashing algorithm (e.g., sha256)")
	// filePtr := hashmeCmd.String("f", "", "Filename to hash")

	// // Check the command and parse the appropriate flags
	// println(os.Args[0], os.Args[1])
	// switch os.Args[1] {
	// case "hashme":
	// 	hashmeCmd.Parse(os.Args[2:])
	// default:
	// 	fmt.Println("Unknown command:", os.Args[1])
	// 	return
	// }

	// if *algorithmPtr == "" {
	// 	fmt.Println("You must specify a hash algorithm")
	// 	return
	// }
	// if *filePtr == "" {
	// 	fmt.Println("You must specify a file to hash")
	// 	return
	// }

	// fmt.Println("Hashing algorithm:", *algorithmPtr)
	// fmt.Println("Filename to hash:", *filePtr)
}
