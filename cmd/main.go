package main

import (
	"microservice-go-template/internal/api"
	"microservice-go-template/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// init db, config
	userService := service.NewUserService()
	userHandler := api.NewUserHandler(userService)

	r := gin.Default()
	r.GET("/ping", userHandler.GetUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
