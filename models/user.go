package models

import (
	"errors"
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/service"
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/utils"
	"gorm.io/gorm"
	"html"
	"strings"
)

type User struct {
	gorm.Model
	Username string `gorm:"varchar(255);not null;unique" json:"username"`
	Password string `gorm:"varchar(255);not null" json:"password"`
}

func (u *User) SaveUser() (*User, error) {
	if u == nil {
		return nil, errors.New("user is nill")
	}
	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil

}

func (u *User) BeforeSave(db *gorm.DB) error {
	hashedPassword, _ := utils.HashPassword(u.Password)

	u.Password = hashedPassword
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil
}

func LoginCheck(username string, password string) (string, error) {
	user := User{}
	err := DB.Where("username = ?", username).Take(&user).Error

	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return "", errors.New("username or password wrong")
	}

	tokenString, _ := service.GenerateToken(user.Username)
	return tokenString, nil

}

func FindByUsername(username string) (user User, err error) {
	err = DB.Where("username = ?", username).Take(&user).Error
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}
