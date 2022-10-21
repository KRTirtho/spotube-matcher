package routes

import (
	"github.com/KRTirtho/spotube-matcher/bootstrappers"
	"github.com/KRTirtho/spotube-matcher/schemas"
	"github.com/gin-gonic/gin"
)

func GetTrack(ctx *gin.Context) {
	spotifyId := ctx.Param("spotifyId")
	var data []schemas.Track

	result := bootstrappers.DB.Where(&schemas.Track{SpotifyId: spotifyId}).Find(&data)

	if result.Error != nil {
		ctx.String(400, result.Error.Error())
		return
	}

	if len(data) == 0 {
		ctx.JSON(404, gin.H{"msg": "No tracks were found"})
		return
	}

	ctx.JSON(200, data)

}

func AddTrack(ctx *gin.Context) {
	var data schemas.Track
	ctx.BindJSON(&data)
	result := bootstrappers.DB.Create(&data)
	if result.Error != nil {
		ctx.String(400, result.Error.Error())
		return
	}
	ctx.JSON(200, data)
}

func UpvoteTrack(ctx *gin.Context) {
	id := ctx.Param("id")
	var data schemas.Track
	result := bootstrappers.DB.First(&data, id)
	if result.Error != nil {
		ctx.String(400, result.Error.Error())
		return
	}

	data.Upvote++

	result = bootstrappers.DB.Save(&data)

	if result.Error != nil {
		ctx.String(400, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{"msg": "No tracks were found"})
		return
	}

	ctx.JSON(200, gin.H{"msg": "+1 for " + id})
}

func DownvoteTrack(ctx *gin.Context) {
	id := ctx.Param("id")
	var data schemas.Track
	result := bootstrappers.DB.First(&data, id)
	if result.Error != nil {
		ctx.String(400, result.Error.Error())
		return
	}

	data.Downvote++

	result = bootstrappers.DB.Save(&data)

	if result.Error != nil {
		ctx.String(400, result.Error.Error())
		return
	}

	if result.RowsAffected == 0 {
		ctx.JSON(404, gin.H{"msg": "No tracks were found"})
		return
	}

	ctx.JSON(200, gin.H{"msg": "-1 for " + id})
}
