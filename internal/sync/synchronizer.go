package sync

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/giskook/polmon/internal/persistence"
)

const (
	blockRange = 100
)

type Synchronizer struct {
	Conf       Configure
	Client     *ethclient.Client
	Store      persistence.Persistence
	exit       chan struct{}
	beginBlock uint64
}

func NewSynchronizer(conf Configure, store persistence.Persistence) *Synchronizer {
	client, err := ethclient.Dial(conf.RpcURL)
	if err != nil {
		log.Fatal("NewSynchronizer ethclient.Dial: ", err)
	}
	return &Synchronizer{
		Conf:       conf,
		Client:     client,
		Store:      store,
		beginBlock: conf.Block,
		exit:       make(chan struct{}),
	}
}

func (s *Synchronizer) Start() {
	for {
		select {
		case <-s.exit:
			log.Println("Synchronizer exit")
			return
		default:
			if err := s.store(); err != nil {
				log.Println("Synchronizer.store: ", err)
			}
		}
	}
}

func (s *Synchronizer) Stop() {
	close(s.exit)
}

func (s *Synchronizer) getBeginBlock(beginBlock uint64) uint64 {
	dbBlock, err := s.Store.GetLatestPersistedBlock()
	if err != nil {
		dbBlock = 0
	}

	return max(beginBlock, dbBlock+1) // 1 for next block
}

func (s *Synchronizer) store() error {
	beginBlock := s.getBeginBlock(s.beginBlock)
	s.beginBlock = beginBlock
	ctx := context.Background()
	logs, pause, err := s.scanBlockRange(ctx, beginBlock, blockRange)
	log.Println("syncing block: ", beginBlock, " to ", beginBlock+blockRange)
	if err != nil {
		return err
	}
	for _, log := range logs {
		fee, err := s.calcFee(ctx, log.TxHash)
		if err != nil {
			return err
		}
		if err := s.Store.SaveSync(log.BlockNumber, log.TxHash.String(), fee); err != nil {
			return err
		}
	}
	s.beginBlock += blockRange + 1
	if pause {
		time.Sleep(5 * time.Minute)
	}

	return nil
}
