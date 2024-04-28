package api

import (
	"coco-server/dao"
	"coco-server/model"
	"coco-server/model/common/request"
	"coco-server/model/common/response"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.POST("/dataConvert/list", DataConvertApi.GetList)
	routerGroup.POST("/dataConvert/create", DataConvertApi.Create)
	routerGroup.POST("/dataConvert/update", DataConvertApi.Update)
	routerGroup.POST("/dataConvert/delete", DataConvertApi.Delete)
}

type dataConvertApi struct{}

var DataConvertApi = new(dataConvertApi)

func (a *dataConvertApi) GetList(c *gin.Context) {
	ctx := c.Request.Context()
	req := &model.DataConvertGetListReq{}
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	filter := &dao.QueryDataConvertFilter{
		ConfName: req.ConfName,
	}
	filter.Offset, filter.Limit = request.FormatPage(req.Page, req.Size)
	list, count := dao.DataConvertDB.QueryDataConvertRecords(ctx, filter)
	res := map[string]interface{}{
		"records": list,
		"total":   count,
	}
	response.OkWithData(res, c)
	return
}

func (a *dataConvertApi) Create(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.DataConvert)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dao.DataConvertDB.InsertDataConvert(ctx, req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
	return
}

func (a *dataConvertApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.DataConvert)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dao.DataConvertDB.UpdateDataConvert(ctx, req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
	return
}

func (a *dataConvertApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.DataConvert)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dao.DataConvertDB.DeleteDataConvert(ctx, req.Id); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
	return
}
