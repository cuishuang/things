// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: dm.proto

package protocolmanage

import (
	"context"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionRespReq                     = dm.ActionRespReq
	ActionSendReq                     = dm.ActionSendReq
	ActionSendResp                    = dm.ActionSendResp
	CommonSchemaCreateReq             = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq              = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp             = dm.CommonSchemaIndexResp
	CommonSchemaInfo                  = dm.CommonSchemaInfo
	CommonSchemaUpdateReq             = dm.CommonSchemaUpdateReq
	CompareInt64                      = dm.CompareInt64
	CompareString                     = dm.CompareString
	CustomTopic                       = dm.CustomTopic
	DeviceCore                        = dm.DeviceCore
	DeviceCountInfo                   = dm.DeviceCountInfo
	DeviceCountReq                    = dm.DeviceCountReq
	DeviceCountResp                   = dm.DeviceCountResp
	DeviceError                       = dm.DeviceError
	DeviceGatewayBindDevice           = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq             = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp            = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq       = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiSaveReq         = dm.DeviceGatewayMultiSaveReq
	DeviceGatewaySign                 = dm.DeviceGatewaySign
	DeviceInfo                        = dm.DeviceInfo
	DeviceInfoBindReq                 = dm.DeviceInfoBindReq
	DeviceInfoCanBindReq              = dm.DeviceInfoCanBindReq
	DeviceInfoCount                   = dm.DeviceInfoCount
	DeviceInfoCountReq                = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq               = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp               = dm.DeviceInfoIndexResp
	DeviceInfoMultiBindReq            = dm.DeviceInfoMultiBindReq
	DeviceInfoMultiBindResp           = dm.DeviceInfoMultiBindResp
	DeviceInfoMultiUpdateReq          = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                 = dm.DeviceInfoReadReq
	DeviceModuleVersion               = dm.DeviceModuleVersion
	DeviceModuleVersionIndexReq       = dm.DeviceModuleVersionIndexReq
	DeviceModuleVersionIndexResp      = dm.DeviceModuleVersionIndexResp
	DeviceModuleVersionReadReq        = dm.DeviceModuleVersionReadReq
	DeviceMoveReq                     = dm.DeviceMoveReq
	DeviceOnlineMultiFix              = dm.DeviceOnlineMultiFix
	DeviceOnlineMultiFixReq           = dm.DeviceOnlineMultiFixReq
	DeviceProfile                     = dm.DeviceProfile
	DeviceProfileIndexReq             = dm.DeviceProfileIndexReq
	DeviceProfileIndexResp            = dm.DeviceProfileIndexResp
	DeviceProfileReadReq              = dm.DeviceProfileReadReq
	DeviceSchema                      = dm.DeviceSchema
	DeviceSchemaIndexReq              = dm.DeviceSchemaIndexReq
	DeviceSchemaIndexResp             = dm.DeviceSchemaIndexResp
	DeviceSchemaMultiCreateReq        = dm.DeviceSchemaMultiCreateReq
	DeviceSchemaMultiDeleteReq        = dm.DeviceSchemaMultiDeleteReq
	DeviceShareInfo                   = dm.DeviceShareInfo
	DeviceTransferReq                 = dm.DeviceTransferReq
	DeviceTypeCountReq                = dm.DeviceTypeCountReq
	DeviceTypeCountResp               = dm.DeviceTypeCountResp
	Empty                             = dm.Empty
	EventLogIndexReq                  = dm.EventLogIndexReq
	EventLogIndexResp                 = dm.EventLogIndexResp
	EventLogInfo                      = dm.EventLogInfo
	Firmware                          = dm.Firmware
	FirmwareFile                      = dm.FirmwareFile
	FirmwareInfo                      = dm.FirmwareInfo
	FirmwareInfoDeleteReq             = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp            = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq              = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp             = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq               = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp              = dm.FirmwareInfoReadResp
	FirmwareResp                      = dm.FirmwareResp
	GatewayCanBindIndexReq            = dm.GatewayCanBindIndexReq
	GatewayCanBindIndexResp           = dm.GatewayCanBindIndexResp
	GatewayGetFoundReq                = dm.GatewayGetFoundReq
	GatewayNotifyBindSendReq          = dm.GatewayNotifyBindSendReq
	GroupDeviceIndexReq               = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp              = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq         = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq           = dm.GroupDeviceMultiSaveReq
	GroupInfo                         = dm.GroupInfo
	GroupInfoCreateReq                = dm.GroupInfoCreateReq
	GroupInfoIndexReq                 = dm.GroupInfoIndexReq
	GroupInfoIndexResp                = dm.GroupInfoIndexResp
	GroupInfoUpdateReq                = dm.GroupInfoUpdateReq
	HubLogIndexReq                    = dm.HubLogIndexReq
	HubLogIndexResp                   = dm.HubLogIndexResp
	HubLogInfo                        = dm.HubLogInfo
	IDPath                            = dm.IDPath
	IDPathWithUpdate                  = dm.IDPathWithUpdate
	OtaFirmwareDeviceCancelReq        = dm.OtaFirmwareDeviceCancelReq
	OtaFirmwareDeviceConfirmReq       = dm.OtaFirmwareDeviceConfirmReq
	OtaFirmwareDeviceIndexReq         = dm.OtaFirmwareDeviceIndexReq
	OtaFirmwareDeviceIndexResp        = dm.OtaFirmwareDeviceIndexResp
	OtaFirmwareDeviceInfo             = dm.OtaFirmwareDeviceInfo
	OtaFirmwareDeviceRetryReq         = dm.OtaFirmwareDeviceRetryReq
	OtaFirmwareFile                   = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq           = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp          = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo               = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp               = dm.OtaFirmwareFileResp
	OtaFirmwareInfo                   = dm.OtaFirmwareInfo
	OtaFirmwareInfoCreateReq          = dm.OtaFirmwareInfoCreateReq
	OtaFirmwareInfoIndexReq           = dm.OtaFirmwareInfoIndexReq
	OtaFirmwareInfoIndexResp          = dm.OtaFirmwareInfoIndexResp
	OtaFirmwareInfoUpdateReq          = dm.OtaFirmwareInfoUpdateReq
	OtaFirmwareJobIndexReq            = dm.OtaFirmwareJobIndexReq
	OtaFirmwareJobIndexResp           = dm.OtaFirmwareJobIndexResp
	OtaFirmwareJobInfo                = dm.OtaFirmwareJobInfo
	OtaJobByDeviceIndexReq            = dm.OtaJobByDeviceIndexReq
	OtaJobDynamicInfo                 = dm.OtaJobDynamicInfo
	OtaJobStaticInfo                  = dm.OtaJobStaticInfo
	OtaModuleInfo                     = dm.OtaModuleInfo
	OtaModuleInfoIndexReq             = dm.OtaModuleInfoIndexReq
	OtaModuleInfoIndexResp            = dm.OtaModuleInfoIndexResp
	PageInfo                          = dm.PageInfo
	PageInfo_OrderBy                  = dm.PageInfo_OrderBy
	Point                             = dm.Point
	ProductCategory                   = dm.ProductCategory
	ProductCategoryIndexReq           = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp          = dm.ProductCategoryIndexResp
	ProductCategorySchemaIndexReq     = dm.ProductCategorySchemaIndexReq
	ProductCategorySchemaIndexResp    = dm.ProductCategorySchemaIndexResp
	ProductCategorySchemaMultiSaveReq = dm.ProductCategorySchemaMultiSaveReq
	ProductCustom                     = dm.ProductCustom
	ProductCustomReadReq              = dm.ProductCustomReadReq
	ProductCustomUi                   = dm.ProductCustomUi
	ProductInfo                       = dm.ProductInfo
	ProductInfoDeleteReq              = dm.ProductInfoDeleteReq
	ProductInfoIndexReq               = dm.ProductInfoIndexReq
	ProductInfoIndexResp              = dm.ProductInfoIndexResp
	ProductInfoReadReq                = dm.ProductInfoReadReq
	ProductInitReq                    = dm.ProductInitReq
	ProductRemoteConfig               = dm.ProductRemoteConfig
	ProductSchemaCreateReq            = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq            = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq             = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp            = dm.ProductSchemaIndexResp
	ProductSchemaInfo                 = dm.ProductSchemaInfo
	ProductSchemaMultiCreateReq       = dm.ProductSchemaMultiCreateReq
	ProductSchemaTslImportReq         = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq           = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp          = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq            = dm.ProductSchemaUpdateReq
	PropertyControlMultiSendReq       = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp      = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg            = dm.PropertyControlSendMsg
	PropertyControlSendReq            = dm.PropertyControlSendReq
	PropertyControlSendResp           = dm.PropertyControlSendResp
	PropertyGetReportSendReq          = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp         = dm.PropertyGetReportSendResp
	PropertyLogIndexReq               = dm.PropertyLogIndexReq
	PropertyLogIndexResp              = dm.PropertyLogIndexResp
	PropertyLogInfo                   = dm.PropertyLogInfo
	PropertyLogLatestIndexReq         = dm.PropertyLogLatestIndexReq
	ProtocolConfigField               = dm.ProtocolConfigField
	ProtocolConfigInfo                = dm.ProtocolConfigInfo
	ProtocolInfo                      = dm.ProtocolInfo
	ProtocolInfoIndexReq              = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp             = dm.ProtocolInfoIndexResp
	ProtocolService                   = dm.ProtocolService
	ProtocolServiceIndexReq           = dm.ProtocolServiceIndexReq
	ProtocolServiceIndexResp          = dm.ProtocolServiceIndexResp
	RemoteConfigCreateReq             = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq              = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp             = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq           = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp          = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq            = dm.RemoteConfigPushAllReq
	RespReadReq                       = dm.RespReadReq
	RootCheckReq                      = dm.RootCheckReq
	SdkLogIndexReq                    = dm.SdkLogIndexReq
	SdkLogIndexResp                   = dm.SdkLogIndexResp
	SdkLogInfo                        = dm.SdkLogInfo
	SendLogIndexReq                   = dm.SendLogIndexReq
	SendLogIndexResp                  = dm.SendLogIndexResp
	SendLogInfo                       = dm.SendLogInfo
	SendMsgReq                        = dm.SendMsgReq
	SendMsgResp                       = dm.SendMsgResp
	SendOption                        = dm.SendOption
	ShadowIndex                       = dm.ShadowIndex
	ShadowIndexResp                   = dm.ShadowIndexResp
	SharePerm                         = dm.SharePerm
	StatusLogIndexReq                 = dm.StatusLogIndexReq
	StatusLogIndexResp                = dm.StatusLogIndexResp
	StatusLogInfo                     = dm.StatusLogInfo
	TimeRange                         = dm.TimeRange
	UserDeviceCollectSave             = dm.UserDeviceCollectSave
	UserDeviceShareIndexReq           = dm.UserDeviceShareIndexReq
	UserDeviceShareIndexResp          = dm.UserDeviceShareIndexResp
	UserDeviceShareInfo               = dm.UserDeviceShareInfo
	UserDeviceShareMultiAcceptReq     = dm.UserDeviceShareMultiAcceptReq
	UserDeviceShareMultiDeleteReq     = dm.UserDeviceShareMultiDeleteReq
	UserDeviceShareMultiInfo          = dm.UserDeviceShareMultiInfo
	UserDeviceShareMultiToken         = dm.UserDeviceShareMultiToken
	UserDeviceShareReadReq            = dm.UserDeviceShareReadReq
	WithID                            = dm.WithID
	WithIDChildren                    = dm.WithIDChildren
	WithIDCode                        = dm.WithIDCode
	WithProfile                       = dm.WithProfile

	ProtocolManage interface {
		// 协议列表
		ProtocolInfoIndex(ctx context.Context, in *ProtocolInfoIndexReq, opts ...grpc.CallOption) (*ProtocolInfoIndexResp, error)
		// 协议详情
		ProtocolInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*ProtocolInfo, error)
		// 协议创建
		ProtocolInfoCreate(ctx context.Context, in *ProtocolInfo, opts ...grpc.CallOption) (*WithID, error)
		// 协议更新
		ProtocolInfoUpdate(ctx context.Context, in *ProtocolInfo, opts ...grpc.CallOption) (*Empty, error)
		// 协议删除
		ProtocolInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		// 更新服务状态,只给服务调用
		ProtocolServiceUpdate(ctx context.Context, in *ProtocolService, opts ...grpc.CallOption) (*Empty, error)
		ProtocolServiceDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		ProtocolServiceIndex(ctx context.Context, in *ProtocolServiceIndexReq, opts ...grpc.CallOption) (*ProtocolServiceIndexResp, error)
	}

	defaultProtocolManage struct {
		cli zrpc.Client
	}

	directProtocolManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.ProtocolManageServer
	}
)

