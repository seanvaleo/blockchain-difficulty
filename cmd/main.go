package main

import (
	"context"

	"github.com/seanvaleo/dsim/internal/algorithms"
	"github.com/seanvaleo/dsim/internal/blockchain"
	"github.com/seanvaleo/dsim/internal/config"
	"github.com/seanvaleo/dsim/internal/simulator"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {

	log.Info("Running Difficulty Simulator v0.1")

	config.Init()
	config.Print()

	log.Info("Please wait for results...")

	ctx := context.Context(context.Background())
	g, _ := errgroup.WithContext(ctx)

	g.Go(simulator.Run(blockchain.New("Blockchain1", algorithms.NewExample())))

	if err := g.Wait(); err != nil {
		log.Error(err)
	}
}
