package simulator

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/seanvaleo/dsim/internal/config"
	"github.com/seanvaleo/dsim/pkg/dsim"
)

// Run executes the simulation
func Run(b dsim.Blockchain) func() error {
	return func() error {
		var difficulty, blockTime uint64

		for i := uint64(0); i < config.Cfg.Blocks; i++ {
			difficulty = b.Difficulty()
			blockTime = difficulty / (config.Cfg.MinerHashTH * config.Cfg.MinerCount)

			b.AddBlock(blockTime)
		}

		printResults(b)

		return nil
	}
}

// printResults prints the results to std output in a tabulated format
func printResults(b dsim.Blockchain) {
	sd, mean := b.Statistics()

	w := tabwriter.NewWriter(os.Stdout, 20, 2, 1, ' ', 0)
	fmt.Fprintln(w, b.Name(), "\t", b.AlgorithmName(), "\tSD:", sd, "\tMean:", mean, "\t")
	w.Flush()
}
