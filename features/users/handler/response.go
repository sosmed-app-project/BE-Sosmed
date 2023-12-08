package handler

import "app-sosmed/features/users"

type UserResponse struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	TanggalLahir string `json:"tanggallahir"`
	Email        string `json:"email"`
	NoHandphone  string `json:"nohandphone"`
	Password     string `json:"password"`
}

func UserCoreToResponseAll(input users.UserCore) UserResponse {
	var userResp = UserResponse{
		Username:     input.Username,
		TanggalLahir: input.TanggalLahir,
		Email:        input.Email,
		NoHandphone:  input.NoHandphone,
	}
	return userResp
}

type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
