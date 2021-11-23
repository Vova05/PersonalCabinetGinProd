package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type ScoresMoneyService struct {
	repo repository.ScoresMoney
}

func NewScoresMoneyService(repo repository.ScoresMoney) *ScoresMoneyService{
	return &ScoresMoneyService{repo: repo}
}

func (s *ScoresMoneyService) Create(scoresMoney models.ScoresMoney) (int, error) {
	return s.repo.Create(scoresMoney)
}


func (s *ScoresMoneyService) GetAll() ([]models.ScoresMoney, error) {
	return s.repo.GetAll()
}

func (s *ScoresMoneyService) GetById(scoresMoneyId int) (models.ScoresMoney, error) {
	return s.repo.GetById(scoresMoneyId)
}

func (s *ScoresMoneyService) Update (scoresMoneyId int, input models.UpdateScoresMoney) error {
	return s.repo.Update(scoresMoneyId,input)
}

func (s *ScoresMoneyService) GetAllScore(scoreId int) ([]models.ScoresMoney, error) {
	return s.repo.GetAllScore(scoreId)
}


