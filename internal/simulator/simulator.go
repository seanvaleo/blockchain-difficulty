package simulator

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"

	"github.com/seanvaleo/dsim/internal/config"
	"github.com/seanvaleo/dsim/pkg/dsim"
)

// Run executes the simulation
func Run(b dsim.Blockchain) func() error {
	return func() error {
		consistentHashPower(b)
		printResults(b)

		return nil
	}
}

func consistentHashPower(b dsim.Blockchain) {
	for i := uint64(0); i < config.Cfg.Blocks; i++ {
		b.AddBlock(config.Cfg.MinerHashTH * config.Cfg.StartMinerCount)
	}
}

// printResults prints the results to std output in a tabulated format
func printResults(b dsim.Blockchain) {
	sd, mean := statistics(b)

	w := tabwriter.NewWriter(os.Stdout, 20, 2, 1, ' ', 0)
	fmt.Fprintln(w, b.Name(), "\t", b.Algorithm().Name(), "\tSD:", sd, "\tMean:", mean, "\t")
	w.Flush()
}

// statistics generates standard deviation and mean values for the block interval time
// Skips the blocks that preceded the difficulty adjustment
func statistics(b dsim.Blockchain) (sd, mean float64) {
	i := b.Algorithm().Window()
	j := b.Algorithm().Window()

	var sum, count float64
	for i < b.Length() {
		fmt.Println(b.GetBlock(i))
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
