package render

import (
	"main/common"
	"main/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	Render(c, gin.H{
		"title": "Go와 Docker를 이용한 웹 서비스",
	}, common.IndexHtml)
}

func MemberLoginPage(c *gin.Context) {
	Render(c, gin.H{
		"title": "로그인 하기",
	}, common.MemberLoginHtml)
}

func MemberRegisterPage(c *gin.Context) {
	Render(c, gin.H{
		"title": "회원 가입 하기",
	}, common.MemberRegisterHtml)
}

func PostListPage(c *gin.Context) {
	list := post.List(c)

	Render(c, gin.H{
		"title": "게시글 목록",
		"posts": list,
	}, common.PostListHtml)
}

func PostRegisterPage(c *gin.Context) {
	Render(c, gin.H{
		"title": "게시글 등록 하기",
	}, common.PostRegisterHtml)
}

func PostArticlePage(c *gin.Context) {
	id := c.Param("id")
	article := post.GetArticle(c, id)

	Render(c, gin.H{
		"title":   "게시글 보기",
		"article": article,
	}, common.PostArticlePage)
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