func NewProtocolManage(cli zrpc.Client) ProtocolManage {
	return &defaultProtocolManage{
		cli: cli,
	}
}

func NewDirectProtocolManage(svcCtx *svc.ServiceContext, svr dm.ProtocolManageServer) ProtocolManage {
	return &directProtocolManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 协议列表
func (m *defaultProtocolManage) ProtocolInfoIndex(ctx context.Context, in *ProtocolInfoIndexReq, opts ...grpc.CallOption) (*ProtocolInfoIndexResp, error) {
	client := dm.NewProtocolManageClient(m.cli.Conn())
	return client.ProtocolInfoIndex(ctx, in, opts...)
}

// 协议列表
func (d *directProtocolManage) ProtocolInfoIndex(ctx context.Context, in *ProtocolInfoIndexReq, opts ...grpc.CallOption) (*ProtocolInfoIndexResp, error) {
	return d.svr.ProtocolInfoIndex(ctx, in)
}

// 协议详情
func (m *defaultProtocolManage) ProtocolInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*ProtocolInfo, error) {
	client := dm.NewProtocolManageClient(m.cli.Conn())
	return client.ProtocolInfoRead(ctx, in, opts...)
}

// 协议详情
func (d *directProtocolManage) ProtocolInfoRead(ctx context.Context, in *WithIDCode, opts ...grpc.CallOption) (*ProtocolInfo, error) {
	return d.svr.ProtocolInfoRead(ctx, in)
}

