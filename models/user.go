package models

type User struct {
	IdUser int `json:"-" `
	Name string `json:"name" binding: "required"`
	Username string `json:"username" binding: "required"`
	Password string `json:"password" binding: "required"`
	RoleUserId int `json:"role_user_id"`
	BankUser int `json:"bank_user"`
	Token string `json:"token"`
	Email string `json:"email"`
}

type UpdateUser struct {
	Name string
	Username string
	Password string
	RoleUserId int
	BankUser int
	Token string
	Email string
}

type UserGet struct {
	IdUser int `json:"id"`
	Username string `json:"username" binding: "required"`
	RoleUserName string `json:"role_user_name"`
	Email string `json:"email"`
	UserAccounts []BankAccountsGet `json:"user_accounts"`
}
