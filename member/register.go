package member

import (
	"errors"
	"main/render"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInfo struct {
	Id       string
	Password string
	Name     string
	Tel      string
}

func Register(c *gin.Context) {
	id := c.PostForm("id")
	password := c.PostForm("password")
	name := c.PostForm("name")
	tel := c.PostForm("tel")

	register := &RegisterInfo{id, password, name, tel}

	err := CheckRegisterMember(register)
	if err != nil {
		render.ErrorPage(c, "회원가입 실패", err)
		return
	}

	c.HTML(http.StatusOK, render.IndexHtml, gin.H{
		"title":  "회원가입 성공",
		"member": register,
	})
}

func CheckRegisterMember(member *RegisterInfo) error {
	var err error

	if len(member.Id) == 0 {
		err = errors.New("아이디를 입력해주세요")
	} else if len(member.Password) == 0 {
		err = errors.New("비밀번호를 입력해 주세요")
	} else if len(member.Name) == 0 {
		err = errors.New("이름을 입력해 주세요")
	} else if len(member.Tel) == 0 {
		err = errors.New("전화번호를 입력해 주세요")
	}

	return err
}
