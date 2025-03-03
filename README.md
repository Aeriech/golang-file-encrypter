## **Getting Started**

### Project Overview
GoCrypt
This project is a secure file encryption and decryption utility built in Golang, utilizing AES-GCM (Galois/Counter Mode) for authenticated encryption. It allows users to encrypt files using a password-derived key and decrypt them seamlessly, ensuring data confidentiality and integrity. The tool automatically renames encrypted and decrypted files for clarity, preventing accidental overwrites.

Key Features:
✔️ AES-GCM Encryption – Secure and authenticated encryption with a randomly generated nonce.
✔️ PBKDF2 Key Derivation – Ensures strong encryption keys from user-provided passwords.
✔️ Automatic File Naming – Adds .encrypted when encrypting and -decrypted when decrypting.
✔️ File Integrity Preservation – Prevents modification or tampering by verifying authenticity.
✔️ Cross-Platform Compatibility – Works on Windows, macOS, and Linux.

Tech Stack:
   - Golang (Primary Language)
   - AES-GCM (Encryption Algorithm)
   - PBKDF2 (Key Derivation)

### Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/doc/install) (version 1.20 or higher)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Aeriech/golang-file-encrypter.git
   ```

2. Install dependencies:
   ```bash
   cd golang-file-encrypter
   go mod tidy
   ```
   
3. Run to display help message:
   ```bash
   go run main.go help
   ```

4. Sample usage:
   ```bash
   go run main.go encrypt fileSamples/image.png
   go run main.go decrypt fileSamples/image.png.encrypted
   ```