package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type StatusCardDB struct {
	db *gorm.DB
}


func NewStatusCardDB(db *gorm.DB) *StatusCardDB{
	return &StatusCardDB{db: db}
}

func (r *StatusCardDB) Create(status models.StatusCard) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)
	row :=	r.db.Table("status_cards").Create(&status)

	if err :=  row.Table("status_cards").Where("name_status_card = ?",status.NameStatusCard).Find(&status).Error; err != nil{
		return 0,err
	}
	r.db.Table("status_cards").Where("name_status_card = ?",status.NameStatusCard).Find(&status)
	id = status.IdStatusCard
	return id, nil
}

func (r *StatusCardDB) GetAll() ([]models.StatusCard, error) {
	var status []models.StatusCard
	err := r.db.Table("status_cards").Find(&status).Error
	return status,  err
}

func (r *StatusCardDB) GetById(statusId int) (models.StatusCard, error) {
	var status models.StatusCard
	err := r.db.Table("status_cards").Where("id_status_card = ?",statusId).Scan(&status).Error
	return status, err
}

func (r *StatusCardDB) Delete(statusId int) error {
	var account models.StatusCard
	err := r.db.Table("status_cards").Where("id_status_card = ? ",statusId).First(&account).Error
	if err == nil {
		err = r.db.Table("status_cards").Where("id_status_card = ? ",statusId).Delete(account).Error
	}else{
		return err
	}
	return err
}

func (r *StatusCardDB) Update(statusId int, input models.UpdateStatusCard) error {
	err := r.db.Table("status_cards").Where("id_status_card = ?", statusId).Updates(models.UpdateStatusCard{
		NameStatusCard: input.NameStatusCard,
		StatusCardCondition: input.StatusCardCondition,
	}).Error
	return err
}


