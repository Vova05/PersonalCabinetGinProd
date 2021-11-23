package models

import "time"

type Card struct {
	IdCard int `json:"id_card"`
	NumberCard int `json:"number_card"`
	StatusCardId int `json:"status_card_id"`
	AmountMoney int `json:"amount_money"`
	CurrencyId int `json:"currency_id"`
	Password string `json:"password"`
	RememberPassId int `json:"remember_pass_id"`
	BankAccountId int `json:"bank_account_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateCard struct {
	StatusCardId int `json:"status_card_id"`
	AmountMoney int `json:"amount_money"`
	CurrencyId int `json:"currency_id"`
	Password string `json:"password"`
	RememberPassId int `json:"remember_pass_id"`
	UpdatedAt time.Time `json:"updated_at"`
}