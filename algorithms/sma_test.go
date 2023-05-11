package algorithms

import (
	"testing"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
)

func TestNextDifficulty(t *testing.T) {
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

	sma := NewSMA(5, 5)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nextDifficulty := sma.NextDifficulty(tc.blockchain, tc.thisBlockTime)
			if nextDifficulty != tc.expected {
				t.Errorf("NextDifficulty: %v, Expected: %v", nextDifficulty, tc.expected)
			}
		})
	}
}

func TestSma(t *testing.T) {
	testCases := []struct {
		name          string
		blockchain    blockchain.Blockchain
		thisBlockTime uint
		expectedSmaD  float64
		expectedSmaBT float64
	}{
		{
			name: "Complete window - Increasing difficulty",
			blockchain: blockchain.Blockchain{
				Chain: []*blockchain.Block{
					{ThisDifficulty: 1200000000, NextDifficulty: 1400000000, BlockTimeSeconds: 1200},
					{ThisDifficulty: 1400000000, NextDifficulty: 1600000000, BlockTimeSeconds: 1400},
					{ThisDifficulty: 1600000000, NextDifficulty: 1800000000, BlockTimeSeconds: 1600},
					{ThisDifficulty: 1800000000, NextDifficulty: 2000000000, BlockTimeSeconds: 1800},
				},
			},
			thisBlockTime: 2000,
			expectedSmaD:  1600000000,
			expectedSmaBT: 1600,
		},
	}

	s := NewSMA(5, 5)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			smaD, smaBT := s.sma(tc.blockchain, tc.thisBlockTime)
			if smaD != tc.expectedSmaD {
				t.Errorf("Expected SMA difficulty: %f, but got: %f", tc.expectedSmaD, smaD)
			}
			if smaBT != tc.expectedSmaBT {
				t.Errorf("Expected SMA block time: %f, but got: %f", tc.expectedSmaBT, smaBT)
			}
		})
	}
}
