package model

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInfo struct {
	Login    string
	Password string
}

type UserSetting struct {
	Id          int64  `json:"id"`
	Nickname    string `json:"nickname"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

type PasswordUpdate struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"oldPassword"`
}
