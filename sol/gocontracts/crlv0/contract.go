// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package crlv0

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

// CRLv0ABI is the input ABI used to generate the binding from.
const CRLv0ABI = "[{\"inputs\":[],\"name\":\"finalised\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"fin\",\"type\":\"bool\"}],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setSigAlgo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setSigValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ref\",\"type\":\"address\"}],\"name\":\"setTBSCertList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sigAlgo\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sigValue\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tbsCertList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CRLv0FuncSigs maps the 4-byte function signature to its string representation.
var CRLv0FuncSigs = map[string]string{
	"214bb60f": "finalised()",
	"6c9789b0": "finalize(bool)",
	"b3f05b97": "finalized()",
	"8d4e4083": "isFinalized()",
	"017af69e": "setSigAlgo(uint8[])",
	"3a399a74": "setSigValue(uint8[])",
	"5cf2e1d8": "setTBSCertList(address)",
	"6ab5e375": "sigAlgo(uint256)",
	"75147665": "sigValue(uint256)",
	"3fd09fed": "tbsCertList()",
}

// CRLv0Bin is the compiled bytecode used for deploying new contracts.
var CRLv0Bin = "0x608060405234801561001057600080fd5b506105d2806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80636ab5e375116100665780636ab5e375146102515780636c9789b01461028457806375147665146102a35780638d4e4083146102c0578063b3f05b97146102c85761009e565b8063017af69e146100a3578063214bb60f146101485780633a399a74146101645780633fd09fed146102075780635cf2e1d81461022b575b600080fd5b610146600480360360208110156100b957600080fd5b8101906020810181356401000000008111156100d457600080fd5b8201836020820111156100e657600080fd5b8035906020019184602083028401116401000000008311171561010857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506102d0945050505050565b005b61015061037a565b604080519115158252519081900360200190f35b6101466004803603602081101561017a57600080fd5b81019060208101813564010000000081111561019557600080fd5b8201836020820111156101a757600080fd5b803590602001918460208302840111640100000000831117156101c957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610388945050505050565b61020f61042d565b604080516001600160a01b039092168252519081900360200190f35b6101466004803603602081101561024157600080fd5b50356001600160a01b0316610442565b61026e6004803603602081101561026757600080fd5b50356104c7565b6040805160ff9092168252519081900360200190f35b6101466004803603602081101561029a57600080fd5b503515156104fb565b61026e600480360360208110156102b957600080fd5b5035610550565b610150610560565b610150610569565b60005460ff16156103125760405162461bcd60e51b815260040180806020018281038252602a815260200180610573602a913960400191505060405180910390fd5b805160005b8181101561037557600183828151811061032d57fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a938402929093021990921617905501610317565b505050565b600054610100900460ff1681565b60005460ff16156103ca5760405162461bcd60e51b815260040180806020018281038252602a815260200180610573602a913960400191505060405180910390fd5b805160005b818110156103755760028382815181106103e557fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a9384029290930219909216179055016103cf565b6000546201000090046001600160a01b031681565b60005460ff16156104845760405162461bcd60e51b815260040180806020018281038252602a815260200180610573602a913960400191505060405180910390fd5b6000546201000090046001600160a01b03161561049d57fe5b600080546001600160a01b03909216620100000262010000600160b01b0319909216919091179055565b600181815481106104d757600080fd5b9060005260206000209060209182820401919006915054906101000a900460ff1681565b60005460ff161561053d5760405162461bcd60e51b815260040180806020018281038252602a815260200180610573602a913960400191505060405180910390fd5b6000805460ff1916911515919091179055565b600281815481106104d757600080fd5b60005460ff1690565b60005460ff168156fe46696e616c697a61626c653a20636f6e747261637420697320616c72656164792066696e616c697a6564a2646970667358221220ae8b963d6f3f3eabe77e11688b8e5956bdc3b5aa93abc4580c63bffd0d0e93c164736f6c63430007040033"

