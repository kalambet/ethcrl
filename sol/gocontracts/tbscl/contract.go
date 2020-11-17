// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tbscl

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

// TBSCertListABI is the input ABI used to generate the binding from.
const TBSCertListABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ref\",\"type\":\"address\"}],\"name\":\"addRevokedCert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"extentions\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"fin\",\"type\":\"bool\"}],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raw\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setExtentions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setRaw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"thisUpd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextUpd\",\"type\":\"uint256\"},{\"internalType\":\"int32\",\"name\":\"version\",\"type\":\"int32\"}],\"name\":\"setTimeAndVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signature\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thisUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verison\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TBSCertListFuncSigs maps the 4-byte function signature to its string representation.
var TBSCertListFuncSigs = map[string]string{
	"39c47dc6": "addRevokedCert(address)",
	"8c4420c4": "extentions(uint256)",
	"6c9789b0": "finalize(bool)",
	"b3f05b97": "finalized()",
	"8d4e4083": "isFinalized()",
	"0068fe8b": "issuer(uint256)",
	"6e9c11f9": "nextUpdate()",
	"302e678f": "raw(uint256)",
	"f6cd261c": "setExtentions(uint8[])",
	"0afe3459": "setIssuer(uint8[])",
	"41e83176": "setRaw(uint8[])",
	"0e075af9": "setSignature(uint8[])",
	"1a1d042c": "setTimeAndVersion(uint256,uint256,int32)",
	"70629548": "signature(uint256)",
	"b4729258": "thisUpdate()",
	"38b9f2a8": "verison()",
}

