package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeControllor{})
	beego.Router("/login", &controllers.LoginController{})
	//beego.Router("/", &controllers.MainController{})
}
