package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
	"time"
)

type ApplicationDB struct {
	db *gorm.DB
}

func NewApplicationDB(db *gorm.DB) *ApplicationDB{
	return &ApplicationDB{db: db}
}

func (r *ApplicationDB) Create(userId int, application models.Application)(int, error){
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)
	application.CreatorId=userId
	application.CreatedAt = time.Now().UTC()
	row := r.db.Table("applications").Create(&application)
	if err := row.Scan(&application).Error; err != nil {
		return 0,err
	}
	id = application.IdApplication
	return id, nil
}

func (r *ApplicationDB) GetAll(userId int)([]models.Application, error){
	var applications []models.Application
	err := r.db.Table("applications").Where("creator_id = ?", userId).Find(&applications).Error
	return applications,  err
}

func (r *ApplicationDB) GetById(userId, applicationId int)(models.Application, error){
	var application models.Application
	err := r.db.Table("applications").Where("creator_id = ? AND id_application = ?",userId,applicationId).Scan(&application).Error
	return application, err
}
func (r *ApplicationDB) Delete(userId, applicationId int)( error){
	var app models.Application
	err := r.db.Where("creator_id = ? AND id_application = ?",userId,applicationId).First(&app).Error
	if err == nil {
		err = r.db.Table("applications").Where("creator_id = ? AND id_application = ?",userId,applicationId).Delete(app).Error
	}else{
		return err
	}
	return err
}

func (r *ApplicationDB) Update(userId,id int ,input models.UpdateApplication)( error){
	//setValues := make([]string, 0)
	//args := make([]interface{}, 0)
	//argId := 1

	//if input.Title != nil{
	//	setValues = append(setValues,fmt.Sprintf("title=$%d",argId))
	//	args = append(args,*input.Title)
	//	argId++
	//}
	//if input.Message != nil{
	//	setValues = append(setValues,fmt.Sprintf("title=$%d",argId))
	//	args = append(args,*input.Message)
	//	argId++
	//}

	//setQuery := strings.Join(setValues,", ")
	var app models.Application
	err := r.db.Model(app).Where("creator_id = ? AND id_application = ?", userId, id).Updates(models.UpdateApplication{
		Title: input.Title,
		Message: input.Message,
		TitleId: input.TitleId,
	}).Error
	return err
}

func (r *ApplicationDB) GetAllUser(userId int)([]models.Application,error){
	var applications []models.Application
	err := r.db.Table("applications").Where("creator_id = ? OR recipient_user_id = ?",userId,userId).Find(&applications).Error
	return applications, err
}


func (r *ApplicationDB) GetAllUserResponse(userId int)([]models.Application,error){
	var applications []models.Application
	err := r.db.Table("applications").Where("recipient_user_id = ?",userId).Find(&applications).Error
	return applications, err
}
