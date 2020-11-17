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
const TBSCertListABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ref\",\"type\":\"address\"}],\"name\":\"addRevokedCert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"extensions\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"fin\",\"type\":\"bool\"}],\"name\":\"finalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"raw\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"revokedCertList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setExtensions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setIssuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setRaw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8[]\",\"name\":\"data\",\"type\":\"uint8[]\"}],\"name\":\"setSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"thisUpd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextUpd\",\"type\":\"uint256\"},{\"internalType\":\"int32\",\"name\":\"ver\",\"type\":\"int32\"}],\"name\":\"setTimeAndVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signature\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"thisUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TBSCertListFuncSigs maps the 4-byte function signature to its string representation.
var TBSCertListFuncSigs = map[string]string{
	"39c47dc6": "addRevokedCert(address)",
	"db85d59c": "extensions(uint256)",
	"6c9789b0": "finalize(bool)",
	"b3f05b97": "finalized()",
	"8d4e4083": "isFinalized()",
	"0068fe8b": "issuer(uint256)",
	"6e9c11f9": "nextUpdate()",
	"302e678f": "raw(uint256)",
	"ec53fc06": "revokedCertList(uint256)",
	"2fc02f22": "setExtensions(uint8[])",
	"0afe3459": "setIssuer(uint8[])",
	"41e83176": "setRaw(uint8[])",
	"0e075af9": "setSignature(uint8[])",
	"1a1d042c": "setTimeAndVersion(uint256,uint256,int32)",
	"70629548": "signature(uint256)",
	"b4729258": "thisUpdate()",
	"54fd4d50": "version()",
}

