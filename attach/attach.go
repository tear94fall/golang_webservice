package attach

import (
	"main/database"

	"github.com/gin-gonic/gin"
)

func GetAttach(c *gin.Context, id string) *database.Attach {
	db, _ := c.MustGet("mysql").(*database.DBHandler)
	attach := &database.Attach{}

	if err := database.GetAttachById(db.DBConn, attach, id); err != nil {
		return nil
	}

	return attach
}
