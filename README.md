# Blockchain in Go

This project demonstrates a simple implementation of a blockchain using Go. It simulates a decentralized ledger for tracking books and their transactions. The blockchain stores book transactions in blocks, where each block is cryptographically linked to its predecessor.

## Features
- **Book Creation**: Allows the creation of books and generates a unique book ID.
- **Blockchain**: Stores book transactions as blocks, with each block being validated.
- **Genesis Block**: The first block in the blockchain, used to initialize the chain.
- **API Endpoints**:
  - `GET /`: Returns the entire blockchain.
  - `POST /`: Adds a new block to the blockchain.
  - `POST /new`: Creates a new book and returns its details.

## Setup

### Prerequisites
- Go (version 1.18 or higher)

### Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/blockchain-go.git
    cd blockchain-go
    ```

2. Install dependencies (if any):
    ```bash
    go mod tidy
    ```

### Running the Server

Run the application:
```bash
go run .
```
This will start a local server on `http://localhost:3000`.

### API Example

- Create a new book:
    - `POST /new` with JSON payload:
        ```json
        {
            "title": "The Go Programming Language",
            "author": "Alan A. A. Donovan",
            "publish_date": "2015-10-26",
            "isbn": "9780134190440"
        }
        ```

- Add a new transaction (book purchase):
    - `POST /` with JSON payload:
        ```json
        {
            "book_id": "book_id_123",
            "user": "John Doe",
            "purchase_date": "2025-03-16",
            "is_genesis": false
        }
        ```

- Get blockchain:
    - `GET /` to retrieve the current blockchain.

## Libraries Used

- **[gorilla/mux](https://github.com/gorilla/mux)**: A powerful URL router and dispatcher for Go, used to manage HTTP routes in the project.
- **[crypto/md5](https://pkg.go.dev/crypto/md5)**: Part of the Go standard library, used for generating MD5 hashes to create unique book IDs.
- **[crypto/sha256](https://pkg.go.dev/crypto/sha256)**: Part of the Go standard library, used for generating SHA-256 hashes for block validation.
- **[encoding/json](https://pkg.go.dev/encoding/json)**: Used for encoding and decoding JSON data, which allows communication with the API.
- **[encoding/hex](https://pkg.go.dev/encoding/hex)**: Used to encode and decode hexadecimal strings, specifically for generating book IDs and hash outputs.
- **[sync](https://pkg.go.dev/sync)**: Used for concurrency and synchronization.
- **[log](https://pkg.go.dev/log)**: Used for logging messages to standard output and error output.
- **[net/http](https://pkg.go.dev/net/http)**: Part of the Go standard library, used to create and manage the HTTP server and handle requests.
- **[fmt](https://pkg.go.dev/fmt)**: Used for formatted I/O operations such as printing to the console.
- **[io](https://pkg.go.dev/io)**: Provides basic I/O functionality, including writing and reading data to and from the HTTP response.

## License

This project is licensed under the MIT License.