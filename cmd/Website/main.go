package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/heroku/go-getting-started/Godeps/_workspace/src/github.com/gin-gonic/gin"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	fmt.Println(port)

	router := gin.New()
	//router.Use(gin.Logger())
	router.LoadHTMLGlob("C:/Users/JRussell/Documents/GoProjects/src/github.com/heroku/Project Pegasus/templates/*.tmpl.html")

	router.Static("/../static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.GET("/discover", func(c *gin.Context) {
		c.HTML(http.StatusOK, "discover.tmpl.html", nil)
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
	router.GET("/viewVideo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "viewVideo.tmpl.html", nil)
	})

	router.POST("/register", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if c.Bind(&form) == nil {
			if form.User == "user" && form.Password == "123" {
				c.HTML(http.StatusOK, "view.tmpl.html", nil)
				fmt.Println("it worked!")
			} else {
				c.HTML(http.StatusUnauthorized, "profile.tmpl.html", nil)
				fmt.Println("You didn't get logged in")
			}
		}
	})
	router.Run(":" + port)
}
