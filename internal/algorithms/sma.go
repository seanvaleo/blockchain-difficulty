package algorithms

import (
	"fmt"

	"github.com/seanvaleo/dsim/internal/config"
	"github.com/seanvaleo/dsim/pkg/dsim"
)

// SMA implements a Simple Moving Average equation, using the average
// block time of the most recent X blocks to estimate a more suitable
// difficulty
type SMA struct {
	name   string
	window uint64
}

// NewSMA instantiates and returns a new SMA
func NewSMA(window uint64) *SMA {
	return &SMA{
		name:   "SMA-" + fmt.Sprint(window),
		window: window,
	}
}

// Name returns the algorithm name
func (s *SMA) Name() string {
	return s.name
}

// Window returns the algorithm window
func (s *SMA) Window() uint64 {
	return s.window
}

// NextDifficulty calculates the next difficulty
func (s *SMA) NextDifficulty(chain []*dsim.Block) uint64 {
	i := uint64(len(chain))
	if i < s.window {
		return chain[i-1].Difficulty
	}

	smaD, smaBT := sma(chain, s.window)

	return uint64(smaD * float64(config.Cfg.TargetBlockTime) / smaBT)
}

// sma calculates the Simple Moving Averages for Difficulty and BlockTime
func sma(chain []*dsim.Block, window uint64) (smaD, smaBT float64) {
	var sumBT, sumD float64

	i := uint64(len(chain))
	j := i - window

	for i > j {
		i--
		sumBT += chain[i].BlockTime
		sumD += float64(chain[i].Difficulty)
	}
	smaBT = sumBT / float64(window)
	smaD = sumD / float64(window)

	return
}
