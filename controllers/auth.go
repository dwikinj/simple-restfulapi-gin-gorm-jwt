package controllers

import (
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context) {
	var input LoginInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	user := models.User{}
	user.Username = input.Username
	user.Password = input.Password

	token, errLogin := models.LoginCheck(user.Username, user.Password)
	if errLogin != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})

}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(ctx *gin.Context) {
	var registerInput RegisterInput

	if err := ctx.ShouldBindJSON(&registerInput); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := &models.User{}
	u.Username = registerInput.Username
	u.Password = registerInput.Password

	_, err := u.SaveUser()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "registration success!"})
}

func CurrentUser(ctx *gin.Context) {
	user, ok := ctx.Get("currentUser")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})

}
