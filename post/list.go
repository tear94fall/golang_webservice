package post

import (
	"main/database"
	"strings"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) *[]database.Post {
	db, _ := c.MustGet("mysql").(*database.DBHandler)
	list := &[]database.Post{}

	if err := database.GetPosts(db.DBConn, list); err != nil {
		return nil
	}

	for i := 0; i < len(*list); i++ {
		slice := strings.Split((*list)[i].RegisterDate, "T")
		(*list)[i].RegisterDate = slice[0]
	}

	return list
}
