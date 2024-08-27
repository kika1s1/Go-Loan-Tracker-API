package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/http/middleware"
)

func RegisterUserRoutes(router *gin.Engine) {
	userHandler := InstantaiteUserHandler()
	userRoute := router.Group("/api/v1/accounts", middleware.AuthMiddleware())
	{
		userRoute.GET("/me", userHandler.GetUser)
		userRoute.DELETE("/me", userHandler.DeleteUser)
		userRoute.PUT("/me", userHandler.UpdateUser)
		userRoute.GET("/any/:id", userHandler.GetAnyUser)
	}

}
