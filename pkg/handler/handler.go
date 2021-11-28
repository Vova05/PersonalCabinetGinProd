package handler

import (
	"PersonalCabinetGin/pkg/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.Default()

	router.StaticFS("/css",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css"))
	router.StaticFS("/fonts",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css\\fonts"))
	router.StaticFS("/img",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\img"))
	router.StaticFS("/js",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\js"))
	router.StaticFS("/less",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\less"))

	router.LoadHTMLGlob("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/*.html")

	//auth := router.Group("/auth")
	//{
	//	auth.Static("/css","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css")
	//	auth.Static("/fonts","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/assets/css/fonts")
	//	auth.Static("/images","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/assets/images")
	//	auth.Static("/js","C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates/js")
	//
	//	auth.POST("/sign-up", h.signUp)
	//
	//	auth.POST("/sign-in", h.signIn)
	//	auth.GET("/sign-in",h.signInGet)
	//
	//}

	router.POST("/auths",  h.signIn)
	router.GET("/auths",h.signInGet)
	router.GET("/test", h.getProfile)
	test := router.Group("/testg/id/ррр")
	{
		test.StaticFS("/css",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css"))
		test.StaticFS("/fonts",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css\\fonts"))
		test.StaticFS("/img",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\img"))
		test.StaticFS("/js",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\js"))
		test.StaticFS("/less",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\less"))


		test.GET("/",h.getProfileUser)
	}
	bank_admin := router.Group("/bank_admin", h.userIdentity)
	{
		bank_admin.StaticFS("/css",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css"))
		bank_admin.StaticFS("/fonts",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css\\fonts"))
		bank_admin.StaticFS("/img",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\img"))
		bank_admin.StaticFS("/js",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\js"))
		bank_admin.StaticFS("/less",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\less"))

		bank_admin.GET("/profile",h.getProfileUser)
		bank_admin.POST("/profile", h.createApplication)

		bank_admin.GET("/bank_account", h.getBankAccount)

		bank_admin.GET("/bank_scores",h.getScores)

		create_admin := bank_admin.Group("/create")
		{
			create_admin.StaticFS("/css",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css"))
			create_admin.StaticFS("/fonts",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\css\\fonts"))
			create_admin.StaticFS("/img",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\img"))
			create_admin.StaticFS("/js",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\js"))
			create_admin.StaticFS("/less",http.Dir("C:\\Users\\VovaGlh\\GolandProjects\\PersonalCabinetGin1\\pkg\\handler\\templates\\less"))

			create_admin.GET("/", h.addUserGet)
			create_admin.POST("/post",h.addUser)
		}
		message := bank_admin.Group("/messages")
		{
			message.GET("/", )
			message.POST("/",)
		}

		application := bank_admin.Group("/applications")
		{
			application.GET("/", h.getAllApplication)
			application.GET("/:id",h.getApplicationById)
			application.POST("/create", h.createApplication)
			application.PUT("/:id", h.updateApplication)
			application.DELETE("/:id",h.deleteApplication)
		}
		role := bank_admin.Group("/role")
		{
			role.GET("/", h.getAllRoles)
			role.GET("/:id",h.getRoleById)
			role.GET("/user",h.getAllRolesUser)
			role.POST("/create", h.createRole)
			role.PUT("/:id", h.updateRole)
			role.DELETE("/:id",h.deleteRole)
		}
		roles_rights := bank_admin.Group("/role_rights")
		{
			roles_rights.GET("/", h.getAllRoleRights)
			roles_rights.GET("/:id",h.getRoleRightsById)
			roles_rights.GET("/role",h.getAllRoleRightsRole)
			roles_rights.POST("/create", h.createRoleRights)
			roles_rights.PUT("/:id", h.updateRoleRights)
			roles_rights.DELETE("/:id",h.deleteRoleRights)
		}
		rights := bank_admin.Group("/rights")
		{
			rights.GET("/", h.getAllRights)
			rights.GET("/:id",h.getRightsById)
			rights.POST("/create", h.createRights)
			rights.PUT("/:id", h.updateRights)
			rights.DELETE("/:id",h.deleteRights)//нужно обратить внимание на сущности которые зависят от этой
		}
		bank_accounts := bank_admin.Group("/bank_accounts")
		{
			bank_accounts.GET("/", h.getAllBankAccounts)
			bank_accounts.GET("/user", h.getAllBankAccountsUser)
			bank_accounts.GET("/:id",h.getBankAccountsById)
			bank_accounts.POST("/create", h.createBankAccounts)
			bank_accounts.PUT("/:id", h.updateBankAccounts)
		}

		status := bank_admin.Group("/status")
		{
			status.GET("/", h.getAllStatus)
			status.GET("/:id",h.getStatusById)
			status.POST("/create", h.createStatus)
			status.PUT("/:id", h.updateStatus)
			status.DELETE("/:id", h.deleteStatus)
		}
		scores := bank_admin.Group("/scores")
		{
			scores.GET("/", h.getAllScores)
			scores.GET("/:id",h.getScoresById)
			scores.GET("/user", h.getAllScoresUser)
			scores.POST("/create", h.createScores)
			scores.PUT("/:id", h.updateScores)
		}

		status_score := bank_admin.Group("/status_score")
		{
			status_score.GET("/", h.getAllStatusScore)
			status_score.GET("/:id",h.getStatusScoreById)
			status_score.POST("/create", h.createStatusScore)
			status_score.PUT("/:id", h.updateStatusScore)
			status_score.DELETE("/:id", h.deleteStatusScore)
		}

		score_money := bank_admin.Group("/score_money")
		{
			score_money.GET("/", h.getAllScoresMoney)
			score_money.GET("/:id",h.getScoresMoneyById)
			score_money.GET("/user", h.getAllScore)
			score_money.POST("/create", h.createScoresMoney)
			score_money.PUT("/:id", h.updateScoresMoney)
		}

		money := bank_admin.Group("/money")
		{
			money.GET("/", h.getAllMoney)
			money.GET("/:id",h.getMoneyById)
			money.POST("/create", h.createMoney)
			money.PUT("/:id", h.updateMoney)
		}

		currency := bank_admin.Group("/currency")
		{
			currency.GET("/", h.getAllCurrency)
			currency.GET("/:id",h.getCurrencyById)
			currency.POST("/create", h.createCurrency)
			currency.PUT("/:id", h.updateCurrency)
		}

		card := bank_admin.Group("/cards")
		{
			card.GET("/", h.getAllCard)
			card.GET("/:id",h.getCardById)
			card.GET("/user", h.getAllCardUser)
			card.POST("/create", h.createCard)
			card.PUT("/:id", h.updateCard)
		}

		status_card := bank_admin.Group("/status_card")
		{
			status_card.GET("/", h.getAllStatusCard)
			status_card.GET("/:id",h.getStatusCardById)
			status_card.POST("/create", h.createStatusCard)
			status_card.PUT("/:id", h.updateStatusCard)
			status_card.DELETE("/:id",h.deleteStatusCard)
		}

		reminder := bank_admin.Group("/reminder")
		{
			reminder.GET("/", h.getAllReminder)
			reminder.GET("/:id",h.getReminderById)
			reminder.POST("/create", h.createReminder)
			reminder.PUT("/:id", h.updateReminder)
			reminder.DELETE("/:id",h.deleteReminder)
		}

		users := bank_admin.Group("/users")
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
