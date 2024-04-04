package persistence

type Persistence interface {
	// Save block number and fee to sync table
	SaveSync(blockNum uint64, txHash string, fee uint64) error
	// GetLatestPersistedBlock returns the latest block number persisted from sync table
	GetLatestPersistedBlock() (uint64, error)
	// GetUnStatisticsBlock returns the block number that has not been processed from sync table
	// blockNumber is the latest block number persisted from statistics table
	GetUnStatisticsBlock(blockNumber uint64) (uint64, uint64, string, error)
	// GetLatestStatisticsBlock returns the latest block number persisted from statistics table
	GetLatestStatisticsBlock() (uint64, error)
	// GetTotalFee returns the total fee from statistic table
	GetTotalFee() (string, error)

	// SaveStatistics saves the block number, transaction hash and fee to statistics table
	SaveStatistics(blockNum uint64, txHash string, fee uint64, total string) error
}
