package form

import (
	"errors"
	"fmt"
)

//登录表单
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

//验证
func (form *LoginForm) Verification() error {
	if form.Username == "" || form.Password == "" {
		return errors.New("用户名或密码不能为空")
	}
	return nil
}

//登录
func (form *LoginForm) Login()  {
	fmt.Println(form)
}

