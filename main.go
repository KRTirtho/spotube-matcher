package main

import (
	"log"

	"github.com/KRTirtho/spotube-matcher/bootstrappers"
	"github.com/KRTirtho/spotube-matcher/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	bootstrappers.BootstrapEnvironment()
	bootstrappers.BootstrapDatabases()
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Message": "Hello World"})
	})
	router.GET("/posts", routes.GetPosts)
	router.GET("/posts/:id", routes.GetPost)
	router.POST("/post", routes.CreatePost)
	log.Fatal(router.Run())
}
