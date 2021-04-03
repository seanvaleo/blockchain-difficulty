package algorithms

import (
	"fmt"

	"github.com/seanvaleo/dsim/internal/config"
	"github.com/seanvaleo/dsim/pkg/dsim"
)

// EMA implements a Simple Moving Average equation, using the average
// block time of the most recent X blocks to estimate a more suitable
// difficulty
type EMA struct {
	name   string
	window uint64
}

// NewEMA instantiates and returns a new EMA
func NewEMA(window uint64) *EMA {
	return &EMA{
		name:   "EMA-" + fmt.Sprint(window),
		window: window,
	}
}

// Name returns the algorithm name
func (s *EMA) Name() string {
	return s.name
}

// NextDifficulty calculates the next difficulty
func (s *EMA) NextDifficulty(chain []*dsim.Block) uint64 {

	var sumBT, meanBT, sumD, meanD uint64

	// k := 2 / (float64(s.window) + 1)
	// fmt.Println(k)

	i := uint64(len(chain))
	if i < s.window {
		return chain[i-1].Difficulty
	}

	j := i - s.window

	for i > j {
		i--
		sumBT += chain[i].BlockTime
		sumD += chain[i].Difficulty
	}
	meanBT = sumBT / s.window
	meanD = sumD / s.window

	return (meanD * config.Cfg.TargetBlockTime) / meanBT
}
