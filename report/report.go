package report

import (
	"fmt"
	"math"

	"github.com/guptarohit/asciigraph"
	"github.com/mesosoftware/blockchain-difficulty/network"
	log "github.com/sirupsen/logrus"
)

// PrintResults prints the results to std output in a tabulated format
func PrintResults(n network.Network) {
	chartWidth := 120
	chartHeight := 10
	sd, mean := statistics(n)
	lastBlock := n.Blockchain.GetLastBlock()
	lastBlockTime := lastBlock.BlockTimeSeconds
	firstBlock := n.Blockchain.GetFirstBlock()
	firstBlockTime := firstBlock.BlockTimeSeconds
	blocksMined := n.Blockchain.GetLength()

	fmt.Println(n.Algorithm.Name())
	fmt.Printf("StdDev:%v  "+
		"Mean Block Time:%vs  "+
		"First Block Time:%vs  "+
		"Last Block Time:%vs  "+
		"Blocks Mined:%v"+
		"\n",
		sd, mean, firstBlockTime, lastBlockTime, blocksMined)

	blockTimes := make([]float64, n.Blockchain.GetLength())
	for i, block := range n.Blockchain.Chain {
		blockTimes[i] = float64(block.BlockTimeSeconds)
	}

	if len(blockTimes) < chartWidth {
		chartWidth = 0
		log.Warn("Chart width reduced due to (original chart width > block count). This is to avoid resizing which can results in steps")
	}

	// Be aware, when setting a width higher than n blocks, the graphing library will
	// render a diagonal line in steps which can be misinterpreted as data points
	blockTimesGraph := asciigraph.Plot(blockTimes,
		asciigraph.SeriesColors(asciigraph.Blue),
		asciigraph.Width(chartWidth),
		asciigraph.Height(chartHeight))

	fmt.Println(blockTimesGraph)
}

// statistics generates standard deviation and mean values for the block interval time
func statistics(n network.Network) (sd, mean float64) {
	var sum, count float64

	for _, block := range n.Blockchain.Chain {
		sum += float64(block.BlockTimeSeconds)
		count++
	}
	mean = sum / count

	for _, block := range n.Blockchain.Chain {
		sd += math.Pow(float64(block.BlockTimeSeconds)-mean, 2)
	}
	sd = math.Sqrt(sd / count)

	return sd, mean
}
