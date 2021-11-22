package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type StatusDB struct {
	db *gorm.DB
}


func NewStatusDB(db *gorm.DB) *StatusDB{
	return &StatusDB{db: db}
}

func (r *StatusDB) Create(status models.Status) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)
	row :=	r.db.Table("status").Create(&status)

	if err :=  row.Table("status").Where("name_status = ?",status.NameStatus).Find(&status).Error; err != nil{
		return 0,err
	}
	r.db.Table("status").Where("name_status = ?",status.NameStatus).Find(&status)
	id = status.IdStatus
	return id, nil
}

func (r *StatusDB) GetAll() ([]models.Status, error) {
	var accounts []models.Status
	err := r.db.Table("status").Find(&accounts).Error
	return accounts,  err
}

func (r *StatusDB) GetById(statusId int) (models.Status, error) {
	var account models.Status
	err := r.db.Table("status").Where("id_status = ?",statusId).Scan(&account).Error
	return account, err
}

func (r *StatusDB) Delete(statusId int) error {
	var account models.Status
	err := r.db.Table("status").Where("id_status = ? ",statusId).First(&account).Error
	if err == nil {
		err = r.db.Table("status").Where("id_status = ? ",statusId).Delete(account).Error
	}else{
		return err
	}
	return err
}

func (r *StatusDB) Update(statusId int, input models.UpdateStatus) error {
	err := r.db.Table("status").Where("id_status = ?", statusId).Updates(models.UpdateStatus{
		NameStatus: input.NameStatus,
		StatusCondition: input.StatusCondition,
	}).Error
	return err
}


