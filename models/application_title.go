package models

type ApplicationTitle struct {
	IdTitlesApp int `json:"id_titles_app"`
	TitleApp string `json:"title_app"`
}

type UpdateApplicationTitle struct {
	TitleApp string `json:"title_app"`
}
