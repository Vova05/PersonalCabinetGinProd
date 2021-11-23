package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
)

type CurrencyDB struct {
	db *gorm.DB
}


func NewCurrencyDB(db *gorm.DB) *CurrencyDB{
	return &CurrencyDB{db: db}
}

func (r *CurrencyDB) Create(currency models.Currency) (int, error) {
	var id int

	row :=	r.db.Table("currency").Create(&currency)

	if err :=  row.Table("currency").Where("id_currency = ?", currency.IdCurrency).Find(&currency).Error; err != nil{
		return 0,err
	}
	r.db.Table("currency").Last(&currency)
	id = currency.IdCurrency
	return id, nil
}

func (r *CurrencyDB) GetAll() ([]models.Currency, error) {
	var currency []models.Currency
	err := r.db.Table("currency").Find(&currency).Error
	return currency,  err
}

func (r *CurrencyDB) GetById(currencyId int) (models.Currency, error) {
	var currency models.Currency
	err := r.db.Table("currency").Where("id_currency = ?", currencyId).Scan(&currency).Error
	return currency, err
}

func (r *CurrencyDB) Update(currencyId int, input models.UpdateCurrency) error {

	err := r.db.Table("currency").Where("id_currency = ?", currencyId).Updates(models.UpdateCurrency{
		NameCurrency: input.NameCurrency,
		Country: input.Country,
	}).Error
	return err
}


