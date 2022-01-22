package post

import (
	"main/common"
	"main/database"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.PostForm("article-id")

	post := GetArticle(c, id)

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	database.DeletePost(db.DBConn, post, id)

	postListPath := filepath.Base(common.PostListHtml)
	postList := strings.TrimSuffix(postListPath, filepath.Ext(postListPath))
	c.Redirect(http.StatusMovedPermanently, postList)
}
