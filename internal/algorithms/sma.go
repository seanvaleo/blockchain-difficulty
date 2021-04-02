package algorithms

import (
	"fmt"
	"math/big"
)

// SMA implements a Simple Moving Average equation, using the average
// block time of the most recent X blocks to estimate a more suitable
// difficulty
type SMA struct {
	name   string
	window uint
}

// NewSMA instantiates and returns a new SMA
func NewSMA(window uint) *SMA {
	return &SMA{
		name:   "SMA-" + fmt.Sprint(window),
		window: window,
	}
}

// Name returns the algorithm name
func (s *SMA) Name() string {
	return s.name
}

// NextDifficulty calculates the next difficulty
func (s *SMA) NextDifficulty() *big.Int {

	// Average time of last -s.window- blocks
	// TARGET_BLOCK_TIME

	return big.NewInt(1)
}
