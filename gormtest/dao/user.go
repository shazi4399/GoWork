package dao

import (
	"log"
)

type User struct {
	ID         int64
	Username   string `gorm:"column:username"`
	Password   string `gorm:"column:password"`
	CreateTime int64  `gorm:"column:createtime"`
}

func (u User) TableName() string {
	return "users"
}

//增加
func SaveUser(user *User) {
	err := DB.Create(user).Error
	if err != nil {
		log.Println("insert user err", err)
	}
}

//查询
func GetById(id int64) User {
	var user User
	err := DB.Where("id=?", id).First(&user)
	if err != nil {
		log.Println("get user by id error", err)
	}
	return user
}

//查找全部
func GetAll() []User {
	var user []User
	err := DB.Find(&user).Error
	if err != nil {
		log.Println("get user by id error", err)
	}
	return user
}

//更新
func UpdateUser(id int64) {
	err := DB.Model(&User{}).Where("id=?", id).Update("username", "wangwu")
	if err != nil {
		log.Println("update user by id error", err)
	}
}

//删除
func DeleteUser(id int64) {
	err := DB.Where("id=?", id).Delete(&User{})
	if err != nil {
		log.Println("del user by id err ", err)
	}
}
