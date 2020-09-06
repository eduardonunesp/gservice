package services

import (
	"github.com/eduardonunesp/gservice/models"
	"github.com/eduardonunesp/gservice/repos"
)

type DataService interface {
	GetAll() ([]models.Data, error)
	GetByTitle(name string) (models.Data, error)
	Insert(title string) error
}

type dataService struct {
	repo repos.DataRepo
}

func NewDataService(repo repos.DataRepo) DataService {
	return &dataService{
		repo: repo,
	}
}

func (s *dataService) Insert(title string) error {
	return s.repo.Insert(title)
}

func (s *dataService) GetAll() ([]models.Data, error) {
	return s.repo.Select()
}

func (s *dataService) GetByTitle(title string) (models.Data, error) {
	return s.repo.SelectByTitle(title)
}
