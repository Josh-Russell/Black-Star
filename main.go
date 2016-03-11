package main

import (
	"database/sql"
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

	//Assign a sql database
	//check if we can connect with on-campus ip
	var db *sql.DB
	fmt.Println("Attempting to connect on-campus...")
	db, err = NewMySQLDB("projp:password@tcp(10.10.14.54:3306)/projectpegasus")
	if err != nil {
		fmt.Println("failed to connect to on-camps.")
		fmt.Println("Attempting to connect off-campus...")
		//check if we can connect with off-campus ip
		db, err = NewMySQLDB("projp:password@tcp(69.27.22.79:3306)/projectpegasus")
		if err != nil {
			fmt.Println("failed to connect to the database..")
			panic(err)
		}
	}
	fmt.Println("Connected to database.\n")
	globalMySQLDB = db

	// Assign an image store
	globalVideoStore = NewDBVideoLocationStore()
}

func main() {

	port := "9419"

	router := gin.Default()

	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")
	router.Static("/videos", "videos")

	router.GET("/", HandleNavigateToHome)
	router.GET("/page:pageNumber", HandleNavigateToHomeWithPageNumber)
	router.GET("/discover", HandleNavigateToDiscoverPage)
	router.GET("/search", HandleNavigateToSearch)
	router.GET("/register", HandleNavigateToRegister)
	router.GET("/login", HandleNavigateToLoginPage)
	router.GET("/viewVideo:videoID", HandleNavigateToVideoViewPage)
	router.GET("/confirmAge:videoID", HandleNavigateToConfirmAge)
	router.GET("/logout", HandleSessionDestroy)
	router.GET("/upload", HandleNavigateToUpload)
	router.POST("/login", HandleSessionCreate)
	router.POST("/register", HandleSessionNew)
	router.POST("/upload", HandleNewVideo)

	router.Run(":" + port)
}
