package controller

import (
	"golang-crud/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		UserService: service,
	}

}

// Create Controller
func (controller *UserController) Create(ctx *gin.Context) {

}

// Update Controller
func (controller *UserController) Update(ctx *gin.Context) {

}

// Delete Controller
func (controller *UserController) Delete(ctx *gin.Context) {

}

// FindById Controller
func (controller *UserController) FindById(ctx *gin.Context) {

}

// FindByAll Controller
func (controller *UserController) FindByAll(ctx *gin.Context) {

}
