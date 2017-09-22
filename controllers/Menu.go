package controllers

import (
	"manage/utils/models/MenuModelUtil"
	"manage/models"
	"manage/utils/rbac"
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
	page.Href = "/menu/index"
	c.Data["page"] = page.GetLinks()
	c.Data["parent"] = parent
	c.Data["subNode"] = subNode

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
		fmt.Println("我看赋值厚", model)
		_, err := model.Save()
		if err == nil {
			//更新 角色路由 与 角色栏目 权限 更新全局栏目缓存
			rbac.RabcUpdate()
			c.Redirect("/menu/index", 302)
			return
		}
		c.SetErrorFlash(err.Error())
	}

	c.Display("_form")
}