package main

import (
    "net/http"

    "github.com/gin-gonic/gin"

	"web-service-gin/controller"
)

var (
	controller controller.Controller = controller.NewController()
)


func main() {
    router := gin.Default()
    router.GET("/albums", controller.getAlbums)
	router.GET("/albums/:id", controller.getAlbumByID)
    router.POST("/albums", controller.postAlbums)

    router.Run("localhost:8080")
}
