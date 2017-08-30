package models

import (
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"
)

type Admin struct {
	Id       int
	Username string
	Nickname string
	Email    string
	Password string
	Salt     string
	Role     int8   //0~127
	Last_ip  string
	Created  int
	Update   int64
}

//表名
func (model *Admin) TableName() string {
	return "admin"
}

//登录
func (model *Admin) Login(username, password, ip string) (*Admin, error) {
	o := orm.NewOrm()
	err := o.QueryTable(model).Filter("username", username).One(model)
	if err != nil {
		return nil, errors.New("账号或密码错误")
	}

	//检查密码是否正确
	pwd := password + model.Salt
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(pwd))
	cipherStr := md5Ctx.Sum(nil)
	md5Psw := hex.EncodeToString(cipherStr)
	if md5Psw != model.Password {
		return nil, errors.New("账号或密码错误")
	}

	model.Last_ip = ip
	model.Update  = time.Now().Unix()
	//更新登录信息
	o.Update(model, "last_ip", "update")

	return model, nil
}