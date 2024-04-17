// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package devicemsg

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionRespReq                       = dm.ActionRespReq
	ActionSendReq                       = dm.ActionSendReq
	ActionSendResp                      = dm.ActionSendResp
	CommonSchemaCreateReq               = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq                = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp               = dm.CommonSchemaIndexResp
	CommonSchemaInfo                    = dm.CommonSchemaInfo
	CommonSchemaUpdateReq               = dm.CommonSchemaUpdateReq
	CustomTopic                         = dm.CustomTopic
	DeviceCore                          = dm.DeviceCore
	DeviceCountInfo                     = dm.DeviceCountInfo
	DeviceCountReq                      = dm.DeviceCountReq
	DeviceCountResp                     = dm.DeviceCountResp
	DeviceGatewayBindDevice             = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq               = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp              = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq         = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq         = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign                   = dm.DeviceGatewaySign
	DeviceInfo                          = dm.DeviceInfo
	DeviceInfoBindReq                   = dm.DeviceInfoBindReq
	DeviceInfoCount                     = dm.DeviceInfoCount
	DeviceInfoCountReq                  = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq                 = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                  = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp                 = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq            = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                   = dm.DeviceInfoReadReq
	DeviceProfile                       = dm.DeviceProfile
	DeviceProfileIndexReq               = dm.DeviceProfileIndexReq
	DeviceProfileIndexResp              = dm.DeviceProfileIndexResp
	DeviceProfileReadReq                = dm.DeviceProfileReadReq
	DeviceTypeCountReq                  = dm.DeviceTypeCountReq
	DeviceTypeCountResp                 = dm.DeviceTypeCountResp
	Empty                               = dm.Empty
	EventLogIndexReq                    = dm.EventLogIndexReq
	EventLogIndexResp                   = dm.EventLogIndexResp
	EventLogInfo                        = dm.EventLogInfo
	Firmware                            = dm.Firmware
	FirmwareFile                        = dm.FirmwareFile
	FirmwareInfo                        = dm.FirmwareInfo
	FirmwareInfoDeleteReq               = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp              = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq                = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp               = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq                 = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp                = dm.FirmwareInfoReadResp
	FirmwareResp                        = dm.FirmwareResp
	GroupDeviceIndexReq                 = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp                = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq           = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq             = dm.GroupDeviceMultiSaveReq
	GroupInfo                           = dm.GroupInfo
	GroupInfoCreateReq                  = dm.GroupInfoCreateReq
	GroupInfoIndexReq                   = dm.GroupInfoIndexReq
	GroupInfoIndexResp                  = dm.GroupInfoIndexResp
	GroupInfoUpdateReq                  = dm.GroupInfoUpdateReq
	HubLogIndexReq                      = dm.HubLogIndexReq
	HubLogIndexResp                     = dm.HubLogIndexResp
	HubLogInfo                          = dm.HubLogInfo
	OtaFirmwareDeviceCancelReq          = dm.OtaFirmwareDeviceCancelReq
	OtaFirmwareDeviceIndexReq           = dm.OtaFirmwareDeviceIndexReq
	OtaFirmwareDeviceIndexResp          = dm.OtaFirmwareDeviceIndexResp
	OtaFirmwareDeviceInfo               = dm.OtaFirmwareDeviceInfo
	OtaFirmwareDeviceRetryReq           = dm.OtaFirmwareDeviceRetryReq
	OtaFirmwareFile                     = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq             = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp            = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo                 = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                  = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp                 = dm.OtaFirmwareFileResp
	OtaFirmwareInfo                     = dm.OtaFirmwareInfo
	OtaFirmwareInfoCreateReq            = dm.OtaFirmwareInfoCreateReq
	OtaFirmwareInfoIndexReq             = dm.OtaFirmwareInfoIndexReq
	OtaFirmwareInfoIndexResp            = dm.OtaFirmwareInfoIndexResp
	OtaFirmwareInfoUpdateReq            = dm.OtaFirmwareInfoUpdateReq
	OtaFirmwareJobIndexReq              = dm.OtaFirmwareJobIndexReq
	OtaFirmwareJobIndexResp             = dm.OtaFirmwareJobIndexResp
	OtaFirmwareJobInfo                  = dm.OtaFirmwareJobInfo
	OtaJobByDeviceIndexReq              = dm.OtaJobByDeviceIndexReq
	OtaJobDynamicInfo                   = dm.OtaJobDynamicInfo
	OtaJobStaticInfo                    = dm.OtaJobStaticInfo
	OtaModuleInfo                       = dm.OtaModuleInfo
	OtaModuleInfoIndexReq               = dm.OtaModuleInfoIndexReq
	OtaModuleInfoIndexResp              = dm.OtaModuleInfoIndexResp
	OtaPromptIndexReq                   = dm.OtaPromptIndexReq
	OtaPromptIndexResp                  = dm.OtaPromptIndexResp
	PageInfo                            = dm.PageInfo
	PageInfo_OrderBy                    = dm.PageInfo_OrderBy
	Point                               = dm.Point
	ProductCategory                     = dm.ProductCategory
	ProductCategoryIndexReq             = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp            = dm.ProductCategoryIndexResp
	ProductCategoryReadReq              = dm.ProductCategoryReadReq
	ProductCategorySchemaIndexResp      = dm.ProductCategorySchemaIndexResp
	ProductCategorySchemaMultiUpdateReq = dm.ProductCategorySchemaMultiUpdateReq
	ProductCustom                       = dm.ProductCustom
	ProductCustomReadReq                = dm.ProductCustomReadReq
	ProductInfo                         = dm.ProductInfo
	ProductInfoDeleteReq                = dm.ProductInfoDeleteReq
	ProductInfoIndexReq                 = dm.ProductInfoIndexReq
	ProductInfoIndexResp                = dm.ProductInfoIndexResp
	ProductInfoReadReq                  = dm.ProductInfoReadReq
	ProductInitReq                      = dm.ProductInitReq
	ProductRemoteConfig                 = dm.ProductRemoteConfig
	ProductSchemaCreateReq              = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq              = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq               = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp              = dm.ProductSchemaIndexResp
	ProductSchemaInfo                   = dm.ProductSchemaInfo
	ProductSchemaMultiCreateReq         = dm.ProductSchemaMultiCreateReq
	ProductSchemaTslImportReq           = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq             = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp            = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq              = dm.ProductSchemaUpdateReq
	PropertyControlMultiSendReq         = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp        = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg              = dm.PropertyControlSendMsg
	PropertyControlSendReq              = dm.PropertyControlSendReq
	PropertyControlSendResp             = dm.PropertyControlSendResp
	PropertyGetReportSendReq            = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp           = dm.PropertyGetReportSendResp
	PropertyLogIndexReq                 = dm.PropertyLogIndexReq
	PropertyLogIndexResp                = dm.PropertyLogIndexResp
	PropertyLogInfo                     = dm.PropertyLogInfo
	PropertyLogLatestIndexReq           = dm.PropertyLogLatestIndexReq
	ProtocolConfigField                 = dm.ProtocolConfigField
	ProtocolConfigInfo                  = dm.ProtocolConfigInfo
	ProtocolInfo                        = dm.ProtocolInfo
	ProtocolInfoIndexReq                = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp               = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq               = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq                = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp               = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq             = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp            = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq              = dm.RemoteConfigPushAllReq
	RespReadReq                         = dm.RespReadReq
	RootCheckReq                        = dm.RootCheckReq
	SdkLogIndexReq                      = dm.SdkLogIndexReq
	SdkLogIndexResp                     = dm.SdkLogIndexResp
	SdkLogInfo                          = dm.SdkLogInfo
	SendLogIndexReq                     = dm.SendLogIndexReq
	SendLogIndexResp                    = dm.SendLogIndexResp
	SendLogInfo                         = dm.SendLogInfo
	SendMsgReq                          = dm.SendMsgReq
	SendMsgResp                         = dm.SendMsgResp
	SendOption                          = dm.SendOption
	ShadowIndex                         = dm.ShadowIndex
	ShadowIndexResp                     = dm.ShadowIndexResp
	StatusLogIndexReq                   = dm.StatusLogIndexReq
	StatusLogIndexResp                  = dm.StatusLogIndexResp
	StatusLogInfo                       = dm.StatusLogInfo
	TimeRange                           = dm.TimeRange
	UserDeviceCollectSave               = dm.UserDeviceCollectSave
	UserDeviceShareIndexReq             = dm.UserDeviceShareIndexReq
	UserDeviceShareIndexResp            = dm.UserDeviceShareIndexResp
	UserDeviceShareInfo                 = dm.UserDeviceShareInfo
	UserDeviceShareMultiDeleteReq       = dm.UserDeviceShareMultiDeleteReq
	UserDeviceShareReadReq              = dm.UserDeviceShareReadReq
	WithID                              = dm.WithID
	WithIDCode                          = dm.WithIDCode

	DeviceMsg interface {
		// 获取设备sdk调试日志
		SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error)
		// 获取设备调试信息记录登入登出,操作
		HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error)
		SendLogIndex(ctx context.Context, in *SendLogIndexReq, opts ...grpc.CallOption) (*SendLogIndexResp, error)
		StatusLogIndex(ctx context.Context, in *StatusLogIndexReq, opts ...grpc.CallOption) (*StatusLogIndexResp, error)
		// 获取设备数据信息
		PropertyLogLatestIndex(ctx context.Context, in *PropertyLogLatestIndexReq, opts ...grpc.CallOption) (*PropertyLogIndexResp, error)
		// 获取设备数据信息
		PropertyLogIndex(ctx context.Context, in *PropertyLogIndexReq, opts ...grpc.CallOption) (*PropertyLogIndexResp, error)
		// 获取设备数据信息
		EventLogIndex(ctx context.Context, in *EventLogIndexReq, opts ...grpc.CallOption) (*EventLogIndexResp, error)
		// 获取设备影子列表
		ShadowIndex(ctx context.Context, in *PropertyLogLatestIndexReq, opts ...grpc.CallOption) (*ShadowIndexResp, error)
		// 主动触发单个设备ota升级推送
		OtaPromptIndex(ctx context.Context, in *OtaPromptIndexReq, opts ...grpc.CallOption) (*OtaPromptIndexResp, error)
	}

	defaultDeviceMsg struct {
		cli zrpc.Client
	}

	directDeviceMsg struct {
		svcCtx *svc.ServiceContext
		svr    dm.DeviceMsgServer
	}
)

