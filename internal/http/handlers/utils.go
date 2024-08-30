package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"
)

func GetClaims(c *gin.Context) (*domain.Claims, error) {
	claims, exists := c.Get("claims")
	if !exists {
		return nil, errors.New("user not found")
	}
	userClaims, ok := claims.(*domain.Claims)
	if !ok {
		return nil, errors.New("user not found")
	}
	return userClaims, nil
}

func GetUserId(c *gin.Context) (string, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return "", err
	}
	return claims.UserID, nil
}

func GetEmail(c *gin.Context) (string, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return "", err
	}
	return claims.Email, nil
}

func GetRole(c *gin.Context) (string, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return "", err
	}
	return claims.Role, nil
}
