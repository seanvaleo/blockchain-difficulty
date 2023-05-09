package blockchain

import "github.com/mesosoftware/blockchain-difficulty/algorithms"

// Block represents a single block
type Block struct {
	Height     uint64
	Difficulty uint64
	BlockTime  float64
}

// Blockchain represents a chain of blocks
type Blockchain struct {
	name      string
	chain     []*Block
	algorithm algorithms.Algorithm
}

// New instantiates and returns a blockchain
func New(name string, algorithm algorithms.Algorithm) *Blockchain {
	blockchain := &Blockchain{
		name:      name,
		chain:     make([]*Block, 0),
		algorithm: algorithm,
	}

	blockchain.addGenesisBlock()

	return blockchain
}

func (b *Blockchain) addGenesisBlock() {
	block := &Block{
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
func (b *Blockchain) Algorithm() algorithms.Algorithm {
	return b.algorithm
}

// Length returns the length of the blockchain
func (b *Blockchain) Length() uint64 {
	return uint64(len(b.chain))
}

// AddBlock appends a block to the blockchain
func (b *Blockchain) AddBlock(hashPower uint64) {
	difficulty := b.algorithm.NextDifficulty(b.chain)

	block := &Block{
		Height:     b.Length(),
		Difficulty: difficulty,
		BlockTime:  float64(difficulty) / float64(hashPower),
	}

	b.chain = append(b.chain, block)
}

// GetBlock reads a block from the blockchain
func (b *Blockchain) GetBlock(height uint64) *Block {
	return b.chain[height]
}
