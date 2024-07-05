package repository

import (
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
	panic("unimplemented")
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(user model.User) {
	panic("unimplemented")
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(user model.User) {
	panic("unimplemented")
}

func NewUserRepsotoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}
