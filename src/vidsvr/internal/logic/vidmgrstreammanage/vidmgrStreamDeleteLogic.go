package vidmgrstreammanagelogic

import (
	"context"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/vidsvr/internal/repo/relationDB"

	"github.com/i-Things/things/src/vidsvr/internal/svc"
	"github.com/i-Things/things/src/vidsvr/pb/vid"

	"github.com/zeromicro/go-zero/core/logx"
)

type VidmgrStreamDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	PiDB *relationDB.VidmgrStreamRepo
}

func NewVidmgrStreamDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VidmgrStreamDeleteLogic {
	return &VidmgrStreamDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		PiDB:   relationDB.NewVidmgrStreamRepo(ctx),
	}
}

// 删除流
func (l *VidmgrStreamDeleteLogic) VidmgrStreamDelete(in *vid.VidmgrStreamDeleteReq) (*vid.Response, error) {
	// todo: add your logic here and delete this line
	err := l.PiDB.DeleteByFilter(l.ctx, relationDB.VidmgrStreamFilter{
		StreamIDs: []int64{in.StreamID},
	})
	if err != nil {
		l.Errorf("%s.Delete err=%v", utils.FuncName(), utils.Fmt(err))
		return nil, err
	}
	return &vid.Response{}, nil
}
