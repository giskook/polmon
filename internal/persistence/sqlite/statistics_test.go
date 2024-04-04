package sqlite

import (
	"testing"
)

func TestPersistence_GetLatestStatisticsBlock(t *testing.T) {
	p := NewPersistence(Configure{Path: "/Users/oker/workspace/polmon/fee.db"})
	t.Log(p.GetLatestStatisticsBlock())
}
