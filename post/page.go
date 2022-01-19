package post

import (
	"main/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Page(c *gin.Context) []int64 {
	count := Count(c)

	var pages []int64

	for i := 1; int64(i) <= count; i++ {
		pages = append(pages, int64(i))
	}

	return pages
}

func RangePage(c *gin.Context, index int) error {
	list := &[]database.Post{}

	start := (index - 1) * 10
	end := index * 10

	db, _ := c.MustGet("mysql").(*database.DBHandler)

	if err := database.GetPostRangeById(db.DBConn, list, strconv.Itoa(start), strconv.Itoa(end)); err != nil {
		return err
	}

	return nil
}
