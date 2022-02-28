package comment

import (
	"main/common"
	"main/database"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentInfo struct {
	ArticleId string
	Content   string
	Writer    string
}

func Register(c *gin.Context) {
	id := c.PostForm("article-id")
	content := c.PostForm("content")
	writer := c.PostForm("writer")

	register := &CommentInfo{id, content, writer}

	comment := NewComment(register)

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	database.CreateComment(db.DBConn, comment)

	postArticle := "/" + strings.TrimSuffix(common.PostArticleHtml, filepath.Ext(common.PostArticleHtml)) + "/" + id
	c.Redirect(http.StatusMovedPermanently, postArticle)
}

func NewComment(register *CommentInfo) *database.Comment {
	comment := &database.Comment{}
	comment.ArticleId = register.ArticleId
	comment.Content = register.Content
	comment.Writer = register.Writer
	comment.RegisterDate = time.Now().Format("2006-01-02")

	return comment
}
