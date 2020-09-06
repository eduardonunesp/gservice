package services

import (
	"gservice/models"
	"gservice/repos"
)

type PostDataService interface {
	GetAll() ([]models.PostData, error)
	GetByTitle(name string) (models.PostData, error)
	Insert(title string) error
}

type postDataService struct {
	repo repos.PostDataRepo
}

func NewPostDataService(repo repos.PostDataRepo) PostDataService {
	return &postDataService{
		repo: repo,
	}
}

func (s *postDataService) Insert(title string) error {
	return s.repo.Insert(title)
}

func (s *postDataService) GetAll() ([]models.PostData, error) {
	return s.repo.Select()
}

func (s *postDataService) GetByTitle(title string) (models.PostData, error) {
	return s.repo.SelectByTitle(title)
}
