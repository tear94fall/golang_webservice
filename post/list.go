package post

import (
	"main/database"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) *[]database.Post {
	db, _ := c.MustGet("mysql").(*database.DBHandler)
	list := &[]database.Post{}

	const max int64 = 10
	var count int64

	if err := database.GetPostsCount(db.DBConn, &count); err != nil {
		return nil
	}

	start := strconv.FormatInt(count, 10)
	end := strconv.FormatInt((count - max), 10)

	if err := database.GetPostRangeById(db.DBConn, list, end, start); err != nil {
		return nil
	}

	for i := 0; i < len(*list); i++ {
		slice := strings.Split((*list)[i].RegisterDate, "T")
		(*list)[i].RegisterDate = slice[0]
	}

	return list
}
