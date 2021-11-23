package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createCurrency(c *gin.Context){

	var input models.Currency
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id,err := h.service.Currency.Create(input)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"id": id,
	})
}

type getAllCurrency struct {
	Data []models.Currency `json:"data"`
}


func (h *Handler)  getAllCurrency(c *gin.Context){

	scores, err := h.service.Currency.GetAll()
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllCurrency{
		Data: scores,
	})
}

func (h *Handler) getCurrencyById(c *gin.Context){

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	account, err := h.service.Currency.GetById(id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,account)
}

func (h *Handler) updateCurrency(c *gin.Context){


	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateCurrency
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	h.service.Currency.Update(id,input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "Updated",
	})
}

