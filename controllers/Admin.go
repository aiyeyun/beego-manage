package controllers

import (
	"manage/utils/rbac"
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

	parent, subNode, authList := rbac.GetRoleAuthMenus(role)
	c.Data["parent"]   = parent
	c.Data["subNode"]  = subNode
	c.Data["authList"] = authList

	c.Display("role/_auth")
}