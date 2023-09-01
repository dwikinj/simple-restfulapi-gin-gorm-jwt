package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// UserClaim adalah struktur khusus untuk menyimpan informasi pengguna dalam token
type UserClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken adalah fungsi untuk membuat token JWT berdasarkan informasi pengguna
func GenerateToken(username string) (string, error) {
	var secretKey = []byte("secret")

	claims := UserClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenStr.SignedString(secretKey)
	return token, err
}

// VerifyToken adalah fungsi untuk memverifikasi token JWT dan mengembalikan klaim pengguna
func VerifyToken(tokenString string) (*UserClaim, error) {
	var secretKey = []byte("secret")
	var claims UserClaim
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrInvalidKey
	}
	return &claims, nil
}
