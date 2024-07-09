package service

import (
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/model"
)

type UserService interface {
	Create(user request.CreateUserRequest) model.User
	Update(user request.UpdateUserRequest)
	Delete(userId int)
	FindById(userId int) response.UserResponse
	FindAll() []response.UserResponse
}
