package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
	"time"
)

type MoneyDB struct {
	db *gorm.DB
}


func NewMoneyDB(db *gorm.DB) *MoneyDB{
	return &MoneyDB{db: db}
}

func (r *MoneyDB) Create(money models.Money) (int, error) {
	var id int

	money.CreatedAt = time.Now().UTC()
	money.UpdatedAt = time.Now().UTC()
	createTime := money.CreatedAt
	row :=	r.db.Table("money").Create(&money)

	if err :=  row.Table("money").Where("created_at = ?", createTime.UTC()).Find(&money).Error; err != nil{
		return 0,err
	}
	r.db.Table("money").Last(&money)
	id = money.IdMoney
	return id, nil
}

func (r *MoneyDB) GetAll() ([]models.Money, error) {
	var money []models.Money
	err := r.db.Table("money").Find(&money).Error
	return money,  err
}

func (r *MoneyDB) GetById(moneyId int) (models.Money, error) {
	var money models.Money
	err := r.db.Table("money").Where("id_money = ?", moneyId).Scan(&money).Error
	return money, err
}

func (r *MoneyDB) Update(moneyId int, input models.UpdateMoney) error {
	input.UpdatedAt = time.Now().UTC()
	err := r.db.Table("money").Where("id_money = ?", moneyId).Updates(models.UpdateMoney{
		Quantity: input.Quantity,
		UpdatedAt: input.UpdatedAt,
	}).Error
	return err
}


