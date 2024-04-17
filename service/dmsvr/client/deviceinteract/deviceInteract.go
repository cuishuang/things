// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package deviceinteract

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

	DeviceInteract interface {
		// 调用设备行为
		ActionSend(ctx context.Context, in *ActionSendReq, opts ...grpc.CallOption) (*ActionSendResp, error)
		// 获取异步调用设备行为的结果
		ActionRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*ActionSendResp, error)
		// 回复调用设备行为
		ActionResp(ctx context.Context, in *ActionRespReq, opts ...grpc.CallOption) (*Empty, error)
		// 请求设备获取设备最新属性
		PropertyGetReportSend(ctx context.Context, in *PropertyGetReportSendReq, opts ...grpc.CallOption) (*PropertyGetReportSendResp, error)
		// 调用设备属性
		PropertyControlSend(ctx context.Context, in *PropertyControlSendReq, opts ...grpc.CallOption) (*PropertyControlSendResp, error)
		// 批量调用设备属性
		PropertyControlMultiSend(ctx context.Context, in *PropertyControlMultiSendReq, opts ...grpc.CallOption) (*PropertyControlMultiSendResp, error)
		// 获取异步调用设备属性的结果
		PropertyControlRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*PropertyControlSendResp, error)
		// 发送消息给设备 -- 调试使用
		SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
	}

	defaultDeviceInteract struct {
		cli zrpc.Client
	}

	directDeviceInteract struct {
		svcCtx *svc.ServiceContext
		svr    dm.DeviceInteractServer
	}
)

func NewDeviceInteract(cli zrpc.Client) DeviceInteract {
	return &defaultDeviceInteract{
		cli: cli,
	}
}

func NewDirectDeviceInteract(svcCtx *svc.ServiceContext, svr dm.DeviceInteractServer) DeviceInteract {
	return &directDeviceInteract{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 调用设备行为
func (m *defaultDeviceInteract) ActionSend(ctx context.Context, in *ActionSendReq, opts ...grpc.CallOption) (*ActionSendResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.ActionSend(ctx, in, opts...)
}

// 调用设备行为
func (d *directDeviceInteract) ActionSend(ctx context.Context, in *ActionSendReq, opts ...grpc.CallOption) (*ActionSendResp, error) {
	return d.svr.ActionSend(ctx, in)
}

// 获取异步调用设备行为的结果
func (m *defaultDeviceInteract) ActionRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*ActionSendResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.ActionRead(ctx, in, opts...)
}

// 获取异步调用设备行为的结果
func (d *directDeviceInteract) ActionRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*ActionSendResp, error) {
	return d.svr.ActionRead(ctx, in)
}

// 回复调用设备行为
func (m *defaultDeviceInteract) ActionResp(ctx context.Context, in *ActionRespReq, opts ...grpc.CallOption) (*Empty, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.ActionResp(ctx, in, opts...)
}

// 回复调用设备行为
func (d *directDeviceInteract) ActionResp(ctx context.Context, in *ActionRespReq, opts ...grpc.CallOption) (*Empty, error) {
	return d.svr.ActionResp(ctx, in)
}

// 请求设备获取设备最新属性
func (m *defaultDeviceInteract) PropertyGetReportSend(ctx context.Context, in *PropertyGetReportSendReq, opts ...grpc.CallOption) (*PropertyGetReportSendResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.PropertyGetReportSend(ctx, in, opts...)
}

// 请求设备获取设备最新属性
func (d *directDeviceInteract) PropertyGetReportSend(ctx context.Context, in *PropertyGetReportSendReq, opts ...grpc.CallOption) (*PropertyGetReportSendResp, error) {
	return d.svr.PropertyGetReportSend(ctx, in)
}

// 调用设备属性
func (m *defaultDeviceInteract) PropertyControlSend(ctx context.Context, in *PropertyControlSendReq, opts ...grpc.CallOption) (*PropertyControlSendResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.PropertyControlSend(ctx, in, opts...)
}

// 调用设备属性
func (d *directDeviceInteract) PropertyControlSend(ctx context.Context, in *PropertyControlSendReq, opts ...grpc.CallOption) (*PropertyControlSendResp, error) {
	return d.svr.PropertyControlSend(ctx, in)
}

// 批量调用设备属性
func (m *defaultDeviceInteract) PropertyControlMultiSend(ctx context.Context, in *PropertyControlMultiSendReq, opts ...grpc.CallOption) (*PropertyControlMultiSendResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.PropertyControlMultiSend(ctx, in, opts...)
}

// 批量调用设备属性
func (d *directDeviceInteract) PropertyControlMultiSend(ctx context.Context, in *PropertyControlMultiSendReq, opts ...grpc.CallOption) (*PropertyControlMultiSendResp, error) {
	return d.svr.PropertyControlMultiSend(ctx, in)
}

// 获取异步调用设备属性的结果
func (m *defaultDeviceInteract) PropertyControlRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*PropertyControlSendResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.PropertyControlRead(ctx, in, opts...)
}

// 获取异步调用设备属性的结果
func (d *directDeviceInteract) PropertyControlRead(ctx context.Context, in *RespReadReq, opts ...grpc.CallOption) (*PropertyControlSendResp, error) {
	return d.svr.PropertyControlRead(ctx, in)
}

// 发送消息给设备 -- 调试使用
func (m *defaultDeviceInteract) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	client := dm.NewDeviceInteractClient(m.cli.Conn())
	return client.SendMsg(ctx, in, opts...)
}

// 发送消息给设备 -- 调试使用
func (d *directDeviceInteract) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	return d.svr.SendMsg(ctx, in)
}
