package sync

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// https://etherscan.io/tx/0xc50c4ce455872168cb4d60d6ee10621dcbc113b19b07190239f371aa2d4ceca0
func TestSynchronizer_scanBlockRange(t *testing.T) {
	const blockNumber = 19580368
	conf := Configure{
		RpcURL:  "https://mainnet.infura.io/v3/0d081e04621c4c69b12649fbb63ef18a",
		Block:   blockNumber,
		Address: common.HexToAddress("0x5132A183E9F3CB7C848b0AAC5Ae0c4f0491B7aB2"),
		Topic1:  common.HexToHash("0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3"),
		Topic2:  common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000003"),
	}

	synchronizer := NewSynchronizer(conf, nil)
	synchronizer.Start()
	ctx := context.Background()

	hash, err := synchronizer.scanBlockRange(ctx, blockNumber, 1)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(hash)
	}
}

func TestCalcFee(t *testing.T) {
	const blockNumber = 19580368
	conf := Configure{
		RpcURL:  "https://mainnet.infura.io/v3/0d081e04621c4c69b12649fbb63ef18a",
		Block:   blockNumber,
		Address: common.HexToAddress("0x5132A183E9F3CB7C848b0AAC5Ae0c4f0491B7aB2"),
		Topic1:  common.HexToHash("0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3"),
		Topic2:  common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000003"),
	}

	synchronizer := NewSynchronizer(conf, nil)
	ctx := context.Background()
	fee, err := synchronizer.calcFee(ctx, common.HexToHash("0xc50c4ce455872168cb4d60d6ee10621dcbc113b19b07190239f371aa2d4ceca0"))
	if err != nil {
		t.Log(err)
	}
	t.Log(fee)
}
