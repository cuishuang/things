package auth5

import (
	"context"
	"gitee.com/unitedrhino/share/ctxs"
	"gitee.com/unitedrhino/things/service/apisvr/internal/logic/things/device"
	"gitee.com/unitedrhino/things/service/apisvr/internal/svc"
	"gitee.com/unitedrhino/things/service/apisvr/internal/types"
	"gitee.com/unitedrhino/things/service/dgsvr/pb/dg"
	"gitee.com/unitedrhino/things/share/devices"
	"github.com/maypok86/otter"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AccessLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAccessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AccessLogic {
	return &AccessLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var accessCache otter.Cache[types.DeviceAuth5AccessReq, struct{}]

func init() {
	cache, err := otter.MustBuilder[types.DeviceAuth5AccessReq, struct{}](10_000).
		CollectStats().
		Cost(func(key types.DeviceAuth5AccessReq, value struct{}) uint32 {
			return 1
		}).
		WithTTL(time.Minute * 10).
		Build()
	logx.Must(err)
	accessCache = cache
}

func (l *AccessLogic) Access(req *types.DeviceAuth5AccessReq) (resp *types.DeviceAuth5AccessResp, err error) {
	_, ok := accessCache.Get(*req)
	if ok {
		return &types.DeviceAuth5AccessResp{Result: "allow"}, nil
	}
	access := req.Action
	//如果是
	switch req.Action {
	case "subscribe":
		access = devices.Sub
	case "publish":
		access = devices.Pub
	}
	l.ctx = ctxs.WithRoot(l.ctx)
	_, err = l.svcCtx.DeviceA.AccessAuth(l.ctx, &dg.AccessAuthReq{
		Username: req.Username,
		Topic:    req.Topic,
		ClientID: req.ClientID,
		Access:   access,
		Ip:       req.Ip,
	})
	if err == nil {
		accessCache.Set(*req, struct{}{})
		return &types.DeviceAuth5AccessResp{Result: "allow"}, nil
	}
	err = device.ThirdProtoAccessAuth(l.ctx, l.svcCtx, &types.DeviceAuthAccessReq{
		Username: req.Username,
		Topic:    req.Topic,
		ClientID: req.ClientID,
		Ip:       req.Ip,
	}, access)
	if err != nil {
		return &types.DeviceAuth5AccessResp{Result: "deny"}, nil
	}
	accessCache.Set(*req, struct{}{})
	return &types.DeviceAuth5AccessResp{Result: "allow"}, nil
}
