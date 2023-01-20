package jwt_service

import (
	"time"

	"github.com/A-Siam/go_auth/src/services/utils"
	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(claims JWTClaims, expirationInMinutes int32) (string, error) {
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(expirationInMinutes) * time.Minute))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.Issuer = "auth"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(utils.GetEnv("JWT_SECRET")))
}
