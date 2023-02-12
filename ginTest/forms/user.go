package forms

type PasswordLoginForm struct {
	// 密码  binding:"required"为必填字段,长度大于3小于20
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20"`
	//用户名
	Username string `form:"name" json:"name" binding:"required,mobile"`
}
