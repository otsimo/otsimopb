// Code generated by protoc-gen-gogo.
// source: api.proto
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

// Client API for ApiService service

type ApiServiceClient interface {
	// Profile
	AddProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Response, error)
	GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*Profile, error)
	UpdateProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Response, error)
	// Child
	AddChild(ctx context.Context, in *Child, opts ...grpc.CallOption) (*Response, error)
	GetChild(ctx context.Context, in *GetChildRequest, opts ...grpc.CallOption) (*Child, error)
	UpdateChild(ctx context.Context, in *Child, opts ...grpc.CallOption) (*Response, error)
	GetChildren(ctx context.Context, in *GetChildrenFromProfileRequest, opts ...grpc.CallOption) (*GetChildrenFromProfileResponse, error)
	UpdateGameEntry(ctx context.Context, in *GameEntryRequest, opts ...grpc.CallOption) (*Response, error)
	ChangeActivation(ctx context.Context, in *ChangeChildActivationRequest, opts ...grpc.CallOption) (*Response, error)
	GetDisabledChildren(ctx context.Context, in *GetChildrenFromProfileRequest, opts ...grpc.CallOption) (*GetChildrenFromProfileResponse, error)
	SoundEnable(ctx context.Context, in *SoundEnableRequest, opts ...grpc.CallOption) (*Response, error)
	UpdateGameIndices(ctx context.Context, in *UpdateIndecesRequest, opts ...grpc.CallOption) (*Child, error)
}

type apiServiceClient struct {
	cc *grpc.ClientConn
}

