package user

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/kika1s1/Go-Loan-Tracker-API/internal/config"
	"github.com/kika1s1/Go-Loan-Tracker-API/pkg/email"
	"github.com/kika1s1/Go-Loan-Tracker-API/pkg/jwt"
)

func (u *UserUsecase) RequestPasswordResetUsecase(userEmail string) error {
	var emailSender email.EmailSender

	Config, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	emailProvider := Config.EMAIL_PROVIDER

	switch emailProvider {
	case "simple":
		emailSender = email.NewSimpleEmailSender(
			Config.SMTP_HOST,
			Config.EMAIL_PORT,
			Config.EMAIL_SENDER_EMAIL,
			Config.EMAIL_SENDER_PASSWORD,
		)
	default:
		emailSender = email.NewSimpleEmailSender(
			Config.SMTP_HOST,
			Config.EMAIL_PORT,
			Config.EMAIL_SENDER_EMAIL,
			Config.EMAIL_SENDER_PASSWORD,
		)
	}
	accessToken, err := jwt.GenerateJWT("password-reset", userEmail, "password-reset", "password-reset")
	if err != nil {
		return err
	}
	go func() {
		err := emailSender.SendPasswordResetEmail(userEmail, accessToken)
		if err != nil {
			log.Printf("Failed to send password reset email: %v", err)
		}
	}()

	return nil
}

func (u *UserUsecase) ResetPassword(resetToken string, newPassword string) error {
	claims, err := jwt.ValidateToken(resetToken)

	if err != nil {
		return err
	}
	fmt.Println(claims, "claims ====================")
	if claims.Role != "password-reset" || claims.UserID != "password-reset" {
		return errors.New("invalid token")
	}
	email := claims.Email
	// hashedPassword, err := hash.HashPassword(newPassword)
	if err != nil {
		return err
	}
	user, err := u.repo.FindUserByEmail(context.Background(), email)

	if user != nil && err == nil {
		// user == hashedPassword
		err = u.repo.UpdateUser(context.Background(), user)
		if err != nil {
			return err
		}
		return nil
	}
	return err
}
