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

// MulticallMetaData contains all meta data concerning the Multicall contract.
var MulticallMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"contracts\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"bytecodes\",\"type\":\"bytes[]\"}],\"name\":\"aggregate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"contracts\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"bytecodes\",\"type\":\"bytes[]\"}],\"name\":\"batch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"result\",\"type\":\"bool[]\"},{\"internalType\":\"bytes[]\",\"name\":\"returnData\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MulticallABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallMetaData.ABI instead.
var MulticallABI = MulticallMetaData.ABI

// Multicall is an auto generated Go binding around an Ethereum contract.
type Multicall struct {
	MulticallCaller     // Read-only binding to the contract
	MulticallTransactor // Write-only binding to the contract
	MulticallFilterer   // Log filterer for contract events
}

// MulticallCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallSession struct {
	Contract     *Multicall        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallCallerSession struct {
	Contract *MulticallCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MulticallTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallTransactorSession struct {
	Contract     *MulticallTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MulticallRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallRaw struct {
	Contract *Multicall // Generic contract binding to access the raw methods on
}

// MulticallCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallCallerRaw struct {
	Contract *MulticallCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallTransactorRaw struct {
	Contract *MulticallTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticall creates a new instance of Multicall, bound to a specific deployed contract.
func NewMulticall(address common.Address, backend bind.ContractBackend) (*Multicall, error) {
	contract, err := bindMulticall(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multicall{MulticallCaller: MulticallCaller{contract: contract}, MulticallTransactor: MulticallTransactor{contract: contract}, MulticallFilterer: MulticallFilterer{contract: contract}}, nil
}

// NewMulticallCaller creates a new read-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallCaller(address common.Address, caller bind.ContractCaller) (*MulticallCaller, error) {
	contract, err := bindMulticall(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallCaller{contract: contract}, nil
}

// NewMulticallTransactor creates a new write-only instance of Multicall, bound to a specific deployed contract.
func NewMulticallTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallTransactor, error) {
	contract, err := bindMulticall(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallTransactor{contract: contract}, nil
}

// NewMulticallFilterer creates a new log filterer instance of Multicall, bound to a specific deployed contract.
func NewMulticallFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallFilterer, error) {
	contract, err := bindMulticall(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallFilterer{contract: contract}, nil
}

// bindMulticall binds a generic wrapper to an already deployed contract.
func bindMulticall(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MulticallMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.MulticallCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.MulticallTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multicall *MulticallCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multicall.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multicall *MulticallTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multicall *MulticallTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multicall.Contract.contract.Transact(opts, method, params...)
}

// Aggregate is a free data retrieval call binding the contract method 0x269d64ae.
//
// Solidity: function aggregate(address[] contracts, bytes[] bytecodes) view returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall *MulticallCaller) Aggregate(opts *bind.CallOpts, contracts []common.Address, bytecodes [][]byte) (struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "aggregate", contracts, bytecodes)

	outstruct := new(struct {
		BlockNumber *big.Int
		ReturnData  [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReturnData = *abi.ConvertType(out[1], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// Aggregate is a free data retrieval call binding the contract method 0x269d64ae.
//
// Solidity: function aggregate(address[] contracts, bytes[] bytecodes) view returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall *MulticallSession) Aggregate(contracts []common.Address, bytecodes [][]byte) (struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}, error) {
	return _Multicall.Contract.Aggregate(&_Multicall.CallOpts, contracts, bytecodes)
}

// Aggregate is a free data retrieval call binding the contract method 0x269d64ae.
//
// Solidity: function aggregate(address[] contracts, bytes[] bytecodes) view returns(uint256 blockNumber, bytes[] returnData)
func (_Multicall *MulticallCallerSession) Aggregate(contracts []common.Address, bytecodes [][]byte) (struct {
	BlockNumber *big.Int
	ReturnData  [][]byte
}, error) {
	return _Multicall.Contract.Aggregate(&_Multicall.CallOpts, contracts, bytecodes)
}

// Batch is a free data retrieval call binding the contract method 0xf38f59d7.
//
// Solidity: function batch(address[] contracts, bytes[] bytecodes) view returns(uint256 blockNumber, bool[] result, bytes[] returnData)
func (_Multicall *MulticallCaller) Batch(opts *bind.CallOpts, contracts []common.Address, bytecodes [][]byte) (struct {
	BlockNumber *big.Int
	Result      []bool
	ReturnData  [][]byte
}, error) {
	var out []interface{}
	err := _Multicall.contract.Call(opts, &out, "batch", contracts, bytecodes)

	outstruct := new(struct {
		BlockNumber *big.Int
		Result      []bool
		ReturnData  [][]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlockNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Result = *abi.ConvertType(out[1], new([]bool)).(*[]bool)
	outstruct.ReturnData = *abi.ConvertType(out[2], new([][]byte)).(*[][]byte)

	return *outstruct, err

}

// Batch is a free data retrieval call binding the contract method 0xf38f59d7.
//
// Solidity: function batch(address[] contracts, bytes[] bytecodes) view returns(uint256 blockNumber, bool[] result, bytes[] returnData)
func (_Multicall *MulticallSession) Batch(contracts []common.Address, bytecodes [][]byte) (struct {
	BlockNumber *big.Int
	Result      []bool
	ReturnData  [][]byte
}, error) {
	return _Multicall.Contract.Batch(&_Multicall.CallOpts, contracts, bytecodes)
}

// Batch is a free data retrieval call binding the contract method 0xf38f59d7.
//
// Solidity: function batch(address[] contracts, bytes[] bytecodes) view returns(uint256 blockNumber, bool[] result, bytes[] returnData)
func (_Multicall *MulticallCallerSession) Batch(contracts []common.Address, bytecodes [][]byte) (struct {
	BlockNumber *big.Int
	Result      []bool
	ReturnData  [][]byte
}, error) {
	return _Multicall.Contract.Batch(&_Multicall.CallOpts, contracts, bytecodes)
}
