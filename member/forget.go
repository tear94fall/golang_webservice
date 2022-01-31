package member

import (
	"errors"
	"main/common"
	"main/database"

	"github.com/gin-gonic/gin"
)

func Forget(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")

	if len(id) == 0 {
		err := errors.New("아이디를 입력해주세요")
		common.ErrorPage(c, "입력된 아이디가 없습니다", err)
		return
	}

	if len(name) == 0 {
		err := errors.New("이름을 입력해주세요")
		common.ErrorPage(c, "입력된 이름이 없습니다", err)
		return
	}

	member := &database.Member{}

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	if err := database.GetMemberByUserId(db.DBConn, member, id); err != nil {
		err := errors.New("등록된 사용자가 없습니다")
		common.ErrorPage(c, "아이디를 다시 한번 확인해주세요", err)
		return
	}

	if member.Name != name {
		err := errors.New("등록된 이름이 없습니다")
		common.ErrorPage(c, "이름을 다시 한번 확인해주세요", err)
		return
	}

	common.Render(c, gin.H{
		"title":  "비밀번호 초기화",
		"userid": id,
	}, common.MemberResetHtml)
}
