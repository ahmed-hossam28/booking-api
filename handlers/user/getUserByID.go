package user

import (
	"api-dev/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	re, _ := regexp.Compile(`[0-9]+`)
	if !re.MatchString(id) {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "int [path parameter], format /{id}").Error())
		return
	}
	user := schemas.User{}
	if err := db.First(&user, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "User Not Found")
		return
	}

	sendSuccess(c, "getUserByID", user)
}
