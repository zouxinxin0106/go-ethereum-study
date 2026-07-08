package main

import (
	"log"
	"math/big"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/common"

	"github.com/FloraZou/go-eth-study/controller"
	"github.com/FloraZou/go-eth-study/ethereum"
	"github.com/FloraZou/go-eth-study/service"
)

func main() {
	// 加载 .env 配置文件
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	rpcUrl := os.Getenv("ETHEREUM_RPC_URL")
	if rpcUrl == "" {
		log.Fatal("ETHEREUM_RPC_URL is not set in .env or environment variables")
	}

	privateKey := os.Getenv("ETHEREUM_PRIVATE_KEY")
	if privateKey == "" {
		log.Fatal("ETHEREUM_PRIVATE_KEY is not set in .env or environment variables")
	}

	ethClient, err := ethereum.NewClient(rpcUrl)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// --- Block 路由 ---
	blockController := controller.NewBlockController(ethClient)
	r.GET("/block/:number", blockController.GetBlock)

	// --- Transfer 路由 ---
	transferController := controller.NewTransferController(ethClient)
	r.POST("/transfer", transferController.Transfer)

	// --- Contract 路由（如果配置了合约地址）---
	contractAddrHex := os.Getenv("CONTRACT_ADDRESS")
	if contractAddrHex != "" && common.IsHexAddress(contractAddrHex) {
		chainID := new(big.Int)
		chainID.SetString(os.Getenv("CHAIN_ID"), 10)

		contractSrv, err := service.NewContractService(
			ethClient,
			common.HexToAddress(contractAddrHex),
			privateKey,
			chainID,
		)
		if err != nil {
			log.Fatalf("Failed to initialize contract service: %v", err)
		}

		contractCtrl := controller.NewContractController(contractSrv)
		r.GET("/contract/info", contractCtrl.GetContractInfo)
		r.GET("/contract/donation", contractCtrl.GetDonation)
		r.GET("/contract/top-donors", contractCtrl.GetTop3Donors)
		r.POST("/contract/donate", contractCtrl.Donate)
		r.POST("/contract/withdraw", contractCtrl.Withdraw)
	} else {
		log.Println("CONTRACT_ADDRESS not set — skipping contract routes")
	}

	r.Run()
}
