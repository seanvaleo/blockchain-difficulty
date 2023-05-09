package network

import (
	"github.com/mesosoftware/blockchain-difficulty/algorithms"
	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
)

// For the sake of simplicity, all clients on the network using the same software
// are grouped into one "network" structure
// They use the same Algorithm
// They have a copy of the same Blockchain
// For simplicity, the hashPower represents not the hashPower but the total network hashing power
type Network struct {
	Algorithm  algorithms.Algorithm
	Blockchain blockchain.Blockchain
	hashPower  uint64
}

func NewNetwork(startHashPower uint64, algorithm algorithms.Algorithm) Network {
	return Network{
		Algorithm:  algorithm,
		hashPower:  startHashPower,
		Blockchain: blockchain.New(),
	}
}

// MiningSimulation runs the network mining simulation
func (n *Network) MiningSimulation() func() error {
	return func() error {
		n.variableHashPowerMining()

		return nil
	}
}

// variableHashPowerMining simulates adding blocks in a varying power network
func (n *Network) variableHashPowerMining() {
	for i := uint64(0); i < internal.Config.Blocks; i++ {

		lastBlock := n.Blockchain.GetBlock(n.Blockchain.GetLength())
		curDifficulty := lastBlock.NextDifficulty

		nextDifficulty := n.Algorithm.NextDifficulty(n.Blockchain)
		blockTime := float64(curDifficulty) / float64(n.hashPower)

		// TODO modify hashpower

		n.Blockchain.AddBlock(nextDifficulty, blockTime)
	}
}
