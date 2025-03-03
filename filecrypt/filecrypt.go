package filecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha1"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	// The file extension for encrypted files
	encryptedFileExtension     = ".encrypted"
	decryptedFileNameExtension = "-decrypted"
)

func EncryptFile(file string, password string) {
	validateFile(file)

	openFile, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}

	defer openFile.Close()

	readFile, err := io.ReadAll(openFile)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, 12)
	_, err = rand.Read(nonce)
	if err != nil {
		panic(err.Error())
	}

	derivedKey, err := pbkdf2.Key(sha1.New, password, nonce, 4096, 32)
	if err != nil {
		panic(err.Error())
	}

	cipherBlock, err := aes.NewCipher(derivedKey)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		panic(err.Error())
	}

	cipherText := aesgcm.Seal(nil, nonce, readFile, nil)
	cipherTextWithNonce := append(cipherText, nonce...)

	// Write encrypted data to a new file
	encryptedFile := file + encryptedFileExtension
	destinationFile, err := os.Create(encryptedFile)
	if err != nil {
		panic(err.Error())
	}

	defer destinationFile.Close()

	_, err = destinationFile.Write(cipherTextWithNonce)
	if err != nil {
		panic(err.Error())
	}
}

func DecryptFile(file string, password string) {
	validateFile(file)

	openFile, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}

	defer openFile.Close()

	ciphertext, err := io.ReadAll(openFile)
	if err != nil {
		panic(err.Error())
	}

	nonce := ciphertext[len(ciphertext)-12:]
	encryptedData := ciphertext[:len(ciphertext)-12]

	derivedKey, err := pbkdf2.Key(sha1.New, password, nonce, 4096, 32)
	if err != nil {
		panic(err.Error())
	}

	cipherBlock, err := aes.NewCipher(derivedKey)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		panic(err.Error())
	}

	plainText, err := aesgcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		panic(err.Error())
	}

	// Write decrypted data to a new file
	decryptedFile := getDecryptedFilename(file)
	destinationFile, err := os.Create(decryptedFile)
	if err != nil {
		panic(err.Error())
	}

	defer destinationFile.Close()

	_, err = destinationFile.Write(plainText)
	if err != nil {
		panic(err.Error())
	}
}

func validateFile(filePath string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		panic(err.Error())
	}
}

// getDecryptedFilename removes `.encrypted` and adds `-decrypted` before extension.
func getDecryptedFilename(file string) string {
	dir := filepath.Dir(file)
	base := filepath.Base(file)

	// Remove ".encrypted" suffix if it exists
	base = strings.TrimSuffix(base, encryptedFileExtension)

	// Extract filename and extension
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)

	// Add "-decrypted" before the extension
	newName := name + decryptedFileNameExtension + ext

	return filepath.Join(dir, newName)
}
