package controllers

import (
	"gin-tutorial/dtos"
	"gin-tutorial/finders"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, dtos.ResultDTO{Result: dtos.Albums})
}

func PostAlbum(c *gin.Context) {

	var newAlbumDTO dtos.AlbumDTO

	err := c.BindJSON(&newAlbumDTO)
	if err != nil {
		panic("I cant bind this!")
	}

	dtos.Albums = append(dtos.Albums, newAlbumDTO)

	c.IndentedJSON(http.StatusCreated, dtos.ResultDTO{Result: newAlbumDTO})
}

func GetAlbum(c *gin.Context) {

	albumID := c.Param("id")

	albumDTO := finders.FindAlbumByIdInAlbums(albumID)
	if albumDTO == nil {
		c.IndentedJSON(http.StatusNotFound, dtos.NewNotFoundDTO())
		return
	}

	c.IndentedJSON(http.StatusOK, dtos.ResultDTO{Result: albumDTO})
}
