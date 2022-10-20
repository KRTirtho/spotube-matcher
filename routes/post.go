package routes

import (
	"github.com/KRTirtho/spotube-matcher/bootstrappers"
	"github.com/KRTirtho/spotube-matcher/schemas"
	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	var body struct {
		Title       string `json:"title"`
		Owner       string `json:"owner"`
		Description string `json:"description"`
	}
	ctx.Bind(&body)

	operation := bootstrappers.DB.Create(&schemas.Post{Owner: body.Owner, Description: body.Description, Title: body.Title})

	if operation.Error != nil {
		ctx.Error(operation.Error)
		ctx.String(400, operation.Error.Error())
		return
	}

	ctx.JSON(201, map[string]string{
		"title":       body.Title,
		"owner":       body.Owner,
		"description": body.Description,
	})
}

func GetPosts(ctx *gin.Context) {
	var posts []schemas.Post
	bootstrappers.DB.Find(&posts)

	ctx.JSON(200, posts)
}

func GetPost(ctx *gin.Context) {
	var post schemas.Post
	bootstrappers.DB.First(&post, ctx.Param("id"))

	ctx.JSON(200, post)
}
