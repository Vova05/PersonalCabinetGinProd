package models

type Application struct {
	IdApplication int `json:"id"`
	Title string `json:"title"`
	Message string `json:"message"`
	CreatorId int `json:"creator"`
	TitleId int `json:"title_id"`
}

type UpdateApplication struct {
	Title string `json:"title"`
	Message string `json:"message"`
	TitleId int `json:"title_id"`
}
