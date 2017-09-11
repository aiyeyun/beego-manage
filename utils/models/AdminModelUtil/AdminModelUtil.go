package AdminModelUtil

const (
	ROLE_SUPER_ADMIN = 1 //超级管理员
	ROLE_AUDITOR     = 2 //审查员
	ROLE_EDITOR      = 3 //编辑员
	ROLE_FINANCE     = 4 //财务
)

//获取所有角色
func GetRoles() map[int]string {
	roles := make(map[int]string)
	roles[ROLE_SUPER_ADMIN] = "超级管理员"
	roles[ROLE_AUDITOR] = "审查员"
	roles[ROLE_EDITOR] = "编辑员"
	roles[ROLE_FINANCE] = "财务"
	return roles
}