package repos

import (
	"time"

	"github.com/eduardonunesp/gservice/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DataRepo interface {
	Select() ([]models.Data, error)
	SelectByName(name string) (models.Data, error)
	SelectById(id string) (models.Data, error)
	Insert(name string, stage, score int) (models.Data, error)
}

type dataRepo struct {
	DB *gorm.DB
}

func NewDataRepo(db *gorm.DB) DataRepo {
	return &dataRepo{DB: db}
}

func (m *dataRepo) Select() ([]models.Data, error) {
	result := []models.Data{}
	err := m.DB.Find(&result).Error
	return result, err
}

func (m *dataRepo) SelectByName(name string) (models.Data, error) {
	result := models.Data{}
	err := m.DB.Where("name = ?", name).First(&result).Error
	return result, err
}

func (m *dataRepo) SelectById(name string) (models.Data, error) {
	result := models.Data{}
	err := m.DB.Where("uuid = ?", name).First(&result).Error
	return result, err
}

func (m *dataRepo) Insert(name string, stage, score int) (models.Data, error) {
	data := models.Data{
		Name:          name,
		Stage:         stage,
		Score:         score,
		UUID4:         uuid.New().String(),
		UnixTimestamp: time.Now().UTC().Unix(),
	}

	if err := m.DB.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
