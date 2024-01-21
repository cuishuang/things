// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/logic/commonschema"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

type CommonSchemaServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedCommonSchemaServer
}

func NewCommonSchemaServer(svcCtx *svc.ServiceContext) *CommonSchemaServer {
	return &CommonSchemaServer{
		svcCtx: svcCtx,
	}
}

// 更新产品物模型
func (s *CommonSchemaServer) CommonSchemaUpdate(ctx context.Context, in *dm.CommonSchemaUpdateReq) (*dm.Response, error) {
	l := commonschemalogic.NewCommonSchemaUpdateLogic(ctx, s.svcCtx)
	return l.CommonSchemaUpdate(in)
}

// 新增产品
func (s *CommonSchemaServer) CommonSchemaCreate(ctx context.Context, in *dm.CommonSchemaCreateReq) (*dm.Response, error) {
	l := commonschemalogic.NewCommonSchemaCreateLogic(ctx, s.svcCtx)
	return l.CommonSchemaCreate(in)
}

// 删除产品
func (s *CommonSchemaServer) CommonSchemaDelete(ctx context.Context, in *dm.WithID) (*dm.Response, error) {
	l := commonschemalogic.NewCommonSchemaDeleteLogic(ctx, s.svcCtx)
	return l.CommonSchemaDelete(in)
}

// 获取产品信息列表
func (s *CommonSchemaServer) CommonSchemaIndex(ctx context.Context, in *dm.CommonSchemaIndexReq) (*dm.CommonSchemaIndexResp, error) {
	l := commonschemalogic.NewCommonSchemaIndexLogic(ctx, s.svcCtx)
	return l.CommonSchemaIndex(in)
}
