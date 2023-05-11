package algorithms

import "github.com/mesosoftware/blockchain-difficulty/blockchain"

type (
	// Algorithm is an interface for difficulty algorithms
	Algorithm interface {
		Name() string
		NextDifficulty(blockchain blockchain.Blockchain, thisBlockTime uint) uint64
	}
)
