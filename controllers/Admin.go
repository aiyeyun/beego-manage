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
	roles := rbac.GetRoles()
	delete(roles, rbac.ROLE_SUPER_ADMIN)
	c.Data["role_list"] = roles
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

//管理员列表
func (c *AdminController) Users()  {
	c.Display("index")
}

//添加编辑 用户信息
func (c *AdminController) Form()  {
	model := &models.Admin{Id: id}

	if c.Ctx.Input.IsPost() {
		c.ParseForm(model)
		err := model.Save(rbac.ROLE_SUPER_ADMIN, c.GetSession("user").(*models.Admin))
		if err == nil {
			c.SetSuccessFlash("管理员添加成功")
			c.Redirect("/admin/index", 302)
			return
		}
		c.SetErrorFlash(err.Error())
	}

	//获取所有角色
	c.Data["role_list"] = rbac.GetRoles()
	c.Data["model"]     = model
	c.Display("_form")
}