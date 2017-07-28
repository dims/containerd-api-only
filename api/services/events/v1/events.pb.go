// Code generated by protoc-gen-gogo.
// source: github.com/containerd/containerd/api/services/events/v1/events.proto
// DO NOT EDIT!

package events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/gogo/protobuf/types"

import time "time"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

type PublishRequest struct {
	Topic string                `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Event *google_protobuf1.Any `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
}

func (m *PublishRequest) Reset()                    { *m = PublishRequest{} }
func (*PublishRequest) ProtoMessage()               {}
func (*PublishRequest) Descriptor() ([]byte, []int) { return fileDescriptorEvents, []int{0} }

type ForwardRequest struct {
	Envelope *Envelope `protobuf:"bytes,1,opt,name=envelope" json:"envelope,omitempty"`
}

func (m *ForwardRequest) Reset()                    { *m = ForwardRequest{} }
func (*ForwardRequest) ProtoMessage()               {}
func (*ForwardRequest) Descriptor() ([]byte, []int) { return fileDescriptorEvents, []int{1} }

type SubscribeRequest struct {
	Filters []string `protobuf:"bytes,1,rep,name=filters" json:"filters,omitempty"`
}

func (m *SubscribeRequest) Reset()                    { *m = SubscribeRequest{} }
func (*SubscribeRequest) ProtoMessage()               {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) { return fileDescriptorEvents, []int{2} }

