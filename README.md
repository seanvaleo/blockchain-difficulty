# Difficulty Simulator

[![Build Status](https://github.com/seanvaleo/dsim/actions/workflows/go.yml/badge.svg)](https://github.com/seanvaleo/dsim/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/seanvaleo/dsim)](https://goreportcard.com/report/github.com/seanvaleo/dsim)

Simulate the effectiveness of various blockchain difficulty algorithms in terms of volatility and accuracy.


### Background

In a decentralized blockchain, where no individual controls the timing of block additions, there
must be a mechanism set in place in order to regulate the desired block frequency.

A so-called 'difficulty algorithm' can be used to adjust the difficulty of mining a new block,
based on an estimate of the network's total problem solving power. With all things being equal, the
block time should approach and maintain its target block time.

The goal of this project is to provide a simulator to report on the effectiveness of various
difficulty algorithms by observing the standard deviation, and mean values of all block intervals
after adding X blocks.


### Performance

```
Blockchain 1         SMA-10             SD: 18.598669175489476  Mean: 60.17482517482517  
Blockchain 2         SMA-20             SD: 13.174972509226185  Mean: 59.324675324675326  
Blockchain 3         SMA-50             SD: 13.718402628006942  Mean: 57.624375624375624  
Blockchain 4         SMA-100            SD: 18.510401708827676  Mean: 55.07992007992008  
Blockchain 5         EMA-10             SD: 18.598669175489476  Mean: 60.17482517482517  
Blockchain 6         EMA-20             SD: 13.174972509226185  Mean: 59.324675324675326 
Blockchain 7         EMA-50             SD: 13.718402628006942  Mean: 57.624375624375624  
Blockchain 8         EMA-100            SD: 18.510401708827676  Mean: 55.07992007992008  
```


### Installation

Install Go: https://golang.org/doc/install

Download sources and install: `go get github.com/seanvaleo/dsim`


### Usage

```
dsim
```


### Configuration

Configure environment variables in the `.env` file.

Default values:
```
TARGET_BLOCK_TIME=60
BLOCKS=1000
MINER_COUNT=100
MINER_HASH_TH=100
```
