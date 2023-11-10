// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	systemapi "github.com/i-Things/things/src/apisvr/internal/handler/system/api"
	systemauth "github.com/i-Things/things/src/apisvr/internal/handler/system/auth"
	systemcommon "github.com/i-Things/things/src/apisvr/internal/handler/system/common"
	systemlog "github.com/i-Things/things/src/apisvr/internal/handler/system/log"
	systemmenu "github.com/i-Things/things/src/apisvr/internal/handler/system/menu"
	systemrole "github.com/i-Things/things/src/apisvr/internal/handler/system/role"
	systemuser "github.com/i-Things/things/src/apisvr/internal/handler/system/user"
	thingsdeviceauth "github.com/i-Things/things/src/apisvr/internal/handler/things/device/auth"
	thingsdeviceauth5 "github.com/i-Things/things/src/apisvr/internal/handler/things/device/auth5"
	thingsdevicegateway "github.com/i-Things/things/src/apisvr/internal/handler/things/device/gateway"
	thingsdeviceinfo "github.com/i-Things/things/src/apisvr/internal/handler/things/device/info"
	thingsdeviceinteract "github.com/i-Things/things/src/apisvr/internal/handler/things/device/interact"
	thingsdevicemsg "github.com/i-Things/things/src/apisvr/internal/handler/things/device/msg"
	thingsgroupdevice "github.com/i-Things/things/src/apisvr/internal/handler/things/group/device"
	thingsgroupinfo "github.com/i-Things/things/src/apisvr/internal/handler/things/group/info"
	thingsproductcustom "github.com/i-Things/things/src/apisvr/internal/handler/things/product/custom"
	thingsproductinfo "github.com/i-Things/things/src/apisvr/internal/handler/things/product/info"
	thingsproductremoteConfig "github.com/i-Things/things/src/apisvr/internal/handler/things/product/remoteConfig"
	thingsproductschema "github.com/i-Things/things/src/apisvr/internal/handler/things/product/schema"
	thingsrulealarmdealRecord "github.com/i-Things/things/src/apisvr/internal/handler/things/rule/alarm/dealRecord"
	thingsrulealarminfo "github.com/i-Things/things/src/apisvr/internal/handler/things/rule/alarm/info"
	thingsrulealarmlog "github.com/i-Things/things/src/apisvr/internal/handler/things/rule/alarm/log"
	thingsrulealarmrecord "github.com/i-Things/things/src/apisvr/internal/handler/things/rule/alarm/record"
	thingsrulealarmscene "github.com/i-Things/things/src/apisvr/internal/handler/things/rule/alarm/scene"
	thingsrulesceneinfo "github.com/i-Things/things/src/apisvr/internal/handler/things/rule/scene/info"
	thingsvidmgrhooks "github.com/i-Things/things/src/apisvr/internal/handler/things/vidmgr/hooks"
	thingsvidmgrindexapi "github.com/i-Things/things/src/apisvr/internal/handler/things/vidmgr/indexapi"
	thingsvidmgrinfo "github.com/i-Things/things/src/apisvr/internal/handler/things/vidmgr/info"
	"github.com/i-Things/things/src/apisvr/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/captcha",
				Handler: systemuser.CaptchaHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: systemuser.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/system/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemuser.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: systemuser.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: systemuser.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: systemuser.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: systemuser.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/resource-read",
					Handler: systemuser.ResourceReadHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/system/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemmenu.MenuCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: systemmenu.MenuIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: systemmenu.MenuUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: systemmenu.MenuDeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/system/menu"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemrole.RoleCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: systemrole.RoleIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: systemrole.RoleUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: systemrole.RoleDeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role-menu/update",
					Handler: systemrole.RoleMenuUpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/system/role"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/login/index",
					Handler: systemlog.LoginIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/oper/index",
					Handler: systemlog.OperIndexHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/system/log"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: systemapi.ApiCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: systemapi.ApiIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: systemapi.ApiUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: systemapi.ApiDeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/system/api"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/config",
				Handler: systemcommon.ConfigHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/upload-url/create",
				Handler: systemcommon.UploadUrlCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/upload-file",
				Handler: systemcommon.UploadFileHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/system/common"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/multiUpdate",
					Handler: systemauth.AuthApiMultiUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: systemauth.AuthApiIndexHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/system/auth/api"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: thingsdeviceauth.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/access",
				Handler: thingsdeviceauth.AccessHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/root-check",
				Handler: thingsdeviceauth.RootCheckHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: thingsdeviceauth.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/device/auth"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: thingsdeviceauth5.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/access",
				Handler: thingsdeviceauth5.AccessHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/device/auth5"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/hub-log/index",
					Handler: thingsdevicemsg.HubLogIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/sdk-log/index",
					Handler: thingsdevicemsg.SdkLogIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property-log/index",
					Handler: thingsdevicemsg.PropertyLogIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property-latest/index",
					Handler: thingsdevicemsg.PropertyLatestIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/event-log/index",
					Handler: thingsdevicemsg.EventLogIndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/shadow/index",
					Handler: thingsdevicemsg.ShadowIndexHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/device/msg"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/multi-create",
					Handler: thingsdevicegateway.MultiCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsdevicegateway.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-delete",
					Handler: thingsdevicegateway.MultiDeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/device/gateway"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsdeviceinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsdeviceinfo.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsdeviceinfo.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsdeviceinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsdeviceinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/count",
					Handler: thingsdeviceinfo.CountHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-import",
					Handler: thingsdeviceinfo.MultiImportHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/device/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/send-action",
					Handler: thingsdeviceinteract.SendActionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/action-read",
					Handler: thingsdeviceinteract.ActionReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/send-property",
					Handler: thingsdeviceinteract.SendPropertyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-send-property",
					Handler: thingsdeviceinteract.MultiSendPropertyHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/property-read",
					Handler: thingsdeviceinteract.PropertyReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/send-msg",
					Handler: thingsdeviceinteract.SendMsgHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/get-property-reply",
					Handler: thingsdeviceinteract.GetPropertyReplyHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/device/interact"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsproductinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsproductinfo.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsproductinfo.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsproductinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsproductinfo.ReadHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/product/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/tsl-import",
					Handler: thingsproductschema.TslImportHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/tsl-read",
					Handler: thingsproductschema.TslReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsproductschema.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsproductschema.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsproductschema.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsproductschema.IndexHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/product/schema"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsproductremoteConfig.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsproductremoteConfig.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/push-all",
					Handler: thingsproductremoteConfig.PushAllHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/lastest-read",
					Handler: thingsproductremoteConfig.LastestReadHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/product/remote-config"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsproductcustom.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsproductcustom.ReadHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/product/custom"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsgroupinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsgroupinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsgroupinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsgroupinfo.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsgroupinfo.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/group/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/multi-create",
					Handler: thingsgroupdevice.MultiCreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsgroupdevice.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/multi-delete",
					Handler: thingsgroupdevice.MultiDeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/group/device"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsrulesceneinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsrulesceneinfo.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsrulesceneinfo.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsrulesceneinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsrulesceneinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/manually-trigger",
					Handler: thingsrulesceneinfo.ManuallyTriggerHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/rule/scene/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsrulealarmdealRecord.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsrulealarmdealRecord.CreateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/rule/alarm/deal-record"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsrulealarminfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsrulealarminfo.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsrulealarminfo.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsrulealarminfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsrulealarminfo.ReadHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/rule/alarm/info"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsrulealarmlog.IndexHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/rule/alarm/log"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsrulealarmrecord.IndexHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/rule/alarm/record"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/multi-update",
					Handler: thingsrulealarmscene.MultiUpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsrulealarmscene.DeleteHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/rule/alarm/scene"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.SetupWare, serverCtx.CheckTokenWare, serverCtx.TeardownWare},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: thingsvidmgrinfo.CreateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: thingsvidmgrinfo.UpdateHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: thingsvidmgrinfo.DeleteHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/index",
					Handler: thingsvidmgrinfo.IndexHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/read",
					Handler: thingsvidmgrinfo.ReadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/count",
					Handler: thingsvidmgrinfo.CountHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/api/v1/things/vidmgr/info"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/onFlowReport",
				Handler: thingsvidmgrhooks.OnFlowReportHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onHttpAccess",
				Handler: thingsvidmgrhooks.OnHttpAccessHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onPlay",
				Handler: thingsvidmgrhooks.OnPlayHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onPublish",
				Handler: thingsvidmgrhooks.OnPublishHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onRecordMp4",
				Handler: thingsvidmgrhooks.OnRecordMp4Handler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onRtspRealm",
				Handler: thingsvidmgrhooks.OnRtspRealmHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onRtspAuth",
				Handler: thingsvidmgrhooks.OnRtspAuthHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onShellLogin",
				Handler: thingsvidmgrhooks.OnShellLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onStreamChanged",
				Handler: thingsvidmgrhooks.OnStreamChangedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onStreamNoneReader",
				Handler: thingsvidmgrhooks.OnStreamNoneReaderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onStreamNotFound",
				Handler: thingsvidmgrhooks.OnStreamNotFoundHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onServerStarted",
				Handler: thingsvidmgrhooks.OnServerStartedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onServerKeepalive",
				Handler: thingsvidmgrhooks.OnServerKeepaliveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/onRtpServerTimeout",
				Handler: thingsvidmgrhooks.OnRtpServerTimeoutHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/vidmgr/hooks"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/getApiList",
				Handler: thingsvidmgrindexapi.GetApiListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getThreadsLoad",
				Handler: thingsvidmgrindexapi.GetThreadsLoadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getWorkThreadsLoad",
				Handler: thingsvidmgrindexapi.GetWorkThreadsLoadHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getServerConfig",
				Handler: thingsvidmgrindexapi.GetServerConfigHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/setServerConfig",
				Handler: thingsvidmgrindexapi.SetServerConfigHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/restartServer",
				Handler: thingsvidmgrindexapi.RestartServerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getMediaList",
				Handler: thingsvidmgrindexapi.GetMediaListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/closeStream",
				Handler: thingsvidmgrindexapi.CloseStreamHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/closeStreams",
				Handler: thingsvidmgrindexapi.CloseStreamsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getAllSession",
				Handler: thingsvidmgrindexapi.GetAllSessionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/kick_session",
				Handler: thingsvidmgrindexapi.KickSessionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/kickSessions",
				Handler: thingsvidmgrindexapi.KickSessionsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addStreamProxy",
				Handler: thingsvidmgrindexapi.AddStreamProxyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delStreamProxy",
				Handler: thingsvidmgrindexapi.DelStreamProxyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addFFmpegSource",
				Handler: thingsvidmgrindexapi.AddFFmpegSourceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delFFmpegSource",
				Handler: thingsvidmgrindexapi.DelFFmpegSourceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/isMediaOnline",
				Handler: thingsvidmgrindexapi.IsMediaOnlineHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getMediaInfo",
				Handler: thingsvidmgrindexapi.GetMediaInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getRtpInfo",
				Handler: thingsvidmgrindexapi.GetRtpInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getMp4RecordFile",
				Handler: thingsvidmgrindexapi.GetMp4RecordFileHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/startRecord",
				Handler: thingsvidmgrindexapi.StartRecordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/stopRecord",
				Handler: thingsvidmgrindexapi.StopRecordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/isRecording",
				Handler: thingsvidmgrindexapi.IsRecordingHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getSnap",
				Handler: thingsvidmgrindexapi.GetSnapHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/openRtpServer",
				Handler: thingsvidmgrindexapi.OpenRtpServerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/closeRtpServer",
				Handler: thingsvidmgrindexapi.CloseRtpServerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/listRtpServer",
				Handler: thingsvidmgrindexapi.ListRtpServerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/startSendRtp",
				Handler: thingsvidmgrindexapi.StartSendRtpHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/startSendRtpPassive",
				Handler: thingsvidmgrindexapi.StartSendRtpPassiveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/stopSendRtp",
				Handler: thingsvidmgrindexapi.StopSendRtpHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getStatistic",
				Handler: thingsvidmgrindexapi.GetStatisticHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/addStreamPusherProxy",
				Handler: thingsvidmgrindexapi.AddStreamPusherProxyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delStreamPusherProxy",
				Handler: thingsvidmgrindexapi.DelStreamPusherProxyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/version",
				Handler: thingsvidmgrindexapi.VersionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getMediaPlayerList",
				Handler: thingsvidmgrindexapi.GetMediaPlayerListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1/things/vidmgr/indexapi"),
	)
}
