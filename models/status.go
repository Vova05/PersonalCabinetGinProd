package models

type Status struct {
	IdStatus int `json:"id_status"`
	NameStatus string `json:"name_status"`
	StatusCondition bool `json:"status_condition"`
}

type UpdateStatus struct {
	NameStatus string `json:"name_status"`
	StatusCondition bool `json:"status_condition"`
}
