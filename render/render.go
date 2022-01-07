package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, IndexHtml, gin.H{
		"title": "Go와 Docker를 이용한 웹 서비스",
	})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, LoginHtml, gin.H{
		"title": "로그인 하기",
	})
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, RegisterHtml, gin.H{
		"title": "회원 가입 하기",
	})
}

func BoardPage(c *gin.Context) {
	c.HTML(http.StatusOK, BoardHtml, gin.H{
		"title": "회원 가입 하기",
	})
}

func ErrorPage(c *gin.Context, titile string, err error) {
	c.HTML(http.StatusOK, NotFoundHtml, gin.H{
		"title":   titile,
		"message": err,
	})
}