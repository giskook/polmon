package sync

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/giskook/polmon/internal/persistence"
)

type Synchronizer struct {
	Conf   Configure
	Client *ethclient.Client
	Store  persistence.Persistence
	exit   chan struct{}
}

func NewSynchronizer(conf Configure, store persistence.Persistence) *Synchronizer {
	client, err := ethclient.Dial(conf.RpcURL)
	if err != nil {
		log.Fatal("NewSynchronizer ethclient.Dial: ", err)
	}
	return &Synchronizer{
		Conf:   conf,
		Client: client,
		Store:  store,
		exit:   make(chan struct{}),
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

func (s *Synchronizer) getBeginBlock() uint64 {
	block, err := s.Store.GetLatestPersistedBlock()
	if err != nil {
		block = 0
	}
	return max(s.Conf.Block, block+1) // 1 for next block
}

func (s *Synchronizer) store() error {
	beginBlock := s.getBeginBlock()
	ctx := context.Background()
	logs, err := s.scanBlockRange(ctx, beginBlock, 100)
	log.Println("syncing block: ", beginBlock, " to ", beginBlock+100)
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
	return nil
}
