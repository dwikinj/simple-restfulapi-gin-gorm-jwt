package middlewares

import (
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/models"
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := utils.ExtractToken(ctx)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		verifiedToken, err := utils.VerifyToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		var user models.User
		user, err = models.FindByUsername(verifiedToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		ctx.Set("currentUser", user)
		ctx.Next()
	}

}
