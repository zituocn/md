package controllers

import (
	"github.com/astaxie/beego"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
)

type MdHandle struct {
	beego.Controller
}

func (this *MdHandle) Get() {
	var (
		dRoot   = "file"
		link    string
		content string
	)

	link = this.GetString(":splat")
	beego.Info(link)

	if len(link) == 0 {
		this.Abort("404")
	}

	//文件的绝对地址
	link = dRoot + "/" + link
	file, err := os.Open(link)
	if err != nil {
		beego.Debug(err.Error())
		this.Abort("404")
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		beego.Debug(err.Error())

	}
	if len(data) > 0 {
		common := blackfriday.MarkdownCommon(data)
		content = string(common)
	} else {
		content = "文章空白"
	}
	this.Data["content"] = content
	this.TplName = "_readfile.html"
}
