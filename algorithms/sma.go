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
func (s *SMA) NextDifficulty(blockchain blockchain.Blockchain) uint64 {
	i := blockchain.GetLength()
	if i < s.window {
		return blockchain.GetLastBlock().NextDifficulty
	}

	smaD, smaBT := sma(blockchain, s.window)

	return uint64(smaD * float64(internal.Config.TargetBlockTimeMinutes) / smaBT)
}

// sma calculates the Simple Moving Averages for Difficulty and BlockTime
func sma(blockchain blockchain.Blockchain, window uint64) (smaD, smaBT float64) {
	var sumBT, sumD float64

	i := blockchain.GetLength()
	j := i - window

	for i > j {
		i--
		sumBT += float64(blockchain.GetLastBlock().BlockTimeSeconds)
		sumD += float64(blockchain.GetLastBlock().NextDifficulty)
	}
	smaBT = sumBT / float64(window)
	smaD = sumD / float64(window)

	return
}
