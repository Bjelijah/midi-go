package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"main_web/bean"
	"main_web/controler"
	"main_web/middleware"
	"net/http"
	"os/exec"
	"path"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CorsHandler())
	h5(r)
	restfulApi(r)
	r.Run(":80")
}

func restfulApi(r *gin.Engine) {
	api := r.Group("/auto_deploy_api")
	{
		api.POST(controler.ProtocolUpload, HandleUploadFile)
		api.POST(controler.ProtocolDownload, HandleDownloadFile)
		api.POST(controler.ProtocolAutoDeploy, HandleAutoDeploy)
	}
}

func h5(r *gin.Engine) {
	h5 := r.Group("/auto_deploy")
	{
		h5.StaticFS("static", http.Dir("template/dist/static"))

		//h5.Static("static/css", "template/dist/static/css")
		//h5.Static("static/fonts", "template/dist/static/fonts")
		//h5.Static("static/js", "template/dist/static/js")
		h5.StaticFile("favicon.ico", "template/dist/favicon.ico")
		h5.StaticFile("index.html", "template/dist/index.html")
		//h5.StaticFile("score.zip", "score.zip")
		h5.StaticFile("index_src.xml", "index_src.xml")

	}
	r.LoadHTMLFiles("template/dist/index.html")
	r.GET("/auto_deploy", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
}

func HandleUploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "文件上传失败"})
		return
	}
	defer file.Close()

	filename := header.Filename
	fileSuffix := path.Ext(filename)
	fmt.Println(filename, fileSuffix)
	if fileSuffix != ".zip" {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "文件类型不是zip"})
		return
	}

	buf := new(bytes.Buffer)
	filesSize, err := io.Copy(buf, file)
	if err != nil {
		panic(err)
	}
	zr, err := zip.NewReader(bytes.NewReader(buf.Bytes()), filesSize)
	zipResBean := controler.UnzipFile(zr)
	if zipResBean.DirIndex > 0 {
		_fileSuffix := path.Ext(filename)
		name := filename[0 : len(filename)-len(_fileSuffix)]
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"resId":  fmt.Sprintf("%d", zipResBean.DirIndex),
				"resUrl": name,
			},
			"message": "上传文件成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "文件解压失败"})
	}

}

func HandleDownloadFile(c *gin.Context) {
	//TODO shell
	var b bean.ProtocolDownLoadBean
	c.BindJSON(&b)
	fmt.Println(b.MidiType)
	var command string = ""
	if b.MidiType == "1" { //钢琴
		command = `./p_run.sh`
	} else if b.MidiType == "2" { // 小提琴
		command = `./v_run.sh`
	} else if b.MidiType == "3" { // 架子鼓
		command = `./d_run.sh`
	}
	cmd := exec.Command("/bin/bash", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		//log.Fatal(err)
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "打包失败"})

	} else {
		fmt.Println(out)

		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": gin.H{
				"resId":  b.MidiType,
				"resUrl": "./score.zip",
			},
			"msg": ""})
	}

}

func HandleAutoDeploy(c *gin.Context) {
	var b bean.AutoDeployBean
	c.BindJSON(&b)
	fmt.Println(b.Content)
	bean := controler.ParseAutoDeploy(b)

	if !bean.Success {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "部署失败",
			"data": gin.H{
				"res": "success",
			},
			"msg": ""})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"res": bean.ResUrl,
		},
		"msg": ""})
}
