package sqlite

import "gorm.io/gorm"

type Statistics struct {
	gorm.Model
	BlockNum uint64 `gorm:"uniqueIndex"`
	TxHash   string
	Fee      uint64
	Total    string
}

func (p *Persistence) GetLatestStatisticsBlock() (uint64, error) {
	var statistics Statistics
	p.db.Last(&statistics, nil)

	return statistics.BlockNum, nil
}

func (p *Persistence) GetTotalFee() (string, error) {
	var statistics Statistics
	p.db.Last(&statistics, nil)

	return statistics.Total, nil
}

func (p *Persistence) SaveStatistics(blockNum uint64, txHash string, fee uint64, total string) error {
	p.db.Create(&Statistics{BlockNum: blockNum, TxHash: txHash, Fee: fee, Total: total})
	return nil
}
