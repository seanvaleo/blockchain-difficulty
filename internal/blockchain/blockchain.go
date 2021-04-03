package blockchain

import (
	"math"

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

// Height returns the height of the blockchain
func (b *Blockchain) Height() uint64 {
	return uint64(len(b.chain))
}

// AddBlock appends a block to the blockchain
func (b *Blockchain) AddBlock(blockTime uint64) {
	block := &dsim.Block{
		Height:     b.Height(),
		Difficulty: b.algorithm.NextDifficulty(b.chain),
		BlockTime:  blockTime,
	}

	b.chain = append(b.chain, block)
}

// Difficulty returns the difficulty of the lastblock
func (b *Blockchain) Difficulty() uint64 {
	return b.chain[b.Height()-1].Difficulty
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
		sum += float64(v.BlockTime)
	}
	mean = sum / count

	for _, v := range b.chain {
		sd += math.Pow(float64(v.BlockTime)-mean, 2)
	}
	sd = math.Sqrt(sd / count)

	return sd, mean
}
