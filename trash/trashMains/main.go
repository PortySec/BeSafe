// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		printUsage()
// 		return
// 	}
// 	// fmt.Println(os.Args)
// 	switch os.Args[1] {
// 	case "passguard":
// 		execute("./bin/passguard.exe", os.Args[2:])
// 		// hashme.HashMain(os.Args[2:])

// 	case "scanner":
// 		execute("./bin/scanner.exe", os.Args[2:])
// 	case "hashme":
// 		execute("./bin/hashme.exe", os.Args[2:])
// 	}
// }

// func execute(tool string, args []string) {
// 	cmd := exec.Command(tool, args...)
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	err := cmd.Run()
// 	if err != nil {
// 		fmt.Println("Error executing tool: ", err)
// 	}
// }

// func printUsage() {
// 	green := "\x1b[32m" // ANSI escape code for green
// 	cyan := "\x1b[36m"  // ANSI escape code for cyan
// 	reset := "\x1b[0m"  // Reset color to default

// 	fmt.Println(green + "**************************************" + reset)
// 	fmt.Println(green + "*                                    *" + reset)
// 	fmt.Println(green + "*      " + cyan + "Welcome to PortyGuard" + green + "         *" + reset)
// 	fmt.Println(green + "*      " + cyan + "Your Trusted Security App" + green + "     *" + reset)
// 	fmt.Println(green + "*                                    *" + reset)
// 	fmt.Println(green + "**************************************" + reset)
// 	fmt.Println("\nUsage: Porty <command> [args]")
// 	fmt.Println("\nCommands:")
// 	fmt.Println("=================== " + cyan + "PASSGUARD" + reset + " ============================")
// 	fmt.Println("passguard <password> - Assess the strength of your password")
// 	fmt.Println("e.g. Porty passguard 1234@abcd")
// 	fmt.Println("=================== " + cyan + "SCANNER" + reset + " ==============================")
// 	fmt.Println("scanner <host> <port> - Check connectivity to host on given port")
// 	fmt.Println("e.g. Porty scanner google.com 80")
// 	fmt.Println("=================== " + cyan + "HASHME" + reset + " ===============================")
// 	fmt.Println("hashme <sha1/sha256/sha512> <plaintext> - Will hash your text based on the selected algorithm")
// 	fmt.Println("e.g. Porty hashme sha256 helloworld")
// 	fmt.Println()
// }
