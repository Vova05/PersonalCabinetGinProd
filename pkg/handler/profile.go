package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getProfileUser (c *gin.Context){


	userId, _ := h.tokenParseToUserId("Token",c)
	userData, _ := h.service.User.GetById(userId)

	userRole, _ := h.service.Role.GetById(userData.RoleUserId)
	applicationsUser, err := h.service.Application.GetAllUser(userId)
	if err != nil {
		c.Redirect(http.StatusUnauthorized,"http://localhost:9090/401")
		return
	}

	for i, j := 0, len(applicationsUser)-1; i < j; i, j = i+1, j-1 {
		applicationsUser[i], applicationsUser[j] = applicationsUser[j], applicationsUser[i]
	}
	var applicationsGet []models.ApplicationGet
	for index, _ := range applicationsUser{
		var app models.ApplicationGet
		app.Message = applicationsUser[index].Message
		app.Title = applicationsUser[index].Title
		app.CreatedAt = applicationsUser[index].CreatedAt
		userCreate , _  := h.service.User.GetById(applicationsUser[index].CreatorId)
		app.Creator = userCreate.Username
		applicationsGet = append(applicationsGet, app)
	}

	data := gin.H{
		"title": "Profile",
		"NameUser": userData.Name,
		"Login": userData.Username,
		"Role": userRole.NameRole,
		"Applications": applicationsGet,
	}
	c.HTML(http.StatusOK,"profile.html",data)
}
