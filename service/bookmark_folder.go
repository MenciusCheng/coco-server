package service

import (
	"coco-server/dao"
	"coco-server/model"
	"coco-server/model/common/request"
	"coco-server/util/log"
	"context"
	"go.uber.org/zap"
)

type bookmarkFolderService struct{}

var BookmarkFolderService = new(bookmarkFolderService)

func (s *bookmarkFolderService) Query(ctx context.Context, req *model.BookmarkFolderQueryReq) (*model.BookmarkFolderQueryRes, error) {
	res := new(model.BookmarkFolderQueryRes)

	filter := &dao.QueryBookmarkFolderFilter{
		Name:     req.Name,
		ParentId: req.ParentId,
	}
	filter.Offset, filter.Limit = request.FormatPage(req.Page, req.Size)
	list, total, err := dao.BookmarkFolderDB.Query(ctx, filter)
	if err != nil {
		log.Error(ctx, "Query failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	res.List = list
	res.Total = total
	return res, nil
}

func (s *bookmarkFolderService) Create(ctx context.Context, req *model.BookmarkFolderCreateReq) (*model.BookmarkFolderCreateRes, error) {
	res := new(model.BookmarkFolderCreateRes)

	v := model.BookmarkFolder{
		Name:     req.Name,
		ParentId: req.ParentId,
	}
	id, err := dao.BookmarkFolderDB.Create(ctx, &v)
	if err != nil {
		log.Error(ctx, "Create failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	res.Id = id
	return res, nil
}

func (s *bookmarkFolderService) Update(ctx context.Context, req *model.BookmarkFolderUpdateReq) (*model.BookmarkFolderUpdateRes, error) {
	res := new(model.BookmarkFolderUpdateRes)

	v := model.BookmarkFolder{
		Id:       req.Id,
		Name:     req.Name,
		ParentId: req.ParentId,
	}
	err := dao.BookmarkFolderDB.Update(ctx, &v)
	if err != nil {
		log.Error(ctx, "Update failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}

func (s *bookmarkFolderService) Delete(ctx context.Context, req *model.BookmarkFolderDeleteReq) (*model.BookmarkFolderDeleteRes, error) {
	res := new(model.BookmarkFolderDeleteRes)

	err := dao.BookmarkFolderDB.Delete(ctx, req.Id)
	if err != nil {
		log.Error(ctx, "Delete failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}
