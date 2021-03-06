// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rc

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

// RevokedCertABI is the input ABI used to generate the binding from.
const RevokedCertABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"extensions\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"fin\",\"type\":\"bool\"}],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revocationTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serialNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"serial\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"setCertData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setExtensions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RevokedCertFuncSigs maps the 4-byte function signature to its string representation.
var RevokedCertFuncSigs = map[string]string{
	"db85d59c": "extensions(uint256)",
	"6c9789b0": "finalize(bool)",
	"b3f05b97": "finalized()",
	"8d4e4083": "isFinalized()",
	"af8ee94a": "revocationTime()",
	"2a0daeea": "serialNumber()",
	"211bd912": "setCertData(uint256,uint256)",
	"2fc02f22": "setExtensions(uint8[])",
}

// RevokedCertBin is the compiled bytecode used for deploying new contracts.
var RevokedCertBin = "0x608060405234801561001057600080fd5b506103eb806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638d4e40831161005b5780638d4e40831461018e578063af8ee94a146101aa578063b3f05b97146101b2578063db85d59c146101ba57610088565b8063211bd9121461008d5780632a0daeea146100b25780632fc02f22146100cc5780636c9789b01461016f575b600080fd5b6100b0600480360360408110156100a357600080fd5b50803590602001356101ed565b005b6100ba61023a565b60408051918252519081900360200190f35b6100b0600480360360208110156100e257600080fd5b8101906020810181356401000000008111156100fd57600080fd5b82018360208201111561010f57600080fd5b8035906020019184602083028401116401000000008311171561013157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610240945050505050565b6100b06004803603602081101561018557600080fd5b503515156102ea565b61019661033f565b604080519115158252519081900360200190f35b6100ba610348565b61019661034e565b6101d7600480360360208110156101d057600080fd5b5035610357565b6040805160ff9092168252519081900360200190f35b60005460ff161561022f5760405162461bcd60e51b815260040180806020018281038252602a81526020018061038c602a913960400191505060405180910390fd5b600191909155600255565b60015481565b60005460ff16156102825760405162461bcd60e51b815260040180806020018281038252602a81526020018061038c602a913960400191505060405180910390fd5b805160005b818110156102e557600383828151811061029d57fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a938402929093021990921617905501610287565b505050565b60005460ff161561032c5760405162461bcd60e51b815260040180806020018281038252602a81526020018061038c602a913960400191505060405180910390fd5b6000805460ff1916911515919091179055565b60005460ff1690565b60025481565b60005460ff1681565b6003818154811061036757600080fd5b9060005260206000209060209182820401919006915054906101000a900460ff168156fe46696e616c697a61626c653a20636f6e747261637420697320616c72656164792066696e616c697a6564a2646970667358221220c131cbc19a993f161f1e27507027522ae8ca4b87f2d8b223f1ecdfe9bc3f502664736f6c63430007040033"

