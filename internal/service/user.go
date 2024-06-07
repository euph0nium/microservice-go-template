package service

import "microservice-go-template/internal/model"

type UserService interface {
	GetUser(id string) (*model.User, error)
}
