package rbac

import (
	"manage/models"
	"sync"
	"manage/utils/global"
)

var once sync.Once

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

//初始化 角色授权 栏目表
func InitRoleAuthMenus()  {
	once.Do(rabcInit)
}

//初始化 rbac
func rabcInit()  {
	//初始化 全局变量
	global.NewRoleAuthMenu()
	global.NewRoleAuthUrl()
	//获取 所有节点
	menuModel := models.Menu{}
	parent, subNode := menuModel.GetNodes()

	authMenuModel := models.MenuAuth{}
	for role, _ := range GetRoles() {
		//获取角色 所有授权的栏目IDS
		authMenuModel.Role = role
		authMenus := authMenuModel.GetAuthMenusByRole()
		UpdateRoleMenus(role, parent, subNode, authMenus)
		UpdateRoleUrl(role, parent, subNode, authMenus)
	}
}

//更新 角色栏目权限
func UpdateRoleMenus(role uint8, parent models.MenuList, subNode map[int]models.MenuList, authMenus map[int]int)  {
	node := make(map[string]interface{})
	//超管的权限
	if role == ROLE_SUPER_ADMIN {
		node["parent"] = parent
		node["subNode"] = subNode
		global.RoleAuthMenu.Set(role, node)
		return
	}

	//解析 角色权限
	newParent, newSubNode := analysisRoleAuthMenus(parent, subNode, authMenus)

	node["parent"] = newParent
	node["subNode"] = newSubNode
	global.RoleAuthMenu.Set(role, node)
}

//更新 角色路由权限
func UpdateRoleUrl(role uint8, parent models.MenuList, subNode map[int]models.MenuList, authMenus map[int]int)  {
	urls := make(map[string]string)
	urls["/console"] = "/console"

	//超管的权限
	if role == ROLE_SUPER_ADMIN {
		for _, parent_v := range parent {
			for _, sub_v:= range subNode[parent_v.Id] {
				if parent_v.Url == "" || sub_v.Url == "" {
					continue
				}
				urls[parent_v.Url + sub_v.Url] = parent_v.Url + sub_v.Url
			}
		}
		global.RoleAuthUrl.Set(role, urls)
		return
	}

	//解析 角色权限
	newParent, newSubNode := analysisRoleAuthMenus(parent, subNode, authMenus)

	for _, parent_v := range newParent {
		for _, sub_v:= range newSubNode[parent_v.Id] {
			if parent_v.Url == "" || sub_v.Url == "" {
				continue
			}
			urls[parent_v.Url + sub_v.Url] = parent_v.Url + sub_v.Url
		}
	}
	global.RoleAuthUrl.Set(role, urls)
}

//解析 角色授权 的栏目
//根据角色 获取所有栏目列表
func analysisRoleAuthMenus(parent models.MenuList, subNode map[int]models.MenuList, authMenus map[int]int) (models.MenuList,  map[int]models.MenuList) {
	newParentNode := make(models.MenuList, 0)
	newSubNode := make(map[int]models.MenuList)

	for _, parent_v := range parent {
		node := make(models.MenuList, 0)
		for _, sub_v := range subNode[parent_v.Id] {
			if _, ok := authMenus[sub_v.Id]; ok {
				node = append(node, sub_v)
			}
		}

		if len(node) >0 {
			newParentNode = append(newParentNode, parent_v)
			newSubNode[parent_v.Id] = node
		}
	}

	return newParentNode, newSubNode
}