package handler

import (
	"PersonalCabinetGin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) addUserGet (c *gin.Context){


	userId, _ := h.tokenParseToUserId("Token",c)
	userData, _ := h.service.User.GetById(userId)


	userAccounts,_  := h.service.User.GetAll()
	var userAccountsGet []models.UserGet
	for index, _ := range userAccounts{
		var userAccountGetOne models.UserGet
		var userBankAccounts []models.BankAccountsGet
		bankAccounts, _ := h.service.BankAccounts.GetAllUser(userAccounts[index].IdUser)
		for index2, _ := range bankAccounts{
			var userBankAccount models.BankAccountsGet
			userBankAccount.NumberAccount = bankAccounts[index2].NumberAccount
			statusAccount, _ := h.service.Status.GetById(bankAccounts[index2].StatusId)
			userBankAccount.StatusName = statusAccount.NameStatus
			userBankAccounts = append(userBankAccounts,userBankAccount)
		}
		userAccountGetOne.UserAccounts = append(userAccountGetOne.UserAccounts,userBankAccounts...)
		roleUser,_ := h.service.Role.GetById(userAccounts[index].RoleUserId)
		userAccountGetOne.RoleUserName=roleUser.NameRole
		userAccountGetOne.Username = userAccounts[index].Username
		userAccountGetOne.Email = userAccounts[index].Email
		userAccountGetOne.IdUser = userAccounts[index].IdUser
		userAccountsGet = append(userAccountsGet,userAccountGetOne)
	}
	data := gin.H{
		"title": "Add user",
		"NameUser": userData.Name,
		"Users": userAccountsGet,
	}
	c.HTML(http.StatusOK,"add_user.html",data)
}


func (h *Handler) addUser (c *gin.Context){


	//userId, err := h.tokenParseToUserId("Token",c)

	//if err != nil{
	//	return
	//}

	Id,_ := c.GetPostForm("id")
	UserName,_ := c.GetPostForm("UserName")
	Password,_ := c.GetPostForm("Password")
	Email,_ := c.GetPostForm("Email")
	Role,_ := c.GetPostForm("Role")
	var input models.User
	IdOldUser, err := strconv.Atoi(Id)
	RoleUserId, err2 := strconv.Atoi(Role)
	if err == nil {
		var userUpdate models.UpdateUser
		userUpdate.Password = Password
		userUpdate.Username = UserName
		userUpdate.Email = Email
		userUpdate.RoleUserId = RoleUserId
		h.service.User.Update( IdOldUser,userUpdate)
	}else{
		if err2 == nil && Password!= "" && UserName!=""{
			input.Password = Password
			input.Username = UserName
			input.Email = Email
			input.RoleUserId = RoleUserId
			input.BankUser = 1
			h.service.Authorisation.CreateUser(input)
		}
	}
	c.Redirect(http.StatusMovedPermanently, "http://localhost:9090/bank_admin/create")
}
