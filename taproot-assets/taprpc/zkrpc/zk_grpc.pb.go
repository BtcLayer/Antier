// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: zkrpc/zk.proto

package zkrpc

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

// ZKClient is the client API for ZK service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ZKClient interface {
	// tapcli: `zk create`
	CreateProof(ctx context.Context, in *CreateProofRequest, opts ...grpc.CallOption) (*CreateProofResponse, error)
	// tapcli: `zk create finalize`
	// FinalizeBatch will attempt to finalize the current pending batch.
	FinalizeBatch(ctx context.Context, in *FinalizeBatchRequest, opts ...grpc.CallOption) (*FinalizeBatchResponse, error)
	// tapcli: `zk create cancel`
	// CancelBatch will attempt to cancel the current pending batch.
	CancelBatch(ctx context.Context, in *CancelBatchRequest, opts ...grpc.CallOption) (*CancelBatchResponse, error)
	// tapcli: `zk create batches`
	// ListBatches lists the set of batches submitted to the daemon, including
	// pending and cancelled batches.
	ListBatches(ctx context.Context, in *ListBatchRequest, opts ...grpc.CallOption) (*ListBatchResponse, error)
}

type zKClient struct {
	cc grpc.ClientConnInterface
}

func NewZKClient(cc grpc.ClientConnInterface) ZKClient {
	return &zKClient{cc}
}

func (c *zKClient) CreateProof(ctx context.Context, in *CreateProofRequest, opts ...grpc.CallOption) (*CreateProofResponse, error) {
	out := new(CreateProofResponse)
	err := c.cc.Invoke(ctx, "/zkrpc.ZK/CreateProof", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zKClient) FinalizeBatch(ctx context.Context, in *FinalizeBatchRequest, opts ...grpc.CallOption) (*FinalizeBatchResponse, error) {
	out := new(FinalizeBatchResponse)
	err := c.cc.Invoke(ctx, "/zkrpc.ZK/FinalizeBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zKClient) CancelBatch(ctx context.Context, in *CancelBatchRequest, opts ...grpc.CallOption) (*CancelBatchResponse, error) {
	out := new(CancelBatchResponse)
	err := c.cc.Invoke(ctx, "/zkrpc.ZK/CancelBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zKClient) ListBatches(ctx context.Context, in *ListBatchRequest, opts ...grpc.CallOption) (*ListBatchResponse, error) {
	out := new(ListBatchResponse)
	err := c.cc.Invoke(ctx, "/zkrpc.ZK/ListBatches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZKServer is the server API for ZK service.
// All implementations must embed UnimplementedZKServer
// for forward compatibility
type ZKServer interface {
	// tapcli: `zk create`
	CreateProof(context.Context, *CreateProofRequest) (*CreateProofResponse, error)
	// tapcli: `zk create finalize`
	// FinalizeBatch will attempt to finalize the current pending batch.
	FinalizeBatch(context.Context, *FinalizeBatchRequest) (*FinalizeBatchResponse, error)
	// tapcli: `zk create cancel`
	// CancelBatch will attempt to cancel the current pending batch.
	CancelBatch(context.Context, *CancelBatchRequest) (*CancelBatchResponse, error)
	// tapcli: `zk create batches`
	// ListBatches lists the set of batches submitted to the daemon, including
	// pending and cancelled batches.
	ListBatches(context.Context, *ListBatchRequest) (*ListBatchResponse, error)
	mustEmbedUnimplementedZKServer()
}

// UnimplementedZKServer must be embedded to have forward compatible implementations.
type UnimplementedZKServer struct {
}

func (UnimplementedZKServer) CreateProof(context.Context, *CreateProofRequest) (*CreateProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProof not implemented")
}
func (UnimplementedZKServer) FinalizeBatch(context.Context, *FinalizeBatchRequest) (*FinalizeBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizeBatch not implemented")
}
func (UnimplementedZKServer) CancelBatch(context.Context, *CancelBatchRequest) (*CancelBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelBatch not implemented")
}
func (UnimplementedZKServer) ListBatches(context.Context, *ListBatchRequest) (*ListBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBatches not implemented")
}
func (UnimplementedZKServer) mustEmbedUnimplementedZKServer() {}

// UnsafeZKServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZKServer will
// result in compilation errors.
type UnsafeZKServer interface {
	mustEmbedUnimplementedZKServer()
}

func RegisterZKServer(s grpc.ServiceRegistrar, srv ZKServer) {
	s.RegisterService(&ZK_ServiceDesc, srv)
}

func _ZK_CreateProof_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZKServer).CreateProof(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zkrpc.ZK/CreateProof",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZKServer).CreateProof(ctx, req.(*CreateProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZK_FinalizeBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZKServer).FinalizeBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zkrpc.ZK/FinalizeBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZKServer).FinalizeBatch(ctx, req.(*FinalizeBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZK_CancelBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZKServer).CancelBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zkrpc.ZK/CancelBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZKServer).CancelBatch(ctx, req.(*CancelBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZK_ListBatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZKServer).ListBatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zkrpc.ZK/ListBatches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZKServer).ListBatches(ctx, req.(*ListBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ZK_ServiceDesc is the grpc.ServiceDesc for ZK service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZK_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zkrpc.ZK",
	HandlerType: (*ZKServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProof",
			Handler:    _ZK_CreateProof_Handler,
		},
		{
			MethodName: "FinalizeBatch",
			Handler:    _ZK_FinalizeBatch_Handler,
		},
		{
			MethodName: "CancelBatch",
			Handler:    _ZK_CancelBatch_Handler,
		},
		{
			MethodName: "ListBatches",
			Handler:    _ZK_ListBatches_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zkrpc/zk.proto",
}