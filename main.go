package main

import (
	_ "github.com/yaolynzc/amecms-g/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/fonts", "static/fonts")
	beego.SetStaticPath("/images", "static/images")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/assets", "static/assets")
	beego.SetStaticPath("/js", "static/js")
	beego.SetStaticPath("/favicon", "static/favicon")
	beego.Run()
}

