package service

import "microservice-go-template/internal/model"

type UserService interface {
	GetUser(id string) (*model.User, error)
}

type UserServiceImpl struct {
}

func (impl *UserServiceImpl) GetUser(id string) (*model.User, error) {
	return &model.User{Name: "test"}, nil
}

func NewUserService() UserService {
	return &UserServiceImpl{}
}
