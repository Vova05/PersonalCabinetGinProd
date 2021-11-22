package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type RightsDB struct {
	db *gorm.DB
}

func NewRightsDB(db *gorm.DB) *RightsDB{
	return &RightsDB{db: db}
}

func (r *RightsDB) Create(right models.Right) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)
	name := right.NameOperation
	row :=	r.db.Table("rights").Create(&right).First(&right)

	if err :=  row.Table("rights").Where("id_rights = ?",right.IdRights).Find(&right).Error; err != nil{
		return 0,err
	}
	r.db.Table("rights").Where("name_operation = ?",name).Find(&right)
	id = right.IdRights
	return id, nil
}

func (r *RightsDB) GetAll() ([]models.Right, error) {
	var rights []models.Right
	err := r.db.Table("rights").Find(&rights).Error
	return rights,  err
}

func (r *RightsDB) GetById(rightId int) (models.Right, error) {
	var right models.Right
	err := r.db.Table("rights").Where("id_rights = ?",rightId).Scan(&right).Error
	return right, err
}

func (r *RightsDB) Delete(rightId int) error {
	var right models.Right
	err := r.db.Table("rights").Where("id_rights = ? ",rightId).First(&right).Error
	if err == nil {
		err = r.db.Table("rights").Where("id_rights = ? ",rightId).Delete(right).Error
	}else{
		return err
	}
	return err
}

func (r *RightsDB) Update(rightId int, input models.UpdateRight) error {
	err := r.db.Table("rights").Where("id_rights = ?", rightId).Updates(models.UpdateRight{
		NameOperation: input.NameOperation,
		Status: input.Status,
	}).Error
	return err
}


