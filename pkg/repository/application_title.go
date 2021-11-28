package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type ApplicationTitleDB struct {
	db *gorm.DB
}


func NewApplicationTitleDB(db *gorm.DB) *ApplicationTitleDB{
	return &ApplicationTitleDB{db: db}
}

func (r *ApplicationTitleDB) Create(input models.ApplicationTitle) (int, error) {
	var id int
	r.db.Table("titles_app").Create(&input)
	r.db.Table("titles_app").Where("title_app = ?",input.TitleApp).First(&input)
	id = input.IdTitlesApp
	return id, nil
}



func (r *ApplicationTitleDB) GetAll() ([]models.ApplicationTitle, error) {
	var applicationTitle []models.ApplicationTitle
	err := r.db.Table("titles_app").Find(&applicationTitle).Error
	return applicationTitle,  err
}

func (r *ApplicationTitleDB) GetById(applicationTitleId int) (models.ApplicationTitle, error) {
	var applicationTitle models.ApplicationTitle
	err := r.db.Table("titles_app").Where("id_titles_app = ?",applicationTitleId).First(&applicationTitle).Error
	return applicationTitle, err
}

func (r *ApplicationTitleDB) Delete(applicationTitleId int) error {
	var applicationTitle models.ApplicationTitle
	err := r.db.Table("titles_app").Where("id_titles_app = ? ",applicationTitleId).First(&applicationTitle).Error
	if err == nil {
		err = r.db.Table("titles_app").Where("id_titles_app = ? ",applicationTitleId).Delete(applicationTitle).Error
	}else{
		return err
	}
	return err
}

func (r *ApplicationTitleDB) Update(applicationTitleId int, input models.UpdateApplicationTitle) error {
	var applicationTitle models.ApplicationTitle
	err := r.db.Model(applicationTitle).Where("id_titles_app = ?", applicationTitleId).Updates(models.UpdateApplicationTitle{
		TitleApp: input.TitleApp,
	}).Error
	return err
}


