package middlewares

import (
	"fmt"
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/controllers"
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := controllers.ExtractToken(ctx)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			fmt.Println("From jwt")
			return
		}
		user, err := service.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}

}
