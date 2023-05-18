package algorithms

import (
	"testing"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
)

func TestETHNextDifficulty(t *testing.T) {
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
			name: "Perfect difficulty",
			blockchain: blockchain.Blockchain{
				Chain: []*blockchain.Block{
					{ThisDifficulty: 15000000, NextDifficulty: 15000000, BlockTimeSeconds: 15},
				},
			},
			thisBlockTime: 15,
			expected:      15000000,
		},
		{
			name: "Low difficulty in acceptable range",
			blockchain: blockchain.Blockchain{
				Chain: []*blockchain.Block{
					{ThisDifficulty: 1000000, NextDifficulty: 1000000, BlockTimeSeconds: 10},
				},
			},
			thisBlockTime: 10,
			expected:      1000163, // TODO should not change
		},
		{
			name: "Low difficulty out of acceptable range",
			blockchain: blockchain.Blockchain{
				Chain: []*blockchain.Block{
					{ThisDifficulty: 9000000, NextDifficulty: 9000000, BlockTimeSeconds: 9},
				},
			},
			thisBlockTime: 9,
			expected:      9000000, // TODO should change
		},
	}

	e := NewETH()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			nextDifficulty := e.NextDifficulty(tc.blockchain, tc.thisBlockTime)
			if nextDifficulty != tc.expected {
				t.Errorf("NextDifficulty: %v, Expected: %v", nextDifficulty, tc.expected)
			}
		})
	}
}
