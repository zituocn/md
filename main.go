package main

import (
	"github.com/astaxie/beego"
	_ "github.com/zituocn/md/routers"
	"html/template"
)

//返回状态名称
func getDirFileStatus(isdir bool) template.HTML {
	var html string
	if isdir {
		html = "<span class=\"dir\">目录</span>"
	} else {
		html = "<span class=\"file\">文件</span>"
	}

	return template.HTML(html)
}

func main() {
	beego.AddFuncMap("getDirFileStatus", getDirFileStatus)
	beego.Run()
}
