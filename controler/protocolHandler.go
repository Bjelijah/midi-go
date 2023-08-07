package controler

import (
	"archive/zip"
	"bufio"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"main_web/bean"
	"os"
	path2 "path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	ProtocolUpload     = "tools/file/upload"
	ProtocolDownload   = "tools/file/download"
	ProtocolAutoDeploy = "tools/autoDeploy"
	FileSystemRoot     = "/data/edith/activity/Hot_sweat_season"
	//FileSystemRoot = "C:/Users/mobi/Desktop/tmp/Hot_sweat_season"
)

var gBean = bean.ZipImgBean{
	DirIndex: -1,
	ImgSize:  0,
}

func copyFile(f []*zip.File, root string) {
	fmt.Printf("zip file size=%d\n\n", len(f))
	for _, file := range f {
		//fmt.Printf("zip  file=%s  isDir=%d\n", file.Name, file.FileInfo().IsDir())
		if file.FileInfo().IsDir() {
			continue
			//空目录
			//subPath := filepath.Join(root, file.Name)
			//fmt.Printf("subPath=%s\n", subPath)
			//read, err := zip.OpenReader(subPath)
			//if err != nil {
			//	fmt.Errorf("")
			//	continue
			//}
			//copyFile(read.File, subPath)
			//read.Close()
		} else {
			name := strings.Split(file.Name, "/") //文件夹abc下文件也会遍历到：  adb/1.jpg
			_fileName := file.Name
			if len(name) > 1 {
				_fileName = name[len(name)-1]
			}
			path := filepath.Join(root, _fileName)
			fmt.Printf("des path=%s\n", path)
			_fileSuffix := path2.Ext(_fileName)
			if _, err := strconv.Atoi(_fileName[0 : len(_fileName)-len(_fileSuffix)]); err != nil {
				continue
			}
			fr, err := file.Open()
			if err != nil {
				fr.Close()
				continue
			}
			fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
			if err != nil {
				fw.Close()
				fr.Close()
				continue
			}
			n, err := io.Copy(fw, fr)
			if err != nil {
				fw.Close()
				fr.Close()
				continue
			}
			fmt.Printf("成功解压%s，写入数据%d\n\n", path, n)
			fw.Close()
			fr.Close()
			gBean.ImgSize += 1
			gBean.ImgSrc.PushBack(path)
		}
	}
}

func UnzipFile(zr *zip.Reader) bean.ZipImgBean {

	//获取路径下项目编号最大的
	files, _ := ioutil.ReadDir(FileSystemRoot + "/" + "img")
	maxTmpDir := 0
	for _, dir := range files {
		if dir.IsDir() {
			tmp, err := strconv.Atoi(dir.Name())
			if err == nil {
				if tmp > maxTmpDir {
					maxTmpDir = tmp
				}
			}
			fmt.Println(dir.Name())
		}
	}
	//生成当前项目编号
	maxTmpDir += 1
	fmt.Printf("next new dir =%d\n", maxTmpDir)
	dst := FileSystemRoot + "/" + "img" + "/" + fmt.Sprintf("%d", maxTmpDir)
	os.MkdirAll(dst, os.ModePerm)

	gBean.DirIndex = maxTmpDir
	gBean.ImgSize = 0
	gBean.ImgSrc = *list.New()
	copyFile(zr.File, dst)

	//for _, file := range zr.File {
	//
	//	//getFile(file,dst)
	//
	//	fmt.Println("zip file =" + file.Name)
	//	path := filepath.Join(dst, file.Name)
	//
	//	if file.FileInfo().IsDir() {
	//		if err := os.MkdirAll(path, file.Mode()); err != nil {
	//			return b
	//		}
	//		continue
	//	}
	//	_fileSuffix := path2.Ext(file.Name)
	//	if _, err := strconv.Atoi(file.Name[0 : len(file.Name)-len(_fileSuffix)]); err != nil {
	//		continue
	//	}
	//
	//	fr, err := file.Open()
	//	if err != nil {
	//		return b
	//	}
	//	fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
	//	if err != nil {
	//		return b
	//	}
	//	n, err := io.Copy(fw, fr)
	//	if err != nil {
	//		return b
	//	}
	//	fmt.Printf("成功解压%s，写入数据%d\n\n", path, n)
	//	fw.Close()
	//	fr.Close()
	//}

	return gBean
}

func ParseAutoDeploy(b bean.AutoDeployBean) bean.DeployResBean {
	res := bean.DeployResBean{
		ResUrl:  "",
		Success: false,
	}

	projectIndex := b.ZipRes
	title := b.Title
	content := b.Content
	imgUrl := b.ImgUrl
	name := b.Name

	dstHtmlFileName := fmt.Sprintf("%s_%s.html", projectIndex, name)

	files, _ := ioutil.ReadDir(FileSystemRoot + "/" + "img" + "/" + projectIndex)

	var fileList []string
	for _, file := range files {
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
			fmt.Println(file.Name())
		}
	}
	sort.Strings(fileList)
	imgSize := len(fileList)

	in, err := os.Open("index_src.html")
	if err != nil {
		return res
	}
	defer in.Close()

	out, err := os.OpenFile(FileSystemRoot+"/"+dstHtmlFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return res
	}
	defer out.Close()
	br := bufio.NewReader(in)
	index := 0
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return res
		}
		new1 := strings.Replace(string(line), "$title$", title, -1)
		new2 := strings.Replace(new1, "$content$", content, -1)
		new3 := strings.Replace(new2, "$img_url$", imgUrl, -1)
		if imgSize < 1 {
			if strings.Contains(new3, "img/16/1.jpg") {
				new3 = ""
			}
		} else {
			if strings.Contains(new3, "img/16/1.jpg") {
				new3 = strings.Replace(new3,
					"img/16/1.jpg",
					fmt.Sprintf("img/%s/%s", projectIndex, fileList[index]),
					-1)
				index++
				for index < imgSize {
					new3 = fmt.Sprintf("%s\n", new3)
					nextLine := fmt.Sprintf("\r<div class=\"img_lit\"><img src=\"img/%s/%s\" /></div>\n", projectIndex, fileList[index])
					new3 = fmt.Sprintf("%s%s", new3, nextLine)
					index++
				}

			}

		}
		_, err = out.WriteString(new3 + "\n")
		if err != nil {
			return res
		}
	}
	res.Success = true
	res.ResUrl = "https://api.mobifitness.cn/edith/activity/Hot_sweat_season/" + dstHtmlFileName

	return res
}
