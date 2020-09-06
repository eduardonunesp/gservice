package utils

import (
	"gservice/models"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetTestDB() *gorm.DB {
	parentDir := os.TempDir()
	dirName := uuid.New().String()
	tmpDir, err := ioutil.TempDir(parentDir, dirName)

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open(tmpDir+"/gdatabase_test.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.PostData{})
	return db
}
