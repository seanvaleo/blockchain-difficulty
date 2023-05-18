package main

import (
	"context"

	"github.com/mesosoftware/blockchain-difficulty/algorithms"
	"github.com/mesosoftware/blockchain-difficulty/internal"
	"github.com/mesosoftware/blockchain-difficulty/network"
	"github.com/mesosoftware/blockchain-difficulty/report"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.Info("Starting Blockchain Difficulty Simulator")

	internal.InitConfig()
	internal.PrintConfig()

	log.Info("Please wait for results...")

	networks := []network.Network{
		network.NewNetwork(600000000, algorithms.NewBTC()),
		network.NewNetwork(15000000, algorithms.NewETH()),
		network.NewNetwork(100000000, algorithms.NewSMA(100, 1440, 10080)),
		network.NewNetwork(100000000, algorithms.NewEMA(100, 1440, 10080)),
		network.NewNetwork(100000000, algorithms.NewLWMA(100, 1440, 10080)),
	}

	ctx := context.Context(context.Background())
	g, _ := errgroup.WithContext(ctx)

	for i := range networks {
		g.Go(networks[i].MiningSimulation())
	}

	if err := g.Wait(); err != nil {
		log.Error(err)
		return
	}

	report.PrintResults(networks)

	log.Info("All results generated")
}
