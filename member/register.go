package member

import (
	"main/auth"
	"main/common"
	"main/database"
	"main/util"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")
	name := c.PostForm("name")
	tel := c.PostForm("tel")

	EncPassword, _ := util.EncStr(password)

	register := &MemberInfo{id, EncPassword, name, tel}

	err := CheckMemberValue(register)
	if err != nil {
		common.ErrorPage(c, "회원가입 실패", err)
		return
	}

	member := NewMember(register)

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	database.CreateMember(db.DBConn, member)

	if id != member.UserId {
		common.ErrorPage(c, "회원가입 실패", err)
		return
	}

	token := auth.GenToken(member)
	c.SetCookie("login_token", token, 3600, "", "", false, true)
	c.Set("login", true)

	database.CreateToken(db.DBConn, token, member.UserId)

	common.Render(c, gin.H{
		"title":  "회원가입 성공",
		"member": register,
	}, common.IndexHtml)
}
