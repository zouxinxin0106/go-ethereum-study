package service

// ──────────────────────────────────────────────
// 合约交互服务层
//
// 这个文件的作用是：把我们和智能合约的交互（比如查捐款、捐款、提现）
// 包装成一个个好用的 Go 函数，这样外面调用起来就很简单。
// ──────────────────────────────────────────────

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/FloraZou/go-eth-study/contract"
)

// ─── ContractService 结构体 ────────────────────
//
// ContractService 是"合约服务员"，帮你做所有和合约有关的事情。
// 它手里拿着三样东西：
//   1. ethClient   — 连接以太坊网络的"网线"
//   2. contract    — abigen 生成的合约"遥控器"（具体看 contract/BeggingContract.go）
//   3. privateKey  — 你的账户私钥，用来签名交易（证明"是你本人在操作"）
//   4. fromAddress — 你的账户地址，从私钥算出来的
//

type ContractService struct {
	client      *ethclient.Client     // 连接以太坊网络的客户端
	contract    *contract.Contract    // abigen 生成的合约绑定对象，相当于合约的"遥控器"
	chainID     *big.Int              // 区块链的 ID（Sepolia 主网是 11155111，主网是 1）
	privateKey  *ecdsa.PrivateKey     // 你的私钥，用来给交易签名
	fromAddress common.Address        // 你的钱包地址，从私钥算出来的
}

// ─── 创建合约服务 ─────────────────────────────
//
// NewContractService 创建一个"合约服务员"。
// 你需要给它：
//   - client        — 连上以太坊的客户端
//   - contractAddr  — 合约部署到链上之后的地址（0x开头的一串）
//   - privateKeyHex — 你的私钥（十六进制字符串）
//   - chainID       — 区块链的 ID（Sepolia 是 11155111）
//
// 它会自动：
//   1. 用私钥算出你的钱包地址
//   2. 在合约地址上绑定好"遥控器"
//   3. 返回一个啥都能干的 ContractService 给你
//

func NewContractService(client *ethclient.Client, contractAddr common.Address, privateKeyHex string, chainID *big.Int) (*ContractService, error) {
	// 第 1 步：把私钥从字符串格式解析成 Go 能用的格式
	// 就好比把一串数字密码解析成真正的钥匙
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}

	// 第 2 步：从私钥算出公钥，再从公钥算出钱包地址
	// 这就好比：你的银行卡密码 → 银行卡号
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 第 3 步：把合约"遥控器"绑定到链上那个真实的合约地址
	// 以后你按遥控器上的按钮，就会真的去调用链上的合约
	instance, err := contract.NewContract(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &ContractService{
		client:      client,
		contract:    instance,
		chainID:     chainID,
		privateKey:  privateKey,
		fromAddress: fromAddress,
	}, nil
}

// ═══════════════════════════════════════════════
// 以下是"读"操作（不花 Gas，免费查询）
// ═══════════════════════════════════════════════

// ─── 查捐款金额 ───────────────────────────────
//
// GetDonation 查询某个地址捐了多少钱。
// 参数 donor 是你要查的那个钱包地址。
// 返回的是捐款金额（单位是 wei，1 ETH = 10^18 wei）。
//
// 举例：GetDonation(ctx, "0x123...") → 返回 1000000000000000000（= 1 ETH）
//

func (s *ContractService) GetDonation(ctx context.Context, donor common.Address) (*big.Int, error) {
	// 直接调用合约上生成的 GetDonation 方法
	// CallOpts 是调用选项，这里只传了 Context，用来控制超时
	return s.contract.GetDonation(&bind.CallOpts{Context: ctx}, donor)
}

// ─── 查前三名捐款者 ───────────────────────────
//
// GetTop3Donors 返回一个列表，里面是捐款最多的三个人的钱包地址。
// 排行第一的在最前面，第二在中间，第三在最后。
//

func (s *ContractService) GetTop3Donors(ctx context.Context) ([]common.Address, error) {
	return s.contract.GetTop3Donors(&bind.CallOpts{Context: ctx})
}

// ─── 查捐款截止时间 ──────────────────────────
//
// GetDonateDeadline 返回捐款什么时间截止。
// 返回的是一个 Unix 时间戳（秒），你可以把它转成人类能看懂的日期。
//
// 合约里设置的是：部署之后 30 天截止。
//

func (s *ContractService) GetDonateDeadline(ctx context.Context) (*big.Int, error) {
	return s.contract.DonateDeadline(&bind.CallOpts{Context: ctx})
}

// ─── 查合约所有者 ─────────────────────────────
//
// GetOwner 返回合约是谁部署的（谁是这个合约的"老板"）。
// 只有这个地址才有权限调用 withdraw（提现）方法。
//

