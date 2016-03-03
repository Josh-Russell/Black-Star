package main

import (
	"fmt"
	"os"

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

	// Assign a sql database
	//	db, err := NewMySQLDB("root:P@ssw0rd! tcp(127.0.0.1:3306)/ProjectPegasus")
	//	if err != nil {
	//		panic(err)
	//	}
	//globalMySQLDB = db

	// Assign an image store
	//globalImageStore = NewDBImageStore()
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	fmt.Println(port)

	router := gin.New()
	router.Use(gin.Logger())

	if port != "3000" {
		router.LoadHTMLGlob("templates/*.tmpl.html")
		router.Static("/static", "static")
	} else {
		router.LoadHTMLGlob("../../templates/*.tmpl.html")
		router.Static("../../static", "static")
	}

	router.GET("/", HandleNavigateToHome)
	router.GET("/discover", HandleNavigateToDiscoverPage)
	router.GET("/upload", HandleNavigateToUpload)
	router.GET("/search", HandleNavigateToSearch)
	router.GET("/register", HandleNavigateToRegister)
	router.GET("/profile", HandleNavigateToProfile)
	router.GET("/login", HandleNavigateToLoginPage)
	router.GET("/viewVideo", HandleNavigateToVideoViewPage)

	router.POST("/logout", HandleSessionDestroy)
	router.POST("/login", HandleSessionCreate)
	router.POST("/register", HandleSessionNew)

	router.Run(":" + port)
}
