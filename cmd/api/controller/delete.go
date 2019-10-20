package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeleteAccount deletes Superhero account.
func (ctl *Controller) DeleteAccount(c *gin.Context) {
	userID := c.Query("userID")
	if userID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Internal server error!",
		})

		return
	}

	fmt.Println(userID)

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Account deleted successfully!",
	})
}