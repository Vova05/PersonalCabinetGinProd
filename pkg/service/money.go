package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type MoneyService struct {
	repo repository.Money
}

func NewMoneyService(repo repository.Money) *MoneyService{
	return &MoneyService{repo: repo}
}

func (s *MoneyService) Create(money models.Money) (int, error) {
	return s.repo.Create(money)
}


func (s *MoneyService) GetAll() ([]models.Money, error) {
	return s.repo.GetAll()
}

func (s *MoneyService) GetById(moneyId int) (models.Money, error) {
	return s.repo.GetById(moneyId)
}

func (s *MoneyService) Update (moneyId int, input models.UpdateMoney) error {
	return s.repo.Update(moneyId,input)
}



