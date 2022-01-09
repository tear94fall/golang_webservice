package main

import (
	"fmt"
	"main/database"
	"main/member"
	"main/post"
	"main/render"

	"github.com/gin-gonic/gin"
)

func main() {
	r := SetupRouter()
	r.Run()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	mysql, err := database.NewDBHandler()
	if err != nil {
		fmt.Println("DB init fail")
		panic("mysql init fail")
	}

	r.Use(MysqlContext(mysql))

	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", render.IndexPage)

	// member
	memberGroup := r.Group("/member")
	{
		// page render
		memberGroup.GET("/login", render.MemberLoginPage)
		memberGroup.GET("/register", render.MemberRegisterPage)

		// business logic
		memberGroup.POST("/login", member.Login)
		memberGroup.POST("/register", member.Register)
		memberGroup.POST("/delete", member.Delete)
	}

	// board
	boardGroup := r.Group("/board")
	{
		// page render
		boardGroup.GET("/board", render.BoardPage)

		// business logic
	}

	// post
	postGroup := r.Group("/post")
	{
		// page render
		postGroup.GET("/register", render.PostRegisterPage)
		postGroup.GET("/list", render.PostListPage)
		postGroup.GET("/delete", render.PostDeletePage)

		// business logic
		postGroup.POST("/register", post.Register)
	}

	return r
}

func MysqlContext(handler *database.DBHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if handler == nil {
			panic("init mysql database conn fail")
		}

		c.Set("mysql", handler)
		c.Next()
	}
}
