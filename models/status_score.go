package models

type StatusScore struct {
	IdStatusScore int `json:"id_status_score"`
	NameStatusScore string `json:"name_status_score"`
	StatusConditionScore bool `json:"status_condition_score"`
}

type UpdateStatusScore struct {
	NameStatusScore string `json:"name_status_score"`
	StatusConditionScore bool `json:"status_condition_score"`
}
