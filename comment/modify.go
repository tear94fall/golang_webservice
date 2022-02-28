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

func Modify(c *gin.Context) {
	id := c.PostForm("comment-id")
	content := c.PostForm("content")
	writer := c.PostForm("writer")

	db, _ := c.MustGet("mysql").(*database.DBHandler)

	article_id, err := database.GetArticleIdByCommentId(db.DBConn, id)
	if err != nil {
		common.ErrorPage(c, "댓글 수정 실패", err)
		return
	}

	modified := &CommentInfo{article_id, content, writer}
	register := &database.Comment{}

	database.GetCommentById(db.DBConn, register, id)

	ModifiedComment(register, modified)
	database.UpdateComment(db.DBConn, register)

	articlePagePath := common.PostArticleHtml
	articlePage := strings.TrimSuffix(articlePagePath, filepath.Ext(articlePagePath))
	articlePageId := "/" + articlePage + "/" + article_id

	c.Redirect(http.StatusMovedPermanently, articlePageId)
}

func ModifiedComment(old *database.Comment, new *CommentInfo) {
	var update bool = false

	if old.Content != new.Content && len(new.Content) != 0 {
		old.Content = new.Content
		update = true
	}

	if old.Writer != new.Writer && len(new.Writer) != 0 {
		old.Writer = new.Writer
	}

	if update {
		old.RegisterDate = time.Now().Format("2006-01-02")
		return
	}

	layout := "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(layout, old.RegisterDate)
	old.RegisterDate = t.Format("2006-01-02")
}
