// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TKTimeLock

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

// TKTimeLockABI is the input ABI used to generate the binding from.
const TKTimeLockABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenLock\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDate\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"tokenLock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrewTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"createdAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"info\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenLock\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unlockDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TKTimeLockBin is the compiled bytecode used for deploying new contracts.
var TKTimeLockBin = "0x6080604052348015620000125760006000fd5b5060405162000b4a38038062000b4a83398181016040528101906200003891906200012a565b5b80600260005081909090555042600360005081909090555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b505050620002345662000233565b600081519050620000f381620001d9565b5b92915050565b6000815190506200010b81620001f7565b5b92915050565b600081519050620001238162000215565b5b92915050565b60006000600060608486031215620001425760006000fd5b60006200015286828701620000e2565b93505060206200016586828701620000fa565b9250506040620001788682870162000112565b9150505b9250925092565b60006200019082620001ad565b90505b919050565b6000620001a58262000183565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b620001e48162000183565b81141515620001f35760006000fd5b505b565b620002028162000198565b81141515620002115760006000fd5b505b565b6200022081620001ce565b811415156200022f5760006000fd5b505b565b5b61090680620002446000396000f3fe60806040523480156100115760006000fd5b50600436106100675760003560e01c8063370158ea1461006d57806369ac57211461008f5780638d8f2adb146100ad5780638da5cb5b146100b7578063cf09e0d0146100d5578063e718234d146100f357610067565b60006000fd5b610075610111565b60405161008695949392919061070d565b60405180910390f35b610097610230565b6040516100a49190610782565b60405180910390f35b6100b5610239565b005b6100bf6104fa565b6040516100cc9190610673565b60405180910390f35b6100dd610520565b6040516100ea9190610782565b60405180910390f35b6100fb610529565b60405161010891906106b9565b60405180910390f35b60006000600060006000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260005054600360005054600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016101c89190610673565b60206040518083038186803b1580156101e15760006000fd5b505afa1580156101f6573d600060003e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061021a91906105aa565b94509450945094509450610229565b9091929394565b60026000505481565b600260005054421015151561024e5760006000fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016102ab9190610673565b60206040518083038186803b1580156102c45760006000fd5b505afa1580156102d9573d600060003e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102fd91906105aa565b905060006000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168560405160240161037692919061068f565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516103c4919061065b565b6000604051808303816000865af19150503d8060008114610401576040519150601f19603f3d011682016040523d82523d6000602084013e610406565b606091505b50915091508180156104345750600081511480610433575080806020019051810190610432919061057f565b5b5b1515610475576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046c90610761565b60405180910390fd5b7fa04c965e7d289704527a948915777d68bb4530fb4b08aa935fb958c0b7901ab9600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040516104ec939291906106d5565b60405180910390a15050505b565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036000505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681566108cf565b60008151905061056281610899565b5b92915050565b600081519050610578816108b4565b5b92915050565b6000602082840312156105925760006000fd5b60006105a084828501610553565b9150505b92915050565b6000602082840312156105bd5760006000fd5b60006105cb84828501610569565b9150505b92915050565b6105de816107c8565b825250505b565b60006105f08261079e565b6105fa81856107aa565b935061060a81856020860161083a565b8084019150505b92915050565b61062081610814565b825250505b565b6000610634601f836107b6565b915061063f8261086f565b6020820190505b919050565b61065481610809565b825250505b565b600061066782846105e5565b91508190505b92915050565b600060208201905061068860008301846105d5565b5b92915050565b60006040820190506106a460008301856105d5565b6106b1602083018461064b565b5b9392505050565b60006020820190506106ce6000830184610617565b5b92915050565b60006060820190506106ea6000830186610617565b6106f760208301856105d5565b610704604083018461064b565b5b949350505050565b600060a0820190506107226000830188610617565b61072f60208301876105d5565b61073c604083018661064b565b610749606083018561064b565b610756608083018461064b565b5b9695505050505050565b6000602082019050818103600083015261077a81610627565b90505b919050565b6000602082019050610797600083018461064b565b5b92915050565b6000815190505b919050565b60008190505b92915050565b60008282526020820190505b92915050565b60006107d3826107e8565b90505b919050565b600081151590505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600061081f82610827565b90505b919050565b6000610832826107e8565b90505b919050565b60005b838110156108595780820151818401525b60208101905061083d565b83811115610868576000848401525b505050505b565b7f5472616e7366657248656c7065723a205452414e534645525f4641494c4544006000820152505b565b6108a2816107db565b811415156108b05760006000fd5b505b565b6108bd81610809565b811415156108cb5760006000fd5b505b565bfea264697066735822122088cf3cd406b46bcae58bff6b360c974ddefd6598f071b43c50293da06e6b346464736f6c63430008040033"

