package service

import (
	"PersonalCabinetGin/models"
	"PersonalCabinetGin/pkg/repository"
)

type Authorisation interface {

	CreateUser(user models.User)(int, error)
	GenerateToken(username, password string)(string, error)
	ParseToken(token string)(int, error)
}

type Application interface {
	Create(userId int, application models.Application)(int, error)
	GetAll(userId int)( []models.Application, error)
	GetById(userId, applicationId int)(models.Application, error)
	Delete(userId, applicationId int)( error)
	Update(userId,id int ,input models.UpdateApplication)( error)
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
type Service struct {
	Authorisation
	Application
	Role
	RoleRights
	Rights
	BankAccounts
	Status
	Scores
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
	}
}