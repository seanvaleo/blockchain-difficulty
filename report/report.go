package report

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"

	"github.com/mesosoftware/blockchain-difficulty/blockchain"
)

// PrintResults prints the results to std output in a tabulated format
func PrintResults(b blockchain.Blockchain) {
	sd, mean := statistics(b)

	w := tabwriter.NewWriter(os.Stdout, 28, 8, 0, ' ', 0)
	fmt.Fprintln(w, b.Name(), "\t", b.Algorithm().Name(), "\tSD:", sd, "\tMean:", mean, "\t")
	w.Flush()
}

// statistics generates standard deviation and mean values for the block interval time
// Skips the blocks that preceded the difficulty adjustment
func statistics(b blockchain.Blockchain) (sd, mean float64) {
	i := b.Algorithm().Window()
	j := b.Algorithm().Window()

	var sum, count float64
	for i < b.Length() {
		sum += float64(b.GetBlock(i).BlockTime)
		count++
		i++
	}
	mean = sum / count

	for j < b.Length() {
		sd += math.Pow(float64(b.GetBlock(j).BlockTime)-mean, 2)
		j++
	}
	sd = math.Sqrt(sd / count)

	return sd, mean
}
