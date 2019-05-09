package api

type AdminLoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type AdminLoginResponse struct {
	Token string `json:"token"`
}

func NewAdminLoginResponse(token string) *AdminLoginResponse {
	return &AdminLoginResponse{
		Token: token,
	}
}
