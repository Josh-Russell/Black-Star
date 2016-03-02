package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToProfile(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.tmpl.html", nil)
}
