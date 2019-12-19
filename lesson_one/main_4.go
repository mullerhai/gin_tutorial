package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()
	v1 := router.Group("v1")
	{
		v1.POST("/login", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"message_v1":"vi login",
			})
		})
		v1.GET("/submit", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"message_v1" :"v1 submit",
			})
		})
	}


	v2 := router.Group("v2")
	{
		v2.POST("/login", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"message_v2":"v2 login",
			})
		})
		v2.GET("/submit", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"message_v2" :"v2 submit",
			})
		})
	}
	router.Run(":8999")
}
