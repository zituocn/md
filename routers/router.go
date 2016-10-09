package routers

import (
	"github.com/astaxie/beego"
	"github.com/zituocn/md/controllers"
)

func init() {
	beego.Router("/md/*", &controllers.MdHandle{})
}