// DeployTKTimeLock deploys a new Ethereum contract, binding an instance of TKTimeLock to it.
func DeployTKTimeLock(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _tokenLock common.Address, _unlockDate *big.Int) (common.Address, *types.Transaction, *TKTimeLock, error) {
	parsed, err := abi.JSON(strings.NewReader(TKTimeLockABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TKTimeLockBin), backend, _owner, _tokenLock, _unlockDate)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TKTimeLock{TKTimeLockCaller: TKTimeLockCaller{contract: contract}, TKTimeLockTransactor: TKTimeLockTransactor{contract: contract}, TKTimeLockFilterer: TKTimeLockFilterer{contract: contract}}, nil
}

// TKTimeLock is an auto generated Go binding around an Ethereum contract.
type TKTimeLock struct {
	TKTimeLockCaller     // Read-only binding to the contract
	TKTimeLockTransactor // Write-only binding to the contract
	TKTimeLockFilterer   // Log filterer for contract events
}

// TKTimeLockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TKTimeLockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TKTimeLockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TKTimeLockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TKTimeLockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TKTimeLockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TKTimeLockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TKTimeLockSession struct {
	Contract     *TKTimeLock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TKTimeLockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TKTimeLockCallerSession struct {
	Contract *TKTimeLockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TKTimeLockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TKTimeLockTransactorSession struct {
	Contract     *TKTimeLockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TKTimeLockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TKTimeLockRaw struct {
	Contract *TKTimeLock // Generic contract binding to access the raw methods on
}

// TKTimeLockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TKTimeLockCallerRaw struct {
	Contract *TKTimeLockCaller // Generic read-only contract binding to access the raw methods on
}

// TKTimeLockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TKTimeLockTransactorRaw struct {
	Contract *TKTimeLockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTKTimeLock creates a new instance of TKTimeLock, bound to a specific deployed contract.
func NewTKTimeLock(address common.Address, backend bind.ContractBackend) (*TKTimeLock, error) {
	contract, err := bindTKTimeLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TKTimeLock{TKTimeLockCaller: TKTimeLockCaller{contract: contract}, TKTimeLockTransactor: TKTimeLockTransactor{contract: contract}, TKTimeLockFilterer: TKTimeLockFilterer{contract: contract}}, nil
}

// NewTKTimeLockCaller creates a new read-only instance of TKTimeLock, bound to a specific deployed contract.
func NewTKTimeLockCaller(address common.Address, caller bind.ContractCaller) (*TKTimeLockCaller, error) {
	contract, err := bindTKTimeLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TKTimeLockCaller{contract: contract}, nil
}

// NewTKTimeLockTransactor creates a new write-only instance of TKTimeLock, bound to a specific deployed contract.
func NewTKTimeLockTransactor(address common.Address, transactor bind.ContractTransactor) (*TKTimeLockTransactor, error) {
	contract, err := bindTKTimeLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TKTimeLockTransactor{contract: contract}, nil
}

// NewTKTimeLockFilterer creates a new log filterer instance of TKTimeLock, bound to a specific deployed contract.
func NewTKTimeLockFilterer(address common.Address, filterer bind.ContractFilterer) (*TKTimeLockFilterer, error) {
	contract, err := bindTKTimeLock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TKTimeLockFilterer{contract: contract}, nil
}

// bindTKTimeLock binds a generic wrapper to an already deployed contract.
func bindTKTimeLock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TKTimeLockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TKTimeLock *TKTimeLockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TKTimeLock.Contract.TKTimeLockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TKTimeLock *TKTimeLockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TKTimeLock.Contract.TKTimeLockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TKTimeLock *TKTimeLockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TKTimeLock.Contract.TKTimeLockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TKTimeLock *TKTimeLockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TKTimeLock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TKTimeLock *TKTimeLockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TKTimeLock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TKTimeLock *TKTimeLockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TKTimeLock.Contract.contract.Transact(opts, method, params...)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_TKTimeLock *TKTimeLockCaller) CreatedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TKTimeLock.contract.Call(opts, &out, "createdAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_TKTimeLock *TKTimeLockSession) CreatedAt() (*big.Int, error) {
	return _TKTimeLock.Contract.CreatedAt(&_TKTimeLock.CallOpts)
}

