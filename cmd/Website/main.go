package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email          string `form:"email" json:"email" binding:"required"`
	Password       string `form:"password" json:"password" binding:"required"`
	RetypePassword string `form:"retype-password" json:retype-password"`
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
		router.Static("/static", "static")
	}

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

	router.POST("/login", func(c *gin.Context) {
		var form Login
		// This will infer what binder to use depending on the content-type header.
		if c.Bind(&form) == nil {
			if form.Email == "user" && form.Password == "123" {
				c.HTML(http.StatusOK, "view.tmpl.html", nil)
				fmt.Println("it worked!")
			} else {
				c.HTML(http.StatusUnauthorized, "profile.tmpl.html", nil)
				fmt.Println("You didn't get logged in")
			}
		}
	})
	router.POST("/register", func(c *gin.Context) {
		var form Login

		if c.Bind(&form) == nil {
			if form.Email == "user" && form.RetypePassword == form.Password {
				c.HTML(http.StatusOK, "view.tmpl.html", nil)
				fmt.Println("it worked!")
			} else {
				fmt.Println("email worked: ", form.Email == "user")
				fmt.Println("password worked: ", form.RetypePassword == form.Password)
				fmt.Println("You didn't get logged in")
				c.HTML(http.StatusUnauthorized, "view.tmpl.html", nil)
			}
		}
	})
	router.Run(":" + port)
}
