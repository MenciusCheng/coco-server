package api

import (
	"coco-server/model/common/response"
	"coco-server/util/log"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.POST("/user/login", UserApi.Login)
	routerGroup.POST("/user/logout", UserApi.Logout)
	routerGroup.GET("/user/info", UserApi.Info)
}

type userApi struct{}

var UserApi = new(userApi)

func (a *userApi) Login(c *gin.Context) {
	ctx := c.Request.Context()

	log.Info(ctx, "login start")
	res := map[string]interface{}{
		"token": "123456",
	}
	response.OkWithData(res, c)
	return
}

func (a *userApi) Logout(c *gin.Context) {
	res := map[string]interface{}{
		"token": "123456",
	}
	response.OkWithData(res, c)
	return
}

func (a *userApi) Info(c *gin.Context) {
	res := map[string]interface{}{
		"name":   "weiweicat",
		"avatar": "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
	}
	response.OkWithData(res, c)
	return
}