// DeployRevokedCert deploys a new Ethereum contract, binding an instance of RevokedCert to it.
func DeployRevokedCert(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RevokedCert, error) {
	parsed, err := abi.JSON(strings.NewReader(RevokedCertABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RevokedCertBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RevokedCert{RevokedCertCaller: RevokedCertCaller{contract: contract}, RevokedCertTransactor: RevokedCertTransactor{contract: contract}, RevokedCertFilterer: RevokedCertFilterer{contract: contract}}, nil
}

// RevokedCert is an auto generated Go binding around an Ethereum contract.
type RevokedCert struct {
	RevokedCertCaller     // Read-only binding to the contract
	RevokedCertTransactor // Write-only binding to the contract
	RevokedCertFilterer   // Log filterer for contract events
}

// RevokedCertCaller is an auto generated read-only Go binding around an Ethereum contract.
type RevokedCertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevokedCertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RevokedCertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevokedCertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RevokedCertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RevokedCertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RevokedCertSession struct {
	Contract     *RevokedCert      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RevokedCertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RevokedCertCallerSession struct {
	Contract *RevokedCertCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RevokedCertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RevokedCertTransactorSession struct {
	Contract     *RevokedCertTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RevokedCertRaw is an auto generated low-level Go binding around an Ethereum contract.
type RevokedCertRaw struct {
	Contract *RevokedCert // Generic contract binding to access the raw methods on
}

// RevokedCertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RevokedCertCallerRaw struct {
	Contract *RevokedCertCaller // Generic read-only contract binding to access the raw methods on
}

// RevokedCertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RevokedCertTransactorRaw struct {
	Contract *RevokedCertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRevokedCert creates a new instance of RevokedCert, bound to a specific deployed contract.
func NewRevokedCert(address common.Address, backend bind.ContractBackend) (*RevokedCert, error) {
	contract, err := bindRevokedCert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RevokedCert{RevokedCertCaller: RevokedCertCaller{contract: contract}, RevokedCertTransactor: RevokedCertTransactor{contract: contract}, RevokedCertFilterer: RevokedCertFilterer{contract: contract}}, nil
}

// NewRevokedCertCaller creates a new read-only instance of RevokedCert, bound to a specific deployed contract.
func NewRevokedCertCaller(address common.Address, caller bind.ContractCaller) (*RevokedCertCaller, error) {
	contract, err := bindRevokedCert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RevokedCertCaller{contract: contract}, nil
}

// NewRevokedCertTransactor creates a new write-only instance of RevokedCert, bound to a specific deployed contract.
func NewRevokedCertTransactor(address common.Address, transactor bind.ContractTransactor) (*RevokedCertTransactor, error) {
	contract, err := bindRevokedCert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RevokedCertTransactor{contract: contract}, nil
}

// NewRevokedCertFilterer creates a new log filterer instance of RevokedCert, bound to a specific deployed contract.
func NewRevokedCertFilterer(address common.Address, filterer bind.ContractFilterer) (*RevokedCertFilterer, error) {
	contract, err := bindRevokedCert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RevokedCertFilterer{contract: contract}, nil
}

// bindRevokedCert binds a generic wrapper to an already deployed contract.
func bindRevokedCert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RevokedCertABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevokedCert *RevokedCertRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevokedCert.Contract.RevokedCertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevokedCert *RevokedCertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevokedCert.Contract.RevokedCertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevokedCert *RevokedCertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevokedCert.Contract.RevokedCertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RevokedCert *RevokedCertCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RevokedCert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RevokedCert *RevokedCertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RevokedCert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RevokedCert *RevokedCertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RevokedCert.Contract.contract.Transact(opts, method, params...)
}

// Extensions is a free data retrieval call binding the contract method 0xdb85d59c.
//
// Solidity: function extensions(uint256 ) view returns(uint8)
func (_RevokedCert *RevokedCertCaller) Extensions(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _RevokedCert.contract.Call(opts, &out, "extensions", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Extensions is a free data retrieval call binding the contract method 0xdb85d59c.
//
// Solidity: function extensions(uint256 ) view returns(uint8)
func (_RevokedCert *RevokedCertSession) Extensions(arg0 *big.Int) (uint8, error) {
	return _RevokedCert.Contract.Extensions(&_RevokedCert.CallOpts, arg0)
}

// Extensions is a free data retrieval call binding the contract method 0xdb85d59c.
//
// Solidity: function extensions(uint256 ) view returns(uint8)
func (_RevokedCert *RevokedCertCallerSession) Extensions(arg0 *big.Int) (uint8, error) {
	return _RevokedCert.Contract.Extensions(&_RevokedCert.CallOpts, arg0)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RevokedCert *RevokedCertCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RevokedCert.contract.Call(opts, &out, "finalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RevokedCert *RevokedCertSession) Finalized() (bool, error) {
	return _RevokedCert.Contract.Finalized(&_RevokedCert.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_RevokedCert *RevokedCertCallerSession) Finalized() (bool, error) {
	return _RevokedCert.Contract.Finalized(&_RevokedCert.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_RevokedCert *RevokedCertCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RevokedCert.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_RevokedCert *RevokedCertSession) IsFinalized() (bool, error) {
	return _RevokedCert.Contract.IsFinalized(&_RevokedCert.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_RevokedCert *RevokedCertCallerSession) IsFinalized() (bool, error) {
	return _RevokedCert.Contract.IsFinalized(&_RevokedCert.CallOpts)
}

// RevocationTime is a free data retrieval call binding the contract method 0xaf8ee94a.
//
// Solidity: function revocationTime() view returns(uint256)
func (_RevokedCert *RevokedCertCaller) RevocationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RevokedCert.contract.Call(opts, &out, "revocationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RevocationTime is a free data retrieval call binding the contract method 0xaf8ee94a.
//
// Solidity: function revocationTime() view returns(uint256)
func (_RevokedCert *RevokedCertSession) RevocationTime() (*big.Int, error) {
	return _RevokedCert.Contract.RevocationTime(&_RevokedCert.CallOpts)
}

// RevocationTime is a free data retrieval call binding the contract method 0xaf8ee94a.
//
// Solidity: function revocationTime() view returns(uint256)
func (_RevokedCert *RevokedCertCallerSession) RevocationTime() (*big.Int, error) {
	return _RevokedCert.Contract.RevocationTime(&_RevokedCert.CallOpts)
}

// SerialNumber is a free data retrieval call binding the contract method 0x2a0daeea.
//
// Solidity: function serialNumber() view returns(uint256)
func (_RevokedCert *RevokedCertCaller) SerialNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RevokedCert.contract.Call(opts, &out, "serialNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SerialNumber is a free data retrieval call binding the contract method 0x2a0daeea.
//
// Solidity: function serialNumber() view returns(uint256)
func (_RevokedCert *RevokedCertSession) SerialNumber() (*big.Int, error) {
	return _RevokedCert.Contract.SerialNumber(&_RevokedCert.CallOpts)
}

// SerialNumber is a free data retrieval call binding the contract method 0x2a0daeea.
//
// Solidity: function serialNumber() view returns(uint256)
func (_RevokedCert *RevokedCertCallerSession) SerialNumber() (*big.Int, error) {
	return _RevokedCert.Contract.SerialNumber(&_RevokedCert.CallOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_RevokedCert *RevokedCertTransactor) Finalize(opts *bind.TransactOpts, fin bool) (*types.Transaction, error) {
	return _RevokedCert.contract.Transact(opts, "finalize", fin)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_RevokedCert *RevokedCertSession) Finalize(fin bool) (*types.Transaction, error) {
	return _RevokedCert.Contract.Finalize(&_RevokedCert.TransactOpts, fin)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_RevokedCert *RevokedCertTransactorSession) Finalize(fin bool) (*types.Transaction, error) {
	return _RevokedCert.Contract.Finalize(&_RevokedCert.TransactOpts, fin)
}

// SetCertData is a paid mutator transaction binding the contract method 0x211bd912.
//
// Solidity: function setCertData(uint256 serial, uint256 time) returns()
func (_RevokedCert *RevokedCertTransactor) SetCertData(opts *bind.TransactOpts, serial *big.Int, time *big.Int) (*types.Transaction, error) {
	return _RevokedCert.contract.Transact(opts, "setCertData", serial, time)
}

// SetCertData is a paid mutator transaction binding the contract method 0x211bd912.
//
// Solidity: function setCertData(uint256 serial, uint256 time) returns()
func (_RevokedCert *RevokedCertSession) SetCertData(serial *big.Int, time *big.Int) (*types.Transaction, error) {
	return _RevokedCert.Contract.SetCertData(&_RevokedCert.TransactOpts, serial, time)
}

// SetCertData is a paid mutator transaction binding the contract method 0x211bd912.
//
// Solidity: function setCertData(uint256 serial, uint256 time) returns()
func (_RevokedCert *RevokedCertTransactorSession) SetCertData(serial *big.Int, time *big.Int) (*types.Transaction, error) {
	return _RevokedCert.Contract.SetCertData(&_RevokedCert.TransactOpts, serial, time)
}

// SetExtensions is a paid mutator transaction binding the contract method 0x2fc02f22.
//
// Solidity: function setExtensions(uint8[] data) returns()
func (_RevokedCert *RevokedCertTransactor) SetExtensions(opts *bind.TransactOpts, data []uint8) (*types.Transaction, error) {
	return _RevokedCert.contract.Transact(opts, "setExtensions", data)
}

// SetExtensions is a paid mutator transaction binding the contract method 0x2fc02f22.
//
// Solidity: function setExtensions(uint8[] data) returns()
func (_RevokedCert *RevokedCertSession) SetExtensions(data []uint8) (*types.Transaction, error) {
	return _RevokedCert.Contract.SetExtensions(&_RevokedCert.TransactOpts, data)
}

// SetExtensions is a paid mutator transaction binding the contract method 0x2fc02f22.
//
// Solidity: function setExtensions(uint8[] data) returns()
func (_RevokedCert *RevokedCertTransactorSession) SetExtensions(data []uint8) (*types.Transaction, error) {
	return _RevokedCert.Contract.SetExtensions(&_RevokedCert.TransactOpts, data)
}
