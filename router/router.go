package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Setup() {
	router := gin.Default()
	routes(router)
	log.Fatal(router.Run(":8080"))
}
