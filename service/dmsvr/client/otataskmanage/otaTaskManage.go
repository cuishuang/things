// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package otataskmanage

import (
	"context"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ActionRespReq                      = dm.ActionRespReq
	ActionSendReq                      = dm.ActionSendReq
	ActionSendResp                     = dm.ActionSendResp
	CommonSchemaCreateReq              = dm.CommonSchemaCreateReq
	CommonSchemaIndexReq               = dm.CommonSchemaIndexReq
	CommonSchemaIndexResp              = dm.CommonSchemaIndexResp
	CommonSchemaInfo                   = dm.CommonSchemaInfo
	CommonSchemaUpdateReq              = dm.CommonSchemaUpdateReq
	CustomTopic                        = dm.CustomTopic
	DeviceCore                         = dm.DeviceCore
	DeviceCountInfo                    = dm.DeviceCountInfo
	DeviceCountReq                     = dm.DeviceCountReq
	DeviceCountResp                    = dm.DeviceCountResp
	DeviceGatewayBindDevice            = dm.DeviceGatewayBindDevice
	DeviceGatewayIndexReq              = dm.DeviceGatewayIndexReq
	DeviceGatewayIndexResp             = dm.DeviceGatewayIndexResp
	DeviceGatewayMultiCreateReq        = dm.DeviceGatewayMultiCreateReq
	DeviceGatewayMultiDeleteReq        = dm.DeviceGatewayMultiDeleteReq
	DeviceGatewaySign                  = dm.DeviceGatewaySign
	DeviceInfo                         = dm.DeviceInfo
	DeviceInfoCount                    = dm.DeviceInfoCount
	DeviceInfoCountReq                 = dm.DeviceInfoCountReq
	DeviceInfoDeleteReq                = dm.DeviceInfoDeleteReq
	DeviceInfoIndexReq                 = dm.DeviceInfoIndexReq
	DeviceInfoIndexResp                = dm.DeviceInfoIndexResp
	DeviceInfoMultiUpdateReq           = dm.DeviceInfoMultiUpdateReq
	DeviceInfoReadReq                  = dm.DeviceInfoReadReq
	DeviceTypeCountReq                 = dm.DeviceTypeCountReq
	DeviceTypeCountResp                = dm.DeviceTypeCountResp
	DynamicUpgradeJobReq               = dm.DynamicUpgradeJobReq
	Empty                              = dm.Empty
	EventIndex                         = dm.EventIndex
	EventIndexResp                     = dm.EventIndexResp
	EventLogIndexReq                   = dm.EventLogIndexReq
	Firmware                           = dm.Firmware
	FirmwareFile                       = dm.FirmwareFile
	FirmwareInfo                       = dm.FirmwareInfo
	FirmwareInfoDeleteReq              = dm.FirmwareInfoDeleteReq
	FirmwareInfoDeleteResp             = dm.FirmwareInfoDeleteResp
	FirmwareInfoIndexReq               = dm.FirmwareInfoIndexReq
	FirmwareInfoIndexResp              = dm.FirmwareInfoIndexResp
	FirmwareInfoReadReq                = dm.FirmwareInfoReadReq
	FirmwareInfoReadResp               = dm.FirmwareInfoReadResp
	FirmwareResp                       = dm.FirmwareResp
	GroupDeviceIndexReq                = dm.GroupDeviceIndexReq
	GroupDeviceIndexResp               = dm.GroupDeviceIndexResp
	GroupDeviceMultiDeleteReq          = dm.GroupDeviceMultiDeleteReq
	GroupDeviceMultiSaveReq            = dm.GroupDeviceMultiSaveReq
	GroupInfo                          = dm.GroupInfo
	GroupInfoCreateReq                 = dm.GroupInfoCreateReq
	GroupInfoIndexReq                  = dm.GroupInfoIndexReq
	GroupInfoIndexResp                 = dm.GroupInfoIndexResp
	GroupInfoUpdateReq                 = dm.GroupInfoUpdateReq
	HubLogIndex                        = dm.HubLogIndex
	HubLogIndexReq                     = dm.HubLogIndexReq
	HubLogIndexResp                    = dm.HubLogIndexResp
	JobReq                             = dm.JobReq
	OTAModuleDeleteReq                 = dm.OTAModuleDeleteReq
	OTAModuleDetail                    = dm.OTAModuleDetail
	OTAModuleIndexReq                  = dm.OTAModuleIndexReq
	OTAModuleIndexResp                 = dm.OTAModuleIndexResp
	OTAModuleReq                       = dm.OTAModuleReq
	OTAModuleVersionsIndexResp         = dm.OTAModuleVersionsIndexResp
	OTATaskByDeviceCancelReq           = dm.OTATaskByDeviceCancelReq
	OTATaskByDeviceNameReq             = dm.OTATaskByDeviceNameReq
	OTATaskByJobCancelReq              = dm.OTATaskByJobCancelReq
	OTATaskByJobIndexReq               = dm.OTATaskByJobIndexReq
	OTATaskConfirmReq                  = dm.OTATaskConfirmReq
	OTATaskReUpgradeReq                = dm.OTATaskReUpgradeReq
	OTAUnfinishedTaskByDeviceIndexReq  = dm.OTAUnfinishedTaskByDeviceIndexReq
	OTAUnfinishedTaskByDeviceIndexResp = dm.OTAUnfinishedTaskByDeviceIndexResp
	OtaCommonResp                      = dm.OtaCommonResp
	OtaFirmwareCreateReq               = dm.OtaFirmwareCreateReq
	OtaFirmwareDeleteReq               = dm.OtaFirmwareDeleteReq
	OtaFirmwareDeviceInfoReq           = dm.OtaFirmwareDeviceInfoReq
	OtaFirmwareDeviceInfoResp          = dm.OtaFirmwareDeviceInfoResp
	OtaFirmwareFile                    = dm.OtaFirmwareFile
	OtaFirmwareFileIndexReq            = dm.OtaFirmwareFileIndexReq
	OtaFirmwareFileIndexResp           = dm.OtaFirmwareFileIndexResp
	OtaFirmwareFileInfo                = dm.OtaFirmwareFileInfo
	OtaFirmwareFileReq                 = dm.OtaFirmwareFileReq
	OtaFirmwareFileResp                = dm.OtaFirmwareFileResp
	OtaFirmwareIndexReq                = dm.OtaFirmwareIndexReq
	OtaFirmwareIndexResp               = dm.OtaFirmwareIndexResp
	OtaFirmwareInfo                    = dm.OtaFirmwareInfo
	OtaFirmwareReadReq                 = dm.OtaFirmwareReadReq
	OtaFirmwareReadResp                = dm.OtaFirmwareReadResp
	OtaFirmwareResp                    = dm.OtaFirmwareResp
	OtaFirmwareUpdateReq               = dm.OtaFirmwareUpdateReq
	OtaFirmwareVerifyReq               = dm.OtaFirmwareVerifyReq
	OtaJobByDeviceIndexReq             = dm.OtaJobByDeviceIndexReq
	OtaJobByFirmwareIndexReq           = dm.OtaJobByFirmwareIndexReq
	OtaJobInfo                         = dm.OtaJobInfo
	OtaJobInfoIndexResp                = dm.OtaJobInfoIndexResp
	OtaModuleInfo                      = dm.OtaModuleInfo
	OtaPageInfo                        = dm.OtaPageInfo
	OtaPromptIndexReq                  = dm.OtaPromptIndexReq
	OtaPromptIndexResp                 = dm.OtaPromptIndexResp
	OtaTaskBatchReq                    = dm.OtaTaskBatchReq
	OtaTaskBatchResp                   = dm.OtaTaskBatchResp
	OtaTaskByJobIndexResp              = dm.OtaTaskByJobIndexResp
	OtaTaskCancleReq                   = dm.OtaTaskCancleReq
	OtaTaskCreatResp                   = dm.OtaTaskCreatResp
	OtaTaskCreateReq                   = dm.OtaTaskCreateReq
	OtaTaskDeviceCancleReq             = dm.OtaTaskDeviceCancleReq
	OtaTaskDeviceIndexReq              = dm.OtaTaskDeviceIndexReq
	OtaTaskDeviceIndexResp             = dm.OtaTaskDeviceIndexResp
	OtaTaskDeviceInfo                  = dm.OtaTaskDeviceInfo
	OtaTaskDeviceProcessReq            = dm.OtaTaskDeviceProcessReq
	OtaTaskDeviceReadReq               = dm.OtaTaskDeviceReadReq
	OtaTaskIndexReq                    = dm.OtaTaskIndexReq
	OtaTaskIndexResp                   = dm.OtaTaskIndexResp
	OtaTaskInfo                        = dm.OtaTaskInfo
	OtaTaskReadReq                     = dm.OtaTaskReadReq
	OtaTaskReadResp                    = dm.OtaTaskReadResp
	OtaUpTaskInfo                      = dm.OtaUpTaskInfo
	PageInfo                           = dm.PageInfo
	PageInfo_OrderBy                   = dm.PageInfo_OrderBy
	Point                              = dm.Point
	ProductCategory                    = dm.ProductCategory
	ProductCategoryIndexReq            = dm.ProductCategoryIndexReq
	ProductCategoryIndexResp           = dm.ProductCategoryIndexResp
	ProductCustom                      = dm.ProductCustom
	ProductCustomReadReq               = dm.ProductCustomReadReq
	ProductInfo                        = dm.ProductInfo
	ProductInfoDeleteReq               = dm.ProductInfoDeleteReq
	ProductInfoIndexReq                = dm.ProductInfoIndexReq
	ProductInfoIndexResp               = dm.ProductInfoIndexResp
	ProductInfoReadReq                 = dm.ProductInfoReadReq
	ProductRemoteConfig                = dm.ProductRemoteConfig
	ProductSchemaCreateReq             = dm.ProductSchemaCreateReq
	ProductSchemaDeleteReq             = dm.ProductSchemaDeleteReq
	ProductSchemaIndexReq              = dm.ProductSchemaIndexReq
	ProductSchemaIndexResp             = dm.ProductSchemaIndexResp
	ProductSchemaInfo                  = dm.ProductSchemaInfo
	ProductSchemaTslImportReq          = dm.ProductSchemaTslImportReq
	ProductSchemaTslReadReq            = dm.ProductSchemaTslReadReq
	ProductSchemaTslReadResp           = dm.ProductSchemaTslReadResp
	ProductSchemaUpdateReq             = dm.ProductSchemaUpdateReq
	PropertyControlMultiSendReq        = dm.PropertyControlMultiSendReq
	PropertyControlMultiSendResp       = dm.PropertyControlMultiSendResp
	PropertyControlSendMsg             = dm.PropertyControlSendMsg
	PropertyControlSendReq             = dm.PropertyControlSendReq
	PropertyControlSendResp            = dm.PropertyControlSendResp
	PropertyGetReportSendReq           = dm.PropertyGetReportSendReq
	PropertyGetReportSendResp          = dm.PropertyGetReportSendResp
	PropertyIndex                      = dm.PropertyIndex
	PropertyIndexResp                  = dm.PropertyIndexResp
	PropertyLatestIndexReq             = dm.PropertyLatestIndexReq
	PropertyLogIndexReq                = dm.PropertyLogIndexReq
	ProtocolInfo                       = dm.ProtocolInfo
	ProtocolInfoIndexReq               = dm.ProtocolInfoIndexReq
	ProtocolInfoIndexResp              = dm.ProtocolInfoIndexResp
	RemoteConfigCreateReq              = dm.RemoteConfigCreateReq
	RemoteConfigIndexReq               = dm.RemoteConfigIndexReq
	RemoteConfigIndexResp              = dm.RemoteConfigIndexResp
	RemoteConfigLastReadReq            = dm.RemoteConfigLastReadReq
	RemoteConfigLastReadResp           = dm.RemoteConfigLastReadResp
	RemoteConfigPushAllReq             = dm.RemoteConfigPushAllReq
	RespReadReq                        = dm.RespReadReq
	RootCheckReq                       = dm.RootCheckReq
	SdkLogIndex                        = dm.SdkLogIndex
	SdkLogIndexReq                     = dm.SdkLogIndexReq
	SdkLogIndexResp                    = dm.SdkLogIndexResp
	SendMsgReq                         = dm.SendMsgReq
	SendMsgResp                        = dm.SendMsgResp
	SendOption                         = dm.SendOption
	ShadowIndex                        = dm.ShadowIndex
	ShadowIndexResp                    = dm.ShadowIndexResp
	StaticUpgradeDeviceInfo            = dm.StaticUpgradeDeviceInfo
	StaticUpgradeJobReq                = dm.StaticUpgradeJobReq
	Tag                                = dm.Tag
	TimeRange                          = dm.TimeRange
	UpgradeJobResp                     = dm.UpgradeJobResp
	UserDeviceCollectSave              = dm.UserDeviceCollectSave
	UserDeviceShareIndexReq            = dm.UserDeviceShareIndexReq
	UserDeviceShareIndexResp           = dm.UserDeviceShareIndexResp
	UserDeviceShareInfo                = dm.UserDeviceShareInfo
	UserDeviceShareReadReq             = dm.UserDeviceShareReadReq
	VerifyOtaFirmwareReq               = dm.VerifyOtaFirmwareReq
	WithID                             = dm.WithID

	OtaTaskManage interface {
		// 创建批量升级任务
		OtaTaskCreate(ctx context.Context, in *OtaTaskCreateReq, opts ...grpc.CallOption) (*OtaTaskCreatResp, error)
		OtaTaskUpdate(ctx context.Context, in *OtaTaskInfo, opts ...grpc.CallOption) (*OtaCommonResp, error)
		// 批量取消升级任务
		OtaTaskCancle(ctx context.Context, in *OtaTaskCancleReq, opts ...grpc.CallOption) (*OtaCommonResp, error)
		OtaTaskIndex(ctx context.Context, in *OtaTaskIndexReq, opts ...grpc.CallOption) (*OtaTaskIndexResp, error)
		// 升级任务详情
		OtaTaskRead(ctx context.Context, in *OtaTaskReadReq, opts ...grpc.CallOption) (*OtaTaskReadResp, error)
		// 升级批次详情列表
		OtaTaskDeviceIndex(ctx context.Context, in *OtaTaskDeviceIndexReq, opts ...grpc.CallOption) (*OtaTaskDeviceIndexResp, error)
		// 设备升级状态详情
		OtaTaskDeviceRead(ctx context.Context, in *OtaTaskDeviceReadReq, opts ...grpc.CallOption) (*OtaTaskDeviceInfo, error)
		// 获取当前可执行批次信息
		OtaTaskDeviceEnableBatch(ctx context.Context, in *OtaTaskBatchReq, opts ...grpc.CallOption) (*OtaTaskBatchResp, error)
		// 升级进度上报
		OtaTaskDeviceProcess(ctx context.Context, in *OtaTaskDeviceProcessReq, opts ...grpc.CallOption) (*OtaCommonResp, error)
		// 取消单个设备的升级
		OtaTaskDeviceCancle(ctx context.Context, in *OtaTaskDeviceCancleReq, opts ...grpc.CallOption) (*OtaCommonResp, error)
	}

	defaultOtaTaskManage struct {
		cli zrpc.Client
	}

	directOtaTaskManage struct {
		svcCtx *svc.ServiceContext
		svr    dm.OtaTaskManageServer
	}
)

