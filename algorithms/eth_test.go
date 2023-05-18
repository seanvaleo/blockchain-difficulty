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
					{ThisDifficulty: 6000000000, NextDifficulty: 6000000000, BlockTimeSeconds: 600},
				},
			},
			thisBlockTime: 600,
			expected:      6000000000,
		},
		{
			name: "Low difficulty",
			blockchain: blockchain.Blockchain{
				Chain: []*blockchain.Block{
					{ThisDifficulty: 5500000000, NextDifficulty: 5500000000, BlockTimeSeconds: 550},
				},
			},
			thisBlockTime: 550,
			expected:      5500223795,
		},
		{
			name: "High difficulty",
			blockchain: blockchain.Blockchain{
				Chain: []*blockchain.Block{
					{ThisDifficulty: 6500000000, NextDifficulty: 6500000000, BlockTimeSeconds: 650},
				},
			},
			thisBlockTime: 650,
			expected:      6500000000,
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
