package bean

import "container/list"

type ProtocolDownLoadBean struct {
	MidiType string `json:"type"`
}

type ZipImgBean struct {
	DirIndex int
	ImgSize  int
	ImgSrc   list.List
}

type AutoDeployBean struct {
	ImgUrl  string `json:"img_url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	ZipRes  string `json:"zipRes"`
	Name    string `json:"name"`
}

type DeployResBean struct {
	Success bool
	ResUrl  string
}
