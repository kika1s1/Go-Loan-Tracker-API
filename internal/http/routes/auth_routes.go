package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/http/handlers/token"
)

func RegisterVerificationRoutes(router *gin.Engine) {

	userHandler := InstantaiteUserHandler()
	authRoutes := router.Group("/api/v1/users")

	{
		authRoutes.POST("/login", userHandler.Login)
		authRoutes.POST("/register", userHandler.Register)
		authRoutes.GET("/google_redirect", userHandler.GoogleCallback)
		authRoutes.POST("/verify-email", userHandler.RequestVerifyEmail)
		authRoutes.GET("/verify/confirm", userHandler.VerifyEmail) //I used this naming to make things clear
		

		//logout and refresh
		tokenHandler := &token.TokenHandler{}
		r2 := authRoutes.Group("/")
		r2.POST("/logout", tokenHandler.LogOut)
		r2.POST("/token/refresh", tokenHandler.RefreshToken)
	}
}
