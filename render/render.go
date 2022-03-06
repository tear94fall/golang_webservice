package render

import (
	"encoding/base64"
	"io/ioutil"
	"main/attach"
	"main/comment"
	"main/common"
	"main/member"
	"main/post"
	"main/util"

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
	member := member.Find(c)

	common.Render(c, gin.H{
		"title":  "게시글 등록 하기",
		"member": member,
	}, common.PostRegisterHtml)
}

func PostModifyPage(c *gin.Context) {
	id := c.Param("id")
	member := member.Find(c)
	article := post.GetArticle(c, id)
	comments := comment.List(c, id)
	attaches := attach.List(c, id)

	common.Render(c, gin.H{
		"title":    "게시글 수정하기",
		"member":   member,
		"article":  article,
		"comments": comments,
		"attaches": attaches,
	}, common.PostModifyHtml)
}

func PostArticlePage(c *gin.Context) {
	id := c.Param("id")
	member := member.Find(c)
	article := post.GetArticle(c, id)
	comments := comment.List(c, id)
	attaches := attach.List(c, id)

	common.Render(c, gin.H{
		"title":    "게시글 보기",
		"member":   member,
		"article":  article,
		"comments": comments,
		"attaches": attaches,
	}, common.PostArticleHtml)
}

func PostArticleAttachPage(c *gin.Context) {
	id := c.Param("id")
	filename := c.Param("file")
	file_dir, _ := util.EncStr(id)
	full_file_path := common.AttachPath + "/" + file_dir + "/" + filename

	var images []string
	bytes, _ := ioutil.ReadFile(full_file_path)
	img := base64.StdEncoding.EncodeToString(bytes)
	images = append(images, img)

	common.Render(c, gin.H{
		"title":  "첨부파일 보기",
		"images": images,
	}, common.PostArticleAttachHtml)
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
