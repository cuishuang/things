package common

import (
	"context"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/i-Things/things/service/apisvr/internal/svc"
	"github.com/i-Things/things/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteLogic) Delete(req *types.CommonSchemaDeleteReq) error {
	dmReq := &dm.WithID{
		Id: req.ID,
	}
	_, err := l.svcCtx.SchemaM.CommonSchemaDelete(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.CommonSchemaDelete req=%v err=%+v", utils.FuncName(), req, er)
		return er
	}
	return nil
}
