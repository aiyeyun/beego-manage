package routers

import (
	"manage/models"
	"strings"
)

//获取我的 menu 列表
func GetMyMenus() (map[string]interface{}) {
	model := models.Menu{}
	parent, subNode, _ := model.GetNodelAll(-1)
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