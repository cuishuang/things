// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/logic/tenantmanage"
	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"
)

type TenantManageServer struct {
	svcCtx *svc.ServiceContext
	sys.UnimplementedTenantManageServer
}

func NewTenantManageServer(svcCtx *svc.ServiceContext) *TenantManageServer {
	return &TenantManageServer{
		svcCtx: svcCtx,
	}
}

// 新增区域
func (s *TenantManageServer) TenantInfoCreate(ctx context.Context, in *sys.TenantInfoCreateReq) (*sys.Response, error) {
	l := tenantmanagelogic.NewTenantInfoCreateLogic(ctx, s.svcCtx)
	return l.TenantInfoCreate(in)
}

// 更新区域
func (s *TenantManageServer) TenantInfoUpdate(ctx context.Context, in *sys.TenantInfo) (*sys.Response, error) {
	l := tenantmanagelogic.NewTenantInfoUpdateLogic(ctx, s.svcCtx)
	return l.TenantInfoUpdate(in)
}

// 删除区域
func (s *TenantManageServer) TenantInfoDelete(ctx context.Context, in *sys.ReqWithIDCode) (*sys.Response, error) {
	l := tenantmanagelogic.NewTenantInfoDeleteLogic(ctx, s.svcCtx)
	return l.TenantInfoDelete(in)
}

// 获取租户信息详情
func (s *TenantManageServer) TenantInfoRead(ctx context.Context, in *sys.ReqWithIDCode) (*sys.TenantInfo, error) {
	l := tenantmanagelogic.NewTenantInfoReadLogic(ctx, s.svcCtx)
	return l.TenantInfoRead(in)
}

// 获取租户信息列表
func (s *TenantManageServer) TenantInfoIndex(ctx context.Context, in *sys.TenantInfoIndexReq) (*sys.TenantInfoIndexResp, error) {
	l := tenantmanagelogic.NewTenantInfoIndexLogic(ctx, s.svcCtx)
	return l.TenantInfoIndex(in)
}

func (s *TenantManageServer) TenantAppIndex(ctx context.Context, in *sys.TenantAppIndexReq) (*sys.TenantAppIndexResp, error) {
	l := tenantmanagelogic.NewTenantAppIndexLogic(ctx, s.svcCtx)
	return l.TenantAppIndex(in)
}

func (s *TenantManageServer) TenantAppMultiUpdate(ctx context.Context, in *sys.TenantAppMultiUpdateReq) (*sys.Response, error) {
	l := tenantmanagelogic.NewTenantAppMultiUpdateLogic(ctx, s.svcCtx)
	return l.TenantAppMultiUpdate(in)
}
