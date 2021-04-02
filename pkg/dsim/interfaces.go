package dsim

import "math/big"

type (
	Blockchain interface {
		Height() uint64
		AddBlock(blockTime uint)
		Name() string
		AlgorithmName() string
		Statistics() (sd, mean float64)
	}
	Algorithm interface {
		Name() string
		NextDifficulty() *big.Int
	}
)
