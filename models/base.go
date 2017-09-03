package models

import (
	"github.com/astaxie/beego/orm"
)

type BaseInterface interface {
	TableName() string
	Validate() error
	Save() (int64, error)
}

type BaseModel struct {

}

//根据ID 获取对象
func (base *BaseModel) GetModelById() error {
	o := orm.NewOrm()
	return o.Read(base)
}