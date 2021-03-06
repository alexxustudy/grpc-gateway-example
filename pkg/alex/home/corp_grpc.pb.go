// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package home

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

// CropServiceClient is the client API for CropService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CropServiceClient interface {
	//扫码
	GetCorpTicket(ctx context.Context, in *GetCorpTicketRequest, opts ...grpc.CallOption) (*GetCorpTicketResponse, error)
	PostCorpTicket(ctx context.Context, in *PostCorpTicketRequest, opts ...grpc.CallOption) (*PostCorpTicketResponse, error)
}

type cropServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCropServiceClient(cc grpc.ClientConnInterface) CropServiceClient {
	return &cropServiceClient{cc}
}

func (c *cropServiceClient) GetCorpTicket(ctx context.Context, in *GetCorpTicketRequest, opts ...grpc.CallOption) (*GetCorpTicketResponse, error) {
	out := new(GetCorpTicketResponse)
	err := c.cc.Invoke(ctx, "/xd.home.CropService/getCorpTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cropServiceClient) PostCorpTicket(ctx context.Context, in *PostCorpTicketRequest, opts ...grpc.CallOption) (*PostCorpTicketResponse, error) {
	out := new(PostCorpTicketResponse)
	err := c.cc.Invoke(ctx, "/xd.home.CropService/postCorpTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CropServiceServer is the server API for CropService service.
// All implementations must embed UnimplementedCropServiceServer
// for forward compatibility
type CropServiceServer interface {
	//扫码
	GetCorpTicket(context.Context, *GetCorpTicketRequest) (*GetCorpTicketResponse, error)
	PostCorpTicket(context.Context, *PostCorpTicketRequest) (*PostCorpTicketResponse, error)
	mustEmbedUnimplementedCropServiceServer()
}

// UnimplementedCropServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCropServiceServer struct {
}

func (UnimplementedCropServiceServer) GetCorpTicket(context.Context, *GetCorpTicketRequest) (*GetCorpTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCorpTicket not implemented")
}
func (UnimplementedCropServiceServer) PostCorpTicket(context.Context, *PostCorpTicketRequest) (*PostCorpTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCorpTicket not implemented")
}
func (UnimplementedCropServiceServer) mustEmbedUnimplementedCropServiceServer() {}

// UnsafeCropServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CropServiceServer will
// result in compilation errors.
type UnsafeCropServiceServer interface {
	mustEmbedUnimplementedCropServiceServer()
}

func RegisterCropServiceServer(s grpc.ServiceRegistrar, srv CropServiceServer) {
	s.RegisterService(&CropService_ServiceDesc, srv)
}

func _CropService_GetCorpTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCorpTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CropServiceServer).GetCorpTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xd.home.CropService/getCorpTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CropServiceServer).GetCorpTicket(ctx, req.(*GetCorpTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CropService_PostCorpTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCorpTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CropServiceServer).PostCorpTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xd.home.CropService/postCorpTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CropServiceServer).PostCorpTicket(ctx, req.(*PostCorpTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CropService_ServiceDesc is the grpc.ServiceDesc for CropService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CropService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xd.home.CropService",
	HandlerType: (*CropServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCorpTicket",
			Handler:    _CropService_GetCorpTicket_Handler,
		},
		{
			MethodName: "postCorpTicket",
			Handler:    _CropService_PostCorpTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "home/corp.proto",
}
