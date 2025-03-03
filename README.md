## **Getting Started**

### Project Overview - GoCrypt
This project is a CLI-based chatbot built in Golang, leveraging Google's Gemini API for natural language processing (NLP). The chatbot allows users to interact with Gemini AI in a conversational manner, processing user inputs and generating intelligent responses. The application manages chat sessions, retrieves an API key from environment variables, and continuously accepts user input until the conversation is ended.

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