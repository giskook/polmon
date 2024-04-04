package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&Sync{})
	if err != nil {
		panic("failed to migrate")
	}
}

type Persistence struct {
	conf Configure
	db   *gorm.DB
}

func NewPersistence(conf Configure) *Persistence {
	db, err := gorm.Open(sqlite.Open(conf.Path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &Persistence{conf: conf, db: db}
}
