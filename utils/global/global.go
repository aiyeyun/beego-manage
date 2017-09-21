package global

import "sync"

var RoleAuthMenu *roleAuthMenus

var RoleAuthUrl *roleAuthUrls

//角色 授权 栏目列表
type roleAuthMenus struct {
	List map[uint8]interface{}
	m sync.RWMutex
}

//角色 授权 栏目URL
type roleAuthUrls struct {
	List map[uint8]map[string]string
	m sync.RWMutex
}

//初始化 角色 授权 栏目列表
func NewRoleAuthMenu()  {
	RoleAuthMenu = &roleAuthMenus{
		List: make(map[uint8]interface{}),
	}
}

//初始化 角色 授权 栏目URL
func NewRoleAuthUrl()  {
	RoleAuthUrl = &roleAuthUrls{
		List: make(map[uint8]map[string]string),
	}
}