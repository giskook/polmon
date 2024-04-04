package sync

import "github.com/ethereum/go-ethereum/common"

type Configure struct {
	RpcURL  string
	Block   uint64
	Address common.Address
	Topic1  common.Hash
	Topic2  common.Hash
}
