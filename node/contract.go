package node

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/harmony-one/harmony/core/types"
	"github.com/harmony-one/harmony/internal/utils/contract"
)

// Constants related to smart contract.
const (
	FaucetContractBinary      = "0x6080604052678ac7230489e8000060015560028054600160a060020a031916331790556101aa806100316000396000f3fe608060405260043610610045577c0100000000000000000000000000000000000000000000000000000000600035046327c78c42811461004a5780634ddd108a1461008c575b600080fd5b34801561005657600080fd5b5061008a6004803603602081101561006d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166100b3565b005b34801561009857600080fd5b506100a1610179565b60408051918252519081900360200190f35b60025473ffffffffffffffffffffffffffffffffffffffff1633146100d757600080fd5b600154303110156100e757600080fd5b73ffffffffffffffffffffffffffffffffffffffff811660009081526020819052604090205460ff161561011a57600080fd5b73ffffffffffffffffffffffffffffffffffffffff8116600081815260208190526040808220805460ff1916600190811790915554905181156108fc0292818181858888f19350505050158015610175573d6000803e3d6000fd5b5050565b30319056fea165627a7a723058203e799228fee2fa7c5d15e71c04267a0cc2687c5eff3b48b98f21f355e1064ab30029"
	FaucetContractFund        = 8000000
	FaucetFreeMoneyMethodCall = "0x27c78c42000000000000000000000000"

	StakingContractBinary = "0x608060405234801561001057600080fd5b50610b51806100206000396000f3fe608060405260043610610072576000357c01000000000000000000000000000000000000000000000000000000009004806325ca4c9c146100775780632e1a7d4d146100e05780634c1b64cb1461012f578063a98e4e7714610194578063d0e30db0146101bf578063e27fd057146101dd575b600080fd5b34801561008357600080fd5b506100c66004803603602081101561009a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610249565b604051808215151515815260200191505060405180910390f35b3480156100ec57600080fd5b506101196004803603602081101561010357600080fd5b8101908080359060200190929190505050610310565b6040518082815260200191505060405180910390f35b34801561013b57600080fd5b5061017e6004803603602081101561015257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506104cb565b6040518082815260200191505060405180910390f35b3480156101a057600080fd5b506101a96106f3565b6040518082815260200191505060405180910390f35b6101c7610700565b6040518082815260200191505060405180910390f35b3480156101e957600080fd5b506101f2610a46565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561023557808201518184015260208101905061021a565b505050509050019250505060405180910390f35b6000806002805490501415610261576000905061030b565b8173ffffffffffffffffffffffffffffffffffffffff166002600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548154811015156102c657fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161490505b919050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211151561048457816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055503373ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156103eb573d6000803e3d6000fd5b5060008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054141561043e5761043c336104cb565b505b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506104c6565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490505b919050565b600080600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000600260016002805490500381548110151561052957fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508060028381548110151561056657fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555060028054809190600190036106079190610ad4565b508373ffffffffffffffffffffffffffffffffffffffff167e1fab73a76dc2de66330e055b1c1e3319c77b736bb4478cc706497f318a4ad7836040518082815260200191505060405180910390a28073ffffffffffffffffffffffffffffffffffffffff167f6095abd20e12b7e743432b409b7879ac77a0b927f89ae330f59c15b32dce0b69836000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808381526020018281526020019250505060405180910390a28192505050919050565b6000600280549050905090565b600061070b33610249565b1561087657346000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055503373ffffffffffffffffffffffffffffffffffffffff167f6095abd20e12b7e743432b409b7879ac77a0b927f89ae330f59c15b32dce0b69600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808381526020018281526020019250505060405180910390a2600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050610a43565b346000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600160023390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555003600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff167fd2ad617bb539c9a6219058035b15d87478e478eb0f74164eae890a0c70fa3f40600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054604051808381526020018281526020019250505060405180910390a260016002805490500390505b90565b60606002805480602002602001604051908101604052809291908181526020018280548015610aca57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610a80575b5050505050905090565b815481835581811115610afb57818360005260206000209182019101610afa9190610b00565b5b505050565b610b2291905b80821115610b1e576000816000905550600101610b06565b5090565b9056fea165627a7a7230582032eb9f748231d6ef2fd0ac6bd603cc3e381b6a6596e87beb1f6ce117e5237f4b0029"
)

// AddStakingContractToPendingTransactions adds the deposit smart contract the genesis block.
func (node *Node) AddStakingContractToPendingTransactions() {
	// Add a contract deployment transaction
	//Generate contract key and associate funds with the smart contract
	priKey, _ := ecdsa.GenerateKey(crypto.S256(), strings.NewReader("Deposit Smart Contract Key"))
	contractAddress := crypto.PubkeyToAddress(priKey.PublicKey)
	//Initially the smart contract should have minimal funds.
	contractFunds := big.NewInt(0)
	contractFunds = contractFunds.Mul(contractFunds, big.NewInt(params.Ether))
	dataEnc := common.FromHex(StakingContractBinary)
	// Unsigned transaction to avoid the case of transaction address.
	mycontracttx, _ := types.SignTx(types.NewContractCreation(uint64(0), node.Consensus.ShardID, contractFunds, params.TxGasContractCreation*10, nil, dataEnc), types.HomesteadSigner{}, priKey)
	node.ContractAddresses = append(node.ContractAddresses, crypto.CreateAddress(contractAddress, uint64(0)))
	node.addPendingTransactions(types.Transactions{mycontracttx})
}

