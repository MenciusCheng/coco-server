package service

import (
	"coco-server/dao"
	"coco-server/model"
	"coco-server/model/common/request"
	"coco-server/util/log"
	"context"
	"go.uber.org/zap"
)

type genStreamDetailTmplService struct{}

var GenStreamDetailTmplService = new(genStreamDetailTmplService)

func (s *genStreamDetailTmplService) Query(ctx context.Context, req *model.GenStreamDetailTmplQueryReq) (*model.GenStreamDetailTmplQueryRes, error) {
	res := new(model.GenStreamDetailTmplQueryRes)

	filter := &dao.QueryGenStreamDetailTmplFilter{
		Name:   req.Name,
		Extend: req.Extend,
	}
	filter.Offset, filter.Limit = request.FormatPage(req.Page, req.Size)
	list, total, err := dao.GenStreamDetailTmplDB.Query(ctx, filter)
	if err != nil {
		log.Error(ctx, "Query failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	res.List = list
	res.Total = total
	return res, nil
}

func (s *genStreamDetailTmplService) Create(ctx context.Context, req *model.GenStreamDetailTmplCreateReq) (*model.GenStreamDetailTmplCreateRes, error) {
	res := new(model.GenStreamDetailTmplCreateRes)

	v := model.GenStreamDetailTmpl{
		Name:   req.Name,
		Extend: req.Extend,
	}
	id, err := dao.GenStreamDetailTmplDB.Create(ctx, &v)
	if err != nil {
		log.Error(ctx, "Create failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	res.Id = id
	return res, nil
}

func (s *genStreamDetailTmplService) Update(ctx context.Context, req *model.GenStreamDetailTmplUpdateReq) (*model.GenStreamDetailTmplUpdateRes, error) {
	res := new(model.GenStreamDetailTmplUpdateRes)

	v := model.GenStreamDetailTmpl{
		Id:     req.Id,
		Name:   req.Name,
		Extend: req.Extend,
	}
	err := dao.GenStreamDetailTmplDB.Update(ctx, &v)
	if err != nil {
		log.Error(ctx, "Update failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}

func (s *genStreamDetailTmplService) Delete(ctx context.Context, req *model.GenStreamDetailTmplDeleteReq) (*model.GenStreamDetailTmplDeleteRes, error) {
	res := new(model.GenStreamDetailTmplDeleteRes)

	err := dao.GenStreamDetailTmplDB.Delete(ctx, req.Id)
	if err != nil {
		log.Error(ctx, "Delete failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}
