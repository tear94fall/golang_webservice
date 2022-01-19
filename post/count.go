package post

import (
	"main/database"

	"github.com/gin-gonic/gin"
)

func Count(c *gin.Context) int64 {
	db, _ := c.MustGet("mysql").(*database.DBHandler)

	var count int64
	if err := database.GetPostsCount(db.DBConn, &count); err != nil {
		return 0
	}

	return count
}
