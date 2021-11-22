package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type RightsService struct {
	repo repository.Rights
}

func NewRightsService(repo repository.Rights) *RightsService{
	return &RightsService{repo: repo}
}

func (s *RightsService) Create(right models.Right) (int, error) {
	return s.repo.Create(right)
}


func (s *RightsService) GetAll() ([]models.Right, error) {
	return s.repo.GetAll()
}

func (s *RightsService) GetById(rightId int) (models.Right, error) {
	return s.repo.GetById(rightId)
}

func (s *RightsService) Delete(rightId int) error {
	return s.repo.Delete(rightId)
}

func (s *RightsService) Update (rightId int, input models.UpdateRight) error {
	return s.repo.Update(rightId,input)
}


