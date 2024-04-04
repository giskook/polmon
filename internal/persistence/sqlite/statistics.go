package sqlite

import "gorm.io/gorm"

type Statistics struct {
	gorm.Model
	BlockNum uint64
	TxHash   string
	Fee      string
}

func (p *Persistence) GetLatestStatisticsBlock() (uint64, error) {
	var statistics Statistics
	p.db.Last(&statistics, 1)

	return statistics.BlockNum, nil
}

func (p *Persistence) GetTotalFee() (string, error) {
	var statistics Statistics
	p.db.Last(&statistics, 1)

	return statistics.Fee, nil
}

func (p *Persistence) SaveStatistics(blockNum uint64, txHash string, fee string) error {
	p.db.Create(&Statistics{BlockNum: blockNum, TxHash: txHash, Fee: fee})
	return nil
}
