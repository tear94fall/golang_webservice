package comment

import (
	"main/database"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context, id string) *[]database.Comment {
	db, _ := c.MustGet("mysql").(*database.DBHandler)
	comments := &[]database.Comment{}

	if err := database.GetCommentByArticleId(db.DBConn, comments, id); err != nil {
		return nil
	}

	return comments
}
