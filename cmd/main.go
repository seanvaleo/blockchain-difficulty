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

	net1 := network.NewNetwork(600000000, algorithms.NewBTC())
	net2 := network.NewNetwork(600000000, algorithms.NewSMA(1440, 10080))
	net3 := network.NewNetwork(600000000, algorithms.NewEMA(1440, 10080))
	net4 := network.NewNetwork(600000000, algorithms.NewLWMA(1440, 10080))

	ctx := context.Context(context.Background())
	g, _ := errgroup.WithContext(ctx)

	g.Go(net1.MiningSimulation())
	g.Go(net2.MiningSimulation())
	g.Go(net3.MiningSimulation())
	g.Go(net4.MiningSimulation())

	if err := g.Wait(); err != nil {
		log.Error(err)
		return
	}

	report.PrintResults([]network.Network{net1, net2, net3, net4})

	log.Info("All results generated")
}
