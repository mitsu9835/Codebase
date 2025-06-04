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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"compareAccessPolicy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"compareMeta1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"compareMeta2\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPKC\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPKF\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getParams\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input1\",\"type\":\"string\"}],\"name\":\"setAccessP\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"input2\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"input3\",\"type\":\"bytes\"}],\"name\":\"setMeta1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"input1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"input2\",\"type\":\"string\"}],\"name\":\"setMeta2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setPKC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"}],\"name\":\"setPKF\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"input1\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"input2\",\"type\":\"bytes\"}],\"name\":\"setParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506111e28061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806392e0bdfd1161008c578063d087d28811610066578063d087d28814610248578063d0ec2e3a14610266578063d41c4b6014610284578063ffd58c6d146102a0576100ea565b806392e0bdfd146101e0578063be6a2a9414610210578063c78e0f711461022c576100ea565b80635e615a6b116100c85780635e615a6b146101575780637bebc55b1461017657806386e08293146101925780638ee3d37a146101c2576100ea565b80630944a67f146100ef57806334e784a11461011f57806339a40fe91461013b575b600080fd5b610109600480360381019061010491906108c7565b6102bc565b604051610116919061092b565b60405180910390f35b610139600480360381019061013491906109e7565b610315565b005b61015560048036038101906101509190610a30565b610328565b005b61015f61034c565b60405161016d929190610b27565b60405180910390f35b610190600480360381019061018b91906109e7565b610473565b005b6101ac60048036038101906101a791906108c7565b610486565b6040516101b9919061092b565b60405180910390f35b6101ca6104df565b6040516101d79190610b5e565b60405180910390f35b6101fa60048036038101906101f591906108c7565b610571565b604051610207919061092b565b60405180910390f35b61022a600480360381019061022591906109e7565b6105ca565b005b61024660048036038101906102419190610b80565b6105dd565b005b610250610612565b60405161025d9190610b5e565b60405180910390f35b61026e6106a4565b60405161027b9190610b5e565b60405180910390f35b61029e60048036038101906102999190610c27565b610736565b005b6102ba60048036038101906102b591906108c7565b61075a565b005b6000816040516020016102cf9190610ce6565b6040516020818303038152906040528051906020012060056040516020016102f79190610df5565b60405160208183030381529060405280519060200120149050919050565b80600390816103249190610fc2565b5050565b81600090816103379190610fc2565b5080600190816103479190610fc2565b505050565b6060806000600181805461035f90610d2c565b80601f016020809104026020016040519081016040528092919081815260200182805461038b90610d2c565b80156103d85780601f106103ad576101008083540402835291602001916103d8565b820191906000526020600020905b8154815290600101906020018083116103bb57829003601f168201915b505050505091508080546103eb90610d2c565b80601f016020809104026020016040519081016040528092919081815260200182805461041790610d2c565b80156104645780601f1061043957610100808354040283529160200191610464565b820191906000526020600020905b81548152906001019060200180831161044757829003601f168201915b50505050509050915091509091565b80600a90816104829190610fc2565b5050565b6000816040516020016104999190610ce6565b6040516020818303038152906040528051906020012060086040516020016104c19190610df5565b60405160208183030381529060405280519060200120149050919050565b6060600380546104ee90610d2c565b80601f016020809104026020016040519081016040528092919081815260200182805461051a90610d2c565b80156105675780601f1061053c57610100808354040283529160200191610567565b820191906000526020600020905b81548152906001019060200180831161054a57829003601f168201915b5050505050905090565b6000816040516020016105849190610ce6565b6040516020818303038152906040528051906020012060096040516020016105ac9190610df5565b60405160208183030381529060405280519060200120149050919050565b80600290816105d99190610fc2565b5050565b82600490816105ec91906110da565b5081600590816105fc91906110da565b50806006908161060c9190610fc2565b50505050565b6060600a805461062190610d2c565b80601f016020809104026020016040519081016040528092919081815260200182805461064d90610d2c565b801561069a5780601f1061066f5761010080835404028352916020019161069a565b820191906000526020600020905b81548152906001019060200180831161067d57829003601f168201915b5050505050905090565b6060600280546106b390610d2c565b80601f01602080910402602001604051908101604052809291908181526020018280546106df90610d2c565b801561072c5780601f106107015761010080835404028352916020019161072c565b820191906000526020600020905b81548152906001019060200180831161070f57829003601f168201915b5050505050905090565b816007908161074591906110da565b50806008908161075591906110da565b505050565b806009908161076991906110da565b5050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6107d48261078b565b810181811067ffffffffffffffff821117156107f3576107f261079c565b5b80604052505050565b600061080661076d565b905061081282826107cb565b919050565b600067ffffffffffffffff8211156108325761083161079c565b5b61083b8261078b565b9050602081019050919050565b82818337600083830152505050565b600061086a61086584610817565b6107fc565b90508281526020810184848401111561088657610885610786565b5b610891848285610848565b509392505050565b600082601f8301126108ae576108ad610781565b5b81356108be848260208601610857565b91505092915050565b6000602082840312156108dd576108dc610777565b5b600082013567ffffffffffffffff8111156108fb576108fa61077c565b5b61090784828501610899565b91505092915050565b60008115159050919050565b61092581610910565b82525050565b6000602082019050610940600083018461091c565b92915050565b600067ffffffffffffffff8211156109615761096061079c565b5b61096a8261078b565b9050602081019050919050565b600061098a61098584610946565b6107fc565b9050828152602081018484840111156109a6576109a5610786565b5b6109b1848285610848565b509392505050565b600082601f8301126109ce576109cd610781565b5b81356109de848260208601610977565b91505092915050565b6000602082840312156109fd576109fc610777565b5b600082013567ffffffffffffffff811115610a1b57610a1a61077c565b5b610a27848285016109b9565b91505092915050565b60008060408385031215610a4757610a46610777565b5b600083013567ffffffffffffffff811115610a6557610a6461077c565b5b610a71858286016109b9565b925050602083013567ffffffffffffffff811115610a9257610a9161077c565b5b610a9e858286016109b9565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610ae2578082015181840152602081019050610ac7565b60008484015250505050565b6000610af982610aa8565b610b038185610ab3565b9350610b13818560208601610ac4565b610b1c8161078b565b840191505092915050565b60006040820190508181036000830152610b418185610aee565b90508181036020830152610b558184610aee565b90509392505050565b60006020820190508181036000830152610b788184610aee565b905092915050565b600080600060608486031215610b9957610b98610777565b5b600084013567ffffffffffffffff811115610bb757610bb661077c565b5b610bc386828701610899565b935050602084013567ffffffffffffffff811115610be457610be361077c565b5b610bf086828701610899565b925050604084013567ffffffffffffffff811115610c1157610c1061077c565b5b610c1d868287016109b9565b9150509250925092565b60008060408385031215610c3e57610c3d610777565b5b600083013567ffffffffffffffff811115610c5c57610c5b61077c565b5b610c6885828601610899565b925050602083013567ffffffffffffffff811115610c8957610c8861077c565b5b610c9585828601610899565b9150509250929050565b600081519050919050565b600081905092915050565b6000610cc082610c9f565b610cca8185610caa565b9350610cda818560208601610ac4565b80840191505092915050565b6000610cf28284610cb5565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610d4457607f821691505b602082108103610d5757610d56610cfd565b5b50919050565b60008190508160005260206000209050919050565b60008154610d7f81610d2c565b610d898186610caa565b94506001821660008114610da45760018114610db957610dec565b60ff1983168652811515820286019350610dec565b610dc285610d5d565b60005b83811015610de457815481890152600182019150602081019050610dc5565b838801955050505b50505092915050565b6000610e018284610d72565b915081905092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610e6e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610e31565b610e788683610e31565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610ebf610eba610eb584610e90565b610e9a565b610e90565b9050919050565b6000819050919050565b610ed983610ea4565b610eed610ee582610ec6565b848454610e3e565b825550505050565b600090565b610f02610ef5565b610f0d818484610ed0565b505050565b5b81811015610f3157610f26600082610efa565b600181019050610f13565b5050565b601f821115610f7657610f4781610e0c565b610f5084610e21565b81016020851015610f5f578190505b610f73610f6b85610e21565b830182610f12565b50505b505050565b600082821c905092915050565b6000610f9960001984600802610f7b565b1980831691505092915050565b6000610fb28383610f88565b9150826002028217905092915050565b610fcb82610aa8565b67ffffffffffffffff811115610fe457610fe361079c565b5b610fee8254610d2c565b610ff9828285610f35565b600060209050601f83116001811461102c576000841561101a578287015190505b6110248582610fa6565b86555061108c565b601f19841661103a86610e0c565b60005b828110156110625784890151825560018201915060208501945060208101905061103d565b8683101561107f578489015161107b601f891682610f88565b8355505b6001600288020188555050505b505050505050565b601f8211156110d5576110a681610d5d565b6110af84610e21565b810160208510156110be578190505b6110d26110ca85610e21565b830182610f12565b50505b505050565b6110e382610c9f565b67ffffffffffffffff8111156110fc576110fb61079c565b5b6111068254610d2c565b611111828285611094565b600060209050601f8311600181146111445760008415611132578287015190505b61113c8582610fa6565b8655506111a4565b601f19841661115286610d5d565b60005b8281101561117a57848901518255600182019150602085019450602081019050611155565b868310156111975784890151611193601f891682610f88565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220b9788fd6725e3417326555ab9031274eb305dd338cc29fa0b4f4f471f5d972de64736f6c634300081e0033",
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

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(bytes)
func (_Contracts *ContractsCaller) GetNonce(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(bytes)
func (_Contracts *ContractsSession) GetNonce() ([]byte, error) {
	return _Contracts.Contract.GetNonce(&_Contracts.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(bytes)
func (_Contracts *ContractsCallerSession) GetNonce() ([]byte, error) {
	return _Contracts.Contract.GetNonce(&_Contracts.CallOpts)
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

// SetNonce is a paid mutator transaction binding the contract method 0x7bebc55b.
//
// Solidity: function setNonce(bytes input) returns()
func (_Contracts *ContractsTransactor) SetNonce(opts *bind.TransactOpts, input []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setNonce", input)
}

// SetNonce is a paid mutator transaction binding the contract method 0x7bebc55b.
//
// Solidity: function setNonce(bytes input) returns()
func (_Contracts *ContractsSession) SetNonce(input []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetNonce(&_Contracts.TransactOpts, input)
}

// SetNonce is a paid mutator transaction binding the contract method 0x7bebc55b.
//
// Solidity: function setNonce(bytes input) returns()
func (_Contracts *ContractsTransactorSession) SetNonce(input []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SetNonce(&_Contracts.TransactOpts, input)
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
