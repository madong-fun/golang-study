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
	this.TplName = "home.html"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	fmt.Println(checkAccount(this.Ctx))
	topics, err := models.GetAllTopics("", "", true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics
}
