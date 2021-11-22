package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type RoleDB struct {
	db *gorm.DB
}

type UserUpdate struct {
	Name string
	Username string
	Password string
	RoleUserId int
}

func NewRoleDB(db *gorm.DB) *RoleDB{
	return &RoleDB{db: db}
}

func (r *RoleDB) Create(role models.Role) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)
	row :=	r.db.Select("NameRole").Create(&role)

	if err :=  row.Table("roles").Where("name_role = ?",role.NameRole).Find(&role).Error; err != nil{
		return 0,err
	}
	r.db.Table("roles").Where("name_role = ?",role.NameRole).Find(&role)
	id = role.IdRole
	return id, nil
}

func (r *RoleDB) GetAllUser(userId int) ([]models.Role, error) {
	var roles []models.Role
	var user UserUpdate
	err := r.db.Table("users").Where("id_user = ?",userId).First(&user).Error
	if err != nil {
		return nil,err
	}
	idRole:= user.RoleUserId
	err = r.db.Table("roles").Where("id_role = ?", idRole).Find(&roles).Error
	return roles,  err
}

func (r *RoleDB) GetAll() ([]models.Role, error) {
	var roles []models.Role
	err := r.db.Table("roles").Find(&roles).Error
	return roles,  err
}

func (r *RoleDB) GetById(roleId int) (models.Role, error) {
	var role models.Role
	err := r.db.Table("roles").Where("id_role = ?",roleId).Scan(&role).Error
	return role, err
}

func (r *RoleDB) Delete(roleId int) error {
	var role models.Role
	err := r.db.Table("roles").Where("id_role = ? ",roleId).First(&role).Error
	if err == nil {
		err = r.db.Table("roles").Where("id_role = ? ",roleId).Delete(role).Error
	}else{
		return err
	}
	return err
}

func (r *RoleDB) Update(roleId int, input models.UpdateRole) error {
	var role models.Role
	err := r.db.Model(role).Where("id_role = ?", roleId).Updates(models.UpdateRole{
		NameRole: input.NameRole,
	}).Error
	return err
}


