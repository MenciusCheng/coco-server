package api

import (
	"coco-server/model"
	"coco-server/model/common/response"
	"coco-server/service"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.POST("/genStreamConf/query", GenStreamConfApi.Query)
	routerGroup.POST("/genStreamConf/create", GenStreamConfApi.Create)
	routerGroup.POST("/genStreamConf/update", GenStreamConfApi.Update)
	routerGroup.POST("/genStreamConf/delete", GenStreamConfApi.Delete)
}

type genStreamConfApi struct{}

var GenStreamConfApi = new(genStreamConfApi)

func (a *genStreamConfApi) Query(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.GenStreamConfQueryReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.GenStreamConfService.Query(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *genStreamConfApi) Create(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.GenStreamConfCreateReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.GenStreamConfService.Create(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *genStreamConfApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.GenStreamConfUpdateReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.GenStreamConfService.Update(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *genStreamConfApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.GenStreamConfDeleteReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.GenStreamConfService.Delete(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
	return
}
