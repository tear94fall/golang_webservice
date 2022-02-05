package post

import (
	"main/common"
	"main/database"
	"main/member"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	id := c.Param("id")

	post := GetArticle(c, id)

	db, _ := c.MustGet("mysql").(*database.DBHandler)
	database.DeletePost(db.DBConn, post, id)

	list, page := List(c)
	member := member.Find(c)

	common.Render(c, gin.H{
		"title":  "게시글 목록",
		"posts":  list,
		"pages":  page,
		"member": member,
	}, common.PostListHtml)
}
