package post

import (
	"errors"
	"main/attach"
	"main/common"
	"main/database"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type PostInfo struct {
	Title   string
	Content string
	Writer  string
}

func Register(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	writer := c.PostForm("writer")

	register := &PostInfo{title, content, writer}

	err := CheckRegisterPost(register)
	if err != nil {
		common.ErrorPage(c, "글등록 실패", err)
		return
	}

	post := NewPost(register)

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	database.CreatePost(db.DBConn, post)

	attach.SaveAttachFile(c, strconv.Itoa(post.Id))

	postListPath := filepath.Base(common.PostListHtml)
	postList := strings.TrimSuffix(postListPath, filepath.Ext(postListPath))
	c.Redirect(http.StatusMovedPermanently, postList)
}

func CheckRegisterPost(post *PostInfo) error {
	var err error

	if len(post.Title) == 0 {
		err = errors.New("게시글의 제목을 입력해주세요")
	} else if len(post.Content) == 0 {
		err = errors.New("게시글의 내용을 입력해 주세요")
	} else if len(post.Writer) == 0 {
		err = errors.New("작성자를 입력해 주세요")
	}

	return err
}

func NewPost(register *PostInfo) *database.Post {
	post := &database.Post{}
	post.Title = register.Title
	post.Content = register.Content
	post.Writer = register.Writer
	post.RegisterDate = time.Now().Format("2006-01-02")
	post.ModifiedDate = time.Now().Format("2006-01-02")
	post.Views = 0

	return post
}