func (s *ContractService) GetOwner(ctx context.Context) (common.Address, error) {
	return s.contract.Owner(&bind.CallOpts{Context: ctx})
}

// ═══════════════════════════════════════════════
// 以下是"写"操作（需要花 Gas，需要签名）
// ═══════════════════════════════════════════════

// ─── 捐款 ─────────────────────────────────────
//
// Donate 往合约里捐款，需要带上 amount（单位 wei）。
//
// 这个操作会：
//   1. 用你的私钥签名一笔交易
//   2. 把签名后的交易发到以太坊网络
//   3. 合约收到后，记录你的捐款，并更新排行榜
//
// 返回的是这笔交易的哈希值（txHash），
// 你可以用 txHash 去区块链浏览器上查这笔交易的状态。
//
// 注意：amount 的单位是 wei。
// 如果你要捐 0.01 ETH，amount 要传 10000000000000000（10^16 wei）。
//

func (s *ContractService) Donate(ctx context.Context, amount *big.Int) (common.Hash, error) {
	// 创建交易选项（包括 Gas 价格、Nonce 等）
	auth, err := s.newTransactOpts(ctx, amount)
	if err != nil {
		return common.Hash{}, err
	}

	// 调用合约的 donate() 函数
	// 注意：合约里的 donate 是 payable 的，意味着可以接收 ETH
	// auth.Value 里放的就是你要捐的金额
	tx, err := s.contract.Donate(auth)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

// ─── 提现（仅合约所有者可调用）────────────────
//
// Withdraw 把合约里收到的所有 ETH 提取到部署者的钱包里。
// 只有合约的 owner（部署者）才能调用，其他人调用会失败。
//
// 返回交易哈希，方便你追踪。
//

func (s *ContractService) Withdraw(ctx context.Context) (common.Hash, error) {
	// 提现不需要额外付钱，所以 value 传 0
	auth, err := s.newTransactOpts(ctx, big.NewInt(0))
	if err != nil {
		return common.Hash{}, err
	}

	tx, err := s.contract.Withdraw(auth)
	if err != nil {
		return common.Hash{}, err
	}
	return tx.Hash(), nil
}

// ═══════════════════════════════════════════════
// 内部工具方法
// ═══════════════════════════════════════════════

// ─── 创建交易选项 ─────────────────────────────
//
// newTransactOpts 是一个内部方法（首字母小写，外部用不了）。
// 它的作用是把"我想做一件事"这个意图，包装成以太坊能理解的交易参数。
//
// 具体来说，它会：
//   1. 问以太坊网络：现在 Gas 价格是多少？（SuggestGasPrice）
//   2. 问以太坊网络：我这个地址第几笔交易了？（PendingNonceAt）
//   3. 用私钥和 chainID 创建一个签名者（NewKeyedTransactorWithChainID）
//   4. 把这些参数组装好，返回给调用者
//
// 这就好比你要寄快递：
//   1. 问快递员：现在邮费多少？
//   2. 看看你发了多少件了（防止重复发件）
//   3. 填好寄件人信息，签字
//   4. 寄出
//

func (s *ContractService) newTransactOpts(ctx context.Context, value *big.Int) (*bind.TransactOpts, error) {
	// 问以太坊节点：当前建议的 Gas 价格是多少？
	// Gas 就是手续费，价格越高，矿工越愿意帮你打包交易
	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	// 问以太坊节点：我这个地址目前发了多少笔交易了？
	// Nonce（交易序号）从 0 开始，每发一笔就 +1
	// 这个序号用来防止"重放攻击"——就是防止你把同一笔交易发两次
	nonce, err := s.client.PendingNonceAt(ctx, s.fromAddress)
	if err != nil {
		return nil, err
	}

	// 用你的私钥创建一个"签名器"
	// 以后每次发交易，都会自动用这个私钥签名
	// chainID 用来防止交易被拿到别的链上重放
	// 比如你在 Sepolia 签名的交易，拿到主网是无效的
	auth, err := bind.NewKeyedTransactorWithChainID(s.privateKey, s.chainID)
	if err != nil {
		return nil, err
	}

	// 把上面查到的参数填进交易选项里
	auth.Nonce = big.NewInt(int64(nonce)) // 交易序号
	auth.Value = value                     // 你要发送的 ETH 数量（wei）
	auth.GasLimit = uint64(300000)         // Gas 上限（合约调用一般 20~30 万就够了）
	auth.GasPrice = gasPrice               // Gas 单价
	auth.Context = ctx                     // 用来控制超时的上下文

	return auth, nil
}
