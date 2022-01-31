package main

import (
	"fmt"
	"main/auth"
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
	r.Use(AuthContext())

	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", render.IndexPage)

	// member
	memberGroup := r.Group("/member")
	{
		// page render
		memberGroup.GET("/login", render.MemberLoginPage)
		memberGroup.GET("/register", render.MemberRegisterPage)
		memberGroup.GET("/modify", render.MemberModifyPage)
		memberGroup.GET("/delete", render.MemberDeletePage)
		memberGroup.GET("/forget", render.MemberForgetPage)
		memberGroup.GET("/reset", render.MemberResetPage)
		memberGroup.GET("/logout", member.Logout)

		// business logic
		memberGroup.POST("/login", member.Login)
		memberGroup.POST("/register", member.Register)
		memberGroup.POST("/modify", member.Modify)
		memberGroup.POST("/delete", member.Delete)
		memberGroup.POST("/forget", member.Forget)
		memberGroup.POST("/reset", member.Reset)
	}

	// post
	postGroup := r.Group("/post")
	{
		// page render
		postGroup.GET("/register", render.PostRegisterPage)
		postGroup.GET("/list/", render.PostListPage)
		postGroup.GET("/list/:index", render.PostListPage)
		postGroup.GET("/article/:id", render.PostArticlePage)

		// business logic
		postGroup.POST("/register", post.Register)
		postGroup.POST("/modify", post.Modify)
		postGroup.POST("/delete", post.Delete)
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

func AuthContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("login_token"); err == nil || token != "" {
			if auth.CheckLoginToken(c, token) != nil {
				c.Set("logon", false)
			} else {
				c.Set("logon", true)
			}
		} else {
			c.Set("logon", false)
		}
	}
}
