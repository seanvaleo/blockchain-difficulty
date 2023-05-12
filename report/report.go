package report

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/mesosoftware/blockchain-difficulty/blockchain"
	"github.com/mesosoftware/blockchain-difficulty/network"
)

// PrintResults prints the results to std output in a tabulated format
func PrintResults(networks []network.Network) {
	page := components.NewPage()
	page.Initialization.PageTitle = "DAA Simulation Results"

	// Define global upper/lower limits for charts
	var minY uint = math.MaxUint
	var maxY uint = 0

	// Add data to charts
	lineCharts := make([]*charts.Line, len(networks))
	for i, n := range networks {
		lineCharts[i] = charts.NewLine()
		lineChartAddData(lineCharts[i], n.Blockchain, &minY, &maxY)
	}

	for i, n := range networks {
		sd, mean := statistics(n)
		lastBlock := n.Blockchain.GetLastBlock()
		lastBlockTime := lastBlock.BlockTimeSeconds
		firstBlock := n.Blockchain.GetFirstBlock()
		firstBlockTime := firstBlock.BlockTimeSeconds
		blocksMined := n.Blockchain.GetLength()
		results := fmt.Sprintf("StdDev:%.4f  "+
			"Mean Block Time:%.4fs  "+
			"First Block Time:%vs  "+
			"Last Block Time:%vs  "+
			"Blocks Mined:%v",
			sd, mean, firstBlockTime, lastBlockTime, blocksMined)

		// Print results to CLI
		fmt.Println(n.Algorithm.Name())
		fmt.Println(results)
		fmt.Println("-----")

		// Add results to graphical charts page
		lineChartSetOpts(lineCharts[i], n.Algorithm.Name(), results, minY, maxY)
		page.AddCharts(lineCharts[i])
	}

	// Create html doc
	f, _ := os.Create("line.html")
	page.Render(f)

	// Open the HTML file with the default web browser
	absFilePath, _ := filepath.Abs(f.Name())
	cmd := exec.Command("open", absFilePath)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

// statistics generates standard deviation and mean values for the block interval time
func statistics(n network.Network) (sd, mean float64) {
	var sum, count float64

	for _, block := range n.Blockchain.Chain {
		sum += float64(block.BlockTimeSeconds)
		count++
	}
	mean = sum / count

	for _, block := range n.Blockchain.Chain {
		sd += math.Pow(float64(block.BlockTimeSeconds)-mean, 2)
	}
	sd = math.Sqrt(sd / count)

	return sd, mean
}

func lineChartAddData(lineChart *charts.Line, blockchain blockchain.Blockchain, minY, maxY *uint) {
	// Extract block time values from blockchain
	x := make([]int, blockchain.GetLength())
	y := make([]opts.LineData, blockchain.GetLength())
	for i, block := range blockchain.Chain {
		x[i] = i
		y[i] = opts.LineData{Value: block.BlockTimeSeconds}
		if block.BlockTimeSeconds > *maxY {
			*maxY = block.BlockTimeSeconds
		}
		if block.BlockTimeSeconds < *minY {
			*minY = block.BlockTimeSeconds
		}
	}

	// Add data to the chart instance
	lineChart.SetXAxis(x).
		AddSeries("Block Time (s)", y).
		SetSeriesOptions(
			charts.WithLineChartOpts(
				opts.LineChart{
					Smooth: false,
				},
			),
		)

	return
}

func lineChartSetOpts(lineChart *charts.Line, name, results string, minY, maxY uint) {
	// Create a new Line Chart instance
	lineChart.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{Title: name, Subtitle: results},
		),
		charts.WithLegendOpts(
			opts.Legend{Show: false},
		),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1600px",
			Height: "600px",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Scale: true,
			Min:   minY,
			Max:   maxY,
		}),
	)
}
