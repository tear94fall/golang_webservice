package post

import (
	"main/database"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context, id string) *database.Post {
	db, _ := c.MustGet("mysql").(*database.DBHandler)
	article := &database.Post{}

	if err := database.GetPostById(db.DBConn, article, id); err != nil {
		return nil
	}

	return article
}
