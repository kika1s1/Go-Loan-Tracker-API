package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/http/middleware"
)

func RegisterAdminUserRoutes(router *gin.Engine) {
	userHandler := InstantaiteUserHandler()



	adminRoute := router.Group("/api/v1/admin/users")
	adminRoute.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{

		adminRoute.GET("/", userHandler.GetAllUsers)
		adminRoute.DELETE("/:id", userHandler.DeleteUser)
		
	}
}
