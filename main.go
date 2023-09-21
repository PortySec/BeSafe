package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}
	// fmt.Println(os.Args)
	switch os.Args[1] {
	case "passguard":
		execute("./bin/passguard.exe", os.Args[2:])
	case "scanner":
		execute("./bin/scanner.exe", os.Args[2:])
	}
}

func execute(tool string, args []string) {
	cmd := exec.Command(tool, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing tool: ", err)
	}
}

func printUsage() {
	fmt.Println("**************************************")
	fmt.Println("*                                    *")
	fmt.Println("*      Welcome to PortyGuard         *")
	fmt.Println("*      Your Trusted Security App     *")
	fmt.Println("*                                    *")
	fmt.Println("**************************************")
	fmt.Println("Usage: Porty <command> [args]")
	fmt.Println("\nCommands:")
	fmt.Println("passguard <password> - Assess the strength of your passsword")
	fmt.Println("scanner <host> <port> - Check connectivity to ost on given port")
	println()
}
