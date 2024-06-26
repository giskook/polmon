package sync

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	blockFinalized = 64
)

func (s *Synchronizer) scanBlockRange(ctx context.Context, from, limit uint64) ([]types.Log, bool, error) {
	pause := false
	latestBlockNumber, err := s.Client.BlockNumber(ctx)
	if err != nil {
		return nil, false, err
	}
	latestBlockNumber -= blockFinalized
	to := from + limit
	if to > latestBlockNumber {
		to = latestBlockNumber
		pause = true
	}

	filterQuery := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(from),
		ToBlock:   new(big.Int).SetUint64(to),
		Addresses: []common.Address{s.Conf.Address},
		Topics:    make([][]common.Hash, 0),
	}
	filterQuery.Topics = append(filterQuery.Topics, []common.Hash{s.Conf.Topic1})
	filterQuery.Topics = append(filterQuery.Topics, []common.Hash{s.Conf.Topic2})

	ret, err := s.Client.FilterLogs(ctx, filterQuery)
	return ret, pause, err
}

func (s *Synchronizer) calcFee(ctx context.Context, txHash common.Hash) (uint64, error) {
	receipt, err := s.Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return 0, err
	}
	gasUsed := receipt.GasUsed
	tx, _, err := s.Client.TransactionByHash(ctx, txHash)
	gasPrice := tx.GasPrice()

	return new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gasUsed)).Uint64(), nil
}
