package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getProfile (c *gin.Context){
	data := gin.H{
		"title": "Profile",
	}
	c.HTML(http.StatusOK,"profile.html",data)
}

func (h *Handler) createUser(c *gin.Context){

	var input models.User
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id,err := h.service.User.Create(input)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"id": id,
	})
}

type getAllUser struct {
	Data []models.User `json:"data"`
}


func (h *Handler)  getAllUser(c *gin.Context){

	users, err := h.service.User.GetAll()
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllUser{
		Data: users,
	})
}

func (h *Handler) getUserById(c *gin.Context){

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.service.User.GetById(id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,user)
}

func (h *Handler) updateUser(c *gin.Context){


	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateUser
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	h.service.User.Update(id,input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "Updated",
	})
}

func (h *Handler) deleteUser(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.User.Delete( id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Deleted",
	})
}