package statistics

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/log"
	"github.com/giskook/polmon/internal/persistence"
)

type Statistics struct {
	Conf        Configure
	persistence persistence.Persistence
	exit        chan struct{}
}

func NewStatistics(conf Configure, store persistence.Persistence) *Statistics {
	return &Statistics{
		Conf:        conf,
		exit:        make(chan struct{}),
		persistence: store,
	}
}

func (s *Statistics) Start() {
	for {
		select {
		case <-s.exit:
			log.Info("Statistics exit")
			return
		case <-time.After(s.Conf.Internal):
			latestStatisticsBlock, err := s.persistence.GetLatestStatisticsBlock()
			if err != nil {
				log.Error("statistics GetLatestStatisticsBlock ", "err", err)
				latestStatisticsBlock = 0
			}
			unStatisticsBlock, fee, txHash, err := s.persistence.GetUnStatisticsBlock(latestStatisticsBlock)
			if err != nil {
				log.Error("statistics GetUnStatisticsBlock ", "err", err)
			}
			if latestStatisticsBlock == unStatisticsBlock || unStatisticsBlock == 0 {
				continue
			}
			totalFee, err := s.persistence.GetTotalFee()
			if err != nil {
				log.Info("statistics GetTotalFee ", "err", err)
				continue
			}
			totalFeeInt := new(big.Int).SetUint64(0)
			if totalFee != "" {
				totalFeeInt, _ = new(big.Int).SetString(totalFee, 10)
			}
			feeInt := new(big.Int).SetUint64(fee)
			totalFeeInt.Add(totalFeeInt, feeInt)
			s.persistence.SaveStatistics(unStatisticsBlock, txHash, fee, totalFeeInt.String())
		}
	}
}

func (s *Statistics) Stop() {
	close(s.exit)
}
