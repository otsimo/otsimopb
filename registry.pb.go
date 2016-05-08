// Code generated by protoc-gen-gogo.
// source: registry.proto
// DO NOT EDIT!

package otsimopb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for RegistryService service

type RegistryServiceClient interface {
	// Get returns game
	Get(ctx context.Context, in *GetGameRequest, opts ...grpc.CallOption) (*Game, error)
	// GetRelease returns GameRelease by given game id and version
	GetRelease(ctx context.Context, in *GetGameReleaseRequest, opts ...grpc.CallOption) (*GameRelease, error)
	// Publish tries to create a new GameRelease by given manifest
	Publish(ctx context.Context, in *GameManifest, opts ...grpc.CallOption) (*PublishResponse, error)
	// ChangeReleaseState changes state of a release, If user is admin than s/he can change
	// from WAITING to REJECTED or VALIDATED, developers can change to any except VALIDATED
	ChangeReleaseState(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*Response, error)
	// GetLatestVersions returns latest versions of given game ids
	GetLatestVersions(ctx context.Context, in *GetLatestVersionsRequest, opts ...grpc.CallOption) (*GameVersionsResponse, error)
	// Search does search
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// ListGames returns all games
	ListGames(ctx context.Context, in *ListGamesRequest, opts ...grpc.CallOption) (RegistryService_ListGamesClient, error)
}

type registryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRegistryServiceClient(cc *grpc.ClientConn) RegistryServiceClient {
	return &registryServiceClient{cc}
}

