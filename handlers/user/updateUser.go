package user

import (
	"github.com/ahmed-hossam28/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	re, _ := regexp.Compile(`[0-9]+`)
	if !re.MatchString(id) {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "int [path parameter], format /{id}").Error())
		return
	}
	newData := UpdateUserRequest{}
	c.BindJSON(&newData)
	if err := newData.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	user := schemas.User{}
	if err := db.First(&user, id).Error; err != nil {
		sendError(c, http.StatusNotFound, err.Error())
		return
	}
	if err := db.Model(&user).Updates(newData).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "Error updating opening")
		return
	}
	sendSuccess(c, "update user data", user)
}
