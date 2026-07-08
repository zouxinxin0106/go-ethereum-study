package controller

// ──────────────────────────────────────────────
// 合约 HTTP 控制器
//
// 这个文件的作用是：把合约的功能暴露成 HTTP 接口（API），
// 这样你就可以用浏览器或者 curl 命令来调用合约了。
//
// 打个比方：
//   service/contract_service.go 是"售货机按钮面板"
//   而这个文件就是"售货机外面的触摸屏"
//   你在触摸屏上点一下（发 HTTP 请求），
//   触摸屏就会去按对应的按钮（调 service 里的方法）。
//
// 每一层各司其职：
//   main.go           — 接线员，把所有东西连起来
//   controller/       — 服务员，接收你的请求、返回结果
//   service/          — 操作员，执行具体业务逻辑
//   contract/         — 翻译官（abigen 自动生成），把 Go 翻译成以太坊能懂的话
// ──────────────────────────────────────────────

import (
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"github.com/FloraZou/go-eth-study/service"
)

// ─── ContractController 结构体 ─────────────────
//
// ContractController 是"合约 HTTP 控制器"。
// 它里面装了一个 ContractService，所有请求来了都交给它处理。
//

type ContractController struct {
	contractSrv *service.ContractService // 底层合约服务，真正干活的人
}

// ─── 创建合约控制器 ──────────────────────────
//
// NewContractController 创建控制器。
// 参数 contractSrv 是一个已经初始化好的合约服务。
//

func NewContractController(contractSrv *service.ContractService) *ContractController {
	return &ContractController{
		contractSrv: contractSrv,
	}
}

// ─── 捐款请求体格式 ───────────────────────────
//
// DonateRequest 定义了 POST /contract/donate 接口的请求体格式。
// 前端发来的 JSON 必须包含 amount 字段，否则会报错（binding:"required"）。
//
// 请求示例：
//   {
//     "amount": "1000000000000000000"
//   }
//
// 注意 amount 的单位是 wei，1 ETH = 10^18 wei。
// 上面这个例子相当于捐了 1 个 ETH。
//

type DonateRequest struct {
	Amount string `json:"amount" binding:"required"` // 捐款金额，单位 wei，字符串形式（因为数字可能超大）
}

// ═══════════════════════════════════════════════
// 以下是 HTTP 接口（API），所有操作都通过 Gin 框架处理
// ═══════════════════════════════════════════════

// ─── GET /contract/donation?address=0x... ─────
//
// GetDonation 查询某个人捐了多少钱。
//
// 使用方式：
//   在浏览器地址栏输入：
//   http://localhost:8080/contract/donation?address=0x123...
//
// 或者用 curl 命令：
//   curl "http://localhost:8080/contract/donation?address=0x123..."
//
// 返回格式：
//   {
//     "donation": "1000000000000000000"
//   }
//
// 如果地址格式不对（不是合法的 0x 开头地址），会返回 400 错误。
// 如果查询过程出问题（比如网络断了），会返回 500 错误。
//

func (cc *ContractController) GetDonation(c *gin.Context) {
	// 从 URL 的查询参数里拿到 address
	// 比如 ?address=0x123... → addrHex = "0x123..."
	addrHex := c.Query("address")

	// 检查地址格式是否合法
	// 必须是以 0x 开头的 42 位十六进制地址
	if !common.IsHexAddress(addrHex) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address"})
		return
	}

	// 把字符串地址转成 Go 的 Address 类型，然后去链上查
	donation, err := cc.contractSrv.GetDonation(c, common.HexToAddress(addrHex))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 查到啦，把结果返回给调用者
	// donation.String() 把大数字转成字符串，防止精度丢失
	c.JSON(http.StatusOK, gin.H{"donation": donation.String()})
}

// ─── GET /contract/top-donors ─────────────────
//
// GetTop3Donors 返回捐款最多的前三名。
//
// 使用方式：
//   curl http://localhost:8080/contract/top-donors
//
// 返回格式：
//   {
//     "topDonors": [
//       "0xaaa...",
//       "0xbbb...",
//       "0xccc..."
//     ]
//   }
//
// 排名第一的（捐最多的）在数组第一个位置。
// 如果某个人没排上，他的位置就是 0x000...000（空地址）。
//

