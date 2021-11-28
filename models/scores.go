package models

import "time"

type Scores struct {
	IdScore int `json:"id_score"`
	NumberScore int `json:"number_score"`
	BankAccountId int `json:"bank_account_id"`
	StatusScoreId int `json:"status_score_id"`
	CreatedAt time.Time
}

type UpdateScores struct {
	NumberScore int `json:"number_score"`
	BankAccountId int `json:"bank_account_id"`
	StatusScoreId int `json:"status_score_id"`
}

type ScoresGet struct {
	NumberScore int `json:"number_score"`
	ScoreStatus string `json:"score_status"`
}

type ScoresGetUser struct {
	NumberScore int `json:"number_score"`
	ScoreStatus string `json:"score_status"`
	MoneyScore []ScoresMoneyGet `json:"money_score"`
}
