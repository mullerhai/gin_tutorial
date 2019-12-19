package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	server := gin.Default()
	server.GET("/index", open_index)
	server.POST("/for_data",post_data)
	server.Run(":8999")
}

func open_index(context *gin.Context)  {
	context.JSON(http.StatusOK,gin.H{
		"message" :"muller",
	})

}

func post_data(c *gin.Context)  {
	id := c.Query("id")

	fmt.Println("id : " ,id )
	name :=c.PostForm("name")
	age := c.PostForm("age")
	gender :=c.DefaultPostForm("gender","male")
	usb := c.DefaultQuery("user","book")
	fmt.Println("name :  age  : " ,name,age,gender,usb )
	c.JSON(http.StatusBadRequest,gin.H{
		"id" :id,
		"name" : name,
		"age" :age,

	})
}