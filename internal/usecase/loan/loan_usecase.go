package loan

import (
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/repository"
)

type LoanUseCase struct {
	repo repository.LoanRepository
}


func NewLoanUseCase(repo repository.LoanRepository) *LoanUseCase {
	return &LoanUseCase{repo: repo}
}

func (uc *LoanUseCase) ApplyForLoan(loan domain.Loan) (*domain.Loan, error) {
	_, err := uc.repo.CreateLoan(loan)
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (uc *LoanUseCase) ViewLoanStatus(id string) (*domain.Loan, error) {
	return uc.repo.FindLoanByID(id)
}

func (uc *LoanUseCase) ViewAllLoans(status string, order string) ([]domain.Loan, error) {
	return uc.repo.FindAllLoans(status, order)
}

func (uc *LoanUseCase) ApproveRejectLoan(id string, status string) (*domain.Loan, error) {
	_, err := uc.repo.UpdateLoanStatus(id, status)
	if err != nil {
		return nil, err
	}
	return uc.repo.FindLoanByID(id)
}

func (uc *LoanUseCase) DeleteLoan(id string) error {
	_, err := uc.repo.DeleteLoan(id)
	return err
}

// ViewSystemLogs implements LoanUseCaseInterface.
func (uc *LoanUseCase) ViewSystemLogs() ([]domain.Log, error) {
	return uc.repo.FindAllLogs()
}