package models

type Currency struct {
	IdCurrency int `json:"id_currency"`
	NameCurrency string `json:"name_currency"`
	Country string `json:"country"`
}

type UpdateCurrency struct {
	NameCurrency string `json:"name_currency"`
	Country string `json:"country"`
}
