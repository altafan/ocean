// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: ocean/v1alpha/account.proto

package oceanv1alpha

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	// CreateAccount creates a new account.
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	// SetAccountTemplate sets the template for the account used to generate new addresses.
	SetAccountTemplate(ctx context.Context, in *SetAccountTemplateRequest, opts ...grpc.CallOption) (*SetAccountTemplateResponse, error)
	// DeriveAddress generates new address(es) for the account.
	DeriveAddress(ctx context.Context, in *DeriveAddressRequest, opts ...grpc.CallOption) (*DeriveAddressResponse, error)
	// DeriveChangeAddress generates new change address(es) for the account.
	DeriveChangeAddress(ctx context.Context, in *DeriveChangeAddressRequest, opts ...grpc.CallOption) (*DeriveChangeAddressResponse, error)
	// ListAddresses returns all derived addresses for the account.
	ListAddresses(ctx context.Context, in *ListAddressesRequest, opts ...grpc.CallOption) (*ListAddressesResponse, error)
	// Balance returns the balance for the account, or for specific list of
	// account's addresses.
	Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error)
	// ListUtxos returns the utxos for the account, or specific list of
	// account's addresses.
	ListUtxos(ctx context.Context, in *ListUtxosRequest, opts ...grpc.CallOption) (*ListUtxosResponse, error)
	// DeleteAccount deletes an existing account. The operation is allowed only
	// if the account has zero balance.
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1alpha.AccountService/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) SetAccountTemplate(ctx context.Context, in *SetAccountTemplateRequest, opts ...grpc.CallOption) (*SetAccountTemplateResponse, error) {
	out := new(SetAccountTemplateResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1alpha.AccountService/SetAccountTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeriveAddress(ctx context.Context, in *DeriveAddressRequest, opts ...grpc.CallOption) (*DeriveAddressResponse, error) {
	out := new(DeriveAddressResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1alpha.AccountService/DeriveAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeriveChangeAddress(ctx context.Context, in *DeriveChangeAddressRequest, opts ...grpc.CallOption) (*DeriveChangeAddressResponse, error) {
	out := new(DeriveChangeAddressResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1alpha.AccountService/DeriveChangeAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ListAddresses(ctx context.Context, in *ListAddressesRequest, opts ...grpc.CallOption) (*ListAddressesResponse, error) {
	out := new(ListAddressesResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1alpha.AccountService/ListAddresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Balance(ctx context.Context, in *BalanceRequest, opts ...grpc.CallOption) (*BalanceResponse, error) {
	out := new(BalanceResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1alpha.AccountService/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ListUtxos(ctx context.Context, in *ListUtxosRequest, opts ...grpc.CallOption) (*ListUtxosResponse, error) {
	out := new(ListUtxosResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1alpha.AccountService/ListUtxos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, "/ocean.v1alpha.AccountService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations should embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	// CreateAccount creates a new account.
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	// SetAccountTemplate sets the template for the account used to generate new addresses.
	SetAccountTemplate(context.Context, *SetAccountTemplateRequest) (*SetAccountTemplateResponse, error)
	// DeriveAddress generates new address(es) for the account.
	DeriveAddress(context.Context, *DeriveAddressRequest) (*DeriveAddressResponse, error)
	// DeriveChangeAddress generates new change address(es) for the account.
	DeriveChangeAddress(context.Context, *DeriveChangeAddressRequest) (*DeriveChangeAddressResponse, error)
	// ListAddresses returns all derived addresses for the account.
	ListAddresses(context.Context, *ListAddressesRequest) (*ListAddressesResponse, error)
	// Balance returns the balance for the account, or for specific list of
	// account's addresses.
	Balance(context.Context, *BalanceRequest) (*BalanceResponse, error)
	// ListUtxos returns the utxos for the account, or specific list of
	// account's addresses.
	ListUtxos(context.Context, *ListUtxosRequest) (*ListUtxosResponse, error)
	// DeleteAccount deletes an existing account. The operation is allowed only
	// if the account has zero balance.
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
}

// UnimplementedAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedAccountServiceServer) SetAccountTemplate(context.Context, *SetAccountTemplateRequest) (*SetAccountTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccountTemplate not implemented")
}
func (UnimplementedAccountServiceServer) DeriveAddress(context.Context, *DeriveAddressRequest) (*DeriveAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeriveAddress not implemented")
}
func (UnimplementedAccountServiceServer) DeriveChangeAddress(context.Context, *DeriveChangeAddressRequest) (*DeriveChangeAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeriveChangeAddress not implemented")
}
func (UnimplementedAccountServiceServer) ListAddresses(context.Context, *ListAddressesRequest) (*ListAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAddresses not implemented")
}
func (UnimplementedAccountServiceServer) Balance(context.Context, *BalanceRequest) (*BalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedAccountServiceServer) ListUtxos(context.Context, *ListUtxosRequest) (*ListUtxosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUtxos not implemented")
}
func (UnimplementedAccountServiceServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1alpha.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_SetAccountTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAccountTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SetAccountTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1alpha.AccountService/SetAccountTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SetAccountTemplate(ctx, req.(*SetAccountTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeriveAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeriveAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeriveAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1alpha.AccountService/DeriveAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeriveAddress(ctx, req.(*DeriveAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeriveChangeAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeriveChangeAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeriveChangeAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1alpha.AccountService/DeriveChangeAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeriveChangeAddress(ctx, req.(*DeriveChangeAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ListAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ListAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1alpha.AccountService/ListAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ListAddresses(ctx, req.(*ListAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1alpha.AccountService/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Balance(ctx, req.(*BalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ListUtxos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUtxosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ListUtxos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1alpha.AccountService/ListUtxos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ListUtxos(ctx, req.(*ListUtxosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocean.v1alpha.AccountService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocean.v1alpha.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "SetAccountTemplate",
			Handler:    _AccountService_SetAccountTemplate_Handler,
		},
		{
			MethodName: "DeriveAddress",
			Handler:    _AccountService_DeriveAddress_Handler,
		},
		{
			MethodName: "DeriveChangeAddress",
			Handler:    _AccountService_DeriveChangeAddress_Handler,
		},
		{
			MethodName: "ListAddresses",
			Handler:    _AccountService_ListAddresses_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _AccountService_Balance_Handler,
		},
		{
			MethodName: "ListUtxos",
			Handler:    _AccountService_ListUtxos_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountService_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocean/v1alpha/account.proto",
}
