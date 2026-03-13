package blockchain

import (
	"testing"
)

func TestNew(t *testing.T) {
	bc := New(2)
	if len(bc.Blocks) != 1 {
		t.Fatalf("expected 1 block (genesis), got %d", len(bc.Blocks))
	}
	if bc.Blocks[0].Index != 0 {
		t.Errorf("genesis index want 0, got %d", bc.Blocks[0].Index)
	}
	if bc.Blocks[0].PrevHash != "0" {
		t.Errorf("genesis PrevHash want \"0\", got %q", bc.Blocks[0].PrevHash)
	}
	if bc.Difficulty != 2 {
		t.Errorf("Difficulty want 2, got %d", bc.Difficulty)
	}
}

func TestIsValid_ValidChain(t *testing.T) {
	bc := New(1) // low difficulty for fast tests
	bc.Add([]Transaction{{From: "A", To: "B", Amount: 10}})
	bc.Add([]Transaction{{From: "B", To: "C", Amount: 5}})

	if err := bc.IsValid(); err != nil {
		t.Errorf("expected valid chain, got error: %v", err)
	}
}

func TestIsValid_InvalidHash(t *testing.T) {
	bc := New(1)
	bc.Add([]Transaction{{From: "A", To: "B", Amount: 10}})
	bc.Blocks[1].Hash = "tampered"

	if err := bc.IsValid(); err == nil {
		t.Fatal("expected error for invalid hash, got nil")
	}
}

func TestIsValid_BrokenLink(t *testing.T) {
	bc := New(1)
	bc.Add([]Transaction{{From: "A", To: "B", Amount: 10}})
	bc.Blocks[1].PrevHash = "wrong-prev-hash"

	if err := bc.IsValid(); err == nil {
		t.Fatal("expected error for broken chain link, got nil")
	}
}

func TestAdd(t *testing.T) {
	bc := New(1)
	bc.Add([]Transaction{{From: "Alice", To: "Bob", Amount: 50}})

	if len(bc.Blocks) != 2 {
		t.Fatalf("expected 2 blocks after Add, got %d", len(bc.Blocks))
	}
	b := bc.Blocks[1]
	if b.Index != 1 {
		t.Errorf("block index want 1, got %d", b.Index)
	}
	if b.PrevHash != bc.Blocks[0].Hash {
		t.Errorf("PrevHash does not match previous block hash")
	}
	if len(b.Transactions) != 1 || b.Transactions[0].Amount != 50 {
		t.Errorf("transactions not stored correctly")
	}
}

func TestBlock_Mine(t *testing.T) {
	bc := New(2)
	bc.Add([]Transaction{{From: "A", To: "B", Amount: 1}})

	b := bc.Blocks[1]
	target := "00"
	if len(b.Hash) < 2 {
		t.Fatalf("hash too short: %s", b.Hash)
	}
	if b.Hash[:2] != target {
		t.Errorf("mined hash should start with %q, got %q", target, b.Hash[:2])
	}
}
