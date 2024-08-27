package account

import (

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"
	"github.com/kika1s1/Go-Loan-Tracker-API/internal/usecase/user"
	"github.com/kika1s1/Go-Loan-Tracker-API/pkg/checker"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UserUsecase user.UserUseCaseInterface
}

func NewUserHandler(uc user.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		UserUsecase: uc,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	var Login *domain.LoginUserDTO

	if err := c.ShouldBindJSON(&Login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, token, err := h.UserUsecase.Login(Login.Email, Login.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	getUser := domain.GetUserDTO{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		Profile: user.Profile,
	}

	// set cookie
	c.SetCookie("x_refresh_token", token.RefreshToken, 60*60*24, "/", "localhost", false, true)
	c.SetCookie("x_access_token", token.AccessToken, 60*60*24, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user":  getUser,
		"token": token,
	})


}

func (h *UserHandler) Register(c *gin.Context) {
	var user *domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}
	if err := checker.IsValidEmail(user.Email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := checker.IsValidPassword(user.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UserUsecase.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Verification email is being sent to " + user.Email + " activate your account!",
	})

}

func (h *UserHandler) GetUser(c *gin.Context) {

	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userClaims, ok := claims.(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	dbUser, err := h.UserUsecase.FindUserById(userClaims.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Me := domain.GetUserDTO{
		ID:      dbUser.ID,
		Email:   dbUser.Email,
		Role:    dbUser.Role,
		Profile: dbUser.Profile,
	}
	c.JSON(http.StatusOK, Me)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	userClaims, ok := claims.(*domain.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := h.UserUsecase.DeleteUser(userClaims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}



func (h *UserHandler) UpdateUser(c *gin.Context) {
    var updateUserDTO domain.UpdateUserDTO
    if err := c.BindJSON(&updateUserDTO); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    claims, exists := c.Get("claims")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }
    userClaims, ok := claims.(*domain.Claims)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    objectID, err := primitive.ObjectIDFromHex(userClaims.UserID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
        return
    }

    dbUser := domain.User{
        ID:       objectID,
        UserName: updateUserDTO.UserName,
        Email:    updateUserDTO.Email,
        Profile:  updateUserDTO.Profile,
    }

    err = h.UserUsecase.UpdateUser(&dbUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}


func (h *UserHandler) GetAnyUser(c *gin.Context) {
	userId := c.Param("userId")
	user, err := h.UserUsecase.FindUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	GetUser := domain.GetUserDTO{
		UserName: user.UserName,
		Email:    user.Email,
		Role:     user.Role,
		Profile:  user.Profile,
	}
	c.JSON(http.StatusOK, GetUser)
}

