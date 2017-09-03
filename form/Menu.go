package form

type MenuForm struct {
	Pid    int                          //父节点ID
	Url    string                       //Url
	Name   string `valid:"Required"`    //菜单栏名称
	Icons  string `` 					//菜单图表
	Status uint8  `valid:"Required"`    //菜单状态
	Sort   int    `valid:"Required"`    //菜单排序
}
