package models

type BankAccounts struct {
	IdBankAccount int `json:"id_bank_account"`
	NumberAccount int `json:"number_account"`
	ClientId int `json:"client_id"`
	StatusId int `json:"status_id"`
}

type UpdateBankAccounts struct {
	NumberAccount int `json:"number_account"`
	ClientId int `json:"client_id"`
	StatusId int `json:"status_id"`
}

type BankAccountsGet struct {
	NumberAccount int `json:"number_account"`
	StatusName string `json:"status_name"`
	ScoresAccount []ScoresGet `json:"scores_account"`
}
