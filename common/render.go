package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorPage(c *gin.Context, titile string, err error) {
	c.HTML(http.StatusOK, NotFoundHtml, gin.H{
		"title":   titile,
		"message": err,
	})
}
