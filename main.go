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
	bootstrappers.BootstrapRateLimiter()
}

func main() {
	router := gin.Default()

	router.Use(bootstrappers.RateLimiterMiddleware)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Message": "Hello World"})
	})

	router.GET("/posts", routes.GetPosts)
	router.GET("/posts/:id", routes.GetPost)
	router.POST("/post", routes.CreatePost)

	router.GET("/track/:spotifyId", routes.GetTrack)
	router.PUT("/track/:id/upvote", routes.UpvoteTrack)
	router.PUT("/track/:id/downvote", routes.DownvoteTrack)
	router.POST("/track", routes.AddTrack)
	log.Fatal(router.Run())
}
