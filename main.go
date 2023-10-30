package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//albums slice to seed record album record data

var albums = []album{
	{ID: "1", Title: "blue train", Artist: "Foyez", Price: 1.8},
	{ID: "2", Title: "red train", Artist: "Abdul", Price: 2.8},
	{ID: "3", Title: "yellow train", Artist: "mumin", Price: 3.8},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {

	router := gin.Default()
	router.GET("/Albums", getAlbums)
	router.POST("/Albums", postAlbums)
	router.GET("/Albums/:id", postAlbums)
	router.Run("localhost:8080")
}

func postAlbums(c *gin.Context) {

	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
