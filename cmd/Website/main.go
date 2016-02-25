package main

import (
	"log"
	"net/http"
	"os"

	"github.com/heroku/go-getting-started/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")

	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.GET("/view", func(c *gin.Context) {
		c.HTML(http.StatusOK, "view.tmpl.html", nil)
	})
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.tmpl.html", nil)
	})
	router.GET("/search", func(c *gin.Context) {
		c.HTML(http.StatusOK, "search.tmpl.html", nil)
	})
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.tmpl.html", nil)
	})
	router.GET("/profile", func(c *gin.Context) {
		c.HTML(http.StatusOK, "profile.tmpl.html", nil)
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl.html", nil)
	})

	router.Run(":" + port)
}
