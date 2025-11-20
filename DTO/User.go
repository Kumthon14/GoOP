package DTO

type UserLoginResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone" `
	Address  string `json:"address"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" `
	Address  string `json:"address"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
