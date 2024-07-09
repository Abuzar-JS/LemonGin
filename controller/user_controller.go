package controller

import (
	"golang-crud/data/request"
	"golang-crud/data/response"
	"golang-crud/helper"
	"golang-crud/service"
	"net/http"
	"strconv"

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
	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ReturnError(err)

	newUser := controller.UserService.Create(createUserRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: gin.H{
			"id":   newUser.Id,
			"Name": newUser.Name,
		},
	}
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)

}

// Update Controller
func (controller *UserController) Update(ctx *gin.Context) {
	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ReturnError(err)

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ReturnError(err)
	updateUserRequest.Id = id

	controller.UserService.Update(updateUserRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   updateUserRequest,
	}
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)
}

// Delete Controller
func (controller *UserController) Delete(ctx *gin.Context) {

	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ReturnError(err)
	controller.UserService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userId,
	}
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(http.StatusOK, webResponse)
}

// FindById Controller
func (controller *UserController) FindById(ctx *gin.Context) {
	userId := ctx.Param("userId")
	id, err := strconv.Atoi(userId)
	helper.ReturnError(err)

	userResponse := controller.UserService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByAll Controller
func (controller *UserController) FindByAll(ctx *gin.Context) {
	userResponse := controller.UserService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   userResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
