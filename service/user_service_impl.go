package service

import (
	"github.com/go-playground/validator/v10"
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	validate       *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{UserRepository: userRepository, validate: validate}
}

func (u *UserServiceImpl) Create(user request.CreateUserRequest) {
	err := u.validate.Struct(user)
	helper.ReturnError(err)
	userModel := model.User{
		Name: user.Name,
	}
	u.UserRepository.Save(userModel)

}

func (u *UserServiceImpl) Delete(userId int) {
	u.UserRepository.Delete(userId)
}

func (u *UserServiceImpl) FindAll() []response.UserResponse {
	result := u.UserRepository.FindAll()

	var users []response.UserResponse

	for _, value := range result {
		user := response.UserResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		users = append(users, user)
	}
	return users
}

func (u *UserServiceImpl) FindById(userId int) response.UserResponse {
	user, err := u.UserRepository.FindById(userId)
	helper.ReturnError(err)
	userResponse := response.UserResponse{
		Id:   user.Id,
		Name: user.Name,
	}
	return userResponse
}

func (u *UserServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := u.UserRepository.FindById(user.Id)
	helper.ReturnError(err)
	userData.Name = user.Name
	u.UserRepository.Update(userData)
}
