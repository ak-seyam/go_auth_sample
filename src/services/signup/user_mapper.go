package signup

import (
	"strings"

	"github.com/A-Siam/go_auth/src/data/entities"
)

func MapToSignUpResponse(user entities.User) SignUpResponse {
	return SignUpResponse{
		ID:       user.ID,
		Username: user.Username,
		Groups:   strings.Split(user.Groups, ","),
	}
}
