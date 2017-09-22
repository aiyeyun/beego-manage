package controllers

import (
	"manage/utils/models/MenuModelUtil"
	"manage/models"
	"fmt"
)

type MenuController struct {
	BaseController
}

//列表
func (c *MenuController) Get()  {
	p, _ := c.GetParamInt("p")
	model := new(models.Menu)

	parent, subNode, page := model.GetNodelAll(p)
	page.Href = "/menu"
	c.Data["page"] = page.GetLinks()
	c.Data["parent"] = parent
	c.Data["subNode"] = subNode

	fmt.Println("sssssssssss", parent[0])
	c.Display("index")
}

//表单
func (c *MenuController) Form()  {
	id, _ := c.GetParamInt("id")

	model := &models.Menu{Id: id}
	err := model.GetModelById()
	if err != nil && id > 0 {
		c.SetErrorFlash("找不到该记录")
		c.Redirect("/menu", 302)
		return
	}

	parentNode, _ := model.GetParentNode()
	c.Data["parentNode"] = parentNode
	c.Data["status"] = MenuModelUtil.GetStatus()
	c.Data["model"] = model

	if c.Ctx.Input.IsPost() {
		c.ParseForm(model)
		_, err := model.Save()
		if err == nil {
			c.Redirect("/menu", 302)
			return
		}
		c.SetErrorFlash(err.Error())
	}

	c.Display("_form")
}