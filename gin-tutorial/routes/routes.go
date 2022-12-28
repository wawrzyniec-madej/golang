package routes

import (
	"gin-tutorial/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/albums", controllers.GetAlbums)

	router.POST("/albums", controllers.PostAlbum)
	router.GET("/albums/:id", controllers.GetAlbum)
}
