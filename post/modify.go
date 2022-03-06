package post

import (
	"main/attach"
	"main/common"
	"main/database"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Modify(c *gin.Context) {
	id := c.PostForm("article-id")
	title := c.PostForm("title")
	content := c.PostForm("content")
	writer := c.PostForm("writer")

	modified := &PostInfo{title, content, writer}
	register := &database.Post{}

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	database.GetPostById(db.DBConn, register, id)

	ModifiedPost(register, modified)
	database.UpdatePost(db.DBConn, register)

	attach.SaveAttachFile(c, id)

	articlePagePath := common.PostArticleHtml
	articlePage := strings.TrimSuffix(articlePagePath, filepath.Ext(articlePagePath))
	articlePageId := "/" + articlePage + "/" + id

	c.Redirect(http.StatusMovedPermanently, articlePageId)
}

func ModifiedPost(old *database.Post, new *PostInfo) {
	var update bool = false

	if old.Title != new.Title && len(new.Title) != 0 {
		old.Title = new.Title
		update = true
	}
	if old.Content != new.Content && len(new.Content) != 0 {
		old.Content = new.Content
		update = true
	}
	if old.Writer != new.Writer && len(new.Writer) != 0 {
		old.Writer = new.Writer
		update = true
	}

	layout := "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(layout, old.RegisterDate)
	old.RegisterDate = t.Format("2006-01-02")

	if update {
		old.ModifiedDate = time.Now().Format("2006-01-02")
	}
}
