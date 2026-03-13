package blockchain

import (
	"errors"
	"fmt"
	"time"
)

// Blockchain holds a slice of blocks and the mining difficulty.
type Blockchain struct {
	Blocks     []Block
	Difficulty int
}

// New creates a new Blockchain with a genesis block (the first block).
// This is Go's convention for constructors — a function named New.
func New(difficulty int) *Blockchain {
	genesis := Block{
		Index:     0,
		Timestamp: time.Now(),
		PrevHash:  "0",
	}
	genesis.Hash = genesis.ComputeHash()

	return &Blockchain{
		Blocks:     []Block{genesis},
		Difficulty: difficulty,
	}
}

// Add creates a new block with the given transactions, mines it, and appends it.
// Pointer receiver (*Blockchain) because we're modifying the chain.
func (bc *Blockchain) Add(transactions []Transaction) {
	prev := bc.Blocks[len(bc.Blocks)-1]

	block := Block{
		Index:        prev.Index + 1,
		Timestamp:    time.Now(),
		Transactions: transactions,
		PrevHash:     prev.Hash,
	}

	block.Mine(bc.Difficulty)
	bc.Blocks = append(bc.Blocks, block)
}

// IsValid checks the chain integrity.
// It returns an error if something is wrong — idiomatic Go error handling.
// Blockchain implements the Validator interface because it has this method.
func (bc *Blockchain) IsValid() error {
	for i := 1; i < len(bc.Blocks); i++ {
		current := bc.Blocks[i]
		prev := bc.Blocks[i-1]

		if current.Hash != current.ComputeHash() {
			return errors.New("invalid hash at block " + fmt.Sprint(i))
		}
		if current.PrevHash != prev.Hash {
			return errors.New("broken chain at block " + fmt.Sprint(i))
		}
	}
	return nil // nil means no error in Go
}
