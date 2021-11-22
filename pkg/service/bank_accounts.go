package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type BankAccountsService struct {
	repo repository.BankAccounts
}

func NewBankAccountsService(repo repository.BankAccounts) *BankAccountsService{
	return &BankAccountsService{repo: repo}
}

func (s *BankAccountsService) Create(account models.BankAccounts) (int, error) {
	return s.repo.Create(account)
}


func (s *BankAccountsService) GetAll() ([]models.BankAccounts, error) {
	return s.repo.GetAll()
}

func (s *BankAccountsService) GetById(accountId int) (models.BankAccounts, error) {
	return s.repo.GetById(accountId)
}

func (s *BankAccountsService) Delete(accountId int) error {
	return s.repo.Delete(accountId)
}

func (s *BankAccountsService) Update (accountId int, input models.UpdateBankAccounts) error {
	return s.repo.Update(accountId,input)
}

func (s *BankAccountsService) GetAllUser(userId int) ([]models.BankAccounts, error) {
	return s.repo.GetAllUser(userId)
}


