package controllers

type MenuController struct {
	BaseController
}

//列表
func (c *MenuController) Get()  {
	c.Display("index.tpl")
}

//表单
func (c *MenuController) Form()  {
	c.Display("_form")
}