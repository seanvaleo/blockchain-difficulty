package blockchain

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Block represents a single block
type Block struct {
	Height           int    // Block number
	ThisDifficulty   uint64 // Difficulty when mining this block
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
func (b *Blockchain) GetLength() int {
	return int(len(b.Chain))
}

// AddBlock appends a block to the blockchain
func (b *Blockchain) AddBlock(thisDifficulty, nextDifficulty uint64, blockTimeSeconds uint) {
	block := &Block{
		Height:           b.GetLength() + 1,
		ThisDifficulty:   thisDifficulty,
		NextDifficulty:   nextDifficulty,
		BlockTimeSeconds: blockTimeSeconds,
	}

	b.Chain = append(b.Chain, block)
	fmt.Println(block)
}

// GetBlock reads a block from the blockchain
func (b *Blockchain) GetBlock(height int) *Block {
	if height-1 > len(b.Chain) {
		log.Errorf("Block does not exist in blockchain")
		os.Exit(1)
	}

	return b.Chain[height-1]
}

// GetFirstBlock reads the first block from the blockchain
func (b *Blockchain) GetFirstBlock() *Block {
	if len(b.Chain) == 0 {
		log.Errorf("No blocks in blockchain")
		os.Exit(1)
	}

	return b.Chain[0]
}

// GetLastBlock reads the last block from the blockchain
func (b *Blockchain) GetLastBlock() *Block {
	if len(b.Chain) == 0 {
		log.Errorf("No blocks in blockchain")
		os.Exit(1)
	}

	return b.Chain[len(b.Chain)-1]
}
