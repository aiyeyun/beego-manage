package controllers

import (
	"github.com/astaxie/beego"
	"manage/models"
	"html/template"
	"strings"
)

// BaseController 基 controller
type BaseController struct {
	beego.Controller
}

var (
	//版权信息
	Copyright = `Cloud 夜版权所有 蜀 ICP 证 15034906 号`
	//管理员 表
	AdminModel models.Admin
)


//构造函数
func (c *BaseController) Prepare() {
	//路由验证
	//c.RouteAuth()
	
	//读取 flash 消息
	beego.ReadFromRequest(&c.Controller)
	
	//xsrf
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["copyright"] = Copyright
	c.Data["wocao"] = "握草"
}

//路由验证
func (c *BaseController) RouteAuth()  {
	//获取请求控制器名
	runControllerName := c.Ctx.Request.RequestURI

	//是否未登录
	isGuest := c.IsGuest()

	//未登录 访问其他路由 重定向至登录页面
	if isGuest == true && runControllerName != "/" && runControllerName != "/logout" {
		//所有访问 其他路由 尚未登录的操作 (除退出登录) 都将 重定向登录页
		c.Redirect("/", 302)
		return
	}

	//登录后 锁屏状态不能访问任何URL (除退出登录)
	lock := c.GetSession("user-lock")
	if !isGuest && lock != nil && lock.(bool) == true && runControllerName != "/logout" && runControllerName != "/lock" {
		//重定向至 锁屏页面
		c.Redirect("/lock", 302)
		return
	}

	//登录了 继续访问登录页面 重定向控制台
	if isGuest == false && runControllerName == "/" {
		c.Redirect("/console", 302)
		return
	}

}

//模板输出
func (c *BaseController) Display(template string)  {
	c.TplExt = "html"
	controllerName, _ := c.GetControllerAndAction()
	cName := strings.ToLower(string([]byte(controllerName)[0:strings.Index(controllerName, "Controller")]))
	tplPath := cName + "/"
	template = tplPath + template

	c.TplName = template + ".html"
}

//是否是宾客
func (c *BaseController) IsGuest() bool {
	user := c.GetSession("user")
	if user != nil {
		return false
	}
	return true
}