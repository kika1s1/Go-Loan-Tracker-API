package account

import (
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	// "github.com/kika1s1/Go-Loan-Tracker-API/internal/domain"
)



func (h *UserHandler) AdminGetUser(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not found"})
		return
	}
	user, err := h.UserUsecase.FindUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) FilterUsers(c *gin.Context) {
	// Define valid filters
	validFilters := []string{"role", "email", "username", "firstName", "lastName"}

	// Create a map to hold the filters
	filters := make(map[string]interface{})

	// Loop through query parameters and add valid ones to the filters map
	for _, key := range validFilters {
		if value := c.Query(key); value != "" {
			filters[key] = value
		}
	}

	// Call the usecase to filter users based on the provided filters
	users, err := h.UserUsecase.FilterUsers(filters)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the filtered users
	c.JSON(http.StatusOK, users)
}