// TBSCertListBin is the compiled bytecode used for deploying new contracts.
var TBSCertListBin = "0x608060405234801561001057600080fd5b50610a16806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ff5760003560e01c80636c9789b0116100975780638d4e4083116100665780638d4e40831461041d578063b3f05b9714610439578063b472925814610441578063f6cd261c14610449576100ff565b80636c9789b0146103aa5780636e9c11f9146103c957806370629548146103e35780638c4420c414610400576100ff565b8063302e678f116100d3578063302e678f146102a757806338b9f2a8146102c457806339c47dc6146102e357806341e8317614610309576100ff565b806268fe8b146101045780630afe3459146101375780630e075af9146101da5780631a1d042c1461027b575b600080fd5b6101216004803603602081101561011a57600080fd5b50356104ea565b6040805160ff9092168252519081900360200190f35b6101d86004803603602081101561014d57600080fd5b810190602081018135600160201b81111561016757600080fd5b82018360208201111561017957600080fd5b803590602001918460208302840111600160201b8311171561019a57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061051e945050505050565b005b6101d8600480360360208110156101f057600080fd5b810190602081018135600160201b81111561020a57600080fd5b82018360208201111561021c57600080fd5b803590602001918460208302840111600160201b8311171561023d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506105c8945050505050565b6101d86004803603606081101561029157600080fd5b508035906020810135906040013560030b61066d565b610121600480360360208110156102bd57600080fd5b50356106bb565b6102cc6106cb565b6040805160039290920b8252519081900360200190f35b6101d8600480360360208110156102f957600080fd5b50356001600160a01b03166106d4565b6101d86004803603602081101561031f57600080fd5b810190602081018135600160201b81111561033957600080fd5b82018360208201111561034b57600080fd5b803590602001918460208302840111600160201b8311171561036c57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506107d9945050505050565b6101d8600480360360208110156103c057600080fd5b5035151561087e565b6103d16108d3565b60408051918252519081900360200190f35b610121600480360360208110156103f957600080fd5b50356108d9565b6101216004803603602081101561041657600080fd5b50356108e9565b6104256108f9565b604080519115158252519081900360200190f35b610425610902565b6103d161090b565b6101d86004803603602081101561045f57600080fd5b810190602081018135600160201b81111561047957600080fd5b82018360208201111561048b57600080fd5b803590602001918460208302840111600160201b831117156104ac57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610911945050505050565b600481815481106104fa57600080fd5b9060005260206000209060209182820401919006915054906101000a900460ff1681565b60005460ff16156105605760405162461bcd60e51b815260040180806020018281038252602a8152602001806109b7602a913960400191505060405180910390fd5b805160005b818110156105c357600483828151811061057b57fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a938402929093021990921617905501610565565b505050565b60005460ff161561060a5760405162461bcd60e51b815260040180806020018281038252602a8152602001806109b7602a913960400191505060405180910390fd5b805160005b818110156105c357600383828151811061062557fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a93840292909302199092161790550161060f565b60005460ff16156106af5760405162461bcd60e51b815260040180806020018281038252602a8152602001806109b7602a913960400191505060405180910390fd5b50600591909155600655565b600181815481106104fa57600080fd5b60025460030b81565b60005460ff16156107165760405162461bcd60e51b815260040180806020018281038252602a8152602001806109b7602a913960400191505060405180910390fd5b6000819050806001600160a01b0316638d4e40836040518163ffffffff1660e01b815260040160206040518083038186803b15801561075457600080fd5b505afa158015610768573d6000803e3d6000fd5b505050506040513d602081101561077e57600080fd5b505161078657fe5b50600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30180546001600160a01b0319166001600160a01b0392909216919091179055565b60005460ff161561081b5760405162461bcd60e51b815260040180806020018281038252602a8152602001806109b7602a913960400191505060405180910390fd5b805160005b818110156105c357600183828151811061083657fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a938402929093021990921617905501610820565b60005460ff16156108c05760405162461bcd60e51b815260040180806020018281038252602a8152602001806109b7602a913960400191505060405180910390fd5b6000805460ff1916911515919091179055565b60065481565b600381815481106104fa57600080fd5b600781815481106104fa57600080fd5b60005460ff1690565b60005460ff1681565b60055481565b60005460ff16156109535760405162461bcd60e51b815260040180806020018281038252602a8152602001806109b7602a913960400191505060405180910390fd5b805160005b818110156105c357600783828151811061096e57fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a93840292909302199092161790550161095856fe46696e616c697a61626c653a20636f6e747261637420697320616c72656164792066696e616c697a6564a26469706673582212204f102bccce5d6ac4a1f297c83e54eba3999269550a8dae42bbb1fbf89e4c5c8564736f6c63430007040033"

