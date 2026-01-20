package main

import (
	"github.com/dtt4h/auth-service/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/health", handler.HealthCheck)
	r.Run()
}
