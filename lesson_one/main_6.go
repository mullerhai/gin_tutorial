package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User string  `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
func main()  {

	router := gin.Default()
	//postman Body raw  {"user":"hh","password":"2343"}
	router.POST("/loginJson", func(context *gin.Context) {
		var json Login
		if err := context.ShouldBindJSON(&json);err == nil{
			fmt.Println("user :  password ",json.User,json.Password)
			if json.User =="nanu" &&json.Password=="123"{
				context.JSON(http.StatusOK,gin.H{
					"user":json.User,
					"passs":json.Password,
				})
			}
		}else {
			context.JSON(http.StatusBadRequest,gin.H{
				"user":err.Error(),
			})
		}
	})
	// 一个 HTML 表单绑定的示例 (user=manu&password=123)
	// postman body x-www-form-urlencoded user:mana
	//password:1224
	router.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 这个将通过 content-type 头去推断绑定器使用哪个依赖。

		if err := c.ShouldBind(&form); err == nil {
			fmt.Println("user :  password ",form.User,form.Password)
			if form.User == "manu" && form.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	router.Run(":8998")
}
