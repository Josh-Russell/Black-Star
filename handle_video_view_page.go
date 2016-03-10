package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToVideoViewPage(c *gin.Context) {
	session := RequestSession(c.Request)
	video, err := globalVideoStore.Find(c.Param("videoID"))

	if err != nil {
		c.String(http.StatusExpectationFailed, "did not find the video.", gin.H{"error": err})
	}
	if session != nil {
		c.HTML(http.StatusOK, "viewVideo.tmpl.html", gin.H{"video": video})
	} else {
		c.HTML(http.StatusOK, "viewVideo.tmpl.html", nil)
	}

}
