package main

import (
	"github.com/astaxie/beego"
	"github.com/smart-as/website/models"
	"github.com/smart-as/website/routers"
)

func main() {
	models.Init()
	routers.Init()
	beego.Run()
}
