package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"

	//"log"
	"net/http"
)



func errorParam(c *gin.Context){
	c.JSON(http.StatusBadRequest,gin.H{
		"code":400,
		"msg":"参数错误",
	})
}

//func HandleUploadFile(c *gin.Context){
//	file, header ,err := c.Request.FormFile("file")
//	if err != nil {
//		c.JSON(http.StatusBadRequest ,gin.H{
//			"code":400,
//			"msg":"文件上传失败"})
//		return
//	}
//
//	content, err := ioutil.ReadAll(file)
//	if err != nil {
//		c.JSON(http.StatusBadRequest ,gin.H{
//			"code":400,
//			"msg":"文件读取失败"})
//		return
//	}
//
//	fmt.Println(header.Filename)
//	fmt.Println(string(content))
//	c.JSON(http.StatusOK,gin.H{
//		"code":200,
//		"msg":"上传文件成功"})
//}
//
//func HandleDownloadFile(c* gin.Context){
//	//TODO shell
//
//	content := c.Query("content")
//	content = "hello world, 我是一个文件，" + content
//
//	c.Writer.WriteHeader(http.StatusOK)
//	c.Header("Content-Disposition", "attachment; filename=hello.txt")
//	c.Header("Content-Type", "application/text/plain")
//	c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
//	c.Writer.Write([]byte(content))
//}