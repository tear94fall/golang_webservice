package member

import (
	"errors"
	"main/common"
	"main/database"
	"main/util"
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

	EncPassword, _ := util.EncStr(password)

	register := &RegisterInfo{id, EncPassword, name, tel}

	err := CheckRegisterMember(register)
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

	c.HTML(http.StatusOK, common.IndexHtml, gin.H{
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

func NewMember(register *RegisterInfo) *database.Member {
	member := &database.Member{}
	member.UserId = register.Id
	member.Password = register.Password
	member.Name = register.Name
	member.Tel = register.Tel

	return member
}