func NewOtaTaskManage(cli zrpc.Client) OtaTaskManage {
	return &defaultOtaTaskManage{
		cli: cli,
	}
}

func NewDirectOtaTaskManage(svcCtx *svc.ServiceContext, svr dm.OtaTaskManageServer) OtaTaskManage {
	return &directOtaTaskManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 创建批量升级任务
func (m *defaultOtaTaskManage) OtaTaskCreate(ctx context.Context, in *OtaTaskCreateReq, opts ...grpc.CallOption) (*OtaTaskCreatResp, error) {
	client := dm.NewOtaTaskManageClient(m.cli.Conn())
	return client.OtaTaskCreate(ctx, in, opts...)
}

// 创建批量升级任务
func (d *directOtaTaskManage) OtaTaskCreate(ctx context.Context, in *OtaTaskCreateReq, opts ...grpc.CallOption) (*OtaTaskCreatResp, error) {
	return d.svr.OtaTaskCreate(ctx, in)
}

func (m *defaultOtaTaskManage) OtaTaskUpdate(ctx context.Context, in *OtaTaskInfo, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	client := dm.NewOtaTaskManageClient(m.cli.Conn())
	return client.OtaTaskUpdate(ctx, in, opts...)
}

func (d *directOtaTaskManage) OtaTaskUpdate(ctx context.Context, in *OtaTaskInfo, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	return d.svr.OtaTaskUpdate(ctx, in)
}

// 批量取消升级任务
func (m *defaultOtaTaskManage) OtaTaskCancle(ctx context.Context, in *OtaTaskCancleReq, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	client := dm.NewOtaTaskManageClient(m.cli.Conn())
	return client.OtaTaskCancle(ctx, in, opts...)
}

// 批量取消升级任务
func (d *directOtaTaskManage) OtaTaskCancle(ctx context.Context, in *OtaTaskCancleReq, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	return d.svr.OtaTaskCancle(ctx, in)
}

func (m *defaultOtaTaskManage) OtaTaskIndex(ctx context.Context, in *OtaTaskIndexReq, opts ...grpc.CallOption) (*OtaTaskIndexResp, error) {
	client := dm.NewOtaTaskManageClient(m.cli.Conn())
	return client.OtaTaskIndex(ctx, in, opts...)
}

func (d *directOtaTaskManage) OtaTaskIndex(ctx context.Context, in *OtaTaskIndexReq, opts ...grpc.CallOption) (*OtaTaskIndexResp, error) {
	return d.svr.OtaTaskIndex(ctx, in)
}

// 升级任务详情
func (m *defaultOtaTaskManage) OtaTaskRead(ctx context.Context, in *OtaTaskReadReq, opts ...grpc.CallOption) (*OtaTaskReadResp, error) {
	client := dm.NewOtaTaskManageClient(m.cli.Conn())
	return client.OtaTaskRead(ctx, in, opts...)
}

// 升级任务详情
func (d *directOtaTaskManage) OtaTaskRead(ctx context.Context, in *OtaTaskReadReq, opts ...grpc.CallOption) (*OtaTaskReadResp, error) {
	return d.svr.OtaTaskRead(ctx, in)
}

// 升级批次详情列表
func (m *defaultOtaTaskManage) OtaTaskDeviceIndex(ctx context.Context, in *OtaTaskDeviceIndexReq, opts ...grpc.CallOption) (*OtaTaskDeviceIndexResp, error) {
	client := dm.NewOtaTaskManageClient(m.cli.Conn())
	return client.OtaTaskDeviceIndex(ctx, in, opts...)
}

// 升级批次详情列表
func (d *directOtaTaskManage) OtaTaskDeviceIndex(ctx context.Context, in *OtaTaskDeviceIndexReq, opts ...grpc.CallOption) (*OtaTaskDeviceIndexResp, error) {
	return d.svr.OtaTaskDeviceIndex(ctx, in)
}

// 设备升级状态详情
func (m *defaultOtaTaskManage) OtaTaskDeviceRead(ctx context.Context, in *OtaTaskDeviceReadReq, opts ...grpc.CallOption) (*OtaTaskDeviceInfo, error) {
	client := dm.NewOtaTaskManageClient(m.cli.Conn())
	return client.OtaTaskDeviceRead(ctx, in, opts...)
}

// 设备升级状态详情
func (d *directOtaTaskManage) OtaTaskDeviceRead(ctx context.Context, in *OtaTaskDeviceReadReq, opts ...grpc.CallOption) (*OtaTaskDeviceInfo, error) {
	return d.svr.OtaTaskDeviceRead(ctx, in)
}

// 获取当前可执行批次信息
func (m *defaultOtaTaskManage) OtaTaskDeviceEnableBatch(ctx context.Context, in *OtaTaskBatchReq, opts ...grpc.CallOption) (*OtaTaskBatchResp, error) {
	client := dm.NewOtaTaskManageClient(m.cli.Conn())
	return client.OtaTaskDeviceEnableBatch(ctx, in, opts...)
}

// 获取当前可执行批次信息
func (d *directOtaTaskManage) OtaTaskDeviceEnableBatch(ctx context.Context, in *OtaTaskBatchReq, opts ...grpc.CallOption) (*OtaTaskBatchResp, error) {
	return d.svr.OtaTaskDeviceEnableBatch(ctx, in)
}

// 升级进度上报
func (m *defaultOtaTaskManage) OtaTaskDeviceProcess(ctx context.Context, in *OtaTaskDeviceProcessReq, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	client := dm.NewOtaTaskManageClient(m.cli.Conn())
	return client.OtaTaskDeviceProcess(ctx, in, opts...)
}

// 升级进度上报
func (d *directOtaTaskManage) OtaTaskDeviceProcess(ctx context.Context, in *OtaTaskDeviceProcessReq, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	return d.svr.OtaTaskDeviceProcess(ctx, in)
}

// 取消单个设备的升级
func (m *defaultOtaTaskManage) OtaTaskDeviceCancle(ctx context.Context, in *OtaTaskDeviceCancleReq, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	client := dm.NewOtaTaskManageClient(m.cli.Conn())
	return client.OtaTaskDeviceCancle(ctx, in, opts...)
}

// 取消单个设备的升级
func (d *directOtaTaskManage) OtaTaskDeviceCancle(ctx context.Context, in *OtaTaskDeviceCancleReq, opts ...grpc.CallOption) (*OtaCommonResp, error) {
	return d.svr.OtaTaskDeviceCancle(ctx, in)
}
