package member

import (
	"errors"
	"main/render"
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

	login := &LoginInfo{id, password}

	err := CheckLoginMember(login)
	if err != nil {
		render.ErrorPage(c, "로그인 실패", err)
		return
	}

	c.HTML(http.StatusOK, render.IndexHtml, gin.H{
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
