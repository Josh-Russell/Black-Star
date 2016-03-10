package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToHome(c *gin.Context) {
	session := RequestSession(c.Request)

	videos, err := globalVideoStore.FindAll(0)
	if err != nil {
		c.String(http.StatusExpectationFailed, "YOU DUn fuuuucked up.", gin.H{"error": err})
	}
	if session != nil {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"currentuser": session.UserID, "videos": videos})
	} else {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	}
}