// CreatedAt is a free data retrieval call binding the contract method 0xcf09e0d0.
//
// Solidity: function createdAt() view returns(uint256)
func (_TKTimeLock *TKTimeLockCallerSession) CreatedAt() (*big.Int, error) {
	return _TKTimeLock.Contract.CreatedAt(&_TKTimeLock.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() view returns(address, address, uint256, uint256, uint256)
func (_TKTimeLock *TKTimeLockCaller) Info(opts *bind.CallOpts) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _TKTimeLock.contract.Call(opts, &out, "info")

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() view returns(address, address, uint256, uint256, uint256)
func (_TKTimeLock *TKTimeLockSession) Info() (common.Address, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _TKTimeLock.Contract.Info(&_TKTimeLock.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() view returns(address, address, uint256, uint256, uint256)
func (_TKTimeLock *TKTimeLockCallerSession) Info() (common.Address, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _TKTimeLock.Contract.Info(&_TKTimeLock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TKTimeLock *TKTimeLockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TKTimeLock.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TKTimeLock *TKTimeLockSession) Owner() (common.Address, error) {
	return _TKTimeLock.Contract.Owner(&_TKTimeLock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TKTimeLock *TKTimeLockCallerSession) Owner() (common.Address, error) {
	return _TKTimeLock.Contract.Owner(&_TKTimeLock.CallOpts)
}

// TokenLock is a free data retrieval call binding the contract method 0xe718234d.
//
// Solidity: function tokenLock() view returns(address)
func (_TKTimeLock *TKTimeLockCaller) TokenLock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TKTimeLock.contract.Call(opts, &out, "tokenLock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenLock is a free data retrieval call binding the contract method 0xe718234d.
//
// Solidity: function tokenLock() view returns(address)
func (_TKTimeLock *TKTimeLockSession) TokenLock() (common.Address, error) {
	return _TKTimeLock.Contract.TokenLock(&_TKTimeLock.CallOpts)
}

// TokenLock is a free data retrieval call binding the contract method 0xe718234d.
//
// Solidity: function tokenLock() view returns(address)
func (_TKTimeLock *TKTimeLockCallerSession) TokenLock() (common.Address, error) {
	return _TKTimeLock.Contract.TokenLock(&_TKTimeLock.CallOpts)
}

// UnlockDate is a free data retrieval call binding the contract method 0x69ac5721.
//
// Solidity: function unlockDate() view returns(uint256)
func (_TKTimeLock *TKTimeLockCaller) UnlockDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TKTimeLock.contract.Call(opts, &out, "unlockDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockDate is a free data retrieval call binding the contract method 0x69ac5721.
//
// Solidity: function unlockDate() view returns(uint256)
func (_TKTimeLock *TKTimeLockSession) UnlockDate() (*big.Int, error) {
	return _TKTimeLock.Contract.UnlockDate(&_TKTimeLock.CallOpts)
}

// UnlockDate is a free data retrieval call binding the contract method 0x69ac5721.
//
// Solidity: function unlockDate() view returns(uint256)
func (_TKTimeLock *TKTimeLockCallerSession) UnlockDate() (*big.Int, error) {
	return _TKTimeLock.Contract.UnlockDate(&_TKTimeLock.CallOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_TKTimeLock *TKTimeLockTransactor) WithdrawTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TKTimeLock.contract.Transact(opts, "withdrawTokens")
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_TKTimeLock *TKTimeLockSession) WithdrawTokens() (*types.Transaction, error) {
	return _TKTimeLock.Contract.WithdrawTokens(&_TKTimeLock.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x8d8f2adb.
//
// Solidity: function withdrawTokens() returns()
func (_TKTimeLock *TKTimeLockTransactorSession) WithdrawTokens() (*types.Transaction, error) {
	return _TKTimeLock.Contract.WithdrawTokens(&_TKTimeLock.TransactOpts)
}

// TKTimeLockWithdrewTokensIterator is returned from FilterWithdrewTokens and is used to iterate over the raw logs and unpacked data for WithdrewTokens events raised by the TKTimeLock contract.
type TKTimeLockWithdrewTokensIterator struct {
	Event *TKTimeLockWithdrewTokens // Event containing the contract specifics and raw log

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
func (it *TKTimeLockWithdrewTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TKTimeLockWithdrewTokens)
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
		it.Event = new(TKTimeLockWithdrewTokens)
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
func (it *TKTimeLockWithdrewTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TKTimeLockWithdrewTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TKTimeLockWithdrewTokens represents a WithdrewTokens event raised by the TKTimeLock contract.
type TKTimeLockWithdrewTokens struct {
	TokenLock common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrewTokens is a free log retrieval operation binding the contract event 0xa04c965e7d289704527a948915777d68bb4530fb4b08aa935fb958c0b7901ab9.
//
// Solidity: event WithdrewTokens(address tokenLock, address to, uint256 amount)
func (_TKTimeLock *TKTimeLockFilterer) FilterWithdrewTokens(opts *bind.FilterOpts) (*TKTimeLockWithdrewTokensIterator, error) {

	logs, sub, err := _TKTimeLock.contract.FilterLogs(opts, "WithdrewTokens")
	if err != nil {
		return nil, err
	}
	return &TKTimeLockWithdrewTokensIterator{contract: _TKTimeLock.contract, event: "WithdrewTokens", logs: logs, sub: sub}, nil
}

// WatchWithdrewTokens is a free log subscription operation binding the contract event 0xa04c965e7d289704527a948915777d68bb4530fb4b08aa935fb958c0b7901ab9.
//
// Solidity: event WithdrewTokens(address tokenLock, address to, uint256 amount)
func (_TKTimeLock *TKTimeLockFilterer) WatchWithdrewTokens(opts *bind.WatchOpts, sink chan<- *TKTimeLockWithdrewTokens) (event.Subscription, error) {

	logs, sub, err := _TKTimeLock.contract.WatchLogs(opts, "WithdrewTokens")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TKTimeLockWithdrewTokens)
				if err := _TKTimeLock.contract.UnpackLog(event, "WithdrewTokens", log); err != nil {
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

// ParseWithdrewTokens is a log parse operation binding the contract event 0xa04c965e7d289704527a948915777d68bb4530fb4b08aa935fb958c0b7901ab9.
//
// Solidity: event WithdrewTokens(address tokenLock, address to, uint256 amount)
func (_TKTimeLock *TKTimeLockFilterer) ParseWithdrewTokens(log types.Log) (*TKTimeLockWithdrewTokens, error) {
	event := new(TKTimeLockWithdrewTokens)
	if err := _TKTimeLock.contract.UnpackLog(event, "WithdrewTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
