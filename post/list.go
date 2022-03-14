package post

import (
	"main/database"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) (*[]database.Post, *PageInfo) {
	current := c.Param("index")

	if len(current) == 0 || current == "" {
		current = "1"
	}

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	list := &[]database.Post{}

	pageInfo := NewPageInfo()

	pageInfo.Current, _ = strconv.Atoi(current)

	var count int64
	if err := database.GetPostsCount(db.DBConn, &count); err != nil {
		return nil, nil
	}

	pageInfo.Total = count
	SetPageInfo(pageInfo)

	if err := database.GetPostPaged(db.DBConn, list, pageInfo.Current, column_max); err != nil {
		return nil, nil
	}

	for i := 0; i < len(*list); i++ {
		slice := strings.Split((*list)[i].RegisterDate, "T")
		(*list)[i].RegisterDate = slice[0]
	}

	return list, pageInfo
}
