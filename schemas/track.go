package schemas

import "gorm.io/gorm"

const (
	ALGORITHM_YOUTUBE_TOP        string = "youtube-top"
	ALGORITHM_AUTHOR_POPULAR     string = "popular-from-author"
	ALGORITHM_POPULAR_ACCURATELY string = "accurately-popular"
)

type Track struct {
	gorm.Model
	YoutubeId string `json:"youtubeId" gorm:"type:varchar(15);not null;uniqueIndex:yt-sp-al"`
	SpotifyId string `json:"spotifyId" gorm:"type:varchar(30);not null;uniqueIndex:yt-sp-al"`
	Algorithm string `json:"algorithm" gorm:"type:varchar(30);not null;uniqueIndex:yt-sp-al;check:algorithm IN ('youtube-top', 'popular-from-author', 'accurately-popular')"`
	Upvote    int    `json:"upvote" gorm:"not null;default:0"`
	Downvote  int    `json:"downvote" gorm:"not null;default:0"`
}
