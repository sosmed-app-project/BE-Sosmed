package handler

import "app-sosmed/features/users"

type UserRequest struct {
	Username     string `json:"username" form:"username"`
	Email        string `json:"email" form:"email"`
	TanggalLahir string `json:"tanggallahir" form:"tanggallahir"`
	NoHandphone  string `json:"nohandphone" form:"nohandphone"`
	Password     string `json:"password" form:"password"`
}

func UserRequestToCore(input UserRequest) users.UserCore {
	var userCore = users.UserCore{
		Username:     input.Username,
		TanggalLahir: input.TanggalLahir,
		Email:        input.Email,
		NoHandphone:  input.NoHandphone,
		Password:     input.Password,
	}
	return userCore
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
