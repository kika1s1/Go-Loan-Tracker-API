package loan

import "github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"

type LoanUseCaseInterface interface {
	ApplyForLoan(loan domain.Loan) (*domain.Loan, error)
	ViewLoanStatus(id string) (*domain.Loan, error)
	ViewAllLoans(status string, order string) ([]domain.Loan, error)
	ApproveRejectLoan(id string, status string) (*domain.Loan, error)
	DeleteLoan(id string) error
	ViewSystemLogs() ([]domain.Log, error)
}