package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Register struct {
	Email          string `form:"email" json:"email" binding:"required"`
	Password       string `form:"password" json:"password" binding:"required"`
	RetypePassword string `form:"repeat-password" json:"repeat-password" binding:"required"`
}

func HandleUserCreate(c *gin.Context) {
	var form Register

	if c.Bind(&form) == nil {
		if form.Email == "user" && form.RetypePassword == form.Password {
			c.HTML(http.StatusOK, "view.tmpl.html", nil)
			fmt.Println("it worked!")
		} else {
			fmt.Println("email worked: ", form.Email == "user")
			fmt.Println("password worked: ", form.RetypePassword == form.Password, form.RetypePassword, form.Password)
			fmt.Println("You didn't get logged in")
			c.HTML(http.StatusUnauthorized, "register.tmpl.html", nil)
		}
	}
}

func HandleNavigateToRegister(c *gin.Context) {
	session := RequestSession(c.Request)

	if session != nil {
		c.HTML(http.StatusOK, "register.tmpl.html", gin.H{"currentuser": session.UserID})
	} else {
		c.HTML(http.StatusOK, "register.tmpl.html", nil)
	}
}
