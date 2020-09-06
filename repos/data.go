package repos

import (
	"time"

	"github.com/eduardonunesp/gservice/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DataRepo interface {
	Select() ([]models.Data, error)
	SelectByTitle(title string) (models.Data, error)
	Insert(title string) error
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

func (m *dataRepo) SelectByTitle(title string) (models.Data, error) {
	result := models.Data{}
	err := m.DB.Where("title = ?", title).First(&result).Error
	return result, err
}

func (m *dataRepo) Insert(title string) error {
	data := models.Data{
		Title:         title,
		UUID4:         uuid.New().String(),
		UnixTimestamp: time.Now().UTC().Unix(),
	}

	return m.DB.Create(&data).Error
}
