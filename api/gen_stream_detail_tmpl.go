package api

import (
	"coco-server/model"
	"coco-server/model/common/response"
	"coco-server/service"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.POST("/genStreamDetailTmpl/query", GenStreamDetailTmplApi.Query)
	routerGroup.POST("/genStreamDetailTmpl/create", GenStreamDetailTmplApi.Create)
	routerGroup.POST("/genStreamDetailTmpl/update", GenStreamDetailTmplApi.Update)
	routerGroup.POST("/genStreamDetailTmpl/delete", GenStreamDetailTmplApi.Delete)
}

type genStreamDetailTmplApi struct{}

var GenStreamDetailTmplApi = new(genStreamDetailTmplApi)

func (a *genStreamDetailTmplApi) Query(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.GenStreamDetailTmplQueryReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.GenStreamDetailTmplService.Query(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *genStreamDetailTmplApi) Create(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.GenStreamDetailTmplCreateReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.GenStreamDetailTmplService.Create(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *genStreamDetailTmplApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.GenStreamDetailTmplUpdateReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.GenStreamDetailTmplService.Update(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *genStreamDetailTmplApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.GenStreamDetailTmplDeleteReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.GenStreamDetailTmplService.Delete(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
	return
}
