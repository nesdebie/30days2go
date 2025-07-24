package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type albumType struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []albumType{
    {ID: "1", Title: "Jane Doe", Artist: "Converge", Price: 56.99},
    {ID: "2", Title: "Mutter", Artist: "Rammstein", Price: 17.99},
    {ID: "3", Title: "Suicide Season", Artist: "Bring Me The Horizon", Price: 39.99},
	{ID: "4", Title: "Wanderlust", Artist: "Carpathian", Price: 666.0},
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, album := range albums {
        if album.ID == id {
            c.IndentedJSON(http.StatusOK, album)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
    var newAlbum albumType

    err := c.BindJSON(&newAlbum); 
	if err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func main() {
    router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
