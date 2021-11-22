package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
	"time"
)

type ScoresDB struct {
	db *gorm.DB
}


func NewScoresDB(db *gorm.DB) *ScoresDB{
	return &ScoresDB{db: db}
}

func (r *ScoresDB) Create(score models.Scores) (int, error) {
	var id int
	score.CreatedAt = time.Now().UTC()
	row :=	r.db.Table("scores").Create(&score)

	if err :=  row.Table("scores").Where("number_score = ?", score.NumberScore).Find(&score).Error; err != nil{
		return 0,err
	}
	r.db.Table("scores").Where("number_score = ?", score.NumberScore).Find(&score)
	id = score.IdScore
	return id, nil
}

func (r *ScoresDB) GetAllAccount(accountId int) ([]models.Scores, error) {
	var scores []models.Scores
	err := r.db.Table("scores").Where("bank_account_id = ?", accountId).Find(&scores).Error
	return scores,  err
}

func (r *ScoresDB) GetAll() ([]models.Scores, error) {
	var scores []models.Scores
	err := r.db.Table("scores").Find(&scores).Error
	return scores,  err
}

func (r *ScoresDB) GetById(scoreId int) (models.Scores, error) {
	var score models.Scores
	err := r.db.Table("scores").Where("id_score = ?", scoreId).Scan(&score).Error
	return score, err
}

func (r *ScoresDB) Update(scoreId int, input models.UpdateScores) error {
	err := r.db.Table("scores").Where("id_score = ?", scoreId).Updates(models.UpdateScores{
		NumberScore: input.NumberScore,
		BankAccountId: input.BankAccountId,
		StatusScoreId: input.StatusScoreId,
	}).Error
	return err
}


