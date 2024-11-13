package user

import (
	"github.com/ahmed-hossam28/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []schemas.User
	if err := db.Find(&users).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error fetching users")
		return
	}
	sendSuccess(c, "list users", users)
}
