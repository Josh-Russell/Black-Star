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

	//	//Assign a sql database
	db, err := NewMySQLDB("projp:password@tcp(127.0.0.1:3306)/projectpegasus")
	if err != nil {
		panic(err)
	}
	globalMySQLDB = db

	// Assign an image store
	globalVideoStore = NewDBVideoLocationStore()
}

func main() {

	port := "9419"

	fmt.Println(port)

	router := gin.Default()

	authorized := router.Group("/registered", gin.BasicAuth(gin.Accounts{"user": "password"}))

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

	authorized.GET("/profile", HandleNavigateToProfile)
	authorized.GET("/upload", HandleNavigateToUpload)

	router.Run(":" + port)
}
