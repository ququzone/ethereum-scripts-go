// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MultiSenderABI is the input ABI used to generate the binding from.
const MultiSenderABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"coinAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// MultiSender is an auto generated Go binding around an Ethereum contract.
type MultiSender struct {
	MultiSenderCaller     // Read-only binding to the contract
	MultiSenderTransactor // Write-only binding to the contract
	MultiSenderFilterer   // Log filterer for contract events
}

// MultiSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSenderSession struct {
	Contract     *MultiSender      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSenderCallerSession struct {
	Contract *MultiSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MultiSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSenderTransactorSession struct {
	Contract     *MultiSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MultiSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSenderRaw struct {
	Contract *MultiSender // Generic contract binding to access the raw methods on
}

// MultiSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSenderCallerRaw struct {
	Contract *MultiSenderCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSenderTransactorRaw struct {
	Contract *MultiSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSender creates a new instance of MultiSender, bound to a specific deployed contract.
func NewMultiSender(address common.Address, backend bind.ContractBackend) (*MultiSender, error) {
	contract, err := bindMultiSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSender{MultiSenderCaller: MultiSenderCaller{contract: contract}, MultiSenderTransactor: MultiSenderTransactor{contract: contract}, MultiSenderFilterer: MultiSenderFilterer{contract: contract}}, nil
}

// NewMultiSenderCaller creates a new read-only instance of MultiSender, bound to a specific deployed contract.
func NewMultiSenderCaller(address common.Address, caller bind.ContractCaller) (*MultiSenderCaller, error) {
	contract, err := bindMultiSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSenderCaller{contract: contract}, nil
}

// NewMultiSenderTransactor creates a new write-only instance of MultiSender, bound to a specific deployed contract.
func NewMultiSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSenderTransactor, error) {
	contract, err := bindMultiSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSenderTransactor{contract: contract}, nil
}

// NewMultiSenderFilterer creates a new log filterer instance of MultiSender, bound to a specific deployed contract.
func NewMultiSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSenderFilterer, error) {
	contract, err := bindMultiSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSenderFilterer{contract: contract}, nil
}

// bindMultiSender binds a generic wrapper to an already deployed contract.
func bindMultiSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSender *MultiSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSender.Contract.MultiSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSender *MultiSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSender.Contract.MultiSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSender *MultiSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSender.Contract.MultiSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSender *MultiSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSender *MultiSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSender *MultiSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSender.Contract.contract.Transact(opts, method, params...)
}

// Send is a paid mutator transaction binding the contract method 0xc6ed1f74.
//
// Solidity: function send(address[] recipients, address token, uint256 coinAmount, uint256 tokenAmount) payable returns()
func (_MultiSender *MultiSenderTransactor) Send(opts *bind.TransactOpts, recipients []common.Address, token common.Address, coinAmount *big.Int, tokenAmount *big.Int) (*types.Transaction, error) {
	return _MultiSender.contract.Transact(opts, "send", recipients, token, coinAmount, tokenAmount)
}

// Send is a paid mutator transaction binding the contract method 0xc6ed1f74.
//
// Solidity: function send(address[] recipients, address token, uint256 coinAmount, uint256 tokenAmount) payable returns()
func (_MultiSender *MultiSenderSession) Send(recipients []common.Address, token common.Address, coinAmount *big.Int, tokenAmount *big.Int) (*types.Transaction, error) {
	return _MultiSender.Contract.Send(&_MultiSender.TransactOpts, recipients, token, coinAmount, tokenAmount)
}

// Send is a paid mutator transaction binding the contract method 0xc6ed1f74.
//
// Solidity: function send(address[] recipients, address token, uint256 coinAmount, uint256 tokenAmount) payable returns()
func (_MultiSender *MultiSenderTransactorSession) Send(recipients []common.Address, token common.Address, coinAmount *big.Int, tokenAmount *big.Int) (*types.Transaction, error) {
	return _MultiSender.Contract.Send(&_MultiSender.TransactOpts, recipients, token, coinAmount, tokenAmount)
}