// DeployTBSCertList deploys a new Ethereum contract, binding an instance of TBSCertList to it.
func DeployTBSCertList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TBSCertList, error) {
	parsed, err := abi.JSON(strings.NewReader(TBSCertListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TBSCertListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TBSCertList{TBSCertListCaller: TBSCertListCaller{contract: contract}, TBSCertListTransactor: TBSCertListTransactor{contract: contract}, TBSCertListFilterer: TBSCertListFilterer{contract: contract}}, nil
}

// TBSCertList is an auto generated Go binding around an Ethereum contract.
type TBSCertList struct {
	TBSCertListCaller     // Read-only binding to the contract
	TBSCertListTransactor // Write-only binding to the contract
	TBSCertListFilterer   // Log filterer for contract events
}

// TBSCertListCaller is an auto generated read-only Go binding around an Ethereum contract.
type TBSCertListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TBSCertListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TBSCertListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TBSCertListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TBSCertListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TBSCertListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TBSCertListSession struct {
	Contract     *TBSCertList      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TBSCertListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TBSCertListCallerSession struct {
	Contract *TBSCertListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TBSCertListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TBSCertListTransactorSession struct {
	Contract     *TBSCertListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TBSCertListRaw is an auto generated low-level Go binding around an Ethereum contract.
type TBSCertListRaw struct {
	Contract *TBSCertList // Generic contract binding to access the raw methods on
}

// TBSCertListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TBSCertListCallerRaw struct {
	Contract *TBSCertListCaller // Generic read-only contract binding to access the raw methods on
}

// TBSCertListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TBSCertListTransactorRaw struct {
	Contract *TBSCertListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTBSCertList creates a new instance of TBSCertList, bound to a specific deployed contract.
func NewTBSCertList(address common.Address, backend bind.ContractBackend) (*TBSCertList, error) {
	contract, err := bindTBSCertList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TBSCertList{TBSCertListCaller: TBSCertListCaller{contract: contract}, TBSCertListTransactor: TBSCertListTransactor{contract: contract}, TBSCertListFilterer: TBSCertListFilterer{contract: contract}}, nil
}

// NewTBSCertListCaller creates a new read-only instance of TBSCertList, bound to a specific deployed contract.
func NewTBSCertListCaller(address common.Address, caller bind.ContractCaller) (*TBSCertListCaller, error) {
	contract, err := bindTBSCertList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TBSCertListCaller{contract: contract}, nil
}

// NewTBSCertListTransactor creates a new write-only instance of TBSCertList, bound to a specific deployed contract.
func NewTBSCertListTransactor(address common.Address, transactor bind.ContractTransactor) (*TBSCertListTransactor, error) {
	contract, err := bindTBSCertList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TBSCertListTransactor{contract: contract}, nil
}

// NewTBSCertListFilterer creates a new log filterer instance of TBSCertList, bound to a specific deployed contract.
func NewTBSCertListFilterer(address common.Address, filterer bind.ContractFilterer) (*TBSCertListFilterer, error) {
	contract, err := bindTBSCertList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TBSCertListFilterer{contract: contract}, nil
}

// bindTBSCertList binds a generic wrapper to an already deployed contract.
func bindTBSCertList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TBSCertListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TBSCertList *TBSCertListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TBSCertList.Contract.TBSCertListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TBSCertList *TBSCertListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TBSCertList.Contract.TBSCertListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TBSCertList *TBSCertListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TBSCertList.Contract.TBSCertListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TBSCertList *TBSCertListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TBSCertList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TBSCertList *TBSCertListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TBSCertList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TBSCertList *TBSCertListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TBSCertList.Contract.contract.Transact(opts, method, params...)
}

// Extentions is a free data retrieval call binding the contract method 0x8c4420c4.
//
// Solidity: function extentions(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListCaller) Extentions(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "extentions", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Extentions is a free data retrieval call binding the contract method 0x8c4420c4.
//
// Solidity: function extentions(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListSession) Extentions(arg0 *big.Int) (uint8, error) {
	return _TBSCertList.Contract.Extentions(&_TBSCertList.CallOpts, arg0)
}

// Extentions is a free data retrieval call binding the contract method 0x8c4420c4.
//
// Solidity: function extentions(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListCallerSession) Extentions(arg0 *big.Int) (uint8, error) {
	return _TBSCertList.Contract.Extentions(&_TBSCertList.CallOpts, arg0)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_TBSCertList *TBSCertListCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "finalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_TBSCertList *TBSCertListSession) Finalized() (bool, error) {
	return _TBSCertList.Contract.Finalized(&_TBSCertList.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() view returns(bool)
func (_TBSCertList *TBSCertListCallerSession) Finalized() (bool, error) {
	return _TBSCertList.Contract.Finalized(&_TBSCertList.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_TBSCertList *TBSCertListCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_TBSCertList *TBSCertListSession) IsFinalized() (bool, error) {
	return _TBSCertList.Contract.IsFinalized(&_TBSCertList.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_TBSCertList *TBSCertListCallerSession) IsFinalized() (bool, error) {
	return _TBSCertList.Contract.IsFinalized(&_TBSCertList.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x0068fe8b.
//
// Solidity: function issuer(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListCaller) Issuer(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "issuer", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Issuer is a free data retrieval call binding the contract method 0x0068fe8b.
//
// Solidity: function issuer(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListSession) Issuer(arg0 *big.Int) (uint8, error) {
	return _TBSCertList.Contract.Issuer(&_TBSCertList.CallOpts, arg0)
}

// Issuer is a free data retrieval call binding the contract method 0x0068fe8b.
//
// Solidity: function issuer(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListCallerSession) Issuer(arg0 *big.Int) (uint8, error) {
	return _TBSCertList.Contract.Issuer(&_TBSCertList.CallOpts, arg0)
}

// NextUpdate is a free data retrieval call binding the contract method 0x6e9c11f9.
//
// Solidity: function nextUpdate() view returns(uint256)
func (_TBSCertList *TBSCertListCaller) NextUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "nextUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextUpdate is a free data retrieval call binding the contract method 0x6e9c11f9.
//
// Solidity: function nextUpdate() view returns(uint256)
func (_TBSCertList *TBSCertListSession) NextUpdate() (*big.Int, error) {
	return _TBSCertList.Contract.NextUpdate(&_TBSCertList.CallOpts)
}

// NextUpdate is a free data retrieval call binding the contract method 0x6e9c11f9.
//
// Solidity: function nextUpdate() view returns(uint256)
func (_TBSCertList *TBSCertListCallerSession) NextUpdate() (*big.Int, error) {
	return _TBSCertList.Contract.NextUpdate(&_TBSCertList.CallOpts)
}

// Raw is a free data retrieval call binding the contract method 0x302e678f.
//
// Solidity: function raw(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListCaller) Raw(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "raw", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Raw is a free data retrieval call binding the contract method 0x302e678f.
//
// Solidity: function raw(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListSession) Raw(arg0 *big.Int) (uint8, error) {
	return _TBSCertList.Contract.Raw(&_TBSCertList.CallOpts, arg0)
}

// Raw is a free data retrieval call binding the contract method 0x302e678f.
//
// Solidity: function raw(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListCallerSession) Raw(arg0 *big.Int) (uint8, error) {
	return _TBSCertList.Contract.Raw(&_TBSCertList.CallOpts, arg0)
}

// Signature is a free data retrieval call binding the contract method 0x70629548.
//
// Solidity: function signature(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListCaller) Signature(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "signature", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Signature is a free data retrieval call binding the contract method 0x70629548.
//
// Solidity: function signature(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListSession) Signature(arg0 *big.Int) (uint8, error) {
	return _TBSCertList.Contract.Signature(&_TBSCertList.CallOpts, arg0)
}

// Signature is a free data retrieval call binding the contract method 0x70629548.
//
// Solidity: function signature(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListCallerSession) Signature(arg0 *big.Int) (uint8, error) {
	return _TBSCertList.Contract.Signature(&_TBSCertList.CallOpts, arg0)
}

// ThisUpdate is a free data retrieval call binding the contract method 0xb4729258.
//
// Solidity: function thisUpdate() view returns(uint256)
func (_TBSCertList *TBSCertListCaller) ThisUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "thisUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ThisUpdate is a free data retrieval call binding the contract method 0xb4729258.
//
// Solidity: function thisUpdate() view returns(uint256)
func (_TBSCertList *TBSCertListSession) ThisUpdate() (*big.Int, error) {
	return _TBSCertList.Contract.ThisUpdate(&_TBSCertList.CallOpts)
}

// ThisUpdate is a free data retrieval call binding the contract method 0xb4729258.
//
// Solidity: function thisUpdate() view returns(uint256)
func (_TBSCertList *TBSCertListCallerSession) ThisUpdate() (*big.Int, error) {
	return _TBSCertList.Contract.ThisUpdate(&_TBSCertList.CallOpts)
}

// Verison is a free data retrieval call binding the contract method 0x38b9f2a8.
//
// Solidity: function verison() view returns(int32)
func (_TBSCertList *TBSCertListCaller) Verison(opts *bind.CallOpts) (int32, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "verison")

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// Verison is a free data retrieval call binding the contract method 0x38b9f2a8.
//
// Solidity: function verison() view returns(int32)
func (_TBSCertList *TBSCertListSession) Verison() (int32, error) {
	return _TBSCertList.Contract.Verison(&_TBSCertList.CallOpts)
}

// Verison is a free data retrieval call binding the contract method 0x38b9f2a8.
//
// Solidity: function verison() view returns(int32)
func (_TBSCertList *TBSCertListCallerSession) Verison() (int32, error) {
	return _TBSCertList.Contract.Verison(&_TBSCertList.CallOpts)
}

// AddRevokedCert is a paid mutator transaction binding the contract method 0x39c47dc6.
//
// Solidity: function addRevokedCert(address ref) returns()
func (_TBSCertList *TBSCertListTransactor) AddRevokedCert(opts *bind.TransactOpts, ref common.Address) (*types.Transaction, error) {
	return _TBSCertList.contract.Transact(opts, "addRevokedCert", ref)
}

// AddRevokedCert is a paid mutator transaction binding the contract method 0x39c47dc6.
//
// Solidity: function addRevokedCert(address ref) returns()
func (_TBSCertList *TBSCertListSession) AddRevokedCert(ref common.Address) (*types.Transaction, error) {
	return _TBSCertList.Contract.AddRevokedCert(&_TBSCertList.TransactOpts, ref)
}

// AddRevokedCert is a paid mutator transaction binding the contract method 0x39c47dc6.
//
// Solidity: function addRevokedCert(address ref) returns()
func (_TBSCertList *TBSCertListTransactorSession) AddRevokedCert(ref common.Address) (*types.Transaction, error) {
	return _TBSCertList.Contract.AddRevokedCert(&_TBSCertList.TransactOpts, ref)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_TBSCertList *TBSCertListTransactor) Finalize(opts *bind.TransactOpts, fin bool) (*types.Transaction, error) {
	return _TBSCertList.contract.Transact(opts, "finalize", fin)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_TBSCertList *TBSCertListSession) Finalize(fin bool) (*types.Transaction, error) {
	return _TBSCertList.Contract.Finalize(&_TBSCertList.TransactOpts, fin)
}

// Finalize is a paid mutator transaction binding the contract method 0x6c9789b0.
//
// Solidity: function finalize(bool fin) returns()
func (_TBSCertList *TBSCertListTransactorSession) Finalize(fin bool) (*types.Transaction, error) {
	return _TBSCertList.Contract.Finalize(&_TBSCertList.TransactOpts, fin)
}

// SetExtentions is a paid mutator transaction binding the contract method 0xf6cd261c.
//
// Solidity: function setExtentions(uint8[] data) returns()
func (_TBSCertList *TBSCertListTransactor) SetExtentions(opts *bind.TransactOpts, data []uint8) (*types.Transaction, error) {
	return _TBSCertList.contract.Transact(opts, "setExtentions", data)
}

// SetExtentions is a paid mutator transaction binding the contract method 0xf6cd261c.
//
// Solidity: function setExtentions(uint8[] data) returns()
func (_TBSCertList *TBSCertListSession) SetExtentions(data []uint8) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetExtentions(&_TBSCertList.TransactOpts, data)
}

