// Code generated by protoc-gen-gogo.
// source: shim.proto
// DO NOT EDIT!

/*
	Package shim is a generated protocol buffer package.

	It is generated from these files:
		shim.proto

	It has these top-level messages:
		CreateRequest
		CreateResponse
		StartRequest
		DeleteRequest
		DeleteResponse
		ExecRequest
		ExecResponse
		PtyRequest
*/
package shim

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptorShim, []int{0} }

type CreateResponse struct {
	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptorShim, []int{1} }

type StartRequest struct {
}

func (m *StartRequest) Reset()                    { *m = StartRequest{} }
func (*StartRequest) ProtoMessage()               {}
func (*StartRequest) Descriptor() ([]byte, []int) { return fileDescriptorShim, []int{2} }

type DeleteRequest struct {
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptorShim, []int{3} }

type DeleteResponse struct {
	ExitStatus uint32 `protobuf:"varint,1,opt,name=exit_status,json=exitStatus,proto3" json:"exit_status,omitempty"`
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptorShim, []int{4} }

type ExecRequest struct {
}

func (m *ExecRequest) Reset()                    { *m = ExecRequest{} }
func (*ExecRequest) ProtoMessage()               {}
func (*ExecRequest) Descriptor() ([]byte, []int) { return fileDescriptorShim, []int{5} }

