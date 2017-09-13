package rbac

import "manage/models"

const (
	ROLE_SUPER_ADMIN uint8 = 1 //超级管理员
	ROLE_AUDITOR     uint8 = 2 //审查员
	ROLE_EDITOR      uint8 = 3 //编辑员
	ROLE_FINANCE     uint8 = 4 //财务
)

//获取所有角色
func GetRoles() map[uint8]string {
	roles := make(map[uint8]string, 0)
	roles[ROLE_SUPER_ADMIN] = "超级管理员"
	roles[ROLE_AUDITOR] = "审查员"
	roles[ROLE_EDITOR] = "编辑员"
	roles[ROLE_FINANCE] = "财务"
	return roles
}

//获取角色授权的 menu 列表
func GetRoleAuthMenus(role uint8) (models.MenuList, map[int]models.MenuList, map[int]int) {
	model := models.MenuAuth{}
	model.Role = role
	return model.GetAuthMenus()
}