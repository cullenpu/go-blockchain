package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Data string `json:"Data" binding:"required"`
}

func getBlockchainHandler(c *gin.Context) {
	if !BC.isValid() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid blockchain"})
		return
	}

	c.JSON(http.StatusOK, BC.chain)
}

func getBlockHandler(c *gin.Context) {
	index, err := strconv.Atoi(c.Param("index"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if index < 0 || index >= len(BC.chain) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "index out of range"})
		return
	}

	block := BC.chain[index]
	c.JSON(http.StatusOK, block)
}

func mineBlockHahdler(c *gin.Context) {
	var req Request

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	block := BC.addBlock(req.Data)
	c.JSON(http.StatusOK, block)
}
