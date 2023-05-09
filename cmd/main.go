package main

import (
	"context"

	"github.com/mesosoftware/blockchain-difficulty/algorithms"
	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/internal"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.Info("Running Difficulty Simulator v0.1")

	internal.InitConfig()
	internal.PrintConfig()

	log.Info("Please wait for results...")

	ctx := context.Context(context.Background())
	g, _ := errgroup.WithContext(ctx)

	g.Go(blockchain.Mine(blockchain.New("Blockchain 1", algorithms.NewSMA(10))))
	g.Go(blockchain.Mine(blockchain.New("Blockchain 2", algorithms.NewSMA(20))))
	g.Go(blockchain.Mine(blockchain.New("Blockchain 3", algorithms.NewSMA(50))))
	g.Go(blockchain.Mine(blockchain.New("Blockchain 4", algorithms.NewSMA(100))))
	g.Go(blockchain.Mine(blockchain.New("Blockchain 5", algorithms.NewEMA(10))))
	g.Go(blockchain.Mine(blockchain.New("Blockchain 6", algorithms.NewEMA(20))))
	g.Go(blockchain.Mine(blockchain.New("Blockchain 7", algorithms.NewEMA(50))))
	g.Go(blockchain.Mine(blockchain.New("Blockchain 8", algorithms.NewEMA(100))))

	if err := g.Wait(); err != nil {
		log.Error(err)
		return
	}

	// TODO
	// for all
	// report.PrintResults(b)

	log.Info("All results generated")
}
