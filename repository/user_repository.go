package repository

import (
	"golang-crud/model"
)

type UserRepository interface {
	Save(user *model.User)
	Update(user model.User)
	Delete(userId int)
	FindById(userId int) (User model.User, err error)
	FindAll() []model.User
}
