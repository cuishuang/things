package deviceMsgEvent

import (
	"context"
	"github.com/i-Things/things/service/dmsvr/internal/domain/deviceLog"
	"strings"
	"time"

	"gitee.com/i-Things/share/domain/deviceMsg"
	"gitee.com/i-Things/share/domain/deviceMsg/msgOta"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type OtaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	topics []string
	dreq   msgOta.Req
	preq   msgOta.Process
}

func NewOtaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OtaLogic {
	return &OtaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *OtaLogic) initMsg(msg *deviceMsg.PublishMsg) error {
	l.topics = strings.Split(msg.Topic, "/")
	if len(l.topics) < 5 || l.topics[1] != "up" {
		return errors.Parameter.AddDetail("initMsg topic is err:" + msg.Topic)
	}
	return nil
}
func (l *OtaLogic) Handle(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	l.Infof("%s req=%+v", utils.FuncName(), msg)
	err = l.initMsg(msg)
	if err != nil {
		return nil, err
	}
	respMsg, err = func() (respMsg *deviceMsg.PublishMsg, err error) {
		switch l.topics[2] {
		case msgOta.TypeReport: //固件升级消息上行 上报版本、模块信息
			return l.HandleDynamicReport(msg)
		case msgOta.TypeProgress: //设备端上报升级进度
			return l.HandleResp(msg)
		default:
			return nil, errors.Parameter.AddDetail("things topic is err:" + msg.Topic)
		}
	}()
	l.svcCtx.HubLogRepo.Insert(l.ctx, &deviceLog.Hub{
		ProductID:  msg.ProductID,
		Action:     "otaLog",
		Timestamp:  l.dreq.GetTimeStamp(), // 操作时间
		DeviceName: msg.DeviceName,
		TraceID:    utils.TraceIdFromContext(l.ctx),
		RequestID:  l.dreq.MsgToken,
		Content:    string(msg.Payload),
		Topic:      msg.Topic,
		ResultCode: errors.Fmt(err).GetCode(),
	})
	return
}

func (l *OtaLogic) HandleDynamicReport(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	//固件升级消息上行 上报版本、模块信息
	err = utils.Unmarshal([]byte(msg.Payload), &l.dreq)
	if err != nil {
		return nil, errors.Parameter.AddDetail("ota topic is err:" + msg.Topic)
	}
	err = l.dreq.VerifyReqParam()
	if err != nil {
		return nil, err
	}
	//获取当前设备是否在动态可执行批次
	//di := &dm.OTATaskByDeviceNameReq{
	//	ProductID:  msg.ProductID,
	//	DeviceName: msg.DeviceName,
	//}
	//taskInfo, err := otatasklogic.NewOtaTaskByDeviceNameLogic(l.ctx, l.svcCtx).OtaTaskByDeviceName(di)
	//if err != nil {
	//	//没有找到可执行的升级任务
	//	return nil, err
	//}
	//firmware, err := relationDB.NewOtaFirmwareInfoRepo(l.ctx).FindOneByFilter(l.ctx, relationDB.OtaFirmwareInfoFilter{ID: taskInfo.FirmwareID, WithFiles: true})
	//var files []msgOta.File
	//_ = copier.Copy(&files, &firmware.Files)
	//data := msgOta.UpgradeParams{
	//	Version:          firmware.Version,
	//	IsDiff:           firmware.IsDiff,
	//	SignMethod:       firmware.SignMethod,
	//	DownloadProtocol: "https",
	//	Files:            files,
	//}
	//return l.DeviceResp(msg, errors.OK, data), nil
	return nil, err
}

func (l *OtaLogic) HandleReport(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	////固件升级消息上行 上报版本、模块信息
	//err = utils.Unmarshal([]byte(msg.Payload), &l.dreq)
	//if err != nil {
	//	return nil, errors.Parameter.AddDetail("ota topic is err:" + msg.Topic)
	//}
	//err = l.dreq.VerifyReqParam()
	//if err != nil {
	//	return nil, err
	//}
	////获取当前设备可用升级包
	//di := &dm.OtaTaskBatchReq{
	//	ID:         l.dreq.Params.ID,
	//	ProductID:  msg.ProductID,
	//	DeviceName: msg.DeviceName,
	//	Version:    l.dreq.GetVersion(),
	//}
	//otd, err := otataskmanage.NewOtaTaskManageServer(l.svcCtx).OtaTaskDeviceEnableBatch(l.ctx, di)
	//if err != nil {
	//	//没找到可执行的升级任务
	//	return nil, err
	//}
	//firmwareInfo, err := relationDB.NewOtaFirmwareInfoRepo(l.ctx).FindOneByFilter(l.ctx, relationDB.OtaFirmwareInfoFilter{ID: otd.FirmwareID, WithFiles: true})
	//if err != nil {
	//	return nil, err
	//}
	//data := map[string]any{
	//	"id":         otd.ID,
	//	"version":    firmwareInfo.Version,
	//	"signMethod": firmwareInfo.SignMethod,
	//	"isDiff":     firmwareInfo.IsDiff,
	//}
	//var files = make([]map[string]any, len(firmwareInfo.Files))
	//for k, v := range firmwareInfo.Files {
	//	url, _ := l.svcCtx.OssClient.PrivateBucket().SignedGetUrl(l.ctx, v.FilePath, 3600*24, common.OptionKv{})
	//	files[k] = map[string]any{
	//		"size":      v.Size,
	//		"name":      v.Name,
	//		"url":       url,
	//		"signature": v.Signature,
	//	}
	//}
	//data["files"] = files
	//return l.DeviceResp(msg, errors.OK, data), nil
	return nil, err
}
func (l *OtaLogic) HandleResp(msg *deviceMsg.PublishMsg) (respMsg *deviceMsg.PublishMsg, err error) {
	//设备端上报升级进度
	err = utils.Unmarshal([]byte(msg.Payload), &l.preq)
	if err != nil {
		return nil, errors.Parameter.AddDetail("ota topic is err:" + msg.Topic)
	}
	err = l.preq.VerifyReqParam()
	if err != nil {
		return nil, err
	}
	//req := dm.OtaTaskDeviceProcessReq{
	//	ProductID:  msg.ProductID,
	//	DeviceName: msg.DeviceName,
	//	Step:       l.preq.Params.Step,
	//}
	//_, err = otatasklogic.NewOtaTaskDeviceProcessLogic(l.ctx, l.svcCtx).OtaTaskDeviceProcess(&req)
	if err != nil {
		return nil, errors.Parameter.AddDetail("ota process rpc is err:" + msg.Topic)
	}
	return l.DeviceResp(msg, errors.OK, nil), nil
}
func (l *OtaLogic) DeviceResp(msg *deviceMsg.PublishMsg, err error, data any) *deviceMsg.PublishMsg {
	resp := &deviceMsg.CommonMsg{
		Method:    deviceMsg.GetRespMethod(l.dreq.Method),
		MsgToken:  l.dreq.MsgToken,
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
	}
	if l.topics[2] == "report" {
		l.topics[2] = "upgrade" //下行 升级包详情
	}
	return &deviceMsg.PublishMsg{
		Handle:       msg.Handle,
		Type:         l.topics[2],
		Payload:      resp.AddStatus(err).Bytes(),
		ProductID:    msg.ProductID,
		DeviceName:   msg.DeviceName,
		ProtocolCode: msg.ProtocolCode,
	}
}
