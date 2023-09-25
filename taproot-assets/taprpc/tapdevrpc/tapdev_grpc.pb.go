// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: tapdevrpc/tapdev.proto

package tapdevrpc

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

// TapDevClient is the client API for TapDev service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TapDevClient interface {
	// tapcli: `dev importproof`
	// ImportProof attempts to import a proof file into the daemon. If successful,
	// a new asset will be inserted on disk, spendable using the specified target
	// script key, and internal key.
	ImportProof(ctx context.Context, in *ImportProofRequest, opts ...grpc.CallOption) (*ImportProofResponse, error)
}

type tapDevClient struct {
	cc grpc.ClientConnInterface
}

func NewTapDevClient(cc grpc.ClientConnInterface) TapDevClient {
	return &tapDevClient{cc}
}

func (c *tapDevClient) ImportProof(ctx context.Context, in *ImportProofRequest, opts ...grpc.CallOption) (*ImportProofResponse, error) {
	out := new(ImportProofResponse)
	err := c.cc.Invoke(ctx, "/tapdevrpc.TapDev/ImportProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TapDevServer is the server API for TapDev service.
// All implementations must embed UnimplementedTapDevServer
// for forward compatibility
type TapDevServer interface {
	// tapcli: `dev importproof`
	// ImportProof attempts to import a proof file into the daemon. If successful,
	// a new asset will be inserted on disk, spendable using the specified target
	// script key, and internal key.
	ImportProof(context.Context, *ImportProofRequest) (*ImportProofResponse, error)
	mustEmbedUnimplementedTapDevServer()
}

// UnimplementedTapDevServer must be embedded to have forward compatible implementations.
type UnimplementedTapDevServer struct {
}

func (UnimplementedTapDevServer) ImportProof(context.Context, *ImportProofRequest) (*ImportProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportProof not implemented")
}
func (UnimplementedTapDevServer) mustEmbedUnimplementedTapDevServer() {}

// UnsafeTapDevServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TapDevServer will
// result in compilation errors.
type UnsafeTapDevServer interface {
	mustEmbedUnimplementedTapDevServer()
}

func RegisterTapDevServer(s grpc.ServiceRegistrar, srv TapDevServer) {
	s.RegisterService(&TapDev_ServiceDesc, srv)
}

func _TapDev_ImportProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TapDevServer).ImportProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tapdevrpc.TapDev/ImportProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TapDevServer).ImportProof(ctx, req.(*ImportProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TapDev_ServiceDesc is the grpc.ServiceDesc for TapDev service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TapDev_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tapdevrpc.TapDev",
	HandlerType: (*TapDevServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportProof",
			Handler:    _TapDev_ImportProof_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tapdevrpc/tapdev.proto",
}
