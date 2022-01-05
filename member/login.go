package member

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInfo struct {
	Username string
	Password string
}

func IndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "환영합니다",
	})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "환영합니다",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	login := &LoginInfo{username, password}
	var err error

	if len(username) == 0 {
		err = errors.New("아이디를 입력해주세요")
	} else if len(password) == 0 {
		err = errors.New("비밀번호를 입력해 주세요")
	}

	if err == nil && username != login.Username && password != login.Password {
		err = errors.New("아이디가 존재하지 않습니다")
	}

	if err != nil {
		ErrorPage(c, "로그인 실패", err)
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":  "로그인 성공",
		"member": login,
	})
}

func ErrorPage(c *gin.Context, titile string, err error) {
	c.HTML(http.StatusOK, "404.html", gin.H{
		"title":   titile,
		"message": err,
	})
}
