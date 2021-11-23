package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}


func NewUserDB(db *gorm.DB) *UserDB{
	return &UserDB{db: db}
}

func (r *UserDB) Create(user models.User) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)

	row :=	r.db.Table("users").Create(&user)

	if err :=  row.Table("users").Where("username = ?",user.Username).Find(&user).Error; err != nil{
		return 0,err
	}
	r.db.Table("users").Last(&user)
	id = user.IdUser
	return id, nil
}

func (r *UserDB) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Table("users").Find(&users).Error
	return users,  err
}

func (r *UserDB) GetById(userId int) (models.User, error) {
	var user models.User
	err := r.db.Table("users").Where("id_user = ?",userId).Scan(&user).Error
	return user, err
}

func (r *UserDB) Delete(userId int) error {
	var user models.User
	err := r.db.Table("users").Where("id_user = ? ", userId).First(&user).Error
	if err == nil {
		err = r.db.Table("users").Where("id_user = ? ", userId).Delete(user).Error
	}else{
		return err
	}
	return err
}

func (r *UserDB) Update(userId int, input models.UpdateUser) error {
	err := r.db.Table("users").Where("id_user = ?", userId).Updates(models.UpdateUser{
		Name: input.Name,
		Username: input.Username,
		Password: input.Password,
		RoleUserId: input.RoleUserId,
		BankUser: input.BankUser,
	}).Error
	return err
}


