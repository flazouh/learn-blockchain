package blockchain

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

// Block is a struct — Go's equivalent of a class.
// It groups related data together.
type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []Transaction // slice of Transaction structs (composition)
	PrevHash     string
	Hash         string
	Nonce        int
}

// ComputeHash is a method on Block.
// The receiver (b Block) means it belongs to Block, like a class method.
func (b Block) ComputeHash() string {
	data := fmt.Sprintf("%d%s%v%s%d", b.Index, b.Timestamp, b.Transactions, b.PrevHash, b.Nonce)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// Mine is a method with a pointer receiver (*Block).
// We use *Block instead of Block because we're mutating the struct (changing Nonce and Hash).
func (b *Block) Mine(difficulty int) {
	target := strings.Repeat("0", difficulty)
	for !strings.HasPrefix(b.Hash, target) {
		b.Nonce++
		b.Hash = b.ComputeHash()
	}
	fmt.Printf("Block %d mined! Nonce: %d Hash: %s\n", b.Index, b.Nonce, b.Hash[:12])
}
