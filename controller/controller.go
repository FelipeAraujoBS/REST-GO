package controller

import (
	"net/http"
	"rest-go/models"

	"github.com/gin-gonic/gin"
)

var Albums = []models.Album{
	{ID: "1", Title: "Mothers Milk", Artist: "Red Hot Chili Peppers", Price: 39.99},
	{ID: "2", Title: "Californication", Artist: "Red Hot Chili Peppers", Price: 49.99},
	{ID: "3", Title: "By the Way", Artist: "Red Hot Chili Peppers", Price: 29.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Albums)
}
