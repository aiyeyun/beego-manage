package models

import (
	"github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"time"
	"github.com/astaxie/beego/validation"
	"manage/utils"
	"manage/utils/page"
)

type AdminList []*Admin

type Admin struct {
	BaseModel
	Id       int
	Username string `form:"username"`
	Nickname string `form:"nickname"`
	Email    string `form:"email"`
	Password string `form:"password"`
	Portrait string `form:"portrait"`
	Salt     string
	Role     uint8  `form:"role"` //0~255
	Last_ip  string
	Created  string
	Update   string
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

//添加管理员验证
func (model *Admin) ValidateSave() error {
	valid := validation.Validation{}
	valid.Required(model.Username, "用户名").Message("请填写用户名")
	valid.Required(model.Password, "密码").Message("请填写密码")
	valid.Required(model.Nickname, "昵称").Message("请填写昵称")
	valid.Required(model.Email, "邮箱").Message("请填写邮箱地址")
	valid.Required(model.Role, "角色").Message("请选择角色")
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
	model.Update = time.Now().Format("2006-01-02 15:04:05")
	model.Update = time.Now().Format("2006-01-02 15:04:05")
	o.Update(model, "last_ip", "update")
	return nil
}

//根据ID 获取model
func (model *Admin) GetModelById() error {
	o := orm.NewOrm()
	return o.Read(model)
}

//添加管理员
func (model *Admin) Save(superAdmin uint8, user *Admin) error {
	err := model.ValidateSave()
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	querySeter := o.QueryTable(model)
	model.Update = time.Now().Format("2006-01-02 15:04:05")

	//新增
	//检查 账号 邮箱 昵称是否存在
	if querySeter.Filter("username", model.Username).Exist() {
		return errors.New("用户名已存在")
	}

	if querySeter.Filter("nickname", model.Nickname).Exist() {
		return errors.New("昵称已存在")
	}

	if querySeter.Filter("email", model.Email).Exist() {
		return errors.New("邮箱已存在")
	}

	model.Created  = time.Now().Format("2006-01-02 15:04:05")
	//加密密码
	model.Salt     = utils.GetRandomString(6)
	model.Password = model.Password + model.Salt
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(model.Password))
	cipherStr := md5Ctx.Sum(nil)
	md5Psw := hex.EncodeToString(cipherStr)
	model.Password = md5Psw
	_, err = o.Insert(model)
	return err
}

//所有管理员
func (model *Admin) Admins(p int, limit int) (AdminList, *page.PageLinks) {
	o := orm.NewOrm()
	adminList := make(AdminList, 0)
	querySeter := o.QueryTable(model)
	page := &page.PageLinks{
		Currentpage: p,
		PageDataSize: limit,
		Query: querySeter,
	}

	querySeter.OrderBy("created").Limit(limit).Offset(page.GetOffset()).All(&adminList)
	return adminList, page
}