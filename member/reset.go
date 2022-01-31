package member

import (
	"errors"
	"main/auth"
	"main/common"
	"main/database"
	"main/util"

	"github.com/gin-gonic/gin"
)

func Reset(c *gin.Context) {
	id := c.PostForm("id")
	password1 := c.PostForm("password1")
	password2 := c.PostForm("password2")

	if len(password1) == 0 || len(password2) == 0 {
		err := errors.New("비밀번호를 입력해주세요")
		common.ErrorPage(c, "패스워드가 전부 입력되지 않았습니다", err)
		return
	}

	if password1 != password2 {
		err := errors.New("비밀번호를 일치하지 않습니다")
		common.ErrorPage(c, "동일한 패스워드를 입력해주세요", err)
		return
	}

	EncPassword, _ := util.EncStr(password1)

	member := &database.Member{}

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	if err := database.GetMemberByUserId(db.DBConn, member, id); err != nil {
		err := errors.New("등록된 사용자가 없습니다")
		common.ErrorPage(c, "아이디를 다시 한번 확인해주세요", err)
		return
	}

	if member.Password == EncPassword {
		err := errors.New("동일한 패스워드로 변경이 불가능 합니다")
		common.ErrorPage(c, "변경할 패스워드를 다시 입력해주세요", err)
		return
	}

	member.Password = EncPassword
	database.UpdateMember(db.DBConn, member)

	NewToken := auth.GenToken(member)
	database.UpdateToken(db.DBConn, NewToken, member.UserId)

	c.SetCookie("login_token", NewToken, 3600, "", "", false, true)
	c.Set("login", true)

	common.Render(c, gin.H{
		"title":  "비밀번호 변경 성공",
		"member": member,
	}, common.MemberResetSuccessHtml)
}
