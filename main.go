package main

import (
	"main/member"
	"main/render"

	"github.com/gin-gonic/gin"
)

func main() {
	r := SetupRouter()
	r.Run()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", render.IndexPage)
	r.GET("/login", render.LoginPage)
	r.GET("/register", render.RegisterPage)
	r.GET("/board", render.BoardPage)

	r.POST("/login", member.Login)
	r.POST("/register", member.Register)

	return r
}
