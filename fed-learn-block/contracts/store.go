// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"compareAccessPolicy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"compareMeta1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"compareMeta2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generate_noise\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPKC\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPKF\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input1\",\"type\":\"string\"}],\"name\":\"setAccessP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"input2\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"input3\",\"type\":\"bytes\"}],\"name\":\"setMeta1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"input2\",\"type\":\"string\"}],\"name\":\"setMeta2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setPKC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setPKF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input2\",\"type\":\"bytes\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50610afd8061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80638ee3d37a1161008c578063c78e0f7111610066578063c78e0f711461019e578063d0ec2e3a146101b1578063d41c4b60146101b9578063ffd58c6d146101cc57600080fd5b80638ee3d37a1461016357806392e0bdfd14610178578063be6a2a941461018b57600080fd5b80630944a67f146100d457806334e784a1146100fc57806339a40fe9146101115780635c7563f4146101245780635e615a6b1461013a57806386e0829314610150575b600080fd5b6100e76100e236600461067d565b6101df565b60405190151581526020015b60405180910390f35b61010f61010a36600461067d565b610238565b005b61010f61011f3660046106ba565b610248565b61012c610266565b6040519081526020016100f3565b610142610336565b6040516100f3929190610773565b6100e761015e36600461067d565b61045d565b61016b610498565b6040516100f391906107a1565b6100e761018636600461067d565b61052a565b61010f61019936600461067d565b610565565b61010f6101ac3660046107bb565b610571565b61016b61059d565b61010f6101c73660046106ba565b6105ac565b61010f6101da36600461067d565b6105c5565b6000816040516020016101f2919061084f565b60405160208183030381529060405280519060200120600560405160200161021a91906108a5565b60405160208183030381529060405280519060200120149050919050565b60036102448282610968565b5050565b60006102548382610968565b5060016102618282610968565b505050565b600080424433610277600143610a3d565b6040805160208101959095528401929092526bffffffffffffffffffffffff19606091821b16908301524060748201526094016040516020818303038152906040528051906020012060001c90506000805b600481101561032157604080516020810185905290810182905260600160408051601f198184030181529190528051602090910120925061030d620f424084610a56565b6103179083610a78565b91506001016102c9565b5061032f621e848082610aa0565b9250505090565b606080600060018180546103499061086b565b80601f01602080910402602001604051908101604052809291908181526020018280546103759061086b565b80156103c25780601f10610397576101008083540402835291602001916103c2565b820191906000526020600020905b8154815290600101906020018083116103a557829003601f168201915b505050505091508080546103d59061086b565b80601f01602080910402602001604051908101604052809291908181526020018280546104019061086b565b801561044e5780601f106104235761010080835404028352916020019161044e565b820191906000526020600020905b81548152906001019060200180831161043157829003601f168201915b50505050509050915091509091565b600081604051602001610470919061084f565b60405160208183030381529060405280519060200120600860405160200161021a91906108a5565b6060600380546104a79061086b565b80601f01602080910402602001604051908101604052809291908181526020018280546104d39061086b565b80156105205780601f106104f557610100808354040283529160200191610520565b820191906000526020600020905b81548152906001019060200180831161050357829003601f168201915b5050505050905090565b60008160405160200161053d919061084f565b60405160208183030381529060405280519060200120600960405160200161021a91906108a5565b60026102448282610968565b600461057d8482610968565b50600561058a8382610968565b5060066105978282610968565b50505050565b6060600280546104a79061086b565b60076105b88382610968565b5060086102618282610968565b60096102448282610968565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126105f857600080fd5b81356020830160008067ffffffffffffffff841115610619576106196105d1565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff82111715610648576106486105d1565b60405283815290508082840187101561066057600080fd5b838360208301376000602085830101528094505050505092915050565b60006020828403121561068f57600080fd5b813567ffffffffffffffff8111156106a657600080fd5b6106b2848285016105e7565b949350505050565b600080604083850312156106cd57600080fd5b823567ffffffffffffffff8111156106e457600080fd5b6106f0858286016105e7565b925050602083013567ffffffffffffffff81111561070d57600080fd5b610719858286016105e7565b9150509250929050565b60005b8381101561073e578181015183820152602001610726565b50506000910152565b6000815180845261075f816020860160208601610723565b601f01601f19169290920160200192915050565b6040815260006107866040830185610747565b82810360208401526107988185610747565b95945050505050565b6020815260006107b46020830184610747565b9392505050565b6000806000606084860312156107d057600080fd5b833567ffffffffffffffff8111156107e757600080fd5b6107f3868287016105e7565b935050602084013567ffffffffffffffff81111561081057600080fd5b61081c868287016105e7565b925050604084013567ffffffffffffffff81111561083957600080fd5b610845868287016105e7565b9150509250925092565b60008251610861818460208701610723565b9190910192915050565b600181811c9082168061087f57607f821691505b60208210810361089f57634e487b7160e01b600052602260045260246000fd5b50919050565b60008083546108b38161086b565b6001821680156108ca57600181146108df5761090f565b60ff198316865281151582028601935061090f565b86600052602060002060005b83811015610907578154888201526001909101906020016108eb565b505081860193505b509195945050505050565b601f82111561026157806000526020600020601f840160051c810160208510156109415750805b601f840160051c820191505b81811015610961576000815560010161094d565b5050505050565b815167ffffffffffffffff811115610982576109826105d1565b61099681610990845461086b565b8461091a565b6020601f8211600181146109ca57600083156109b25750848201515b600019600385901b1c1916600184901b178455610961565b600084815260208120601f198516915b828110156109fa57878501518255602094850194600190920191016109da565b5084821015610a185786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b81810381811115610a5057610a50610a27565b92915050565b600082610a7357634e487b7160e01b600052601260045260246000fd5b500690565b8082018281126000831280158216821582161715610a9857610a98610a27565b505092915050565b8181036000831280158383131683831282161715610ac057610ac0610a27565b509291505056fea26469706673582212200cf0236c8cdca246a689abaf06826ac4c3c70096bb2bbfedd99112ca3a511fc464736f6c634300081e0033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// CompareAccessPolicy is a free data retrieval call binding the contract method 0x92e0bdfd.
//
// Solidity: function compareAccessPolicy(string req) view returns(bool)
func (_Contracts *ContractsCaller) CompareAccessPolicy(opts *bind.CallOpts, req string) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "compareAccessPolicy", req)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareAccessPolicy is a free data retrieval call binding the contract method 0x92e0bdfd.
//
// Solidity: function compareAccessPolicy(string req) view returns(bool)
func (_Contracts *ContractsSession) CompareAccessPolicy(req string) (bool, error) {
	return _Contracts.Contract.CompareAccessPolicy(&_Contracts.CallOpts, req)
}

