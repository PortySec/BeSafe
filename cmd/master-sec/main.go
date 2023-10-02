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
	ValidArgsMiddleware(arguments[1], arguments)
	// if err != nil {
	// 	fmt.Println("Invalid arguments", err)
	// 	return
	// }
	switch os.Args[1] {
	case "passguard":
		PassGuardHandler(arguments)
	case "scanner":
		ScannerHandler(arguments)
	case "hashme":
		HashmeHandler(arguments)
	default:
		fmt.Println("Unknown command")
		return
	}
}

func printUsage() {
	fmt.Println("============================================")
	fmt.Println("               PortyGuard")
	fmt.Println("       Your Trusted Security App")
	fmt.Println("============================================")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  porty <command> [options] [args]")
	fmt.Println()
	fmt.Println("Commands:")

	fmt.Println("▶ PASSGUARD")
	fmt.Println("   - Assess the strength of a password.")
	fmt.Println("     $ porty passguard <password>")
	fmt.Println("     Example: porty passguard 1234@abcd")
	fmt.Println()

	fmt.Println("▶ SCANNER")
	fmt.Println("   - Check connectivity to a host on a specific port.")
	fmt.Println("     $ porty scanner <host> <port>")
	fmt.Println("     Example: porty scanner google.com 80")
	fmt.Println()

	fmt.Println("▶ HASHME")
	fmt.Println("   - Hash text or file content.")
	fmt.Println("     Options:")
	fmt.Println("       -a, --algorithm <algorithm>   Specify the hashing algorithm.")
	fmt.Println("       -f, --file <filename>         Hash the content of a file.")
	fmt.Println("       -s, --salt <salt>             Add a salt to the hashing process.")
	fmt.Println("       -v, --verify <hash>           Verify a hash against the input.")
	fmt.Println()
	fmt.Println("     Examples:")
	fmt.Println("       $ porty hashme -a sha256 helloworld")
	fmt.Println("       $ porty hashme -a sha256 -f myfile.txt")
	fmt.Println("       $ porty hashme -a sha256 -s mysalt helloworld")
	fmt.Println("       $ porty hashme -a sha256 -v <hash> helloworld")
	fmt.Println()

	fmt.Println("Supported Algorithms for HASHME:")
	fmt.Println("   • sha1")
	fmt.Println("   • sha256")
	fmt.Println("   • sha512")
	fmt.Println()

	fmt.Println("============================================")
	fmt.Println("For more details and advanced usage, visit: www.portyguard.com")
	fmt.Println("============================================")
}
