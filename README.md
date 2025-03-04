## **Getting Started**

### Project Overview - GoCrypt
GoCrypt is a lightweight, command-line file encryption and decryption tool built in Golang, designed to secure sensitive files using AES-GCM encryption with PBKDF2-based key derivation. It provides an easy way to encrypt files with a password and decrypt them when needed, ensuring data security and integrity.

Key Features:
   - AES-GCM Encryption - Provides strong encryption with authentication.
   - PBKDF2 Key Derivation - Strengthens password security with salt and iterations.
   - Simple CLI Interface - Encrypt and decrypt files with intuitive commands.
   - Automatic File Naming -
      - Adds .encrypted when encrypting.
      - Renames to -decrypted before the file extension when decrypting.
   - Cross-Platform Support - Works on Windows, macOS, and Linux.

Tech Stack:
   - Golang (Primary Language)
   - AES-GCM (Encryption Algorithm)
   - PBKDF2 (Key Derivation)
   - CLI (Command-Line Interface)

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
