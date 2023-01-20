package signup

import (
	"log"
	"net/http"
	"strings"

	"github.com/A-Siam/go_auth/src/data/dao"
	"github.com/A-Siam/go_auth/src/data/system_errors"
	"golang.org/x/crypto/bcrypt"
)

func SingUp(singUpBody SignUpRequest) (SignUpResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(singUpBody.Password), 0)
	if err != nil {
		log.Printf("something went wrong while hashing password {%s}\n", err.Error())
		return SignUpResponse{}, system_errors.New(
			"Cannot create a user now please try again later",
			http.StatusInternalServerError,
		)
	}
	user, err := dao.SaveUser(singUpBody.Username, string(hashedPassword), strings.Join(singUpBody.Groups, ","))
	if err != nil {
		log.Printf("error storing user with username {%s} errorMsg {%s}\n", singUpBody.Username, err.Error())
		return SignUpResponse{}, system_errors.New(
			"Cannot create a user now please try again later",
			http.StatusInternalServerError,
		)
	}
	return MapToSignUpResponse(user), nil
}
