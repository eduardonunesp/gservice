package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Data{})
	return db
}
