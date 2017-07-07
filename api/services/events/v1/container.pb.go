// Code generated by protoc-gen-gogo.
// source: github.com/containerd/containerd/api/services/events/v1/container.proto
// DO NOT EDIT!

/*
	Package events is a generated protocol buffer package.

	It is generated from these files:
		github.com/containerd/containerd/api/services/events/v1/container.proto
		github.com/containerd/containerd/api/services/events/v1/content.proto
		github.com/containerd/containerd/api/services/events/v1/events.proto
		github.com/containerd/containerd/api/services/events/v1/image.proto
		github.com/containerd/containerd/api/services/events/v1/namespace.proto
		github.com/containerd/containerd/api/services/events/v1/runtime.proto
		github.com/containerd/containerd/api/services/events/v1/snapshot.proto
		github.com/containerd/containerd/api/services/events/v1/task.proto

	It has these top-level messages:
		ContainerCreate
		ContainerUpdate
		ContainerDelete
		ContentDelete
		StreamEventsRequest
		PostEventRequest
		Envelope
		ImageUpdate
		ImageDelete
		NamespaceCreate
		NamespaceUpdate
		NamespaceDelete
		RuntimeIO
		RuntimeMount
		RuntimeCreate
		RuntimeEvent
		RuntimeDelete
		SnapshotPrepare
		SnapshotCommit
		SnapshotRemove
		TaskCreate
		TaskStart
		TaskDelete
*/
package events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

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

type ContainerCreate struct {
	ID      string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image   string                   `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Runtime *ContainerCreate_Runtime `protobuf:"bytes,3,opt,name=runtime" json:"runtime,omitempty"`
}

func (m *ContainerCreate) Reset()                    { *m = ContainerCreate{} }
func (*ContainerCreate) ProtoMessage()               {}
func (*ContainerCreate) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{0} }

type ContainerCreate_Runtime struct {
	Name    string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Options *google_protobuf1.Any `protobuf:"bytes,2,opt,name=options" json:"options,omitempty"`
}

func (m *ContainerCreate_Runtime) Reset()      { *m = ContainerCreate_Runtime{} }
func (*ContainerCreate_Runtime) ProtoMessage() {}
func (*ContainerCreate_Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptorContainer, []int{0, 0}
}

type ContainerUpdate struct {
	ID     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Image  string            `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RootFS string            `protobuf:"bytes,4,opt,name=rootfs,proto3" json:"rootfs,omitempty"`
}

func (m *ContainerUpdate) Reset()                    { *m = ContainerUpdate{} }
func (*ContainerUpdate) ProtoMessage()               {}
func (*ContainerUpdate) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{1} }

type ContainerDelete struct {
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *ContainerDelete) Reset()                    { *m = ContainerDelete{} }
func (*ContainerDelete) ProtoMessage()               {}
func (*ContainerDelete) Descriptor() ([]byte, []int) { return fileDescriptorContainer, []int{2} }

func init() {
	proto.RegisterType((*ContainerCreate)(nil), "containerd.services.events.v1.ContainerCreate")
	proto.RegisterType((*ContainerCreate_Runtime)(nil), "containerd.services.events.v1.ContainerCreate.Runtime")
	proto.RegisterType((*ContainerUpdate)(nil), "containerd.services.events.v1.ContainerUpdate")
	proto.RegisterType((*ContainerDelete)(nil), "containerd.services.events.v1.ContainerDelete")
}
func (m *ContainerCreate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerCreate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Image) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.Image)))
		i += copy(dAtA[i:], m.Image)
	}
	if m.Runtime != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Runtime.Size()))
		n1, err := m.Runtime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *ContainerCreate_Runtime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerCreate_Runtime) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Options != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintContainer(dAtA, i, uint64(m.Options.Size()))
		n2, err := m.Options.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *ContainerUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerUpdate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Image) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.Image)))
		i += copy(dAtA[i:], m.Image)
	}
	if len(m.Labels) > 0 {
		for k, _ := range m.Labels {
			dAtA[i] = 0x1a
			i++
			v := m.Labels[k]
			mapSize := 1 + len(k) + sovContainer(uint64(len(k))) + 1 + len(v) + sovContainer(uint64(len(v)))
			i = encodeVarintContainer(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintContainer(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintContainer(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	if len(m.RootFS) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.RootFS)))
		i += copy(dAtA[i:], m.RootFS)
	}
	return i, nil
}

