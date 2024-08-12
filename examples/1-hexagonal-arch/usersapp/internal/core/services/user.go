package services

import "usersapp/internal/ports"

type UserService struct {
	userPort ports.UserPort
}

func NewUserService(userPort ports.UserPort) *UserService {
	return &UserService{userPort: userPort}
}
