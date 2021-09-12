// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TKLockTime

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

// TKLockTimeABI is the input ABI used to generate the binding from.
const TKLockTimeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Created\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getWallets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_tokenLock\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlockDate\",\"type\":\"uint256\"}],\"name\":\"newTimeLockedWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TKLockTimeBin is the compiled bytecode used for deploying new contracts.
var TKLockTimeBin = "0x60806040523480156100115760006000fd5b505b61002f61002461003560201b60201c565b61004260201b60201c565b5b610109565b600033905061003f565b90565b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505b50565b611a5c806101186000396000f3fe6080604052348015620000125760006000fd5b5060043610620000645760003560e01c80633f106cde146200006a578063422c29a414620000a0578063715018a614620000d65780638da5cb5b14620000e2578063f2fde38b14620001045762000064565b60006000fd5b6200008860048036038101906200008291906200091c565b62000124565b60405162000097919062000b2f565b60405180910390f35b620000be6004803603810190620000b89190620008ee565b620004c0565b604051620000cd919062000c55565b60405180910390f35b620000e06200059b565b005b620000ec62000642565b604051620000fb919062000b2f565b60405180910390f35b6200012260048036038101906200011c9190620008ee565b62000672565b005b6000620001366200078e63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff166200015c6200064263ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16141515620001b7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001ae9062000cc0565b60405180910390fd5b828473ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401620001f592919062000b4d565b60206040518083038186803b1580156200020f5760006000fd5b505afa15801562000225573d600060003e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200024b9190620009ba565b101515156200025a5760006000fd5b42821115156200026a5760006000fd5b8473ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff16141515620002dd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002d49062000c7a565b60405180910390fd5b6000858584604051620002f09062000863565b620002fe9392919062000c17565b604051809103906000f0801580156200031c573d600060003e3d6000fd5b5090508473ffffffffffffffffffffffffffffffffffffffff166323b872dd3383876040518463ffffffff1660e01b81526004016200035e9392919062000b7b565b602060405180830381600087803b1580156200037a5760006000fd5b505af115801562000390573d600060003e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003b691906200098c565b1515620003c35760006000fd5b600160005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005081908060018154018082558091505060019003906000526020600020900160005b9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2abec6ac791844b0f5490c49d67e9eede10b60ff00b9e1180467b970e8fc8b008187428688604051620004a495949392919062000bb9565b60405180910390a180915050620004b756505b5b949350505050565b6060600160005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508054806020026020016040519081016040528092919081815260200182805480156200058957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116200053e575b5050505050905062000596565b919050565b620005ab6200078e63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16620005d16200064263ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff161415156200062c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620006239062000cc0565b60405180910390fd5b6200063e60006200079c63ffffffff16565b5b5b565b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506200066f565b90565b620006826200078e63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16620006a86200064263ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1614151562000703576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620006fa9062000cc0565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151562000778576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200076f9062000c9d565b60405180910390fd5b62000789816200079c63ffffffff16565b5b5b50565b600033905062000799565b90565b6000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505b50565b610b4a8062000edd833901905662000edb565b600081359050620008878162000e63565b5b92915050565b6000815190506200089f8162000e81565b5b92915050565b600081359050620008b78162000e9f565b5b92915050565b600081359050620008cf8162000ebd565b5b92915050565b600081519050620008e78162000ebd565b5b92915050565b600060208284031215620009025760006000fd5b6000620009128482850162000876565b9150505b92915050565b600060006000600060808587031215620009365760006000fd5b6000620009468782880162000876565b94505060206200095987828801620008a6565b93505060406200096c87828801620008be565b92505060606200097f87828801620008be565b9150505b92959194509250565b600060208284031215620009a05760006000fd5b6000620009b0848285016200088e565b9150505b92915050565b600060208284031215620009ce5760006000fd5b6000620009de84828501620008d6565b9150505b92915050565b6000620009f6838362000a03565b6020830190505b92915050565b62000a0e8162000d32565b825250505b565b62000a208162000d32565b825250505b565b600062000a348262000cf4565b62000a40818562000d0e565b935062000a4d8362000ce3565b8060005b8381101562000a8557815162000a688882620009e8565b975062000a758362000d00565b9250505b60018101905062000a51565b508593505050505b92915050565b62000a9e8162000d95565b825250505b565b600062000ab4600f8362000d20565b915062000ac18262000dbf565b6020820190505b919050565b600062000adc60268362000d20565b915062000ae98262000de9565b6040820190505b919050565b600062000b0460208362000d20565b915062000b118262000e39565b6020820190505b919050565b62000b288162000d8a565b825250505b565b600060208201905062000b46600083018462000a15565b5b92915050565b600060408201905062000b64600083018562000a15565b62000b73602083018462000a15565b5b9392505050565b600060608201905062000b92600083018662000a15565b62000ba1602083018562000a15565b62000bb0604083018462000b1d565b5b949350505050565b600060a08201905062000bd0600083018862000a15565b62000bdf602083018762000a15565b62000bee604083018662000b1d565b62000bfd606083018562000b1d565b62000c0c608083018462000b1d565b5b9695505050505050565b600060608201905062000c2e600083018662000a15565b62000c3d602083018562000a93565b62000c4c604083018462000b1d565b5b949350505050565b6000602082019050818103600083015262000c71818462000a27565b90505b92915050565b6000602082019050818103600083015262000c958162000aa5565b90505b919050565b6000602082019050818103600083015262000cb88162000acd565b90505b919050565b6000602082019050818103600083015262000cdb8162000af5565b90505b919050565b60008190506020820190505b919050565b6000815190505b919050565b60006020820190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b600062000d3f8262000d69565b90505b919050565b600081151590505b919050565b600062000d618262000d32565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600062000da28262000daa565b90505b919050565b600062000db78262000d69565b90505b919050565b7f496e76616c6964206164647265737300000000000000000000000000000000006000820152505b565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f64647265737300000000000000000000000000000000000000000000000000006020820152505b565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000820152505b565b62000e6e8162000d32565b8114151562000e7d5760006000fd5b505b565b62000e8c8162000d47565b8114151562000e9b5760006000fd5b505b565b62000eaa8162000d54565b8114151562000eb95760006000fd5b505b565b62000ec88162000d8a565b8114151562000ed75760006000fd5b505b565bfe6080604052348015620000125760006000fd5b5060405162000b4a38038062000b4a83398181016040528101906200003891906200012a565b5b80600260005081909090555042600360005081909090555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b505050620002345662000233565b600081519050620000f381620001d9565b5b92915050565b6000815190506200010b81620001f7565b5b92915050565b600081519050620001238162000215565b5b92915050565b60006000600060608486031215620001425760006000fd5b60006200015286828701620000e2565b93505060206200016586828701620000fa565b9250506040620001788682870162000112565b9150505b9250925092565b60006200019082620001ad565b90505b919050565b6000620001a58262000183565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b620001e48162000183565b81141515620001f35760006000fd5b505b565b620002028162000198565b81141515620002115760006000fd5b505b565b6200022081620001ce565b811415156200022f5760006000fd5b505b565b5b61090680620002446000396000f3fe60806040523480156100115760006000fd5b50600436106100675760003560e01c8063370158ea1461006d57806369ac57211461008f5780638d8f2adb146100ad5780638da5cb5b146100b7578063cf09e0d0146100d5578063e718234d146100f357610067565b60006000fd5b610075610111565b60405161008695949392919061070d565b60405180910390f35b610097610230565b6040516100a49190610782565b60405180910390f35b6100b5610239565b005b6100bf6104fa565b6040516100cc9190610673565b60405180910390f35b6100dd610520565b6040516100ea9190610782565b60405180910390f35b6100fb610529565b60405161010891906106b9565b60405180910390f35b60006000600060006000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260005054600360005054600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016101c89190610673565b60206040518083038186803b1580156101e15760006000fd5b505afa1580156101f6573d600060003e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061021a91906105aa565b94509450945094509450610229565b9091929394565b60026000505481565b600260005054421015151561024e5760006000fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016102ab9190610673565b60206040518083038186803b1580156102c45760006000fd5b505afa1580156102d9573d600060003e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102fd91906105aa565b905060006000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168560405160240161037692919061068f565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516103c4919061065b565b6000604051808303816000865af19150503d8060008114610401576040519150601f19603f3d011682016040523d82523d6000602084013e610406565b606091505b50915091508180156104345750600081511480610433575080806020019051810190610432919061057f565b5b5b1515610475576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046c90610761565b60405180910390fd5b7fa04c965e7d289704527a948915777d68bb4530fb4b08aa935fb958c0b7901ab9600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856040516104ec939291906106d5565b60405180910390a15050505b565b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036000505481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681566108cf565b60008151905061056281610899565b5b92915050565b600081519050610578816108b4565b5b92915050565b6000602082840312156105925760006000fd5b60006105a084828501610553565b9150505b92915050565b6000602082840312156105bd5760006000fd5b60006105cb84828501610569565b9150505b92915050565b6105de816107c8565b825250505b565b60006105f08261079e565b6105fa81856107aa565b935061060a81856020860161083a565b8084019150505b92915050565b61062081610814565b825250505b565b6000610634601f836107b6565b915061063f8261086f565b6020820190505b919050565b61065481610809565b825250505b565b600061066782846105e5565b91508190505b92915050565b600060208201905061068860008301846105d5565b5b92915050565b60006040820190506106a460008301856105d5565b6106b1602083018461064b565b5b9392505050565b60006020820190506106ce6000830184610617565b5b92915050565b60006060820190506106ea6000830186610617565b6106f760208301856105d5565b610704604083018461064b565b5b949350505050565b600060a0820190506107226000830188610617565b61072f60208301876105d5565b61073c604083018661064b565b610749606083018561064b565b610756608083018461064b565b5b9695505050505050565b6000602082019050818103600083015261077a81610627565b90505b919050565b6000602082019050610797600083018461064b565b5b92915050565b6000815190505b919050565b60008190505b92915050565b60008282526020820190505b92915050565b60006107d3826107e8565b90505b919050565b600081151590505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600061081f82610827565b90505b919050565b6000610832826107e8565b90505b919050565b60005b838110156108595780820151818401525b60208101905061083d565b83811115610868576000848401525b505050505b565b7f5472616e7366657248656c7065723a205452414e534645525f4641494c4544006000820152505b565b6108a2816107db565b811415156108b05760006000fd5b505b565b6108bd81610809565b811415156108cb5760006000fd5b505b565bfea264697066735822122088cf3cd406b46bcae58bff6b360c974ddefd6598f071b43c50293da06e6b346464736f6c63430008040033a26469706673582212201f548fdf8f2184556b7fd6d35293159419e4ab1318136914477cea9b5801fb1464736f6c63430008040033"

