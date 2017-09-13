package models

import (
	"github.com/astaxie/beego/orm"
)

type MenuAuthList []*MenuAuth

type MenuAuth struct {
	BaseModel
	Id      int  //主键ID
	Role    uint8 //角色
	MenuId  int  //菜单ID
	Created int  //创建时间
	Update  int  //更新时间
}

//表名
func (model *MenuAuth) TableName() string {
	return "menu_auth"
}

//获取已授权的栏目列表 超管授权使用方法
func (model *MenuAuth) GetAuthMenus() (MenuList, map[int]MenuList, map[int]int) {
	o := orm.NewOrm()
	//查询角色 授权栏目ID 列表
	menuAuthList := make(MenuAuthList, 0)
	o.QueryTable(model).Filter("role", model.Role).All(&menuAuthList, "menu_id")
	authMenus := make(map[int]int)
	for _, auth := range menuAuthList {
		authMenus[auth.MenuId] = auth.MenuId
	}

	//获取所有栏目
	menuModel := Menu{}
	parent, subNode := menuModel.GetNodes()

	return parent, subNode, authMenus
}