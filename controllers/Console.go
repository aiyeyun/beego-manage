package controllers

type ConsoleController struct {
	BaseController
}

func (c *ConsoleController) Get()  {
	c.Display("index.tpl")
}