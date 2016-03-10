package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	// Assign a user store
	store, err := NewFileUserStore("./data/users.json")
	if err != nil {
		panic(fmt.Errorf("Error creating user store: %s", err))
	}
	globalUserStore = store

	// Assign a session store
	sessionStore, err := NewFileSessionStore("./data/sessions.json")
	if err != nil {
		panic(fmt.Errorf("Error creating session store: %s", err))
	}
	globalSessionStore = sessionStore

	//Assign an image store
	//	imagestore, err := NewFileSessionStore(".data/imageInformation.json")
	//	if err != nil {
	//		panic(fmt.Errorf("Error creating session store: %s", err))
	//	}
	//	globalImageStore = imagestore

	//Assign a sql database
	//check if we can connect with on-campus ip
	db, err := NewMySQLDB("projp:password@tcp(69.27.22.79:3306)/projectpegasus")
	if err != nil {
		//check if we can connect with off-campus ip
		//db, err := NewMySQLDB("projp:password@tcp(69.27.22.79:3306)/projectpegasus")
		if err != nil {
			panic(err)
		}
	}
	globalMySQLDB = db

	// Assign an image store
	globalImageStore = NewDBImageStore()
}

func main() {

	port := "9419"

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", HandleNavigateToHome)
	router.GET("/discover", HandleNavigateToDiscoverPage)
	router.GET("/search", HandleNavigateToSearch)
	router.GET("/register", HandleNavigateToRegister)
	router.GET("/login", HandleNavigateToLoginPage)
	router.GET("/viewVideo", HandleNavigateToVideoViewPage)
	router.GET("/logout", HandleSessionDestroy)
	router.GET("/upload", HandleNavigateToUpload)
	router.POST("/login", HandleSessionCreate)
	router.POST("/register", HandleSessionNew)
	router.POST("/upload", HandleNewVideo)

	router.Run(":" + port)
}
