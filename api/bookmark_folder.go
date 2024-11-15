package api

import (
	"coco-server/model"
	"coco-server/model/common/response"
	"coco-server/service"
	"github.com/gin-gonic/gin"
)

func init() {
	routerGroup := GetRouterGroup()
	routerGroup.POST("/bookmarkFolder/query", BookmarkFolderApi.Query)
	routerGroup.POST("/bookmarkFolder/create", BookmarkFolderApi.Create)
	routerGroup.POST("/bookmarkFolder/update", BookmarkFolderApi.Update)
	routerGroup.POST("/bookmarkFolder/delete", BookmarkFolderApi.Delete)
}

type bookmarkFolderApi struct{}

var BookmarkFolderApi = new(bookmarkFolderApi)

func (a *bookmarkFolderApi) Query(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.BookmarkFolderQueryReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.BookmarkFolderService.Query(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *bookmarkFolderApi) Create(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.BookmarkFolderCreateReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.BookmarkFolderService.Create(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *bookmarkFolderApi) Update(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.BookmarkFolderUpdateReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.BookmarkFolderService.Update(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}

func (a *bookmarkFolderApi) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(model.BookmarkFolderDeleteReq)
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := service.BookmarkFolderService.Delete(ctx, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
	return
}
