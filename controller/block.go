package controller

import (
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type BlockController struct {
	ethClient *ethclient.Client
}

func NewBlockController(ethClient *ethclient.Client) *BlockController {
	return &BlockController{
		ethClient: ethClient,
	}
}

func (bc *BlockController) GetBlock(c *gin.Context) {
	blockNumberStr := c.Param("number")
	blockNumber := new(big.Int)
	blockNumber.SetString(blockNumberStr, 10)

	block, err := bc.ethClient.BlockByNumber(c, blockNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ─── 控制台输出（作业要求：输出查询结果到控制台）───
	log.Printf("====== 查询到区块信息 ======")
	log.Printf("  区块号: %s", block.Number().String())
	log.Printf("  区块哈希: %s", block.Hash().Hex())
	log.Printf("  交易数量: %d", len(block.Transactions()))
	log.Printf("  时间戳: %d", block.Time())
	log.Printf("============================")

	c.JSON(http.StatusOK, gin.H{
		"number":           block.Number().String(),
		"hash":             block.Hash().Hex(),
		"transactionCount": len(block.Transactions()),
		"timestamp":        block.Time(),
	})
}
