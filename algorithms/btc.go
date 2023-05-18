package algorithms

import (
	"fmt"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
)

// BTC implements the Bitcoin difficulty adjustment algorithm, which is:
// New Difficulty = Old Difficulty * ((2016 * 10 minutes) / Actual Time for last 2016 blocks)
// to estimate a more suitable difficulty
type BTC struct {
	name                   string
	intervalBlocks         int // Frequency of difficulty re-calculation
	windowBlocks           int // Block data in sample
	nextRecalculationBlock int
}

// NewBTC instantiates and returns a new BTC
func NewBTC() *BTC {
	return &BTC{
		name: fmt.Sprintf("Bitcoin: Recalculate at every 2016 blocks using a 2016 block window. Target is %ds",
			internal.Config.TargetBlockTimeSeconds),
		intervalBlocks:         2016, // fixed
		windowBlocks:           2016, // fixed
		nextRecalculationBlock: 2016, // fixed
	}
}

// Name returns the algorithm name
func (b *BTC) Name() string {
	return b.name
}

// NextDifficulty calculates the next difficulty
// We must account for the block time of the current block not yet added
func (b *BTC) NextDifficulty(blockchain blockchain.Blockchain, thisBlockTime uint) uint64 {
	lenBlocks := blockchain.GetLength()

	if lenBlocks == 0 {
		return blockchain.StartDifficulty
	}

	// Don't start calculating until we have a complete window (including this block)
	if lenBlocks < b.windowBlocks-1 {
		return blockchain.GetLastBlock().NextDifficulty
	}

	// Only recalculate on the desired interval
	if lenBlocks < b.nextRecalculationBlock-1 {
		return blockchain.GetLastBlock().NextDifficulty
	}

	b.nextRecalculationBlock += b.intervalBlocks

	time := b.sumBlockTimes(blockchain, thisBlockTime)

	return uint64(float64(blockchain.GetLastBlock().NextDifficulty) * (float64(b.windowBlocks*600) / float64(time)))
}

func (b *BTC) sumBlockTimes(blockchain blockchain.Blockchain, thisBlockTime uint) (sum uint) {
	i := blockchain.GetLength()                      // Number of last block added (not including This block)
	j := blockchain.GetLength() - b.windowBlocks + 1 // Number of first block in window

	// Add values from the current block being processed
	sum += thisBlockTime

	// Add the rest of the values
	for ; i > j; i-- {
		sum += blockchain.GetBlock(i).BlockTimeSeconds
	}

	return
}
