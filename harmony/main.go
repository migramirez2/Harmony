package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	"github.com/harmony-one/harmony/consensus"
	"github.com/harmony-one/harmony/core"
	"github.com/harmony-one/harmony/core/types"
	"github.com/harmony-one/harmony/core/vm"
	"github.com/harmony-one/harmony/db"
	"math/big"
)

var (

	// Test accounts
	testBankKey, _  = crypto.GenerateKey()
	testBankAddress = crypto.PubkeyToAddress(testBankKey.PublicKey)
	testBankFunds   = big.NewInt(1000000000000000000)

	testUserKey, _  = crypto.GenerateKey()
	testUserAddress = crypto.PubkeyToAddress(testUserKey.PublicKey)

	chainConfig = params.TestChainConfig

	// Test transactions
	pendingTxs []*types.Transaction
	newTxs     []*types.Transaction
)

func init() {
	tx1, _ := types.SignTx(types.NewTransaction(0, testUserAddress, big.NewInt(1000), params.TxGas, nil, nil), types.HomesteadSigner{}, testBankKey)
	tx2, _ := types.SignTx(types.NewTransaction(1, testUserAddress, big.NewInt(1000), params.TxGas, nil, nil), types.HomesteadSigner{}, testBankKey)
	tx3, _ := types.SignTx(types.NewTransaction(2, testUserAddress, big.NewInt(1000), params.TxGas, nil, nil), types.HomesteadSigner{}, testBankKey)
	pendingTxs = append(pendingTxs, tx1)
	pendingTxs = append(pendingTxs, tx2)
	pendingTxs = append(pendingTxs, tx3)
	tx4, _ := types.SignTx(types.NewTransaction(1, testUserAddress, big.NewInt(1000), params.TxGas, nil, nil), types.HomesteadSigner{}, testBankKey)
	newTxs = append(newTxs, tx4)
}

type testWorkerBackend struct {
	db     db.Database
	txPool *core.TxPool
	chain  *core.BlockChain
}

func main() {
	var (
		database = db.NewMemDatabase()
		gspec    = core.Genesis{
			Config: chainConfig,
			Alloc:  core.GenesisAlloc{testBankAddress: {Balance: testBankFunds}},
		}
	)

	genesis := gspec.MustCommit(database)

	chain, _ := core.NewBlockChain(database, nil, gspec.Config, consensus.NewFaker(), vm.Config{}, nil)

	fmt.Println(testBankAddress)
	fmt.Println(testUserAddress)
	fmt.Println(genesis.Root())

	txpool := core.NewTxPool(core.DefaultTxPoolConfig, chainConfig, chain)

	backend := &testWorkerBackend{
		db:     database,
		chain:  chain,
		txPool: txpool,
	}
	backend.txPool.AddLocals(pendingTxs)

	// Generate a small n-block chain and an uncle block for it
	n := 3
	if n > 0 {
		blocks, _ := core.GenerateChain(chainConfig, genesis, consensus.NewFaker(), database, n, func(i int, gen *core.BlockGen) {
			gen.SetCoinbase(testBankAddress)
			gen.AddTx(pendingTxs[i])
		})
		fmt.Println("Printing blocks 1")
		fmt.Println(blocks[0].Root())
		fmt.Println(blocks[1].Root())
		fmt.Println(blocks[2].Root())
		if _, err := chain.InsertChain(blocks); err != nil {
			fmt.Errorf("failed to insert origin chain: %v", err)
		}
		fmt.Println("Printing blocks 2")
		fmt.Println(chain.GetBlockByNumber(0).Root())
		fmt.Println(chain.GetBlockByNumber(0).Coinbase())
		fmt.Println(chain.GetBlockByNumber(0).Header())
		fmt.Println(chain.GetBlockByNumber(0).Transactions())
		fmt.Println(chain.GetBlockByNumber(0).ParentHash())
		fmt.Println(chain.GetBlockByNumber(0))
		fmt.Println(genesis.Root())
		fmt.Println(chain.GetBlockByNumber(1).Root())
		fmt.Println(chain.GetBlockByNumber(1).Coinbase())
		fmt.Println(chain.GetBlockByNumber(1).Transactions()[0].Value())
		fmt.Println(blocks[0].Root())
		fmt.Println(chain.GetBlockByNumber(2).Root())
		fmt.Println(blocks[1].Root())
		fmt.Println(chain.GetBlockByNumber(3).Root())
		fmt.Println(blocks[2].Root())
		fmt.Println("Yeah")
	}
}
