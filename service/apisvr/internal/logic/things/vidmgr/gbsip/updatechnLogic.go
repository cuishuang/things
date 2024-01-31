package gbsip

import (
	"context"
	"encoding/json"
	"fmt"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/apisvr/internal/svc"
	"github.com/i-Things/things/service/apisvr/internal/types"
	"github.com/i-Things/things/service/vidsip/pb/sip"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatechnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatechnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatechnLogic {
	return &UpdatechnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatechnLogic) Updatechn(req *types.VidmgrSipUpdateChnReq) error {
	// todo: add your logic here and delete this line
	if req.ChannelID == "" {
		return errors.MediaSipUpdateError.AddDetail("通道IO为空")
	}

	vidReq := &sip.SipChnUpdateReq{
		ChannelID:  req.ChannelID,
		Memo:       req.Memo,
		StreamType: req.StreamType,
		Url:        req.URL,
	}
	jsonStr, _ := json.Marshal(req)
	fmt.Println("airgens Updatedev:", string(jsonStr))
	_, err := l.svcCtx.SipRpc.SipChannelUpdate(l.ctx, vidReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.ManageVidmgr req=%v err=%v", utils.FuncName(), req, er)
		return er
	}
	return nil
}
