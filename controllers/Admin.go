package controllers

type AdminController struct {
	BaseController
}

//管理员列表
func (c *AdminController) Get()  {
	c.Ctx.WriteString("list")
}

//表单
func (c *AdminController) Form()  {
	c.Ctx.WriteString("form")
}