func NewDeviceMsg(cli zrpc.Client) DeviceMsg {
	return &defaultDeviceMsg{
		cli: cli,
	}
}

func NewDirectDeviceMsg(svcCtx *svc.ServiceContext, svr dm.DeviceMsgServer) DeviceMsg {
	return &directDeviceMsg{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 获取设备sdk调试日志
func (m *defaultDeviceMsg) SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.SdkLogIndex(ctx, in, opts...)
}

// 获取设备sdk调试日志
func (d *directDeviceMsg) SdkLogIndex(ctx context.Context, in *SdkLogIndexReq, opts ...grpc.CallOption) (*SdkLogIndexResp, error) {
	return d.svr.SdkLogIndex(ctx, in)
}

// 获取设备调试信息记录登入登出,操作
func (m *defaultDeviceMsg) HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.HubLogIndex(ctx, in, opts...)
}

// 获取设备调试信息记录登入登出,操作
func (d *directDeviceMsg) HubLogIndex(ctx context.Context, in *HubLogIndexReq, opts ...grpc.CallOption) (*HubLogIndexResp, error) {
	return d.svr.HubLogIndex(ctx, in)
}

func (m *defaultDeviceMsg) SendLogIndex(ctx context.Context, in *SendLogIndexReq, opts ...grpc.CallOption) (*SendLogIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.SendLogIndex(ctx, in, opts...)
}

