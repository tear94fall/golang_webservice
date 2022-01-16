package render

import (
	"main/common"
	"main/member"
	"main/post"

	"github.com/gin-gonic/gin"
)

func IndexPage(c *gin.Context) {
	common.Render(c, gin.H{
		"title": "Go와 Docker를 이용한 웹 서비스",
	}, common.IndexHtml)
}

func MemberLoginPage(c *gin.Context) {
	common.Render(c, gin.H{
		"title": "로그인 하기",
	}, common.MemberLoginHtml)
}

func MemberRegisterPage(c *gin.Context) {
	common.Render(c, gin.H{
		"title": "회원 가입 하기",
	}, common.MemberRegisterHtml)
}

func MemberModifyPage(c *gin.Context) {
	member := member.Find(c)

	common.Render(c, gin.H{
		"title":  "회원 정보 변경 하기",
		"member": member,
	}, common.MemberModifyHtml)
}

func PostListPage(c *gin.Context) {
	list := post.List(c)

	common.Render(c, gin.H{
		"title": "게시글 목록",
		"posts": list,
	}, common.PostListHtml)
}

func PostRegisterPage(c *gin.Context) {
	common.Render(c, gin.H{
		"title": "게시글 등록 하기",
	}, common.PostRegisterHtml)
}

func PostArticlePage(c *gin.Context) {
	id := c.Param("id")
	article := post.GetArticle(c, id)

	common.Render(c, gin.H{
		"title":   "게시글 보기",
		"article": article,
	}, common.PostArticlePage)
}
