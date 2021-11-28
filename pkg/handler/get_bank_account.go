package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)



func (h *Handler) getBankAccount (c *gin.Context){
	userId, _ := h.tokenParseToUserId("Token",c)
	userData, _ := h.service.User.GetById(userId)

	banksAccountsUser, _ := h.service.BankAccounts.GetAllUser(userId)
	var banksAccountsUserGet []models.BankAccountsGet
	for index, _:= range banksAccountsUser{
		statusAccount,_ := h.service.Status.GetById(banksAccountsUser[index].StatusId)
		scoreAccount, _ := h.service.Scores.GetAllAccount(banksAccountsUser[index].IdBankAccount)
		var scoreAccountGet []models.ScoresGet
		for index2, _ := range scoreAccount{
			var scoreAccountGetOne models.ScoresGet
			scoreAccountGetOne.NumberScore = scoreAccount[index2].NumberScore
			statusScore, _ := h.service.StatusScore.GetById(scoreAccount[index2].StatusScoreId)
			scoreAccountGetOne.ScoreStatus = statusScore.NameStatusScore
			scoreAccountGet = append(scoreAccountGet,scoreAccountGetOne)
		}
		var account models.BankAccountsGet
		account.NumberAccount = banksAccountsUser[index].NumberAccount
		account.ScoresAccount = scoreAccountGet
		account.StatusName = statusAccount.NameStatus
		banksAccountsUserGet = append(banksAccountsUserGet, account)
	}

	//:= h.service.Status.GetById()
	//banksScores

	data := gin.H{
		"title": "Bank accounts",
		"NameUser": userData.Name,
		"Accounts": banksAccountsUserGet,
	}
	c.HTML(http.StatusOK,"bank_account.html",data)

}
