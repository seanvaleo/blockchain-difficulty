package report

import (
	"fmt"
	"math"

	"github.com/guptarohit/asciigraph"
	"github.com/mesosoftware/blockchain-difficulty/network"
)

// TODO type Results?

// PrintResults prints the results to std output in a tabulated format
func PrintResults(n network.Network) {
	sd, mean := statistics(n)

	fmt.Println(n.Algorithm.Name())
	fmt.Printf("StdDev:%v\tMean Block Time:%vm\tFirst Block Time:%vm\tLast Block Time:%vm\tBlocks Mined:%v\n", sd, mean, 20, 30, 34000)

	// fmt.Println(n.Blockchain.Chain)
	/*
		blockTimes := make([]float64, n.Blockchain.GetLength())
		for i, block := range n.Blockchain.Chain {
			blockTimes[i] = float64(block.BlockTimeSeconds)
		}
		fmt.Println(blockTimes)
	*/

	blockTimes := []float64{3, 2, 6, 2, 6, 2, 5, 6, 7, 8, 2, 2, 6, 2, 6, 2, 5, 6, 7, 8, 2, 6, 2, 6, 2, 5, 6, 7, 8, 2}

	blockTimesGraph := asciigraph.Plot(blockTimes, asciigraph.SeriesColors(asciigraph.Blue))

	fmt.Println(blockTimesGraph)
}

// statistics generates standard deviation and mean values for the block interval time
// Skips the blocks that preceded the difficulty adjustment
func statistics(n network.Network) (sd, mean float64) {
	i := n.Algorithm.Window()
	j := n.Algorithm.Window()

	var sum, count float64
	for i < n.Blockchain.GetLength() {
		sum += float64(n.Blockchain.GetLastBlock().BlockTimeSeconds)
		count++
		i++
	}
	mean = sum / count

	for j < n.Blockchain.GetLength() {
		sd += math.Pow(float64(n.Blockchain.GetLastBlock().BlockTimeSeconds)-mean, 2)
		j++
	}
	sd = math.Sqrt(sd / count)

	return sd, mean
}
