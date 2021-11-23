package models

import "time"

type Money struct {
	IdMoney int `json:"id_money"`
	Quantity int `json:"quantity"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateMoney struct {
	Quantity int `json:"quantity"`
	UpdatedAt time.Time
}
