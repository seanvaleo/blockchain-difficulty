package simulator

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/seanvaleo/dsim/internal/config"
	"github.com/seanvaleo/dsim/pkg/dsim"
)

func Run(b dsim.Blockchain) func() error {
	return func() error {
		for i := uint(0); i < config.Cfg.Blocks; i++ {
			b.AddBlock(60)
		}

		printResults(b)

		return nil
	}
}

func printResults(b dsim.Blockchain) {

	sd, mean := b.Statistics()

	w := tabwriter.NewWriter(os.Stdout, 20, 2, 1, ' ', 0)
	fmt.Fprintln(w, b.Name(), "\t", b.AlgorithmName(), "\tSD:", sd, "\tMean:", mean, "\t")
	w.Flush()
}
