package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type Authorisation interface {

	CreateUser(user models.User)(int, error)
	GenerateToken(username, password string)(int,string, error)
	ParseToken(token string)(int, error)
	SaveToken(userId int,token string)(error)
	TakeToken(userId int)(string,error)

}

type Application interface {
	Create(userId int, application models.Application)(int, error)
	GetAll(userId int)( []models.Application, error)
	GetById(userId, applicationId int)(models.Application, error)
	Delete(userId, applicationId int)( error)
	Update(userId,id int ,input models.UpdateApplication)( error)
	GetAllUser(token int)([]models.Application,error)
	GetAllUserResponse(token int)([]models.Application,error)
}

type Role interface {
	Create(role models.Role)(int, error)
	GetAllUser(userId int)( []models.Role, error)
	GetAll()([]models.Role, error)
	GetById(roleId int)(models.Role, error)
	Delete(roleId int)( error)
	Update(roleId int ,input models.UpdateRole)( error)
}

type RoleRights interface {
	Create(role models.RoleRight)(int, error)
	GetAllRole(roleId int)( []models.RoleRight, error)
	GetAll()([]models.RoleRight, error)
	GetById(roleId int)(models.RoleRight, error)
	Delete(roleId int)( error)
	Update(roleId int ,input models.UpdateRoleRight)( error)
}

type Rights interface {
	Create(right models.Right)(int, error)
	GetAll()([]models.Right, error)
	GetById(rightId int)(models.Right, error)
	Delete(rightId int)( error)
	Update(rightId int ,input models.UpdateRight)( error)
}

type BankAccounts interface{
	Create(account models.BankAccounts)(int, error)
	GetAll()([]models.BankAccounts, error)
	GetById(accountId int)(models.BankAccounts, error)
	Update(accountId int ,input models.UpdateBankAccounts)( error)
	GetAllUser(userId int)( []models.BankAccounts, error)
}

type Status interface {
	Create(status models.Status)(int, error)
	GetAll()([]models.Status, error)
	GetById(statusId int)(models.Status, error)
	Update(statusId int ,input models.UpdateStatus)( error)
	Delete(statusId int)( error)
}

type Scores interface {
	Create(score models.Scores)(int, error)
	GetAll()([]models.Scores, error)
	GetById(scoreId int)(models.Scores, error)
	Update(scoreId int ,input models.UpdateScores)( error)
	GetAllAccount(accountId int)( []models.Scores, error)
}

type StatusScore interface {
	Create(statusScore models.StatusScore)(int, error)
	GetAll()([]models.StatusScore, error)
	GetById(statusScoreId int)(models.StatusScore, error)
	Update(statusScoreId int ,input models.UpdateStatusScore)( error)
	Delete(statusScoreId int)( error)
}

type ScoresMoney interface {
	Create(scoreMoney models.ScoresMoney)(int, error)
	GetAll()([]models.ScoresMoney, error)
	GetById(scoreMoneyId int)(models.ScoresMoney, error)
	Update(scoreMoneyId int ,input models.UpdateScoresMoney)( error)
	GetAllScore(scoreId int)( []models.ScoresMoney, error)
}

type Money interface {
	Create(money models.Money)(int, error)
	GetAll()([]models.Money, error)
	GetById(moneyId int)(models.Money, error)
	Update(moneyId int ,input models.UpdateMoney)( error)
}

type Currency interface {
	Create(currency models.Currency)(int, error)
	GetAll()([]models.Currency, error)
	GetById(currencyId int)(models.Currency, error)
	Update(currencyId int ,input models.UpdateCurrency)( error)
}

type Card interface {
	Create(card models.Card)(int, error)
	GetAll()([]models.Card, error)
	GetById(cardId int)(models.Card, error)
	Update(cardId int ,input models.UpdateCard)( error)
	GetAllAccount(accountId int)( []models.Card, error)
}

type StatusCard interface {
	Create(status models.StatusCard)(int, error)
	GetAll()([]models.StatusCard, error)
	GetById(statusId int)(models.StatusCard, error)
	Update(statusId int ,input models.UpdateStatusCard)( error)
	Delete(statusId int)( error)
}

type Reminder interface {
	Create(reminder models.Reminder)(int, error)
	GetAll()([]models.Reminder, error)
	GetById(reminderId int)(models.Reminder, error)
	Update(reminderId int ,input models.UpdateReminder)( error)
	Delete(reminderId int)( error)
}

type User interface {
	Create(user models.User)(int, error)
	GetAll()([]models.User, error)
	GetById(userId int)(models.User, error)
	Update(userId int ,input models.UpdateUser)( error)
	Delete(userId int)( error)
	GetByToken(token string)(int,error)
}

type ApplicationTitle interface {
	Create(applicationTitle models.ApplicationTitle)(int, error)
	GetAll()([]models.ApplicationTitle, error)
	GetById(applicationTitleId int)(models.ApplicationTitle, error)
	Delete(applicationTitleId int)( error)
	Update(applicationTitleId int ,input models.UpdateApplicationTitle)( error)
}
type Profile interface {
	GetProfile(token string)(models.User, []models.Application)
}
type Service struct {
	Authorisation
	Application
	Role
	RoleRights
	Rights
	BankAccounts
	Status
	Scores
	StatusScore
	ScoresMoney
	Money
	Currency
	Card
	StatusCard
	Reminder
	User
	ApplicationTitle
	Profile
}

func NewService(repos *repository.Repository) *Service{
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
		Application: NewApplicationService(repos.Application),
		Role: NewRoleService(repos.Role),
		RoleRights: NewRoleRightsService (repos.RoleRight),
		Rights: NewRightsService (repos.Rights),
		BankAccounts: NewBankAccountsService (repos.BankAccounts),
		Status: NewStatusService (repos.Status),
		Scores: NewScoresService(repos.Scores),
		StatusScore: NewStatusScoreService (repos.StatusScore),
		ScoresMoney: NewScoresMoneyService (repos.ScoresMoney),
		Money: NewMoneyService (repos.Money),
		Currency: NewCurrencyService (repos.Currency),
		Card: NewCardService (repos.Card),
		StatusCard: NewStatusCardService(repos.StatusCard),
		Reminder: NewReminderService(repos.Reminder),
		User: NewUserService (repos.User),
		ApplicationTitle: NewApplicationTitleService(repos.ApplicationTitle),
		Profile: NewProfileService(repos.Profile),
	}
}