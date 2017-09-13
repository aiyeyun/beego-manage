package models

import (
	"github.com/astaxie/beego/orm"
	"manage/utils/models/MenuModelUtil"
	"time"
	"github.com/astaxie/beego/validation"
	"errors"
	"manage/utils/page"
)

type Menu struct {
	BaseModel
	Id      int     `form:"id"`      //主键ID
	Pid     int     `form:"pid"`     //父节点ID
	Url     string  `form:"url"`     //Url
	Name    string  `form:"name"`    //菜单栏名称
	Icons   string  `form:"icons"`   //菜单图标
	Status  int     `form:"status"`  //菜单状态
	Sort    int     `form:"sort"`    //菜单排序
	Created int64                    //创建时间
	Update  int64                    //更新时间
}

//菜单列表
type MenuList []*Menu

//表名
func (model *Menu) TableName() string {
	return "menu"
}

//数据验证
func (model *Menu) Validate() error {
	valid := validation.Validation{}
	valid.Min(model.Pid, 0, "父节点").Message("父节点 错误类型")
	valid.Required(model.Name, "栏目名").Message("栏目名 不能为空")
	valid.Range(model.Status, 0, 1, "状态").Message("请选择状态")
	valid.Min(model.Sort, 1, "排序").Message("排序值最小为1")

	//场景 顶级栏目
	if model.Pid > 0 {
		valid.Required(model.Url, "Url").Message("Url 不能为空")
	} else {
		//场景 子栏目
		valid.Required(model.Icons, "Icons").Message("Icons 不能为空")
	}

	if valid.HasErrors() {
		return errors.New(valid.Errors[0].Message)
	}
	return nil
}

//新增 or 新增
func (model *Menu) Save() (int64, error) {
	err := model.Validate()
	if err != nil {
		return 0, err
	}

	o := orm.NewOrm()
	model.Update = time.Now().Unix()
	//编辑
	if model.Id > 0 {
		num, err := o.Update(model)
		return num, err
	}

	//新增
	model.Created = time.Now().Unix()
	id, err := o.Insert(model)
	return id, err
}

//获取所有父节点
func (model *Menu) GetParentNode() (MenuList, error) {
	o := orm.NewOrm()
	menuList := make(MenuList, 0)
	_, err := o.QueryTable(model).Filter("status", MenuModelUtil.STATUS_OPEN).Filter("pid", MenuModelUtil.PARENT_NODE).
		OrderBy("-sort").Limit(-1).All(&menuList)
	if err != nil {
		return nil, err
	}

	return menuList, nil
}

//分页返回
func (model *Menu) GetParentNodePaging(p, limit int) (MenuList, *page.PageLinks, error) {
	o := orm.NewOrm()
	query := o.QueryTable(model).Filter("status", MenuModelUtil.STATUS_OPEN).Filter("pid", MenuModelUtil.PARENT_NODE)
	pg := &page.PageLinks{
		Currentpage: p,
		PageDataSize: limit,
		Query: query,
	}


	menuList := make(MenuList, 0)
	_, err := o.QueryTable(model).Filter("status", MenuModelUtil.STATUS_OPEN).Filter("pid", MenuModelUtil.PARENT_NODE).
		OrderBy("-sort").Limit(limit).Offset(pg.GetOffset()).All(&menuList)

	if err != nil {
		return nil, nil, err
	}

	return menuList, pg, nil
}

//根据ID 获取model
func (model *Menu) GetModelById() error {
	o := orm.NewOrm()
	return o.Read(model)
}

//获取所有菜单栏目
func (model *Menu) GetNodelAll(p int) (MenuList, map[int]MenuList, *page.PageLinks) {
	//获取所有父节点
	var parent MenuList
	var pg *page.PageLinks
	//parent, _ := model.GetParentNode()
	//parent, pg, _ = model.GetParentNodePaging(p, 1)

	if p >= 0 {
		//需要分页
		p_num := p
		if p_num == 0 {
			p_num = 1
		}
		parent, pg, _ = model.GetParentNodePaging(p_num, 1)
	} else {
		//不需要分页
		parent, _ = model.GetParentNode()
	}

	node := make(map[int]MenuList, 0)
	nodeQuery := func(pid int) (MenuList, error) {
		o := orm.NewOrm()
		menuList := make(MenuList, 0)
		_, err := o.QueryTable(model).Filter("status", MenuModelUtil.STATUS_OPEN).
			Filter("pid", pid).
			OrderBy("-sort").Limit(-1).All(&menuList)
		if err != nil {
			return nil, err
		}
		return menuList, nil
	}
	for i := range parent {
		subNode, _ := nodeQuery(parent[i].Id)
		node[parent[i].Id] = subNode
	}

	//不需要分页 返回
	if p < 0 {
		return parent, node, nil
	}

	return parent, node, pg
}

//获取所有 菜单栏目
func (model *Menu) GetNodes() (MenuList, map[int]MenuList) {
	parent, _ := model.GetParentNode()
	nodes := make(map[int]MenuList, 0)

	nodeQuery := func(pid int) (MenuList, error) {
		o := orm.NewOrm()
		menuList := make(MenuList, 0)
		_, err := o.QueryTable(model).Filter("pid", pid).
			OrderBy("-sort").Limit(-1).All(&menuList)
		if err != nil {
			return nil, err
		}
		return menuList, nil
	}

	for i := range parent {
		subNode, _ := nodeQuery(parent[i].Id)
		nodes[parent[i].Id] = subNode
	}

	return parent, nodes
}