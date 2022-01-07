package main

import (
	"fmt"
	"main/database"
	"main/member"
	"main/render"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewDBHandler()
	if err != nil {
		fmt.Println("DB init fail")
		return
	}

	r := SetupRouter(db)
	r.Run()
}

func SetupRouter(db_handler *database.DBHandler) *gin.Engine {
	r := gin.Default()
	r.Use(MysqlContext(db_handler))

	r.LoadHTMLGlob("templates/*")

	r.GET("/", render.IndexPage)
	r.GET("/login", render.LoginPage)
	r.GET("/register", render.RegisterPage)
	r.GET("/board", render.BoardPage)

	r.POST("/login", member.Login)
	r.POST("/register", member.Register)

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
