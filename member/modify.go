package member

import (
	"errors"
	"main/auth"
	"main/common"
	"main/database"
	"main/util"

	"github.com/gin-gonic/gin"
)

func Modify(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")
	name := c.PostForm("name")
	tel := c.PostForm("tel")

	if len(password) == 0 {
		err := errors.New("비밀번호를 입력해주세요")
		common.ErrorPage(c, "변경 내용이 없습니다", err)
		return
	}

	EncPassword, _ := util.EncStr(password)

	modify := &MemberInfo{id, EncPassword, name, tel}

	err := CheckMemberValue(modify)
	if err != nil {
		common.ErrorPage(c, "회원정보 변경 실패", err)
		return
	}

	member := Find(c)

	if !CheckChangeMember(member, modify) {
		common.ErrorPage(c, "변경 내용이 없습니다", err)
		return
	}

	UpdateMember(member, modify)

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	database.UpdateMember(db.DBConn, member)

	token := auth.GenToken(member)
	c.SetCookie("login_token", token, 3600, "", "", false, true)
	c.Set("login", true)

	database.CreateToken(db.DBConn, token, member.UserId)

	common.Render(c, gin.H{
		"title":  "회원가입 성공",
		"member": modify,
	}, common.IndexHtml)
}
