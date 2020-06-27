package main

import (
	"./bean"
	"./controler"
	"./middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
)

func main()  {
	r := gin.Default()
	r.Use(middleware.CorsHandler())
	h5(r)
	restfulApi(r)
	r.Run("192.168.1.12:8093")
}

func restfulApi(r *gin.Engine)  {
	api := r.Group("/api")
	{
		api.POST(controler.ProtocolUpload, HandleUploadFile)
		api.POST(controler.ProtocolDownload, HandleDownloadFile)
	}
}

func h5(r *gin.Engine)  {
	h5 := r.Group("/")
	{
		h5.Static("static/css","template/dist/static/css")
		h5.Static("static/fonts","template/dist/static/fonts")
		h5.Static("static/js","template/dist/static/js")
		h5.StaticFile("favicon.ico","template/dist/favicon.ico")
		h5.StaticFile("score.zip","score.zip")
		r.LoadHTMLFiles("template/dist/index.html")
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
}

func HandleUploadFile(c *gin.Context){
	file, header ,err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK ,gin.H{
			"code":400,
			"msg":"文件上传失败"})
		return
	}
	defer file.Close()

	filename := header.Filename
	fileSuffix := path.Ext(filename)
	fmt.Println(filename, fileSuffix)
	if fileSuffix == ".xml" {
		filename = "music.xml"
	} else if fileSuffix == ".mid" {
		filename = "score.mid"
	}

	out,err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()
	_,err = io.Copy(out,file)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"data":gin.H{
			"resId":"",
			"resUrl":filename,
		},
		"msg":"上传文件成功"})

}

func HandleDownloadFile(c* gin.Context){
	//TODO shell
	var b bean.ProtocolDownLoadBean
	c.BindJSON(&b)
	fmt.Println(b.MidiType)
	var command string = ""
	if b.MidiType=="1" {//钢琴
		command = `./p_run.sh`
	} else if b.MidiType== "2" { // 小提琴
		command = `./v_run.sh`
	} else if b.MidiType == "3"{ // 架子鼓
		command = `./d_run.sh`
	}
	cmd := exec.Command("/bin/bash","-c",command)
	out,err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
		c.JSON(http.StatusOK,gin.H{
			"code":400,
			"msg":"打包失败"})

	} else {
		fmt.Println(out)

		c.JSON(http.StatusOK,gin.H{
			"code":200,
			"data":gin.H{
				"resId":b.MidiType,
				"resUrl":"./score.zip",
			},
			"msg":""})
	}



}