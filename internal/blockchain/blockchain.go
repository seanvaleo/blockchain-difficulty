package blockchain

import (
	"math/big"
	"time"

	"github.com/seanvaleo/dsim/pkg/dsim"
)

type Blockchain struct {
	name      string
	chain     []*Block
	algorithm dsim.Algorithm
}

type Block struct {
	height     uint64
	difficulty *big.Int
	timestamp  time.Time
}

func New(name string, algorithm dsim.Algorithm) *Blockchain {
	return &Blockchain{
		name:      name,
		chain:     make([]*Block, 0),
		algorithm: algorithm,
	}
}

func (b *Blockchain) Height() uint64 {
	return uint64(len(b.chain)) - 1
}

func (b *Blockchain) AddBlock() {
	block := &Block{
		height:     b.Height(),
		difficulty: b.algorithm.NextDifficulty(),
		timestamp:  time.Now(),
	}

	b.chain = append(b.chain, block)
}
