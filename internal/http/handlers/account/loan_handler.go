package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/usecase/loan"
)

type LoanHandler struct {
	LoanUseCase loan.LoanUseCaseInterface
}

func NewLoanHandler(uc loan.LoanUseCaseInterface) *LoanHandler {
	return &LoanHandler{LoanUseCase: uc}
}


func (h *LoanHandler) ApplyForLoan(c *gin.Context) {
	var loan domain.Loan
	if err := c.ShouldBindJSON(&loan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.LoanUseCase.ApplyForLoan(loan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *LoanHandler) ViewLoanStatus(c *gin.Context) {
	id := c.Param("id")

	result, err := h.LoanUseCase.ViewLoanStatus(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *LoanHandler) ViewAllLoans(c *gin.Context) {
	status := c.Query("status")
	order := c.Query("order")

	result, err := h.LoanUseCase.ViewAllLoans(status, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *LoanHandler) ApproveRejectLoan(c *gin.Context) {
	id := c.Param("id")
	status := c.Query("status")

	result, err := h.LoanUseCase.ApproveRejectLoan(id, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *LoanHandler) DeleteLoan(c *gin.Context) {
	id := c.Param("id")

	err := h.LoanUseCase.DeleteLoan(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan deleted successfully"})
}


func (h *LoanHandler) ViewSystemLogs(c *gin.Context) {
	logs, err := h.LoanUseCase.ViewSystemLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}