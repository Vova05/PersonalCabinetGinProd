package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type ApplicationTitleService struct {
	repo repository.ApplicationTitle
}

func NewApplicationTitleService(repo repository.ApplicationTitle) *ApplicationTitleService{
	return &ApplicationTitleService {repo: repo}
}

func (s *ApplicationTitleService) Create(applicationTitle models.ApplicationTitle) (int, error) {
	return s.repo.Create(applicationTitle)
}

func (s *ApplicationTitleService) GetAll() ([]models.ApplicationTitle, error) {
	return s.repo.GetAll()
}

func (s *ApplicationTitleService) GetById(applicationTitleId int) (models.ApplicationTitle, error) {
	return s.repo.GetById(applicationTitleId)
}

func (s *ApplicationTitleService) Delete(applicationTitleId int) error {
	return s.repo.Delete(applicationTitleId)
}

func (s *ApplicationTitleService) Update(applicationTitleId int, input models.UpdateApplicationTitle) error {
	return s.repo.Update(applicationTitleId,input)
}


