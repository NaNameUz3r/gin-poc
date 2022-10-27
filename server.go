package main

import (
	"fmt"
	"gin-poc/controllers"
	"gin-poc/middlewares"
	"gin-poc/services"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    services.VideoService       = services.NewVideoService()
	videoController controllers.VideoController = controllers.NewVideoController(videoService)
)

func setupLogFile() {
	f, err := os.Create("gin.log")
	if err != nil {
		fmt.Println("[WARNING] Cannot create logfile, will use stdout.")
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

}

func main() {
	server := gin.New()
	setupLogFile()
	server.Use(gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
		gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.GetAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Saved successfully."})

	})
	server.Run(":8080")
}
