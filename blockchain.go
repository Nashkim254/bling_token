package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	data         map[string]any
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type BlockChain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (b Block) CalculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.CalculateHash()
	}
}

func createBlockchain(difficulty int) BlockChain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return BlockChain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}
