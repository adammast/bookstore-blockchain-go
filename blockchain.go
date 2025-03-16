package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Mutex to ensure thread safety when modifying the blockchain
var blockchainMutex sync.Mutex

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(transaction BookTransaction) {
	// Lock blockchain to ensure thread safety
	blockchainMutex.Lock()
	defer blockchainMutex.Unlock()

	// Get the previous block to chain the new block
	prevBlock := bc.blocks[len(bc.blocks)-1]

	// Create a new block and validate it
	block := CreateBlock(prevBlock, transaction)
	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

// validBlock validates the integrity of a block based on its data
func validBlock(block, prevBlock *Block) bool {
	if prevBlock.Hash != block.PreviousHash {
		return false
	}
	if prevBlock.Position+1 != block.Position {
		return false
	}
	if !block.validateHash(block.Hash) {
		return false
	}
	return true
}

// validateHash checks if the block's hash is valid
func (block *Block) validateHash(hash string) bool {
	block.generateHash()
	return block.Hash == hash
}

// CreateBlock generates a new block based on a previous block and a transaction
func CreateBlock(prevBlock *Block, transaction BookTransaction) *Block {
	block := &Block{
		Position:     prevBlock.Position + 1,
		Data:         transaction,
		Timestamp:    time.Now().String(),
		PreviousHash: prevBlock.Hash,
	}
	block.generateHash()
	return block
}

// generateHash calculates the hash of the block based on its data
func (block *Block) generateHash() {
	bytes, _ := json.Marshal(block.Data)
	data := fmt.Sprint(block.Position) + block.Timestamp + string(bytes) + block.PreviousHash
	hash := sha256.New()
	hash.Write([]byte(data))
	block.Hash = hex.EncodeToString(hash.Sum(nil))
}

// NewBlockchain initializes a new blockchain with the genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

// GenesisBlock generates the first block in the blockchain
func GenesisBlock() *Block {
	return CreateBlock(&Block{}, BookTransaction{IsGenesis: true})
}
