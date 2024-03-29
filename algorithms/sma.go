package algorithms

import (
	"fmt"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
)

// SMA implements a Simple Moving Average equation, using the average
// block time of the most recent X blocks to estimate a more suitable
// difficulty
type SMA struct {
	name                   string
	target                 uint
	intervalBlocks         int // Frequency of difficulty re-calculation
	windowBlocks           int // Block data in sample
	nextRecalculationBlock int
}

// NewSMA instantiates and returns a new SMA
func NewSMA(target uint, intervalBlocks, windowBlocks int) *SMA {
	return &SMA{
		name: fmt.Sprintf("SMA: Recalculate at every %v blocks using a %v block window. Target is %ds",
			intervalBlocks,
			windowBlocks,
			target),
		target:                 target,
		intervalBlocks:         intervalBlocks,
		windowBlocks:           windowBlocks,
		nextRecalculationBlock: intervalBlocks,
	}
}

// Name returns the algorithm name
func (s *SMA) Name() string {
	return s.name
}

// NextDifficulty calculates the next difficulty
// We must account for the block time of the current block not yet added
func (s *SMA) NextDifficulty(blockchain blockchain.Blockchain, thisBlockTime uint) uint64 {
	lenBlocks := blockchain.GetLength()

	if lenBlocks == 0 {
		return blockchain.StartDifficulty
	}

	// Don't start calculating until we have a complete window (including this block)
	if lenBlocks < s.windowBlocks-1 {
		return blockchain.GetLastBlock().NextDifficulty
	}

	// Only recalculate on the desired interval
	if lenBlocks < s.nextRecalculationBlock-1 {
		return blockchain.GetLastBlock().NextDifficulty
	}

	s.nextRecalculationBlock += s.intervalBlocks

	smaD, smaBT := s.sma(blockchain, thisBlockTime)

	return uint64(smaD * (float64(s.target) / smaBT))
}

// sma calculates the Simple Moving Averages for Difficulty and BlockTime
func (s *SMA) sma(blockchain blockchain.Blockchain, thisBlockTime uint) (smaD, smaBT float64) {
	var sumBT, sumD float64

	i := blockchain.GetLength()                      // Number of last block added (not including This block)
	j := blockchain.GetLength() - s.windowBlocks + 1 // Number of first block in window

	// Add values from the current block being processed
	sumBT += float64(thisBlockTime)
	sumD += float64(blockchain.GetBlock(i).NextDifficulty)

	// Add the rest of the values
	for ; i > j; i-- {
		sumBT += float64(blockchain.GetBlock(i).BlockTimeSeconds)
		sumD += float64(blockchain.GetBlock(i).ThisDifficulty)
	}
	smaBT = sumBT / float64(s.windowBlocks)
	smaD = sumD / float64(s.windowBlocks)

	return
}
