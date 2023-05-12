package algorithms

import (
	"testing"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
)

func TestEMANextDifficulty(t *testing.T) {
	internal.InitConfig()
	testCases := []struct {
		name          string
		blockchain    blockchain.Blockchain
		thisBlockTime uint
		expected      uint64
	}{
		{
			name:          "Empty blockchain",
			blockchain:    blockchain.New(1000000),
			thisBlockTime: 60,
			expected:      1000000,
		},
		{
			name: "Incomplete window",
			blockchain: blockchain.Blockchain{
				StartDifficulty: 1000000,
				Chain: []*blockchain.Block{
					{ThisDifficulty: 2000000, NextDifficulty: 3000000, BlockTimeSeconds: 60},
				},
			},
			thisBlockTime: 60,
			expected:      3000000,
		},
		{
			name: "Complete window",
			blockchain: blockchain.Blockchain{
				Chain: []*blockchain.Block{
					{ThisDifficulty: 1200000000, NextDifficulty: 1400000000, BlockTimeSeconds: 1200},
					{ThisDifficulty: 1400000000, NextDifficulty: 1600000000, BlockTimeSeconds: 1400},
					{ThisDifficulty: 1600000000, NextDifficulty: 1800000000, BlockTimeSeconds: 1600},
					{ThisDifficulty: 1800000000, NextDifficulty: 2000000000, BlockTimeSeconds: 1800},
				},
			},
			thisBlockTime: 2000,
			expected:      600000000,
		},
	}

	e := NewEMA(5, 5)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nextDifficulty := e.NextDifficulty(tc.blockchain, tc.thisBlockTime)
			if nextDifficulty != tc.expected {
				t.Errorf("NextDifficulty: %v, Expected: %v", nextDifficulty, tc.expected)
			}
		})
	}
}

func TestEma(t *testing.T) {
	testCases := []struct {
		name          string
		blockchain    blockchain.Blockchain
		lastEmaBT     float64
		lastEmaD      float64
		thisBlockTime uint
		expectedEmaBT float64
		expectedEmaD  float64
	}{
		{
			name: "Complete window - Increasing difficulty - No last EMA",
			blockchain: blockchain.Blockchain{
				Chain: []*blockchain.Block{
					{ThisDifficulty: 1200000000, NextDifficulty: 1400000000, BlockTimeSeconds: 1200},
					{ThisDifficulty: 1400000000, NextDifficulty: 1600000000, BlockTimeSeconds: 1400},
					{ThisDifficulty: 1600000000, NextDifficulty: 1800000000, BlockTimeSeconds: 1600},
					{ThisDifficulty: 1800000000, NextDifficulty: 2000000000, BlockTimeSeconds: 1800},
				},
			},
			lastEmaBT:     0,
			lastEmaD:      0,
			thisBlockTime: 2000,
			expectedEmaBT: 1600,
			expectedEmaD:  1600000000,
		},
		{
			name: "Complete window - Increasing difficulty - Existing last EMA",
			blockchain: blockchain.Blockchain{
				Chain: []*blockchain.Block{
					{ThisDifficulty: 1200000000, NextDifficulty: 1400000000, BlockTimeSeconds: 1200},
					{ThisDifficulty: 1400000000, NextDifficulty: 1600000000, BlockTimeSeconds: 1400},
					{ThisDifficulty: 1600000000, NextDifficulty: 1800000000, BlockTimeSeconds: 1600},
					{ThisDifficulty: 1800000000, NextDifficulty: 2000000000, BlockTimeSeconds: 1800},
				},
			},
			lastEmaBT:     1600,
			lastEmaD:      1600000000,
			thisBlockTime: 2000,
			expectedEmaBT: 1600,
			expectedEmaD:  1600000000,
		},
	}

	e := NewEMA(5, 5)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			emaD, emaBT := e.ema(tc.blockchain, tc.thisBlockTime)
			if emaD != tc.expectedEmaD {
				t.Errorf("Expected EMA difficulty: %f, but got: %f", tc.expectedEmaD, emaD)
			}
			if emaBT != tc.expectedEmaBT {
				t.Errorf("Expected EMA block time: %f, but got: %f", tc.expectedEmaBT, emaBT)
			}
		})
	}
}
