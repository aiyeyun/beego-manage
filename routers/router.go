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
	beego.Router("/menu/index", &controllers.MenuController{})
	//菜单栏目 分页请求页
	beego.Router("/menu/index/?:p([0-9]+)", &controllers.MenuController{})
	//添加 编辑 菜单栏 Form 表单 视图页面
	beego.Router("/menu/form/?:id([0-9]+)", &controllers.MenuController{}, "get:Form;post:Form")
	//删除 栏目
	beego.Router("/menu/del/?:id([0-9]+)", &controllers.MenuController{}, "get:Del")


	//----------------------------------  管理员  ----------------------------------
	//管理员列表
	beego.Router("/admin/index/?:p([0-9])+", &controllers.AdminController{}, "get:Users")
	//添加管理员
	beego.Router("/admin/form", &controllers.AdminController{}, "*:Form")
	//角色列表
	beego.Router("/admin/role", &controllers.AdminController{})
	//角色授权
	beego.Router("/admin/auth/?:role([0-9]+)", &controllers.AdminController{}, "get:Auth;post:BatchAuth")


	//Errors 错误渲染
	beego.ErrorController(&controllers.ErrrosController{})

	//Filter 过滤器
	//beego.InsertFilter("/*", beego.BeforeExec, controllers.BeforeExec)
	//Filter 过滤器
	//beego.InsertFilter("/*", beego.BeforeStatic, controllers.BeforeStatic)
	//beego.InsertFilter("/*", beego.BeforeRouter, controllers.BeforeRouter)

	//自定义模板函数
	addFuncMap()
}

//自定义模板函数
func addFuncMap()  {
	beego.AddFuncMap("Menus", GetMenus)
	beego.AddFuncMap("MenuActive", MenuActive)
	beego.AddFuncMap("GetSuperAdmin", GetSuperAdmin)
	beego.AddFuncMap("SubMenuBoxOpen", SubMenuBoxOpen)
	beego.AddFuncMap("SubMenuBoxOpenStyle", SubMenuBoxOpenStyle)
}