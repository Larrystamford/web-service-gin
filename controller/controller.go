package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetAlbums(c *gin.Context)
	GetAlbumByID(c *gin.Context)
	PostAlbums(c *gin.Context)
}

type controller struct{}

func NewController() Controller {
	return &controller{}
}

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "4", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "5", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "6", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func (*controller) GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func (*controller) GetAlbumByID(c *gin.Context) {
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

// postAlbums adds an album from JSON received in the request body.
func (*controller) PostAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
