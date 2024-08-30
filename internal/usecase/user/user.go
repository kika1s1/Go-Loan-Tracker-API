package user

import "github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"

type UserUseCaseInterface interface {
	RegisterUser(user *domain.User) (*domain.User, error)
	Login(email, password string) (*domain.User, *domain.Token, error)
	RequestEmailVerification(user domain.User) error
	RequestPasswordResetUsecase(userEmail string) error
	ResetPassword(token string, password string) error
	VerifyEmail(token string) error

	FindUserById(id string) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
	FindUserByUserName(username string) (*domain.User, error)

	UpdateUser(user *domain.User) error
	
	GetAllUsers() ([]*domain.GetUserDTO, error)
	
	DeleteUser(id string) error
	FilterUsers(filter map[string]interface{}) ([]*domain.User, error)

	GoogleCallback(code string) (*domain.User, *domain.Token, error)
}
