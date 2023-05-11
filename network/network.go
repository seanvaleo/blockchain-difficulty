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

			var thisDifficulty uint64
			var nextDifficulty uint64

			if n.Blockchain.GetLength() == 0 {
				thisDifficulty = n.Blockchain.StartDifficulty
			} else {
				thisDifficulty = n.Blockchain.GetLastBlock().NextDifficulty
			}

			thisBlockTimeSeconds := uint(float64(thisDifficulty) / float64(n.hashPower))

			// Give the algorithm a chance to modify the difficulty
			// We need to hand it information on the current block time and difficulty
			nextDifficulty = n.Algorithm.NextDifficulty(n.Blockchain, thisBlockTimeSeconds)

			n.Blockchain.AddBlock(thisDifficulty, nextDifficulty, thisBlockTimeSeconds)

			timeElapsedSeconds += uint64(thisBlockTimeSeconds)
			// If a day or more has elapsed, update days elapsed, and modify the network hashpower
			if timeElapsedSeconds >= uint64((timeElapsedDays+1)*24*60*60) {
				timeElapsedDays = uint32(timeElapsedSeconds / (24 * 60 * 60))

				// Modify hashpower by a random amount within configured limits from initial hashpower
				rand.Seed(time.Now().UnixNano())
				max_change := internal.Config.LimitNetworkHashPowerPctChange
				pctChange := (rand.Float64() / (float64(max_change) / 2)) - float64(max_change)/100
				n.hashPower = uint64(float64(internal.Config.InitialNetworkHashPower) * (1 + pctChange))
			}
		}
		return nil
	}
}
