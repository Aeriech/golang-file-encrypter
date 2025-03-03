package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/aeriech/golang-file-encrypter/filecrypt"
)

var (
	reader = bufio.NewReader(os.Stdin)
)

func main() {
	if len(os.Args) < 2 {
		forMoreInfo()
		os.Exit(0)
	}

	functionName := os.Args[1]
	switch functionName {
	case "help":
		printHelp()
	case "encrypt":
		encryptFile()
	case "decrypt":
		decryptFile()
	default:
		forMoreInfo()
		os.Exit(0)
	}
}

func forMoreInfo() {
	fmt.Println("For more information, run: go run main.go help")
}

func printHelp() {
	helpSentences := []string{
		"",
		"Welcome to GoCrypt!",
		"Simple file encryption and decryption tool",
		"",
		"Usage:",
		"Run: go run main.go <function> /path/to/file",
		"",
		"Functions:",
		"help: Display help message",
		"encrypt: Encrypt a file with .encrypted extension",
		"decrypt: Decrypt a file with -decrypted file name extension",
		"",
		"Examples Usage:",
		"go run main.go help",
		"go run main.go encrypt fileSamples/image.png",
		"go run main.go decrypt fileSamples/image.png.encrypted",
		"",
	}
	for _, sentence := range helpSentences {
		fmt.Println(sentence)
	}
}

func encryptFile() {
	file := validateFile()

	password := getPassword()

	fmt.Println("\nEncrypting file...")
	filecrypt.EncryptFile(file, password)
	fmt.Println("File encrypted successfully!")
}

func decryptFile() {
	file := validateFile()

	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')

	fmt.Println("\nDecrypting file...")
	filecrypt.DecryptFile(file, password)
	fmt.Println("File decrypted successfully!")
}

func getPassword() string {
	fmt.Print("Enter password: ")
	password, _ := reader.ReadString('\n')
	fmt.Print("Confirm password: ")
	confirmPassword, _ := reader.ReadString('\n')

	if !isValidPassword(password, confirmPassword) {
		fmt.Print("\nPasswords do not match, please try again\n")
		getPassword()
	}

	return password
}

func isValidPassword(password, confirmPassword string) bool {
	return password == confirmPassword
}

func validateFile() (file string) {
	if len(os.Args) < 3 {
		fmt.Println("Error: No file path provided")
		forMoreInfo()
		os.Exit(0)
	}

	file = os.Args[2]
	if !isValidFile(file) {
		fmt.Println("Error: Invalid file path")
		forMoreInfo()
		os.Exit(0)
	}

	return file
}

func isValidFile(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
