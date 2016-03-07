package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func HandleNavigateToLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl.html", nil)
}
