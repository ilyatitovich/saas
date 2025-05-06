package routes

import (
	"saas/auth/internal/handlers"
	"saas/auth/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/register",
			middlewares.ValidateRegisterInput(),
			handlers.Register)
	}
}
