package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService{
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetProfile(token string)(models.User, []models.Application){
	return s.repo.GetProfile(token)
}
