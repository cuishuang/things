// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: proto/vid.proto

package vid

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VidmgrMange_VidmgrInfoCreate_FullMethodName = "/vid.VidmgrMange/vidmgrInfoCreate"
	VidmgrMange_VidmgrInfoUpdate_FullMethodName = "/vid.VidmgrMange/vidmgrInfoUpdate"
	VidmgrMange_VidmgrInfoDelete_FullMethodName = "/vid.VidmgrMange/vidmgrInfoDelete"
	VidmgrMange_VidmgrInfoIndex_FullMethodName  = "/vid.VidmgrMange/vidmgrInfoIndex"
	VidmgrMange_VidmgrInfoRead_FullMethodName   = "/vid.VidmgrMange/vidmgrInfoRead"
	VidmgrMange_VidmgrInfoCount_FullMethodName  = "/vid.VidmgrMange/vidmgrInfoCount"
)

// VidmgrMangeClient is the client API for VidmgrMange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VidmgrMangeClient interface {
	// 新建服务
	VidmgrInfoCreate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error)
	// 更新服务
	VidmgrInfoUpdate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error)
	// 删除服务
	VidmgrInfoDelete(ctx context.Context, in *VidmgrInfoDeleteReq, opts ...grpc.CallOption) (*Response, error)
	// 获取服务列表
	VidmgrInfoIndex(ctx context.Context, in *VidmgrInfoIndexReq, opts ...grpc.CallOption) (*VidmgrInfoIndexResp, error)
	// 获取服务信息详情
	VidmgrInfoRead(ctx context.Context, in *VidmgrInfoReadReq, opts ...grpc.CallOption) (*VidmgrInfo, error)
	// 获取服务统计  在线，离线，未激活
	VidmgrInfoCount(ctx context.Context, in *VidmgrInfoCountReq, opts ...grpc.CallOption) (*VidmgrInfoCountResp, error)
}

type vidmgrMangeClient struct {
	cc grpc.ClientConnInterface
}

func NewVidmgrMangeClient(cc grpc.ClientConnInterface) VidmgrMangeClient {
	return &vidmgrMangeClient{cc}
}

