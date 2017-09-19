package initialize

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"manage/models"
	"manage/utils/rbac"
)

//ORM 初始化
func init()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	dataSource := beego.AppConfig.String("mysql_user") +
		":" +
		beego.AppConfig.String("mysql_pass") +
		"@tcp(" +
		beego.AppConfig.String("mysql_host") +
		":" +
		beego.AppConfig.String("mysql_port") +
		")/" +
		beego.AppConfig.String("mysql_db_name") +
		"?charset=" +
		beego.AppConfig.String("mysql_charset")

	//表前缀
	db_prefix := beego.AppConfig.String("mysql_db_prefix")

	orm.RegisterDataBase("default", "mysql", dataSource)

	orm.Debug = true

	//根据数据库的别名，设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)
	//根据数据库的别名，设置数据库的最大数据库连接
	orm.SetMaxOpenConns("default", 30)

	//注册 Model
	orm.RegisterModelWithPrefix(db_prefix, new(models.Admin))
	orm.RegisterModelWithPrefix(db_prefix, new(models.Menu))
	orm.RegisterModelWithPrefix(db_prefix, new(models.MenuAuth))

	//获取所有角色授权的栏目列表
	rbac.InitRoleAuthMenus()
}