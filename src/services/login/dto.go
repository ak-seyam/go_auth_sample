package login

import "time"

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginError struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

func NewLoginError(message string) LoginError {
	return LoginError{
		Message: message,
		Date:    time.Now(),
	}
}
