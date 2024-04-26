package api

import (
	"coco-server/model/common/response"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.GET("/ping", Ping)
}

func Ping(c *gin.Context) {
	response.Ok(c)
}
