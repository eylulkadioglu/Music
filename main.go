package main

import (
	"fmt"
	"os"

	"github.com/eylulkadioglu/Music/appconfig"
	"github.com/eylulkadioglu/Music/db"
	"github.com/eylulkadioglu/Music/mw"
	"github.com/eylulkadioglu/Music/routes"
	"github.com/eylulkadioglu/Music/salt"
	"github.com/gin-gonic/gin"
)

func main() {
	config := appconfig.ReadConfig()
	if config.DbDSN == "" {
		fmt.Printf("Database configuration not found\n")
		os.Exit(-1)
	}

	salt.SetSalt(config.Salt)
	db.InitDB()

	ginEn := gin.Default()

	ginEn.GET("/", routes.Landing)
	ginEn.POST("/login", routes.Login)
	ginEn.POST("/lostPassword", routes.LostPassword)
	ginEn.POST("/changePassword", routes.ChangePassword)
	ginEn.Use(mw.CheckAuthorization)
	ginEn.POST("/user", routes.CreateUser)
	ginEn.GET("/artists", routes.GetArtists)
	ginEn.POST("/artists/add", routes.CreateArtist)
	ginEn.DELETE("/artists/delete", routes.DeleteArtist)

	ginEn.Run(config.ListenAddress)
}
