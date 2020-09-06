package repos

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/eduardonunesp/gservice/models"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getTestDB() *gorm.DB {
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

	db.AutoMigrate(&models.Data{})
	return db
}

func TestSelect(t *testing.T) {
	db := models.GetTestDB()
	repo := NewDataRepo(db)

	// Testing select
	result, err := repo.Select()

	if err != nil {
		t.Error(err)
		return
	}

	if len(result) != 0 {
		t.Error("Result should be 0")
	}

	title := "Title Test"

	// Inserting register on db to test
	db.Create(models.Data{
		Title:         title,
		UUID4:         uuid.New().String(),
		UnixTimestamp: time.Now().UTC().Unix(),
	})

	// Testing if select returns the register
	results, err := repo.Select()

	if len(results) == 0 {
		t.Error("Result should greater than 0")
	}

	// Test if find the register by title
	_, err = repo.SelectByTitle(title)

	if err != nil {
		t.Error(err)
	}

	// Test not found the register by title
	_, err = repo.SelectByTitle("Some Not Inserted Title")

	if err == nil {
		t.Error("Title should be not found")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("Error shoudl be not found")
	}
}

func TestInsert(t *testing.T) {
	db := models.GetTestDB()
	repo := NewDataRepo(db)

	title := "Title Test"

	// Testing insert
	err := repo.Insert(title)

	if err != nil {
		t.Error(err)
		return
	}

	// Check if has inserted
	var result []models.Data
	err = db.Find(&result).Error

	if err != nil {
		t.Error(err)
		return
	}

	if len(result) == 0 {
		t.Error("Result should greater than 0")
	}

	// Testing unique register
	err = repo.Insert(title)

	if err == nil {
		t.Error("Expecting error of unique")
	}
}
