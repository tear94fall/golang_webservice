package attach

import (
	"main/database"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context, id string) *[]database.Attach {
	db, _ := c.MustGet("mysql").(*database.DBHandler)
	attaches := &[]database.Attach{}

	if err := database.GetAttachByArticleId(db.DBConn, attaches, id); err != nil {
		return nil
	}

	return attaches
}
