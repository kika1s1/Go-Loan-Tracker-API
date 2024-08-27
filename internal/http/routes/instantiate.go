package routes

import (
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/http/handlers/account"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/repository/mongodb"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/usecase/loan"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/usecase/user"
	localMongo "github.com/kika1s1/Go-Loan-Tracker-API/pkg/mongo"
)



func InstantaiteUserHandler() *account.UserHandler {
	usersCollection := localMongo.UserCollection
	userRepo := &mongodb.UserRepositoryMongo{Collection: usersCollection}
	userUsecase := user.NewUserUsecase(userRepo)
	userHandler := account.NewUserHandler(userUsecase)
	return userHandler
}


func InstantaiteLoandHandler() *account.LoanHandler {
	loansCollection := localMongo.LoanCollection
	loanRepo := &mongodb.LoanRepositoryMongo{Collection: loansCollection}
	loanUsecase := loan.NewLoanUseCase(loanRepo)
	loanHandler := account.NewLoanHandler(loanUsecase)
	return loanHandler
}
