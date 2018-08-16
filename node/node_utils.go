package node

import (
	"github.com/simple-rules/harmony-benchmark/blockchain"
	"github.com/simple-rules/harmony-benchmark/crypto/pki"
)

// AddTestingAddresses creates in genesis block numAddress transactions which assign 1000 token to each address in [0 - numAddress)
// This is used by client code.
// TODO: Consider to remove it later when moving to production.
func (node *Node) AddTestingAddresses(numAddress int) {
	txs := make([]*blockchain.Transaction, numAddress)
	for i := range txs {
		txs[i] = blockchain.NewCoinbaseTX(pki.GetAddressFromInt(i+1), "", node.Consensus.ShardID)
	}
	node.blockchain.Blocks[0].Transactions = append(node.blockchain.Blocks[0].Transactions, txs...)
	node.UtxoPool.Update(txs)
}
