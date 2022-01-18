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

func Render(c *gin.Context, data gin.H, PageHtml string) {
	login, _ := c.Get("logon")
	data["login"] = login.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data)
	case "application/xml":
		c.XML(http.StatusOK, data)
	default:
		c.HTML(http.StatusOK, PageHtml, data)
	}
}
