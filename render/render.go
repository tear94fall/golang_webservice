package render

import (
	"main/common"
	"main/post"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, common.IndexHtml, gin.H{
		"title": "Go와 Docker를 이용한 웹 서비스",
	})
}

func MemberLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, common.MemberLoginHtml, gin.H{
		"title": "로그인 하기",
	})
}

func MemberRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, common.MemberRegisterHtml, gin.H{
		"title": "회원 가입 하기",
	})
}

func PostListPage(c *gin.Context) {
	list := post.List()
	c.HTML(http.StatusOK, common.PostListHtml, gin.H{
		"title": "게시글 목록",
		"posts": list,
	})
}

func PostRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, common.PostRegisterHtml, gin.H{
		"title": "게시글 등록 하기",
	})
}

func PostDeletePage(c *gin.Context) {
	c.HTML(http.StatusOK, common.PostDeleteHtml, gin.H{
		"title": "게시글 삭제 하기",
	})
}
