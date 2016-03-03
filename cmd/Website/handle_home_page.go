package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToHome(c *gin.Context) {
	session := RequestSession(c.Request)

	if session != nil {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"currentuser": session.UserID})
	} else {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	}
}
