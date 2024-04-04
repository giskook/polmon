package sqlite

import (
	"gorm.io/gorm"
)

type Sync struct {
	gorm.Model
	BlockNum uint64
	TxHash   string
	Fee      uint64
}

func (p *Persistence) SaveSync(blockNum uint64, txHash string, fee uint64) error {
	p.db.Create(&Sync{BlockNum: blockNum, TxHash: txHash, Fee: fee})
	return nil
}

func (p *Persistence) GetLatestPersistedBlock() (uint64, error) {
	var syncedBlock Sync
	p.db.Last(&syncedBlock, 1)

	return syncedBlock.BlockNum, nil
}

func (p *Persistence) GetUnStatisticsBlock(blockNumber uint64) (uint64, uint64, string, error) {
	var syncedBlock Sync
	p.db.Where("block_num > ?", blockNumber).First(&syncedBlock)

	return syncedBlock.BlockNum, syncedBlock.Fee, syncedBlock.TxHash, nil
}
