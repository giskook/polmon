package persistence

type Persistence interface {
	Save(blockNum int, fee string)
	GetTotalFee() string
}
