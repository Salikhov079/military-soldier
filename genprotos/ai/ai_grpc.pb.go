// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: ai.proto

package ai

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

// AiServiceClient is the client API for AiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AiServiceClient interface {
	CHat(ctx context.Context, in *AiCHat, opts ...grpc.CallOption) (*AiCHat, error)
}

type aiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAiServiceClient(cc grpc.ClientConnInterface) AiServiceClient {
	return &aiServiceClient{cc}
}

func (c *aiServiceClient) CHat(ctx context.Context, in *AiCHat, opts ...grpc.CallOption) (*AiCHat, error) {
	out := new(AiCHat)
	err := c.cc.Invoke(ctx, "/AI.AiService/CHat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AiServiceServer is the server API for AiService service.
// All implementations must embed UnimplementedAiServiceServer
// for forward compatibility
type AiServiceServer interface {
	CHat(context.Context, *AiCHat) (*AiCHat, error)
	mustEmbedUnimplementedAiServiceServer()
}

// UnimplementedAiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAiServiceServer struct {
}

func (UnimplementedAiServiceServer) CHat(context.Context, *AiCHat) (*AiCHat, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CHat not implemented")
}
func (UnimplementedAiServiceServer) mustEmbedUnimplementedAiServiceServer() {}

// UnsafeAiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AiServiceServer will
// result in compilation errors.
type UnsafeAiServiceServer interface {
	mustEmbedUnimplementedAiServiceServer()
}

func RegisterAiServiceServer(s grpc.ServiceRegistrar, srv AiServiceServer) {
	s.RegisterService(&AiService_ServiceDesc, srv)
}

func _AiService_CHat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AiCHat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AiServiceServer).CHat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AI.AiService/CHat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AiServiceServer).CHat(ctx, req.(*AiCHat))
	}
	return interceptor(ctx, in, info, handler)
}

// AiService_ServiceDesc is the grpc.ServiceDesc for AiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AI.AiService",
	HandlerType: (*AiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CHat",
			Handler:    _AiService_CHat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ai.proto",
}
