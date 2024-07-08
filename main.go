package main

import (
	"golang-crud/config"
	"golang-crud/controller"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
	"golang-crud/router"
	"golang-crud/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started Server!")

	//Database

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("user").AutoMigrate(&model.User{})

	userRepository := repository.NewUserRepositoryImpl(db)

	userService := service.NewUserServiceImpl(userRepository, validate)

	userController := controller.NewUserController(userService)

	//Router
	routes := router.NewRouter(userController)

	server := &http.Server{
		Addr:    ":8800",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ReturnError(err)

}
