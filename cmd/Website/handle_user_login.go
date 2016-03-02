package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func HandleUserLogin(c *gin.Context) {
	var form Login
	// This will infer what binder to use depending on the content-type header.
	if c.Bind(&form) == nil {
		if form.Email == "user" && form.Password == "123" {
			c.HTML(http.StatusOK, "index.tmpl.html", nil)
			fmt.Println("it worked!")
		} else {
			c.HTML(http.StatusUnauthorized, "login.tmpl.html", nil)
			fmt.Println("You didn't get logged in")
		}
	}
}
func HandleNavigateToLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl.html", nil)
}
