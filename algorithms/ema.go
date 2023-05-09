package algorithms

import (
	"fmt"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
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
func (e *EMA) NextDifficulty(blockchain blockchain.Blockchain) uint64 {
	i := blockchain.GetLength()
	if i < e.window {
		return blockchain.GetLastBlock().NextDifficulty
	}

	emaD, emaBT := ema(blockchain, e.window, e.lastBlockTimeEMA, e.lastDifficultyEMA)

	e.lastBlockTimeEMA = emaBT
	e.lastDifficultyEMA = emaD

	return uint64(emaD * float64(internal.Config.TargetBlockTimeMinutes) / emaBT)
}

// ema calculates the Exponential Moving Averages for Difficulty and BlockTime
// uses SMA as the first EMA
func ema(blockchain blockchain.Blockchain, window uint64, lastBlockTimeEMA, lastDifficultyEMA float64) (emaD, emaBT float64) {
	i := blockchain.GetLength()
	if i == window {
		return sma(blockchain, window)
	}

	j := i - window
	for i > j {
		i--
		emaBT = (float64(blockchain.GetLastBlock().BlockTimeSeconds)-lastBlockTimeEMA)*(2/(float64(window)+1)) + lastBlockTimeEMA
		emaD = (float64(blockchain.GetLastBlock().NextDifficulty)-lastDifficultyEMA)*(2/(float64(window)+1)) + lastDifficultyEMA
	}

	return
}
