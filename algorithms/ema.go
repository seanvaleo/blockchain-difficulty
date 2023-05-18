package algorithms

import (
	"fmt"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
)

// EMA implements an Exponential Moving Average equation, using an exponentially weighted average
// block time of the most recent X blocks to estimate a more suitable
// difficulty
type EMA struct {
	name                   string
	target                 uint
	intervalBlocks         int // Frequency of difficulty re-calculation
	windowBlocks           int // Block data in sample
	nextRecalculationBlock int
	lastEmaBT              float64
	lastEmaD               float64
}

// NewEMA instantiates and returns a new EMA
func NewEMA(target uint, intervalBlocks, windowBlocks int) *EMA {
	return &EMA{
		name: fmt.Sprintf("EMA: Recalculate at every %v blocks using a %v block window. Target is %ds",
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
func (e *EMA) Name() string {
	return e.name
}

// NextDifficulty calculates the next difficulty
// We must account for the block time of the current block not yet added
func (e *EMA) NextDifficulty(blockchain blockchain.Blockchain, thisBlockTime uint) uint64 {
	lenBlocks := blockchain.GetLength()

	if lenBlocks == 0 {
		return blockchain.StartDifficulty
	}

	// Don't start calculating until we have a complete window (including this block)
	if lenBlocks < e.windowBlocks-1 {
		return blockchain.GetLastBlock().NextDifficulty
	}

	// Only recalculate on the desired interval
	if lenBlocks < e.nextRecalculationBlock-1 {
		return blockchain.GetLastBlock().NextDifficulty
	}

	e.nextRecalculationBlock += e.intervalBlocks

	emaD, emaBT := e.ema(blockchain, thisBlockTime)

	e.lastEmaD = emaD
	e.lastEmaBT = emaBT

	return uint64(emaD * (float64(e.target) / emaBT))
}

// ema calculates the Exponential Moving Averages for Difficulty and BlockTime
func (e *EMA) ema(blockchain blockchain.Blockchain, thisBlockTime uint) (emaD, emaBT float64) {
	var sumBT, sumD float64

	// For the first EMA calculation, use the SMA
	if e.lastEmaD == 0 {
		s := NewSMA(e.target, e.windowBlocks, e.intervalBlocks)
		return s.sma(blockchain, thisBlockTime)
	}

	// Average the values over the last interval (i.e. 1 day)
	// instead of just using the last block's data
	i := blockchain.GetLength()                        // Number of last block added (not including This block)
	j := blockchain.GetLength() - e.intervalBlocks + 1 // Number of first block in interval

	// Add values from the current block being processed
	sumBT += float64(thisBlockTime)
	sumD += float64(blockchain.GetBlock(i).NextDifficulty)

	// Add the rest of the values
	for ; i > j; i-- {
		sumBT += float64(blockchain.GetBlock(i).BlockTimeSeconds)
		sumD += float64(blockchain.GetBlock(i).ThisDifficulty)
	}

	avgBT := sumBT / float64(e.intervalBlocks)
	avgD := sumD / float64(e.intervalBlocks)

	// Standard EMA multiplier
	multiplier := float64(2) / (float64(e.windowBlocks) + 1)

	// Standard EMA formula
	emaBT = ((avgBT - e.lastEmaBT) * multiplier) + e.lastEmaBT
	emaD = ((avgD - e.lastEmaD) * multiplier) + e.lastEmaD

	return
}
