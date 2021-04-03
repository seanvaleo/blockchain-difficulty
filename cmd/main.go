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

	g.Go(simulator.Run(blockchain.New("Blockchain 1", algorithms.NewSMA(10))))
	g.Go(simulator.Run(blockchain.New("Blockchain 2", algorithms.NewSMA(20))))
	g.Go(simulator.Run(blockchain.New("Blockchain 3", algorithms.NewSMA(50))))
	g.Go(simulator.Run(blockchain.New("Blockchain 4", algorithms.NewSMA(100))))
	g.Go(simulator.Run(blockchain.New("Blockchain 5", algorithms.NewEMA(10))))
	g.Go(simulator.Run(blockchain.New("Blockchain 6", algorithms.NewEMA(20))))
	g.Go(simulator.Run(blockchain.New("Blockchain 7", algorithms.NewEMA(50))))
	g.Go(simulator.Run(blockchain.New("Blockchain 8", algorithms.NewEMA(100))))

	if err := g.Wait(); err != nil {
		log.Error(err)
	} else {
		log.Info("All results generated")
	}
}
