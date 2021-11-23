package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type CurrencyService struct {
	repo repository.Currency
}

func NewCurrencyService(repo repository.Currency) *CurrencyService{
	return &CurrencyService{repo: repo}
}

func (s *CurrencyService) Create(currency models.Currency) (int, error) {
	return s.repo.Create(currency)
}


func (s *CurrencyService) GetAll() ([]models.Currency, error) {
	return s.repo.GetAll()
}

func (s *CurrencyService) GetById(currencyId int) (models.Currency, error) {
	return s.repo.GetById(currencyId)
}

func (s *CurrencyService) Update (currencyId int, input models.UpdateCurrency) error {
	return s.repo.Update(currencyId,input)
}



