package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToVideoViewPage(c *gin.Context) {
	c.HTML(http.StatusOK, "viewVideo.tmpl.html", nil)
}
