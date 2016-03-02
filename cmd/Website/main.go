package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// keep the users logged in for 3 days
	sessionLength     = 24 * 3 * time.Hour
	sessionCookieName = "GophrSession"
	sessionIDLength   = 20
)

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

	router.POST("/login", HandleUserLogin)
	router.POST("/register", HandleUserCreate)

	router.Run(":" + port)
}
