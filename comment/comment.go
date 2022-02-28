package comment

import (
	"main/database"

	"github.com/gin-gonic/gin"
)

func GetComment(c *gin.Context, id string) *database.Comment {
	db, _ := c.MustGet("mysql").(*database.DBHandler)
	comment := &database.Comment{}

	if err := database.GetCommentById(db.DBConn, comment, id); err != nil {
		return nil
	}

	return comment
}
