package blockchain

// Transaction represents a transfer from one address to another.
// This is a plain struct — Go's equivalent of a class with only data fields.
type Transaction struct {
	From   string
	To     string
	Amount float64
}
