package user

import (
	"api-dev/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func CreateUser(c *gin.Context) {
	request := CreateUserRequest{}
	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := request.Validate(); err != nil {
		sendError(c, http.StatusBadRequest, "CreatUser:"+err.Error())
		return
	}
	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error handling user input")
		log.Printf("CreateUser: Failed to hash password: %v", err)
		return
	}

	user := schemas.User{
		Username: request.Username,
		Password: string(hashedPassword),
		Email:    request.Email,
		Role:     "user",
	}

	if err := db.Create(&user).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error creating user:"+err.Error())
		return
	}
	sendSuccess(c, "CreateUser", user)
}
