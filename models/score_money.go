package models

type ScoresMoney struct {
	IdScoreMoney int `json:"id_score_money"`
	ScoreMoneyId int `json:"score_money_id"`
	ScoreCurrencyId int `json:"score_currency_id"`
	ScoreId int `json:"score_id"`
}

type UpdateScoresMoney struct {
	ScoreMoneyId int `json:"score_money_id"`
	ScoreCurrencyId int `json:"score_currency_id"`
	ScoreId int `json:"score_id"`
}