func (c *vidmgrMangeClient) VidmgrInfoCreate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrMange_VidmgrInfoCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrMangeClient) VidmgrInfoUpdate(ctx context.Context, in *VidmgrInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrMange_VidmgrInfoUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrMangeClient) VidmgrInfoDelete(ctx context.Context, in *VidmgrInfoDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrMange_VidmgrInfoDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrMangeClient) VidmgrInfoIndex(ctx context.Context, in *VidmgrInfoIndexReq, opts ...grpc.CallOption) (*VidmgrInfoIndexResp, error) {
	out := new(VidmgrInfoIndexResp)
	err := c.cc.Invoke(ctx, VidmgrMange_VidmgrInfoIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrMangeClient) VidmgrInfoRead(ctx context.Context, in *VidmgrInfoReadReq, opts ...grpc.CallOption) (*VidmgrInfo, error) {
	out := new(VidmgrInfo)
	err := c.cc.Invoke(ctx, VidmgrMange_VidmgrInfoRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrMangeClient) VidmgrInfoCount(ctx context.Context, in *VidmgrInfoCountReq, opts ...grpc.CallOption) (*VidmgrInfoCountResp, error) {
	out := new(VidmgrInfoCountResp)
	err := c.cc.Invoke(ctx, VidmgrMange_VidmgrInfoCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VidmgrMangeServer is the server API for VidmgrMange service.
// All implementations must embed UnimplementedVidmgrMangeServer
// for forward compatibility
type VidmgrMangeServer interface {
	// 新建服务
	VidmgrInfoCreate(context.Context, *VidmgrInfo) (*Response, error)
	// 更新服务
	VidmgrInfoUpdate(context.Context, *VidmgrInfo) (*Response, error)
	// 删除服务
	VidmgrInfoDelete(context.Context, *VidmgrInfoDeleteReq) (*Response, error)
	// 获取服务列表
	VidmgrInfoIndex(context.Context, *VidmgrInfoIndexReq) (*VidmgrInfoIndexResp, error)
	// 获取服务信息详情
	VidmgrInfoRead(context.Context, *VidmgrInfoReadReq) (*VidmgrInfo, error)
	// 获取服务统计  在线，离线，未激活
	VidmgrInfoCount(context.Context, *VidmgrInfoCountReq) (*VidmgrInfoCountResp, error)
	mustEmbedUnimplementedVidmgrMangeServer()
}

// UnimplementedVidmgrMangeServer must be embedded to have forward compatible implementations.
type UnimplementedVidmgrMangeServer struct {
}

func (UnimplementedVidmgrMangeServer) VidmgrInfoCreate(context.Context, *VidmgrInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoCreate not implemented")
}
func (UnimplementedVidmgrMangeServer) VidmgrInfoUpdate(context.Context, *VidmgrInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoUpdate not implemented")
}
func (UnimplementedVidmgrMangeServer) VidmgrInfoDelete(context.Context, *VidmgrInfoDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoDelete not implemented")
}
func (UnimplementedVidmgrMangeServer) VidmgrInfoIndex(context.Context, *VidmgrInfoIndexReq) (*VidmgrInfoIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoIndex not implemented")
}
func (UnimplementedVidmgrMangeServer) VidmgrInfoRead(context.Context, *VidmgrInfoReadReq) (*VidmgrInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoRead not implemented")
}
func (UnimplementedVidmgrMangeServer) VidmgrInfoCount(context.Context, *VidmgrInfoCountReq) (*VidmgrInfoCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrInfoCount not implemented")
}
func (UnimplementedVidmgrMangeServer) mustEmbedUnimplementedVidmgrMangeServer() {}

// UnsafeVidmgrMangeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VidmgrMangeServer will
// result in compilation errors.
type UnsafeVidmgrMangeServer interface {
	mustEmbedUnimplementedVidmgrMangeServer()
}

func RegisterVidmgrMangeServer(s grpc.ServiceRegistrar, srv VidmgrMangeServer) {
	s.RegisterService(&VidmgrMange_ServiceDesc, srv)
}

func _VidmgrMange_VidmgrInfoCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrMangeServer).VidmgrInfoCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrMange_VidmgrInfoCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrMangeServer).VidmgrInfoCreate(ctx, req.(*VidmgrInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrMange_VidmgrInfoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrMangeServer).VidmgrInfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrMange_VidmgrInfoUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrMangeServer).VidmgrInfoUpdate(ctx, req.(*VidmgrInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrMange_VidmgrInfoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfoDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrMangeServer).VidmgrInfoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrMange_VidmgrInfoDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrMangeServer).VidmgrInfoDelete(ctx, req.(*VidmgrInfoDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrMange_VidmgrInfoIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfoIndexReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrMangeServer).VidmgrInfoIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrMange_VidmgrInfoIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrMangeServer).VidmgrInfoIndex(ctx, req.(*VidmgrInfoIndexReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrMange_VidmgrInfoRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfoReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrMangeServer).VidmgrInfoRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrMange_VidmgrInfoRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrMangeServer).VidmgrInfoRead(ctx, req.(*VidmgrInfoReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrMange_VidmgrInfoCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrInfoCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrMangeServer).VidmgrInfoCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrMange_VidmgrInfoCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrMangeServer).VidmgrInfoCount(ctx, req.(*VidmgrInfoCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VidmgrMange_ServiceDesc is the grpc.ServiceDesc for VidmgrMange service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VidmgrMange_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vid.VidmgrMange",
	HandlerType: (*VidmgrMangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "vidmgrInfoCreate",
			Handler:    _VidmgrMange_VidmgrInfoCreate_Handler,
		},
		{
			MethodName: "vidmgrInfoUpdate",
			Handler:    _VidmgrMange_VidmgrInfoUpdate_Handler,
		},
		{
			MethodName: "vidmgrInfoDelete",
			Handler:    _VidmgrMange_VidmgrInfoDelete_Handler,
		},
		{
			MethodName: "vidmgrInfoIndex",
			Handler:    _VidmgrMange_VidmgrInfoIndex_Handler,
		},
		{
			MethodName: "vidmgrInfoRead",
			Handler:    _VidmgrMange_VidmgrInfoRead_Handler,
		},
		{
			MethodName: "vidmgrInfoCount",
			Handler:    _VidmgrMange_VidmgrInfoCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vid.proto",
}

const (
	ViddevMange_VideoaddInfroCreate_FullMethodName = "/vid.ViddevMange/videoaddInfroCreate"
	ViddevMange_VideoaddInfroUpdate_FullMethodName = "/vid.ViddevMange/videoaddInfroUpdate"
	ViddevMange_VideoaddInfroDelete_FullMethodName = "/vid.ViddevMange/videoaddInfroDelete"
	ViddevMange_VideoaddInfroIndex_FullMethodName  = "/vid.ViddevMange/videoaddInfroIndex"
	ViddevMange_VideoaddInfroRead_FullMethodName   = "/vid.ViddevMange/videoaddInfroRead"
	ViddevMange_VideoaddInfroCount_FullMethodName  = "/vid.ViddevMange/videoaddInfroCount"
)

// ViddevMangeClient is the client API for ViddevMange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ViddevMangeClient interface {
	// 流添加
	VideoaddInfroCreate(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
	// 流更新
	VideoaddInfroUpdate(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
	// 删除流
	VideoaddInfroDelete(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
	// 获取流列表
	VideoaddInfroIndex(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
	// 获取流信息详情
	VideoaddInfroRead(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
	// 统计流 在线，离线，未激活
	VideoaddInfroCount(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
}

type viddevMangeClient struct {
	cc grpc.ClientConnInterface
}

func NewViddevMangeClient(cc grpc.ClientConnInterface) ViddevMangeClient {
	return &viddevMangeClient{cc}
}

func (c *viddevMangeClient) VideoaddInfroCreate(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ViddevMange_VideoaddInfroCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viddevMangeClient) VideoaddInfroUpdate(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ViddevMange_VideoaddInfroUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viddevMangeClient) VideoaddInfroDelete(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ViddevMange_VideoaddInfroDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viddevMangeClient) VideoaddInfroIndex(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ViddevMange_VideoaddInfroIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viddevMangeClient) VideoaddInfroRead(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ViddevMange_VideoaddInfroRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viddevMangeClient) VideoaddInfroCount(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ViddevMange_VideoaddInfroCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViddevMangeServer is the server API for ViddevMange service.
// All implementations must embed UnimplementedViddevMangeServer
// for forward compatibility
type ViddevMangeServer interface {
	// 流添加
	VideoaddInfroCreate(context.Context, *ViddevInfo) (*Response, error)
	// 流更新
	VideoaddInfroUpdate(context.Context, *ViddevInfo) (*Response, error)
	// 删除流
	VideoaddInfroDelete(context.Context, *ViddevInfo) (*Response, error)
	// 获取流列表
	VideoaddInfroIndex(context.Context, *ViddevInfo) (*Response, error)
	// 获取流信息详情
	VideoaddInfroRead(context.Context, *ViddevInfo) (*Response, error)
	// 统计流 在线，离线，未激活
	VideoaddInfroCount(context.Context, *ViddevInfo) (*Response, error)
	mustEmbedUnimplementedViddevMangeServer()
}

// UnimplementedViddevMangeServer must be embedded to have forward compatible implementations.
type UnimplementedViddevMangeServer struct {
}

func (UnimplementedViddevMangeServer) VideoaddInfroCreate(context.Context, *ViddevInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoaddInfroCreate not implemented")
}
func (UnimplementedViddevMangeServer) VideoaddInfroUpdate(context.Context, *ViddevInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoaddInfroUpdate not implemented")
}
func (UnimplementedViddevMangeServer) VideoaddInfroDelete(context.Context, *ViddevInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoaddInfroDelete not implemented")
}
func (UnimplementedViddevMangeServer) VideoaddInfroIndex(context.Context, *ViddevInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoaddInfroIndex not implemented")
}
func (UnimplementedViddevMangeServer) VideoaddInfroRead(context.Context, *ViddevInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoaddInfroRead not implemented")
}
func (UnimplementedViddevMangeServer) VideoaddInfroCount(context.Context, *ViddevInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VideoaddInfroCount not implemented")
}
func (UnimplementedViddevMangeServer) mustEmbedUnimplementedViddevMangeServer() {}

// UnsafeViddevMangeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ViddevMangeServer will
// result in compilation errors.
type UnsafeViddevMangeServer interface {
	mustEmbedUnimplementedViddevMangeServer()
}

func RegisterViddevMangeServer(s grpc.ServiceRegistrar, srv ViddevMangeServer) {
	s.RegisterService(&ViddevMange_ServiceDesc, srv)
}

func _ViddevMange_VideoaddInfroCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViddevInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViddevMangeServer).VideoaddInfroCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViddevMange_VideoaddInfroCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViddevMangeServer).VideoaddInfroCreate(ctx, req.(*ViddevInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViddevMange_VideoaddInfroUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViddevInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViddevMangeServer).VideoaddInfroUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViddevMange_VideoaddInfroUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViddevMangeServer).VideoaddInfroUpdate(ctx, req.(*ViddevInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViddevMange_VideoaddInfroDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViddevInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViddevMangeServer).VideoaddInfroDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViddevMange_VideoaddInfroDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViddevMangeServer).VideoaddInfroDelete(ctx, req.(*ViddevInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViddevMange_VideoaddInfroIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViddevInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViddevMangeServer).VideoaddInfroIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViddevMange_VideoaddInfroIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViddevMangeServer).VideoaddInfroIndex(ctx, req.(*ViddevInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViddevMange_VideoaddInfroRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViddevInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViddevMangeServer).VideoaddInfroRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViddevMange_VideoaddInfroRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViddevMangeServer).VideoaddInfroRead(ctx, req.(*ViddevInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViddevMange_VideoaddInfroCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViddevInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViddevMangeServer).VideoaddInfroCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ViddevMange_VideoaddInfroCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViddevMangeServer).VideoaddInfroCount(ctx, req.(*ViddevInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// ViddevMange_ServiceDesc is the grpc.ServiceDesc for ViddevMange service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ViddevMange_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vid.ViddevMange",
	HandlerType: (*ViddevMangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "videoaddInfroCreate",
			Handler:    _ViddevMange_VideoaddInfroCreate_Handler,
		},
		{
			MethodName: "videoaddInfroUpdate",
			Handler:    _ViddevMange_VideoaddInfroUpdate_Handler,
		},
		{
			MethodName: "videoaddInfroDelete",
			Handler:    _ViddevMange_VideoaddInfroDelete_Handler,
		},
		{
			MethodName: "videoaddInfroIndex",
			Handler:    _ViddevMange_VideoaddInfroIndex_Handler,
		},
		{
			MethodName: "videoaddInfroRead",
			Handler:    _ViddevMange_VideoaddInfroRead_Handler,
		},
		{
			MethodName: "videoaddInfroCount",
			Handler:    _ViddevMange_VideoaddInfroCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vid.proto",
}

const (
	VidmgrConfigMange_VidmgrConfigCreate_FullMethodName = "/vid.VidmgrConfigMange/vidmgrConfigCreate"
	VidmgrConfigMange_VidmgrConfigUpdate_FullMethodName = "/vid.VidmgrConfigMange/vidmgrConfigUpdate"
	VidmgrConfigMange_VidmgrConfigDelete_FullMethodName = "/vid.VidmgrConfigMange/vidmgrConfigDelete"
	VidmgrConfigMange_VidmgrConfigIndex_FullMethodName  = "/vid.VidmgrConfigMange/vidmgrConfigIndex"
	VidmgrConfigMange_VidmgrConfigRead_FullMethodName   = "/vid.VidmgrConfigMange/vidmgrConfigRead"
)

// VidmgrConfigMangeClient is the client API for VidmgrConfigMange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VidmgrConfigMangeClient interface {
	// 新建配置
	VidmgrConfigCreate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error)
	// 更新配置
	VidmgrConfigUpdate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error)
	// 删除配置
	VidmgrConfigDelete(ctx context.Context, in *VidmgrConfigDeleteReq, opts ...grpc.CallOption) (*Response, error)
	// 配置列表
	VidmgrConfigIndex(ctx context.Context, in *VidmgrConfigReadReq, opts ...grpc.CallOption) (*VidmgrConfigIndexResp, error)
	// 获取配置信息详情
	VidmgrConfigRead(ctx context.Context, in *VidmgrConfigReadReq, opts ...grpc.CallOption) (*VidmgrConfig, error)
}

type vidmgrConfigMangeClient struct {
	cc grpc.ClientConnInterface
}

func NewVidmgrConfigMangeClient(cc grpc.ClientConnInterface) VidmgrConfigMangeClient {
	return &vidmgrConfigMangeClient{cc}
}

func (c *vidmgrConfigMangeClient) VidmgrConfigCreate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrConfigMange_VidmgrConfigCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrConfigMangeClient) VidmgrConfigUpdate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrConfigMange_VidmgrConfigUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrConfigMangeClient) VidmgrConfigDelete(ctx context.Context, in *VidmgrConfigDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, VidmgrConfigMange_VidmgrConfigDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrConfigMangeClient) VidmgrConfigIndex(ctx context.Context, in *VidmgrConfigReadReq, opts ...grpc.CallOption) (*VidmgrConfigIndexResp, error) {
	out := new(VidmgrConfigIndexResp)
	err := c.cc.Invoke(ctx, VidmgrConfigMange_VidmgrConfigIndex_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vidmgrConfigMangeClient) VidmgrConfigRead(ctx context.Context, in *VidmgrConfigReadReq, opts ...grpc.CallOption) (*VidmgrConfig, error) {
	out := new(VidmgrConfig)
	err := c.cc.Invoke(ctx, VidmgrConfigMange_VidmgrConfigRead_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VidmgrConfigMangeServer is the server API for VidmgrConfigMange service.
// All implementations must embed UnimplementedVidmgrConfigMangeServer
// for forward compatibility
type VidmgrConfigMangeServer interface {
	// 新建配置
	VidmgrConfigCreate(context.Context, *VidmgrConfig) (*Response, error)
	// 更新配置
	VidmgrConfigUpdate(context.Context, *VidmgrConfig) (*Response, error)
	// 删除配置
	VidmgrConfigDelete(context.Context, *VidmgrConfigDeleteReq) (*Response, error)
	// 配置列表
	VidmgrConfigIndex(context.Context, *VidmgrConfigReadReq) (*VidmgrConfigIndexResp, error)
	// 获取配置信息详情
	VidmgrConfigRead(context.Context, *VidmgrConfigReadReq) (*VidmgrConfig, error)
	mustEmbedUnimplementedVidmgrConfigMangeServer()
}

// UnimplementedVidmgrConfigMangeServer must be embedded to have forward compatible implementations.
type UnimplementedVidmgrConfigMangeServer struct {
}

func (UnimplementedVidmgrConfigMangeServer) VidmgrConfigCreate(context.Context, *VidmgrConfig) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrConfigCreate not implemented")
}
func (UnimplementedVidmgrConfigMangeServer) VidmgrConfigUpdate(context.Context, *VidmgrConfig) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrConfigUpdate not implemented")
}
func (UnimplementedVidmgrConfigMangeServer) VidmgrConfigDelete(context.Context, *VidmgrConfigDeleteReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrConfigDelete not implemented")
}
func (UnimplementedVidmgrConfigMangeServer) VidmgrConfigIndex(context.Context, *VidmgrConfigReadReq) (*VidmgrConfigIndexResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrConfigIndex not implemented")
}
func (UnimplementedVidmgrConfigMangeServer) VidmgrConfigRead(context.Context, *VidmgrConfigReadReq) (*VidmgrConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VidmgrConfigRead not implemented")
}
func (UnimplementedVidmgrConfigMangeServer) mustEmbedUnimplementedVidmgrConfigMangeServer() {}

// UnsafeVidmgrConfigMangeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VidmgrConfigMangeServer will
// result in compilation errors.
type UnsafeVidmgrConfigMangeServer interface {
	mustEmbedUnimplementedVidmgrConfigMangeServer()
}

func RegisterVidmgrConfigMangeServer(s grpc.ServiceRegistrar, srv VidmgrConfigMangeServer) {
	s.RegisterService(&VidmgrConfigMange_ServiceDesc, srv)
}

func _VidmgrConfigMange_VidmgrConfigCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrConfigMangeServer).VidmgrConfigCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrConfigMange_VidmgrConfigCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrConfigMangeServer).VidmgrConfigCreate(ctx, req.(*VidmgrConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrConfigMange_VidmgrConfigUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrConfigMangeServer).VidmgrConfigUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrConfigMange_VidmgrConfigUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrConfigMangeServer).VidmgrConfigUpdate(ctx, req.(*VidmgrConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrConfigMange_VidmgrConfigDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrConfigDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrConfigMangeServer).VidmgrConfigDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrConfigMange_VidmgrConfigDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrConfigMangeServer).VidmgrConfigDelete(ctx, req.(*VidmgrConfigDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrConfigMange_VidmgrConfigIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrConfigReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrConfigMangeServer).VidmgrConfigIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrConfigMange_VidmgrConfigIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrConfigMangeServer).VidmgrConfigIndex(ctx, req.(*VidmgrConfigReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VidmgrConfigMange_VidmgrConfigRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VidmgrConfigReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VidmgrConfigMangeServer).VidmgrConfigRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VidmgrConfigMange_VidmgrConfigRead_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VidmgrConfigMangeServer).VidmgrConfigRead(ctx, req.(*VidmgrConfigReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VidmgrConfigMange_ServiceDesc is the grpc.ServiceDesc for VidmgrConfigMange service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VidmgrConfigMange_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vid.VidmgrConfigMange",
	HandlerType: (*VidmgrConfigMangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "vidmgrConfigCreate",
			Handler:    _VidmgrConfigMange_VidmgrConfigCreate_Handler,
		},
		{
			MethodName: "vidmgrConfigUpdate",
			Handler:    _VidmgrConfigMange_VidmgrConfigUpdate_Handler,
		},
		{
			MethodName: "vidmgrConfigDelete",
			Handler:    _VidmgrConfigMange_VidmgrConfigDelete_Handler,
		},
		{
			MethodName: "vidmgrConfigIndex",
			Handler:    _VidmgrConfigMange_VidmgrConfigIndex_Handler,
		},
		{
			MethodName: "vidmgrConfigRead",
			Handler:    _VidmgrConfigMange_VidmgrConfigRead_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/vid.proto",
}
