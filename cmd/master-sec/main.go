package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	arguments := os.Args
	err := ValidArgsMiddleware(arguments[1], arguments)
	if err != nil {
		fmt.Println("Invalid arguments", err)
		os.Exit(0)
	}
	switch os.Args[1] {
	case "passguard":
		PassGuardHandler(arguments)
	case "scanner":
		ScannerHandler(arguments)
	case "hashme":
		HashmeHandler(arguments)
	default:
		fmt.Println("Unknown command")
		os.Exit(0)
	}
}

func printUsage() {
	fmt.Println("**************************************")
	fmt.Println("*                                    *")
	fmt.Println("*      ╔══════════════════╗          *")
	fmt.Println("*      ║   Welcome to     ║          *")
	fmt.Println("*      ║   PortyGuard     ║          *")
	fmt.Println("*      ║   Your Trusted   ║          *")
	fmt.Println("*      ║   Security App   ║          *")
	fmt.Println("*      ║                  ║          *")
	fmt.Println("*      ╚══════════════════╝          *")
	fmt.Println("*                                    *")
	fmt.Println("**************************************")
	fmt.Println("\nUsage: porty <command> [args]")
	fmt.Println("Commands:")
	fmt.Println("=================== PASSGUARD ============================")
	fmt.Println("\npassguard <password> - Assess the strength of your password")
	fmt.Println("e.g. porty passguard 1234@abcd")
	fmt.Println("\n=================== SCANNER ==============================")
	fmt.Println("\nscanner <host> <port> - Check connectivity to host on given port")
	fmt.Println("e.g. porty scanner google.com 80")
	fmt.Println("\n=================== HASHME ===============================")
	fmt.Println("\nhashme <sha1/sha256/sha512> <plaintext> - Will hash your text based on the selected algorithm")
	fmt.Println("e.g. porty hashme sha256 helloworld")
	fmt.Println()
}
