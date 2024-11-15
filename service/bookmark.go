package service

import (
	"coco-server/dao"
	"coco-server/model"
	"coco-server/model/common/request"
	"coco-server/util/log"
	"context"
	"go.uber.org/zap"
)

type bookmarkService struct{}

var BookmarkService = new(bookmarkService)

func (s *bookmarkService) Query(ctx context.Context, req *model.BookmarkQueryReq) (*model.BookmarkQueryRes, error) {
	res := new(model.BookmarkQueryRes)

	filter := &dao.QueryBookmarkFilter{
		Name:     req.Name,
		FolderId: req.FolderId,
	}
	filter.Offset, filter.Limit = request.FormatPage(req.Page, req.Size)
	list, total, err := dao.BookmarkDB.Query(ctx, filter)
	if err != nil {
		log.Error(ctx, "Query failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	res.List = list
	res.Total = total
	return res, nil
}

func (s *bookmarkService) Create(ctx context.Context, req *model.BookmarkCreateReq) (*model.BookmarkCreateRes, error) {
	res := new(model.BookmarkCreateRes)

	v := model.Bookmark{
		Name:     req.Name,
		Url:      req.Url,
		Icon:     req.Icon,
		Remark:   req.Remark,
		FolderId: req.FolderId,
	}
	id, err := dao.BookmarkDB.Create(ctx, &v)
	if err != nil {
		log.Error(ctx, "Create failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	res.Id = id
	return res, nil
}

func (s *bookmarkService) Update(ctx context.Context, req *model.BookmarkUpdateReq) (*model.BookmarkUpdateRes, error) {
	res := new(model.BookmarkUpdateRes)

	v := model.Bookmark{
		Id:       req.Id,
		Name:     req.Name,
		Url:      req.Url,
		Icon:     req.Icon,
		Remark:   req.Remark,
		FolderId: req.FolderId,
	}
	err := dao.BookmarkDB.Update(ctx, &v)
	if err != nil {
		log.Error(ctx, "Update failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}

func (s *bookmarkService) Delete(ctx context.Context, req *model.BookmarkDeleteReq) (*model.BookmarkDeleteRes, error) {
	res := new(model.BookmarkDeleteRes)

	err := dao.BookmarkDB.Delete(ctx, req.Id)
	if err != nil {
		log.Error(ctx, "Delete failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}

func (s *bookmarkService) GetBookmarkTree(ctx context.Context, req *model.BookmarkTreeReq) (*model.BookmarkTreeRes, error) {
	folders, _, err := dao.BookmarkFolderDB.Query(ctx, &dao.QueryBookmarkFolderFilter{Limit: 100000})
	if err != nil {
		return nil, err
	}

	bookmarks, _, err := dao.BookmarkDB.Query(ctx, &dao.QueryBookmarkFilter{Limit: 100000})
	if err != nil {
		return nil, err
	}

	res := s.buildBookmarkTree(ctx, folders, bookmarks)
	return res, nil
}

func (s *bookmarkService) buildBookmarkTree(ctx context.Context, folders []*model.BookmarkFolder, bookmarks []*model.Bookmark) *model.BookmarkTreeRes {
	res := new(model.BookmarkTreeRes)

	folderMap := make(map[int64]*model.BookmarkTreeData)
	for _, folder := range folders {
		folderMap[folder.Id] = &model.BookmarkTreeData{
			FolderId: folder.Id,
			Type:     model.TypeBookmarkTreeFolder,
			Name:     folder.Name,
		}
	}

	rootBookmarks := make([]*model.BookmarkTreeData, 0)
	for _, bookmark := range bookmarks {
		info := &model.BookmarkTreeData{
			BookmarkId: bookmark.Id,
			Type:       model.TypeBookmarkTreeBookmark,
			Name:       bookmark.Name,
			Url:        bookmark.Url,
			Icon:       bookmark.Icon,
			Remark:     bookmark.Remark,
		}
		if folder, ok := folderMap[bookmark.FolderId]; ok {
			folder.List = append(folder.List, info)
		} else {
			rootBookmarks = append(rootBookmarks, info)
		}
	}

	rootFolders := make([]*model.BookmarkTreeData, 0)
	for _, folder := range folders {
		if folder.ParentId == 0 {
			rootFolders = append(rootFolders, folderMap[folder.Id])
		} else {
			if parentFolder, ok := folderMap[folder.ParentId]; ok {
				parentFolder.List = append(parentFolder.List, folderMap[folder.Id])
			}
		}
	}

	res.List = append(res.List, rootFolders...)
	res.List = append(res.List, rootBookmarks...)
	return res
}
