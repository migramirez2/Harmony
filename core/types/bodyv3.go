package types

import (
	"io"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/harmony-one/harmony/block"
	staking "github.com/harmony-one/harmony/staking/types"
)

// BodyV3 is the V3 block body
type BodyV3 struct {
	f bodyFieldsV3
}

type bodyFieldsV3 struct {
	EthTransactions     []*EthTransaction
	Transactions        []*Transaction
	StakingTransactions []*staking.StakingTransaction
	Uncles              []*block.Header
	IncomingReceipts    CXReceiptsProofs
}

// Transactions returns the list of transactions.
//
// The returned list is a deep copy; the caller may do anything with it without
// affecting the original.
func (b *BodyV3) Transactions() (txs []*Transaction) {
	for _, tx := range b.f.Transactions {
		txs = append(txs, tx.Copy())
	}
	return txs
}

// StakingTransactions returns the list of staking transactions.
// The returned list is a deep copy; the caller may do anything with it without
// affecting the original.
func (b *BodyV3) StakingTransactions() (txs []*staking.StakingTransaction) {
	for _, tx := range b.f.StakingTransactions {
		txs = append(txs, tx.Copy())
	}
	return txs
}

// TransactionAt returns the transaction at the given index in this block.
// It returns nil if index is out of bounds.
func (b *BodyV3) TransactionAt(index int) *Transaction {
	if index < 0 || index >= len(b.f.Transactions) {
		return nil
	}
	return b.f.Transactions[index].Copy()
}

// StakingTransactionAt returns the staking transaction at the given index in this block.
// It returns nil if index is out of bounds.
func (b *BodyV3) StakingTransactionAt(index int) *staking.StakingTransaction {
	if index < 0 || index >= len(b.f.StakingTransactions) {
		return nil
	}
	return b.f.StakingTransactions[index].Copy()
}

// CXReceiptAt returns the CXReceipt at given index in this block
// It returns nil if index is out of bounds
func (b *BodyV3) CXReceiptAt(index int) *CXReceipt {
	if index < 0 {
		return nil
	}
	for _, cxp := range b.f.IncomingReceipts {
		cxs := cxp.Receipts
		if index < len(cxs) {
			return cxs[index].Copy()
		}
		index -= len(cxs)
	}
	return nil
}

// SetTransactions sets the list of transactions with a deep copy of the given
// list.
func (b *BodyV3) SetTransactions(newTransactions []*Transaction) {
	var txs []*Transaction
	for _, tx := range newTransactions {
		txs = append(txs, tx.Copy())
	}
	b.f.Transactions = txs
}

// SetStakingTransactions sets the list of staking transactions with a deep copy of the given
// list.
func (b *BodyV3) SetStakingTransactions(newStakingTransactions []*staking.StakingTransaction) {
	var txs []*staking.StakingTransaction
	for _, tx := range newStakingTransactions {
		txs = append(txs, tx.Copy())
	}
	b.f.StakingTransactions = txs
}

// Uncles returns a deep copy of the list of uncle headers of this block.
func (b *BodyV3) Uncles() (uncles []*block.Header) {
	for _, uncle := range b.f.Uncles {
		uncles = append(uncles, CopyHeader(uncle))
	}
	return uncles
}

// SetUncles sets the list of uncle headers with a deep copy of the given list.
func (b *BodyV3) SetUncles(newUncle []*block.Header) {
	var uncles []*block.Header
	for _, uncle := range newUncle {
		uncles = append(uncles, CopyHeader(uncle))
	}
	b.f.Uncles = uncles
}

// IncomingReceipts returns a deep copy of the list of incoming cross-shard
// transaction receipts of this block.
func (b *BodyV3) IncomingReceipts() (incomingReceipts CXReceiptsProofs) {
	return b.f.IncomingReceipts.Copy()
}

// SetIncomingReceipts sets the list of incoming cross-shard transaction
// receipts of this block with a dep copy of the given list.
func (b *BodyV3) SetIncomingReceipts(newIncomingReceipts CXReceiptsProofs) {
	b.f.IncomingReceipts = newIncomingReceipts.Copy()
}

// EncodeRLP RLP-encodes the block body into the given writer.
func (b *BodyV3) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &b.f)
}

// DecodeRLP RLP-decodes a block body from the given RLP stream into the
// receiver.
func (b *BodyV3) DecodeRLP(s *rlp.Stream) error {
	return s.Decode(&b.f)
}

// EthTransactions returns the list of transactions that's ethereum-compatible.
//
// The returned list is a deep copy; the caller may do anything with it without
// affecting the original.
func (b *BodyV3) EthTransactions() (txs []*EthTransaction) {
	for _, tx := range b.f.EthTransactions {
		txs = append(txs, tx.Copy())
	}
	return txs
}

// EthTransactionAt returns the ethereum-compatible transaction at the given index in this block.
// It returns nil if index is out of bounds.
func (b *BodyV3) EthTransactionAt(index int) *EthTransaction {
	if index < 0 || index >= len(b.f.EthTransactions) {
		return nil
	}
	return b.f.EthTransactions[index].Copy()
}

// SetEthTransactions sets the list of ethereum-compatible transactions with a deep copy of the given
// list.
func (b *BodyV3) SetEthTransactions(newTransactions []*EthTransaction) {
	var txs []*EthTransaction
	for _, tx := range newTransactions {
		txs = append(txs, tx.Copy())
	}
	b.f.EthTransactions = txs
}
