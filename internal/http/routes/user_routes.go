package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/http/middleware"
)

func RegisterUserRoutes(router *gin.Engine) {
	userHandler := InstantaiteUserHandler()
	userRoute := router.Group("/api/v1/users", middleware.AuthMiddleware())
	{
		userRoute.GET("/profile", userHandler.GetUser)
		userRoute.DELETE("/delete", userHandler.DeleteUser)
		userRoute.PUT("/password-update", userHandler.UpdateUser)
		userRoute.POST("/password-reset", userHandler.ResetPasswordRequest)
		userRoute.POST("/password-update", userHandler.ResetPassword)

	}

}
