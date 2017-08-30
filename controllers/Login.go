package controllers

import (
	"github.com/astaxie/beego"
	"manage/form"
	"crypto/md5"
	"encoding/hex"
	"manage/models"
)

type LoginContrller struct {
	BaseController
}

//视图
func (c *LoginContrller) Get()  {
	c.TplName = "login/index.tpl"
}

//登录
func (c *LoginContrller) Post()  {
	loginForm := new(form.LoginForm)
	c.ParseForm(loginForm)
	err := loginForm.Verification()
	if err != nil {
		flash := beego.NewFlash()
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/", 302)
	}

	//登录
	adminUser, err := AdminModel.Login(loginForm.Username, loginForm.Password, c.Ctx.Input.IP())
	if err != nil {
		flash := beego.NewFlash()
		flash.Error(err.Error())
		flash.Store(&c.Controller)
		c.Redirect("/", 302)

		//此处要 return 不然 会继续往下执行
		return
	}

	//登录成功 Set Session
	c.SetSession("user", adminUser)

	//重定向控制台
	c.Redirect("/console", 302)
}

//退出登录
func (c *LoginContrller) Logout()  {
	c.DelSession("user")
	c.Redirect("/", 302)
}

//锁定屏幕
func (c *LoginContrller) Lock()  {
	user := c.GetSession("user")
	//未登录用户 重定向到 登录页面
	if user == nil {
		c.Redirect("/", 302)
		return
	}
	//登录用户 锁定
	c.SetSession("user-lock", true)
	c.TplName = "login/lock.tpl"
}

//解锁屏幕
func (c *LoginContrller) Unlock()  {
	user := c.GetSession("user")
	//未登录用户 重定向到 登录页面
	if user == nil {
		c.Redirect("/", 302)
		return
	}

	//获取来源地址
	refererUrl := c.Ctx.Request.Referer()
	if refererUrl == "" {
		//来源地址为空时 返回主页
		refererUrl = "/"
	}

	//登录用户 解锁
	lock := c.GetSession("user-lock")
	//如果未锁定 就返回来源地址
	if lock == nil || lock.(bool) != true {
		c.Redirect(refererUrl, 302)
		return
	}

	//获取输入的解锁密码
	password := c.GetString("password")
	if password == "" {
		flash := beego.NewFlash()
		flash.Error("请填写解锁密码")
		flash.Store(&c.Controller)
		c.Redirect("/lock", 302)
		return
	}

	//加密输入的密码
	password += user.(*models.Admin).Salt
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	cipherStr := md5Ctx.Sum(nil)
	md5Psw := hex.EncodeToString(cipherStr)

	//验证密码是否正确
	if user.(*models.Admin).Password != md5Psw {
		flash := beego.NewFlash()
		flash.Error("密码错误")
		flash.Store(&c.Controller)
		c.Redirect("/lock", 302)
		return
	}

	//解锁成功
	c.DelSession("user-lock")
	c.Redirect("/console", 302)
}