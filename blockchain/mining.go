package blockchain

import "github.com/mesosoftware/blockchain-difficulty/internal"

// Mine runs the mining simulation
func Mine(b Blockchain) func() error {
	return func() error {
		consistentHashPower(b)

		return nil
	}
}

// consistentHashPower simulates adding blocks in a consistent network
func consistentHashPower(b Blockchain) {
	for i := uint64(0); i < internal.Config.Blocks; i++ {
		b.AddBlock(internal.Config.MinerHashTH * internal.Config.StartMinerCount)
	}
}