// CompareAccessPolicy is a free data retrieval call binding the contract method 0x92e0bdfd.
//
// Solidity: function compareAccessPolicy(string req) view returns(bool)
func (_Contracts *ContractsCallerSession) CompareAccessPolicy(req string) (bool, error) {
	return _Contracts.Contract.CompareAccessPolicy(&_Contracts.CallOpts, req)
}

// CompareMeta1 is a free data retrieval call binding the contract method 0x0944a67f.
//
// Solidity: function compareMeta1(string req) view returns(bool)
func (_Contracts *ContractsCaller) CompareMeta1(opts *bind.CallOpts, req string) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "compareMeta1", req)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareMeta1 is a free data retrieval call binding the contract method 0x0944a67f.
//
// Solidity: function compareMeta1(string req) view returns(bool)
func (_Contracts *ContractsSession) CompareMeta1(req string) (bool, error) {
	return _Contracts.Contract.CompareMeta1(&_Contracts.CallOpts, req)
}

// CompareMeta1 is a free data retrieval call binding the contract method 0x0944a67f.
//
// Solidity: function compareMeta1(string req) view returns(bool)
func (_Contracts *ContractsCallerSession) CompareMeta1(req string) (bool, error) {
	return _Contracts.Contract.CompareMeta1(&_Contracts.CallOpts, req)
}

// CompareMeta2 is a free data retrieval call binding the contract method 0x86e08293.
//
// Solidity: function compareMeta2(string req) view returns(bool)
func (_Contracts *ContractsCaller) CompareMeta2(opts *bind.CallOpts, req string) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "compareMeta2", req)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CompareMeta2 is a free data retrieval call binding the contract method 0x86e08293.
//
// Solidity: function compareMeta2(string req) view returns(bool)
func (_Contracts *ContractsSession) CompareMeta2(req string) (bool, error) {
	return _Contracts.Contract.CompareMeta2(&_Contracts.CallOpts, req)
}

// CompareMeta2 is a free data retrieval call binding the contract method 0x86e08293.
//
// Solidity: function compareMeta2(string req) view returns(bool)
func (_Contracts *ContractsCallerSession) CompareMeta2(req string) (bool, error) {
	return _Contracts.Contract.CompareMeta2(&_Contracts.CallOpts, req)
}

