package global

import "sync"

var RoleAuthMenu *roleAuthMenus

//角色 授权 栏目列表
type roleAuthMenus struct {
	List map[uint8]interface{}
	m *sync.RWMutex
}