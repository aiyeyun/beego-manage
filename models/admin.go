package models

import (
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"
	"github.com/astaxie/beego/validation"
)

type Admin struct {
	BaseModel
	Id       int
	Username string `form:"username"`
	Nickname string
	Email    string
	Password string `form:"password"`
	Salt     string
	Role     uint8   //0~255
	Last_ip  string
	Created  int
	Update   int64
}

//表名
func (model *Admin) TableName() string {
	return "admin"
}

//登录验证
func (model *Admin) Validate() error {
	valid := validation.Validation{}
	valid.Required(model.Username, "用户名").Message("请填写用户名")
	valid.Required(model.Password, "密码").Message("请填写密码")
	if valid.HasErrors() {
		return errors.New(valid.Errors[0].Message)
	}
	return nil
}

//登录
func (model *Admin) Login() error {
	password := model.Password
	ip := model.Last_ip
	o := orm.NewOrm()
	err := o.QueryTable(model).Filter("username", model.Username).One(model)
	if err != nil {
		return err
	}

	//检查密码是否正确
	pwd := password + model.Salt
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(pwd))
	cipherStr := md5Ctx.Sum(nil)
	md5Psw := hex.EncodeToString(cipherStr)
	if md5Psw != model.Password {
		return errors.New("账号或密码错误")
	}

	//更新登录信息
	model.Last_ip = ip
	model.Update = time.Now().Unix()
	o.Update(model, "last_ip", "update")
	return nil
}