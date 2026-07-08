package controller

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransferController struct {
	ethClient *ethclient.Client
}

func NewTransferController(ethClient *ethclient.Client) *TransferController {
	return &TransferController{
		ethClient: ethClient,
	}
}

type TransferRequest struct {
	ToAddress string `json:"to_address" binding:"required"`
	Amount    string `json:"amount"    binding:"required"`
}

func (t *TransferController) Transfer(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()

	// 从 .env 或环境变量中读取账户私钥（不通过网络传输，更安全）
	privateKeyHex := os.Getenv("ETHEREUM_PRIVATE_KEY")
	if privateKeyHex == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ETHEREUM_PRIVATE_KEY is not set in .env",
		})
		return
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ETHEREUM_PRIVATE_KEY in .env"})
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get public key"})
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := t.ethClient.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get nonce"})
		return
	}

	amount := new(big.Int)
	if _, ok := amount.SetString(req.Amount, 10); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	gasLimit := uint64(21000)
	gasPrice, err := t.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get gas price"})
		return
	}

	toAddress := common.HexToAddress(req.ToAddress)
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	chainId, err := t.ethClient.NetworkID(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get chain ID"})
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign transaction"})
		return
	}

	if err = t.ethClient.SendTransaction(ctx, signedTx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send transaction"})
		return
	}

	// ─── 控制台输出交易哈希 ───
	log.Printf("====== 交易已发送 ======")
	log.Printf("  交易哈希: %s", signedTx.Hash().Hex())
	log.Printf("  发送方: %s", fromAddress.Hex())
	log.Printf("  接收方: %s", toAddress.Hex())
	log.Printf("  金额: %s wei", amount.String())
	log.Printf("=========================")

	c.JSON(http.StatusOK, gin.H{
		"txHash": signedTx.Hash().Hex(),
		"from":   fromAddress.Hex(),
		"to":     toAddress.Hex(),
	})
}
