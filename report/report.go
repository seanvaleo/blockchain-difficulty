package report

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"

	"github.com/mesosoftware/blockchain-difficulty/network"
)

// TODO type Results?

// PrintResults prints the results to std output in a tabulated format
func PrintResults(n network.Network) {
	sd, mean := statistics(n)

	w := tabwriter.NewWriter(os.Stdout, 28, 8, 0, ' ', 0)
	fmt.Fprintln(w, n.Algorithm.Name(), "\tSD:", sd, "\tMean:", mean, "\t")
	w.Flush()
}

// statistics generates standard deviation and mean values for the block interval time
// Skips the blocks that preceded the difficulty adjustment
func statistics(n network.Network) (sd, mean float64) {
	i := n.Algorithm.Window()
	j := n.Algorithm.Window()

	var sum, count float64
	for i < n.Blockchain.GetLength() {
		sum += float64(n.Blockchain.GetLastBlock().BlockTime)
		count++
		i++
	}
	mean = sum / count

	for j < n.Blockchain.GetLength() {
		sd += math.Pow(float64(n.Blockchain.GetLastBlock().BlockTime)-mean, 2)
		j++
	}
	sd = math.Sqrt(sd / count)

	return sd, mean
}
