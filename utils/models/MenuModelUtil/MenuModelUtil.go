package MenuModelUtil

const (
	//菜单 显示状态 开
	STATUS_OPEN  = 1
	//菜单 显示状态 开
	STATUS_CLOSE = 0
	//菜单 父节点 Pid值
	PARENT_NODE  = 0
)

type Status struct {
	Status int
	Name   string
}

type StatusList []*Status

//返回 所有状态码
func GetStatus() StatusList {
	status := make(StatusList, 0)
	status = append(status, &Status{Status: STATUS_CLOSE, Name: "关"})
	status = append(status, &Status{Status: STATUS_OPEN, Name: "开"})
	return status
}