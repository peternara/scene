// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serverpb/rpc.proto

/*
Package serverpb is a generated protocol buffer package.

It is generated from these files:
	serverpb/rpc.proto

It has these top-level messages:
	AuthRequest
	AuthResponse
	GeoPosition
	SignSpot
	PhotoPredictRequest
	PhotoPredictResponse
*/
package serverpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignSpot_SignSpotType int32

const (
	SignSpot_Museum SignSpot_SignSpotType = 0
)

var SignSpot_SignSpotType_name = map[int32]string{
	0: "Museum",
}
var SignSpot_SignSpotType_value = map[string]int32{
	"Museum": 0,
}

func (x SignSpot_SignSpotType) String() string {
	return proto.EnumName(SignSpot_SignSpotType_name, int32(x))
}
func (SignSpot_SignSpotType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type PhotoPredictRequest_PhotoType int32

const (
	PhotoPredictRequest_PNG PhotoPredictRequest_PhotoType = 0
	PhotoPredictRequest_JPG PhotoPredictRequest_PhotoType = 1
)

var PhotoPredictRequest_PhotoType_name = map[int32]string{
	0: "PNG",
	1: "JPG",
}
var PhotoPredictRequest_PhotoType_value = map[string]int32{
	"PNG": 0,
	"JPG": 1,
}

func (x PhotoPredictRequest_PhotoType) String() string {
	return proto.EnumName(PhotoPredictRequest_PhotoType_name, int32(x))
}
func (PhotoPredictRequest_PhotoType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

type AuthRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Token    string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *AuthRequest) Reset()                    { *m = AuthRequest{} }
func (m *AuthRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()               {}
func (*AuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthResponse struct {
	RequireLogin bool   `protobuf:"varint,1,opt,name=require_login,json=requireLogin" json:"require_login,omitempty"`
	Token        string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Msg          string `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
}

func (m *AuthResponse) Reset()                    { *m = AuthResponse{} }
func (m *AuthResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()               {}
func (*AuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthResponse) GetRequireLogin() bool {
	if m != nil {
		return m.RequireLogin
	}
	return false
}

func (m *AuthResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AuthResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GeoPosition struct {
	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude" json:"longitude,omitempty"`
}

func (m *GeoPosition) Reset()                    { *m = GeoPosition{} }
func (m *GeoPosition) String() string            { return proto.CompactTextString(m) }
func (*GeoPosition) ProtoMessage()               {}
func (*GeoPosition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GeoPosition) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *GeoPosition) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

type SignSpot struct {
	Id   uint64                `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string                `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Type SignSpot_SignSpotType `protobuf:"varint,3,opt,name=type,enum=serverpb.SignSpot_SignSpotType" json:"type,omitempty"`
	Geo  *GeoPosition          `protobuf:"bytes,4,opt,name=geo" json:"geo,omitempty"`
}

func (m *SignSpot) Reset()                    { *m = SignSpot{} }
func (m *SignSpot) String() string            { return proto.CompactTextString(m) }
func (*SignSpot) ProtoMessage()               {}
func (*SignSpot) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SignSpot) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SignSpot) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SignSpot) GetType() SignSpot_SignSpotType {
	if m != nil {
		return m.Type
	}
	return SignSpot_Museum
}

func (m *SignSpot) GetGeo() *GeoPosition {
	if m != nil {
		return m.Geo
	}
	return nil
}

type PhotoPredictRequest struct {
	Type          PhotoPredictRequest_PhotoType `protobuf:"varint,1,opt,name=type,enum=serverpb.PhotoPredictRequest_PhotoType" json:"type,omitempty"`
	Data          []byte                        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Base64Encoded bool                          `protobuf:"varint,3,opt,name=base64_encoded,json=base64Encoded" json:"base64_encoded,omitempty"`
	Geo           *GeoPosition                  `protobuf:"bytes,4,opt,name=geo" json:"geo,omitempty"`
	AcquireText   bool                          `protobuf:"varint,5,opt,name=acquire_text,json=acquireText" json:"acquire_text,omitempty"`
	AcquireAudio  bool                          `protobuf:"varint,6,opt,name=acquire_audio,json=acquireAudio" json:"acquire_audio,omitempty"`
	AcquireVideo  bool                          `protobuf:"varint,7,opt,name=acquire_video,json=acquireVideo" json:"acquire_video,omitempty"`
}

