package blockchain

import (
	"github.com/seanvaleo/dsim/pkg/dsim"
)

// Blockchain represents a blockchain object
type Blockchain struct {
	name      string
	chain     []*dsim.Block
	algorithm dsim.Algorithm
}

// New instantiates and returns a blockchain
func New(name string, algorithm dsim.Algorithm) *Blockchain {
	blockchain := &Blockchain{
		name:      name,
		chain:     make([]*dsim.Block, 0),
		algorithm: algorithm,
	}

	blockchain.addGenesisBlock()

	return blockchain
}

func (b *Blockchain) addGenesisBlock() {
	block := &dsim.Block{
		Height:     0,
		Difficulty: 10000,
		BlockTime:  1,
	}

	b.chain = append(b.chain, block)
}

// Name returns the name of the blockchain
func (b *Blockchain) Name() string {
	return b.name
}

// Algorithm returns the algorithm of the blockchain
func (b *Blockchain) Algorithm() dsim.Algorithm {
	return b.algorithm
}

// Length returns the length of the blockchain
func (b *Blockchain) Length() uint64 {
	return uint64(len(b.chain))
}

// AddBlock appends a block to the blockchain
func (b *Blockchain) AddBlock(hashPower uint64) {
	difficulty := b.algorithm.NextDifficulty(b.chain)

	block := &dsim.Block{
		Height:     b.Length(),
		Difficulty: difficulty,
		BlockTime:  float64(difficulty) / float64(hashPower),
	}

	b.chain = append(b.chain, block)
}

// GetBlock reads a block from the blockchain
func (b *Blockchain) GetBlock(height uint64) *dsim.Block {
	return b.chain[height]
}
