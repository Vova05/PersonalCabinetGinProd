package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createStatusCard(c *gin.Context){

	var input models.StatusCard
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id,err := h.service.StatusCard.Create(input)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"id": id,
	})
}

type getAllStatusCard struct {
	Data []models.StatusCard `json:"data"`
}


func (h *Handler)  getAllStatusCard(c *gin.Context){

	status, err := h.service.StatusCard.GetAll()
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllStatusCard{
		Data: status,
	})
}

func (h *Handler) getStatusCardById(c *gin.Context){

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	application, err := h.service.StatusCard.GetById(id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,application)
}

func (h *Handler) updateStatusCard(c *gin.Context){


	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateStatusCard
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	h.service.StatusCard.Update(id,input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "Updated",
	})
}

func (h *Handler) deleteStatusCard(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.StatusCard.Delete( id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Deleted",
	})
}