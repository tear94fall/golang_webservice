package attach

import (
	"main/common"
	"main/database"
	"main/util"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	attach := GetAttach(c, id)
	file_name := attach.FileName
	article_id := attach.ArticleId

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	database.DeleteAttach(db.DBConn, attach)

	// Delete local file
	file_dir, _ := util.EncStr(article_id)
	file_path := common.AttachPath + "/" + file_dir
	full_file_path := file_path + "/" + file_name
	util.RemoveFile(full_file_path)

	articlePagePath := common.PostArticleHtml
	articlePage := strings.TrimSuffix(articlePagePath, filepath.Ext(articlePagePath))
	articlePageId := "/" + articlePage + "/" + article_id

	c.Redirect(http.StatusMovedPermanently, articlePageId)
}
