package member

import (
	"errors"
	"fmt"
	"main/auth"
	"main/common"
	"main/database"
	"main/util"

	"github.com/gin-gonic/gin"
)

type DeleteInfo struct {
	Id       string
	Password string
}

func Delete(c *gin.Context) {
	password1 := c.PostForm("password1")
	password2 := c.PostForm("password2")

	if password1 != password2 {
		err := errors.New("비밀번호를 불일치")
		common.ErrorPage(c, "비밀번호 두개가 일치하지 않습니다", err)
		return
	}

	if len(password1) == 0 || len(password2) == 0 {
		err := errors.New("비밀번호를 입력해주세요")
		common.ErrorPage(c, "입력된 비밀번호가 없습니다", err)
		return
	}

	member := Find(c)

	EncPassword1, _ := util.EncStr(password1)

	fmt.Println(member.Password)
	fmt.Println(EncPassword1)

	if member.Password != EncPassword1 {
		err := errors.New("비밀번호 오류")
		common.ErrorPage(c, "비밀번호가 일치하지 않습니다", err)
		return
	}

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	if err := database.DeleteMember(db.DBConn, member, member.UserId); err != nil {
		fmt.Println(err)
		err := errors.New("존재하지 않는 사용자")
		common.ErrorPage(c, "회원 탈퇴에 실패하였습니다", err)
		return
	}

	token := auth.GenToken(member)
	database.DeleteToken(db.DBConn, token, member.UserId)

	Logout(c)

	common.Render(c, gin.H{
		"title": "회원 탈퇴 성공",
	}, common.IndexHtml)
}

func CheckDeleteMember(member *DeleteInfo) error {
	var err error

	if len(member.Id) == 0 {
		err = errors.New("아이디를 입력해주세요")
	} else if len(member.Password) == 0 {
		err = errors.New("비밀번호를 입력해 주세요")
	}

	return err
}
