package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type StatusScoreService struct {
	repo repository.StatusScore
}

func NewStatusScoreService(repo repository.StatusScore) *StatusScoreService{
	return &StatusScoreService{repo: repo}
}

func (s *StatusScoreService) Create(statusScore models.StatusScore) (int, error) {
	return s.repo.Create(statusScore)
}


func (s *StatusScoreService) GetAll() ([]models.StatusScore, error) {
	return s.repo.GetAll()
}

func (s *StatusScoreService) GetById(statusScoreId int) (models.StatusScore, error) {
	return s.repo.GetById(statusScoreId)
}

func (s *StatusScoreService) Delete(statusScoreId int) error {
	return s.repo.Delete(statusScoreId)
}

func (s *StatusScoreService) Update (statusScoreId int, input models.UpdateStatusScore) error {
	return s.repo.Update(statusScoreId,input)
}



