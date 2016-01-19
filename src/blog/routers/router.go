package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeControllor{})
	//beego.Router("/", &controllers.MainController{})
}
