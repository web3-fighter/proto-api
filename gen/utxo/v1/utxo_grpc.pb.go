// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: dapplink/utxo.proto

package v1

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
	WalletUtxoService_GetSupportChains_FullMethodName        = "/dapplink.utxo.WalletUtxoService/getSupportChains"
	WalletUtxoService_ConvertAddress_FullMethodName          = "/dapplink.utxo.WalletUtxoService/convertAddress"
	WalletUtxoService_ValidAddress_FullMethodName            = "/dapplink.utxo.WalletUtxoService/validAddress"
	WalletUtxoService_GetFee_FullMethodName                  = "/dapplink.utxo.WalletUtxoService/getFee"
	WalletUtxoService_GetAccount_FullMethodName              = "/dapplink.utxo.WalletUtxoService/getAccount"
	WalletUtxoService_GetUnspentOutputs_FullMethodName       = "/dapplink.utxo.WalletUtxoService/getUnspentOutputs"
	WalletUtxoService_GetBlockByNumber_FullMethodName        = "/dapplink.utxo.WalletUtxoService/getBlockByNumber"
	WalletUtxoService_GetBlockByHash_FullMethodName          = "/dapplink.utxo.WalletUtxoService/getBlockByHash"
	WalletUtxoService_GetBlockHeaderByHash_FullMethodName    = "/dapplink.utxo.WalletUtxoService/getBlockHeaderByHash"
	WalletUtxoService_GetBlockHeaderByNumber_FullMethodName  = "/dapplink.utxo.WalletUtxoService/getBlockHeaderByNumber"
	WalletUtxoService_SendTx_FullMethodName                  = "/dapplink.utxo.WalletUtxoService/SendTx"
	WalletUtxoService_GetTxByAddress_FullMethodName          = "/dapplink.utxo.WalletUtxoService/getTxByAddress"
	WalletUtxoService_GetTxByHash_FullMethodName             = "/dapplink.utxo.WalletUtxoService/getTxByHash"
	WalletUtxoService_CreateUnSignTransaction_FullMethodName = "/dapplink.utxo.WalletUtxoService/createUnSignTransaction"
	WalletUtxoService_BuildSignedTransaction_FullMethodName  = "/dapplink.utxo.WalletUtxoService/buildSignedTransaction"
	WalletUtxoService_DecodeTransaction_FullMethodName       = "/dapplink.utxo.WalletUtxoService/decodeTransaction"
	WalletUtxoService_VerifySignedTransaction_FullMethodName = "/dapplink.utxo.WalletUtxoService/verifySignedTransaction"
)

// WalletUtxoServiceClient is the client API for WalletUtxoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WalletUtxoServiceClient interface {
	GetSupportChains(ctx context.Context, in *SupportChainsRequest, opts ...grpc.CallOption) (*SupportChainsResponse, error)
	ConvertAddress(ctx context.Context, in *ConvertAddressRequest, opts ...grpc.CallOption) (*ConvertAddressResponse, error)
	ValidAddress(ctx context.Context, in *ValidAddressRequest, opts ...grpc.CallOption) (*ValidAddressResponse, error)
	GetFee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error)
	GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	GetUnspentOutputs(ctx context.Context, in *UnspentOutputsRequest, opts ...grpc.CallOption) (*UnspentOutputsResponse, error)
	GetBlockByNumber(ctx context.Context, in *BlockNumberRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	GetBlockByHash(ctx context.Context, in *BlockHashRequest, opts ...grpc.CallOption) (*BlockResponse, error)
	GetBlockHeaderByHash(ctx context.Context, in *BlockHeaderHashRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
	GetBlockHeaderByNumber(ctx context.Context, in *BlockHeaderNumberRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error)
	SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*SendTxResponse, error)
	GetTxByAddress(ctx context.Context, in *TxAddressRequest, opts ...grpc.CallOption) (*TxAddressResponse, error)
	GetTxByHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*TxHashResponse, error)
	CreateUnSignTransaction(ctx context.Context, in *UnSignTransactionRequest, opts ...grpc.CallOption) (*UnSignTransactionResponse, error)
	BuildSignedTransaction(ctx context.Context, in *SignedTransactionRequest, opts ...grpc.CallOption) (*SignedTransactionResponse, error)
	DecodeTransaction(ctx context.Context, in *DecodeTransactionRequest, opts ...grpc.CallOption) (*DecodeTransactionResponse, error)
	VerifySignedTransaction(ctx context.Context, in *VerifyTransactionRequest, opts ...grpc.CallOption) (*VerifyTransactionResponse, error)
}

type walletUtxoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWalletUtxoServiceClient(cc grpc.ClientConnInterface) WalletUtxoServiceClient {
	return &walletUtxoServiceClient{cc}
}

func (c *walletUtxoServiceClient) GetSupportChains(ctx context.Context, in *SupportChainsRequest, opts ...grpc.CallOption) (*SupportChainsResponse, error) {
	out := new(SupportChainsResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_GetSupportChains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) ConvertAddress(ctx context.Context, in *ConvertAddressRequest, opts ...grpc.CallOption) (*ConvertAddressResponse, error) {
	out := new(ConvertAddressResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_ConvertAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) ValidAddress(ctx context.Context, in *ValidAddressRequest, opts ...grpc.CallOption) (*ValidAddressResponse, error) {
	out := new(ValidAddressResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_ValidAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) GetFee(ctx context.Context, in *FeeRequest, opts ...grpc.CallOption) (*FeeResponse, error) {
	out := new(FeeResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_GetFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_GetAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) GetUnspentOutputs(ctx context.Context, in *UnspentOutputsRequest, opts ...grpc.CallOption) (*UnspentOutputsResponse, error) {
	out := new(UnspentOutputsResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_GetUnspentOutputs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) GetBlockByNumber(ctx context.Context, in *BlockNumberRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_GetBlockByNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) GetBlockByHash(ctx context.Context, in *BlockHashRequest, opts ...grpc.CallOption) (*BlockResponse, error) {
	out := new(BlockResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_GetBlockByHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) GetBlockHeaderByHash(ctx context.Context, in *BlockHeaderHashRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_GetBlockHeaderByHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) GetBlockHeaderByNumber(ctx context.Context, in *BlockHeaderNumberRequest, opts ...grpc.CallOption) (*BlockHeaderResponse, error) {
	out := new(BlockHeaderResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_GetBlockHeaderByNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*SendTxResponse, error) {
	out := new(SendTxResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_SendTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) GetTxByAddress(ctx context.Context, in *TxAddressRequest, opts ...grpc.CallOption) (*TxAddressResponse, error) {
	out := new(TxAddressResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_GetTxByAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) GetTxByHash(ctx context.Context, in *TxHashRequest, opts ...grpc.CallOption) (*TxHashResponse, error) {
	out := new(TxHashResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_GetTxByHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) CreateUnSignTransaction(ctx context.Context, in *UnSignTransactionRequest, opts ...grpc.CallOption) (*UnSignTransactionResponse, error) {
	out := new(UnSignTransactionResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_CreateUnSignTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) BuildSignedTransaction(ctx context.Context, in *SignedTransactionRequest, opts ...grpc.CallOption) (*SignedTransactionResponse, error) {
	out := new(SignedTransactionResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_BuildSignedTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) DecodeTransaction(ctx context.Context, in *DecodeTransactionRequest, opts ...grpc.CallOption) (*DecodeTransactionResponse, error) {
	out := new(DecodeTransactionResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_DecodeTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletUtxoServiceClient) VerifySignedTransaction(ctx context.Context, in *VerifyTransactionRequest, opts ...grpc.CallOption) (*VerifyTransactionResponse, error) {
	out := new(VerifyTransactionResponse)
	err := c.cc.Invoke(ctx, WalletUtxoService_VerifySignedTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletUtxoServiceServer is the server API for WalletUtxoService service.
// All implementations should embed UnimplementedWalletUtxoServiceServer
// for forward compatibility
type WalletUtxoServiceServer interface {
	GetSupportChains(context.Context, *SupportChainsRequest) (*SupportChainsResponse, error)
	ConvertAddress(context.Context, *ConvertAddressRequest) (*ConvertAddressResponse, error)
	ValidAddress(context.Context, *ValidAddressRequest) (*ValidAddressResponse, error)
	GetFee(context.Context, *FeeRequest) (*FeeResponse, error)
	GetAccount(context.Context, *AccountRequest) (*AccountResponse, error)
	GetUnspentOutputs(context.Context, *UnspentOutputsRequest) (*UnspentOutputsResponse, error)
	GetBlockByNumber(context.Context, *BlockNumberRequest) (*BlockResponse, error)
	GetBlockByHash(context.Context, *BlockHashRequest) (*BlockResponse, error)
	GetBlockHeaderByHash(context.Context, *BlockHeaderHashRequest) (*BlockHeaderResponse, error)
	GetBlockHeaderByNumber(context.Context, *BlockHeaderNumberRequest) (*BlockHeaderResponse, error)
	SendTx(context.Context, *SendTxRequest) (*SendTxResponse, error)
	GetTxByAddress(context.Context, *TxAddressRequest) (*TxAddressResponse, error)
	GetTxByHash(context.Context, *TxHashRequest) (*TxHashResponse, error)
	CreateUnSignTransaction(context.Context, *UnSignTransactionRequest) (*UnSignTransactionResponse, error)
	BuildSignedTransaction(context.Context, *SignedTransactionRequest) (*SignedTransactionResponse, error)
	DecodeTransaction(context.Context, *DecodeTransactionRequest) (*DecodeTransactionResponse, error)
	VerifySignedTransaction(context.Context, *VerifyTransactionRequest) (*VerifyTransactionResponse, error)
}

// UnimplementedWalletUtxoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWalletUtxoServiceServer struct {
}

func (UnimplementedWalletUtxoServiceServer) GetSupportChains(context.Context, *SupportChainsRequest) (*SupportChainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportChains not implemented")
}
func (UnimplementedWalletUtxoServiceServer) ConvertAddress(context.Context, *ConvertAddressRequest) (*ConvertAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConvertAddress not implemented")
}
func (UnimplementedWalletUtxoServiceServer) ValidAddress(context.Context, *ValidAddressRequest) (*ValidAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidAddress not implemented")
}
func (UnimplementedWalletUtxoServiceServer) GetFee(context.Context, *FeeRequest) (*FeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFee not implemented")
}
func (UnimplementedWalletUtxoServiceServer) GetAccount(context.Context, *AccountRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccount not implemented")
}
func (UnimplementedWalletUtxoServiceServer) GetUnspentOutputs(context.Context, *UnspentOutputsRequest) (*UnspentOutputsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnspentOutputs not implemented")
}
func (UnimplementedWalletUtxoServiceServer) GetBlockByNumber(context.Context, *BlockNumberRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByNumber not implemented")
}
func (UnimplementedWalletUtxoServiceServer) GetBlockByHash(context.Context, *BlockHashRequest) (*BlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockByHash not implemented")
}
func (UnimplementedWalletUtxoServiceServer) GetBlockHeaderByHash(context.Context, *BlockHeaderHashRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByHash not implemented")
}
func (UnimplementedWalletUtxoServiceServer) GetBlockHeaderByNumber(context.Context, *BlockHeaderNumberRequest) (*BlockHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockHeaderByNumber not implemented")
}
func (UnimplementedWalletUtxoServiceServer) SendTx(context.Context, *SendTxRequest) (*SendTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTx not implemented")
}
func (UnimplementedWalletUtxoServiceServer) GetTxByAddress(context.Context, *TxAddressRequest) (*TxAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByAddress not implemented")
}
func (UnimplementedWalletUtxoServiceServer) GetTxByHash(context.Context, *TxHashRequest) (*TxHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTxByHash not implemented")
}
func (UnimplementedWalletUtxoServiceServer) CreateUnSignTransaction(context.Context, *UnSignTransactionRequest) (*UnSignTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUnSignTransaction not implemented")
}
func (UnimplementedWalletUtxoServiceServer) BuildSignedTransaction(context.Context, *SignedTransactionRequest) (*SignedTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildSignedTransaction not implemented")
}
func (UnimplementedWalletUtxoServiceServer) DecodeTransaction(context.Context, *DecodeTransactionRequest) (*DecodeTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeTransaction not implemented")
}
func (UnimplementedWalletUtxoServiceServer) VerifySignedTransaction(context.Context, *VerifyTransactionRequest) (*VerifyTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySignedTransaction not implemented")
}

// UnsafeWalletUtxoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WalletUtxoServiceServer will
// result in compilation errors.
type UnsafeWalletUtxoServiceServer interface {
	mustEmbedUnimplementedWalletUtxoServiceServer()
}

func RegisterWalletUtxoServiceServer(s grpc.ServiceRegistrar, srv WalletUtxoServiceServer) {
	s.RegisterService(&WalletUtxoService_ServiceDesc, srv)
}

func _WalletUtxoService_GetSupportChains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportChainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).GetSupportChains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_GetSupportChains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).GetSupportChains(ctx, req.(*SupportChainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_ConvertAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).ConvertAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_ConvertAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).ConvertAddress(ctx, req.(*ConvertAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_ValidAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).ValidAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_ValidAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).ValidAddress(ctx, req.(*ValidAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_GetFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).GetFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_GetFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).GetFee(ctx, req.(*FeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_GetAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).GetAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_GetUnspentOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnspentOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).GetUnspentOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_GetUnspentOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).GetUnspentOutputs(ctx, req.(*UnspentOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_GetBlockByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).GetBlockByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_GetBlockByNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).GetBlockByNumber(ctx, req.(*BlockNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_GetBlockByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).GetBlockByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_GetBlockByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).GetBlockByHash(ctx, req.(*BlockHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_GetBlockHeaderByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHeaderHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).GetBlockHeaderByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_GetBlockHeaderByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).GetBlockHeaderByHash(ctx, req.(*BlockHeaderHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_GetBlockHeaderByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockHeaderNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).GetBlockHeaderByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_GetBlockHeaderByNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).GetBlockHeaderByNumber(ctx, req.(*BlockHeaderNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_SendTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).SendTx(ctx, req.(*SendTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_GetTxByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).GetTxByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_GetTxByAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).GetTxByAddress(ctx, req.(*TxAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_GetTxByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).GetTxByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_GetTxByHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).GetTxByHash(ctx, req.(*TxHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_CreateUnSignTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSignTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).CreateUnSignTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_CreateUnSignTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).CreateUnSignTransaction(ctx, req.(*UnSignTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_BuildSignedTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignedTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).BuildSignedTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_BuildSignedTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).BuildSignedTransaction(ctx, req.(*SignedTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_DecodeTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).DecodeTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_DecodeTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).DecodeTransaction(ctx, req.(*DecodeTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WalletUtxoService_VerifySignedTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletUtxoServiceServer).VerifySignedTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WalletUtxoService_VerifySignedTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletUtxoServiceServer).VerifySignedTransaction(ctx, req.(*VerifyTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WalletUtxoService_ServiceDesc is the grpc.ServiceDesc for WalletUtxoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WalletUtxoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapplink.utxo.WalletUtxoService",
	HandlerType: (*WalletUtxoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSupportChains",
			Handler:    _WalletUtxoService_GetSupportChains_Handler,
		},
		{
			MethodName: "convertAddress",
			Handler:    _WalletUtxoService_ConvertAddress_Handler,
		},
		{
			MethodName: "validAddress",
			Handler:    _WalletUtxoService_ValidAddress_Handler,
		},
		{
			MethodName: "getFee",
			Handler:    _WalletUtxoService_GetFee_Handler,
		},
		{
			MethodName: "getAccount",
			Handler:    _WalletUtxoService_GetAccount_Handler,
		},
		{
			MethodName: "getUnspentOutputs",
			Handler:    _WalletUtxoService_GetUnspentOutputs_Handler,
		},
		{
			MethodName: "getBlockByNumber",
			Handler:    _WalletUtxoService_GetBlockByNumber_Handler,
		},
		{
			MethodName: "getBlockByHash",
			Handler:    _WalletUtxoService_GetBlockByHash_Handler,
		},
		{
			MethodName: "getBlockHeaderByHash",
			Handler:    _WalletUtxoService_GetBlockHeaderByHash_Handler,
		},
		{
			MethodName: "getBlockHeaderByNumber",
			Handler:    _WalletUtxoService_GetBlockHeaderByNumber_Handler,
		},
		{
			MethodName: "SendTx",
			Handler:    _WalletUtxoService_SendTx_Handler,
		},
		{
			MethodName: "getTxByAddress",
			Handler:    _WalletUtxoService_GetTxByAddress_Handler,
		},
		{
			MethodName: "getTxByHash",
			Handler:    _WalletUtxoService_GetTxByHash_Handler,
		},
		{
			MethodName: "createUnSignTransaction",
			Handler:    _WalletUtxoService_CreateUnSignTransaction_Handler,
		},
		{
			MethodName: "buildSignedTransaction",
			Handler:    _WalletUtxoService_BuildSignedTransaction_Handler,
		},
		{
			MethodName: "decodeTransaction",
			Handler:    _WalletUtxoService_DecodeTransaction_Handler,
		},
		{
			MethodName: "verifySignedTransaction",
			Handler:    _WalletUtxoService_VerifySignedTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dapplink/utxo.proto",
}