func (cc *ContractController) GetTop3Donors(c *gin.Context) {
	// 去合约里查前三名
	donors, err := cc.contractSrv.GetTop3Donors(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 把地址从"Go 的内部格式"转成"人能看懂的 0x 字符串"
	addresses := make([]string, len(donors))
	for i, d := range donors {
		addresses[i] = d.Hex() // 比如 "0xabcd...1234"
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{"topDonors": addresses})
}

// ─── GET /contract/info ────────────────────────
//
// GetContractInfo 返回合约的基本信息。
//
// 使用方式：
//   curl http://localhost:8080/contract/info
//
// 返回格式：
//   {
//     "owner": "0xaaa...",
//     "donateDeadline": "1734567890"
//   }
//
// owner：合约部署者的钱包地址
// donateDeadline：捐款截止时间（Unix 时间戳，单位秒）
//   你可以用这个网站转成人类时间：https://tool.chinaz.com/tools/unixtime.aspx
//

func (cc *ContractController) GetContractInfo(c *gin.Context) {
	// 查合约所有者
	owner, err := cc.contractSrv.GetOwner(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 查捐款截止时间
	deadline, err := cc.contractSrv.GetDonateDeadline(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 把结果拼起来返回
	c.JSON(http.StatusOK, gin.H{
		"owner":          owner.Hex(),
		"donateDeadline": deadline.String(),
	})
}

// ─── POST /contract/donate ────────────────────
//
// Donate 往合约里捐款。
//
// 使用方式：
//   curl -X POST http://localhost:8080/contract/donate \
//     -H "Content-Type: application/json" \
//     -d '{"amount": "1000000000000000000"}'
//
// 请求体：
//   {
//     "amount": "1000000000000000000"
//   }
//
// amount 是你想捐的金额，单位 wei。
//   1 ETH    = 1000000000000000000 wei（18 个零）
//   0.01 ETH = 10000000000000000  wei（16 个零）
//
// 返回格式：
//   {
//     "txHash": "0xabc...123"
//   }
//
// txHash 是这笔交易的哈希值。
// 你可以在 Sepolia 区块链浏览器上查这笔交易的状态：
//   https://sepolia.etherscan.io/tx/0xabc...123
//
// 注意：
//   - 这笔操作需要花 Gas（手续费），会从你的钱包里扣
//   - amount 不能是 0 或负数
//   - 如果合约捐款已截止，会返回错误
//

func (cc *ContractController) Donate(c *gin.Context) {
	// 第 1 步：解析请求体里的 JSON
	var req DonateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// JSON 格式不对，或者缺少 amount 字段
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 第 2 步：把 amount 从字符串解析成大整数
	// 因为金额可能超级大（10^18 级别），Go 的普通 int 装不下
	// 所以要用 big.Int（大整数）
	amount := new(big.Int)
	if _, ok := amount.SetString(req.Amount, 10); !ok || amount.Sign() <= 0 {
		// 金额不是合法的数字，或者小于等于 0
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	// 第 3 步：调用合约服务去捐款
	// 这个过程会：签名交易 → 发到以太坊 → 等待结果
	txHash, err := cc.contractSrv.Donate(c, amount)
	if err != nil {
		// 捐款失败（可能是 Gas 不够、网络问题、合约拒绝等）
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 第 4 步：捐款成功，返回交易哈希
	c.JSON(http.StatusOK, gin.H{"txHash": txHash.Hex()})
}

// ─── POST /contract/withdraw ──────────────────
//
// Withdraw 把合约里收到的所有捐款提取到部署者钱包。
// 只有合约的 owner（部署合约的人）才能调用成功，其他人调用会报错。
//
// 使用方式：
//   curl -X POST http://localhost:8080/contract/withdraw
//
// 返回格式：
//   {
//     "txHash": "0xabc...123"
//   }
//
// txHash 是提现交易的哈希值。
// 你可以在 Sepolia 区块链浏览器上追踪：
//   https://sepolia.etherscan.io/tx/0xabc...123
//

func (cc *ContractController) Withdraw(c *gin.Context) {
	txHash, err := cc.contractSrv.Withdraw(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("提现交易已发送，txHash: %s", txHash.Hex())

	c.JSON(http.StatusOK, gin.H{"txHash": txHash.Hex()})
}

// ──────────────────────────────────────────────
// 合约 HTTP 控制器 — 结束
// ──────────────────────────────────────────────