// GenerateNoise is a free data retrieval call binding the contract method 0x5c7563f4.
//
// Solidity: function generate_noise() view returns(int256)
func (_Contracts *ContractsCaller) GenerateNoise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "generate_noise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenerateNoise is a free data retrieval call binding the contract method 0x5c7563f4.
//
// Solidity: function generate_noise() view returns(int256)
func (_Contracts *ContractsSession) GenerateNoise() (*big.Int, error) {
	return _Contracts.Contract.GenerateNoise(&_Contracts.CallOpts)
}

// GenerateNoise is a free data retrieval call binding the contract method 0x5c7563f4.
//
// Solidity: function generate_noise() view returns(int256)
func (_Contracts *ContractsCallerSession) GenerateNoise() (*big.Int, error) {
	return _Contracts.Contract.GenerateNoise(&_Contracts.CallOpts)
}

// GetPKC is a free data retrieval call binding the contract method 0xd0ec2e3a.
//
// Solidity: function getPKC() view returns(bytes)
func (_Contracts *ContractsCaller) GetPKC(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPKC")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPKC is a free data retrieval call binding the contract method 0xd0ec2e3a.
//
// Solidity: function getPKC() view returns(bytes)
func (_Contracts *ContractsSession) GetPKC() ([]byte, error) {
	return _Contracts.Contract.GetPKC(&_Contracts.CallOpts)
}

// GetPKC is a free data retrieval call binding the contract method 0xd0ec2e3a.
//
// Solidity: function getPKC() view returns(bytes)
func (_Contracts *ContractsCallerSession) GetPKC() ([]byte, error) {
	return _Contracts.Contract.GetPKC(&_Contracts.CallOpts)
}

// GetPKF is a free data retrieval call binding the contract method 0x8ee3d37a.
//
// Solidity: function getPKF() view returns(bytes)
func (_Contracts *ContractsCaller) GetPKF(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getPKF")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPKF is a free data retrieval call binding the contract method 0x8ee3d37a.
//
// Solidity: function getPKF() view returns(bytes)
func (_Contracts *ContractsSession) GetPKF() ([]byte, error) {
	return _Contracts.Contract.GetPKF(&_Contracts.CallOpts)
}

// GetPKF is a free data retrieval call binding the contract method 0x8ee3d37a.
//
// Solidity: function getPKF() view returns(bytes)
func (_Contracts *ContractsCallerSession) GetPKF() ([]byte, error) {
	return _Contracts.Contract.GetPKF(&_Contracts.CallOpts)
}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns(bytes, bytes)
func (_Contracts *ContractsCaller) GetParams(opts *bind.CallOpts) ([]byte, []byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getParams")

	if err != nil {
		return *new([]byte), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns(bytes, bytes)
func (_Contracts *ContractsSession) GetParams() ([]byte, []byte, error) {
	return _Contracts.Contract.GetParams(&_Contracts.CallOpts)
}

// GetParams is a free data retrieval call binding the contract method 0x5e615a6b.
//
// Solidity: function getParams() view returns(bytes, bytes)
func (_Contracts *ContractsCallerSession) GetParams() ([]byte, []byte, error) {
	return _Contracts.Contract.GetParams(&_Contracts.CallOpts)
}

// SetAccessP is a paid mutator transaction binding the contract method 0xffd58c6d.
//
// Solidity: function setAccessP(string input1) returns()
func (_Contracts *ContractsTransactor) SetAccessP(opts *bind.TransactOpts, input1 string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setAccessP", input1)
}

// SetAccessP is a paid mutator transaction binding the contract method 0xffd58c6d.
//
// Solidity: function setAccessP(string input1) returns()
func (_Contracts *ContractsSession) SetAccessP(input1 string) (*types.Transaction, error) {
	return _Contracts.Contract.SetAccessP(&_Contracts.TransactOpts, input1)
}

// SetAccessP is a paid mutator transaction binding the contract method 0xffd58c6d.
//
// Solidity: function setAccessP(string input1) returns()
func (_Contracts *ContractsTransactorSession) SetAccessP(input1 string) (*types.Transaction, error) {
	return _Contracts.Contract.SetAccessP(&_Contracts.TransactOpts, input1)
}

// SetMeta1 is a paid mutator transaction binding the contract method 0xc78e0f71.
//
// Solidity: function setMeta1(string input1, string input2, bytes input3) returns()
func (_Contracts *ContractsTransactor) SetMeta1(opts *bind.TransactOpts, input1 string, input2 string, input3 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMeta1", input1, input2, input3)
}

// SetMeta1 is a paid mutator transaction binding the contract method 0xc78e0f71.
//
// Solidity: function setMeta1(string input1, string input2, bytes input3) returns()
func (_Contracts *ContractsSession) SetMeta1(input1 string, input2 string, input3 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetMeta1(&_Contracts.TransactOpts, input1, input2, input3)
}

// SetMeta1 is a paid mutator transaction binding the contract method 0xc78e0f71.
//
// Solidity: function setMeta1(string input1, string input2, bytes input3) returns()
func (_Contracts *ContractsTransactorSession) SetMeta1(input1 string, input2 string, input3 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetMeta1(&_Contracts.TransactOpts, input1, input2, input3)
}

// SetMeta2 is a paid mutator transaction binding the contract method 0xd41c4b60.
//
// Solidity: function setMeta2(string input1, string input2) returns()
func (_Contracts *ContractsTransactor) SetMeta2(opts *bind.TransactOpts, input1 string, input2 string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setMeta2", input1, input2)
}

// SetMeta2 is a paid mutator transaction binding the contract method 0xd41c4b60.
//
// Solidity: function setMeta2(string input1, string input2) returns()
func (_Contracts *ContractsSession) SetMeta2(input1 string, input2 string) (*types.Transaction, error) {
	return _Contracts.Contract.SetMeta2(&_Contracts.TransactOpts, input1, input2)
}

// SetMeta2 is a paid mutator transaction binding the contract method 0xd41c4b60.
//
// Solidity: function setMeta2(string input1, string input2) returns()
func (_Contracts *ContractsTransactorSession) SetMeta2(input1 string, input2 string) (*types.Transaction, error) {
	return _Contracts.Contract.SetMeta2(&_Contracts.TransactOpts, input1, input2)
}

// SetPKC is a paid mutator transaction binding the contract method 0xbe6a2a94.
//
// Solidity: function setPKC(bytes input) returns()
func (_Contracts *ContractsTransactor) SetPKC(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPKC", input)
}

// SetPKC is a paid mutator transaction binding the contract method 0xbe6a2a94.
//
// Solidity: function setPKC(bytes input) returns()
func (_Contracts *ContractsSession) SetPKC(input []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetPKC(&_Contracts.TransactOpts, input)
}

// SetPKC is a paid mutator transaction binding the contract method 0xbe6a2a94.
//
// Solidity: function setPKC(bytes input) returns()
func (_Contracts *ContractsTransactorSession) SetPKC(input []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetPKC(&_Contracts.TransactOpts, input)
}

// SetPKF is a paid mutator transaction binding the contract method 0x34e784a1.
//
// Solidity: function setPKF(bytes input) returns()
func (_Contracts *ContractsTransactor) SetPKF(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setPKF", input)
}

// SetPKF is a paid mutator transaction binding the contract method 0x34e784a1.
//
// Solidity: function setPKF(bytes input) returns()
func (_Contracts *ContractsSession) SetPKF(input []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetPKF(&_Contracts.TransactOpts, input)
}

// SetPKF is a paid mutator transaction binding the contract method 0x34e784a1.
//
// Solidity: function setPKF(bytes input) returns()
func (_Contracts *ContractsTransactorSession) SetPKF(input []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetPKF(&_Contracts.TransactOpts, input)
}

// SetParams is a paid mutator transaction binding the contract method 0x39a40fe9.
//
// Solidity: function setParams(bytes input1, bytes input2) returns()
func (_Contracts *ContractsTransactor) SetParams(opts *bind.TransactOpts, input1 []byte, input2 []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setParams", input1, input2)
}

// SetParams is a paid mutator transaction binding the contract method 0x39a40fe9.
//
// Solidity: function setParams(bytes input1, bytes input2) returns()
func (_Contracts *ContractsSession) SetParams(input1 []byte, input2 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetParams(&_Contracts.TransactOpts, input1, input2)
}

// SetParams is a paid mutator transaction binding the contract method 0x39a40fe9.
//
// Solidity: function setParams(bytes input1, bytes input2) returns()
func (_Contracts *ContractsTransactorSession) SetParams(input1 []byte, input2 []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetParams(&_Contracts.TransactOpts, input1, input2)
}
