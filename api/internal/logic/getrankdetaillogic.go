package logic

import (
	"context"
	"fmt"

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
	res, _ := l.svcCtx.ItemModel.FindOne(req.Id)
	fmt.Println(res)
	return &types.InfoRes{Data: res}, nil
}
