package main

// Book represents a book with basic details
type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}

// BookTransaction represents a transaction involving a book
type BookTransaction struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	PurchaseDate string `json:"purchase_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

// Block represents a block in the blockchain
type Block struct {
	Position     int
	Data         BookTransaction
	Timestamp    string
	Hash         string
	PreviousHash string
}

// Blockchain represents the entire chain of blocks
type Blockchain struct {
	blocks []*Block
}
