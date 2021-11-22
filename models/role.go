package models

type Role struct {
	IdRole int `json:"id_role"`
	NameRole string `json:"name_role"`
}

type RoleRight struct {
	IdRolesRights int `json:"id_role_right"`
	RoleId int `json:"role_id"`
	RightId int `json:"right_id"`
}

type Right struct {
	IdRights int `json:"id_right"`
	NameOperation string `json:"name_operation"`
	Status bool `json:"status"`
}

type UpdateRole struct {
	NameRole string `json:"name_role"`
}

type UpdateRoleRight struct {
	RoleId int `json:"role_id"`
	RightId int `json:"right_id"`
}

type UpdateRight struct {
	NameOperation string `json:"name_operation"`
	Status bool `json:"status"`
}
