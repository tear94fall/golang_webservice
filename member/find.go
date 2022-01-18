package member

import (
	"main/database"

	"github.com/gin-gonic/gin"
)

func Find(c *gin.Context) *database.Member {
	db, _ := c.MustGet("mysql").(*database.DBHandler)

	token, err := c.Request.Cookie("login_token")
	if err != nil {
		return nil
	}

	t := &database.Token{}
	tokenErr := database.GetIdByToken(db.DBConn, t, token.Value)
	if tokenErr != nil {
		return nil
	}

	member := &database.Member{}
	memberErr := database.GetMemberByUserId(db.DBConn, member, t.UserId)
	if memberErr != nil {
		return nil
	}

	return member
}
