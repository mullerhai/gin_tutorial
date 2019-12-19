package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main()  {
	gin.DisableConsoleColor()
	f,_ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	authorized :=router.Group("/")
	authorized.Use(Auth())
	{
		authorized.POST("/login",loginEnpoint)
		authorized.GET("/submit",submitEnpoint)
	}
	gin.SetMode(gin.ReleaseMode)
	router.Run(":8999")
}
func  loginEnpoint(ctx *gin.Context)  {
	fmt.Println("login Enpoint")
	ctx.JSON(http.StatusOK,gin.H{
		"login" :"fail",

	})
}

func  submitEnpoint(ctx *gin.Context)  {
	fmt.Println("submitEnpoint")
	ctx.JSON(http.StatusOK,gin.H{
		"submit" :"fail",

	})
}
func  Auth()gin.HandlerFunc  {
	fmt.Println("auth Enpoint")
	return func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"mess":"auth",
		})
	}
}
