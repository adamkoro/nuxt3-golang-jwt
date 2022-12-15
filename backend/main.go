package main

import (
	"nuxt3-golang-jwt-backend/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/ping", api.Ping)
	}

	router.Run("127.0.0.1:8080")
}