// DeployTKLockTime deploys a new Ethereum contract, binding an instance of TKLockTime to it.
func DeployTKLockTime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TKLockTime, error) {
	parsed, err := abi.JSON(strings.NewReader(TKLockTimeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TKLockTimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TKLockTime{TKLockTimeCaller: TKLockTimeCaller{contract: contract}, TKLockTimeTransactor: TKLockTimeTransactor{contract: contract}, TKLockTimeFilterer: TKLockTimeFilterer{contract: contract}}, nil
}

// TKLockTime is an auto generated Go binding around an Ethereum contract.
type TKLockTime struct {
	TKLockTimeCaller     // Read-only binding to the contract
	TKLockTimeTransactor // Write-only binding to the contract
	TKLockTimeFilterer   // Log filterer for contract events
}

// TKLockTimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TKLockTimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TKLockTimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TKLockTimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TKLockTimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TKLockTimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TKLockTimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TKLockTimeSession struct {
	Contract     *TKLockTime       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TKLockTimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TKLockTimeCallerSession struct {
	Contract *TKLockTimeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TKLockTimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TKLockTimeTransactorSession struct {
	Contract     *TKLockTimeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TKLockTimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TKLockTimeRaw struct {
	Contract *TKLockTime // Generic contract binding to access the raw methods on
}

// TKLockTimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TKLockTimeCallerRaw struct {
	Contract *TKLockTimeCaller // Generic read-only contract binding to access the raw methods on
}

// TKLockTimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TKLockTimeTransactorRaw struct {
	Contract *TKLockTimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTKLockTime creates a new instance of TKLockTime, bound to a specific deployed contract.
func NewTKLockTime(address common.Address, backend bind.ContractBackend) (*TKLockTime, error) {
	contract, err := bindTKLockTime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TKLockTime{TKLockTimeCaller: TKLockTimeCaller{contract: contract}, TKLockTimeTransactor: TKLockTimeTransactor{contract: contract}, TKLockTimeFilterer: TKLockTimeFilterer{contract: contract}}, nil
}

// NewTKLockTimeCaller creates a new read-only instance of TKLockTime, bound to a specific deployed contract.
func NewTKLockTimeCaller(address common.Address, caller bind.ContractCaller) (*TKLockTimeCaller, error) {
	contract, err := bindTKLockTime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TKLockTimeCaller{contract: contract}, nil
}

// NewTKLockTimeTransactor creates a new write-only instance of TKLockTime, bound to a specific deployed contract.
func NewTKLockTimeTransactor(address common.Address, transactor bind.ContractTransactor) (*TKLockTimeTransactor, error) {
	contract, err := bindTKLockTime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TKLockTimeTransactor{contract: contract}, nil
}

// NewTKLockTimeFilterer creates a new log filterer instance of TKLockTime, bound to a specific deployed contract.
func NewTKLockTimeFilterer(address common.Address, filterer bind.ContractFilterer) (*TKLockTimeFilterer, error) {
	contract, err := bindTKLockTime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TKLockTimeFilterer{contract: contract}, nil
}

// bindTKLockTime binds a generic wrapper to an already deployed contract.
func bindTKLockTime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TKLockTimeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TKLockTime *TKLockTimeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TKLockTime.Contract.TKLockTimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TKLockTime *TKLockTimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TKLockTime.Contract.TKLockTimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TKLockTime *TKLockTimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TKLockTime.Contract.TKLockTimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TKLockTime *TKLockTimeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TKLockTime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TKLockTime *TKLockTimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TKLockTime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TKLockTime *TKLockTimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TKLockTime.Contract.contract.Transact(opts, method, params...)
}

// GetWallets is a free data retrieval call binding the contract method 0x422c29a4.
//
// Solidity: function getWallets(address _user) view returns(address[])
func (_TKLockTime *TKLockTimeCaller) GetWallets(opts *bind.CallOpts, _user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _TKLockTime.contract.Call(opts, &out, "getWallets", _user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWallets is a free data retrieval call binding the contract method 0x422c29a4.
//
// Solidity: function getWallets(address _user) view returns(address[])
func (_TKLockTime *TKLockTimeSession) GetWallets(_user common.Address) ([]common.Address, error) {
	return _TKLockTime.Contract.GetWallets(&_TKLockTime.CallOpts, _user)
}

// GetWallets is a free data retrieval call binding the contract method 0x422c29a4.
//
// Solidity: function getWallets(address _user) view returns(address[])
func (_TKLockTime *TKLockTimeCallerSession) GetWallets(_user common.Address) ([]common.Address, error) {
	return _TKLockTime.Contract.GetWallets(&_TKLockTime.CallOpts, _user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TKLockTime *TKLockTimeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TKLockTime.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TKLockTime *TKLockTimeSession) Owner() (common.Address, error) {
	return _TKLockTime.Contract.Owner(&_TKLockTime.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TKLockTime *TKLockTimeCallerSession) Owner() (common.Address, error) {
	return _TKLockTime.Contract.Owner(&_TKLockTime.CallOpts)
}

// NewTimeLockedWallet is a paid mutator transaction binding the contract method 0x3f106cde.
//
// Solidity: function newTimeLockedWallet(address _owner, address _tokenLock, uint256 _amount, uint256 _unlockDate) returns(address)
func (_TKLockTime *TKLockTimeTransactor) NewTimeLockedWallet(opts *bind.TransactOpts, _owner common.Address, _tokenLock common.Address, _amount *big.Int, _unlockDate *big.Int) (*types.Transaction, error) {
	return _TKLockTime.contract.Transact(opts, "newTimeLockedWallet", _owner, _tokenLock, _amount, _unlockDate)
}

// NewTimeLockedWallet is a paid mutator transaction binding the contract method 0x3f106cde.
//
// Solidity: function newTimeLockedWallet(address _owner, address _tokenLock, uint256 _amount, uint256 _unlockDate) returns(address)
func (_TKLockTime *TKLockTimeSession) NewTimeLockedWallet(_owner common.Address, _tokenLock common.Address, _amount *big.Int, _unlockDate *big.Int) (*types.Transaction, error) {
	return _TKLockTime.Contract.NewTimeLockedWallet(&_TKLockTime.TransactOpts, _owner, _tokenLock, _amount, _unlockDate)
}

// NewTimeLockedWallet is a paid mutator transaction binding the contract method 0x3f106cde.
//
// Solidity: function newTimeLockedWallet(address _owner, address _tokenLock, uint256 _amount, uint256 _unlockDate) returns(address)
func (_TKLockTime *TKLockTimeTransactorSession) NewTimeLockedWallet(_owner common.Address, _tokenLock common.Address, _amount *big.Int, _unlockDate *big.Int) (*types.Transaction, error) {
	return _TKLockTime.Contract.NewTimeLockedWallet(&_TKLockTime.TransactOpts, _owner, _tokenLock, _amount, _unlockDate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TKLockTime *TKLockTimeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TKLockTime.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TKLockTime *TKLockTimeSession) RenounceOwnership() (*types.Transaction, error) {
	return _TKLockTime.Contract.RenounceOwnership(&_TKLockTime.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TKLockTime *TKLockTimeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TKLockTime.Contract.RenounceOwnership(&_TKLockTime.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TKLockTime *TKLockTimeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TKLockTime.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TKLockTime *TKLockTimeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TKLockTime.Contract.TransferOwnership(&_TKLockTime.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TKLockTime *TKLockTimeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TKLockTime.Contract.TransferOwnership(&_TKLockTime.TransactOpts, newOwner)
}

// TKLockTimeCreatedIterator is returned from FilterCreated and is used to iterate over the raw logs and unpacked data for Created events raised by the TKLockTime contract.
type TKLockTimeCreatedIterator struct {
	Event *TKLockTimeCreated // Event containing the contract specifics and raw log

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
func (it *TKLockTimeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TKLockTimeCreated)
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
		it.Event = new(TKLockTimeCreated)
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
func (it *TKLockTimeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TKLockTimeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TKLockTimeCreated represents a Created event raised by the TKLockTime contract.
type TKLockTimeCreated struct {
	Wallet     common.Address
	To         common.Address
	CreatedAt  *big.Int
	UnlockDate *big.Int
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCreated is a free log retrieval operation binding the contract event 0x2abec6ac791844b0f5490c49d67e9eede10b60ff00b9e1180467b970e8fc8b00.
//
// Solidity: event Created(address wallet, address to, uint256 createdAt, uint256 unlockDate, uint256 amount)
func (_TKLockTime *TKLockTimeFilterer) FilterCreated(opts *bind.FilterOpts) (*TKLockTimeCreatedIterator, error) {

	logs, sub, err := _TKLockTime.contract.FilterLogs(opts, "Created")
	if err != nil {
		return nil, err
	}
	return &TKLockTimeCreatedIterator{contract: _TKLockTime.contract, event: "Created", logs: logs, sub: sub}, nil
}

// WatchCreated is a free log subscription operation binding the contract event 0x2abec6ac791844b0f5490c49d67e9eede10b60ff00b9e1180467b970e8fc8b00.
//
// Solidity: event Created(address wallet, address to, uint256 createdAt, uint256 unlockDate, uint256 amount)
func (_TKLockTime *TKLockTimeFilterer) WatchCreated(opts *bind.WatchOpts, sink chan<- *TKLockTimeCreated) (event.Subscription, error) {

	logs, sub, err := _TKLockTime.contract.WatchLogs(opts, "Created")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TKLockTimeCreated)
				if err := _TKLockTime.contract.UnpackLog(event, "Created", log); err != nil {
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

// ParseCreated is a log parse operation binding the contract event 0x2abec6ac791844b0f5490c49d67e9eede10b60ff00b9e1180467b970e8fc8b00.
//
// Solidity: event Created(address wallet, address to, uint256 createdAt, uint256 unlockDate, uint256 amount)
func (_TKLockTime *TKLockTimeFilterer) ParseCreated(log types.Log) (*TKLockTimeCreated, error) {
	event := new(TKLockTimeCreated)
	if err := _TKLockTime.contract.UnpackLog(event, "Created", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TKLockTimeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TKLockTime contract.
type TKLockTimeOwnershipTransferredIterator struct {
	Event *TKLockTimeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TKLockTimeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TKLockTimeOwnershipTransferred)
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
		it.Event = new(TKLockTimeOwnershipTransferred)
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
func (it *TKLockTimeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TKLockTimeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TKLockTimeOwnershipTransferred represents a OwnershipTransferred event raised by the TKLockTime contract.
type TKLockTimeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TKLockTime *TKLockTimeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TKLockTimeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TKLockTime.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TKLockTimeOwnershipTransferredIterator{contract: _TKLockTime.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TKLockTime *TKLockTimeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TKLockTimeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TKLockTime.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TKLockTimeOwnershipTransferred)
				if err := _TKLockTime.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TKLockTime *TKLockTimeFilterer) ParseOwnershipTransferred(log types.Log) (*TKLockTimeOwnershipTransferred, error) {
	event := new(TKLockTimeOwnershipTransferred)
	if err := _TKLockTime.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
