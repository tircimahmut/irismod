// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: irismod/token/v1/tx.proto

package tokenv1

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
	Msg_IssueToken_FullMethodName         = "/irismod.token.v1.Msg/IssueToken"
	Msg_EditToken_FullMethodName          = "/irismod.token.v1.Msg/EditToken"
	Msg_MintToken_FullMethodName          = "/irismod.token.v1.Msg/MintToken"
	Msg_BurnToken_FullMethodName          = "/irismod.token.v1.Msg/BurnToken"
	Msg_TransferTokenOwner_FullMethodName = "/irismod.token.v1.Msg/TransferTokenOwner"
	Msg_SwapFeeToken_FullMethodName       = "/irismod.token.v1.Msg/SwapFeeToken"
	Msg_SwapToERC20_FullMethodName        = "/irismod.token.v1.Msg/SwapToERC20"
	Msg_SwapFromERC20_FullMethodName      = "/irismod.token.v1.Msg/SwapFromERC20"
	Msg_UpdateParams_FullMethodName       = "/irismod.token.v1.Msg/UpdateParams"
	Msg_DeployERC20_FullMethodName        = "/irismod.token.v1.Msg/DeployERC20"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// IssueToken defines a method for issuing a new token
	IssueToken(ctx context.Context, in *MsgIssueToken, opts ...grpc.CallOption) (*MsgIssueTokenResponse, error)
	// EditToken defines a method for editing a token
	EditToken(ctx context.Context, in *MsgEditToken, opts ...grpc.CallOption) (*MsgEditTokenResponse, error)
	// MintToken defines a method for minting some tokens
	MintToken(ctx context.Context, in *MsgMintToken, opts ...grpc.CallOption) (*MsgMintTokenResponse, error)
	// BurnToken defines a method for burning some tokens
	BurnToken(ctx context.Context, in *MsgBurnToken, opts ...grpc.CallOption) (*MsgBurnTokenResponse, error)
	// TransferTokenOwner defines a method for transfering token owner
	TransferTokenOwner(ctx context.Context, in *MsgTransferTokenOwner, opts ...grpc.CallOption) (*MsgTransferTokenOwnerResponse, error)
	// SwapFeeToken defines a method for swapping some fee token
	SwapFeeToken(ctx context.Context, in *MsgSwapFeeToken, opts ...grpc.CallOption) (*MsgSwapFeeTokenResponse, error)
	// SwapToERC20 defines a method for swapping some erc20 token
	SwapToERC20(ctx context.Context, in *MsgSwapToERC20, opts ...grpc.CallOption) (*MsgSwapToERC20Response, error)
	// SwapFromERC20 defines a method for swapping some native token
	SwapFromERC20(ctx context.Context, in *MsgSwapFromERC20, opts ...grpc.CallOption) (*MsgSwapFromERC20Response, error)
	// UpdateParams defines a governance operation for updating the token
	// module parameters. The authority is defined in the keeper.
	//
	// Since: cosmos-sdk 0.47
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// DeployERC20 defines a governance operation for deploying erc20 token for
	// the native token
	DeployERC20(ctx context.Context, in *MsgDeployERC20, opts ...grpc.CallOption) (*MsgDeployERC20Response, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) IssueToken(ctx context.Context, in *MsgIssueToken, opts ...grpc.CallOption) (*MsgIssueTokenResponse, error) {
	out := new(MsgIssueTokenResponse)
	err := c.cc.Invoke(ctx, Msg_IssueToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EditToken(ctx context.Context, in *MsgEditToken, opts ...grpc.CallOption) (*MsgEditTokenResponse, error) {
	out := new(MsgEditTokenResponse)
	err := c.cc.Invoke(ctx, Msg_EditToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MintToken(ctx context.Context, in *MsgMintToken, opts ...grpc.CallOption) (*MsgMintTokenResponse, error) {
	out := new(MsgMintTokenResponse)
	err := c.cc.Invoke(ctx, Msg_MintToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) BurnToken(ctx context.Context, in *MsgBurnToken, opts ...grpc.CallOption) (*MsgBurnTokenResponse, error) {
	out := new(MsgBurnTokenResponse)
	err := c.cc.Invoke(ctx, Msg_BurnToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) TransferTokenOwner(ctx context.Context, in *MsgTransferTokenOwner, opts ...grpc.CallOption) (*MsgTransferTokenOwnerResponse, error) {
	out := new(MsgTransferTokenOwnerResponse)
	err := c.cc.Invoke(ctx, Msg_TransferTokenOwner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapFeeToken(ctx context.Context, in *MsgSwapFeeToken, opts ...grpc.CallOption) (*MsgSwapFeeTokenResponse, error) {
	out := new(MsgSwapFeeTokenResponse)
	err := c.cc.Invoke(ctx, Msg_SwapFeeToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapToERC20(ctx context.Context, in *MsgSwapToERC20, opts ...grpc.CallOption) (*MsgSwapToERC20Response, error) {
	out := new(MsgSwapToERC20Response)
	err := c.cc.Invoke(ctx, Msg_SwapToERC20_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapFromERC20(ctx context.Context, in *MsgSwapFromERC20, opts ...grpc.CallOption) (*MsgSwapFromERC20Response, error) {
	out := new(MsgSwapFromERC20Response)
	err := c.cc.Invoke(ctx, Msg_SwapFromERC20_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeployERC20(ctx context.Context, in *MsgDeployERC20, opts ...grpc.CallOption) (*MsgDeployERC20Response, error) {
	out := new(MsgDeployERC20Response)
	err := c.cc.Invoke(ctx, Msg_DeployERC20_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// IssueToken defines a method for issuing a new token
	IssueToken(context.Context, *MsgIssueToken) (*MsgIssueTokenResponse, error)
	// EditToken defines a method for editing a token
	EditToken(context.Context, *MsgEditToken) (*MsgEditTokenResponse, error)
	// MintToken defines a method for minting some tokens
	MintToken(context.Context, *MsgMintToken) (*MsgMintTokenResponse, error)
	// BurnToken defines a method for burning some tokens
	BurnToken(context.Context, *MsgBurnToken) (*MsgBurnTokenResponse, error)
	// TransferTokenOwner defines a method for transfering token owner
	TransferTokenOwner(context.Context, *MsgTransferTokenOwner) (*MsgTransferTokenOwnerResponse, error)
	// SwapFeeToken defines a method for swapping some fee token
	SwapFeeToken(context.Context, *MsgSwapFeeToken) (*MsgSwapFeeTokenResponse, error)
	// SwapToERC20 defines a method for swapping some erc20 token
	SwapToERC20(context.Context, *MsgSwapToERC20) (*MsgSwapToERC20Response, error)
	// SwapFromERC20 defines a method for swapping some native token
	SwapFromERC20(context.Context, *MsgSwapFromERC20) (*MsgSwapFromERC20Response, error)
	// UpdateParams defines a governance operation for updating the token
	// module parameters. The authority is defined in the keeper.
	//
	// Since: cosmos-sdk 0.47
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// DeployERC20 defines a governance operation for deploying erc20 token for
	// the native token
	DeployERC20(context.Context, *MsgDeployERC20) (*MsgDeployERC20Response, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) IssueToken(context.Context, *MsgIssueToken) (*MsgIssueTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}
func (UnimplementedMsgServer) EditToken(context.Context, *MsgEditToken) (*MsgEditTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditToken not implemented")
}
func (UnimplementedMsgServer) MintToken(context.Context, *MsgMintToken) (*MsgMintTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintToken not implemented")
}
func (UnimplementedMsgServer) BurnToken(context.Context, *MsgBurnToken) (*MsgBurnTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnToken not implemented")
}
func (UnimplementedMsgServer) TransferTokenOwner(context.Context, *MsgTransferTokenOwner) (*MsgTransferTokenOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferTokenOwner not implemented")
}
func (UnimplementedMsgServer) SwapFeeToken(context.Context, *MsgSwapFeeToken) (*MsgSwapFeeTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapFeeToken not implemented")
}
func (UnimplementedMsgServer) SwapToERC20(context.Context, *MsgSwapToERC20) (*MsgSwapToERC20Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapToERC20 not implemented")
}
func (UnimplementedMsgServer) SwapFromERC20(context.Context, *MsgSwapFromERC20) (*MsgSwapFromERC20Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapFromERC20 not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) DeployERC20(context.Context, *MsgDeployERC20) (*MsgDeployERC20Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployERC20 not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgIssueToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_IssueToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).IssueToken(ctx, req.(*MsgIssueToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EditToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEditToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EditToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_EditToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EditToken(ctx, req.(*MsgEditToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MintToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MintToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintToken(ctx, req.(*MsgMintToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_BurnToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_BurnToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnToken(ctx, req.(*MsgBurnToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_TransferTokenOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferTokenOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferTokenOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_TransferTokenOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferTokenOwner(ctx, req.(*MsgTransferTokenOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapFeeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapFeeToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapFeeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SwapFeeToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapFeeToken(ctx, req.(*MsgSwapFeeToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapToERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapToERC20)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapToERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SwapToERC20_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapToERC20(ctx, req.(*MsgSwapToERC20))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapFromERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapFromERC20)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapFromERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SwapFromERC20_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapFromERC20(ctx, req.(*MsgSwapFromERC20))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeployERC20_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeployERC20)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeployERC20(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeployERC20_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeployERC20(ctx, req.(*MsgDeployERC20))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.token.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueToken",
			Handler:    _Msg_IssueToken_Handler,
		},
		{
			MethodName: "EditToken",
			Handler:    _Msg_EditToken_Handler,
		},
		{
			MethodName: "MintToken",
			Handler:    _Msg_MintToken_Handler,
		},
		{
			MethodName: "BurnToken",
			Handler:    _Msg_BurnToken_Handler,
		},
		{
			MethodName: "TransferTokenOwner",
			Handler:    _Msg_TransferTokenOwner_Handler,
		},
		{
			MethodName: "SwapFeeToken",
			Handler:    _Msg_SwapFeeToken_Handler,
		},
		{
			MethodName: "SwapToERC20",
			Handler:    _Msg_SwapToERC20_Handler,
		},
		{
			MethodName: "SwapFromERC20",
			Handler:    _Msg_SwapFromERC20_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "DeployERC20",
			Handler:    _Msg_DeployERC20_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "irismod/token/v1/tx.proto",
}
