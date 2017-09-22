package routers

import (
	"strings"
	"manage/utils/global"
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

//被点击后 的menu Calss 显示
func ParentMenuActive(url, testUrl string) string {
	if testUrl == "" {
		return ""
	}
	if strings.Contains(url, testUrl) {
		return " active "
	}
	return ""
}

//子栏目 class 显示
func SubMenuActive(url, parentUrl, subUrl string) string {
	if parentUrl + subUrl == url {
		return  " active "
	}
	return ""
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
	//if url == testUrl || strings.Contains(url, testUrl) {
	if url == testUrl {
		return  " active "
	}
	return ""
}