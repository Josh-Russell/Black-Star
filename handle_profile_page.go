package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToProfile(c *gin.Context) {
	session := RequestSession(c.Request)

	if session != nil {
		c.HTML(http.StatusOK, "profile.tmpl.html", gin.H{"currentuser": session.UserID})
	} else {
		c.HTML(http.StatusOK, "profile.tmpl.html", nil)
	}
}
