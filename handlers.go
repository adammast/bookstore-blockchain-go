package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// newBook handles the creation of a new book and assigns it an ID
func newBook(writer http.ResponseWriter, request *http.Request) {
	var book Book

	// Decode the book data from the request body
	if err := json.NewDecoder(request.Body).Decode(&book); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not create: %v", err)
		writer.Write([]byte("Could not create new book"))
		return
	}

	// Generate a unique ID for the book based on ISBN and PublishDate
	hash := md5.New()
	io.WriteString(hash, book.ISBN+book.PublishDate)
	book.ID = fmt.Sprintf("%x", hash.Sum(nil))

	// Marshal the book data and send it in the response
	response, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not marshal payload: %v", err)
		writer.Write([]byte("Could not save book data"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(response)
}

// writeBlock handles adding a new block to the blockchain
func writeBlock(writer http.ResponseWriter, request *http.Request) {
	var transaction BookTransaction

	// Decode the transaction data from the request body
	if err := json.NewDecoder(request.Body).Decode(&transaction); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Printf("Could not write block: %v", err)
		writer.Write([]byte("Could not write block"))
		return
	}

	// Add the transaction as a new block in the blockchain
	BlockChain.AddBlock(transaction)
}

// getBlockchain returns the entire blockchain in the response
func getBlockchain(writer http.ResponseWriter, request *http.Request) {
	bytes, err := json.MarshalIndent(BlockChain.blocks, "", " ")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(err)
		return
	}

	io.WriteString(writer, string(bytes))
}
