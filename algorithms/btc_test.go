package algorithms

import (
	"testing"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
)

func TestBTCNextDifficulty(t *testing.T) {
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
			expected:      750000000,
		},
	}

	b := NewBTC()
	b.windowBlocks = 5
	b.intervalBlocks = 5
	b.nextRecalculationBlock = 5

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nextDifficulty := b.NextDifficulty(tc.blockchain, tc.thisBlockTime)
			if nextDifficulty != tc.expected {
				t.Errorf("NextDifficulty: %v, Expected: %v", nextDifficulty, tc.expected)
			}
		})
	}
}

func TestSumBlockTimes(t *testing.T) {
	testCases := []struct {
		name          string
		blockchain    blockchain.Blockchain
		thisBlockTime uint
		expectedSum   uint
	}{
		{
			name: "5 Blocks",
			blockchain: blockchain.Blockchain{
				Chain: []*blockchain.Block{
					{BlockTimeSeconds: 1200},
					{BlockTimeSeconds: 1400},
					{BlockTimeSeconds: 1600},
					{BlockTimeSeconds: 1800},
				},
			},
			thisBlockTime: 2000,
			expectedSum:   8000,
		},
	}

	b := NewBTC()
	b.windowBlocks = 5

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sum := b.sumBlockTimes(tc.blockchain, tc.thisBlockTime)
			if sum != tc.expectedSum {
				t.Errorf("Expected Sum of block times: %d, but got: %d", tc.expectedSum, sum)
			}
		})
	}
}
