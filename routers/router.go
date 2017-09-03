package routers

import (
	"manage/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//路由注册
    //beego.Router("/", &controllers.MainController{})

	//----------------------------------  登录  ----------------------------------
	//登录
	beego.Router("/", &controllers.LoginContrller{})
    //退出登录
	beego.Router("/logout", &controllers.LoginContrller{}, "*:Logout")
	//锁定屏幕 解锁屏幕
	beego.Router("/lock", &controllers.LoginContrller{}, "get:Lock;post:Unlock")



	//----------------------------------  控制台  ----------------------------------
	//控制台
	beego.Router("/console", &controllers.ConsoleController{})



	//----------------------------------  菜单栏  ----------------------------------
	//菜单栏管理
	beego.Router("/menu", &controllers.MenuController{})
	//添加 编辑 菜单栏 Form 表单 视图页面
	beego.Router("/menu/form/?:id([0-9]+)", &controllers.MenuController{}, "get:Form;post:Form")


	//Errors 错误渲染
	beego.ErrorController(&controllers.ErrrosController{})

	//Filter 过滤器
	//beego.InsertFilter("/*", beego.BeforeExec, controllers.BeforeExec)
	//Filter 过滤器
	//beego.InsertFilter("/*", beego.BeforeStatic, controllers.BeforeStatic)
	//beego.InsertFilter("/*", beego.BeforeRouter, controllers.BeforeRouter)
}