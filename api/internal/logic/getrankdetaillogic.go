package logic

import (
	"context"

	"api/internal/svc"
	"api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRankDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRankDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRankDetailLogic {
	return GetRankDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRankDetailLogic) GetRankDetail(req types.InfoReq) (*types.InfoRes, error) {
	detail, _ := l.svcCtx.ItemModel.FindOne(req.Id)
	list, _ := l.svcCtx.ItemDetailModel.GetAll(req.Id)
	return &types.InfoRes{detail, list}, nil
}
