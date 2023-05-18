package algorithms

import (
	"fmt"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
)

// LWMA implements a Linear Weighted Moving Average equation, using the
// linearly weighted average block time of the most recent X blocks to
// estimate a more suitable difficulty
type LWMA struct {
	name                   string
	intervalBlocks         int // Frequency of difficulty re-calculation
	windowBlocks           int // Block data in sample
	nextRecalculationBlock int
}

// NewLWMA instantiates and returns a new LWMA
func NewLWMA(intervalBlocks, windowBlocks int) *LWMA {
	return &LWMA{
		name: fmt.Sprintf("LWMA: Recalculate at every %v blocks using a %v block window. Target is %ds",
			intervalBlocks,
			windowBlocks,
			internal.Config.TargetBlockTimeSeconds),
		intervalBlocks:         intervalBlocks,
		windowBlocks:           windowBlocks,
		nextRecalculationBlock: intervalBlocks,
	}
}

// Name returns the algorithm name
func (l *LWMA) Name() string {
	return l.name
}

// NextDifficulty calculates the next difficulty
// We must account for the block time of the current block not yet added
func (l *LWMA) NextDifficulty(blockchain blockchain.Blockchain, thisBlockTime uint) uint64 {
	lenBlocks := blockchain.GetLength()

	if lenBlocks == 0 {
		return blockchain.StartDifficulty
	}

	// Don't start calculating until we have a complete window (including this block)
	if lenBlocks < l.windowBlocks-1 {
		return blockchain.GetLastBlock().NextDifficulty
	}

	// Only recalculate on the desired interval
	if lenBlocks < l.nextRecalculationBlock-1 {
		return blockchain.GetLastBlock().NextDifficulty
	}

	l.nextRecalculationBlock += l.intervalBlocks

	lwmaD, lwmaBT := l.lwma(blockchain, thisBlockTime)

	return uint64(lwmaD * (float64(internal.Config.TargetBlockTimeSeconds) / lwmaBT))
}

// lwma calculates the Linear Weighted Moving Averages for Difficulty and BlockTime
func (l *LWMA) lwma(blockchain blockchain.Blockchain, thisBlockTime uint) (lwmaD, lwmaBT float64) {
	var sumBT, sumD float64

	i := blockchain.GetLength()                      // Number of last block added (not including This block)
	j := blockchain.GetLength() - l.windowBlocks + 1 // Number of first block in window

	weight := float64(l.windowBlocks) // Weight starts at window size and decreases to 1
	sumWeights := float64(0)          // Sum of all Weights

	// Add values from the current block being processed
	sumBT += float64(thisBlockTime) * weight
	sumD += float64(blockchain.GetBlock(i).NextDifficulty) * weight
	sumWeights += weight
	weight--

	// Add the rest of the values
	for ; i > j; i-- {
		sumBT += (float64(blockchain.GetBlock(i).BlockTimeSeconds) * weight)
		sumD += (float64(blockchain.GetBlock(i).ThisDifficulty) * weight)
		sumWeights += weight
		weight--
	}
	lwmaBT = sumBT / sumWeights
	lwmaD = sumD / sumWeights

	return
}
