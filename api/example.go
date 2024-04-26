package api

import (
	"coco-server/dao"
	"coco-server/model"
	"coco-server/model/common/response"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.POST("/example/list", ExampleApi.GetList)
	routerGroup.POST("/example/create", ExampleApi.Create)
	routerGroup.POST("/example/update", ExampleApi.Update)
	routerGroup.POST("/example/delete", ExampleApi.Delete)
}

type exampleApi struct{}

var ExampleApi = new(exampleApi)

func (a *exampleApi) GetList(c *gin.Context) {
	ctx := c.Request.Context()
	req := &dao.QueryExampleFilter{}
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, count := dao.ExampleDB.QueryExampleRecords(ctx, req)
	res := map[string]interface{}{
		"list":  list,
		"count": count,
	}
	response.OkWithData(res, c)
	return
}

func (a *exampleApi) Create(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.Example)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dao.ExampleDB.InsertExample(ctx, req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
	return
}

func (a *exampleApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.Example)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dao.ExampleDB.UpdateExample(ctx, req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
	return
}

func (a *exampleApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.Example)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dao.ExampleDB.DeleteExample(ctx, req.Id); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Ok(c)
	return
}
