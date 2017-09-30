package controllers

import (
	"manage/utils/rbac"
	"manage/models"
	"strconv"
	"manage/utils/file"
	"fmt"
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
	p, _ := c.GetParamInt("p")
	model := models.Admin{}
	admins, page := model.Admins(p, 2)
	page.Href = "/admin/index"
	c.Data["admins"] = admins
	c.Data["page"] = page.GetLinks()
	c.Display("index")
}

//添加编辑 用户信息
func (c *AdminController) Form()  {
	model := &models.Admin{}
	if c.Ctx.Input.IsPost() {
		c.ParseForm(model)
		filePath, filerr := file.ImagesUpload(c.Ctx.Request.FormFile("portrait"))
		fmt.Println("上传是纳闷回事", filePath, filerr)
		if filePath != "" {
			model.Portrait = filePath
		}
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

//修改密码
func (c *AdminController) UpdatePsw()  {
	newPsw := c.GetString("newPsw")
	uid, _ := c.GetInt("uid")
	model := models.Admin{Id:uid}
	err := model.UpdatePassword(rbac.ROLE_SUPER_ADMIN, newPsw, c.GetSession("user").(*models.Admin))

	data := make(map[string]interface{})
	c.Data["json"] = data
	if err != nil {
		data["status"] = false
		data["msg"] = err.Error()
		c.ServeJSON()
	}

	data["status"] = true
	data["msg"] = "修改成功"
	c.ServeJSON()
}