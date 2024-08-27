package repository

import (
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoanRepository interface {
	CreateLoan(loan domain.Loan) (*mongo.InsertOneResult, error)
	FindLoanByID(id string) (*domain.Loan, error)
	FindAllLoans(status string, order string) ([]domain.Loan, error)
	UpdateLoanStatus(id string, status string) (*mongo.UpdateResult, error)
	DeleteLoan(id string) (*mongo.DeleteResult, error)
	FindAllLogs() ([]domain.Log, error)
}