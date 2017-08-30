package main

import (
	_ "manage/routers"
	_ "manage/models"
	_ "manage/initialize"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

