// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.2
// source: pkg/testing/remote/loader.proto

package remote

import (
	context "context"
	server "github.com/linuxsuren/api-testing/pkg/server"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LoaderClient is the client API for Loader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoaderClient interface {
	ListTestSuite(ctx context.Context, in *server.Empty, opts ...grpc.CallOption) (*TestSuites, error)
	CreateTestSuite(ctx context.Context, in *TestSuite, opts ...grpc.CallOption) (*server.Empty, error)
	GetTestSuite(ctx context.Context, in *TestSuite, opts ...grpc.CallOption) (*TestSuite, error)
	UpdateTestSuite(ctx context.Context, in *TestSuite, opts ...grpc.CallOption) (*TestSuite, error)
	DeleteTestSuite(ctx context.Context, in *TestSuite, opts ...grpc.CallOption) (*server.Empty, error)
	ListTestCases(ctx context.Context, in *TestSuite, opts ...grpc.CallOption) (*server.TestCases, error)
	CreateTestCase(ctx context.Context, in *server.TestCase, opts ...grpc.CallOption) (*server.Empty, error)
	GetTestCase(ctx context.Context, in *server.TestCase, opts ...grpc.CallOption) (*server.TestCase, error)
	UpdateTestCase(ctx context.Context, in *server.TestCase, opts ...grpc.CallOption) (*server.TestCase, error)
	DeleteTestCase(ctx context.Context, in *server.TestCase, opts ...grpc.CallOption) (*server.Empty, error)
	Verify(ctx context.Context, in *server.Empty, opts ...grpc.CallOption) (*server.CommonResult, error)
}

type loaderClient struct {
	cc grpc.ClientConnInterface
}

func NewLoaderClient(cc grpc.ClientConnInterface) LoaderClient {
	return &loaderClient{cc}
}

func (c *loaderClient) ListTestSuite(ctx context.Context, in *server.Empty, opts ...grpc.CallOption) (*TestSuites, error) {
	out := new(TestSuites)
	err := c.cc.Invoke(ctx, "/remote.Loader/ListTestSuite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) CreateTestSuite(ctx context.Context, in *TestSuite, opts ...grpc.CallOption) (*server.Empty, error) {
	out := new(server.Empty)
	err := c.cc.Invoke(ctx, "/remote.Loader/CreateTestSuite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) GetTestSuite(ctx context.Context, in *TestSuite, opts ...grpc.CallOption) (*TestSuite, error) {
	out := new(TestSuite)
	err := c.cc.Invoke(ctx, "/remote.Loader/GetTestSuite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) UpdateTestSuite(ctx context.Context, in *TestSuite, opts ...grpc.CallOption) (*TestSuite, error) {
	out := new(TestSuite)
	err := c.cc.Invoke(ctx, "/remote.Loader/UpdateTestSuite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) DeleteTestSuite(ctx context.Context, in *TestSuite, opts ...grpc.CallOption) (*server.Empty, error) {
	out := new(server.Empty)
	err := c.cc.Invoke(ctx, "/remote.Loader/DeleteTestSuite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) ListTestCases(ctx context.Context, in *TestSuite, opts ...grpc.CallOption) (*server.TestCases, error) {
	out := new(server.TestCases)
	err := c.cc.Invoke(ctx, "/remote.Loader/ListTestCases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) CreateTestCase(ctx context.Context, in *server.TestCase, opts ...grpc.CallOption) (*server.Empty, error) {
	out := new(server.Empty)
	err := c.cc.Invoke(ctx, "/remote.Loader/CreateTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) GetTestCase(ctx context.Context, in *server.TestCase, opts ...grpc.CallOption) (*server.TestCase, error) {
	out := new(server.TestCase)
	err := c.cc.Invoke(ctx, "/remote.Loader/GetTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) UpdateTestCase(ctx context.Context, in *server.TestCase, opts ...grpc.CallOption) (*server.TestCase, error) {
	out := new(server.TestCase)
	err := c.cc.Invoke(ctx, "/remote.Loader/UpdateTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) DeleteTestCase(ctx context.Context, in *server.TestCase, opts ...grpc.CallOption) (*server.Empty, error) {
	out := new(server.Empty)
	err := c.cc.Invoke(ctx, "/remote.Loader/DeleteTestCase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loaderClient) Verify(ctx context.Context, in *server.Empty, opts ...grpc.CallOption) (*server.CommonResult, error) {
	out := new(server.CommonResult)
	err := c.cc.Invoke(ctx, "/remote.Loader/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoaderServer is the server API for Loader service.
// All implementations must embed UnimplementedLoaderServer
// for forward compatibility
type LoaderServer interface {
	ListTestSuite(context.Context, *server.Empty) (*TestSuites, error)
	CreateTestSuite(context.Context, *TestSuite) (*server.Empty, error)
	GetTestSuite(context.Context, *TestSuite) (*TestSuite, error)
	UpdateTestSuite(context.Context, *TestSuite) (*TestSuite, error)
	DeleteTestSuite(context.Context, *TestSuite) (*server.Empty, error)
	ListTestCases(context.Context, *TestSuite) (*server.TestCases, error)
	CreateTestCase(context.Context, *server.TestCase) (*server.Empty, error)
	GetTestCase(context.Context, *server.TestCase) (*server.TestCase, error)
	UpdateTestCase(context.Context, *server.TestCase) (*server.TestCase, error)
	DeleteTestCase(context.Context, *server.TestCase) (*server.Empty, error)
	Verify(context.Context, *server.Empty) (*server.CommonResult, error)
	mustEmbedUnimplementedLoaderServer()
}

// UnimplementedLoaderServer must be embedded to have forward compatible implementations.
type UnimplementedLoaderServer struct {
}

func (UnimplementedLoaderServer) ListTestSuite(context.Context, *server.Empty) (*TestSuites, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTestSuite not implemented")
}
func (UnimplementedLoaderServer) CreateTestSuite(context.Context, *TestSuite) (*server.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTestSuite not implemented")
}
func (UnimplementedLoaderServer) GetTestSuite(context.Context, *TestSuite) (*TestSuite, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestSuite not implemented")
}
func (UnimplementedLoaderServer) UpdateTestSuite(context.Context, *TestSuite) (*TestSuite, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTestSuite not implemented")
}
func (UnimplementedLoaderServer) DeleteTestSuite(context.Context, *TestSuite) (*server.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTestSuite not implemented")
}
func (UnimplementedLoaderServer) ListTestCases(context.Context, *TestSuite) (*server.TestCases, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTestCases not implemented")
}
func (UnimplementedLoaderServer) CreateTestCase(context.Context, *server.TestCase) (*server.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTestCase not implemented")
}
func (UnimplementedLoaderServer) GetTestCase(context.Context, *server.TestCase) (*server.TestCase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTestCase not implemented")
}
func (UnimplementedLoaderServer) UpdateTestCase(context.Context, *server.TestCase) (*server.TestCase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTestCase not implemented")
}
func (UnimplementedLoaderServer) DeleteTestCase(context.Context, *server.TestCase) (*server.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTestCase not implemented")
}
func (UnimplementedLoaderServer) Verify(context.Context, *server.Empty) (*server.CommonResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedLoaderServer) mustEmbedUnimplementedLoaderServer() {}

// UnsafeLoaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoaderServer will
// result in compilation errors.
type UnsafeLoaderServer interface {
	mustEmbedUnimplementedLoaderServer()
}

func RegisterLoaderServer(s grpc.ServiceRegistrar, srv LoaderServer) {
	s.RegisterService(&Loader_ServiceDesc, srv)
}

func _Loader_ListTestSuite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).ListTestSuite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/ListTestSuite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).ListTestSuite(ctx, req.(*server.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_CreateTestSuite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestSuite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).CreateTestSuite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/CreateTestSuite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).CreateTestSuite(ctx, req.(*TestSuite))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_GetTestSuite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestSuite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).GetTestSuite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/GetTestSuite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).GetTestSuite(ctx, req.(*TestSuite))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_UpdateTestSuite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestSuite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).UpdateTestSuite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/UpdateTestSuite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).UpdateTestSuite(ctx, req.(*TestSuite))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_DeleteTestSuite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestSuite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).DeleteTestSuite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/DeleteTestSuite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).DeleteTestSuite(ctx, req.(*TestSuite))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_ListTestCases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestSuite)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).ListTestCases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/ListTestCases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).ListTestCases(ctx, req.(*TestSuite))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_CreateTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.TestCase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).CreateTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/CreateTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).CreateTestCase(ctx, req.(*server.TestCase))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_GetTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.TestCase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).GetTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/GetTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).GetTestCase(ctx, req.(*server.TestCase))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_UpdateTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.TestCase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).UpdateTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/UpdateTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).UpdateTestCase(ctx, req.(*server.TestCase))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_DeleteTestCase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.TestCase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).DeleteTestCase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/DeleteTestCase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).DeleteTestCase(ctx, req.(*server.TestCase))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loader_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoaderServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.Loader/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoaderServer).Verify(ctx, req.(*server.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Loader_ServiceDesc is the grpc.ServiceDesc for Loader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Loader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remote.Loader",
	HandlerType: (*LoaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTestSuite",
			Handler:    _Loader_ListTestSuite_Handler,
		},
		{
			MethodName: "CreateTestSuite",
			Handler:    _Loader_CreateTestSuite_Handler,
		},
		{
			MethodName: "GetTestSuite",
			Handler:    _Loader_GetTestSuite_Handler,
		},
		{
			MethodName: "UpdateTestSuite",
			Handler:    _Loader_UpdateTestSuite_Handler,
		},
		{
			MethodName: "DeleteTestSuite",
			Handler:    _Loader_DeleteTestSuite_Handler,
		},
		{
			MethodName: "ListTestCases",
			Handler:    _Loader_ListTestCases_Handler,
		},
		{
			MethodName: "CreateTestCase",
			Handler:    _Loader_CreateTestCase_Handler,
		},
		{
			MethodName: "GetTestCase",
			Handler:    _Loader_GetTestCase_Handler,
		},
		{
			MethodName: "UpdateTestCase",
			Handler:    _Loader_UpdateTestCase_Handler,
		},
		{
			MethodName: "DeleteTestCase",
			Handler:    _Loader_DeleteTestCase_Handler,
		},
		{
			MethodName: "Verify",
			Handler:    _Loader_Verify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/testing/remote/loader.proto",
}

// SecretServiceClient is the client API for SecretService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecretServiceClient interface {
	GetSecret(ctx context.Context, in *server.Secret, opts ...grpc.CallOption) (*server.Secret, error)
	GetSecrets(ctx context.Context, in *server.Empty, opts ...grpc.CallOption) (*server.Secrets, error)
	CreateSecret(ctx context.Context, in *server.Secret, opts ...grpc.CallOption) (*server.CommonResult, error)
	DeleteSecret(ctx context.Context, in *server.Secret, opts ...grpc.CallOption) (*server.CommonResult, error)
	UpdateSecret(ctx context.Context, in *server.Secret, opts ...grpc.CallOption) (*server.CommonResult, error)
}

type secretServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecretServiceClient(cc grpc.ClientConnInterface) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) GetSecret(ctx context.Context, in *server.Secret, opts ...grpc.CallOption) (*server.Secret, error) {
	out := new(server.Secret)
	err := c.cc.Invoke(ctx, "/remote.SecretService/GetSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) GetSecrets(ctx context.Context, in *server.Empty, opts ...grpc.CallOption) (*server.Secrets, error) {
	out := new(server.Secrets)
	err := c.cc.Invoke(ctx, "/remote.SecretService/GetSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) CreateSecret(ctx context.Context, in *server.Secret, opts ...grpc.CallOption) (*server.CommonResult, error) {
	out := new(server.CommonResult)
	err := c.cc.Invoke(ctx, "/remote.SecretService/CreateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) DeleteSecret(ctx context.Context, in *server.Secret, opts ...grpc.CallOption) (*server.CommonResult, error) {
	out := new(server.CommonResult)
	err := c.cc.Invoke(ctx, "/remote.SecretService/DeleteSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) UpdateSecret(ctx context.Context, in *server.Secret, opts ...grpc.CallOption) (*server.CommonResult, error) {
	out := new(server.CommonResult)
	err := c.cc.Invoke(ctx, "/remote.SecretService/UpdateSecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretServiceServer is the server API for SecretService service.
// All implementations must embed UnimplementedSecretServiceServer
// for forward compatibility
type SecretServiceServer interface {
	GetSecret(context.Context, *server.Secret) (*server.Secret, error)
	GetSecrets(context.Context, *server.Empty) (*server.Secrets, error)
	CreateSecret(context.Context, *server.Secret) (*server.CommonResult, error)
	DeleteSecret(context.Context, *server.Secret) (*server.CommonResult, error)
	UpdateSecret(context.Context, *server.Secret) (*server.CommonResult, error)
	mustEmbedUnimplementedSecretServiceServer()
}

// UnimplementedSecretServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecretServiceServer struct {
}

func (UnimplementedSecretServiceServer) GetSecret(context.Context, *server.Secret) (*server.Secret, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecret not implemented")
}
func (UnimplementedSecretServiceServer) GetSecrets(context.Context, *server.Empty) (*server.Secrets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecrets not implemented")
}
func (UnimplementedSecretServiceServer) CreateSecret(context.Context, *server.Secret) (*server.CommonResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSecret not implemented")
}
func (UnimplementedSecretServiceServer) DeleteSecret(context.Context, *server.Secret) (*server.CommonResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecret not implemented")
}
func (UnimplementedSecretServiceServer) UpdateSecret(context.Context, *server.Secret) (*server.CommonResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSecret not implemented")
}
func (UnimplementedSecretServiceServer) mustEmbedUnimplementedSecretServiceServer() {}

// UnsafeSecretServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecretServiceServer will
// result in compilation errors.
type UnsafeSecretServiceServer interface {
	mustEmbedUnimplementedSecretServiceServer()
}

func RegisterSecretServiceServer(s grpc.ServiceRegistrar, srv SecretServiceServer) {
	s.RegisterService(&SecretService_ServiceDesc, srv)
}

func _SecretService_GetSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).GetSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.SecretService/GetSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).GetSecret(ctx, req.(*server.Secret))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_GetSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).GetSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.SecretService/GetSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).GetSecrets(ctx, req.(*server.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_CreateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).CreateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.SecretService/CreateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).CreateSecret(ctx, req.(*server.Secret))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_DeleteSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).DeleteSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.SecretService/DeleteSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).DeleteSecret(ctx, req.(*server.Secret))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_UpdateSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(server.Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).UpdateSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.SecretService/UpdateSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).UpdateSecret(ctx, req.(*server.Secret))
	}
	return interceptor(ctx, in, info, handler)
}

// SecretService_ServiceDesc is the grpc.ServiceDesc for SecretService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecretService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remote.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSecret",
			Handler:    _SecretService_GetSecret_Handler,
		},
		{
			MethodName: "GetSecrets",
			Handler:    _SecretService_GetSecrets_Handler,
		},
		{
			MethodName: "CreateSecret",
			Handler:    _SecretService_CreateSecret_Handler,
		},
		{
			MethodName: "DeleteSecret",
			Handler:    _SecretService_DeleteSecret_Handler,
		},
		{
			MethodName: "UpdateSecret",
			Handler:    _SecretService_UpdateSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/testing/remote/loader.proto",
}
