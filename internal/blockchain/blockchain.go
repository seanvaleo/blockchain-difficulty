package blockchain

import (
	"math"
	"math/big"

	"github.com/seanvaleo/dsim/pkg/dsim"
)

// Blockchain represents a blockchain object
type Blockchain struct {
	name      string
	chain     []*Block
	algorithm dsim.Algorithm
}

// Block represents a single block
type Block struct {
	height     uint64
	difficulty *big.Int
	blockTime  uint
}

// New instantiates and returns a blockchain
func New(name string, algorithm dsim.Algorithm) *Blockchain {
	return &Blockchain{
		name:      name,
		chain:     make([]*Block, 0),
		algorithm: algorithm,
	}
}

// Height returns the height of the blockchain
func (b *Blockchain) Height() uint64 {
	return uint64(len(b.chain)) - 1
}

// AddBlock appends a block to the blockchain
func (b *Blockchain) AddBlock(blockTime uint) {
	block := &Block{
		height:     b.Height(),
		difficulty: b.algorithm.NextDifficulty(),
		blockTime:  blockTime,
	}

	b.chain = append(b.chain, block)
}

// Name returns the name of the blockchain
func (b *Blockchain) Name() string {
	return b.name
}

// AlgorithmName returns the name of the blockchain's difficulty algorithm
func (b *Blockchain) AlgorithmName() string {
	return b.algorithm.Name()
}

// Statistics generates standard deviation and mean values for the block interval time
func (b *Blockchain) Statistics() (sd, mean float64) {

	var sum float64
	count := float64(len(b.chain))

	for _, v := range b.chain {
		sum += float64(v.blockTime)
	}
	mean = sum / count

	for _, v := range b.chain {
		sd += math.Pow(float64(v.blockTime)-mean, 2)
	}
	sd = math.Sqrt(sd / count)

	return sd, mean
}