// 协议创建
func (m *defaultProtocolManage) ProtocolInfoCreate(ctx context.Context, in *ProtocolInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := dm.NewProtocolManageClient(m.cli.Conn())
	return client.ProtocolInfoCreate(ctx, in, opts...)
}

// 协议创建
func (d *directProtocolManage) ProtocolInfoCreate(ctx context.Context, in *ProtocolInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.ProtocolInfoCreate(ctx, in)
}

// 协议更新
func (m *defaultProtocolManage) ProtocolInfoUpdate(ctx context.Context, in *ProtocolInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProtocolManageClient(m.cli.Conn())
	return client.ProtocolInfoUpdate(ctx, in, opts...)
}

// 协议更新
func (d *directProtocolManage) ProtocolInfoUpdate(ctx context.Context, in *ProtocolInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProtocolInfoUpdate(ctx, in)
}

// 协议删除
func (m *defaultProtocolManage) ProtocolInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProtocolManageClient(m.cli.Conn())
	return client.ProtocolInfoDelete(ctx, in, opts...)
}

// 协议删除
func (d *directProtocolManage) ProtocolInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProtocolInfoDelete(ctx, in)
}

// 更新服务状态,只给服务调用
func (m *defaultProtocolManage) ProtocolServiceUpdate(ctx context.Context, in *ProtocolService, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProtocolManageClient(m.cli.Conn())
	return client.ProtocolServiceUpdate(ctx, in, opts...)
}

// 更新服务状态,只给服务调用
func (d *directProtocolManage) ProtocolServiceUpdate(ctx context.Context, in *ProtocolService, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProtocolServiceUpdate(ctx, in)
}

func (m *defaultProtocolManage) ProtocolServiceDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewProtocolManageClient(m.cli.Conn())
	return client.ProtocolServiceDelete(ctx, in, opts...)
}

func (d *directProtocolManage) ProtocolServiceDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ProtocolServiceDelete(ctx, in)
}

func (m *defaultProtocolManage) ProtocolServiceIndex(ctx context.Context, in *ProtocolServiceIndexReq, opts ...grpc.CallOption) (*ProtocolServiceIndexResp, error) {
	client := dm.NewProtocolManageClient(m.cli.Conn())
	return client.ProtocolServiceIndex(ctx, in, opts...)
}

func (d *directProtocolManage) ProtocolServiceIndex(ctx context.Context, in *ProtocolServiceIndexReq, opts ...grpc.CallOption) (*ProtocolServiceIndexResp, error) {
	return d.svr.ProtocolServiceIndex(ctx, in)
}
