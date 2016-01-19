package controllers

import (
	"blog/models"

	"fmt"
	"github.com/astaxie/beego"
)

type HomeControllor struct {
	beego.Controller
}

func (this *HomeControllor) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "home.html"
	//his.Data["IsLogin"] = checkAccount(this.Ctx)
	str := models.GetAllTopics(true)
	fmt.Println(str)
	this.Data["Topics"] = str
}
