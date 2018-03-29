package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.gncoder.com"
	c.Data["Email"] = "yaolynzc@126.com"
	c.TplName = "login.html"
}

func (c *MainController) Index(){
	c.TplName = "index.html"
}
