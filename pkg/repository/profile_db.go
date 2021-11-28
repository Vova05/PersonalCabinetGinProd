package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type ProfileDB struct {
	db *gorm.DB
}

func NewProfileDB(db *gorm.DB) *ProfileDB{
	return &ProfileDB{db: db}
}

func (r *ProfileDB) GetProfile(token string)(models.User, []models.Application){
	var user models.User
	return user ,nil
}
