package algorithms

import (
	"testing"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
)

func TestNextDifficulty(t *testing.T) {
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
				StartDifficulty: 1000000,
				Chain: []*blockchain.Block{
					{ThisDifficulty: 2000000, NextDifficulty: 3000000, BlockTimeSeconds: 60},
					{ThisDifficulty: 3000000, NextDifficulty: 4000000, BlockTimeSeconds: 60},
					{ThisDifficulty: 4000000, NextDifficulty: 5000000, BlockTimeSeconds: 60},
					{ThisDifficulty: 5000000, NextDifficulty: 6000000, BlockTimeSeconds: 60},
					{ThisDifficulty: 6000000, NextDifficulty: 7000000, BlockTimeSeconds: 60},
					{ThisDifficulty: 7000000, NextDifficulty: 8000000, BlockTimeSeconds: 60},
				},
			},
			thisBlockTime: 60,
			expected:      36000000,
		},
	}

	sma := NewSMA(5, 5)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Calculate the next difficulty
			nextDifficulty := sma.NextDifficulty(tc.blockchain, tc.thisBlockTime)

			// Check if the result matches the expected value
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
				StartDifficulty: 1000000000,
				Chain: []*blockchain.Block{
					{ThisDifficulty: 1000000000, NextDifficulty: 1000000000, BlockTimeSeconds: 200},
					{ThisDifficulty: 1200000000, NextDifficulty: 1200000000, BlockTimeSeconds: 250},
					{ThisDifficulty: 1500000000, NextDifficulty: 1500000000, BlockTimeSeconds: 280},
					{ThisDifficulty: 1800000000, NextDifficulty: 1800000000, BlockTimeSeconds: 300},
					{ThisDifficulty: 1800000000, NextDifficulty: 1800000000, BlockTimeSeconds: 300},
				},
			},
			thisBlockTime: 350,
			expectedSmaD:  1625000000,
			expectedSmaBT: 282.5,
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
