package repository

import (
	"PersonalCabinetGin/models"
	"gorm.io/gorm"
	"time"
)

type CardDB struct {
	db *gorm.DB
}


func NewCardDB(db *gorm.DB) *CardDB{
	return &CardDB{db: db}
}

func (r *CardDB) Create(card models.Card) (int, error) {
	var id int
	card.CreatedAt = time.Now().UTC()
	row :=	r.db.Table("cards").Create(&card)
	if err :=  row.Table("cards").Where("number_card = ?", card.NumberCard).Find(&card).Error; err != nil{
		return 0,err
	}
	r.db.Table("cards").Where("number_card = ?", card.NumberCard).Find(&card)
	id = card.IdCard
	return id, nil
}

func (r *CardDB) GetAllAccount(accountId int) ([]models.Card, error) {
	var cards []models.Card
	err := r.db.Table("cards").Where("bank_account_id = ?", accountId).Find(&cards).Error
	return cards,  err
}

func (r *CardDB) GetAll() ([]models.Card, error) {
	var cards []models.Card
	err := r.db.Table("cards").Find(&cards).Error
	return cards,  err
}

func (r *CardDB) GetById(cardId int) (models.Card, error) {
	var card models.Card
	err := r.db.Table("cards").Where("id_card = ?", cardId).Scan(&card).Error
	return card, err
}

func (r *CardDB) Update(cardId int, input models.UpdateCard) error {
	err := r.db.Table("cards").Where("id_card = ?", cardId).Updates(models.UpdateCard{
		StatusCardId: input.StatusCardId,
		AmountMoney: input.AmountMoney,
		CurrencyId: input.CurrencyId,
		Password: input.Password,
		RememberPassId: input.RememberPassId,
		UpdatedAt: time.Now().UTC(),
	}).Error
	return err
}


