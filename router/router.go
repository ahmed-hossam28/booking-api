package router

import (
	"github.com/ahmed-hossam28/config"
	"github.com/gin-gonic/gin"
)

func Setup() {
	logger := config.NewSingletonLogger()
	router := gin.Default()
	routes(router)
	logger.Fatal(router.Run(":8080"))
}
