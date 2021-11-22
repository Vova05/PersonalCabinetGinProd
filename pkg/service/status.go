package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type StatusService struct {
	repo repository.Status
}

func NewStatusService(repo repository.Status) *StatusService{
	return &StatusService{repo: repo}
}

func (s *StatusService) Create(status models.Status) (int, error) {
	return s.repo.Create(status)
}


func (s *StatusService) GetAll() ([]models.Status, error) {
	return s.repo.GetAll()
}

func (s *StatusService) GetById(statusId int) (models.Status, error) {
	return s.repo.GetById(statusId)
}

func (s *StatusService) Delete(statusId int) error {
	return s.repo.Delete(statusId)
}

func (s *StatusService) Update (statusId int, input models.UpdateStatus) error {
	return s.repo.Update(statusId,input)
}



