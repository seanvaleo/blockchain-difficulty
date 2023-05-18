package algorithms

import (
	"fmt"
	"math"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
)

// ETH implements the Ethereum difficulty adjustment algorithm, which is:
// New Difficulty = Old Difficulty + (Old Difficulty // 2048 *
//     max(1 - (block timestamp - parent timestamp) // 10, -99) +
//     int(2**((block number // 100000) - 2)))
// to estimate a more suitable difficulty
type ETH struct {
	name                   string
	target                 uint
	intervalBlocks         int // Frequency of difficulty re-calculation
	windowBlocks           int // Block data in sample
	nextRecalculationBlock int
}

// NewETH instantiates and returns a new ETH
func NewETH() *ETH {
	return &ETH{
		name:                   fmt.Sprintf("Ethereum: Recalculate at every 1 block using a 1 block window. Target is 15s"),
		target:                 15, // fixed
		intervalBlocks:         1,  // fixed
		windowBlocks:           1,  // fixed
		nextRecalculationBlock: 1,  // fixed
	}
}

// Name returns the algorithm name
func (e *ETH) Name() string {
	return e.name
}

// NextDifficulty calculates the next difficulty
// We must account for the block time of the current block not yet added
func (e *ETH) NextDifficulty(blockchain blockchain.Blockchain, thisBlockTime uint) uint64 {
	lenBlocks := blockchain.GetLength()

	if lenBlocks == 0 {
		return blockchain.StartDifficulty
	}

	// Don't start calculating until we have a complete window (including this block)
	if lenBlocks < e.windowBlocks-1 {
		return blockchain.GetLastBlock().NextDifficulty
	}

	// Only recalculate on the desired interval
	if lenBlocks < e.nextRecalculationBlock-1 {
		return blockchain.GetLastBlock().NextDifficulty
	}

	e.nextRecalculationBlock += e.intervalBlocks

	oldDifficulty := float64(blockchain.GetLastBlock().NextDifficulty)
	thisBlockNumber := float64(lenBlocks + 1)

	return uint64(oldDifficulty +
		float64((oldDifficulty/2048)*
			math.Max(-99, 1-
				(float64(thisBlockTime)/float64(e.target)))+
			math.Pow(2,
				float64((int(thisBlockNumber)/100000)-
					2))))
}
