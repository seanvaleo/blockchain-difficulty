package dsim

import "math/big"

type (
	// Blockchain is an interface for blockchains
	Blockchain interface {
		Height() uint64
		AddBlock(blockTime uint)
		Name() string
		AlgorithmName() string
		Statistics() (sd, mean float64)
	}
	// Algorithm is an interface for difficulty algorithms
	Algorithm interface {
		Name() string
		NextDifficulty() *big.Int
	}
)
