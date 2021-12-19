package api

import (
	"chat/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func UserRegister(c *gin.Context) {
	var userRegisterService service.UserRegisterService //相当于创建了一个UserRegisterService对象，调用这个对象中的Register方法。
	if err := c.ShouldBind(&userRegisterService); err == nil {
		res := userRegisterService.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		logging.Info(err)
	}
}
