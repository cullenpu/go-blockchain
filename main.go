package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Block struct {
	Index     int
	Timestamp time.Time
	Hash      string
	PrevHash  string
	Data      string
	Pow       int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

var BC Blockchain // In-memory blockchain

func calculateHash(block Block) string {
	blockData := strconv.Itoa(block.Index) + block.Timestamp.String() + block.Data + block.PrevHash + strconv.Itoa(block.Pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Pow++
		b.Hash = calculateHash(*b)
	}
}

// Mine a new block and add it to the chain
func (bc *Blockchain) addBlock(Data string) Block {
	lastBlock := bc.chain[len(bc.chain)-1]
	newBlock := Block{
		Index:     lastBlock.Index + 1,
		Timestamp: time.Now(),
		Data:      Data,
		PrevHash:  lastBlock.Hash,
		Pow:       0,
	}
	newBlock.mine(bc.difficulty)
	bc.chain = append(bc.chain, newBlock)
	return newBlock
}

// Check validity of a block and its previous block
func isBlockValid(currBlock, prevBlock Block) bool {
	if prevBlock.Index+1 != currBlock.Index {
		return false
	}
	if currBlock.Hash != calculateHash(currBlock) {
		return false
	}
	if prevBlock.Hash != currBlock.PrevHash {
		return false
	}

	return true
}

// Check validity of entire chain
func (bc *Blockchain) isValid() bool {
	for i := 1; i < len(bc.chain); i++ {
		currBlock, prevBlock := bc.chain[i], bc.chain[i-1]
		if !isBlockValid(currBlock, prevBlock) {
			return false
		}
	}
	return true
}

// Initialize a new blockchain
func createBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		Hash:      "0",
		Timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func main() {
	BC = createBlockchain(2)

	router := gin.Default()

	router.GET("/", getBlockchainHandler)
	router.GET("/:index", getBlockHandler)
	router.POST("/mine", mineBlockHahdler)

	port := ":8080"
	router.Run(port)
}