func (c *registryServiceClient) Get(ctx context.Context, in *GetGameRequest, opts ...grpc.CallOption) (*Game, error) {
	out := new(Game)
	err := grpc.Invoke(ctx, "/apipb.RegistryService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) GetRelease(ctx context.Context, in *GetGameReleaseRequest, opts ...grpc.CallOption) (*GameRelease, error) {
	out := new(GameRelease)
	err := grpc.Invoke(ctx, "/apipb.RegistryService/GetRelease", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) Publish(ctx context.Context, in *GameManifest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := grpc.Invoke(ctx, "/apipb.RegistryService/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) ChangeReleaseState(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.RegistryService/ChangeReleaseState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) GetLatestVersions(ctx context.Context, in *GetLatestVersionsRequest, opts ...grpc.CallOption) (*GameVersionsResponse, error) {
	out := new(GameVersionsResponse)
	err := grpc.Invoke(ctx, "/apipb.RegistryService/GetLatestVersions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := grpc.Invoke(ctx, "/apipb.RegistryService/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryServiceClient) ListGames(ctx context.Context, in *ListGamesRequest, opts ...grpc.CallOption) (RegistryService_ListGamesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RegistryService_serviceDesc.Streams[0], c.cc, "/apipb.RegistryService/ListGames", opts...)
	if err != nil {
		return nil, err
	}
	x := &registryServiceListGamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RegistryService_ListGamesClient interface {
	Recv() (*ListItem, error)
	grpc.ClientStream
}

type registryServiceListGamesClient struct {
	grpc.ClientStream
}

func (x *registryServiceListGamesClient) Recv() (*ListItem, error) {
	m := new(ListItem)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RegistryService service

type RegistryServiceServer interface {
	// Get returns game
	Get(context.Context, *GetGameRequest) (*Game, error)
	// GetRelease returns GameRelease by given game id and version
	GetRelease(context.Context, *GetGameReleaseRequest) (*GameRelease, error)
	// Publish tries to create a new GameRelease by given manifest
	Publish(context.Context, *GameManifest) (*PublishResponse, error)
	// ChangeReleaseState changes state of a release, If user is admin than s/he can change
	// from WAITING to REJECTED or VALIDATED, developers can change to any except VALIDATED
	ChangeReleaseState(context.Context, *ValidateRequest) (*Response, error)
	// GetLatestVersions returns latest versions of given game ids
	GetLatestVersions(context.Context, *GetLatestVersionsRequest) (*GameVersionsResponse, error)
	// Search does search
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// ListGames returns all games
	ListGames(*ListGamesRequest, RegistryService_ListGamesServer) error
}

func RegisterRegistryServiceServer(s *grpc.Server, srv RegistryServiceServer) {
	s.RegisterService(&_RegistryService_serviceDesc, srv)
}

func _RegistryService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.RegistryService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).Get(ctx, req.(*GetGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_GetRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGameReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).GetRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.RegistryService/GetRelease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).GetRelease(ctx, req.(*GetGameReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GameManifest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.RegistryService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).Publish(ctx, req.(*GameManifest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_ChangeReleaseState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).ChangeReleaseState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.RegistryService/ChangeReleaseState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).ChangeReleaseState(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_GetLatestVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).GetLatestVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.RegistryService/GetLatestVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).GetLatestVersions(ctx, req.(*GetLatestVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistryServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.RegistryService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistryServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistryService_ListGames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListGamesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RegistryServiceServer).ListGames(m, &registryServiceListGamesServer{stream})
}

type RegistryService_ListGamesServer interface {
	Send(*ListItem) error
	grpc.ServerStream
}

type registryServiceListGamesServer struct {
	grpc.ServerStream
}

func (x *registryServiceListGamesServer) Send(m *ListItem) error {
	return x.ServerStream.SendMsg(m)
}

var _RegistryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.RegistryService",
	HandlerType: (*RegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RegistryService_Get_Handler,
		},
		{
			MethodName: "GetRelease",
			Handler:    _RegistryService_GetRelease_Handler,
		},
		{
			MethodName: "Publish",
			Handler:    _RegistryService_Publish_Handler,
		},
		{
			MethodName: "ChangeReleaseState",
			Handler:    _RegistryService_ChangeReleaseState_Handler,
		},
		{
			MethodName: "GetLatestVersions",
			Handler:    _RegistryService_GetLatestVersions_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _RegistryService_Search_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListGames",
			Handler:       _RegistryService_ListGames_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptorRegistry = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x37, 0x87, 0x53, 0xa3, 0x6c, 0x18, 0x9d, 0xc2, 0x94, 0x09, 0xde, 0x3c, 0xd8, 0x89,
	0x1e, 0xe6, 0x49, 0x41, 0x0f, 0x43, 0x98, 0x38, 0x5a, 0xd9, 0xc1, 0x5b, 0xda, 0xbd, 0x4b, 0x03,
	0xed, 0x52, 0x93, 0x54, 0xf0, 0x5b, 0xf8, 0x19, 0xfc, 0x20, 0x9e, 0x77, 0xf4, 0x23, 0xf8, 0xe7,
	0x8b, 0x98, 0xa5, 0xcd, 0x56, 0x7b, 0x08, 0xe4, 0xfd, 0xbd, 0xcf, 0xf3, 0xbc, 0x09, 0x09, 0x6a,
	0x08, 0xa0, 0x4c, 0x2a, 0xf1, 0xea, 0x24, 0x82, 0x2b, 0x8e, 0x57, 0x49, 0xc2, 0x12, 0xbf, 0xdd,
	0x88, 0x41, 0x4a, 0x42, 0x41, 0x66, 0xb8, 0xbd, 0x15, 0xf3, 0x31, 0x44, 0xb6, 0x3a, 0xa5, 0x4c,
	0x85, 0xa9, 0xef, 0x04, 0x3c, 0xee, 0x52, 0x4e, 0x79, 0xd7, 0x60, 0x3f, 0x9d, 0x98, 0xca, 0x14,
	0x66, 0x97, 0xc9, 0xcf, 0x3f, 0x6a, 0xa8, 0xe9, 0xe6, 0x63, 0x3c, 0x10, 0x2f, 0x2c, 0x00, 0x7c,
	0x82, 0x6a, 0x7d, 0x50, 0xb8, 0xe5, 0x98, 0x79, 0x8e, 0xde, 0xf7, 0x49, 0x0c, 0x2e, 0x3c, 0xa7,
	0x20, 0x55, 0x7b, 0xd3, 0x62, 0xcd, 0xf0, 0x15, 0x42, 0xba, 0xed, 0x42, 0x04, 0x44, 0x02, 0x3e,
	0x2c, 0x3b, 0x0c, 0xb6, 0x46, 0x5c, 0x30, 0x5a, 0xc7, 0x25, 0x5a, 0x1b, 0xa6, 0x7e, 0xc4, 0x64,
	0x88, 0x77, 0x0a, 0xed, 0x7b, 0x32, 0x65, 0x93, 0xb9, 0x67, 0x2f, 0x87, 0xb9, 0xc8, 0x05, 0x99,
	0xf0, 0xa9, 0x84, 0xe3, 0x0a, 0xbe, 0x46, 0xf8, 0x36, 0x24, 0x53, 0x6a, 0xa3, 0x3c, 0x45, 0x14,
	0x60, 0xab, 0x1f, 0x91, 0x88, 0x8d, 0x35, 0xb0, 0xb3, 0x9b, 0x39, 0x2f, 0x04, 0x78, 0x68, 0x5b,
	0x9f, 0x73, 0xa0, 0x45, 0x52, 0x8d, 0x40, 0x48, 0xa6, 0x39, 0x3e, 0x5a, 0xde, 0xe0, 0x7f, 0xc7,
	0x06, 0x1d, 0x14, 0x4e, 0xb9, 0xec, 0x2d, 0x42, 0x7b, 0xa8, 0xee, 0x01, 0x11, 0x41, 0x88, 0x77,
	0x73, 0x61, 0x56, 0x5a, 0x7b, 0xab, 0x44, 0x0b, 0xc6, 0x8d, 0x81, 0x7e, 0x84, 0x79, 0xac, 0xc4,
	0xfb, 0xb9, 0x6a, 0x41, 0xca, 0xd7, 0x98, 0x37, 0xee, 0x14, 0xc4, 0x67, 0xd5, 0x9b, 0xde, 0xec,
	0xbb, 0x53, 0x99, 0xfd, 0x74, 0xaa, 0x9f, 0x7a, 0x7d, 0xe9, 0xf5, 0xf6, 0xdb, 0xa9, 0xa0, 0xa6,
	0x7e, 0x7e, 0x87, 0x2b, 0xc9, 0x62, 0xee, 0x50, 0x91, 0x04, 0xc3, 0xea, 0xd3, 0x7a, 0x56, 0x26,
	0xfe, 0xfb, 0x4a, 0xed, 0xe1, 0xd1, 0xf3, 0xeb, 0xe6, 0x03, 0x5c, 0xfc, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x72, 0xc8, 0x6e, 0xee, 0x66, 0x02, 0x00, 0x00,
}
