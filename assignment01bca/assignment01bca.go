package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Block represents a single block in the blockchain.
type Block struct {
	Transaction   string
	Nonce         int
	PreviousHash  string
	CurrentHash   string
}

// NewBlock adds a new block to the blockchain with the given transaction, nonce, and previous hash.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.CalculateHash()
	return block
}

// calculateHash calculates and sets the hash of the block based on its data.
func (b *Block) CalculateHash() {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hashBytes := sha256.Sum256([]byte(data))
	b.CurrentHash = hex.EncodeToString(hashBytes[:])
}

// DisplayBlocks prints all the blocks in the blockchain in a nice format.
func DisplayBlocks(blocks []*Block) {
	fmt.Println("Blockchain:")
	for _, block := range blocks {
		block.displayBlockInfo()
	}
}

// ChangeBlockTransaction updates the transaction of a given block.
func ChangeBlock(block *Block, newTransaction string){

	// Save the original hash for later comparison
	originalHash := block.CurrentHash

	block.Transaction = newTransaction
	block.CalculateHash() // Recalculate the hash after changing the transaction

	// Check if the recalculated hash is different from the original hash
	if block.CurrentHash != originalHash {
		fmt.Printf("Error: Invalid transaction in Block: %s\n", newTransaction)
	}
}


// VerifyChain verifies the integrity of the blockchain in case any changes are made.
func VerifyChain(blocks []*Block) bool {
	for i := 1; i < len(blocks); i++ {
		currentBlock := blocks[i]
		previousBlock := blocks[i-1]

		// Verify hash
		if currentBlock.CurrentHash != currentBlock.calculateHashValue() {
			return false
		}

		// Verify previous hash reference
		if currentBlock.PreviousHash != previousBlock.CurrentHash {
			return false
		}
	}
	return true
}

// CalculateHash calculates the hash of a given string.
func CalculateHash(stringToHash string) string {
	hashBytes := sha256.Sum256([]byte(stringToHash))
	return hex.EncodeToString(hashBytes[:])
}

// calculateHashValue calculates and returns the hash of the block's data.
func (b *Block) calculateHashValue() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hashBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashBytes[:])
}

// displayBlockInfo prints the block's information.
func (b *Block) displayBlockInfo() {
	fmt.Println("Transaction:", b.Transaction)
	fmt.Println("Nonce:", b.Nonce)
	fmt.Println("Previous Hash:", b.PreviousHash)
	fmt.Println("Current Hash:", b.CurrentHash)
	fmt.Println("--------------------------")
}
