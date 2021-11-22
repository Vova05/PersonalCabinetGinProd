package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createRoleRights(c *gin.Context){

	var input models.RoleRight
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	id,err := h.service.RoleRights.Create(input)

	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"id": id,
	})
}

type getAllRoleRights struct {
	Data []models.RoleRight `json:"data"`
}

func (h *Handler)  getAllRoleRightsRole(c *gin.Context){

	userId, err := getUserId(c)

	if err != nil{
		return
	}

	role, err := h.service.Role.GetAllUser(userId)
	roles, err := h.service.RoleRights.GetAllRole(role[0].IdRole)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllRoleRights{
		Data: roles,
	})
}

func (h *Handler)  getAllRoleRights(c *gin.Context){

	roles, err := h.service.RoleRights.GetAll()
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,getAllRoleRights{
		Data: roles,
	})
}

func (h *Handler) getRoleRightsById(c *gin.Context){

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	application, err := h.service.RoleRights.GetById(id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK,application)
}

func (h *Handler) updateRoleRights(c *gin.Context){


	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input models.UpdateRoleRight
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c,http.StatusBadRequest, err.Error())
		return
	}

	h.service.RoleRights.Update(id,input)

	c.JSON(http.StatusOK, statusResponse{
		Status: "Updated",
	})
}

func (h *Handler) deleteRoleRights(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c,http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.service.RoleRights.Delete( id)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "Deleted",
	})
}