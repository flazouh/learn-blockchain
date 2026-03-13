# hello-world-go

A small Go playground project featuring a minimal blockchain demo.

## What this project is

This repository is a learning-oriented Go project with:

- A toy blockchain implementation (`blockchain/`)
- A runnable demo in `main.go`
- A couple of basic Go examples (`greet/`, `direction.go`)

## Requirements

- Go `1.25.0` (or compatible newer version)

## Run

From the project root:

```bash
go run .
```

Expected behavior:

- Creates a blockchain
- Adds sample transactions
- Mines blocks using a simple proof-of-work rule
- Validates chain integrity
- Prints a short summary of blocks

## Project structure

```text
.
├── main.go
├── direction.go
├── greet/
│   └── greet.go
├── blockchain/
│   ├── block.go
│   ├── chain.go
│   ├── transaction.go
│   └── validator.go
└── go.mod
```

## Blockchain package overview

- `block.go`
  - Defines `Block`
  - Computes SHA-256 hash
  - Mines by incrementing nonce until hash has a prefix of `0`s

- `chain.go`
  - Defines `Blockchain`
  - Creates genesis block
  - Adds new blocks with transactions
  - Validates hash and previous-hash links

- `transaction.go`
  - Defines `Transaction` (`From`, `To`, `Amount`)

- `validator.go`
  - Defines `Validator` interface (`IsValid() error`)

## Notes

- This is an educational example, not a production blockchain.
- There is no networking, consensus, signatures, mempool, or persistence.

## Next ideas

- Add unit tests for `IsValid` and mining behavior
- Add CLI flags for difficulty and transaction input
- Add JSON export for blocks
