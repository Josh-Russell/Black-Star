package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToUpload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.tmpl.html", nil)
}
