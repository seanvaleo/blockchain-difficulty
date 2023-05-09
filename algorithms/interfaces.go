package algorithms

import "github.com/mesosoftware/blockchain-difficulty/blockchain"

type (
	// Algorithm is an interface for difficulty algorithms
	Algorithm interface {
		Name() string
		Window() uint64
		NextDifficulty(blockchain.Blockchain) uint64
	}
)
