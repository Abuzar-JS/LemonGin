package router

import (
	"golang-crud/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserController) *gin.Engine {

	router := gin.Default()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome User")
	})

	baseRouter := router.Group("/api")
	userRouter := baseRouter.Group("/user")
	userRouter.GET("", userController.FindByAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.POST("", userController.Create)
	userRouter.PATCH("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)

	return router
}
