package member

import (
	"errors"
	"main/common"
	"main/database"
	"main/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Id       string
	Password string
}

func Login(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")

	EncPassword, _ := util.EncStr(password)

	login := &LoginInfo{id, password}

	err := CheckLoginMember(login)
	if err != nil {
		common.ErrorPage(c, "로그인 실패", err)
		return
	}

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	member := &database.Member{}
	if err2 := database.GetMemberByUserId(db.DBConn, member, id); err2 != nil {
		common.ErrorPage(c, "로그인 실패", err)
		return
	}

	if EncPassword != member.Password {
		common.ErrorPage(c, "가입되지 않은 사용자 입니다.", errors.New("invalid user id"))
		return
	}

	c.HTML(http.StatusOK, common.IndexHtml, gin.H{
		"title":  "로그인 성공",
		"member": login,
	})
}

func CheckLoginMember(member *LoginInfo) error {
	var err error

	if len(member.Id) == 0 {
		err = errors.New("아이디를 입력해주세요")
	} else if len(member.Password) == 0 {
		err = errors.New("비밀번호를 입력해 주세요")
	}

	return err
}
