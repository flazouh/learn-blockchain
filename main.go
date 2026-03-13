package main

import (
	"fmt"
	"hello-world/blockchain"
)

func main() {
	// Create a new blockchain with difficulty 3 (hash must start with "000")
	bc := blockchain.New(3)

	// Add a block with two transactions
	bc.Add([]blockchain.Transaction{
		{From: "Alice", To: "Bob", Amount: 50},
		{From: "Bob", To: "Charlie", Amount: 20},
	})

	bc.Add([]blockchain.Transaction{
		{From: "Charlie", To: "Alice", Amount: 10},
	})

	// Validate the chain — IsValid() returns an error or nil
	if err := bc.IsValid(); err != nil {
		fmt.Println("Chain is invalid:", err)
	} else {
		fmt.Println("Chain is valid!")
	}

	// Print all blocks
	for _, block := range bc.Blocks {
		fmt.Printf("\nBlock %d | Hash: %s... | Txs: %d\n", block.Index, block.Hash[:12], len(block.Transactions))
	}
}
