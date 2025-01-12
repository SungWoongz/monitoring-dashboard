package models

import (
	v "server/variables"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func init() {
	var err error
	Db, err = gorm.Open(sqlite.Open(v.DATABASE), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Db.AutoMigrate(&CPU_Status{})
}
