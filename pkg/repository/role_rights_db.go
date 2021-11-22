package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type RoleRightsDB struct {
	db *gorm.DB
}

type RoleUpdate struct {
	IdRole int
	NameRole string
}

func NewRoleRightsDB(db *gorm.DB) *RoleRightsDB{
	return &RoleRightsDB{db: db}
}

func (r *RoleRightsDB) Create(role models.RoleRight) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", Table)

	tmp := role
	row :=	r.db.Table("roles_rights").Create(&role)

	if err :=  row.Table("roles_rights").Find(&role).Error; err != nil{
		return 0,err
	}
	r.db.Table("roles_rights").Where("role_id = ? AND right_id = ?",tmp.RightId,tmp.RoleId).Find(&role)
	//row.Table("roles_rights").Find(&role)
	id = role.IdRolesRights
	return id, nil
}

func (r *RoleRightsDB) GetAllRole(roleId int) ([]models.RoleRight, error) {
	var roles []models.RoleRight //roleId=userId вверху
	var roleUp RoleUpdate
	err := r.db.Table("roles").Where("id_role = ?",roleId).First(&roleUp).Error
	if err != nil {
		return nil,err
	}
	idRole:= roleUp.IdRole
	err = r.db.Table("roles_rights").Where("role_id = ?", idRole).Find(&roles).Error
	return roles,  err
}

func (r *RoleRightsDB) GetAll() ([]models.RoleRight, error) {
	var roles []models.RoleRight
	err := r.db.Table("roles_rights").Find(&roles).Error
	return roles,  err
}

func (r *RoleRightsDB) GetById(roleId int) (models.RoleRight, error) {
	var role models.RoleRight
	err := r.db.Table("roles_rights").Where("id_roles_rights = ?",roleId).Scan(&role).Error
	return role, err
}

func (r *RoleRightsDB) Delete(roleId int) error {
	var role models.RoleRight
	err := r.db.Table("roles_rights").Where("id_roles_rights = ? ",roleId).First(&role).Error
	if err == nil {
		err = r.db.Table("roles_rights").Where("id_roles_rights = ? ",roleId).Delete(role).Error
	}else{
		return err
	}
	return err
}

func (r *RoleRightsDB) Update(roleId int, input models.UpdateRoleRight) error {
	//var role models.RoleRight
	err := r.db.Table("roles_rights").Where("id_roles_rights = ?", roleId).Updates(models.UpdateRoleRight{
		RoleId: input.RoleId,
		RightId: input.RightId,
	}).Error
	return err
}