// DeployCRLv0 deploys a new Ethereum contract, binding an instance of CRLv0 to it.
func DeployCRLv0(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CRLv0, error) {
	parsed, err := abi.JSON(strings.NewReader(CRLv0ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CRLv0Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CRLv0{CRLv0Caller: CRLv0Caller{contract: contract}, CRLv0Transactor: CRLv0Transactor{contract: contract}, CRLv0Filterer: CRLv0Filterer{contract: contract}}, nil
}

// CRLv0 is an auto generated Go binding around an Ethereum contract.
type CRLv0 struct {
	CRLv0Caller     // Read-only binding to the contract
	CRLv0Transactor // Write-only binding to the contract
	CRLv0Filterer   // Log filterer for contract events
}

// CRLv0Caller is an auto generated read-only Go binding around an Ethereum contract.
type CRLv0Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CRLv0Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CRLv0Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CRLv0Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CRLv0Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CRLv0Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CRLv0Session struct {
	Contract     *CRLv0            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CRLv0CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CRLv0CallerSession struct {
	Contract *CRLv0Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CRLv0TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CRLv0TransactorSession struct {
	Contract     *CRLv0Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CRLv0Raw is an auto generated low-level Go binding around an Ethereum contract.
type CRLv0Raw struct {
	Contract *CRLv0 // Generic contract binding to access the raw methods on
}

// CRLv0CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CRLv0CallerRaw struct {
	Contract *CRLv0Caller // Generic read-only contract binding to access the raw methods on
}

// CRLv0TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CRLv0TransactorRaw struct {
	Contract *CRLv0Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCRLv0 creates a new instance of CRLv0, bound to a specific deployed contract.
func NewCRLv0(address common.Address, backend bind.ContractBackend) (*CRLv0, error) {
	contract, err := bindCRLv0(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CRLv0{CRLv0Caller: CRLv0Caller{contract: contract}, CRLv0Transactor: CRLv0Transactor{contract: contract}, CRLv0Filterer: CRLv0Filterer{contract: contract}}, nil
}

// NewCRLv0Caller creates a new read-only instance of CRLv0, bound to a specific deployed contract.
func NewCRLv0Caller(address common.Address, caller bind.ContractCaller) (*CRLv0Caller, error) {
	contract, err := bindCRLv0(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CRLv0Caller{contract: contract}, nil
}

// NewCRLv0Transactor creates a new write-only instance of CRLv0, bound to a specific deployed contract.
func NewCRLv0Transactor(address common.Address, transactor bind.ContractTransactor) (*CRLv0Transactor, error) {
	contract, err := bindCRLv0(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CRLv0Transactor{contract: contract}, nil
}

// NewCRLv0Filterer creates a new log filterer instance of CRLv0, bound to a specific deployed contract.
func NewCRLv0Filterer(address common.Address, filterer bind.ContractFilterer) (*CRLv0Filterer, error) {
	contract, err := bindCRLv0(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CRLv0Filterer{contract: contract}, nil
}

// bindCRLv0 binds a generic wrapper to an already deployed contract.
func bindCRLv0(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CRLv0ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CRLv0 *CRLv0Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CRLv0.Contract.CRLv0Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CRLv0 *CRLv0Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CRLv0.Contract.CRLv0Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CRLv0 *CRLv0Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CRLv0.Contract.CRLv0Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CRLv0 *CRLv0CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CRLv0.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CRLv0 *CRLv0TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CRLv0.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CRLv0 *CRLv0TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CRLv0.Contract.contract.Transact(opts, method, params...)
}

// Finalised is a free data retrieval call binding the contract method 0x214bb60f.
//
// Solidity: function finalised() view returns(bool)
func (_CRLv0 *CRLv0Caller) Finalised(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CRLv0.contract.Call(opts, &out, "finalised")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalised is a free data retrieval call binding the contract method 0x214bb60f.
//
// Solidity: function finalised() view returns(bool)
func (_CRLv0 *CRLv0Session) Finalised() (bool, error) {
	return _CRLv0.Contract.Finalised(&_CRLv0.CallOpts)
}

// Finalised is a free data retrieval call binding the contract method 0x214bb60f.
//
// Solidity: function finalised() view returns(bool)
func (_CRLv0 *CRLv0CallerSession) Finalised() (bool, error) {
	return _CRLv0.Contract.Finalised(&_CRLv0.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_CRLv0 *CRLv0Caller) Finalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CRLv0.contract.Call(opts, &out, "finalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_CRLv0 *CRLv0Session) Finalized() (bool, error) {
	return _CRLv0.Contract.Finalized(&_CRLv0.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_CRLv0 *CRLv0CallerSession) Finalized() (bool, error) {
	return _CRLv0.Contract.Finalized(&_CRLv0.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_CRLv0 *CRLv0Caller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CRLv0.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_CRLv0 *CRLv0Session) IsFinalized() (bool, error) {
	return _CRLv0.Contract.IsFinalized(&_CRLv0.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_CRLv0 *CRLv0CallerSession) IsFinalized() (bool, error) {
	return _CRLv0.Contract.IsFinalized(&_CRLv0.CallOpts)
}

// SigAlgo is a free data retrieval call binding the contract method 0x6ab5e375.
//
// Solidity: function sigAlgo(uint256 ) view returns(uint8)
func (_CRLv0 *CRLv0Caller) SigAlgo(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _CRLv0.contract.Call(opts, &out, "sigAlgo", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SigAlgo is a free data retrieval call binding the contract method 0x6ab5e375.
//
// Solidity: function sigAlgo(uint256 ) view returns(uint8)
func (_CRLv0 *CRLv0Session) SigAlgo(arg0 *big.Int) (uint8, error) {
	return _CRLv0.Contract.SigAlgo(&_CRLv0.CallOpts, arg0)
}

// SigAlgo is a free data retrieval call binding the contract method 0x6ab5e375.
//
// Solidity: function sigAlgo(uint256 ) view returns(uint8)
func (_CRLv0 *CRLv0CallerSession) SigAlgo(arg0 *big.Int) (uint8, error) {
	return _CRLv0.Contract.SigAlgo(&_CRLv0.CallOpts, arg0)
}

// SigValue is a free data retrieval call binding the contract method 0x75147665.
//
// Solidity: function sigValue(uint256 ) view returns(uint8)
func (_CRLv0 *CRLv0Caller) SigValue(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _CRLv0.contract.Call(opts, &out, "sigValue", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SigValue is a free data retrieval call binding the contract method 0x75147665.
//
// Solidity: function sigValue(uint256 ) view returns(uint8)
func (_CRLv0 *CRLv0Session) SigValue(arg0 *big.Int) (uint8, error) {
	return _CRLv0.Contract.SigValue(&_CRLv0.CallOpts, arg0)
}

// SigValue is a free data retrieval call binding the contract method 0x75147665.
//
// Solidity: function sigValue(uint256 ) view returns(uint8)
func (_CRLv0 *CRLv0CallerSession) SigValue(arg0 *big.Int) (uint8, error) {
	return _CRLv0.Contract.SigValue(&_CRLv0.CallOpts, arg0)
}

// TbsCertList is a free data retrieval call binding the contract method 0x3fd09fed.
//
// Solidity: function tbsCertList() view returns(address)
func (_CRLv0 *CRLv0Caller) TbsCertList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CRLv0.contract.Call(opts, &out, "tbsCertList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TbsCertList is a free data retrieval call binding the contract method 0x3fd09fed.
//
// Solidity: function tbsCertList() view returns(address)
func (_CRLv0 *CRLv0Session) TbsCertList() (common.Address, error) {
	return _CRLv0.Contract.TbsCertList(&_CRLv0.CallOpts)
}

// TbsCertList is a free data retrieval call binding the contract method 0x3fd09fed.
//
// Solidity: function tbsCertList() view returns(address)
func (_CRLv0 *CRLv0CallerSession) TbsCertList() (common.Address, error) {
	return _CRLv0.Contract.TbsCertList(&_CRLv0.CallOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_CRLv0 *CRLv0Transactor) Finalize(opts *bind.TransactOpts, fin bool) (*types.Transaction, error) {
	return _CRLv0.contract.Transact(opts, "finalize", fin)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_CRLv0 *CRLv0Session) Finalize(fin bool) (*types.Transaction, error) {
	return _CRLv0.Contract.Finalize(&_CRLv0.TransactOpts, fin)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_CRLv0 *CRLv0TransactorSession) Finalize(fin bool) (*types.Transaction, error) {
	return _CRLv0.Contract.Finalize(&_CRLv0.TransactOpts, fin)
}

// SetSigAlgo is a paid mutator transaction binding the contract method 0x017af69e.
//
// Solidity: function setSigAlgo(uint8[] data) returns()
func (_CRLv0 *CRLv0Transactor) SetSigAlgo(opts *bind.TransactOpts, data []uint8) (*types.Transaction, error) {
	return _CRLv0.contract.Transact(opts, "setSigAlgo", data)
}

// SetSigAlgo is a paid mutator transaction binding the contract method 0x017af69e.
//
// Solidity: function setSigAlgo(uint8[] data) returns()
func (_CRLv0 *CRLv0Session) SetSigAlgo(data []uint8) (*types.Transaction, error) {
	return _CRLv0.Contract.SetSigAlgo(&_CRLv0.TransactOpts, data)
}

// SetSigAlgo is a paid mutator transaction binding the contract method 0x017af69e.
//
// Solidity: function setSigAlgo(uint8[] data) returns()
func (_CRLv0 *CRLv0TransactorSession) SetSigAlgo(data []uint8) (*types.Transaction, error) {
	return _CRLv0.Contract.SetSigAlgo(&_CRLv0.TransactOpts, data)
}

// SetSigValue is a paid mutator transaction binding the contract method 0x3a399a74.
//
// Solidity: function setSigValue(uint8[] data) returns()
func (_CRLv0 *CRLv0Transactor) SetSigValue(opts *bind.TransactOpts, data []uint8) (*types.Transaction, error) {
	return _CRLv0.contract.Transact(opts, "setSigValue", data)
}

// SetSigValue is a paid mutator transaction binding the contract method 0x3a399a74.
//
// Solidity: function setSigValue(uint8[] data) returns()
func (_CRLv0 *CRLv0Session) SetSigValue(data []uint8) (*types.Transaction, error) {
	return _CRLv0.Contract.SetSigValue(&_CRLv0.TransactOpts, data)
}

// SetSigValue is a paid mutator transaction binding the contract method 0x3a399a74.
//
// Solidity: function setSigValue(uint8[] data) returns()
func (_CRLv0 *CRLv0TransactorSession) SetSigValue(data []uint8) (*types.Transaction, error) {
	return _CRLv0.Contract.SetSigValue(&_CRLv0.TransactOpts, data)
}

// SetTBSCertList is a paid mutator transaction binding the contract method 0x5cf2e1d8.
//
// Solidity: function setTBSCertList(address ref) returns()
func (_CRLv0 *CRLv0Transactor) SetTBSCertList(opts *bind.TransactOpts, ref common.Address) (*types.Transaction, error) {
	return _CRLv0.contract.Transact(opts, "setTBSCertList", ref)
}

// SetTBSCertList is a paid mutator transaction binding the contract method 0x5cf2e1d8.
//
// Solidity: function setTBSCertList(address ref) returns()
func (_CRLv0 *CRLv0Session) SetTBSCertList(ref common.Address) (*types.Transaction, error) {
	return _CRLv0.Contract.SetTBSCertList(&_CRLv0.TransactOpts, ref)
}

// SetTBSCertList is a paid mutator transaction binding the contract method 0x5cf2e1d8.
//
// Solidity: function setTBSCertList(address ref) returns()
func (_CRLv0 *CRLv0TransactorSession) SetTBSCertList(ref common.Address) (*types.Transaction, error) {
	return _CRLv0.Contract.SetTBSCertList(&_CRLv0.TransactOpts, ref)
}

// FinalizableABI is the input ABI used to generate the binding from.
const FinalizableABI = "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"fin\",\"type\":\"bool\"}],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// FinalizableFuncSigs maps the 4-byte function signature to its string representation.
var FinalizableFuncSigs = map[string]string{
	"6c9789b0": "finalize(bool)",
	"b3f05b97": "finalized()",
	"8d4e4083": "isFinalized()",
}

// FinalizableBin is the compiled bytecode used for deploying new contracts.
var FinalizableBin = "0x608060405234801561001057600080fd5b50610152806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80636c9789b0146100465780638d4e408314610067578063b3f05b9714610083575b600080fd5b6100656004803603602081101561005c57600080fd5b5035151561008b565b005b61006f6100e0565b604080519115158252519081900360200190f35b61006f6100e9565b60005460ff16156100cd5760405162461bcd60e51b815260040180806020018281038252602a8152602001806100f3602a913960400191505060405180910390fd5b6000805460ff1916911515919091179055565b60005460ff1690565b60005460ff168156fe46696e616c697a61626c653a20636f6e747261637420697320616c72656164792066696e616c697a6564a2646970667358221220d406032a903ff199ab48db7d5053c56534c5e126effd42477e0495eab40944c564736f6c63430007040033"

// DeployFinalizable deploys a new Ethereum contract, binding an instance of Finalizable to it.
func DeployFinalizable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Finalizable, error) {
	parsed, err := abi.JSON(strings.NewReader(FinalizableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FinalizableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Finalizable{FinalizableCaller: FinalizableCaller{contract: contract}, FinalizableTransactor: FinalizableTransactor{contract: contract}, FinalizableFilterer: FinalizableFilterer{contract: contract}}, nil
}

// Finalizable is an auto generated Go binding around an Ethereum contract.
type Finalizable struct {
	FinalizableCaller     // Read-only binding to the contract
	FinalizableTransactor // Write-only binding to the contract
	FinalizableFilterer   // Log filterer for contract events
}

// FinalizableCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalizableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalizableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalizableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalizableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalizableSession struct {
	Contract     *Finalizable      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FinalizableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalizableCallerSession struct {
	Contract *FinalizableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FinalizableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalizableTransactorSession struct {
	Contract     *FinalizableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FinalizableRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalizableRaw struct {
	Contract *Finalizable // Generic contract binding to access the raw methods on
}

// FinalizableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalizableCallerRaw struct {
	Contract *FinalizableCaller // Generic read-only contract binding to access the raw methods on
}

// FinalizableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalizableTransactorRaw struct {
	Contract *FinalizableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalizable creates a new instance of Finalizable, bound to a specific deployed contract.
func NewFinalizable(address common.Address, backend bind.ContractBackend) (*Finalizable, error) {
	contract, err := bindFinalizable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Finalizable{FinalizableCaller: FinalizableCaller{contract: contract}, FinalizableTransactor: FinalizableTransactor{contract: contract}, FinalizableFilterer: FinalizableFilterer{contract: contract}}, nil
}

// NewFinalizableCaller creates a new read-only instance of Finalizable, bound to a specific deployed contract.
func NewFinalizableCaller(address common.Address, caller bind.ContractCaller) (*FinalizableCaller, error) {
	contract, err := bindFinalizable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableCaller{contract: contract}, nil
}

// NewFinalizableTransactor creates a new write-only instance of Finalizable, bound to a specific deployed contract.
func NewFinalizableTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalizableTransactor, error) {
	contract, err := bindFinalizable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalizableTransactor{contract: contract}, nil
}

// NewFinalizableFilterer creates a new log filterer instance of Finalizable, bound to a specific deployed contract.
func NewFinalizableFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalizableFilterer, error) {
	contract, err := bindFinalizable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalizableFilterer{contract: contract}, nil
}

// bindFinalizable binds a generic wrapper to an already deployed contract.
func bindFinalizable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FinalizableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Finalizable *FinalizableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Finalizable.Contract.FinalizableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Finalizable *FinalizableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Finalizable.Contract.FinalizableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Finalizable *FinalizableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Finalizable.Contract.FinalizableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Finalizable *FinalizableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Finalizable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Finalizable *FinalizableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Finalizable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Finalizable *FinalizableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Finalizable.Contract.contract.Transact(opts, method, params...)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_Finalizable *FinalizableCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Finalizable.contract.Call(opts, &out, "finalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_Finalizable *FinalizableSession) Finalized() (bool, error) {
	return _Finalizable.Contract.Finalized(&_Finalizable.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_Finalizable *FinalizableCallerSession) Finalized() (bool, error) {
	return _Finalizable.Contract.Finalized(&_Finalizable.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Finalizable *FinalizableCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Finalizable.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Finalizable *FinalizableSession) IsFinalized() (bool, error) {
	return _Finalizable.Contract.IsFinalized(&_Finalizable.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Finalizable *FinalizableCallerSession) IsFinalized() (bool, error) {
	return _Finalizable.Contract.IsFinalized(&_Finalizable.CallOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_Finalizable *FinalizableTransactor) Finalize(opts *bind.TransactOpts, fin bool) (*types.Transaction, error) {
	return _Finalizable.contract.Transact(opts, "finalize", fin)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_Finalizable *FinalizableSession) Finalize(fin bool) (*types.Transaction, error) {
	return _Finalizable.Contract.Finalize(&_Finalizable.TransactOpts, fin)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_Finalizable *FinalizableTransactorSession) Finalize(fin bool) (*types.Transaction, error) {
	return _Finalizable.Contract.Finalize(&_Finalizable.TransactOpts, fin)
}
