package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"strings"
	"strconv"
	"manage/models"
	"manage/utils/global"
	"manage/utils/rbac"
)

// BaseController 基 controller
type BaseController struct {
	beego.Controller
}

var (
	//版权信息
	Copyright = `Cloud 夜版权所有 蜀 ICP 证 15034906 号`
)

//构造函数
func (c *BaseController) Prepare() {
	//路由验证
	c.RouteAuth()

	//读取 flash 消息
	beego.ReadFromRequest(&c.Controller)
	
	//xsrf
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["copyright"] = Copyright
	c.Data["requestUrl"] = c.Ctx.Request.RequestURI
	if !c.IsGuest() {
		c.Data["current_role"] = c.GetSession("user").(*models.Admin).Role
		c.Data["my_user"] = c.GetSession("user").(*models.Admin)
	}
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

	//检查角色路由权限
	if isGuest {
		return
	}

	//任何角色都有权限访问 退出登录接口 与 锁屏接口
	if runControllerName == "/logout" || runControllerName == "/lock" {
		return
	}

	//超管不用检查路由权限
	if c.GetSession("user").(*models.Admin).Role == rbac.ROLE_SUPER_ADMIN {
		return
	}

	auth := false
	roleUrls := global.RoleAuthUrl.Get(c.GetSession("user").(*models.Admin).Role)
	for k, _ := range roleUrls {
		if runControllerName == k || strings.Contains(runControllerName, k) {
			auth = true
		}
	}

	if !auth {
		c.Jump("大兄弟你没有权限啊")
	}
}

//跳转路由
func (c *BaseController) Jump(bodyTxt string, args ...interface{})  {
	c.Data["bodyTxt"] = bodyTxt
	//跳转来源地址
	referer := c.Ctx.Request.Referer()
	if referer == "" {
		referer = "/"
	}
	c.Data["url"] = referer
	c.Data["second"] = 6

	if len(args) > 0 {
		c.Data["url"] = args[0]
	}

	if len(args) >= 2 {
		c.Data["second"] = args[1]
	}

	c.TplName = "errors/jump.html"
	c.Render()
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

//设置错误 flash 消息
func (c *BaseController) SetErrorFlash(message string)  {
	beego.NewFlash()
	flash := beego.NewFlash()
	flash.Error(message)
	flash.Store(&c.Controller)
}

//设置成功 flash 消息
func (c *BaseController) SetSuccessFlash(message string) {
	beego.NewFlash()
	flash := beego.NewFlash()
	flash.Success(message)
	flash.Store(&c.Controller)
}

//获取 GET 参数 正则路由 与普通路由的参数
func (c *BaseController) GetParamInt(key string) (int, error) {
	val := c.Ctx.Input.Param(":"+key)
	if val != "" {
		val, err := strconv.Atoi(val)
		return val, err
	}
	return c.GetInt(key)
}

//获取 GET 参数 正则路由 与普通路由的参数
func (c *BaseController) GetParamUint8(key string) (uint8, error) {
	val := c.Ctx.Input.Param(":"+key)
	int64, err := strconv.ParseUint(val, 10, 8)
	return uint8(int64), err
}