func (m *ContainerDelete) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContainerDelete) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContainer(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	return i, nil
}

func encodeFixed64Container(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Container(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintContainer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ContainerCreate) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	if m.Runtime != nil {
		l = m.Runtime.Size()
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func (m *ContainerCreate_Runtime) Size() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	if m.Options != nil {
		l = m.Options.Size()
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func (m *ContainerUpdate) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovContainer(uint64(len(k))) + 1 + len(v) + sovContainer(uint64(len(v)))
			n += mapEntrySize + 1 + sovContainer(uint64(mapEntrySize))
		}
	}
	l = len(m.RootFS)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func (m *ContainerDelete) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovContainer(uint64(l))
	}
	return n
}

func sovContainer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozContainer(x uint64) (n int) {
	return sovContainer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ContainerCreate) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContainerCreate{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`Runtime:` + strings.Replace(fmt.Sprintf("%v", this.Runtime), "ContainerCreate_Runtime", "ContainerCreate_Runtime", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContainerCreate_Runtime) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContainerCreate_Runtime{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Options:` + strings.Replace(fmt.Sprintf("%v", this.Options), "Any", "google_protobuf1.Any", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContainerUpdate) String() string {
	if this == nil {
		return "nil"
	}
	keysForLabels := make([]string, 0, len(this.Labels))
	for k, _ := range this.Labels {
		keysForLabels = append(keysForLabels, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForLabels)
	mapStringForLabels := "map[string]string{"
	for _, k := range keysForLabels {
		mapStringForLabels += fmt.Sprintf("%v: %v,", k, this.Labels[k])
	}
	mapStringForLabels += "}"
	s := strings.Join([]string{`&ContainerUpdate{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`Image:` + fmt.Sprintf("%v", this.Image) + `,`,
		`Labels:` + mapStringForLabels + `,`,
		`RootFS:` + fmt.Sprintf("%v", this.RootFS) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContainerDelete) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContainerDelete{`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringContainer(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ContainerCreate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: ContainerCreate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerCreate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Runtime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Runtime == nil {
				m.Runtime = &ContainerCreate_Runtime{}
			}
			if err := m.Runtime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
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
func (m *ContainerCreate_Runtime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: Runtime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Runtime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Options == nil {
				m.Options = &google_protobuf1.Any{}
			}
			if err := m.Options.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
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
func (m *ContainerUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: ContainerUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthContainer
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(dAtA[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowContainer
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowContainer
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthContainer
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(dAtA[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Labels[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Labels[mapkey] = mapvalue
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootFS", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootFS = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
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
func (m *ContainerDelete) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContainer
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
			return fmt.Errorf("proto: ContainerDelete: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContainerDelete: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContainer
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
				return ErrInvalidLengthContainer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContainer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContainer
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
func skipContainer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContainer
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
					return 0, ErrIntOverflowContainer
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
					return 0, ErrIntOverflowContainer
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
				return 0, ErrInvalidLengthContainer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowContainer
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
				next, err := skipContainer(dAtA[start:])
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
	ErrInvalidLengthContainer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContainer   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/containerd/containerd/api/services/events/v1/container.proto", fileDescriptorContainer)
}

var fileDescriptorContainer = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x41, 0x8b, 0xd4, 0x30,
	0x18, 0xdd, 0xb4, 0x6b, 0x07, 0xd3, 0x83, 0x12, 0x06, 0xa9, 0x05, 0xbb, 0x43, 0x4f, 0xe3, 0x25,
	0x65, 0x47, 0x10, 0x5d, 0x41, 0x70, 0x77, 0x55, 0x04, 0x05, 0x89, 0x08, 0xe2, 0x2d, 0x9d, 0x66,
	0x6a, 0xb0, 0x4d, 0x4a, 0x9b, 0x16, 0x7a, 0xf3, 0xe7, 0xcd, 0xd1, 0xa3, 0xa7, 0x61, 0xa6, 0x3f,
	0xc1, 0x5f, 0x20, 0x4d, 0x5a, 0xa7, 0x08, 0x8a, 0x7a, 0x7b, 0x5f, 0xbe, 0xf7, 0xbe, 0xef, 0xbd,
	0x24, 0xf0, 0x65, 0xca, 0xd5, 0xa7, 0x3a, 0xc6, 0x6b, 0x99, 0x47, 0x6b, 0x29, 0x14, 0xe5, 0x82,
	0x95, 0xc9, 0x14, 0xd2, 0x82, 0x47, 0x15, 0x2b, 0x1b, 0xbe, 0x66, 0x55, 0xc4, 0x1a, 0x26, 0x54,
	0x15, 0x35, 0xe7, 0x47, 0x06, 0x2e, 0x4a, 0xa9, 0x24, 0xba, 0x77, 0x94, 0xe0, 0x91, 0x8e, 0x0d,
	0x1d, 0x37, 0xe7, 0xfe, 0x3c, 0x95, 0xa9, 0xd4, 0xcc, 0xa8, 0x47, 0x46, 0xe4, 0xdf, 0x4d, 0xa5,
	0x4c, 0x33, 0x16, 0xe9, 0x2a, 0xae, 0x37, 0x11, 0x15, 0xad, 0x69, 0x85, 0x7b, 0x00, 0x6f, 0x5d,
	0x8d, 0x23, 0xaf, 0x4a, 0x46, 0x15, 0x43, 0x77, 0xa0, 0xc5, 0x13, 0x0f, 0x2c, 0xc0, 0xf2, 0xe6,
	0xa5, 0xd3, 0xed, 0xce, 0xac, 0x57, 0xd7, 0xc4, 0xe2, 0x09, 0x9a, 0xc3, 0x1b, 0x3c, 0xa7, 0x29,
	0xf3, 0xac, 0xbe, 0x45, 0x4c, 0x81, 0xde, 0xc2, 0x59, 0x59, 0x0b, 0xc5, 0x73, 0xe6, 0xd9, 0x0b,
	0xb0, 0x74, 0x57, 0x0f, 0xf1, 0x1f, 0x3d, 0xe2, 0x5f, 0xd6, 0x61, 0x62, 0xd4, 0x64, 0x1c, 0xe3,
	0xbf, 0x81, 0xb3, 0xe1, 0x0c, 0x21, 0x78, 0x2a, 0x68, 0xce, 0x8c, 0x19, 0xa2, 0x31, 0xc2, 0x70,
	0x26, 0x0b, 0xc5, 0xa5, 0xa8, 0xb4, 0x11, 0x77, 0x35, 0xc7, 0x26, 0x1f, 0x1e, 0xf3, 0xe1, 0x67,
	0xa2, 0x25, 0x23, 0x29, 0xfc, 0x3e, 0x8d, 0xf8, 0xbe, 0x48, 0xfe, 0x3d, 0x22, 0x81, 0x4e, 0x46,
	0x63, 0x96, 0x55, 0x9e, 0xbd, 0xb0, 0x97, 0xee, 0xea, 0xe2, 0x6f, 0x13, 0x9a, 0x6d, 0xf8, 0xb5,
	0x16, 0x3f, 0x17, 0xaa, 0x6c, 0xc9, 0x30, 0x09, 0x85, 0xd0, 0x29, 0xa5, 0x54, 0x9b, 0xca, 0x3b,
	0xd5, 0x2e, 0x60, 0xb7, 0x3b, 0x73, 0x88, 0x94, 0xea, 0xc5, 0x3b, 0x32, 0x74, 0xfc, 0xc7, 0xd0,
	0x9d, 0x48, 0xd1, 0x6d, 0x68, 0x7f, 0x66, 0xed, 0x70, 0x17, 0x3d, 0xec, 0xed, 0x36, 0x34, 0xab,
	0x7f, 0xda, 0xd5, 0xc5, 0x85, 0xf5, 0x08, 0x84, 0xf7, 0x27, 0x99, 0xaf, 0x59, 0xc6, 0x7e, 0x9f,
	0xf9, 0xf2, 0xc3, 0xf6, 0x10, 0x9c, 0x7c, 0x3b, 0x04, 0x27, 0x5f, 0xba, 0x00, 0x6c, 0xbb, 0x00,
	0x7c, 0xed, 0x02, 0xb0, 0xef, 0x02, 0xf0, 0xf1, 0xe9, 0x7f, 0xfe, 0xda, 0x27, 0x06, 0xc5, 0x8e,
	0x7e, 0x90, 0x07, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0xeb, 0xf5, 0x3f, 0xfe, 0x02, 0x00,
	0x00,
}
