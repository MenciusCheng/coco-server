package api

import (
	"coco-server/model"
	"coco-server/model/common/response"
	"coco-server/service"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.POST("/actJsonConf/query", ActJsonConfApi.Query)
	routerGroup.POST("/actJsonConf/create", ActJsonConfApi.Create)
	routerGroup.POST("/actJsonConf/update", ActJsonConfApi.Update)
	routerGroup.POST("/actJsonConf/delete", ActJsonConfApi.Delete)
}

type actJsonConfApi struct{}

var ActJsonConfApi = new(actJsonConfApi)

func (a *actJsonConfApi) Query(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.ActJsonConfQueryReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.ActJsonConfService.Query(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *actJsonConfApi) Create(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.ActJsonConfCreateReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.ActJsonConfService.Create(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *actJsonConfApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.ActJsonConfUpdateReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.ActJsonConfService.Update(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *actJsonConfApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.ActJsonConfDeleteReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.ActJsonConfService.Delete(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
	return
}
