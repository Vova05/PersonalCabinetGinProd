package models

import "time"

type Application struct {
	IdApplication int `json:"id"`
	Title string `json:"title"`
	Message string `json:"message"`
	CreatorId int `json:"creator"`
	TitleId int `json:"title_id"`
	RecipientUserId int `json:"recipient_user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateApplication struct {
	Title string `json:"title"`
	Message string `json:"message"`
	TitleId int `json:"title_id"`
	RecipientUserId int `json:"recipient_user_id"`
}

type ApplicationGet struct {
	Title string `json:"title"`
	Message string `json:"message"`
	Creator string `json:"creator"`
	CreatedAt time.Time `json:"created_at"`
}
