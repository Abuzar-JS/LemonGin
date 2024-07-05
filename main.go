package main

import (
	"golang-crud/helper"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Started Server!")

	routes := gin.Default()

	routes.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Get",
		})
	})

	routes.POST("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Post",
		})
	})

	routes.DELETE("", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Delete",
		})
	})

	server := &http.Server{
		Addr:    ":8800",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ReturnError(err)

}
