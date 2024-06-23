package service

import (
	"coco-server/dao"
	"coco-server/model"
	"coco-server/model/common/request"
	"coco-server/util/log"
	"context"
	"go.uber.org/zap"
)

type genStreamConfService struct{}

var GenStreamConfService = new(genStreamConfService)

func (s *genStreamConfService) Query(ctx context.Context, req *model.GenStreamConfQueryReq) (*model.GenStreamConfQueryRes, error) {
	res := new(model.GenStreamConfQueryRes)

	filter := &dao.QueryGenStreamConfFilter{
		ConfNameLike: req.ConfNameLike,
	}
	filter.Offset, filter.Limit = request.FormatPage(req.Page, req.Size)
	list, total, err := dao.GenStreamConfDB.Query(ctx, filter)
	if err != nil {
		log.Error(ctx, "Query failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	res.List = list
	res.Total = total
	return res, nil
}

func (s *genStreamConfService) Create(ctx context.Context, req *model.GenStreamConfCreateReq) (*model.GenStreamConfCreateRes, error) {
	res := new(model.GenStreamConfCreateRes)

	v := model.GenStreamConf{
		ConfName: req.ConfName,
		Extend:   req.Extend,
	}
	id, err := dao.GenStreamConfDB.Create(ctx, &v)
	if err != nil {
		log.Error(ctx, "Create failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	res.Id = id
	return res, nil
}

func (s *genStreamConfService) Update(ctx context.Context, req *model.GenStreamConfUpdateReq) (*model.GenStreamConfUpdateRes, error) {
	res := new(model.GenStreamConfUpdateRes)

	v := model.GenStreamConf{
		Id:       req.Id,
		ConfName: req.ConfName,
		Extend:   req.Extend,
	}
	err := dao.GenStreamConfDB.Update(ctx, &v)
	if err != nil {
		log.Error(ctx, "Update failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}

func (s *genStreamConfService) Delete(ctx context.Context, req *model.GenStreamConfDeleteReq) (*model.GenStreamConfDeleteRes, error) {
	res := new(model.GenStreamConfDeleteRes)

	err := dao.GenStreamConfDB.Delete(ctx, req.Id)
	if err != nil {
		log.Error(ctx, "Delete failed", zap.Error(err), zap.Any("req", req))
		return nil, err
	}

	return res, nil
}
