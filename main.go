package main

import (
	"main/member"

	"github.com/gin-gonic/gin"
)

func main() {
	r := SetupRouter()
	r.Run()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", member.IndexPage)
	r.GET("/login", member.LoginPage)
	r.POST("/login", member.Login)

	return r
}