func (m *PhotoPredictRequest) Reset()                    { *m = PhotoPredictRequest{} }
func (m *PhotoPredictRequest) String() string            { return proto.CompactTextString(m) }
func (*PhotoPredictRequest) ProtoMessage()               {}
func (*PhotoPredictRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PhotoPredictRequest) GetType() PhotoPredictRequest_PhotoType {
	if m != nil {
		return m.Type
	}
	return PhotoPredictRequest_PNG
}

func (m *PhotoPredictRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *PhotoPredictRequest) GetBase64Encoded() bool {
	if m != nil {
		return m.Base64Encoded
	}
	return false
}

func (m *PhotoPredictRequest) GetGeo() *GeoPosition {
	if m != nil {
		return m.Geo
	}
	return nil
}

func (m *PhotoPredictRequest) GetAcquireText() bool {
	if m != nil {
		return m.AcquireText
	}
	return false
}

func (m *PhotoPredictRequest) GetAcquireAudio() bool {
	if m != nil {
		return m.AcquireAudio
	}
	return false
}

func (m *PhotoPredictRequest) GetAcquireVideo() bool {
	if m != nil {
		return m.AcquireVideo
	}
	return false
}

type PhotoPredictResponse struct {
	Results []*PhotoPredictResponse_Result `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *PhotoPredictResponse) Reset()                    { *m = PhotoPredictResponse{} }
func (m *PhotoPredictResponse) String() string            { return proto.CompactTextString(m) }
func (*PhotoPredictResponse) ProtoMessage()               {}
func (*PhotoPredictResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PhotoPredictResponse) GetResults() []*PhotoPredictResponse_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

type PhotoPredictResponse_Result struct {
	Text     string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl" json:"image_url,omitempty"`
	AudioUrl string `protobuf:"bytes,3,opt,name=audio_url,json=audioUrl" json:"audio_url,omitempty"`
	VideoUrl string `protobuf:"bytes,4,opt,name=video_url,json=videoUrl" json:"video_url,omitempty"`
}

func (m *PhotoPredictResponse_Result) Reset()                    { *m = PhotoPredictResponse_Result{} }
func (m *PhotoPredictResponse_Result) String() string            { return proto.CompactTextString(m) }
func (*PhotoPredictResponse_Result) ProtoMessage()               {}
func (*PhotoPredictResponse_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *PhotoPredictResponse_Result) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PhotoPredictResponse_Result) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *PhotoPredictResponse_Result) GetAudioUrl() string {
	if m != nil {
		return m.AudioUrl
	}
	return ""
}

