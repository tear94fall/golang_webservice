package post

import (
	"time"
)

type Post struct {
	Title        string
	Content      string
	Writer       string
	RegisterDate string
	ModifiedDate string
	View         int
}

func List() *[]Post {
	now := time.Now()

	// dummky data
	list := &[]Post{
		{"테스트 게시글 입니다.", "테스트 내용", "관리자", now.Format("2006-01-02"), now.Format("2006-01-02 15:04:05"), 1},
		{"공지사항 입니다.", "공지 사항 내용", "admin", now.Format("2006-01-02"), now.Format("2006-01-02 15:04:05"), 92},
		{"안내입니다.", "안내 내용", "root", now.Format("2006-01-02"), now.Format("2006-01-02 15:04:05"), 73},
		{"알려드립니다.", "알려드릴 내용", "manager", now.Format("2006-01-02"), now.Format("2006-01-02 15:04:05"), 3},
	}

	return list
}