func (d *directDeviceMsg) SendLogIndex(ctx context.Context, in *SendLogIndexReq, opts ...grpc.CallOption) (*SendLogIndexResp, error) {
	return d.svr.SendLogIndex(ctx, in)
}

func (m *defaultDeviceMsg) StatusLogIndex(ctx context.Context, in *StatusLogIndexReq, opts ...grpc.CallOption) (*StatusLogIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.StatusLogIndex(ctx, in, opts...)
}

func (d *directDeviceMsg) StatusLogIndex(ctx context.Context, in *StatusLogIndexReq, opts ...grpc.CallOption) (*StatusLogIndexResp, error) {
	return d.svr.StatusLogIndex(ctx, in)
}

// 获取设备数据信息
func (m *defaultDeviceMsg) PropertyLogLatestIndex(ctx context.Context, in *PropertyLogLatestIndexReq, opts ...grpc.CallOption) (*PropertyLogIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.PropertyLogLatestIndex(ctx, in, opts...)
}

// 获取设备数据信息
func (d *directDeviceMsg) PropertyLogLatestIndex(ctx context.Context, in *PropertyLogLatestIndexReq, opts ...grpc.CallOption) (*PropertyLogIndexResp, error) {
	return d.svr.PropertyLogLatestIndex(ctx, in)
}

