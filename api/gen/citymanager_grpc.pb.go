// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: citymanager.proto

package proto

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

// CityManagerClient is the client API for CityManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CityManagerClient interface {
	AddCity(ctx context.Context, in *AddCityRequest, opts ...grpc.CallOption) (*AddCityResponse, error)
	RemoveCity(ctx context.Context, in *RemoveCityRequest, opts ...grpc.CallOption) (*RemoveCityResponse, error)
	GetCity(ctx context.Context, in *GetCityRequest, opts ...grpc.CallOption) (*GetCityResponse, error)
	GetNearestCities(ctx context.Context, in *GetNearestCitiesRequest, opts ...grpc.CallOption) (*GetNearestCitiesResponse, error)
}

type cityManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewCityManagerClient(cc grpc.ClientConnInterface) CityManagerClient {
	return &cityManagerClient{cc}
}

func (c *cityManagerClient) AddCity(ctx context.Context, in *AddCityRequest, opts ...grpc.CallOption) (*AddCityResponse, error) {
	out := new(AddCityResponse)
	err := c.cc.Invoke(ctx, "/citymanager.CityManager/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityManagerClient) RemoveCity(ctx context.Context, in *RemoveCityRequest, opts ...grpc.CallOption) (*RemoveCityResponse, error) {
	out := new(RemoveCityResponse)
	err := c.cc.Invoke(ctx, "/citymanager.CityManager/RemoveCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityManagerClient) GetCity(ctx context.Context, in *GetCityRequest, opts ...grpc.CallOption) (*GetCityResponse, error) {
	out := new(GetCityResponse)
	err := c.cc.Invoke(ctx, "/citymanager.CityManager/GetCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cityManagerClient) GetNearestCities(ctx context.Context, in *GetNearestCitiesRequest, opts ...grpc.CallOption) (*GetNearestCitiesResponse, error) {
	out := new(GetNearestCitiesResponse)
	err := c.cc.Invoke(ctx, "/citymanager.CityManager/GetNearestCities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CityManagerServer is the server API for CityManager service.
// All implementations must embed UnimplementedCityManagerServer
// for forward compatibility
type CityManagerServer interface {
	AddCity(context.Context, *AddCityRequest) (*AddCityResponse, error)
	RemoveCity(context.Context, *RemoveCityRequest) (*RemoveCityResponse, error)
	GetCity(context.Context, *GetCityRequest) (*GetCityResponse, error)
	GetNearestCities(context.Context, *GetNearestCitiesRequest) (*GetNearestCitiesResponse, error)
	mustEmbedUnimplementedCityManagerServer()
}

// UnimplementedCityManagerServer must be embedded to have forward compatible implementations.
type UnimplementedCityManagerServer struct {
}

func (UnimplementedCityManagerServer) AddCity(context.Context, *AddCityRequest) (*AddCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedCityManagerServer) RemoveCity(context.Context, *RemoveCityRequest) (*RemoveCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCity not implemented")
}
func (UnimplementedCityManagerServer) GetCity(context.Context, *GetCityRequest) (*GetCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCity not implemented")
}
func (UnimplementedCityManagerServer) GetNearestCities(context.Context, *GetNearestCitiesRequest) (*GetNearestCitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNearestCities not implemented")
}
func (UnimplementedCityManagerServer) mustEmbedUnimplementedCityManagerServer() {}

// UnsafeCityManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CityManagerServer will
// result in compilation errors.
type UnsafeCityManagerServer interface {
	mustEmbedUnimplementedCityManagerServer()
}

func RegisterCityManagerServer(s grpc.ServiceRegistrar, srv CityManagerServer) {
	s.RegisterService(&CityManager_ServiceDesc, srv)
}

func _CityManager_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityManagerServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citymanager.CityManager/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityManagerServer).AddCity(ctx, req.(*AddCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityManager_RemoveCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityManagerServer).RemoveCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citymanager.CityManager/RemoveCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityManagerServer).RemoveCity(ctx, req.(*RemoveCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityManager_GetCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityManagerServer).GetCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citymanager.CityManager/GetCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityManagerServer).GetCity(ctx, req.(*GetCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CityManager_GetNearestCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNearestCitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CityManagerServer).GetNearestCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/citymanager.CityManager/GetNearestCities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CityManagerServer).GetNearestCities(ctx, req.(*GetNearestCitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CityManager_ServiceDesc is the grpc.ServiceDesc for CityManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CityManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "citymanager.CityManager",
	HandlerType: (*CityManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCity",
			Handler:    _CityManager_AddCity_Handler,
		},
		{
			MethodName: "RemoveCity",
			Handler:    _CityManager_RemoveCity_Handler,
		},
		{
			MethodName: "GetCity",
			Handler:    _CityManager_GetCity_Handler,
		},
		{
			MethodName: "GetNearestCities",
			Handler:    _CityManager_GetNearestCities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "citymanager.proto",
}
