// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: blockchain.proto

package pactus

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

// BlockchainClient is the client API for Blockchain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainClient interface {
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error)
	GetBlockHash(ctx context.Context, in *GetBlockHashRequest, opts ...grpc.CallOption) (*GetBlockHashResponse, error)
	GetBlockHeight(ctx context.Context, in *GetBlockHeightRequest, opts ...grpc.CallOption) (*GetBlockHeightResponse, error)
	GetBlockchainInfo(ctx context.Context, in *GetBlockchainInfoRequest, opts ...grpc.CallOption) (*GetBlockchainInfoResponse, error)
	GetConsensusInfo(ctx context.Context, in *GetConsensusInfoRequest, opts ...grpc.CallOption) (*GetConsensusInfoResponse, error)
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error)
	GetValidator(ctx context.Context, in *GetValidatorRequest, opts ...grpc.CallOption) (*GetValidatorResponse, error)
	GetValidatorByNumber(ctx context.Context, in *GetValidatorByNumberRequest, opts ...grpc.CallOption) (*GetValidatorResponse, error)
	GetValidatorAddresses(ctx context.Context, in *GetValidatorAddressesRequest, opts ...grpc.CallOption) (*GetValidatorAddressesResponse, error)
	GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error)
}

type blockchainClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainClient(cc grpc.ClientConnInterface) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error) {
	out := new(GetBlockResponse)
	err := c.cc.Invoke(ctx, "/pactus.Blockchain/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetBlockHash(ctx context.Context, in *GetBlockHashRequest, opts ...grpc.CallOption) (*GetBlockHashResponse, error) {
	out := new(GetBlockHashResponse)
	err := c.cc.Invoke(ctx, "/pactus.Blockchain/GetBlockHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetBlockHeight(ctx context.Context, in *GetBlockHeightRequest, opts ...grpc.CallOption) (*GetBlockHeightResponse, error) {
	out := new(GetBlockHeightResponse)
	err := c.cc.Invoke(ctx, "/pactus.Blockchain/GetBlockHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetBlockchainInfo(ctx context.Context, in *GetBlockchainInfoRequest, opts ...grpc.CallOption) (*GetBlockchainInfoResponse, error) {
	out := new(GetBlockchainInfoResponse)
	err := c.cc.Invoke(ctx, "/pactus.Blockchain/GetBlockchainInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetConsensusInfo(ctx context.Context, in *GetConsensusInfoRequest, opts ...grpc.CallOption) (*GetConsensusInfoResponse, error) {
	out := new(GetConsensusInfoResponse)
	err := c.cc.Invoke(ctx, "/pactus.Blockchain/GetConsensusInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...grpc.CallOption) (*GetAccountResponse, error) {
	out := new(GetAccountResponse)
	err := c.cc.Invoke(ctx, "/pactus.Blockchain/GetAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetValidator(ctx context.Context, in *GetValidatorRequest, opts ...grpc.CallOption) (*GetValidatorResponse, error) {
	out := new(GetValidatorResponse)
	err := c.cc.Invoke(ctx, "/pactus.Blockchain/GetValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetValidatorByNumber(ctx context.Context, in *GetValidatorByNumberRequest, opts ...grpc.CallOption) (*GetValidatorResponse, error) {
	out := new(GetValidatorResponse)
	err := c.cc.Invoke(ctx, "/pactus.Blockchain/GetValidatorByNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetValidatorAddresses(ctx context.Context, in *GetValidatorAddressesRequest, opts ...grpc.CallOption) (*GetValidatorAddressesResponse, error) {
	out := new(GetValidatorAddressesResponse)
	err := c.cc.Invoke(ctx, "/pactus.Blockchain/GetValidatorAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetPublicKey(ctx context.Context, in *GetPublicKeyRequest, opts ...grpc.CallOption) (*GetPublicKeyResponse, error) {
	out := new(GetPublicKeyResponse)
	err := c.cc.Invoke(ctx, "/pactus.Blockchain/GetPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServer is the server API for Blockchain service.
// All implementations should embed UnimplementedBlockchainServer
// for forward compatibility
type BlockchainServer interface {
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error)
	GetBlockHash(context.Context, *GetBlockHashRequest) (*GetBlockHashResponse, error)
	GetBlockHeight(context.Context, *GetBlockHeightRequest) (*GetBlockHeightResponse, error)
	GetBlockchainInfo(context.Context, *GetBlockchainInfoRequest) (*GetBlockchainInfoResponse, error)
	GetConsensusInfo(context.Context, *GetConsensusInfoRequest) (*GetConsensusInfoResponse, error)
	GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error)
	GetValidator(context.Context, *GetValidatorRequest) (*GetValidatorResponse, error)
	GetValidatorByNumber(context.Context, *GetValidatorByNumberRequest) (*GetValidatorResponse, error)
	GetValidatorAddresses(context.Context, *GetValidatorAddressesRequest) (*GetValidatorAddressesResponse, error)
	GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error)
}

// UnimplementedBlockchainServer should be embedded to have forward compatible implementations.
type UnimplementedBlockchainServer struct {
}

func (UnimplementedBlockchainServer) GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedBlockchainServer) GetBlockHash(context.Context, *GetBlockHashRequest) (*GetBlockHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHash not implemented")
}
func (UnimplementedBlockchainServer) GetBlockHeight(context.Context, *GetBlockHeightRequest) (*GetBlockHeightResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeight not implemented")
}
func (UnimplementedBlockchainServer) GetBlockchainInfo(context.Context, *GetBlockchainInfoRequest) (*GetBlockchainInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockchainInfo not implemented")
}
func (UnimplementedBlockchainServer) GetConsensusInfo(context.Context, *GetConsensusInfoRequest) (*GetConsensusInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsensusInfo not implemented")
}
func (UnimplementedBlockchainServer) GetAccount(context.Context, *GetAccountRequest) (*GetAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedBlockchainServer) GetValidator(context.Context, *GetValidatorRequest) (*GetValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidator not implemented")
}
func (UnimplementedBlockchainServer) GetValidatorByNumber(context.Context, *GetValidatorByNumberRequest) (*GetValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorByNumber not implemented")
}
func (UnimplementedBlockchainServer) GetValidatorAddresses(context.Context, *GetValidatorAddressesRequest) (*GetValidatorAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetValidatorAddresses not implemented")
}
func (UnimplementedBlockchainServer) GetPublicKey(context.Context, *GetPublicKeyRequest) (*GetPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicKey not implemented")
}

// UnsafeBlockchainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockchainServer will
// result in compilation errors.
type UnsafeBlockchainServer interface {
	mustEmbedUnimplementedBlockchainServer()
}

func RegisterBlockchainServer(s grpc.ServiceRegistrar, srv BlockchainServer) {
	s.RegisterService(&Blockchain_ServiceDesc, srv)
}

func _Blockchain_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Blockchain/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetBlockHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlockHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Blockchain/GetBlockHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlockHash(ctx, req.(*GetBlockHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockHeightRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Blockchain/GetBlockHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlockHeight(ctx, req.(*GetBlockHeightRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetBlockchainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockchainInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlockchainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Blockchain/GetBlockchainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlockchainInfo(ctx, req.(*GetBlockchainInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetConsensusInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsensusInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetConsensusInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Blockchain/GetConsensusInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetConsensusInfo(ctx, req.(*GetConsensusInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Blockchain/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetAccount(ctx, req.(*GetAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Blockchain/GetValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetValidator(ctx, req.(*GetValidatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetValidatorByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorByNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetValidatorByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Blockchain/GetValidatorByNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetValidatorByNumber(ctx, req.(*GetValidatorByNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetValidatorAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetValidatorAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetValidatorAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Blockchain/GetValidatorAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetValidatorAddresses(ctx, req.(*GetValidatorAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pactus.Blockchain/GetPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetPublicKey(ctx, req.(*GetPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Blockchain_ServiceDesc is the grpc.ServiceDesc for Blockchain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blockchain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pactus.Blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlock",
			Handler:    _Blockchain_GetBlock_Handler,
		},
		{
			MethodName: "GetBlockHash",
			Handler:    _Blockchain_GetBlockHash_Handler,
		},
		{
			MethodName: "GetBlockHeight",
			Handler:    _Blockchain_GetBlockHeight_Handler,
		},
		{
			MethodName: "GetBlockchainInfo",
			Handler:    _Blockchain_GetBlockchainInfo_Handler,
		},
		{
			MethodName: "GetConsensusInfo",
			Handler:    _Blockchain_GetConsensusInfo_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Blockchain_GetAccount_Handler,
		},
		{
			MethodName: "GetValidator",
			Handler:    _Blockchain_GetValidator_Handler,
		},
		{
			MethodName: "GetValidatorByNumber",
			Handler:    _Blockchain_GetValidatorByNumber_Handler,
		},
		{
			MethodName: "GetValidatorAddresses",
			Handler:    _Blockchain_GetValidatorAddresses_Handler,
		},
		{
			MethodName: "GetPublicKey",
			Handler:    _Blockchain_GetPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockchain.proto",
}
