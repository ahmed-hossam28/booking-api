package router

import (
	"github.com/ahmed-hossam28/handlers/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func index(c *gin.Context) {
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "Welcome to Booking api",
	//})
	c.HTML(http.StatusOK, "index.html", nil)
}
func routes(router *gin.Engine) {
	router.GET("/", index)
	//User EndPoints
	router.GET("/user", user.GetUsers)
	router.POST("/user", user.CreateUser)
	router.GET("/user/:id", user.GetUserByID)
	router.PUT("/user/:id", user.UpdateUser)
	router.DELETE("/user/:id", user.DeleteUser)

}