// SetExtentions is a paid mutator transaction binding the contract method 0xf6cd261c.
//
// Solidity: function setExtentions(uint8[] data) returns()
func (_TBSCertList *TBSCertListTransactorSession) SetExtentions(data []uint8) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetExtentions(&_TBSCertList.TransactOpts, data)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x0afe3459.
//
// Solidity: function setIssuer(uint8[] data) returns()
func (_TBSCertList *TBSCertListTransactor) SetIssuer(opts *bind.TransactOpts, data []uint8) (*types.Transaction, error) {
	return _TBSCertList.contract.Transact(opts, "setIssuer", data)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x0afe3459.
//
// Solidity: function setIssuer(uint8[] data) returns()
func (_TBSCertList *TBSCertListSession) SetIssuer(data []uint8) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetIssuer(&_TBSCertList.TransactOpts, data)
}

// SetIssuer is a paid mutator transaction binding the contract method 0x0afe3459.
//
// Solidity: function setIssuer(uint8[] data) returns()
func (_TBSCertList *TBSCertListTransactorSession) SetIssuer(data []uint8) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetIssuer(&_TBSCertList.TransactOpts, data)
}

// SetRaw is a paid mutator transaction binding the contract method 0x41e83176.
//
// Solidity: function setRaw(uint8[] data) returns()
func (_TBSCertList *TBSCertListTransactor) SetRaw(opts *bind.TransactOpts, data []uint8) (*types.Transaction, error) {
	return _TBSCertList.contract.Transact(opts, "setRaw", data)
}

