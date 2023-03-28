package main

import (
	"cdec-api/pkg/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // create server
	router.POST("/calculate", handler.Calculate)
	router.Run("localhost:8080") //listen  on host 8080
}
