package handler

import (
	"PersonalCabinetGin/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	router.Static("/css","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css")
	router.Static("/fonts","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/css/fonts")
	router.Static("/images","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/images")
	router.Static("/js","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/js")

	router.LoadHTMLGlob("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/*.html")

	auth := router.Group("/auth")
	{
		auth.Static("/css","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css")
		auth.Static("/fonts","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/assets/css/fonts")
		auth.Static("/images","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/assets/images")
		auth.Static("/js","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/js")

		auth.POST("/sign-up", h.signUp)

		auth.POST("/sign-in", h.signIn)
		auth.GET("/sign-in",h.signInGet)
	}

	api :=router.Group("/api",h.userIdentity)
	{
		application := api.Group("/applications")
		{
			application.GET("/", h.getAllApplication)
			application.GET("/:id",h.getApplicationById)
			application.POST("/create", h.createApplication)
			application.PUT("/:id", h.updateApplication)
			application.DELETE("/:id",h.deleteApplication)
		}
		role := api.Group("/role")
		{
			role.GET("/", h.getAllRoles)
			role.GET("/:id",h.getRoleById)
			role.GET("/user",h.getAllRolesUser)
			role.POST("/create", h.createRole)
			role.PUT("/:id", h.updateRole)
			role.DELETE("/:id",h.deleteRole)
		}
		roles_rights := api.Group("/role_rights")
		{
			roles_rights.GET("/", h.getAllRoleRights)
			roles_rights.GET("/:id",h.getRoleRightsById)
			roles_rights.GET("/role",h.getAllRoleRightsRole)
			roles_rights.POST("/create", h.createRoleRights)
			roles_rights.PUT("/:id", h.updateRoleRights)
			roles_rights.DELETE("/:id",h.deleteRoleRights)
		}
		rights := api.Group("/rights")
		{
			rights.GET("/", h.getAllRights)
			rights.GET("/:id",h.getRightsById)
			rights.POST("/create", h.createRights)
			rights.PUT("/:id", h.updateRights)
			rights.DELETE("/:id",h.deleteRights)//нужно обратить внимание на сущности которые зависят от этой
		}
		bank_accounts := api.Group("/bank_accounts")
		{
			bank_accounts.GET("/", h.getAllBankAccounts)
			bank_accounts.GET("/user", h.getAllBankAccountsUser)
			bank_accounts.GET("/:id",h.getBankAccountsById)
			bank_accounts.POST("/create", h.createBankAccounts)
			bank_accounts.PUT("/:id", h.updateBankAccounts)
		}

		status := api.Group("/status")
		{
			status.GET("/", h.getAllStatus)
			status.GET("/:id",h.getStatusById)
			status.POST("/create", h.createStatus)
			status.PUT("/:id", h.updateStatus)
			status.DELETE("/:id", h.deleteStatus)
		}
		scores := api.Group("/scores")
		{
			scores.GET("/", h.getAllScores)
			scores.GET("/:id",h.getScoresById)
			scores.GET("/user", h.getAllScoresUser)
			scores.POST("/create", h.createScores)
			scores.PUT("/:id", h.updateScores)
		}

		status_score := api.Group("/status_score")
		{
			status_score.GET("/", h.getAllStatusScore)
			status_score.GET("/:id",h.getStatusScoreById)
			status_score.POST("/create", h.createStatusScore)
			status_score.PUT("/:id", h.updateStatusScore)
			status_score.DELETE("/:id", h.deleteStatusScore)
		}

		score_money := api.Group("/score_money")
		{
			score_money.GET("/", h.getAllScoresMoney)
			score_money.GET("/:id",h.getScoresMoneyById)
			score_money.GET("/user", h.getAllScore)
			score_money.POST("/create", h.createScoresMoney)
			score_money.PUT("/:id", h.updateScoresMoney)
		}

		money := api.Group("/money")
		{
			money.GET("/", h.getAllMoney)
			money.GET("/:id",h.getMoneyById)
			money.POST("/create", h.createMoney)
			money.PUT("/:id", h.updateMoney)
		}

		currency := api.Group("/currency")
		{
			currency.GET("/", h.getAllCurrency)
			currency.GET("/:id",h.getCurrencyById)
			currency.POST("/create", h.createCurrency)
			currency.PUT("/:id", h.updateCurrency)
		}

		card := api.Group("/cards")
		{
			card.GET("/", h.getAllCard)
			card.GET("/:id",h.getCardById)
			card.GET("/user", h.getAllCardUser)
			card.POST("/create", h.createCard)
			card.PUT("/:id", h.updateCard)
		}

		status_card := api.Group("/status_card")
		{
			status_card.GET("/", h.getAllStatusCard)
			status_card.GET("/:id",h.getStatusCardById)
			status_card.POST("/create", h.createStatusCard)
			status_card.PUT("/:id", h.updateStatusCard)
			status_card.DELETE("/:id",h.deleteStatusCard)
		}

		reminder := api.Group("/reminder")
		{
			reminder.GET("/", h.getAllReminder)
			reminder.GET("/:id",h.getReminderById)
			reminder.POST("/create", h.createReminder)
			reminder.PUT("/:id", h.updateReminder)
			reminder.DELETE("/:id",h.deleteReminder)
		}

		users := api.Group("/users")
		{
			users.GET("/", h.getAllUser)
			users.GET("/:id",h.getUserById)
			users.POST("/create", h.createUser)
			users.PUT("/:id", h.updateUser)
			users.DELETE("/:id",h.deleteUser)
		}
	}
	return router
}
