package controllers

import (
	"manage/utils/rbac"
	"manage/models"
)

type AdminController struct {
	BaseController
}

//管理角色列表
func (c *AdminController) Get()  {
	c.Data["role_list"] = rbac.GetRoles()
	c.Display("role/_list")
}

//管理角色 授权
func (c *AdminController) Auth()  {
	role, _ := c.GetParamUint8("role")

	if c.Ctx.Input.IsPost() {
		model := models.MenuAuth{}
		model.BatchAuth(role, c.GetStrings("auth[]"))
	}

	parent, subNode, authList := rbac.GetRoleAuthMenus(role)
	c.Data["parent"]   = parent
	c.Data["subNode"]  = subNode
	c.Data["authList"] = authList
	c.SetSuccessFlash("批量授权成功")

	c.Display("role/_auth")
}