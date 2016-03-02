package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToDiscoverPage(c *gin.Context) {
	c.HTML(http.StatusOK, "discover.tmpl.html", nil)
}
