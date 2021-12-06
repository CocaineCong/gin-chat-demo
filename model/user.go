package model

import (
	"github.com/jinzhu/gorm"
)

//User 用户模型
type User struct {
	gorm.Model
	UserName       string
	Email          string 	//`gorm:"unique"`
	Avatar         string 	`gorm:"size:1000"`
	Phone 		   string
	BanTime 	   int
	OpenID 		   string 	`gorm:"unique"`
}

const (
	PassWordCost        = 12         //密码加密难度
	Active       string = "active"   //激活用户
	Inactive     string = "inactive" //未激活用户
	Suspend      string = "suspend"  //被封禁用户
)

//GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := DB.First(&user, ID)
	return user, result.Error
}

//AvatarUrl 封面地址
func (user *User) AvatarURL() string {
	signedGetURL := user.Avatar
	return signedGetURL
}
