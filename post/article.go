package post

import (
	"main/database"
	"time"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context, id string) *database.Post {
	db, _ := c.MustGet("mysql").(*database.DBHandler)
	article := &database.Post{}

	if err := database.GetPostById(db.DBConn, article, id); err != nil {
		return nil
	}

	layout := "2006-01-02T15:04:05.000Z"
	t, _ := time.Parse(layout, article.RegisterDate)
	article.RegisterDate = t.Format("2006-01-02")

	t2, _ := time.Parse(layout, article.ModifiedDate)
	article.ModifiedDate = t2.Format("2006-01-02")

	article.Views += 1
	database.UpdatePost(db.DBConn, article)

	return article
}
