package network

import (
	"math/rand"
	"time"

	"github.com/mesosoftware/blockchain-difficulty/algorithms"
	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
)

// For the sake of simplicity, all clients on the network using the same software
// are grouped into one "network" structure
// They use the same Algorithm
// They have a copy of the same Blockchain
// For simplicity, the hashPower represents collective network hashing power
type Network struct {
	Algorithm  algorithms.Algorithm
	Blockchain blockchain.Blockchain
	hashPower  uint64
}

func NewNetwork(startDifficulty uint64, algorithm algorithms.Algorithm) Network {
	return Network{
		Algorithm:  algorithm,
		Blockchain: blockchain.New(startDifficulty),
		hashPower:  internal.Config.InitialNetworkHashPower,
	}
}

// MiningSimulation simulates adding blocks in a varying power network
func (n *Network) MiningSimulation() func() error {
	timeElapsedDays := uint32(0)
	timeElapsedSeconds := uint64(0)

	return func() error {
		// Add blocks continuously until the simulation is complete
		for timeElapsedDays < internal.Config.SimulationDays {

			var curDifficulty uint64
			nextDifficulty := curDifficulty
			lastBlock := n.Blockchain.GetLastBlock()
			if lastBlock == nil {
				curDifficulty = n.Blockchain.StartDifficulty
			} else {
				curDifficulty = lastBlock.NextDifficulty

				// Give the algorithm a chance to modify the difficulty
				nextDifficulty = n.Algorithm.NextDifficulty(n.Blockchain)
			}

			blockTimeSeconds := uint(float64(curDifficulty) / float64(n.hashPower))

			timeElapsedSeconds += uint64(blockTimeSeconds)

			n.Blockchain.AddBlock(nextDifficulty, blockTimeSeconds)

			// If a day or more has elapsed, update days elapsed, and modify the network hashpower
			if timeElapsedSeconds >= uint64((timeElapsedDays+1)*24*60*60) {
				timeElapsedDays = uint32(timeElapsedSeconds / (24 * 60 * 60))

				// Modify hashpower by a random amount with limits of +-25%
				rand.Seed(time.Now().UnixNano())
				pctChange := (rand.Float64() * 0.5) - 0.25
				n.hashPower = uint64(float64(n.hashPower) * (1 + pctChange))
			}
		}
		return nil
	}
}
