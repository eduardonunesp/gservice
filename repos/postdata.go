package repos

import (
	"time"

	"github.com/eduardonunesp/gservice/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostDataRepo interface {
	Select() ([]models.PostData, error)
	SelectByTitle(title string) (models.PostData, error)
	Insert(title string) error
}

type postDataRepo struct {
	DB *gorm.DB
}

func NewPostDataRepo(db *gorm.DB) PostDataRepo {
	return &postDataRepo{DB: db}
}

func (m *postDataRepo) Select() ([]models.PostData, error) {
	result := []models.PostData{}
	err := m.DB.Find(&result).Error
	return result, err
}

func (m *postDataRepo) SelectByTitle(title string) (models.PostData, error) {
	result := models.PostData{}
	err := m.DB.Where("title = ?", title).First(&result).Error
	return result, err
}

func (m *postDataRepo) Insert(title string) error {
	postData := models.PostData{
		Title:         title,
		UUID4:         uuid.New().String(),
		UnixTimestamp: time.Now().UTC().Unix(),
	}

	return m.DB.Create(&postData).Error
}
