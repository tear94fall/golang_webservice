package auth

import (
	"main/database"
	"main/util"

	"github.com/gin-gonic/gin"
)

func CheckLoginToken(c *gin.Context, token string) error {
	db, _ := c.MustGet("mysql").(*database.DBHandler)
	if count, err := database.CheckToken(db.DBConn, token); err != nil || count == 0 {
		return err
	}

	return nil
}

func GetIdByToken(c *gin.Context, token string) string {
	db, _ := c.MustGet("mysql").(*database.DBHandler)

	t := &database.Token{}
	if err := database.GetIdByToken(db.DBConn, t, token); err != nil {
		return ""
	}

	return t.UserId
}

func GenToken(member *database.Member) string {
	id, _ := util.EncStr(member.UserId)
	passwd := member.Password

	token, _ := util.EncStr(id + passwd)

	return token
}
