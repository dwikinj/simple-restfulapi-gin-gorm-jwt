package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/models"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type UserClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(user *models.User) (token string, err error) {
	var secretKey = []byte("secret")

	claims := UserClaim{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ = tokenStr.SignedString(secretKey)
	return token, err
}

func ExtractToken(ctx *gin.Context) string {
	var token string

	authorization := ctx.Request.Header.Get("Authorization")
	authorizationField := strings.Fields(authorization)

	if len(authorizationField) == 0 && authorizationField[0] != "Bearer" {
		return ""
	} else {
		token = authorizationField[1]
		return token
	}

}

func VerifyToken(tokenStr string) (string, error) {

	keyFun := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
		}
		return []byte("secret"), nil
	}

	var claims UserClaim

	tok, err := jwt.ParseWithClaims(tokenStr, claims, keyFun)
	if err != nil && !tok.Valid {
		return "", fmt.Errorf("invalid token")
	}
	_, err = models.FindByUsername(claims.Username)
	if err != nil && !tok.Valid {
		return "", fmt.Errorf("error :%v", err.Error())
	}

	return claims.Username, nil

}
