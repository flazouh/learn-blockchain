# learn-blockchain

**Build a mini blockchain in Go from scratch.** No frameworks, no fluff — just hashes, blocks, and proof-of-work in ~200 lines.

```
Block 0 mined! Nonce: 0 Hash: 5feceb66...
Block 1 mined! Nonce: 42 Hash: 000a1b2c...
Block 2 mined! Nonce: 17 Hash: 000c3d4e...
Chain is valid!
```

## Quick start

**Requires:** Go 1.25+ (or any compatible version)

```bash
git clone https://github.com/flazouh/learn-blockchain.git
cd learn-blockchain
go run .
```

You’ll see a chain created, transactions added, blocks mined (SHA-256 with leading zeros), and a quick validation. That’s the whole demo.

## What’s inside

| Piece | What it does |
|-------|----------------|
| **Block** | Index, timestamp, transactions, previous hash, nonce. Hash = SHA-256 of the lot. |
| **Mining** | Brute-force nonce until the hash starts with `0` × difficulty (e.g. `000` for difficulty 3). |
| **Chain** | Genesis block + append-only blocks; each block links to the previous via `PrevHash`. |
| **Validator** | Walks the chain, checks every hash and link — `IsValid() error` tells you if something’s broken. |

So: **blocks**, **proof-of-work**, and **chain integrity**. No networking, no consensus, no wallets — ideal for learning the core ideas.

## Project layout

```
.
├── main.go              # Demo: create chain, add txs, mine, validate, print
├── direction.go         # Small Go examples
├── greet/
│   └── greet.go
├── blockchain/
│   ├── block.go         # Block struct, ComputeHash(), Mine(difficulty)
│   ├── chain.go         # Blockchain, New(), Add(), IsValid()
│   ├── transaction.go   # From, To, Amount
│   └── validator.go     # Validator interface
└── go.mod
```

## Run the tests

```bash
go test ./blockchain/...
```

Covers genesis block, `Add`, validation (valid chain, invalid hash, broken link), and mining.

## Contributing

Pull requests and ideas are welcome. Open an issue to discuss bigger changes, or just fork, tweak, and send a PR. Keep the scope small and educational.

## License

This project is for educational use. Use, modify, and share as you like. If you add a `LICENSE` file to the repo, you can mention it here (e.g. MIT, Apache 2.0).

## Disclaimer

Educational toy only. No persistence, no signatures, no mempool, no P2P — just the data structures and mining logic to get the concepts clear.

## Possible next steps

- CLI flags for difficulty and custom transactions  
- JSON export for blocks  
- More tests or edge cases  

**Have fun breaking and extending it.**