// 获取设备数据信息
func (m *defaultDeviceMsg) PropertyLogIndex(ctx context.Context, in *PropertyLogIndexReq, opts ...grpc.CallOption) (*PropertyLogIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.PropertyLogIndex(ctx, in, opts...)
}

// 获取设备数据信息
func (d *directDeviceMsg) PropertyLogIndex(ctx context.Context, in *PropertyLogIndexReq, opts ...grpc.CallOption) (*PropertyLogIndexResp, error) {
	return d.svr.PropertyLogIndex(ctx, in)
}

// 获取设备数据信息
func (m *defaultDeviceMsg) EventLogIndex(ctx context.Context, in *EventLogIndexReq, opts ...grpc.CallOption) (*EventLogIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.EventLogIndex(ctx, in, opts...)
}

// 获取设备数据信息
func (d *directDeviceMsg) EventLogIndex(ctx context.Context, in *EventLogIndexReq, opts ...grpc.CallOption) (*EventLogIndexResp, error) {
	return d.svr.EventLogIndex(ctx, in)
}

// 获取设备影子列表
func (m *defaultDeviceMsg) ShadowIndex(ctx context.Context, in *PropertyLogLatestIndexReq, opts ...grpc.CallOption) (*ShadowIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.ShadowIndex(ctx, in, opts...)
}

// 获取设备影子列表
func (d *directDeviceMsg) ShadowIndex(ctx context.Context, in *PropertyLogLatestIndexReq, opts ...grpc.CallOption) (*ShadowIndexResp, error) {
	return d.svr.ShadowIndex(ctx, in)
}

// 主动触发单个设备ota升级推送
func (m *defaultDeviceMsg) OtaPromptIndex(ctx context.Context, in *OtaPromptIndexReq, opts ...grpc.CallOption) (*OtaPromptIndexResp, error) {
	client := dm.NewDeviceMsgClient(m.cli.Conn())
	return client.OtaPromptIndex(ctx, in, opts...)
}

// 主动触发单个设备ota升级推送
func (d *directDeviceMsg) OtaPromptIndex(ctx context.Context, in *OtaPromptIndexReq, opts ...grpc.CallOption) (*OtaPromptIndexResp, error) {
	return d.svr.OtaPromptIndex(ctx, in)
}
