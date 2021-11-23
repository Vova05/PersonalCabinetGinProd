package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type StatusScoreDB struct {
	db *gorm.DB
}


func NewStatusScoreDB(db *gorm.DB) *StatusScoreDB{
	return &StatusScoreDB{db: db}
}

func (r *StatusScoreDB) Create(statusScore models.StatusScore) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)
	row :=	r.db.Table("status_scores").Create(&statusScore)

	if err :=  row.Table("status_scores").Where("name_status_score = ?",statusScore.NameStatusScore).Find(&statusScore).Error; err != nil{
		return 0,err
	}
	r.db.Table("status_scores").Where("name_status_score = ?",statusScore.NameStatusScore).Find(&statusScore)
	id = statusScore.IdStatusScore
	return id, nil
}

func (r *StatusScoreDB) GetAll() ([]models.StatusScore, error) {
	var statusScore []models.StatusScore
	err := r.db.Table("status_scores").Find(&statusScore).Error
	return statusScore,  err
}

func (r *StatusScoreDB) GetById(statusScoreId int) (models.StatusScore, error) {
	var statusScore models.StatusScore
	err := r.db.Table("status_scores").Where("id_status_score = ?",statusScoreId).Scan(&statusScore).Error
	return statusScore, err
}

func (r *StatusScoreDB) Delete(statusScoreId int) error {
	var account models.StatusScore
	err := r.db.Table("status_scores").Where("id_status_score = ? ",statusScoreId).First(&account).Error
	if err == nil {
		err = r.db.Table("status_scores").Where("id_status_score = ? ",statusScoreId).Delete(account).Error
	}else{
		return err
	}
	return err
}

func (r *StatusScoreDB) Update(statusScoreId int, input models.UpdateStatusScore) error {
	err := r.db.Table("status_scores").Where("id_status_score = ?", statusScoreId).Updates(models.UpdateStatusScore{
		NameStatusScore: input.NameStatusScore,
		StatusConditionScore: input.StatusConditionScore,
	}).Error
	return err
}


