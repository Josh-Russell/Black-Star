package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToVideoViewPage(c *gin.Context) {
	session := RequestSession(c.Request)
	video, err := globalVideoStore.Find(c.Param("videoID"))
	fmt.Println("--------------------------------------------------")
	fmt.Println(c.Request.Header.Get("Referer"))
	fmt.Println("--------------------------------------------------")

	if video.mature && !strings.Contains(c.Request.Header.Get("Referer"), "confirmAge") {
		c.Redirect(303, "/confirmAge"+c.Param("videoID"))
	} else if err != nil {
		c.String(http.StatusExpectationFailed, "did not find the video.", gin.H{"error": err, "currentuser": session})
	} else if session != nil {
		c.HTML(http.StatusOK, "viewVideo.tmpl.html", gin.H{"video": video, "currentuser": session})
	} else {
		c.HTML(http.StatusOK, "viewVideo.tmpl.html", nil)
	}

}

func HandleNavigateToConfirmAge(c *gin.Context) {
	video, _ := globalVideoStore.Find(c.Param("videoID"))
	session := RequestSession(c.Request)
	c.HTML(http.StatusAccepted, "confirmAge.tmpl.html", gin.H{"video": video, "currentuser": session})
}
