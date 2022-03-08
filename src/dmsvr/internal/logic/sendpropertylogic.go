package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-things/things/shared/errors"
	"github.com/go-things/things/src/dmsvr/internal/domain/deviceTemplate"
	"github.com/go-things/things/src/dmsvr/internal/repo/mysql"
	"github.com/hashicorp/go-uuid"
	"time"

	"github.com/go-things/things/src/dmsvr/dm"
	"github.com/go-things/things/src/dmsvr/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendPropertyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	pt       *mysql.ProductTemplate
	template *deviceTemplate.Template
}

func NewSendPropertyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendPropertyLogic {
	return &SendPropertyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendPropertyLogic) initMsg(productID string) error {
	var err error
	l.pt, err = l.svcCtx.ProductTemplate.FindOne(productID)
	if err != nil {
		return err
	}
	l.template, err = deviceTemplate.NewTemplate([]byte(l.pt.Template))
	if err != nil {
		return err
	}
	return nil
}

func (l *SendPropertyLogic) SendProperty(in *dm.SendPropertyReq) (*dm.SendPropertyResp, error) {
	l.Infof("SendProperty|req=%+v", in)
	err := l.initMsg(in.ProductID)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	err = json.Unmarshal([]byte(in.Data), &param)
	if err != nil {
		return nil, errors.Parameter.AddDetail(
			"SendProperty|Data not right:", in.Data)
	}
	uuid, err := uuid.GenerateUUID()
	if err != nil {
		l.Errorf("SendProperty|GenerateUUID err:%v", err)
		return nil, errors.System.AddDetail(err)
	}
	req := deviceTemplate.DeviceReq{
		Method:      deviceTemplate.CONTROL,
		ClientToken: uuid,
		//ClientToken:"de65377c-4041-565d-0b5e-67b664a06be8",//这个是测试代码
		Timestamp: time.Now().UnixMilli(),
		Params:    param}
	l.template.VerifyReqParam(req, deviceTemplate.ACTION_INPUT)
	pubTopic := fmt.Sprintf("$thing/down/property/%s/%s", in.ProductID, in.DeviceName)
	subTopic := fmt.Sprintf("$thing/up/property/%s/%s", in.ProductID, in.DeviceName)

	resp, err := l.svcCtx.DevClient.DeviceReq(l.ctx, req, pubTopic, subTopic)
	if err != nil {
		return nil, err
	}

	return &dm.SendPropertyResp{
		ClientToken: resp.ClientToken,
		Status:      resp.Status,
		Code:        resp.Code,
	}, nil

	return &dm.SendPropertyResp{}, nil
}
