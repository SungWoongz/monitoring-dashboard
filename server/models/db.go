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
<<<<<<< HEAD
	Db.AutoMigrate(&CPU_Status{}, &MEM_Status{})
=======
	Db.AutoMigrate(&CPU_Status{})
>>>>>>> 97e9e548f2ce395e9ae5a36f2f4c41359c1607b7
}
