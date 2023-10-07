// main.go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"blockchain/assignment01bca"
)

func main() {
	// Create a blockchain and add some blocks
	blockchain := []*assignment01bca.Block{}
	genesisBlock := assignment01bca.NewBlock("Genesis Transaction", 0, "")
	blockchain = append(blockchain, genesisBlock)

	// Add more blocks with random nonces
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	block1 := assignment01bca.NewBlock("Alice to Bob", rand.Int(), genesisBlock.CurrentHash)
	blockchain = append(blockchain, block1)

	block2 := assignment01bca.NewBlock("Bob to Carol", rand.Int(), block1.CurrentHash)
	blockchain = append(blockchain, block2)

	block3 := assignment01bca.NewBlock("Carol to Susan", rand.Int(), block2.CurrentHash)
	blockchain = append(blockchain, block3)

	block4 := assignment01bca.NewBlock("Susan to Mark", rand.Int(), block3.CurrentHash)
	blockchain = append(blockchain, block4)

	block5 := assignment01bca.NewBlock("Mark to Alice", rand.Int(), block4.CurrentHash)
	blockchain = append(blockchain, block5)

	block6 := assignment01bca.NewBlock("Hello, I am Khadija", rand.Int(), block5.CurrentHash)
	blockchain = append(blockchain, block6)

	// Display all blocks in the blockchain
	assignment01bca.DisplayBlocks(blockchain)

	// Attempt to modify a block's transaction (Simulating an invalid blockchain)
	/* modifiedBlockIndex := 2
	modifiedTransaction := "Modified Transaction"
	assignment01bca.ChangeBlock(blockchain[modifiedBlockIndex], modifiedTransaction)
	fmt.Printf("Modifications made at Block %d: The modification made was: %s\n", modifiedBlockIndex, modifiedTransaction)
	*/

	// Verify the blockchain's integrity
	isValid := assignment01bca.VerifyChain(blockchain)
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
