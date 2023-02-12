package controller

import (
	"ginTest/Response"
	"ginTest/forms"
	"ginTest/utils"
	"github.com/gin-gonic/gin"
)

//// PasswordLogin 登录
//func PasswordLogin(c *gin.Context) {
//	PasswordLoginForm := forms.PasswordLoginForm{}
//	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
//		color.Blue(err.Error())
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"err": err.Error(),
//		})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"msg": "sussess",
//	})
//}
// PasswordLogin 登录
func PasswordLogin(c *gin.Context) {
	PasswordLoginForm := forms.PasswordLoginForm{}
	if err := c.ShouldBind(&PasswordLoginForm); err != nil {
		// 统一处理异常
		utils.HandleValidatorError(c, err)
		return
	}
	Response.Success(c, 200, "success", "test")
}
