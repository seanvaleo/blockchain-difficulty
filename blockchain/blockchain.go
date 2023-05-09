package blockchain

// Block represents a single block
type Block struct {
	Height           uint64
	NextDifficulty   uint64
	BlockTimeSeconds uint // Time differential since last block
}

// Blockchain represents a chain of blocks
type Blockchain struct {
	Chain           []*Block
	StartDifficulty uint64
}

// New instantiates and returns a blockchain
func New(startDifficulty uint64) Blockchain {
	blockchain := Blockchain{
		Chain:           make([]*Block, 0),
		StartDifficulty: startDifficulty,
	}

	return blockchain
}

// Length returns the length of the blockchain
func (b *Blockchain) GetLength() uint64 {
	return uint64(len(b.Chain))
}

// AddBlock appends a block to the blockchain
func (b *Blockchain) AddBlock(nextDifficulty uint64, blockTimeSeconds uint) {
	block := &Block{
		Height:           b.GetLength(),
		NextDifficulty:   nextDifficulty,
		BlockTimeSeconds: blockTimeSeconds,
	}

	b.Chain = append(b.Chain, block)
}

// GetBlock reads a block from the blockchain
func (b *Blockchain) GetBlock(height uint64) *Block {
	return b.Chain[height]
}

// GetLastBlock reads a block from the blockchain
func (b *Blockchain) GetLastBlock() *Block {
	if len(b.Chain) == 0 {
		return nil
	}

	return b.Chain[len(b.Chain)-1]
}
