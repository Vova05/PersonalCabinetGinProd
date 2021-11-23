package models

import "time"

type Reminder struct {
	IdReminder int `json:"id_reminder"`
	Hint string `json:"hint"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateReminder struct {
	Hint string `json:"hint"`
	UpdatedAt time.Time `json:"updated_at"`
}
