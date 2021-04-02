package dsim

import "math/big"

type (
	Blockchain interface {
		Height() uint64
		AddBlock()
	}
	Algorithm interface {
		Name() string
		NextDifficulty() *big.Int
	}
)
