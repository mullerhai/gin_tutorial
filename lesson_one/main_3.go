package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func  main()  {

	router := gin.Default()
	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.JSON(http.StatusOK,gin.H{
			"message":name,
		})
	})
	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		context.JSON(http.StatusBadRequest,gin.H{
			"name" :name,
			"action": action,
		})
		context.Redirect(http.StatusPermanentRedirect,"www.baidu.com")
	})

	// enctype必须设置： enctype:multipart/form-data ,必须删除 content-type :application/json
	router.POST("/upload", func(context *gin.Context) {
		file,_ := context.FormFile("file")
		fmt.Println(file.Filename)
		context.SaveUploadedFile(file,"/Users/muller/Documents/gowork/src/gin_tutorial/lesson_one/hh.txt")
		context.JSON(http.StatusOK,gin.H{
			"hh":"success",
		})
	})
	// 多个参数重名 都叫 "uplaod[]"
	router.POST("/uploadarr", func(context *gin.Context) {
		form,_ :=context.MultipartForm()
		files :=form.File["upload[]"]
		for _,file := range  files{
			fmt.Println(file.Filename)
		}
		context.JSON(http.StatusOK,gin.H{
			"message ":"uplaod cess",
		})
	})
	router.Run(":8999")
}
