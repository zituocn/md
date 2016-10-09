package controllers

import (
	"github.com/astaxie/beego"
)

type IndexHandle struct {
	beego.Controller
}

func (this *IndexHandle) Get() {

	this.TplName = "_index.html"
}
