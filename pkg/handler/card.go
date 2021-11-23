package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createCard(c *gin.Context){

	var input models.Card
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id,err := h.service.Card.Create(input)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"id": id,
	})
}

type getAllCard struct {
	Data []models.Card `json:"data"`
}

func (h *Handler)  getAllCardUser(c *gin.Context){

	userId, err := getUserId(c)

	if err != nil{
		return
	}

	account, err := h.service.BankAccounts.GetAllUser(userId)
	cards, err := h.service.Card.GetAllAccount(account[0].IdBankAccount)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllCard{
		Data: cards,
	})
}

func (h *Handler)  getAllCard(c *gin.Context){

	cards, err := h.service.Card.GetAll()
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllCard{
		Data: cards,
	})
}

func (h *Handler) getCardById(c *gin.Context){

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	account, err := h.service.Card.GetById(id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,account)
}

func (h *Handler) updateCard(c *gin.Context){


	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateCard
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	h.service.Card.Update(id,input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "Updated",
	})
}

