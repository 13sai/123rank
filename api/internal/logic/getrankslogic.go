package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRanksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRanksLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRanksLogic {
	return GetRanksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRanksLogic) GetRanks(req types.ListReq) (*types.ListRes, error) {
	res, _ := l.svcCtx.ItemModel.GetList(req.Page, req.Limit)
	return &types.ListRes{Data: res, Page: req.Page, Limit: req.Limit}, nil
}
