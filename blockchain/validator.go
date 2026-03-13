package blockchain

// Validator is an interface — it defines behavior, not data.
// Any type that has an IsValid() method automatically implements this interface.
// This is Go's way of polymorphism — no "implements" keyword needed.
type Validator interface {
	IsValid() error
}
