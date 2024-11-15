package api

import (
	"coco-server/model"
	"coco-server/model/common/response"
	"coco-server/service"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.POST("/bookmark/query", BookmarkApi.Query)
	routerGroup.POST("/bookmark/create", BookmarkApi.Create)
	routerGroup.POST("/bookmark/update", BookmarkApi.Update)
	routerGroup.POST("/bookmark/delete", BookmarkApi.Delete)
	routerGroup.POST("/bookmark/tree", BookmarkApi.Tree)
}

type bookmarkApi struct{}

var BookmarkApi = new(bookmarkApi)

func (a *bookmarkApi) Query(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.BookmarkQueryReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.BookmarkService.Query(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *bookmarkApi) Create(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.BookmarkCreateReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.BookmarkService.Create(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *bookmarkApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.BookmarkUpdateReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.BookmarkService.Update(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *bookmarkApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.BookmarkDeleteReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.BookmarkService.Delete(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
	return
}

func (a *bookmarkApi) Tree(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.BookmarkTreeReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.BookmarkService.GetBookmarkTree(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}