func (m *PhotoPredictResponse_Result) GetVideoUrl() string {
	if m != nil {
		return m.VideoUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "serverpb.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "serverpb.AuthResponse")
	proto.RegisterType((*GeoPosition)(nil), "serverpb.GeoPosition")
	proto.RegisterType((*SignSpot)(nil), "serverpb.SignSpot")
	proto.RegisterType((*PhotoPredictRequest)(nil), "serverpb.PhotoPredictRequest")
	proto.RegisterType((*PhotoPredictResponse)(nil), "serverpb.PhotoPredictResponse")
	proto.RegisterType((*PhotoPredictResponse_Result)(nil), "serverpb.PhotoPredictResponse.Result")
	proto.RegisterEnum("serverpb.SignSpot_SignSpotType", SignSpot_SignSpotType_name, SignSpot_SignSpotType_value)
	proto.RegisterEnum("serverpb.PhotoPredictRequest_PhotoType", PhotoPredictRequest_PhotoType_name, PhotoPredictRequest_PhotoType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/serverpb.Auth/Authenticate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Authenticate(context.Context, *AuthRequest) (*AuthResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.Auth/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _Auth_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serverpb/rpc.proto",
}

// Client API for Predict service

type PredictClient interface {
	PredictPhoto(ctx context.Context, in *PhotoPredictRequest, opts ...grpc.CallOption) (*PhotoPredictResponse, error)
}

type predictClient struct {
	cc *grpc.ClientConn
}

func NewPredictClient(cc *grpc.ClientConn) PredictClient {
	return &predictClient{cc}
}

func (c *predictClient) PredictPhoto(ctx context.Context, in *PhotoPredictRequest, opts ...grpc.CallOption) (*PhotoPredictResponse, error) {
	out := new(PhotoPredictResponse)
	err := grpc.Invoke(ctx, "/serverpb.Predict/PredictPhoto", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Predict service

type PredictServer interface {
	PredictPhoto(context.Context, *PhotoPredictRequest) (*PhotoPredictResponse, error)
}

func RegisterPredictServer(s *grpc.Server, srv PredictServer) {
	s.RegisterService(&_Predict_serviceDesc, srv)
}

func _Predict_PredictPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhotoPredictRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictServer).PredictPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.Predict/PredictPhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictServer).PredictPhoto(ctx, req.(*PhotoPredictRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Predict_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.Predict",
	HandlerType: (*PredictServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PredictPhoto",
			Handler:    _Predict_PredictPhoto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "serverpb/rpc.proto",
}

func init() { proto.RegisterFile("serverpb/rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x1d, 0x37, 0x1f, 0x13, 0x37, 0x8a, 0x96, 0x82, 0xac, 0x40, 0xa1, 0x18, 0x95, 0xf6,
	0x64, 0xa4, 0x14, 0x71, 0xe1, 0x50, 0xb5, 0x12, 0x8a, 0x84, 0xf8, 0xb0, 0xdc, 0x96, 0x03, 0x1c,
	0x2a, 0xc7, 0x1e, 0xb9, 0x2b, 0x1c, 0xaf, 0xbb, 0xbb, 0x2e, 0xed, 0xdf, 0xe1, 0xc4, 0xef, 0xe0,
	0xcc, 0x8f, 0x42, 0xbb, 0x6b, 0x27, 0x2e, 0x82, 0x8a, 0xdb, 0xec, 0x7b, 0xb3, 0xcf, 0xf3, 0x66,
	0x67, 0x0c, 0x44, 0x20, 0xbf, 0x42, 0x5e, 0x2e, 0x5e, 0xf0, 0x32, 0x09, 0x4a, 0xce, 0x24, 0x23,
	0x83, 0x06, 0xf3, 0x4f, 0x60, 0x74, 0x54, 0xc9, 0x8b, 0x08, 0x2f, 0x2b, 0x14, 0x92, 0x10, 0x70,
	0x8a, 0x78, 0x89, 0x9e, 0xb5, 0x63, 0xed, 0x0f, 0x23, 0x1d, 0x93, 0x29, 0x0c, 0xca, 0x58, 0x88,
	0x6f, 0x8c, 0xa7, 0x9e, 0xad, 0xf1, 0xd5, 0x99, 0x6c, 0xc1, 0x86, 0x64, 0x5f, 0xb1, 0xf0, 0xba,
	0x9a, 0x30, 0x07, 0xff, 0x0b, 0xb8, 0x46, 0x54, 0x94, 0xac, 0x10, 0x48, 0x9e, 0xc1, 0x26, 0xc7,
	0xcb, 0x8a, 0x72, 0x3c, 0xcf, 0x59, 0x46, 0x0b, 0x2d, 0x3f, 0x88, 0xdc, 0x1a, 0x7c, 0xa7, 0xb0,
	0xb5, 0x94, 0xdd, 0x92, 0x22, 0x13, 0xe8, 0x2e, 0x45, 0x56, 0xcb, 0xab, 0xd0, 0x9f, 0xc3, 0x68,
	0x8e, 0x2c, 0x64, 0x82, 0x4a, 0xca, 0x0a, 0x55, 0x5d, 0x1e, 0x4b, 0x2a, 0xab, 0xd4, 0x54, 0x6d,
	0x45, 0xab, 0x33, 0x79, 0x04, 0xc3, 0x9c, 0x15, 0x99, 0x21, 0x6d, 0x4d, 0xae, 0x01, 0xff, 0x87,
	0x05, 0x83, 0x13, 0x9a, 0x15, 0x27, 0x25, 0x93, 0x64, 0x0c, 0x36, 0x4d, 0xb5, 0x80, 0x13, 0xd9,
	0x34, 0x5d, 0x35, 0xc2, 0x6e, 0x35, 0xe2, 0x00, 0x1c, 0x79, 0x53, 0xa2, 0x2e, 0x66, 0x3c, 0x7b,
	0x12, 0x34, 0x4d, 0x0c, 0x1a, 0x95, 0x55, 0x70, 0x7a, 0x53, 0x62, 0xa4, 0x93, 0xc9, 0x1e, 0x74,
	0x33, 0x64, 0x9e, 0xb3, 0x63, 0xed, 0x8f, 0x66, 0xf7, 0xd7, 0x77, 0x5a, 0x1e, 0x22, 0x95, 0xe1,
	0x4f, 0xc1, 0x6d, 0x5f, 0x27, 0x00, 0xbd, 0xf7, 0x95, 0xc0, 0x6a, 0x39, 0xe9, 0xf8, 0x3f, 0x6d,
	0xb8, 0x17, 0x5e, 0x30, 0xc9, 0x42, 0x8e, 0x29, 0x4d, 0x64, 0xf3, 0x5c, 0xaf, 0xeb, 0x8a, 0x2c,
	0x5d, 0xd1, 0xde, 0x5a, 0xfd, 0x2f, 0xc9, 0x06, 0x6b, 0x55, 0x46, 0xc0, 0x49, 0x63, 0x19, 0x6b,
	0x8b, 0x6e, 0xa4, 0x63, 0xb2, 0x0b, 0xe3, 0x45, 0x2c, 0xf0, 0xd5, 0xcb, 0x73, 0x2c, 0x12, 0x96,
	0x62, 0xaa, 0xcd, 0x0e, 0xa2, 0x4d, 0x83, 0xbe, 0x31, 0xe0, 0x7f, 0x9b, 0x22, 0x4f, 0xc1, 0x8d,
	0x13, 0xf3, 0xf2, 0x12, 0xaf, 0xa5, 0xb7, 0xa1, 0xd5, 0x46, 0x35, 0x76, 0x8a, 0xd7, 0x52, 0x0d,
	0x47, 0x93, 0x12, 0x57, 0x29, 0x65, 0x5e, 0xcf, 0x0c, 0x47, 0x0d, 0x1e, 0x29, 0xac, 0x9d, 0x74,
	0x45, 0x53, 0x64, 0x5e, 0xff, 0x56, 0xd2, 0x27, 0x85, 0xf9, 0xdb, 0x30, 0x5c, 0x79, 0x24, 0x7d,
	0xe8, 0x86, 0x1f, 0xe6, 0x93, 0x8e, 0x0a, 0xde, 0x86, 0xf3, 0x89, 0xe5, 0xff, 0xb2, 0x60, 0xeb,
	0x76, 0x5f, 0xea, 0xf1, 0x3c, 0x84, 0x3e, 0x47, 0x51, 0xe5, 0x52, 0x78, 0xd6, 0x4e, 0x77, 0x7f,
	0x34, 0xdb, 0xfd, 0x57, 0x23, 0xcd, 0x85, 0x20, 0xd2, 0xd9, 0x51, 0x73, 0x6b, 0x2a, 0xa0, 0x67,
	0x20, 0xd5, 0x53, 0xed, 0xb3, 0xde, 0x1f, 0x15, 0x93, 0x87, 0x30, 0xa4, 0xcb, 0x38, 0xc3, 0xf3,
	0x8a, 0xe7, 0xcd, 0x02, 0x69, 0xe0, 0x8c, 0xe7, 0x8a, 0xd4, 0xae, 0x35, 0x69, 0xa6, 0x7c, 0xa0,
	0x81, 0x9a, 0xd4, 0x6e, 0x35, 0xe9, 0x18, 0x52, 0x03, 0x67, 0x3c, 0x9f, 0xcd, 0xc1, 0x51, 0x4b,
	0x46, 0x0e, 0xcd, 0xb2, 0x61, 0x21, 0x69, 0x12, 0x4b, 0x24, 0xad, 0xe7, 0x68, 0x6d, 0xf6, 0xf4,
	0xc1, 0x9f, 0xb0, 0xf1, 0xe2, 0x77, 0x66, 0x9f, 0xa1, 0x5f, 0x1b, 0x24, 0x1f, 0xc1, 0xad, 0x43,
	0xed, 0x9b, 0x6c, 0xdf, 0x39, 0x51, 0xd3, 0xc7, 0x77, 0xf7, 0xc9, 0xef, 0x1c, 0x3f, 0x87, 0x71,
	0xc2, 0x96, 0x41, 0x4c, 0x25, 0xab, 0x78, 0xc0, 0xcb, 0xe4, 0x78, 0xa8, 0xbe, 0x1e, 0xaa, 0xbf,
	0x50, 0x68, 0x7d, 0xb7, 0x7b, 0x86, 0x59, 0xf4, 0xf4, 0x7f, 0xe9, 0xe0, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x17, 0xd9, 0xaf, 0x86, 0xad, 0x04, 0x00, 0x00,
}