// TBSCertListBin is the compiled bytecode used for deploying new contracts.
var TBSCertListBin = "0x608060405234801561001057600080fd5b50610aa2806100206000396000f3fe608060405234801561001057600080fd5b506004361061010a5760003560e01c806354fd4d50116100a25780638d4e4083116100715780638d4e4083146104ac578063b3f05b97146104c8578063b4729258146104d0578063db85d59c146104d8578063ec53fc06146104f55761010a565b806354fd4d50146104375780636c9789b0146104565780636e9c11f914610475578063706295481461048f5761010a565b80632fc02f22116100de5780632fc02f22146102b2578063302e678f1461035357806339c47dc61461037057806341e83176146103965761010a565b806268fe8b1461010f5780630afe3459146101425780630e075af9146101e55780631a1d042c14610286575b600080fd5b61012c6004803603602081101561012557600080fd5b503561052e565b6040805160ff9092168252519081900360200190f35b6101e36004803603602081101561015857600080fd5b810190602081018135600160201b81111561017257600080fd5b82018360208201111561018457600080fd5b803590602001918460208302840111600160201b831117156101a557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610562945050505050565b005b6101e3600480360360208110156101fb57600080fd5b810190602081018135600160201b81111561021557600080fd5b82018360208201111561022757600080fd5b803590602001918460208302840111600160201b8311171561024857600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061060c945050505050565b6101e36004803603606081101561029c57600080fd5b508035906020810135906040013560030b6106b1565b6101e3600480360360208110156102c857600080fd5b810190602081018135600160201b8111156102e257600080fd5b8201836020820111156102f457600080fd5b803590602001918460208302840111600160201b8311171561031557600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061071d945050505050565b61012c6004803603602081101561036957600080fd5b50356107c2565b6101e36004803603602081101561038657600080fd5b50356001600160a01b03166107d2565b6101e3600480360360208110156103ac57600080fd5b810190602081018135600160201b8111156103c657600080fd5b8201836020820111156103d857600080fd5b803590602001918460208302840111600160201b831117156103f957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506108d7945050505050565b61043f61097c565b6040805160039290920b8252519081900360200190f35b6101e36004803603602081101561046c57600080fd5b50351515610985565b61047d6109da565b60408051918252519081900360200190f35b61012c600480360360208110156104a557600080fd5b50356109e0565b6104b46109f0565b604080519115158252519081900360200190f35b6104b46109f9565b61047d610a02565b61012c600480360360208110156104ee57600080fd5b5035610a08565b6105126004803603602081101561050b57600080fd5b5035610a18565b604080516001600160a01b039092168252519081900360200190f35b6004818154811061053e57600080fd5b9060005260206000209060209182820401919006915054906101000a900460ff1681565b60005460ff16156105a45760405162461bcd60e51b815260040180806020018281038252602a815260200180610a43602a913960400191505060405180910390fd5b805160005b818110156106075760048382815181106105bf57fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a9384029290930219909216179055016105a9565b505050565b60005460ff161561064e5760405162461bcd60e51b815260040180806020018281038252602a815260200180610a43602a913960400191505060405180910390fd5b805160005b8181101561060757600383828151811061066957fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a938402929093021990921617905501610653565b60005460ff16156106f35760405162461bcd60e51b815260040180806020018281038252602a815260200180610a43602a913960400191505060405180910390fd5b6005929092556006556002805460039290920b63ffffffff1663ffffffff19909216919091179055565b60005460ff161561075f5760405162461bcd60e51b815260040180806020018281038252602a815260200180610a43602a913960400191505060405180910390fd5b805160005b8181101561060757600783828151811061077a57fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a938402929093021990921617905501610764565b6001818154811061053e57600080fd5b60005460ff16156108145760405162461bcd60e51b815260040180806020018281038252602a815260200180610a43602a913960400191505060405180910390fd5b6000819050806001600160a01b0316638d4e40836040518163ffffffff1660e01b815260040160206040518083038186803b15801561085257600080fd5b505afa158015610866573d6000803e3d6000fd5b505050506040513d602081101561087c57600080fd5b505161088457fe5b50600880546001810182556000919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30180546001600160a01b0319166001600160a01b0392909216919091179055565b60005460ff16156109195760405162461bcd60e51b815260040180806020018281038252602a815260200180610a43602a913960400191505060405180910390fd5b805160005b8181101561060757600183828151811061093457fe5b602090810291909101810151825460018181018555600094855293839020928104909201805460ff928316601f9094166101000a93840292909302199092161790550161091e565b60025460030b81565b60005460ff16156109c75760405162461bcd60e51b815260040180806020018281038252602a815260200180610a43602a913960400191505060405180910390fd5b6000805460ff1916911515919091179055565b60065481565b6003818154811061053e57600080fd5b60005460ff1690565b60005460ff1681565b60055481565b6007818154811061053e57600080fd5b60088181548110610a2857600080fd5b6000918252602090912001546001600160a01b031690508156fe46696e616c697a61626c653a20636f6e747261637420697320616c72656164792066696e616c697a6564a2646970667358221220bd52fe7f817ee23f9243c0cb2b02365e0edb8984639e3f3da960ff9b147a42d964736f6c63430007040033"

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

// Extensions is a free data retrieval call binding the contract method 0xdb85d59c.
//
// Solidity: function extensions(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListCaller) Extensions(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "extensions", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Extensions is a free data retrieval call binding the contract method 0xdb85d59c.
//
// Solidity: function extensions(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListSession) Extensions(arg0 *big.Int) (uint8, error) {
	return _TBSCertList.Contract.Extensions(&_TBSCertList.CallOpts, arg0)
}

