package main

import (
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/controllers"
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/middlewares"
	"github.com/dwikinj/simple-restfulapi-gin-gorm-jwt/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")
}
