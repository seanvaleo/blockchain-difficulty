package blockchain

// Block represents a single block
type Block struct {
	Height         uint64
	NextDifficulty uint64
	BlockTime      float64 // Time differential since last block
}

// Blockchain represents a chain of blocks
type Blockchain struct {
	Chain []*Block
}

// New instantiates and returns a blockchain
func New() Blockchain {
	blockchain := Blockchain{
		Chain: make([]*Block, 0),
	}

	blockchain.addGenesisBlock()

	return blockchain
}

func (b *Blockchain) addGenesisBlock() {
	block := &Block{
		Height:         0,
		NextDifficulty: 10000,
		BlockTime:      1,
	}

	b.Chain = append(b.Chain, block)
}

// Length returns the length of the blockchain
func (b *Blockchain) GetLength() uint64 {
	return uint64(len(b.Chain))
}

// AddBlock appends a block to the blockchain
func (b *Blockchain) AddBlock(nextDifficulty uint64, blockTime float64) {
	block := &Block{
		Height:         b.GetLength(),
		NextDifficulty: nextDifficulty,
		BlockTime:      blockTime,
	}

	b.Chain = append(b.Chain, block)
}

// GetBlock reads a block from the blockchain
func (b *Blockchain) GetBlock(height uint64) *Block {
	return b.Chain[height]
}

// GetLastBlock reads a block from the blockchain
func (b *Blockchain) GetLastBlock() *Block {
	return b.Chain[len(b.Chain)-1]
}
