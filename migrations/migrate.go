package main

import (
	"github.com/KRTirtho/spotube-matcher/bootstrappers"
	"github.com/KRTirtho/spotube-matcher/schemas"
)

func init() {
	bootstrappers.BootstrapEnvironment()
	bootstrappers.BootstrapDatabases()

}

func main() {
	bootstrappers.DB.AutoMigrate(&schemas.Post{})
}
