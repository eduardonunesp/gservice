package services

import (
	"github.com/eduardonunesp/gservice/models"
	"github.com/eduardonunesp/gservice/repos"
)

type DataService interface {
	GetAll() ([]models.Data, error)
	GetByName(name string) (models.Data, error)
	Insert(name string, stage, score int) (models.Data, error)
}

type dataService struct {
	repo repos.DataRepo
}

func NewDataService(repo repos.DataRepo) DataService {
	return &dataService{
		repo: repo,
	}
}

func (s *dataService) Insert(name string, stage, score int) (models.Data, error) {
	return s.repo.Insert(name, stage, score)
}

func (s *dataService) GetAll() ([]models.Data, error) {
	return s.repo.Select()
}

func (s *dataService) GetByName(name string) (models.Data, error) {
	return s.repo.SelectByName(name)
}

func (s *dataService) GetByID(name string) (models.Data, error) {
	return s.repo.SelectById(name)
}
