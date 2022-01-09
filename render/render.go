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

func MemberLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, MemberLoginHtml, gin.H{
		"title": "로그인 하기",
	})
}

func MemberRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, MemberRegisterHtml, gin.H{
		"title": "회원 가입 하기",
	})
}

func BoardPage(c *gin.Context) {
	c.HTML(http.StatusOK, BoardHtml, gin.H{
		"title": "회원 가입 하기",
	})
}

func PostListPage(c *gin.Context) {
	c.HTML(http.StatusOK, PostListHtml, gin.H{
		"title": "게시글 목록",
	})
}

func PostRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, PostRegisterHtml, gin.H{
		"title": "게시글 등록 하기",
	})
}

func PostDeletePage(c *gin.Context) {
	c.HTML(http.StatusOK, PostDeleteHtml, gin.H{
		"title": "게시글 삭제 하기",
	})
}

func ErrorPage(c *gin.Context, titile string, err error) {
	c.HTML(http.StatusOK, NotFoundHtml, gin.H{
		"title":   titile,
		"message": err,
	})
}