type Envelope struct {
	Timestamp time.Time             `protobuf:"bytes,1,opt,name=timestamp,stdtime" json:"timestamp"`
	Namespace string                `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Topic     string                `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Event     *google_protobuf1.Any `protobuf:"bytes,4,opt,name=event" json:"event,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptorEvents, []int{3} }

func init() {
	proto.RegisterType((*PublishRequest)(nil), "containerd.services.events.v1.PublishRequest")
	proto.RegisterType((*ForwardRequest)(nil), "containerd.services.events.v1.ForwardRequest")
	proto.RegisterType((*SubscribeRequest)(nil), "containerd.services.events.v1.SubscribeRequest")
	proto.RegisterType((*Envelope)(nil), "containerd.services.events.v1.Envelope")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Events service

type EventsClient interface {
	// Publish an event to a topic.
	//
	// The event will be packed into a timestamp envelope with the namespace
	// introspected from the context. The envelope will then be dispatched.
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	// Forward sends an event that has already been packaged into an envelope
	// with a timestamp and namespace.
	//
	// This is useful if earlier timestamping is required or when fowarding on
	// behalf of another component, namespace or publisher.
	Forward(ctx context.Context, in *ForwardRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error)
	// Subscribe to a stream of events, possibly returning only that match any
	// of the provided filters.
	//
	// Unlike many other methods in containerd, subscribers will get messages
	// from all namespaces unless otherwise specified. If this is not desired,
	// a filter can be provided in the format 'namespace==<namespace>' to
	// restrict the received events.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Events_SubscribeClient, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/containerd.services.events.v1.Events/Publish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) Forward(ctx context.Context, in *ForwardRequest, opts ...grpc.CallOption) (*google_protobuf2.Empty, error) {
	out := new(google_protobuf2.Empty)
	err := grpc.Invoke(ctx, "/containerd.services.events.v1.Events/Forward", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Events_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Events_serviceDesc.Streams[0], c.cc, "/containerd.services.events.v1.Events/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Events_SubscribeClient interface {
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type eventsSubscribeClient struct {
	grpc.ClientStream
}

func (x *eventsSubscribeClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Events service

type EventsServer interface {
	// Publish an event to a topic.
	//
	// The event will be packed into a timestamp envelope with the namespace
	// introspected from the context. The envelope will then be dispatched.
	Publish(context.Context, *PublishRequest) (*google_protobuf2.Empty, error)
	// Forward sends an event that has already been packaged into an envelope
	// with a timestamp and namespace.
	//
	// This is useful if earlier timestamping is required or when fowarding on
	// behalf of another component, namespace or publisher.
	Forward(context.Context, *ForwardRequest) (*google_protobuf2.Empty, error)
	// Subscribe to a stream of events, possibly returning only that match any
	// of the provided filters.
	//
	// Unlike many other methods in containerd, subscribers will get messages
	// from all namespaces unless otherwise specified. If this is not desired,
	// a filter can be provided in the format 'namespace==<namespace>' to
	// restrict the received events.
	Subscribe(*SubscribeRequest, Events_SubscribeServer) error
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.events.v1.Events/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_Forward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).Forward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.services.events.v1.Events/Forward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).Forward(ctx, req.(*ForwardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Events_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsServer).Subscribe(m, &eventsSubscribeServer{stream})
}

type Events_SubscribeServer interface {
	Send(*Envelope) error
	grpc.ServerStream
}

type eventsSubscribeServer struct {
	grpc.ServerStream
}

func (x *eventsSubscribeServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.services.events.v1.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Events_Publish_Handler,
		},
		{
			MethodName: "Forward",
			Handler:    _Events_Forward_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Events_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/containerd/containerd/api/services/events/v1/events.proto",
}

func (m *PublishRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PublishRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Topic) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Topic)))
		i += copy(dAtA[i:], m.Topic)
	}
	if m.Event != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.Event.Size()))
		n1, err := m.Event.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *ForwardRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForwardRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Envelope != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.Envelope.Size()))
		n2, err := m.Envelope.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *SubscribeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubscribeRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *Envelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Envelope) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintEvents(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)))
	n3, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	if len(m.Namespace) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Namespace)))
		i += copy(dAtA[i:], m.Namespace)
	}
	if len(m.Topic) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Topic)))
		i += copy(dAtA[i:], m.Topic)
	}
	if m.Event != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.Event.Size()))
		n4, err := m.Event.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func encodeFixed64Events(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Events(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PublishRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *ForwardRequest) Size() (n int) {
	var l int
	_ = l
	if m.Envelope != nil {
		l = m.Envelope.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *SubscribeRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.Filters) > 0 {
		for _, s := range m.Filters {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *Envelope) Size() (n int) {
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovEvents(uint64(l))
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Topic)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PublishRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PublishRequest{`,
		`Topic:` + fmt.Sprintf("%v", this.Topic) + `,`,
		`Event:` + strings.Replace(fmt.Sprintf("%v", this.Event), "Any", "google_protobuf1.Any", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ForwardRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ForwardRequest{`,
		`Envelope:` + strings.Replace(fmt.Sprintf("%v", this.Envelope), "Envelope", "Envelope", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SubscribeRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SubscribeRequest{`,
		`Filters:` + fmt.Sprintf("%v", this.Filters) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Envelope) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Envelope{`,
		`Timestamp:` + strings.Replace(strings.Replace(this.Timestamp.String(), "Timestamp", "google_protobuf3.Timestamp", 1), `&`, ``, 1) + `,`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`Topic:` + fmt.Sprintf("%v", this.Topic) + `,`,
		`Event:` + strings.Replace(fmt.Sprintf("%v", this.Event), "Any", "google_protobuf1.Any", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringEvents(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PublishRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: PublishRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PublishRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &google_protobuf1.Any{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
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
func (m *ForwardRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: ForwardRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForwardRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Envelope", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Envelope == nil {
				m.Envelope = &Envelope{}
			}
			if err := m.Envelope.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
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
func (m *SubscribeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: SubscribeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubscribeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Filters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Filters = append(m.Filters, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
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
func (m *Envelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: Envelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Envelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Topic", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Topic = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &google_protobuf1.Any{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEvents
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
				next, err := skipEvents(dAtA[start:])
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
	ErrInvalidLengthEvents = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/services/events/v1/events.proto", fileDescriptorEvents)
}

var fileDescriptorEvents = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0x8d, 0x6d, 0x8d, 0x27, 0x4d, 0xc8, 0xaa, 0x50, 0x08, 0x90, 0x56, 0xb9, 0x50,
	0x21, 0xb0, 0x59, 0x39, 0x22, 0x21, 0x51, 0x28, 0xe7, 0xc9, 0x80, 0x84, 0xb8, 0x25, 0xd9, 0x5b,
	0x66, 0x29, 0x89, 0x43, 0xec, 0x04, 0xed, 0xc6, 0x47, 0xe0, 0xc2, 0xd7, 0xe0, 0x73, 0xf4, 0xc8,
	0x91, 0x13, 0xb0, 0x7c, 0x12, 0xd4, 0x24, 0x6e, 0x58, 0x0b, 0x04, 0x71, 0x7b, 0xce, 0xfb, 0xbf,
	0x5f, 0xfc, 0xfe, 0xff, 0x04, 0x3f, 0x8f, 0x84, 0x3e, 0x2f, 0x02, 0x1a, 0xca, 0x84, 0x85, 0x32,
	0xd5, 0xbe, 0x48, 0x21, 0x3f, 0xfd, 0xb5, 0xf4, 0x33, 0xc1, 0x14, 0xe4, 0xa5, 0x08, 0x41, 0x31,
	0x28, 0x21, 0xd5, 0x8a, 0x95, 0xc7, 0x6d, 0x45, 0xb3, 0x5c, 0x6a, 0x49, 0xee, 0x74, 0x7a, 0x6a,
	0xb4, 0xb4, 0x55, 0x94, 0xc7, 0xce, 0x28, 0x92, 0x91, 0xac, 0x95, 0x6c, 0x55, 0x35, 0x43, 0xce,
	0xcd, 0x48, 0xca, 0x28, 0x06, 0x56, 0x9f, 0x82, 0xe2, 0x8c, 0xf9, 0xe9, 0x45, 0xdb, 0xba, 0xb5,
	0xd9, 0x82, 0x24, 0xd3, 0xa6, 0x39, 0xde, 0x6c, 0x6a, 0x91, 0x80, 0xd2, 0x7e, 0x92, 0x35, 0x02,
	0x8f, 0xe3, 0xa3, 0x93, 0x22, 0x88, 0x85, 0x3a, 0xe7, 0xf0, 0xae, 0x00, 0xa5, 0xc9, 0x08, 0xef,
	0x69, 0x99, 0x89, 0xd0, 0x46, 0x13, 0x34, 0xb5, 0x78, 0x73, 0x20, 0xf7, 0xf0, 0x5e, 0x7d, 0x47,
	0x7b, 0x67, 0x82, 0xa6, 0x87, 0xb3, 0x11, 0x6d, 0xc0, 0xd4, 0x80, 0xe9, 0xd3, 0xf4, 0x82, 0x37,
	0x12, 0xef, 0x35, 0x3e, 0x7a, 0x21, 0xf3, 0xf7, 0x7e, 0x7e, 0x6a, 0x98, 0xcf, 0xf0, 0x10, 0xd2,
	0x12, 0x62, 0x99, 0x41, 0x8d, 0x3d, 0x9c, 0xdd, 0xa5, 0x7f, 0xb5, 0x81, 0x2e, 0x5a, 0x39, 0x5f,
	0x0f, 0x7a, 0xf7, 0xf1, 0xf5, 0x97, 0x45, 0xa0, 0xc2, 0x5c, 0x04, 0x60, 0xc0, 0x36, 0x3e, 0x38,
	0x13, 0xb1, 0x86, 0x5c, 0xd9, 0x68, 0xb2, 0x3b, 0xb5, 0xb8, 0x39, 0x7a, 0x9f, 0x11, 0x1e, 0x1a,
	0x08, 0x99, 0x63, 0x6b, 0xbd, 0x78, 0x7b, 0x01, 0x67, 0x6b, 0x83, 0x57, 0x46, 0x31, 0x1f, 0x2e,
	0xbf, 0x8d, 0x07, 0x1f, 0xbf, 0x8f, 0x11, 0xef, 0xc6, 0xc8, 0x6d, 0x6c, 0xa5, 0x7e, 0x02, 0x2a,
	0xf3, 0x43, 0xa8, 0x5d, 0xb0, 0x78, 0xf7, 0xa0, 0x73, 0x6d, 0xf7, 0xb7, 0xae, 0x5d, 0xeb, 0x75,
	0x6d, 0xf6, 0x69, 0x07, 0xef, 0x2f, 0xea, 0xfd, 0xc9, 0x09, 0x3e, 0x68, 0x43, 0x21, 0x0f, 0x7a,
	0x7c, 0xba, 0x1a, 0x9e, 0x73, 0x63, 0xeb, 0x0d, 0x8b, 0xd5, 0xd7, 0xb0, 0x22, 0xb6, 0x91, 0xf4,
	0x12, 0xaf, 0x46, 0xf7, 0x47, 0x62, 0x84, 0xad, 0x75, 0x1a, 0x84, 0xf5, 0x30, 0x37, 0x73, 0x73,
	0xfe, 0x35, 0xfe, 0x87, 0x68, 0xfe, 0x66, 0x79, 0xe9, 0x0e, 0xbe, 0x5e, 0xba, 0x83, 0x0f, 0x95,
	0x8b, 0x96, 0x95, 0x8b, 0xbe, 0x54, 0x2e, 0xfa, 0x51, 0xb9, 0xe8, 0xed, 0x93, 0xff, 0xfc, 0x1f,
	0x1f, 0x37, 0x55, 0xb0, 0x5f, 0xaf, 0xf4, 0xe8, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x5d,
	0x09, 0xd6, 0xd8, 0x03, 0x00, 0x00,
}
