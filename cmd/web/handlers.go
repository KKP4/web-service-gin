package main

import (
	"encoding/json"
	"example.com/web-service-gin/pkg/models"
	"example.com/web-service-gin/pkg/validation"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getAllAlbumsHandler(c *gin.Context) {

	result, err := models.GetAllAlbums()

	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"albums": result})

}

func insertAlbumsHandler(c *gin.Context) {
	var input struct {
		Title  string  `json:"title"`
		Artist string  `json:"artist"`
		Price  float64 `json:"price"`
	}
	err := json.NewDecoder(c.Request.Body).Decode(&input)

	if err != nil {
		fmt.Println(err)
	}
	var album = &validation.AlbumSchema{Title: input.Title, Artist: input.Artist, Price: input.Price}
	err = validation.ValidateAlbumSchema(album)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	err = models.InsertAlbum(input.Title, input.Artist, input.Price)
	if err != nil {
		fmt.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
