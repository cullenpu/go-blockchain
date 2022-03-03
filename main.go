package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	index     int
	timestamp string
	hash      string
	prevHash  string
	data      string
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := strconv.Itoa(block.index) + block.timestamp + block.data + block.prevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, data string) (Block, error) {
	var newBlock Block
	t := time.Now()

	newBlock.index = oldBlock.index + 1
	newBlock.timestamp = t.String()
	newBlock.data = data
	newBlock.prevHash = oldBlock.hash
	newBlock.hash = calculateHash(newBlock)

	return newBlock, nil
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

func replaceChain(newBlockchain []Block) {
	if len(newBlockchain) > len(Blockchain) {
		Blockchain = newBlockchain
	}
}
