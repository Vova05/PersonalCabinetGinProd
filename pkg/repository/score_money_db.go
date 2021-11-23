package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type ScoresMoneyDB struct {
	db *gorm.DB
}


func NewScoresMoneyDB(db *gorm.DB) *ScoresMoneyDB{
	return &ScoresMoneyDB{db: db}
}

func (r *ScoresMoneyDB) Create(scoresMoney models.ScoresMoney) (int, error) {
	var id int

	row :=	r.db.Table("score_money").Create(&scoresMoney)

	if err :=  row.Table("score_money").Where("id_score_money = ?", scoresMoney.IdScoreMoney).Find(&scoresMoney).Error; err != nil{
		return 0,err
	}
	r.db.Table("score_money").Where("score_money_id = ? AND score_currency_id = ? AND score_id = ?", scoresMoney.ScoreMoneyId,scoresMoney.ScoreCurrencyId,scoresMoney.ScoreId).Find(&scoresMoney)
	id = scoresMoney.IdScoreMoney
	return id, nil
}

func (r *ScoresMoneyDB) GetAllScore(scoreId int) ([]models.ScoresMoney, error) {
	var scoreMoney []models.ScoresMoney
	err := r.db.Table("score_money").Where("score_id = ?", scoreId).Find(&scoreMoney).Error
	return scoreMoney,  err
}

func (r *ScoresMoneyDB) GetAll() ([]models.ScoresMoney, error) {
	var scoreMoney []models.ScoresMoney
	err := r.db.Table("score_money").Find(&scoreMoney).Error
	return scoreMoney,  err
}

func (r *ScoresMoneyDB) GetById(scoreMoneyId int) (models.ScoresMoney, error) {
	var scoreMoney models.ScoresMoney
	err := r.db.Table("score_money").Where("id_score_money = ?", scoreMoneyId).Scan(&scoreMoney).Error
	return scoreMoney, err
}

func (r *ScoresMoneyDB) Update(scoreMoneyId int, input models.UpdateScoresMoney) error {
	err := r.db.Table("score_money").Where("id_score_money = ?", scoreMoneyId).Updates(models.UpdateScoresMoney{
		ScoreMoneyId: input.ScoreMoneyId,
		ScoreCurrencyId: input.ScoreCurrencyId,
		ScoreId: input.ScoreId,
	}).Error
	return err
}