type ExecResponse struct {
	Pid uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (m *ExecResponse) Reset()                    { *m = ExecResponse{} }
func (*ExecResponse) ProtoMessage()               {}
func (*ExecResponse) Descriptor() ([]byte, []int) { return fileDescriptorShim, []int{6} }

type PtyRequest struct {
	ID     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Width  uint32 `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *PtyRequest) Reset()                    { *m = PtyRequest{} }
func (*PtyRequest) ProtoMessage()               {}
func (*PtyRequest) Descriptor() ([]byte, []int) { return fileDescriptorShim, []int{7} }

func init() {
	proto.RegisterType((*CreateRequest)(nil), "containerd.v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "containerd.v1.CreateResponse")
	proto.RegisterType((*StartRequest)(nil), "containerd.v1.StartRequest")
	proto.RegisterType((*DeleteRequest)(nil), "containerd.v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "containerd.v1.DeleteResponse")
	proto.RegisterType((*ExecRequest)(nil), "containerd.v1.ExecRequest")
	proto.RegisterType((*ExecResponse)(nil), "containerd.v1.ExecResponse")
	proto.RegisterType((*PtyRequest)(nil), "containerd.v1.PtyRequest")
}
func (this *CreateRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&shim.CreateRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CreateResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&shim.CreateResponse{")
	s = append(s, "Pid: "+fmt.Sprintf("%#v", this.Pid)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StartRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&shim.StartRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DeleteRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&shim.DeleteRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DeleteResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&shim.DeleteResponse{")
	s = append(s, "ExitStatus: "+fmt.Sprintf("%#v", this.ExitStatus)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ExecRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&shim.ExecRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ExecResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&shim.ExecResponse{")
	s = append(s, "Pid: "+fmt.Sprintf("%#v", this.Pid)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PtyRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&shim.PtyRequest{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "Width: "+fmt.Sprintf("%#v", this.Width)+",\n")
	s = append(s, "Height: "+fmt.Sprintf("%#v", this.Height)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringShim(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringShim(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ShimService service

type ShimServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
	Pty(ctx context.Context, in *PtyRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
}

type shimServiceClient struct {
	cc *grpc.ClientConn
}

func NewShimServiceClient(cc *grpc.ClientConn) ShimServiceClient {
	return &shimServiceClient{cc}
}

func (c *shimServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := grpc.Invoke(ctx, "/containerd.v1.ShimService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shimServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/containerd.v1.ShimService/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shimServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := grpc.Invoke(ctx, "/containerd.v1.ShimService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shimServiceClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := grpc.Invoke(ctx, "/containerd.v1.ShimService/Exec", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shimServiceClient) Pty(ctx context.Context, in *PtyRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/containerd.v1.ShimService/Pty", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShimService service

type ShimServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Start(context.Context, *StartRequest) (*google_protobuf.Empty, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
	Pty(context.Context, *PtyRequest) (*google_protobuf.Empty, error)
}

func RegisterShimServiceServer(s *grpc.Server, srv ShimServiceServer) {
	s.RegisterService(&_ShimService_serviceDesc, srv)
}

func _ShimService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShimServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.ShimService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShimServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShimService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShimServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.ShimService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShimServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShimService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShimServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.ShimService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShimServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShimService_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShimServiceServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.ShimService/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShimServiceServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShimService_Pty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PtyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShimServiceServer).Pty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.v1.ShimService/Pty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShimServiceServer).Pty(ctx, req.(*PtyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShimService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.v1.ShimService",
	HandlerType: (*ShimServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ShimService_Create_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _ShimService_Start_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ShimService_Delete_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _ShimService_Exec_Handler,
		},
		{
			MethodName: "Pty",
			Handler:    _ShimService_Pty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shim.proto",
}

func (m *CreateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *CreateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Pid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintShim(dAtA, i, uint64(m.Pid))
	}
	return i, nil
}

func (m *StartRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StartRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *DeleteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *DeleteResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeleteResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ExitStatus != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintShim(dAtA, i, uint64(m.ExitStatus))
	}
	return i, nil
}

func (m *ExecRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ExecResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExecResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Pid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintShim(dAtA, i, uint64(m.Pid))
	}
	return i, nil
}

func (m *PtyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PtyRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintShim(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Width != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintShim(dAtA, i, uint64(m.Width))
	}
	if m.Height != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintShim(dAtA, i, uint64(m.Height))
	}
	return i, nil
}

func encodeFixed64Shim(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Shim(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintShim(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CreateRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *CreateResponse) Size() (n int) {
	var l int
	_ = l
	if m.Pid != 0 {
		n += 1 + sovShim(uint64(m.Pid))
	}
	return n
}

func (m *StartRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *DeleteRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *DeleteResponse) Size() (n int) {
	var l int
	_ = l
	if m.ExitStatus != 0 {
		n += 1 + sovShim(uint64(m.ExitStatus))
	}
	return n
}

func (m *ExecRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ExecResponse) Size() (n int) {
	var l int
	_ = l
	if m.Pid != 0 {
		n += 1 + sovShim(uint64(m.Pid))
	}
	return n
}

func (m *PtyRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovShim(uint64(l))
	}
	if m.Width != 0 {
		n += 1 + sovShim(uint64(m.Width))
	}
	if m.Height != 0 {
		n += 1 + sovShim(uint64(m.Height))
	}
	return n
}

func sovShim(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozShim(x uint64) (n int) {
	return sovShim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CreateRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateRequest{`,
		`}`,
	}, "")
	return s
}
func (this *CreateResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CreateResponse{`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StartRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StartRequest{`,
		`}`,
	}, "")
	return s
}
func (this *DeleteRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DeleteRequest{`,
		`}`,
	}, "")
	return s
}
func (this *DeleteResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DeleteResponse{`,
		`ExitStatus:` + fmt.Sprintf("%v", this.ExitStatus) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ExecRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ExecRequest{`,
		`}`,
	}, "")
	return s
}
func (this *ExecResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ExecResponse{`,
		`Pid:` + fmt.Sprintf("%v", this.Pid) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PtyRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PtyRequest{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Width:` + fmt.Sprintf("%v", this.Width) + `,`,
		`Height:` + fmt.Sprintf("%v", this.Height) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringShim(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CreateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipShim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StartRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StartRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StartRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipShim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeleteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipShim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DeleteResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeleteResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeleteResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExitStatus", wireType)
			}
			m.ExitStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExitStatus |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExecRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipShim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExecResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ExecResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExecResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pid", wireType)
			}
			m.Pid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pid |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PtyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShim
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PtyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PtyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthShim
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthShim
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipShim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShim
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowShim
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthShim
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowShim
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipShim(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthShim = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShim   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("shim.proto", fileDescriptorShim) }

var fileDescriptorShim = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x51, 0xbd, 0x6e, 0xf2, 0x40,
	0x10, 0xc4, 0xe6, 0xc3, 0xd2, 0xb7, 0x60, 0xf8, 0x74, 0x42, 0x88, 0xcf, 0x10, 0x83, 0x5c, 0xa5,
	0x32, 0x22, 0x69, 0x52, 0x44, 0x8a, 0x44, 0xa0, 0x48, 0x87, 0xec, 0x07, 0x88, 0x0c, 0xde, 0xd8,
	0x27, 0x01, 0xe7, 0xd8, 0x07, 0x81, 0x2e, 0x8f, 0x47, 0x99, 0x22, 0x45, 0xaa, 0x28, 0xf8, 0x09,
	0xf2, 0x08, 0x91, 0xff, 0x04, 0x98, 0xd0, 0xed, 0xce, 0xce, 0x8d, 0xe6, 0x66, 0x00, 0x02, 0x97,
	0xce, 0x75, 0xcf, 0x67, 0x9c, 0x11, 0x79, 0xca, 0x16, 0xdc, 0xa2, 0x0b, 0xf4, 0x6d, 0x7d, 0xd5,
	0x57, 0x5a, 0x0e, 0x63, 0xce, 0x0c, 0x7b, 0xf1, 0x71, 0xb2, 0x7c, 0xea, 0xe1, 0xdc, 0xe3, 0x9b,
	0x84, 0xab, 0xd4, 0x1d, 0xe6, 0xb0, 0x78, 0xec, 0x45, 0x53, 0x82, 0x6a, 0x35, 0x90, 0xef, 0x7d,
	0xb4, 0x38, 0x1a, 0xf8, 0xbc, 0xc4, 0x80, 0x6b, 0x1a, 0x54, 0x33, 0x20, 0xf0, 0xd8, 0x22, 0x40,
	0xf2, 0x0f, 0x8a, 0x1e, 0xb5, 0x9b, 0x42, 0x57, 0xb8, 0x94, 0x8d, 0x68, 0xd4, 0xaa, 0x50, 0x31,
	0xb9, 0xe5, 0xf3, 0xec, 0x4d, 0x0d, 0xe4, 0x21, 0xce, 0x70, 0x2f, 0xd2, 0x87, 0x6a, 0x06, 0xa4,
	0x22, 0x1d, 0x28, 0xe3, 0x9a, 0xf2, 0xc7, 0x80, 0x5b, 0x7c, 0x19, 0xa4, 0x62, 0x10, 0x41, 0x66,
	0x8c, 0x68, 0x32, 0x94, 0x47, 0x6b, 0x9c, 0x66, 0x0a, 0x5d, 0xa8, 0x24, 0xeb, 0x59, 0x13, 0x06,
	0xc0, 0x98, 0x6f, 0x52, 0x3e, 0x69, 0x80, 0x98, 0x9e, 0xff, 0x0e, 0xa4, 0xf0, 0xb3, 0x23, 0x3e,
	0x0c, 0x0d, 0x91, 0xda, 0xa4, 0x0e, 0xa5, 0x17, 0x6a, 0x73, 0xb7, 0x29, 0xc6, 0x2f, 0x93, 0x85,
	0x34, 0x40, 0x72, 0x91, 0x3a, 0x2e, 0x6f, 0x16, 0x63, 0x38, 0xdd, 0xae, 0xde, 0x45, 0x28, 0x9b,
	0x2e, 0x9d, 0x9b, 0xe8, 0xaf, 0xe8, 0x14, 0xc9, 0x08, 0xa4, 0x24, 0x0c, 0xd2, 0xd6, 0x8f, 0xa2,
	0xd6, 0x8f, 0x42, 0x53, 0x2e, 0xce, 0x5c, 0x53, 0xf3, 0xb7, 0x50, 0x8a, 0xf3, 0x22, 0xad, 0x1c,
	0xef, 0x30, 0x45, 0xa5, 0xa1, 0x27, 0xf5, 0xe9, 0x59, 0x7d, 0xfa, 0x28, 0xaa, 0x2f, 0x32, 0x91,
	0x84, 0x79, 0x62, 0xe2, 0x28, 0xf4, 0x13, 0x13, 0xb9, 0x06, 0xee, 0xe0, 0x4f, 0x94, 0x28, 0x51,
	0x72, 0xb4, 0x83, 0xd4, 0x95, 0xd6, 0xaf, 0xb7, 0x54, 0xe0, 0x06, 0x8a, 0x63, 0xbe, 0x21, 0xff,
	0x73, 0x9c, 0x7d, 0x09, 0xe7, 0x7e, 0x30, 0x68, 0x6f, 0x77, 0x6a, 0xe1, 0x63, 0xa7, 0x16, 0xbe,
	0x77, 0xaa, 0xf0, 0x1a, 0xaa, 0xc2, 0x36, 0x54, 0x85, 0xb7, 0x50, 0x15, 0xbe, 0x42, 0x55, 0x98,
	0x48, 0x31, 0xfb, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x9a, 0x8d, 0xf1, 0xd9, 0x02, 0x00,
	0x00,
}