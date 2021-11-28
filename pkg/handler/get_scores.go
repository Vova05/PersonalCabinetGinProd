package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getScores (c *gin.Context){
	userId, _ := h.tokenParseToUserId("Token",c)
	userData, _ := h.service.User.GetById(userId)

	banksAccountsUser, _ := h.service.BankAccounts.GetAllUser(userId)
	//var userScores []models.ScoresGetUser
	var scoresUser []models.ScoresGetUser
	for index, _:= range banksAccountsUser{
		scoreAccount, _ := h.service.Scores.GetAllAccount(banksAccountsUser[index].IdBankAccount)
		for index2, _ := range scoreAccount{
			var scoreAccountGetOne models.ScoresGetUser
			scoreAccountGetOne.NumberScore = scoreAccount[index2].NumberScore
			statusScore, _ := h.service.StatusScore.GetById(scoreAccount[index2].StatusScoreId)
			scoreAccountGetOne.ScoreStatus = statusScore.NameStatusScore
			scoreMoney, _ := h.service.ScoresMoney.GetAllScore(scoreAccount[index2].IdScore)
			var scoreMoneyGet []models.ScoresMoneyGet
			for index3, _:=range scoreMoney{
				var scoreMoneyGetOne models.ScoresMoneyGet
				quantity, _ := h.service.Money.GetById(scoreMoney[index3].ScoreMoneyId)
				scoreMoneyGetOne.Quantity = quantity.Quantity
				currency, _ := h.service.Currency.GetById(scoreMoney[index3].ScoreCurrencyId)
				scoreMoneyGetOne.CurrencyMoney = currency.NameCurrency
				scoreMoneyGet = append(scoreMoneyGet, scoreMoneyGetOne)
			}
			scoreAccountGetOne.MoneyScore = scoreMoneyGet
			scoresUser = append(scoresUser,scoreAccountGetOne)
		}
	}

	var cardsUser []models.CardGet
	for index, _:= range banksAccountsUser{
		cardAccount, _ := h.service.Card.GetAllAccount(banksAccountsUser[index].IdBankAccount)
		for index2, _ := range cardAccount {
			var cardGetOne models.CardGet
			cardGetOne.NumberCard = cardAccount[index2].NumberCard
			statusCard, _ := h.service.StatusCard.GetById(cardAccount[index2].StatusCardId)
			cardGetOne.StatusCardName = statusCard.NameStatusCard
			cardGetOne.AmountMoney = cardAccount[index2].AmountMoney
			rememberPass, _ := h.service.Reminder.GetById(cardAccount[index2].RememberPassId)
			cardGetOne.RememberPass = rememberPass.Hint
			cardGetOne.CreatedAt = cardAccount[index2].CreatedAt
			currencyCard, _ := h.service.Currency.GetById(cardAccount[index2].CurrencyId)
			cardGetOne.CurrencyName = currencyCard.NameCurrency
			cardsUser = append(cardsUser,cardGetOne)
		}
	}


	data := gin.H{
		"title": "Bank accounts",
		"NameUser": userData.Name,
		"Login": userData.Username,
		"Scores": scoresUser,
		"Cards": cardsUser,
	}
	c.HTML(http.StatusOK,"scores.html",data)

}