// SetRaw is a paid mutator transaction binding the contract method 0x41e83176.
//
// Solidity: function setRaw(uint8[] data) returns()
func (_TBSCertList *TBSCertListSession) SetRaw(data []uint8) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetRaw(&_TBSCertList.TransactOpts, data)
}

// SetRaw is a paid mutator transaction binding the contract method 0x41e83176.
//
// Solidity: function setRaw(uint8[] data) returns()
func (_TBSCertList *TBSCertListTransactorSession) SetRaw(data []uint8) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetRaw(&_TBSCertList.TransactOpts, data)
}

// SetSignature is a paid mutator transaction binding the contract method 0x0e075af9.
//
// Solidity: function setSignature(uint8[] data) returns()
func (_TBSCertList *TBSCertListTransactor) SetSignature(opts *bind.TransactOpts, data []uint8) (*types.Transaction, error) {
	return _TBSCertList.contract.Transact(opts, "setSignature", data)
}

// SetSignature is a paid mutator transaction binding the contract method 0x0e075af9.
//
// Solidity: function setSignature(uint8[] data) returns()
func (_TBSCertList *TBSCertListSession) SetSignature(data []uint8) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetSignature(&_TBSCertList.TransactOpts, data)
}

// SetSignature is a paid mutator transaction binding the contract method 0x0e075af9.
//
// Solidity: function setSignature(uint8[] data) returns()
func (_TBSCertList *TBSCertListTransactorSession) SetSignature(data []uint8) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetSignature(&_TBSCertList.TransactOpts, data)
}

