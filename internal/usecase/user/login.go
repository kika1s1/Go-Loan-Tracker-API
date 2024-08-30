package user

import (
	"context"
	"errors"

	"github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"
	"github.com/kika1s1/Go-Loan-Tracker-API/pkg/hash"
	"github.com/kika1s1/Go-Loan-Tracker-API/pkg/jwt"
)

func (u *UserUsecase) Login(email string, password string) (*domain.User, *domain.Token, error) {
	user, err := u.repo.FindUserByEmail(context.TODO(), email)

	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, nil, errors.New("invalid credentials")
	}
	if !user.Verified {
		return nil, nil, errors.New("you need to verify your email first")
	}

	if !hash.CheckPasswordHash(password, user.Password) {
		return nil, nil, errors.New("invalid credentials")

	}
	// hashedPassword, err := hash.HashPassword(password)
	// if err != nil {
	// 	return nil, nil, err

	// }

	// if user.Password != hashedPassword {
	// 	return nil, nil, errors.New("invalid credentials")
	// }
	accessToken, err := jwt.GenerateJWT(user.ID.Hex(), user.UserName, user.Email, user.Role)

	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := jwt.GenerateRefreshToken(user.ID.Hex(), user.UserName, user.Email, user.Role)
	if err != nil {
		return nil, nil, err
	}

	return user, &domain.Token{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}
