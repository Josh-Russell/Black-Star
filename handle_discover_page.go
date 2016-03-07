package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToDiscoverPage(c *gin.Context) {
	session := RequestSession(c.Request)

	if session != nil {
		c.HTML(http.StatusOK, "discover.tmpl.html", gin.H{"currentuser": session.UserID})
	} else {
		c.HTML(http.StatusOK, "discover.tmpl.html", nil)
	}
}
