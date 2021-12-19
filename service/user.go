package service

import (
	"chat/model"
	"chat/pkg/e"
	"chat/serializer"
	logging "github.com/sirupsen/logrus"
)

//UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	NickName string `form:"nick_name" json:"nick_name" binding:"required,min=2,max=10"`
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=16"`
}

func (service UserRegisterService) Register() serializer.Response {
	var user model.User
	var count int
	code := e.SUCCESS
	model.DB.Model(&model.User{}).Where("user_name=?",service.UserName).Count(&count)
	if count == 1 {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:"已经存在该用户了",
		}
	}
	user = model.User{
		UserName: service.UserName,
		Status:   model.Active,
	}
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.Avatar = "http://q1.qlogo.cn/g?b=qq&nk=294350394&s=640"
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}