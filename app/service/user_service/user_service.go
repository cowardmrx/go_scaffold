package user_service

import (
	"go_scaffold/app/http/request/user_request"
	"go_scaffold/app/model"
	"go_scaffold/app/repository"
)

type userService struct{}

type UserService interface {
	AddUser(request *user_request.UserAddRequest) error
}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) AddUser(request *user_request.UserAddRequest) error {
	userRepo := repository.NewUserRepository()

	return userRepo.Add(&model.Model{
		UserModel: &model.UserModel{
			Name: request.Name,
			Age:  request.Age,
		},
	})

}
