package models

type StatusCard struct {
	IdStatusCard int `json:"id_status_card"`
	NameStatusCard  string `json:"name_status_card"`
	StatusCardCondition bool `json:"status_card_condition"`
}

type UpdateStatusCard struct {
	NameStatusCard  string `json:"name_status_card"`
	StatusCardCondition bool `json:"status_card_condition"`
}
