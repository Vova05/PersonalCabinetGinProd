package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type CardService struct {
	repo repository.Card
}

func NewCardService(repo repository.Card) *CardService{
	return &CardService{repo: repo}
}

func (s *CardService) Create(card models.Card) (int, error) {
	return s.repo.Create(card)
}


func (s *CardService) GetAll() ([]models.Card, error) {
	return s.repo.GetAll()
}

func (s *CardService) GetById(cardId int) (models.Card, error) {
	return s.repo.GetById(cardId)
}

func (s *CardService) Update (scoreId int, input models.UpdateCard) error {
	return s.repo.Update(scoreId,input)
}

func (s *CardService) GetAllAccount(accountId int) ([]models.Card, error) {
	return s.repo.GetAllAccount(accountId)
}