func NewApiServiceClient(cc *grpc.ClientConn) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) AddProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.ApiService/AddProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...grpc.CallOption) (*Profile, error) {
	out := new(Profile)
	err := grpc.Invoke(ctx, "/apipb.ApiService/GetProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) UpdateProfile(ctx context.Context, in *Profile, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.ApiService/UpdateProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) AddChild(ctx context.Context, in *Child, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.ApiService/AddChild", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetChild(ctx context.Context, in *GetChildRequest, opts ...grpc.CallOption) (*Child, error) {
	out := new(Child)
	err := grpc.Invoke(ctx, "/apipb.ApiService/GetChild", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) UpdateChild(ctx context.Context, in *Child, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.ApiService/UpdateChild", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetChildren(ctx context.Context, in *GetChildrenFromProfileRequest, opts ...grpc.CallOption) (*GetChildrenFromProfileResponse, error) {
	out := new(GetChildrenFromProfileResponse)
	err := grpc.Invoke(ctx, "/apipb.ApiService/GetChildren", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) UpdateGameEntry(ctx context.Context, in *GameEntryRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.ApiService/UpdateGameEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) ChangeActivation(ctx context.Context, in *ChangeChildActivationRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.ApiService/ChangeActivation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetDisabledChildren(ctx context.Context, in *GetChildrenFromProfileRequest, opts ...grpc.CallOption) (*GetChildrenFromProfileResponse, error) {
	out := new(GetChildrenFromProfileResponse)
	err := grpc.Invoke(ctx, "/apipb.ApiService/GetDisabledChildren", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) SoundEnable(ctx context.Context, in *SoundEnableRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/apipb.ApiService/SoundEnable", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) UpdateGameIndices(ctx context.Context, in *UpdateIndecesRequest, opts ...grpc.CallOption) (*Child, error) {
	out := new(Child)
	err := grpc.Invoke(ctx, "/apipb.ApiService/UpdateGameIndices", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApiService service

type ApiServiceServer interface {
	// Profile
	AddProfile(context.Context, *Profile) (*Response, error)
	GetProfile(context.Context, *GetProfileRequest) (*Profile, error)
	UpdateProfile(context.Context, *Profile) (*Response, error)
	// Child
	AddChild(context.Context, *Child) (*Response, error)
	GetChild(context.Context, *GetChildRequest) (*Child, error)
	UpdateChild(context.Context, *Child) (*Response, error)
	GetChildren(context.Context, *GetChildrenFromProfileRequest) (*GetChildrenFromProfileResponse, error)
	UpdateGameEntry(context.Context, *GameEntryRequest) (*Response, error)
	ChangeActivation(context.Context, *ChangeChildActivationRequest) (*Response, error)
	GetDisabledChildren(context.Context, *GetChildrenFromProfileRequest) (*GetChildrenFromProfileResponse, error)
	SoundEnable(context.Context, *SoundEnableRequest) (*Response, error)
	UpdateGameIndices(context.Context, *UpdateIndecesRequest) (*Child, error)
}

func RegisterApiServiceServer(s *grpc.Server, srv ApiServiceServer) {
	s.RegisterService(&_ApiService_serviceDesc, srv)
}

func _ApiService_AddProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).AddProfile(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_GetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).GetProfile(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Profile)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).UpdateProfile(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_AddChild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Child)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).AddChild(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_GetChild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetChildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).GetChild(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_UpdateChild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Child)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).UpdateChild(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_GetChildren_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetChildrenFromProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).GetChildren(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_UpdateGameEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GameEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).UpdateGameEntry(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_ChangeActivation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChangeChildActivationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).ChangeActivation(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_GetDisabledChildren_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetChildrenFromProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).GetDisabledChildren(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_SoundEnable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SoundEnableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).SoundEnable(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _ApiService_UpdateGameIndices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(UpdateIndecesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServiceServer).UpdateGameIndices(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProfile",
			Handler:    _ApiService_AddProfile_Handler,
		},
		{
			MethodName: "GetProfile",
			Handler:    _ApiService_GetProfile_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _ApiService_UpdateProfile_Handler,
		},
		{
			MethodName: "AddChild",
			Handler:    _ApiService_AddChild_Handler,
		},
		{
			MethodName: "GetChild",
			Handler:    _ApiService_GetChild_Handler,
		},
		{
			MethodName: "UpdateChild",
			Handler:    _ApiService_UpdateChild_Handler,
		},
		{
			MethodName: "GetChildren",
			Handler:    _ApiService_GetChildren_Handler,
		},
		{
			MethodName: "UpdateGameEntry",
			Handler:    _ApiService_UpdateGameEntry_Handler,
		},
		{
			MethodName: "ChangeActivation",
			Handler:    _ApiService_ChangeActivation_Handler,
		},
		{
			MethodName: "GetDisabledChildren",
			Handler:    _ApiService_GetDisabledChildren_Handler,
		},
		{
			MethodName: "SoundEnable",
			Handler:    _ApiService_SoundEnable_Handler,
		},
		{
			MethodName: "UpdateGameIndices",
			Handler:    _ApiService_UpdateGameIndices_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptorApi = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x53, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x6d, 0x41, 0x54, 0xe5, 0xb6, 0xb4, 0x60, 0x24, 0x1e, 0x41, 0xea, 0x82, 0xc7, 0x0a, 0x91,
	0x4a, 0x65, 0x03, 0x42, 0x20, 0x95, 0x52, 0x0a, 0x2b, 0xaa, 0x16, 0x36, 0x88, 0x4d, 0x12, 0xdf,
	0xa6, 0x96, 0x12, 0xdb, 0xc4, 0x4e, 0x25, 0xb6, 0x7c, 0xc1, 0x7c, 0xc3, 0x7c, 0x4d, 0x97, 0xfd,
	0x84, 0x79, 0xfc, 0xc8, 0xb8, 0x4e, 0xd2, 0xcc, 0x54, 0xad, 0x66, 0x66, 0x31, 0x0b, 0x4b, 0x3e,
	0xe7, 0xde, 0x7b, 0xce, 0xb1, 0x13, 0xc3, 0x5d, 0x4f, 0x32, 0x57, 0x26, 0x42, 0x0b, 0x72, 0xc7,
	0x6c, 0xa5, 0xef, 0xb4, 0x62, 0x54, 0xca, 0x0b, 0x51, 0x65, 0xb4, 0xd3, 0x8c, 0x05, 0xc5, 0xa8,
	0x40, 0x6f, 0x42, 0xa6, 0xe7, 0xa9, 0xef, 0x06, 0x22, 0xee, 0x86, 0x22, 0x14, 0x5d, 0x4b, 0xfb,
	0xe9, 0xcc, 0x22, 0x0b, 0xec, 0x2e, 0x6b, 0xef, 0xfd, 0xaf, 0x01, 0xf4, 0x25, 0x9b, 0x62, 0xb2,
	0x60, 0x01, 0x92, 0xae, 0x41, 0x94, 0x8e, 0x13, 0x31, 0x63, 0x11, 0x92, 0x96, 0x6b, 0x1d, 0xdd,
	0x1c, 0x3b, 0xed, 0x1c, 0x4f, 0x50, 0x49, 0xc1, 0x15, 0x3e, 0xaf, 0x90, 0x77, 0x00, 0x23, 0xd4,
	0xc5, 0xc0, 0x93, 0xbc, 0xa1, 0xa4, 0x26, 0xf8, 0x37, 0x45, 0xa5, 0x9d, 0x2d, 0x29, 0x33, 0xd9,
	0x83, 0x7b, 0xbf, 0x24, 0xf5, 0x34, 0x5e, 0xc3, 0xed, 0x35, 0xd4, 0x4d, 0xbc, 0xc1, 0x9c, 0x45,
	0x94, 0x34, 0xf3, 0xb2, 0x45, 0xbb, 0x9a, 0x7b, 0x50, 0x37, 0x39, 0xb2, 0xe6, 0x47, 0x65, 0x30,
	0x4b, 0x14, 0xb1, 0x2e, 0x88, 0x98, 0x19, 0x17, 0x1a, 0x59, 0xa8, 0x2b, 0x7a, 0xfc, 0x81, 0x46,
	0x21, 0x99, 0x20, 0x27, 0x2f, 0xb7, 0x6c, 0x0c, 0xf7, 0x35, 0x11, 0xf1, 0xd6, 0x5d, 0xbc, 0xba,
	0xa4, 0x6b, 0xa3, 0xfe, 0x11, 0xda, 0x59, 0x9a, 0x91, 0x17, 0xe3, 0x90, 0xeb, 0xe4, 0x1f, 0x79,
	0x5c, 0xcc, 0x16, 0x4c, 0x21, 0xba, 0x23, 0xdc, 0x37, 0xb8, 0x3f, 0x98, 0x7b, 0x3c, 0xc4, 0x7e,
	0xa0, 0xd9, 0xc2, 0xd3, 0x4c, 0x70, 0xf2, 0x62, 0x73, 0xa2, 0x75, 0xc1, 0xda, 0x97, 0xd5, 0x7d,
	0x5a, 0xc4, 0x87, 0x87, 0x26, 0xec, 0x17, 0xa6, 0x3c, 0x3f, 0x42, 0x7a, 0x23, 0xc7, 0x25, 0xef,
	0xa1, 0x31, 0x15, 0x29, 0xa7, 0x43, 0xbe, 0x36, 0x21, 0x4f, 0xf3, 0xa9, 0x73, 0xdc, 0xde, 0x78,
	0x9f, 0xe0, 0x41, 0x79, 0x4f, 0xdf, 0x39, 0x35, 0x7f, 0xb2, 0x22, 0xcf, 0xf2, 0xae, 0xac, 0x62,
	0x58, 0x34, 0xec, 0xce, 0xef, 0xfe, 0xf9, 0xc3, 0xf2, 0xb8, 0x53, 0x59, 0x99, 0xb5, 0x3c, 0xe9,
	0x54, 0x57, 0x66, 0x1d, 0x99, 0x75, 0x70, 0xda, 0xa9, 0x40, 0xdb, 0x3c, 0x23, 0x57, 0x68, 0xc5,
	0x62, 0xe1, 0x86, 0x89, 0x0c, 0xc6, 0xd5, 0xdf, 0xf5, 0x0c, 0x4a, 0xff, 0xf0, 0xd6, 0xed, 0x1f,
	0x3f, 0xa7, 0x7e, 0xcd, 0x3e, 0xa4, 0xb7, 0x67, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x05, 0x6b,
	0xda, 0xa9, 0x03, 0x00, 0x00,
}
