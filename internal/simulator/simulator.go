package simulator

import (
	"github.com/seanvaleo/dsim/pkg/dsim"
	log "github.com/sirupsen/logrus"
)

func Run(b dsim.Blockchain) func() error {
	return func() error {
		for i := 0; i < 10; i++ {
			b.AddBlock()
		}

		Report(b)

		return nil
	}
}

func Report(b dsim.Blockchain) {
	// b.name
	// algo.name
	// s.d.
	// mean
	log.Info(b)
}
