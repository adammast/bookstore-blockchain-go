package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var BlockChain *Blockchain

// Initialize the blockchain and start the server
func main() {
	// Create a new blockchain
	BlockChain = NewBlockchain()

	// Initialize the HTTP router
	r := mux.NewRouter()

	// Define the routes and their handlers
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newBook).Methods("POST")

	// Run a goroutine to log blockchain details
	go logBlockchainDetails()

	// Start the server
	log.Println("Listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}

// logBlockchainDetails logs details of each block in the blockchain
func logBlockchainDetails() {
	for _, block := range BlockChain.blocks {
		log.Printf("Previous hash: %x\n", block.PreviousHash)
		bytes, _ := json.MarshalIndent(block.Data, "", " ")
		log.Printf("Data: %v\n", string(bytes))
		log.Printf("Hash: %x\n", block.Hash)
		log.Println()
	}
}
