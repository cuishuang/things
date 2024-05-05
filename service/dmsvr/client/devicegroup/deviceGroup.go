// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package devicegroup

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

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
	CustomTopic                       = dm.CustomTopic
	DeviceCore                        = dm.DeviceCore
	DeviceCountInfo                   = dm.DeviceCountInfo
	DeviceCountReq                    = dm.DeviceCountReq
	DeviceCountResp                   = dm.DeviceCountResp
	DeviceGatewayBindDevice           = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq             = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp            = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq       = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiSaveReq         = dm.DeviceGatewayMultiSaveReq
	DeviceGatewaySign                 = dm.DeviceGatewaySign
	DeviceInfo                        = dm.DeviceInfo
	DeviceInfoBindReq                 = dm.DeviceInfoBindReq
	DeviceInfoCount                   = dm.DeviceInfoCount
	DeviceInfoCountReq                = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq               = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp               = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq          = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                 = dm.DeviceInfoReadReq
	DeviceProfile                     = dm.DeviceProfile
	DeviceProfileIndexReq             = dm.DeviceProfileIndexReq
	DeviceProfileIndexResp            = dm.DeviceProfileIndexResp
	DeviceProfileReadReq              = dm.DeviceProfileReadReq
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
	ManufacturerInfo                  = dm.ManufacturerInfo
	OtaFirmwareDeviceCancelReq        = dm.OtaFirmwareDeviceCancelReq
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
	OtaPromptIndexReq                 = dm.OtaPromptIndexReq
	OtaPromptIndexResp                = dm.OtaPromptIndexResp
	PageInfo                          = dm.PageInfo
	PageInfo_OrderBy                  = dm.PageInfo_OrderBy
	Point                             = dm.Point
	ProductCategory                   = dm.ProductCategory
	ProductCategoryIndexReq           = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp          = dm.ProductCategoryIndexResp
	ProductCategoryReadReq            = dm.ProductCategoryReadReq
	ProductCategorySchemaIndexReq     = dm.ProductCategorySchemaIndexReq
	ProductCategorySchemaIndexResp    = dm.ProductCategorySchemaIndexResp
	ProductCategorySchemaMultiSaveReq = dm.ProductCategorySchemaMultiSaveReq
	ProductCustom                     = dm.ProductCustom
	ProductCustomReadReq              = dm.ProductCustomReadReq
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
	UserDeviceShareMultiDeleteReq     = dm.UserDeviceShareMultiDeleteReq
	UserDeviceShareReadReq            = dm.UserDeviceShareReadReq
	WithID                            = dm.WithID
	WithIDCode                        = dm.WithIDCode

	DeviceGroup interface {
		// 创建分组
		GroupInfoCreate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*WithID, error)
		// 获取分组信息列表
		GroupInfoIndex(ctx context.Context, in *GroupInfoIndexReq, opts ...grpc.CallOption) (*GroupInfoIndexResp, error)
		// 获取分组信息详情
		GroupInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*GroupInfo, error)
		// 更新分组
		GroupInfoUpdate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error)
		// 删除分组
		GroupInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error)
		// 创建分组设备
		GroupDeviceMultiCreate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
		// 更新分组设备
		GroupDeviceMultiUpdate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error)
		// 获取分组设备信息列表
		GroupDeviceIndex(ctx context.Context, in *GroupDeviceIndexReq, opts ...grpc.CallOption) (*GroupDeviceIndexResp, error)
		// 删除分组设备
		GroupDeviceMultiDelete(ctx context.Context, in *GroupDeviceMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultDeviceGroup struct {
		cli zrpc.Client
	}

	directDeviceGroup struct {
		svcCtx *svc.ServiceContext
		svr    dm.DeviceGroupServer
	}
)

func NewDeviceGroup(cli zrpc.Client) DeviceGroup {
	return &defaultDeviceGroup{
		cli: cli,
	}
}

func NewDirectDeviceGroup(svcCtx *svc.ServiceContext, svr dm.DeviceGroupServer) DeviceGroup {
	return &directDeviceGroup{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 创建分组
func (m *defaultDeviceGroup) GroupInfoCreate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*WithID, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoCreate(ctx, in, opts...)
}

// 创建分组
func (d *directDeviceGroup) GroupInfoCreate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*WithID, error) {
	return d.svr.GroupInfoCreate(ctx, in)
}

// 获取分组信息列表
func (m *defaultDeviceGroup) GroupInfoIndex(ctx context.Context, in *GroupInfoIndexReq, opts ...grpc.CallOption) (*GroupInfoIndexResp, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoIndex(ctx, in, opts...)
}

// 获取分组信息列表
func (d *directDeviceGroup) GroupInfoIndex(ctx context.Context, in *GroupInfoIndexReq, opts ...grpc.CallOption) (*GroupInfoIndexResp, error) {
	return d.svr.GroupInfoIndex(ctx, in)
}

// 获取分组信息详情
func (m *defaultDeviceGroup) GroupInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*GroupInfo, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoRead(ctx, in, opts...)
}

// 获取分组信息详情
func (d *directDeviceGroup) GroupInfoRead(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*GroupInfo, error) {
	return d.svr.GroupInfoRead(ctx, in)
}

// 更新分组
func (m *defaultDeviceGroup) GroupInfoUpdate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoUpdate(ctx, in, opts...)
}

// 更新分组
func (d *directDeviceGroup) GroupInfoUpdate(ctx context.Context, in *GroupInfo, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupInfoUpdate(ctx, in)
}

// 删除分组
func (m *defaultDeviceGroup) GroupInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupInfoDelete(ctx, in, opts...)
}

// 删除分组
func (d *directDeviceGroup) GroupInfoDelete(ctx context.Context, in *WithID, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupInfoDelete(ctx, in)
}

// 创建分组设备
func (m *defaultDeviceGroup) GroupDeviceMultiCreate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupDeviceMultiCreate(ctx, in, opts...)
}

// 创建分组设备
func (d *directDeviceGroup) GroupDeviceMultiCreate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupDeviceMultiCreate(ctx, in)
}

// 更新分组设备
func (m *defaultDeviceGroup) GroupDeviceMultiUpdate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupDeviceMultiUpdate(ctx, in, opts...)
}

// 更新分组设备
func (d *directDeviceGroup) GroupDeviceMultiUpdate(ctx context.Context, in *GroupDeviceMultiSaveReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupDeviceMultiUpdate(ctx, in)
}

// 获取分组设备信息列表
func (m *defaultDeviceGroup) GroupDeviceIndex(ctx context.Context, in *GroupDeviceIndexReq, opts ...grpc.CallOption) (*GroupDeviceIndexResp, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupDeviceIndex(ctx, in, opts...)
}

// 获取分组设备信息列表
func (d *directDeviceGroup) GroupDeviceIndex(ctx context.Context, in *GroupDeviceIndexReq, opts ...grpc.CallOption) (*GroupDeviceIndexResp, error) {
	return d.svr.GroupDeviceIndex(ctx, in)
}

// 删除分组设备
func (m *defaultDeviceGroup) GroupDeviceMultiDelete(ctx context.Context, in *GroupDeviceMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceGroupClient(m.cli.Conn())
	return client.GroupDeviceMultiDelete(ctx, in, opts...)
}

// 删除分组设备
func (d *directDeviceGroup) GroupDeviceMultiDelete(ctx context.Context, in *GroupDeviceMultiDeleteReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.GroupDeviceMultiDelete(ctx, in)
}
