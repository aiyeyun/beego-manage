package routers

import (
	"strings"
	"manage/utils/global"
	"manage/utils/rbac"
)

//获取角色 menu 列表
func GetMenus(role uint8) (map[string]interface{}) {
	parent  := global.RoleAuthMenu.Get(role).(map[string]interface{})["parent"]
	subNode := global.RoleAuthMenu.Get(role).(map[string]interface{})["subNode"]
	nodes := make(map[string]interface{})
	nodes["parent"] = parent
	nodes["subNode"] = subNode
	return nodes
}

//获取超级管理员 角色
func GetSuperAdmin() uint8 {
	return rbac.ROLE_SUPER_ADMIN
}

//获取所有 角色
func GetRoles() map[uint8]string {
	return rbac.GetRoles()
}

//子栏目 是否展开
func SubMenuBoxOpen(url, testUrl string) string {
	if testUrl == "" {
		return ""
	}
	if strings.Contains(url, testUrl) {
		return " in "
	}
	return ""
}

//子栏目 是否展开的 style 样式
func SubMenuBoxOpenStyle(url, testUrl string) string {
	if testUrl == "" {
		return ""
	}
	if strings.Contains(url, testUrl) {
		return ""
	}
	return "height: 0px;"
}

//栏目是否有 active class
func MenuActive(url string, agrs ...string) string {
	var testUrl string = ""
	for _, v := range agrs {
		testUrl += v
	}

	if url == "" || testUrl == "" {
		return  ""
	}

	if url == testUrl {
		return  " active "
	}
	return ""
}

//URL 后缀是GET参数 ID
func urlSuffixIsGetParam(url string) bool {
	if strings.HasSuffix(url, "1") ||
		strings.HasSuffix(url, "2") ||
		strings.HasSuffix(url, "3") ||
		strings.HasSuffix(url, "4") ||
		strings.HasSuffix(url, "5") ||
		strings.HasSuffix(url, "6") ||
		strings.HasSuffix(url, "7") ||
		strings.HasSuffix(url, "8") ||
		strings.HasSuffix(url, "9") ||
		strings.HasSuffix(url, "0") {
			return true
	}
	return false
}