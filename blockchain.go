package main

import "time"

type Block struct {
	data         map[string]any
	hash         string
	previoushash string
	timestamp    time.Time
	pow          int
}

type BlockChain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}
