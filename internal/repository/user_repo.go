package repository

import (
	"context"

	"github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	FindUserByUserName(ctx context.Context, username string) (*domain.User, error)
	FindUserById(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id string) error

	GetAllUsers(ctx context.Context) ([]*domain.GetUserDTO, error)
	FilterUsers(ctx context.Context, filter map[string]interface{}) ([]*domain.User, error)

	IsEmptyCollection(ctx context.Context) (bool, error)

	RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error)

	GoogleCallback(ctx context.Context, code string) (*domain.User, error)
}
