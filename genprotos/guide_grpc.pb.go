// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: staffer-protos/guide.proto

package genprotos

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

// GuideServiceClient is the client API for GuideService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GuideServiceClient interface {
	CreateGuide(ctx context.Context, in *CreateGuideRequest, opts ...grpc.CallOption) (*Void, error)
	GetGuide(ctx context.Context, in *GetGuideRequest, opts ...grpc.CallOption) (*GuideResponse, error)
	UpdateGuide(ctx context.Context, in *UpdateGuideRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteGuide(ctx context.Context, in *DeleteGuideRequest, opts ...grpc.CallOption) (*Void, error)
	ListAllGuides(ctx context.Context, in *ListAllGuidesRequest, opts ...grpc.CallOption) (*ListAllGuidesResponse, error)
	SearchGuides(ctx context.Context, in *SearchGuidesRequest, opts ...grpc.CallOption) (*ListAllGuidesResponse, error)
}

type guideServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGuideServiceClient(cc grpc.ClientConnInterface) GuideServiceClient {
	return &guideServiceClient{cc}
}

func (c *guideServiceClient) CreateGuide(ctx context.Context, in *CreateGuideRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/staff.GuideService/CreateGuide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideServiceClient) GetGuide(ctx context.Context, in *GetGuideRequest, opts ...grpc.CallOption) (*GuideResponse, error) {
	out := new(GuideResponse)
	err := c.cc.Invoke(ctx, "/staff.GuideService/GetGuide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideServiceClient) UpdateGuide(ctx context.Context, in *UpdateGuideRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/staff.GuideService/UpdateGuide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideServiceClient) DeleteGuide(ctx context.Context, in *DeleteGuideRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/staff.GuideService/DeleteGuide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideServiceClient) ListAllGuides(ctx context.Context, in *ListAllGuidesRequest, opts ...grpc.CallOption) (*ListAllGuidesResponse, error) {
	out := new(ListAllGuidesResponse)
	err := c.cc.Invoke(ctx, "/staff.GuideService/ListAllGuides", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideServiceClient) SearchGuides(ctx context.Context, in *SearchGuidesRequest, opts ...grpc.CallOption) (*ListAllGuidesResponse, error) {
	out := new(ListAllGuidesResponse)
	err := c.cc.Invoke(ctx, "/staff.GuideService/SearchGuides", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuideServiceServer is the server API for GuideService service.
// All implementations must embed UnimplementedGuideServiceServer
// for forward compatibility
type GuideServiceServer interface {
	CreateGuide(context.Context, *CreateGuideRequest) (*Void, error)
	GetGuide(context.Context, *GetGuideRequest) (*GuideResponse, error)
	UpdateGuide(context.Context, *UpdateGuideRequest) (*Void, error)
	DeleteGuide(context.Context, *DeleteGuideRequest) (*Void, error)
	ListAllGuides(context.Context, *ListAllGuidesRequest) (*ListAllGuidesResponse, error)
	SearchGuides(context.Context, *SearchGuidesRequest) (*ListAllGuidesResponse, error)
	mustEmbedUnimplementedGuideServiceServer()
}

// UnimplementedGuideServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGuideServiceServer struct {
}

func (UnimplementedGuideServiceServer) CreateGuide(context.Context, *CreateGuideRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGuide not implemented")
}
func (UnimplementedGuideServiceServer) GetGuide(context.Context, *GetGuideRequest) (*GuideResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuide not implemented")
}
func (UnimplementedGuideServiceServer) UpdateGuide(context.Context, *UpdateGuideRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGuide not implemented")
}
func (UnimplementedGuideServiceServer) DeleteGuide(context.Context, *DeleteGuideRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGuide not implemented")
}
func (UnimplementedGuideServiceServer) ListAllGuides(context.Context, *ListAllGuidesRequest) (*ListAllGuidesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllGuides not implemented")
}
func (UnimplementedGuideServiceServer) SearchGuides(context.Context, *SearchGuidesRequest) (*ListAllGuidesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchGuides not implemented")
}
func (UnimplementedGuideServiceServer) mustEmbedUnimplementedGuideServiceServer() {}

// UnsafeGuideServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GuideServiceServer will
// result in compilation errors.
type UnsafeGuideServiceServer interface {
	mustEmbedUnimplementedGuideServiceServer()
}

func RegisterGuideServiceServer(s grpc.ServiceRegistrar, srv GuideServiceServer) {
	s.RegisterService(&GuideService_ServiceDesc, srv)
}

func _GuideService_CreateGuide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGuideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).CreateGuide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.GuideService/CreateGuide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).CreateGuide(ctx, req.(*CreateGuideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideService_GetGuide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGuideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).GetGuide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.GuideService/GetGuide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).GetGuide(ctx, req.(*GetGuideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideService_UpdateGuide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGuideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).UpdateGuide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.GuideService/UpdateGuide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).UpdateGuide(ctx, req.(*UpdateGuideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideService_DeleteGuide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGuideRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).DeleteGuide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.GuideService/DeleteGuide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).DeleteGuide(ctx, req.(*DeleteGuideRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideService_ListAllGuides_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllGuidesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).ListAllGuides(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.GuideService/ListAllGuides",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).ListAllGuides(ctx, req.(*ListAllGuidesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideService_SearchGuides_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchGuidesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideServiceServer).SearchGuides(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/staff.GuideService/SearchGuides",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideServiceServer).SearchGuides(ctx, req.(*SearchGuidesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GuideService_ServiceDesc is the grpc.ServiceDesc for GuideService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GuideService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "staff.GuideService",
	HandlerType: (*GuideServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGuide",
			Handler:    _GuideService_CreateGuide_Handler,
		},
		{
			MethodName: "GetGuide",
			Handler:    _GuideService_GetGuide_Handler,
		},
		{
			MethodName: "UpdateGuide",
			Handler:    _GuideService_UpdateGuide_Handler,
		},
		{
			MethodName: "DeleteGuide",
			Handler:    _GuideService_DeleteGuide_Handler,
		},
		{
			MethodName: "ListAllGuides",
			Handler:    _GuideService_ListAllGuides_Handler,
		},
		{
			MethodName: "SearchGuides",
			Handler:    _GuideService_SearchGuides_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staffer-protos/guide.proto",
}
