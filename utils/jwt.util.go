package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const secret = "secret-key"

type jwtCustomClaims struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	RoleName string `json:"role_name"`
	jwt.StandardClaims
}

func GenerateAccessToken(id int, email string, username string, roleName string) (string, error) {
	claims := &jwtCustomClaims{
		Id:       id,
		Email:    email,
		Username: username,
		RoleName: roleName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
