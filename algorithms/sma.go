package algorithms

import (
	"fmt"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
)

// SMA implements a Simple Moving Average equation, using the average
// block time of the most recent X blocks to estimate a more suitable
// difficulty
type SMA struct {
	name                   string
	windowBlocks           int // Block data in sample
	intervalBlocks         int // Frequency of difficulty re-calculation
	nextRecalculationBlock int
}

// NewSMA instantiates and returns a new SMA
func NewSMA(windowBlocks, intervalBlocks int) *SMA {
	return &SMA{
		name:                   fmt.Sprintf("SMA-Window-%v-Interval-%v", windowBlocks, intervalBlocks),
		windowBlocks:           windowBlocks,
		intervalBlocks:         intervalBlocks,
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
		fmt.Println(1)
		return blockchain.StartDifficulty
	}

	// Don't start calculating until we have a complete window (including this block)
	if lenBlocks < s.windowBlocks-1 {
		fmt.Println(2)
		return blockchain.GetLastBlock().NextDifficulty
	}

	// Only recalculate on the desired interval
	if lenBlocks < s.nextRecalculationBlock-1 {
		fmt.Println(3)
		return blockchain.GetLastBlock().NextDifficulty
	}

	fmt.Println(4)
	s.nextRecalculationBlock += s.intervalBlocks

	smaD, smaBT := sma(blockchain, s.windowBlocks, thisBlockTime)

	// For example:
	// smaD = 1,000,000,000 ; smaBT = 1,000 ; targetBT = 6000
	// new difficulty should become 6,000,000,000
	fmt.Println("smaD: ", smaD)
	fmt.Println("smaBT: ", smaBT)
	fmt.Println("new difficulty:", uint64(smaD*(float64(internal.Config.TargetBlockTimeSeconds)/smaBT)))
	return uint64(smaD * (float64(internal.Config.TargetBlockTimeSeconds) / smaBT))
}

// sma calculates the Simple Moving Averages for Difficulty and BlockTime
func sma(blockchain blockchain.Blockchain, windowBlocks int, thisBlockTime uint) (smaD, smaBT float64) {
	var sumBT, sumD float64

	i := blockchain.GetLength()                    // Number of last block added (not including This block)
	j := blockchain.GetLength() - windowBlocks + 1 // Number of first block in window

	// Add values from the current block being processed
	sumBT += float64(thisBlockTime)
	sumD += float64(blockchain.GetBlock(i).NextDifficulty)

	// Add the rest of the values
	for ; i > j; i-- {
		sumBT += float64(blockchain.GetBlock(i).BlockTimeSeconds)
		sumD += float64(blockchain.GetBlock(i).ThisDifficulty)
	}
	smaBT = sumBT / float64(windowBlocks)
	smaD = sumD / float64(windowBlocks)

	return
}
