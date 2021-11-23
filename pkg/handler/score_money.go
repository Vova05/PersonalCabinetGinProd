package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createScoresMoney(c *gin.Context){

	var input models.ScoresMoney
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id,err := h.service.ScoresMoney.Create(input)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"id": id,
	})
}

type getAllScoresMoney struct {
	Data []models.ScoresMoney `json:"data"`
}

func (h *Handler)  getAllScore(c *gin.Context){

	userId, err := getUserId(c)

	if err != nil{
		return
	}


	account, err := h.service.BankAccounts.GetAllUser(userId)
	score, err := h.service.Scores.GetAllAccount(account[0].IdBankAccount)//поменять на срез
	moneyScore, err := h.service.ScoresMoney.GetAllScore(score[0].IdScore)//тоже

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllScoresMoney{
		Data: moneyScore,
	})
}

func (h *Handler)  getAllScoresMoney(c *gin.Context){

	scores, err := h.service.ScoresMoney.GetAll()
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllScoresMoney{
		Data: scores,
	})
}

func (h *Handler) getScoresMoneyById(c *gin.Context){

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	account, err := h.service.ScoresMoney.GetById(id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,account)
}

func (h *Handler) updateScoresMoney(c *gin.Context){


	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateScoresMoney
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	h.service.ScoresMoney.Update(id,input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "Updated",
	})
}

