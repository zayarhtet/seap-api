package dto

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Username     string `json:"username"`
	Role         string `json:"role"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
