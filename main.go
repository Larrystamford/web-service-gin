package main

import (
	"os"
	"github.com/Larrystamford/web-service-gin/controller"
	"github.com/gin-gonic/gin"
)

var (
	webController controller.Controller = controller.NewController()
)

func main() {
    router := gin.Default()
    router.GET("/", webController.GetAlbums)
	router.GET("/albums/:id", webController.GetAlbumByID)
    router.POST("/albums", webController.PostAlbums)

    router.Run(":" + os.Getenv("PORT"))
}