// SetTimeAndVersion is a paid mutator transaction binding the contract method 0x1a1d042c.
//
// Solidity: function setTimeAndVersion(uint256 thisUpd, uint256 nextUpd, int32 version) returns()
func (_TBSCertList *TBSCertListTransactor) SetTimeAndVersion(opts *bind.TransactOpts, thisUpd *big.Int, nextUpd *big.Int, version int32) (*types.Transaction, error) {
	return _TBSCertList.contract.Transact(opts, "setTimeAndVersion", thisUpd, nextUpd, version)
}

// SetTimeAndVersion is a paid mutator transaction binding the contract method 0x1a1d042c.
//
// Solidity: function setTimeAndVersion(uint256 thisUpd, uint256 nextUpd, int32 version) returns()
func (_TBSCertList *TBSCertListSession) SetTimeAndVersion(thisUpd *big.Int, nextUpd *big.Int, version int32) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetTimeAndVersion(&_TBSCertList.TransactOpts, thisUpd, nextUpd, version)
}

// SetTimeAndVersion is a paid mutator transaction binding the contract method 0x1a1d042c.
//
// Solidity: function setTimeAndVersion(uint256 thisUpd, uint256 nextUpd, int32 version) returns()
func (_TBSCertList *TBSCertListTransactorSession) SetTimeAndVersion(thisUpd *big.Int, nextUpd *big.Int, version int32) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetTimeAndVersion(&_TBSCertList.TransactOpts, thisUpd, nextUpd, version)
}
