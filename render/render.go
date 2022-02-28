package render

import (
	"main/comment"
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

func MemberDeletePage(c *gin.Context) {
	common.Render(c, gin.H{
		"title": "탈퇴 하기",
	}, common.MemberDeleteHtml)
}

func MemberForgetPage(c *gin.Context) {
	common.Render(c, gin.H{
		"title": "비밀번호 찾기",
	}, common.MemberForgetHtml)
}

func MemberResetPage(c *gin.Context) {
	common.Render(c, gin.H{
		"title": "비밀번호 초기화",
	}, common.MemberResetHtml)
}

func PostListPage(c *gin.Context) {
	list, page := post.List(c)
	member := member.Find(c)

	common.Render(c, gin.H{
		"title":  "게시글 목록",
		"posts":  list,
		"pages":  page,
		"member": member,
	}, common.PostListHtml)
}

func PostRegisterPage(c *gin.Context) {
	common.Render(c, gin.H{
		"title": "게시글 등록 하기",
	}, common.PostRegisterHtml)
}

func PostModifyPage(c *gin.Context) {
	id := c.Param("id")
	member := member.Find(c)
	article := post.GetArticle(c, id)
	comments := comment.List(c, id)

	common.Render(c, gin.H{
		"title":    "게시글 수정하기",
		"member":   member,
		"article":  article,
		"comments": comments,
	}, common.PostModifyHtml)
}

func PostArticlePage(c *gin.Context) {
	id := c.Param("id")
	member := member.Find(c)
	article := post.GetArticle(c, id)
	comments := comment.List(c, id)

	common.Render(c, gin.H{
		"title":    "게시글 보기",
		"member":   member,
		"article":  article,
		"comments": comments,
	}, common.PostArticleHtml)
}

func CommentModifyPage(c *gin.Context) {
	id := c.Param("id")
	member := member.Find(c)
	comment := comment.GetComment(c, id)
	article := post.GetArticle(c, comment.ArticleId)

	common.Render(c, gin.H{
		"title":   "댓글 수정하기",
		"member":  member,
		"article": article,
		"comment": comment,
	}, common.CommenttModifyHtml)
}
