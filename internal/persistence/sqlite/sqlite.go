package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Persistence struct {
	conf Configure
	db   *gorm.DB
}

func NewPersistence(conf Configure) *Persistence {
	db, err := gorm.Open(sqlite.Open(conf.Path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Sync{})
	if err != nil {
		panic("failed to migrate")

	}
	err = db.AutoMigrate(&Statistics{})
	if err != nil {
		panic("failed to migrate")
	}

	return &Persistence{conf: conf, db: db}
}
