package user

import (
	"github.com/ahmed-hossam28/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "id[int] empty")
		return
	}
	user := schemas.User{}
	if err := db.First(&user, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "user not found")
		return
	}
	if err := db.Delete(&user).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "Error deleting user")
		return
	}
	sendSuccess(c, "delete user", user)
}
