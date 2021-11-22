package models

type User struct {
	IdUser int `json:"-" `
	Name string `json:"name" binding: "required"`
	Username string `json:"username" binding: "required"`
	Password string `json:"password" binding: "required"`
	RoleUserId int `json:"role_user_id"`
}

type UpdateUser struct {
	Name string
	Username string
	Password string
	RoleUserId int
}
