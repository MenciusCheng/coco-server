package service

import (
	"coco-server/dao"
	"coco-server/model"
	"coco-server/model/common/request"
	"coco-server/util/log"
	"context"
	"go.uber.org/zap"
)

type actJsonConfService struct{}

var ActJsonConfService = new(actJsonConfService)

func (s *actJsonConfService) Query(ctx context.Context, req *model.ActJsonConfQueryReq) (*model.ActJsonConfQueryRes, error) {
	res := new(model.ActJsonConfQueryRes)

	filter := &dao.QueryActJsonConfFilter{
		Name:   req.Name,
		Extend: req.Extend,
	}
	filter.Offset, filter.Limit = request.FormatPage(req.Page, req.Size)
	list, total, err := dao.ActJsonConfDB.Query(ctx, filter)
	if err != nil {
		log.Error(ctx, "Query failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	res.List = list
	res.Total = total
	return res, nil
}

func (s *actJsonConfService) Create(ctx context.Context, req *model.ActJsonConfCreateReq) (*model.ActJsonConfCreateRes, error) {
	res := new(model.ActJsonConfCreateRes)

	v := model.ActJsonConf{
		Name:   req.Name,
		Extend: req.Extend,
	}
	id, err := dao.ActJsonConfDB.Create(ctx, &v)
	if err != nil {
		log.Error(ctx, "Create failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	res.Id = id
	return res, nil
}

func (s *actJsonConfService) Update(ctx context.Context, req *model.ActJsonConfUpdateReq) (*model.ActJsonConfUpdateRes, error) {
	res := new(model.ActJsonConfUpdateRes)

	v := model.ActJsonConf{
		Id:     req.Id,
		Name:   req.Name,
		Extend: req.Extend,
	}
	err := dao.ActJsonConfDB.Update(ctx, &v)
	if err != nil {
		log.Error(ctx, "Update failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}

func (s *actJsonConfService) Delete(ctx context.Context, req *model.ActJsonConfDeleteReq) (*model.ActJsonConfDeleteRes, error) {
	res := new(model.ActJsonConfDeleteRes)

	err := dao.ActJsonConfDB.Delete(ctx, req.Id)
	if err != nil {
		log.Error(ctx, "Delete failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}
