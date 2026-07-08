// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Donate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donateDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"donations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"}],\"name\":\"getDonation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTop3Donors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"topDonations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"topDonors\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b503360015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062278d0042605c9190609c565b60028190555060c8565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f60a4826066565b915060ad836066565b925082820190508082111560c25760c1606f565b5b92915050565b611196806100d55f395ff3fe608060405260043610610089575f3560e01c80635d2fae29116100585780635d2fae29146101545780638da5cb5b1461017e578063b76a4824146101a8578063cc6cb19a146101e4578063ed88c68e1461022057610098565b80632d1e25181461009c5780633ccfd60b146100c6578063410a1d32146100dc5780635cbaddab1461011857610098565b366100985761009661022a565b005b5f5ffd5b3480156100a7575f5ffd5b506100b061039a565b6040516100bd9190610c38565b60405180910390f35b3480156100d1575f5ffd5b506100da61048e565b005b3480156100e7575f5ffd5b5061010260048036038101906100fd9190610c86565b610630565b60405161010f9190610cc9565b60405180910390f35b348015610123575f5ffd5b5061013e60048036038101906101399190610d0c565b610675565b60405161014b9190610d46565b60405180910390f35b34801561015f575f5ffd5b506101686106a9565b6040516101759190610cc9565b60405180910390f35b348015610189575f5ffd5b506101926106af565b60405161019f9190610d46565b60405180910390f35b3480156101b3575f5ffd5b506101ce60048036038101906101c99190610d0c565b6106d4565b6040516101db9190610cc9565b60405180910390f35b3480156101ef575f5ffd5b5061020a60048036038101906102059190610c86565b6106ed565b6040516102179190610cc9565b60405180910390f35b61022861022a565b005b5f341161026c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026390610db9565b60405180910390fd5b60025442106102b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102a790610e21565b60405180910390fd5b345f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8282546102fb9190610e6c565b92505081905550610348335f5f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610701565b3373ffffffffffffffffffffffffffffffffffffffff167f49bb2edd8608e46c3b9823ec2dabe8ddf219c00870a4a112ff60a1a70669f4573442604051610390929190610e9f565b60405180910390a2565b60605f600367ffffffffffffffff8111156103b8576103b7610ec6565b5b6040519080825280602002602001820160405280156103e65781602001602082028036833780820191505090505b5090505f5f90505b6003811015610486576003816003811061040b5761040a610ef3565b5b015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682828151811061043f5761043e610ef3565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806001019150506103ee565b508091505090565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461051d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051490610f90565b60405180910390fd5b5f4790505f8111610563576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055a90610ff8565b60405180910390fd5b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826040516105a990611043565b5f6040518083038185875af1925050503d805f81146105e3576040519150601f19603f3d011682016040523d82523d5f602084013e6105e8565b606091505b505090508061062c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610623906110a1565b60405180910390fd5b5050565b5f5f5f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b60038160038110610684575f80fd5b015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600681600381106106e3575f80fd5b015f915090505481565b5f602052805f5260405f205f915090505481565b5f5f90505b60038110156109b0578273ffffffffffffffffffffffffffffffffffffffff166003826003811061073a57610739610ef3565b5b015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361099d57816006826003811061078b5761078a610ef3565b5b01819055505b5f811180156107d5575060066001826107aa91906110bf565b600381106107bb576107ba610ef3565b5b0154600682600381106107d1576107d0610ef3565b5b0154115b156109975760036001826107e991906110bf565b600381106107fa576107f9610ef3565b5b015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166003826003811061082f5761082e610ef3565b5b015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff166003836003811061086457610863610ef3565b5b015f600360018661087591906110bf565b6003811061088657610885610ef3565b5b015f8491906101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508391906101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600660018261091791906110bf565b6003811061092857610927610ef3565b5b01546006826003811061093e5761093d610ef3565b5b01546006836003811061095457610953610ef3565b5b015f600660018661096591906110bf565b6003811061097657610975610ef3565b5b015f849190505583919050555050808061098f906110f2565b915050610791565b50610b4d565b80806109a890611119565b915050610706565b505f5f90505b6003811015610b4b57600681600381106109d3576109d2610ef3565b5b0154821115610b3e575f600290505b81811115610acb5760036001826109f991906110bf565b60038110610a0a57610a09610ef3565b5b015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660038260038110610a3f57610a3e610ef3565b5b015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506006600182610a8c91906110bf565b60038110610a9d57610a9c610ef3565b5b015460068260038110610ab357610ab2610ef3565b5b01819055508080610ac3906110f2565b9150506109e2565b508260038260038110610ae157610ae0610ef3565b5b015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160068260038110610b3457610b33610ef3565b5b0181905550610b4b565b80806001019150506109b6565b505b5050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ba382610b7a565b9050919050565b610bb381610b99565b82525050565b5f610bc48383610baa565b60208301905092915050565b5f602082019050919050565b5f610be682610b51565b610bf08185610b5b565b9350610bfb83610b6b565b805f5b83811015610c2b578151610c128882610bb9565b9750610c1d83610bd0565b925050600181019050610bfe565b5085935050505092915050565b5f6020820190508181035f830152610c508184610bdc565b905092915050565b5f5ffd5b610c6581610b99565b8114610c6f575f5ffd5b50565b5f81359050610c8081610c5c565b92915050565b5f60208284031215610c9b57610c9a610c58565b5b5f610ca884828501610c72565b91505092915050565b5f819050919050565b610cc381610cb1565b82525050565b5f602082019050610cdc5f830184610cba565b92915050565b610ceb81610cb1565b8114610cf5575f5ffd5b50565b5f81359050610d0681610ce2565b92915050565b5f60208284031215610d2157610d20610c58565b5b5f610d2e84828501610cf8565b91505092915050565b610d4081610b99565b82525050565b5f602082019050610d595f830184610d37565b92915050565b5f82825260208201905092915050565b7f446f6e6174696f6e206d7573742062652067726561746572207468616e20302e5f82015250565b5f610da3602083610d5f565b9150610dae82610d6f565b602082019050919050565b5f6020820190508181035f830152610dd081610d97565b9050919050565b7f446f6e6174696f6e20706572696f642068617320656e6465642e0000000000005f82015250565b5f610e0b601a83610d5f565b9150610e1682610dd7565b602082019050919050565b5f6020820190508181035f830152610e3881610dff565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610e7682610cb1565b9150610e8183610cb1565b9250828201905080821115610e9957610e98610e3f565b5b92915050565b5f604082019050610eb25f830185610cba565b610ebf6020830184610cba565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4f6e6c7920746865206f776e65722063616e2063616c6c20746869732066756e5f8201527f6374696f6e2e0000000000000000000000000000000000000000000000000000602082015250565b5f610f7a602683610d5f565b9150610f8582610f20565b604082019050919050565b5f6020820190508181035f830152610fa781610f6e565b9050919050565b7f4e6f2066756e647320746f2077697468647261772e00000000000000000000005f82015250565b5f610fe2601583610d5f565b9150610fed82610fae565b602082019050919050565b5f6020820190508181035f83015261100f81610fd6565b9050919050565b5f81905092915050565b50565b5f61102e5f83611016565b915061103982611020565b5f82019050919050565b5f61104d82611023565b9150819050919050565b7f5769746864726177206661696c65642e000000000000000000000000000000005f82015250565b5f61108b601083610d5f565b915061109682611057565b602082019050919050565b5f6020820190508181035f8301526110b88161107f565b9050919050565b5f6110c982610cb1565b91506110d483610cb1565b92508282039050818111156110ec576110eb610e3f565b5b92915050565b5f6110fc82610cb1565b91505f820361110e5761110d610e3f565b5b600182039050919050565b5f61112382610cb1565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361115557611154610e3f565b5b60018201905091905056fea2646970667358221220dc819060f40e2ce11cdc76929dccb70c59607b1c3528205f8f47bf90eba9b0bb64736f6c634300081c0033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DonateDeadline is a free data retrieval call binding the contract method 0x5d2fae29.
//
// Solidity: function donateDeadline() view returns(uint256)
func (_Contract *ContractCaller) DonateDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "donateDeadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DonateDeadline is a free data retrieval call binding the contract method 0x5d2fae29.
//
// Solidity: function donateDeadline() view returns(uint256)
func (_Contract *ContractSession) DonateDeadline() (*big.Int, error) {
	return _Contract.Contract.DonateDeadline(&_Contract.CallOpts)
}

