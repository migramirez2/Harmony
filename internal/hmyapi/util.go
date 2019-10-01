package hmyapi

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/harmony-one/harmony/core/types"
	common2 "github.com/harmony-one/harmony/internal/common"
	"github.com/harmony-one/harmony/internal/utils"
)

// defaultOffset is to have default pagination.
const (
	defaultPageSize = 100
)

// ReturnWithPagination returns result with pagination (offset, page in TxHistoryArgs).
func ReturnWithPagination(hashes []common.Hash, args TxHistoryArgs) []common.Hash {
	pageSize := defaultPageSize
	pageIndex := args.PageIndex
	if args.PageSize > 0 {
		pageSize = args.PageSize
	}
	if pageSize*pageIndex >= len(hashes) {
		return make([]common.Hash, 0)
	}
	if pageSize*pageIndex+pageSize > len(hashes) {
		return hashes[pageSize*pageIndex:]
	}
	return hashes[pageSize*pageIndex : pageSize*pageIndex+pageSize]
}

// SubmitTransaction is a helper function that submits tx to txPool and logs a message.
func SubmitTransaction(ctx context.Context, b Backend, tx *types.Transaction) (common.Hash, error) {
	if err := b.SendTx(ctx, tx); err != nil {
		return common.Hash{}, err
	}
	if tx.To() == nil {
		signer := types.MakeSigner(b.ChainConfig(), b.CurrentBlock().Epoch())
		from, err := types.Sender(signer, tx)
		if err != nil {
			return common.Hash{}, err
		}
		addr := crypto.CreateAddress(from, tx.Nonce())
		utils.Logger().Info().
			Str("fullhash", tx.Hash().Hex()).
			Str("contract", common2.MustAddressToBech32(addr)).
			Msg("Submitted contract creation")
	} else {
		utils.Logger().Info().
			Str("fullhash", tx.Hash().Hex()).
			Str("recipient", tx.To().Hex()).
			Msg("Submitted transaction")
	}
	return tx.Hash(), nil
}
