package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type RoleRightsService struct {
	repo repository.RoleRight
}

func NewRoleRightsService(repo repository.RoleRight) *RoleRightsService{
	return &RoleRightsService{repo: repo}
}

func (s *RoleRightsService) Create(role models.RoleRight) (int, error) {
	return s.repo.Create(role)
}

func (s *RoleRightsService) GetAllRole(roleId int) ([]models.RoleRight, error) {
	return s.repo.GetAllRole(roleId)
}

func (s *RoleRightsService) GetAll() ([]models.RoleRight, error) {
	return s.repo.GetAll()
}

func (s *RoleRightsService) GetById(roleId int) (models.RoleRight, error) {
	return s.repo.GetById(roleId)
}

func (s *RoleRightsService) Delete(roleId int) error {
	return s.repo.Delete(roleId)
}

func (s *RoleRightsService) Update (roleId int, input models.UpdateRoleRight) error {
	return s.repo.Update(roleId,input)
}


