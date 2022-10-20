package schemas

import "gorm.io/gorm"

type Post struct {
	gorm.Model

	Owner       string
	Description string
	Title       string
}
