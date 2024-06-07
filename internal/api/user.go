package api

import (
	"microservice-go-template/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserService
}

func (h *UserHandler) GetUser(c *gin.Context) {
	user, _ := h.UserService.GetUser("")
	c.JSON(http.StatusOK, user)
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{UserService: service}
}