// CreateStakingDepositTransaction creates a new deposit staking transaction
func (node *Node) CreateStakingDepositTransaction(stake int) (*types.Transaction, error) {
	//These should be read from somewhere.
	DepositContractPriKey, _ := ecdsa.GenerateKey(crypto.S256(), strings.NewReader("Deposit Smart Contract Key")) //DepositContractPriKey is pk for contract
	DepositContractAddress := crypto.PubkeyToAddress(DepositContractPriKey.PublicKey)                             //DepositContractAddress is the address for the contract
	state, err := node.blockchain.State()
	if err != nil {
		log.Error("Failed to get chain state", "Error", err)
	}
	nonce := state.GetNonce(crypto.PubkeyToAddress(DepositContractPriKey.PublicKey))
	callingFunction := "0xd0e30db0"
	dataEnc := common.FromHex(callingFunction) //Deposit Does not take a argument, stake is transferred via amount.
	tx, err := types.SignTx(types.NewTransaction(nonce, DepositContractAddress, node.Consensus.ShardID, big.NewInt(int64(stake)), params.TxGasContractCreation*10, nil, dataEnc), types.HomesteadSigner{}, node.AccountKey)
	return tx, err
}

//CreateStakingWithdrawTransaction creates a new withdraw stake transaction
func (node *Node) CreateStakingWithdrawTransaction(stake int) (*types.Transaction, error) {
	//These should be read from somewhere.
	DepositContractPriKey, _ := ecdsa.GenerateKey(crypto.S256(), strings.NewReader("Deposit Smart Contract Key")) //DepositContractPriKey is pk for contract
	DepositContractAddress := crypto.PubkeyToAddress(DepositContractPriKey.PublicKey)                             //DepositContractAddress is the address for the contract
	state, err := node.blockchain.State()
	if err != nil {
		log.Error("Failed to get chain state", "Error", err)
	}
	nonce := state.GetNonce(crypto.PubkeyToAddress(DepositContractPriKey.PublicKey))
	callingFunction := "0x2e1a7d4d"
	contractData := callingFunction + hex.EncodeToString([]byte(strconv.Itoa(stake)))
	dataEnc := common.FromHex(contractData)
	tx, err := types.SignTx(types.NewTransaction(nonce, DepositContractAddress, node.Consensus.ShardID, big.NewInt(0), params.TxGasContractCreation*10, nil, dataEnc), types.HomesteadSigner{}, node.AccountKey)
	return tx, err
}

// AddFaucetContractToPendingTransactions adds the faucet contract the genesis block.
func (node *Node) AddFaucetContractToPendingTransactions() {
	// Add a contract deployment transactionv
	priKey := node.ContractKeys[0]
	dataEnc := common.FromHex(FaucetContractBinary)
	// Unsigned transaction to avoid the case of transaction address.

	contractFunds := big.NewInt(FaucetContractFund)
	contractFunds = contractFunds.Mul(contractFunds, big.NewInt(params.Ether))
	mycontracttx, _ := types.SignTx(
		types.NewContractCreation(uint64(0), node.Consensus.ShardID, contractFunds, params.TxGasContractCreation*10, nil, dataEnc),
		types.HomesteadSigner{},
		priKey)
	node.ContractAddresses = append(node.ContractAddresses, crypto.CreateAddress(crypto.PubkeyToAddress(priKey.PublicKey), uint64(0)))
	node.addPendingTransactions(types.Transactions{mycontracttx})
}

// CallFaucetContract invokes the faucet contract to give the walletAddress initial money
func (node *Node) CallFaucetContract(walletAddress common.Address) common.Hash {
	return node.createSendingMoneyTransaction(walletAddress)
}

func (node *Node) createSendingMoneyTransaction(walletAddress common.Address) common.Hash {
	state, err := node.blockchain.State()
	if err != nil {
		log.Error("Failed to get chain state", "Error", err)
	}
	nonce := state.GetNonce(crypto.PubkeyToAddress(node.ContractKeys[0].PublicKey))
	contractData := FaucetFreeMoneyMethodCall + hex.EncodeToString(walletAddress.Bytes())
	dataEnc := common.FromHex(contractData)
	tx, _ := types.SignTx(types.NewTransaction(nonce, node.ContractAddresses[0], node.Consensus.ShardID, big.NewInt(0), params.TxGasContractCreation*10, nil, dataEnc), types.HomesteadSigner{}, node.ContractKeys[0])

	node.addPendingTransactions(types.Transactions{tx})
	return tx.Hash()
}

// DepositToFakeAccounts invokes the faucet contract to give the walletAddress initial money
func (node *Node) DepositToFakeAccounts() {
	for _, deployAccount := range contract.FakeAccounts {
		address := common.HexToAddress(deployAccount.Address)
		node.createSendingMoneyTransaction(address)
	}
}
