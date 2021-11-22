package models

type Application struct {
	IdApplication int `json:"id"`
	Title string `json:"title"`
	Message string `json:"message"`
	CreatorId int `json:"creator"`
}

type UpdateApplication struct {
	Title string `json:"title"`
	Message string `json:"message"`
}
