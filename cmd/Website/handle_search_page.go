package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleNavigateToSearch(c *gin.Context) {
	c.HTML(http.StatusOK, "search.tmpl.html", nil)
}
