package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToUpload(c *gin.Context) {
	session := RequestSession(c.Request)

	if session != nil {
		c.HTML(http.StatusOK, "upload.tmpl.html", gin.H{"currentuser": session.UserID})
	} else {
		c.HTML(http.StatusOK, "login.tmpl.html", nil)
	}
}
