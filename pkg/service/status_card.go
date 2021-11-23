package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type StatusCardService struct {
	repo repository.StatusCard
}

func NewStatusCardService(repo repository.StatusCard) *StatusCardService{
	return &StatusCardService{repo: repo}
}

func (s *StatusCardService) Create(status models.StatusCard) (int, error) {
	return s.repo.Create(status)
}


func (s *StatusCardService) GetAll() ([]models.StatusCard, error) {
	return s.repo.GetAll()
}

func (s *StatusCardService) GetById(statusId int) (models.StatusCard, error) {
	return s.repo.GetById(statusId)
}

func (s *StatusCardService) Delete(statusId int) error {
	return s.repo.Delete(statusId)
}

func (s *StatusCardService) Update (statusId int, input models.UpdateStatusCard) error {
	return s.repo.Update(statusId,input)
}



