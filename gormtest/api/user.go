package api

import (
	"github.com/gin-gonic/gin"
	"test.com/gormtest/dao"
	"time"
)

func SaveUser(c *gin.Context) {
	user := &dao.User{
		Username:   "lisi",
		Password:   "123456",
		CreateTime: time.Now().UnixMilli(),
	}
	dao.SaveUser(user)
	c.JSON(200, user)
}

func GetUser(c *gin.Context) {
	user := dao.GetById(1)
	c.JSON(200, user)
}

func GetAllUser(c *gin.Context) {
	user := dao.GetAll()
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	dao.UpdateUser(1)
	user := dao.GetById(1)
	c.JSON(200, user)
}
func DeleteUser(c *gin.Context) {
	dao.DeleteUser(1)
	user := dao.GetById(1)
	c.JSON(200, user)
}
