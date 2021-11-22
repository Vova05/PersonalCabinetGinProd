package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type ScoresService struct {
	repo repository.Scores
}

func NewScoresService(repo repository.Scores) *ScoresService{
	return &ScoresService{repo: repo}
}

func (s *ScoresService) Create(score models.Scores) (int, error) {
	return s.repo.Create(score)
}


func (s *ScoresService) GetAll() ([]models.Scores, error) {
	return s.repo.GetAll()
}

func (s *ScoresService) GetById(scoreId int) (models.Scores, error) {
	return s.repo.GetById(scoreId)
}

func (s *ScoresService) Update (scoreId int, input models.UpdateScores) error {
	return s.repo.Update(scoreId,input)
}

func (s *ScoresService) GetAllAccount(accountId int) ([]models.Scores, error) {
	return s.repo.GetAllAccount(accountId)
}


