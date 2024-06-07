package api

import (
	"microservice-go-template/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserService
}

func (h *UserHandler) GetUser(c *gin.Context) {
}
