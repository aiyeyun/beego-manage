package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
	"github.com/astaxie/beego"
	"time"
)

type MenuAuthList []*MenuAuth

type MenuAuth struct {
	BaseModel
	Id      int    //主键ID
	Role    uint8  //角色
	MenuId  int    //菜单ID
	Created string //创建时间
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

//批量授权 超管授权使用方法
func (model *MenuAuth) BatchAuth(role uint8, menuIds []string) {
	//删除现在有的授权 准备注入角色新的授权
	o := orm.NewOrm()
	dbPrefix := beego.AppConfig.String("mysql_db_prefix")
	o.Raw("DELETE FROM "+ dbPrefix + model.TableName() +" WHERE role=" + strconv.Itoa(int(role)) ).Exec()

	//注入新的授权
	for _, v := range menuIds {
		authMenuId, _ := strconv.Atoi(v)
		o.Insert(&MenuAuth{
			Role: role,
			MenuId: authMenuId,
			Created: time.Now().Format("2006-01-02 15:04:05"),
		})
	}
}

//根据角色 获取授权栏目的列表 ids
func (model *MenuAuth) GetAuthMenusByRole() map[int]int {
	o := orm.NewOrm()
	menuAuthList := make(MenuAuthList, 0)
	o.QueryTable(model).Filter("role", model.Role).All(&menuAuthList, "menu_id")

	authMenus := make(map[int]int)
	for _, auth := range menuAuthList {
		authMenus[auth.MenuId] = auth.MenuId
	}

	return authMenus
}