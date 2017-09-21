package controllers

import (
	"manage/utils/rbac"
	"manage/models"
	"strconv"
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
	c.Data["role"]     = role
	c.Data["parent"]   = parent
	c.Data["subNode"]  = subNode
	c.Data["authList"] = authList
	c.Display("role/_auth")
}

//栏目批量授权 POST
func (c *AdminController) BatchAuth()  {
	role, _ := c.GetUint8("role")
	roleStr := strconv.Itoa(int(role))
	model := models.MenuAuth{}
	model.BatchAuth(role, c.GetStrings("auth[]"))

	//更新 角色路由与栏目权限
	menuModel := models.Menu{}
	parent, subNode := menuModel.GetNodes()
	model.Role = role
	authMenus := model.GetAuthMenusByRole()
	rbac.UpdateRoleMenus(role, parent, subNode, authMenus)
	rbac.UpdateRoleUrl(role, parent, subNode, authMenus)

	c.SetSuccessFlash("批量授权成功")
	c.Redirect("/admin/auth/" + roleStr, 302)
}