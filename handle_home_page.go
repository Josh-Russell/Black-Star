package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToHome(c *gin.Context) {
	session := RequestSession(c.Request)

	videos, err := globalVideoStore.FindAll(0)

	if err != nil {
		c.String(http.StatusExpectationFailed, "Was unable to find the videos.", gin.H{"error": err})
	}
	if session != nil {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"currentuser": session.UserID, "videos": videos})
	} else {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	}
}
func HandleNavigateToHomeWithPageNumber(c *gin.Context) {
	session := RequestSession(c.Request)
	pageNumber, _ := strconv.Atoi(c.Param("pageNumber"))

	videos, err := globalVideoStore.FindAll(pageSize * (pageNumber - 1))

	if err != nil {
		c.String(http.StatusExpectationFailed, "Was unable to find the videos.", gin.H{"error": err})
	}
	if session != nil {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"currentuser": session.UserID, "videos": videos})
	} else {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	}
}
