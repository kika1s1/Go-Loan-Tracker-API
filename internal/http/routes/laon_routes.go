package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/http/middleware"
)

func RegisterLoanRoutes(router *gin.Engine) {
	loanHandler := InstantaiteLoandHandler()

	r := router.Group("api/v1/loans")
	{
		r.POST("/", middleware.AuthMiddleware(), loanHandler.ApplyForLoan)
		r.GET("/:id", middleware.AuthMiddleware(), loanHandler.ViewLoanStatus)
	}

	admin := router.Group("api/v1/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.GET("/loans", loanHandler.ViewAllLoans)
		admin.PATCH("/loans/:id/status", loanHandler.ApproveRejectLoan)
		admin.DELETE("/loans/:id", loanHandler.DeleteLoan)
		admin.GET("/logs", loanHandler.ViewSystemLogs)
	}
}
