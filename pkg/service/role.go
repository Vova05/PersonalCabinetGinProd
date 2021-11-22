package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type RoleService struct {
	repo repository.Role
}

func NewRoleService(repo repository.Role) *RoleService{
	return &RoleService{repo: repo}
}

func (s *RoleService) Create(role models.Role) (int, error) {
	return s.repo.Create(role)
}

func (s *RoleService) GetAllUser(userId int) ([]models.Role, error) {
	return s.repo.GetAllUser(userId)
}

func (s *RoleService) GetAll() ([]models.Role, error) {
	return s.repo.GetAll()
}

func (s *RoleService) GetById(roleId int) (models.Role, error) {
	return s.repo.GetById(roleId)
}

func (s *RoleService) Delete(roleId int) error {
	return s.repo.Delete(roleId)
}

func (s *RoleService) Update(roleId int, input models.UpdateRole) error {
	return s.repo.Update(roleId,input)
}


