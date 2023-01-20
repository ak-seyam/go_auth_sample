package login

import (
	"errors"
	"log"
	"strings"

	"github.com/A-Siam/go_auth/src/data/dao"
	"github.com/A-Siam/go_auth/src/services/jwt_service"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, password string) (*LoginResponse, error) {
	user, err := dao.GetUser(username)
	if err != nil {
		log.Printf("Error while retrieving user {%s} errorMsg {%s}\n", username, err.Error())
		return nil, errors.New("Cannot fetch user")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Fatal("Invalid password")
		return nil, errors.New("Invalid password")
	}
	groups := strings.Split(user.Groups, ",")
	accessToken, err := jwt_service.CreateJWT(jwt_service.JWTClaims{
		Groups: groups,
		RegisteredClaims: jwt.RegisteredClaims{
			ID: user.ID,
		},
	}, 15)
	if err != nil {
		log.Printf("Cannot create access jwt for user {%s} errorMsg {%s}\n", username, err.Error())
		return nil, errors.New("Error happened")
	}
	refreshToken, err := jwt_service.CreateJWT(jwt_service.JWTClaims{
		Groups: groups,
		RegisteredClaims: jwt.RegisteredClaims{
			ID: user.ID,
		},
	}, 15)
	if err != nil {
		log.Printf("Cannot create access jwt for user {%s} errorMsg {%s}\n", username, err.Error())
		return nil, errors.New("Error happened")
	}
	loginResponse := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return &loginResponse, nil
}
