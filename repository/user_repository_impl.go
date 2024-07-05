package repository

import (
	"errors"
	"golang-crud/data/request"
	"golang-crud/helper"
	"golang-crud/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u *UserRepositoryImpl) Delete(userId int) {
	var user model.User
	result := u.Db.Where("id = ?", userId).Delete(&user)
	helper.ReturnError(result.Error)
}

func (u *UserRepositoryImpl) FindAll() []model.User {
	var user []model.User
	result := u.Db.Find(&user)
	helper.ReturnError(result.Error)
	return user
}

func (u *UserRepositoryImpl) FindById(userId int) (User model.User, err error) {
	var user model.User
	result := u.Db.Find(&user, userId)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}

}

func (u *UserRepositoryImpl) Save(user model.User) {
	result := u.Db.Create(&user)
	helper.ReturnError(result.Error)

}

func (u *UserRepositoryImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		Id:   user.Id,
		Name: user.Name,
	}

	result := u.Db.Model(&user).Updates(updateUser)
	helper.ReturnError(result.Error)
}