// Extensions is a free data retrieval call binding the contract method 0xdb85d59c.
//
// Solidity: function extensions(uint256 ) view returns(uint8)
func (_TBSCertList *TBSCertListCallerSession) Extensions(arg0 *big.Int) (uint8, error) {
	return _TBSCertList.Contract.Extensions(&_TBSCertList.CallOpts, arg0)
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

// RevokedCertList is a free data retrieval call binding the contract method 0xec53fc06.
//
// Solidity: function revokedCertList(uint256 ) view returns(address)
func (_TBSCertList *TBSCertListCaller) RevokedCertList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "revokedCertList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RevokedCertList is a free data retrieval call binding the contract method 0xec53fc06.
//
// Solidity: function revokedCertList(uint256 ) view returns(address)
func (_TBSCertList *TBSCertListSession) RevokedCertList(arg0 *big.Int) (common.Address, error) {
	return _TBSCertList.Contract.RevokedCertList(&_TBSCertList.CallOpts, arg0)
}

// RevokedCertList is a free data retrieval call binding the contract method 0xec53fc06.
//
// Solidity: function revokedCertList(uint256 ) view returns(address)
func (_TBSCertList *TBSCertListCallerSession) RevokedCertList(arg0 *big.Int) (common.Address, error) {
	return _TBSCertList.Contract.RevokedCertList(&_TBSCertList.CallOpts, arg0)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(int32)
func (_TBSCertList *TBSCertListCaller) Version(opts *bind.CallOpts) (int32, error) {
	var out []interface{}
	err := _TBSCertList.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(int32)
func (_TBSCertList *TBSCertListSession) Version() (int32, error) {
	return _TBSCertList.Contract.Version(&_TBSCertList.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(int32)
func (_TBSCertList *TBSCertListCallerSession) Version() (int32, error) {
	return _TBSCertList.Contract.Version(&_TBSCertList.CallOpts)
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

// SetExtensions is a paid mutator transaction binding the contract method 0x2fc02f22.
//
// Solidity: function setExtensions(uint8[] data) returns()
func (_TBSCertList *TBSCertListTransactor) SetExtensions(opts *bind.TransactOpts, data []uint8) (*types.Transaction, error) {
	return _TBSCertList.contract.Transact(opts, "setExtensions", data)
}

// SetExtensions is a paid mutator transaction binding the contract method 0x2fc02f22.
//
// Solidity: function setExtensions(uint8[] data) returns()
func (_TBSCertList *TBSCertListSession) SetExtensions(data []uint8) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetExtensions(&_TBSCertList.TransactOpts, data)
}

// SetExtensions is a paid mutator transaction binding the contract method 0x2fc02f22.
//
// Solidity: function setExtensions(uint8[] data) returns()
func (_TBSCertList *TBSCertListTransactorSession) SetExtensions(data []uint8) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetExtensions(&_TBSCertList.TransactOpts, data)
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
// Solidity: function setTimeAndVersion(uint256 thisUpd, uint256 nextUpd, int32 ver) returns()
func (_TBSCertList *TBSCertListTransactor) SetTimeAndVersion(opts *bind.TransactOpts, thisUpd *big.Int, nextUpd *big.Int, ver int32) (*types.Transaction, error) {
	return _TBSCertList.contract.Transact(opts, "setTimeAndVersion", thisUpd, nextUpd, ver)
}

// SetTimeAndVersion is a paid mutator transaction binding the contract method 0x1a1d042c.
//
// Solidity: function setTimeAndVersion(uint256 thisUpd, uint256 nextUpd, int32 ver) returns()
func (_TBSCertList *TBSCertListSession) SetTimeAndVersion(thisUpd *big.Int, nextUpd *big.Int, ver int32) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetTimeAndVersion(&_TBSCertList.TransactOpts, thisUpd, nextUpd, ver)
}

// SetTimeAndVersion is a paid mutator transaction binding the contract method 0x1a1d042c.
//
// Solidity: function setTimeAndVersion(uint256 thisUpd, uint256 nextUpd, int32 ver) returns()
func (_TBSCertList *TBSCertListTransactorSession) SetTimeAndVersion(thisUpd *big.Int, nextUpd *big.Int, ver int32) (*types.Transaction, error) {
	return _TBSCertList.Contract.SetTimeAndVersion(&_TBSCertList.TransactOpts, thisUpd, nextUpd, ver)
}
