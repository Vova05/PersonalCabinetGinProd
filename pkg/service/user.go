package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService{
	return &UserService{repo: repo}
}

func (s *UserService) Create(user models.User) (int, error) {
	user.Password= generatePasswordHash(user.Password)
	return s.repo.Create(user)
}


func (s *UserService) GetAll() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetById(userId int) (models.User, error) {
	return s.repo.GetById(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *UserService) Update (userId int, input models.UpdateUser) error {
	if input.Password!= ""{
		input.Password= generatePasswordHash(input.Password)
	}
	return s.repo.Update(userId,input)
}

func (s *UserService) GetByToken(token string)(int,error){
	return s.repo.GetByToken(token )
}



