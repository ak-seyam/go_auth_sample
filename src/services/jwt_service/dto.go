package jwt_service

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	Groups []string `json:"groups"`
	jwt.RegisteredClaims
}
