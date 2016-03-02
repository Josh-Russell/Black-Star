package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}
