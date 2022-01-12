package post

import (
	"main/database"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) *[]database.Post {
	db, _ := c.MustGet("mysql").(*database.DBHandler)
	list := &[]database.Post{}

	if err := database.GetPosts(db.DBConn, list); err != nil {
		return nil
	}

	return list
}
