package comment

import (
	"main/common"
	"main/database"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	comment := GetComment(c, id)
	article_id := comment.ArticleId

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	database.DeleteComment(db.DBConn, comment)

	articlePagePath := common.PostArticleHtml
	articlePage := strings.TrimSuffix(articlePagePath, filepath.Ext(articlePagePath))
	articlePageId := "/" + articlePage + "/" + article_id

	c.Redirect(http.StatusMovedPermanently, articlePageId)
}
