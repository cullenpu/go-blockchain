package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	index     int
	timestamp time.Time
	hash      string
	prevHash  string
	data      string
	pow       int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func calculateHash(block Block) string {
	blockData := strconv.Itoa(block.index) + block.timestamp.String() + block.data + block.prevHash + strconv.Itoa(block.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.index+1 != newBlock.index {
		return false
	}
	if newBlock.hash != calculateHash(newBlock) {
		return false
	}
	if oldBlock.hash != newBlock.prevHash {
		return false
	}

	return true
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = calculateHash(*b)
	}
}

func (bc *Blockchain) getLastBlock() Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) addBlock(data string) {
	lastBlock := bc.getLastBlock()
	newBlock := Block{
		index:     lastBlock.index + 1,
		timestamp: time.Now(),
		data:      data,
		prevHash:  lastBlock.hash,
	}
	newBlock.mine(bc.difficulty)
	bc.chain = append(bc.chain, newBlock)
}

func (bc *Blockchain) isValid() bool {
	for i := range bc.chain[1:] {
		currBlock, prevBlock := bc.chain[i], bc.chain[i-1]
		if !isBlockValid(currBlock, prevBlock) {
			return false
		}
	}
	return true
}

func createBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}
