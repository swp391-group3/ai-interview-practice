package auth

const (
	RefreshTokenCookie = "refresh"
	RefreshPath        = "/auth/refresh"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"jane@example.com"`
	Password string `json:"password" binding:"required" example:"SuperSecret123"`
}

type TokenPair struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}