// DonateDeadline is a free data retrieval call binding the contract method 0x5d2fae29.
//
// Solidity: function donateDeadline() view returns(uint256)
func (_Contract *ContractCallerSession) DonateDeadline() (*big.Int, error) {
	return _Contract.Contract.DonateDeadline(&_Contract.CallOpts)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_Contract *ContractCaller) Donations(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "donations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_Contract *ContractSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Donations(&_Contract.CallOpts, arg0)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_Contract *ContractCallerSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.Donations(&_Contract.CallOpts, arg0)
}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address donor) view returns(uint256)
func (_Contract *ContractCaller) GetDonation(opts *bind.CallOpts, donor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getDonation", donor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address donor) view returns(uint256)
func (_Contract *ContractSession) GetDonation(donor common.Address) (*big.Int, error) {
	return _Contract.Contract.GetDonation(&_Contract.CallOpts, donor)
}

// GetDonation is a free data retrieval call binding the contract method 0x410a1d32.
//
// Solidity: function getDonation(address donor) view returns(uint256)
func (_Contract *ContractCallerSession) GetDonation(donor common.Address) (*big.Int, error) {
	return _Contract.Contract.GetDonation(&_Contract.CallOpts, donor)
}

// GetTop3Donors is a free data retrieval call binding the contract method 0x2d1e2518.
//
// Solidity: function getTop3Donors() view returns(address[])
func (_Contract *ContractCaller) GetTop3Donors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTop3Donors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTop3Donors is a free data retrieval call binding the contract method 0x2d1e2518.
//
// Solidity: function getTop3Donors() view returns(address[])
func (_Contract *ContractSession) GetTop3Donors() ([]common.Address, error) {
	return _Contract.Contract.GetTop3Donors(&_Contract.CallOpts)
}

// GetTop3Donors is a free data retrieval call binding the contract method 0x2d1e2518.
//
// Solidity: function getTop3Donors() view returns(address[])
func (_Contract *ContractCallerSession) GetTop3Donors() ([]common.Address, error) {
	return _Contract.Contract.GetTop3Donors(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// TopDonations is a free data retrieval call binding the contract method 0xb76a4824.
//
// Solidity: function topDonations(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) TopDonations(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "topDonations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TopDonations is a free data retrieval call binding the contract method 0xb76a4824.
//
// Solidity: function topDonations(uint256 ) view returns(uint256)
func (_Contract *ContractSession) TopDonations(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.TopDonations(&_Contract.CallOpts, arg0)
}

// TopDonations is a free data retrieval call binding the contract method 0xb76a4824.
//
// Solidity: function topDonations(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) TopDonations(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.TopDonations(&_Contract.CallOpts, arg0)
}

// TopDonors is a free data retrieval call binding the contract method 0x5cbaddab.
//
// Solidity: function topDonors(uint256 ) view returns(address)
func (_Contract *ContractCaller) TopDonors(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "topDonors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TopDonors is a free data retrieval call binding the contract method 0x5cbaddab.
//
// Solidity: function topDonors(uint256 ) view returns(address)
func (_Contract *ContractSession) TopDonors(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.TopDonors(&_Contract.CallOpts, arg0)
}

// TopDonors is a free data retrieval call binding the contract method 0x5cbaddab.
//
// Solidity: function topDonors(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) TopDonors(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.TopDonors(&_Contract.CallOpts, arg0)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Contract *ContractTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Contract *ContractSession) Donate() (*types.Transaction, error) {
	return _Contract.Contract.Donate(&_Contract.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Contract *ContractTransactorSession) Donate() (*types.Transaction, error) {
	return _Contract.Contract.Donate(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contract *ContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Contract *ContractTransactorSession) Receive() (*types.Transaction, error) {
	return _Contract.Contract.Receive(&_Contract.TransactOpts)
}

// ContractDonateIterator is returned from FilterDonate and is used to iterate over the raw logs and unpacked data for Donate events raised by the Contract contract.
type ContractDonateIterator struct {
	Event *ContractDonate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDonateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDonate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDonate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDonateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDonateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDonate represents a Donate event raised by the Contract contract.
type ContractDonate struct {
	Donor     common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDonate is a free log retrieval operation binding the contract event 0x49bb2edd8608e46c3b9823ec2dabe8ddf219c00870a4a112ff60a1a70669f457.
//
// Solidity: event Donate(address indexed donor, uint256 amount, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDonate(opts *bind.FilterOpts, donor []common.Address) (*ContractDonateIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Donate", donorRule)
	if err != nil {
		return nil, err
	}
	return &ContractDonateIterator{contract: _Contract.contract, event: "Donate", logs: logs, sub: sub}, nil
}

// WatchDonate is a free log subscription operation binding the contract event 0x49bb2edd8608e46c3b9823ec2dabe8ddf219c00870a4a112ff60a1a70669f457.
//
// Solidity: event Donate(address indexed donor, uint256 amount, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDonate(opts *bind.WatchOpts, sink chan<- *ContractDonate, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Donate", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDonate)
				if err := _Contract.contract.UnpackLog(event, "Donate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDonate is a log parse operation binding the contract event 0x49bb2edd8608e46c3b9823ec2dabe8ddf219c00870a4a112ff60a1a70669f457.
//
// Solidity: event Donate(address indexed donor, uint256 amount, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDonate(log types.Log) (*ContractDonate, error) {
	event := new(ContractDonate)
	if err := _Contract.contract.UnpackLog(event, "Donate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
