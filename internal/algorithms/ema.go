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
	name              string
	window            uint64
	lastDifficultyEMA float64
	lastBlockTimeEMA  float64
}

// NewEMA instantiates and returns a new EMA
func NewEMA(window uint64) *EMA {
	return &EMA{
		name:   "EMA-" + fmt.Sprint(window),
		window: window,
	}
}

// Name returns the algorithm name
func (e *EMA) Name() string {
	return e.name
}

// Window returns the algorithm window
func (e *EMA) Window() uint64 {
	return e.window
}

// NextDifficulty calculates the next difficulty
func (e *EMA) NextDifficulty(chain []*dsim.Block) uint64 {
	i := uint64(len(chain))
	if i < e.window {
		return chain[i-1].Difficulty
	}

	emaD, emaBT := ema(chain, e.window, e.lastBlockTimeEMA, e.lastDifficultyEMA)

	e.lastBlockTimeEMA = emaBT
	e.lastDifficultyEMA = emaD

	return uint64(emaD * float64(config.Cfg.TargetBlockTime) / emaBT)
}

// ema calculates the Exponential Moving Averages for Difficulty and BlockTime
// uses SMA as the first EMA
func ema(chain []*dsim.Block, window uint64, lastBlockTimeEMA, lastDifficultyEMA float64) (emaD, emaBT float64) {
	i := uint64(len(chain))
	if i == window {
		return sma(chain, window)
	}

	j := i - window
	for i > j {
		i--
		emaBT = (chain[i].BlockTime-lastBlockTimeEMA)*(2/(float64(window)+1)) + lastBlockTimeEMA
		emaD = (float64(chain[i].Difficulty)-lastDifficultyEMA)*(2/(float64(window)+1)) + lastDifficultyEMA
	}

	return
}
