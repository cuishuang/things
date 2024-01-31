package firmware

import (
	"context"

	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/i-Things/things/service/apisvr/internal/svc"
	"github.com/i-Things/things/service/apisvr/internal/types"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.OtaFirmwareInfoUpdateReq) error {
	dmReq := &dm.FirmwareInfo{
		FirmwareID: req.FirmwareID,
		Name:       req.Name,
	}
	if req.Desc != nil {
		dmReq.Desc = &wrappers.StringValue{
			Value: *req.Desc,
		}
	}
	if req.ExtData != nil {
		dmReq.ExtData = &wrappers.StringValue{
			Value: *req.ExtData,
		}
	}
	_, err := l.svcCtx.FirmwareM.FirmwareInfoUpdate(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s|rpc.firmwaremanage|req=%v|err=%+v", utils.FuncName(), utils.Fmt(req), er)
		return er
	}
	return nil
}
