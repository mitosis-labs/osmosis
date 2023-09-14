// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/incentives/group_gauge.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/osmosis-labs/osmosis/v19/x/lockup/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SplittingPolicy determines the way we want to split incentives in groupGauges
type SplittingPolicy int32

const (
	Volume SplittingPolicy = 0
	Evenly SplittingPolicy = 1
)

var SplittingPolicy_name = map[int32]string{
	0: "Volume",
	1: "Evenly",
}

var SplittingPolicy_value = map[string]int32{
	"Volume": 0,
	"Evenly": 1,
}

func (x SplittingPolicy) String() string {
	return proto.EnumName(SplittingPolicy_name, int32(x))
}

func (SplittingPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_693357a19afe66cf, []int{0}
}

// Note that while both InternalGaugeInfo and InternalGaugeRecord could
// technically be replaced by DistrInfo and DistrRecord from the pool-incentives
// module, we create separate types here to keep our abstractions clean and
// readable (pool-incentives distribution abstractions are used in a very
// specific way that does not directly relate to gauge logic). This also helps
// us sidestep a refactor to avoid an import cycle.
type InternalGaugeInfo struct {
	TotalWeight  cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=total_weight,json=totalWeight,proto3,customtype=cosmossdk.io/math.Int" json:"total_weight" yaml:"total_weight"`
	GaugeRecords []InternalGaugeRecord `protobuf:"bytes,2,rep,name=gauge_records,json=gaugeRecords,proto3" json:"gauge_records"`
}

func (m *InternalGaugeInfo) Reset()         { *m = InternalGaugeInfo{} }
func (m *InternalGaugeInfo) String() string { return proto.CompactTextString(m) }
func (*InternalGaugeInfo) ProtoMessage()    {}
func (*InternalGaugeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_693357a19afe66cf, []int{0}
}
func (m *InternalGaugeInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InternalGaugeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InternalGaugeInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InternalGaugeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalGaugeInfo.Merge(m, src)
}
func (m *InternalGaugeInfo) XXX_Size() int {
	return m.Size()
}
func (m *InternalGaugeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalGaugeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InternalGaugeInfo proto.InternalMessageInfo

func (m *InternalGaugeInfo) GetGaugeRecords() []InternalGaugeRecord {
	if m != nil {
		return m.GaugeRecords
	}
	return nil
}

type InternalGaugeRecord struct {
	GaugeId          uint64                `protobuf:"varint,1,opt,name=gauge_id,json=gaugeId,proto3" json:"gauge_id,omitempty" yaml:"gauge_id"`
	CurrentWeight    cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=current_weight,json=currentWeight,proto3,customtype=cosmossdk.io/math.Int" json:"current_weight"`
	CumulativeWeight cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=cumulative_weight,json=cumulativeWeight,proto3,customtype=cosmossdk.io/math.Int" json:"cumulative_weight"`
}

func (m *InternalGaugeRecord) Reset()         { *m = InternalGaugeRecord{} }
func (m *InternalGaugeRecord) String() string { return proto.CompactTextString(m) }
func (*InternalGaugeRecord) ProtoMessage()    {}
func (*InternalGaugeRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_693357a19afe66cf, []int{1}
}
func (m *InternalGaugeRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InternalGaugeRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InternalGaugeRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InternalGaugeRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalGaugeRecord.Merge(m, src)
}
func (m *InternalGaugeRecord) XXX_Size() int {
	return m.Size()
}
func (m *InternalGaugeRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalGaugeRecord.DiscardUnknown(m)
}

var xxx_messageInfo_InternalGaugeRecord proto.InternalMessageInfo

func (m *InternalGaugeRecord) GetGaugeId() uint64 {
	if m != nil {
		return m.GaugeId
	}
	return 0
}

// GroupGauge is an object that stores GroupGaugeId, a list of gauge info, and a
// splitting policy. These are grouped into a single abstraction to allow for
// distribution of group incentives to internal gauges according to the
// specified splitting policy.
type GroupGauge struct {
	GroupGaugeId      uint64            `protobuf:"varint,1,opt,name=group_gauge_id,json=groupGaugeId,proto3" json:"group_gauge_id,omitempty"`
	InternalGaugeInfo InternalGaugeInfo `protobuf:"bytes,2,opt,name=internal_gauge_info,json=internalGaugeInfo,proto3" json:"internal_gauge_info"`
	SplittingPolicy   SplittingPolicy   `protobuf:"varint,3,opt,name=splitting_policy,json=splittingPolicy,proto3,enum=osmosis.incentives.SplittingPolicy" json:"splitting_policy,omitempty"`
}

func (m *GroupGauge) Reset()         { *m = GroupGauge{} }
func (m *GroupGauge) String() string { return proto.CompactTextString(m) }
func (*GroupGauge) ProtoMessage()    {}
func (*GroupGauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_693357a19afe66cf, []int{2}
}
func (m *GroupGauge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupGauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupGauge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupGauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupGauge.Merge(m, src)
}
func (m *GroupGauge) XXX_Size() int {
	return m.Size()
}
func (m *GroupGauge) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupGauge.DiscardUnknown(m)
}

var xxx_messageInfo_GroupGauge proto.InternalMessageInfo

func (m *GroupGauge) GetGroupGaugeId() uint64 {
	if m != nil {
		return m.GroupGaugeId
	}
	return 0
}

func (m *GroupGauge) GetInternalGaugeInfo() InternalGaugeInfo {
	if m != nil {
		return m.InternalGaugeInfo
	}
	return InternalGaugeInfo{}
}

func (m *GroupGauge) GetSplittingPolicy() SplittingPolicy {
	if m != nil {
		return m.SplittingPolicy
	}
	return Volume
}

func init() {
	proto.RegisterEnum("osmosis.incentives.SplittingPolicy", SplittingPolicy_name, SplittingPolicy_value)
	proto.RegisterType((*InternalGaugeInfo)(nil), "osmosis.incentives.InternalGaugeInfo")
	proto.RegisterType((*InternalGaugeRecord)(nil), "osmosis.incentives.InternalGaugeRecord")
	proto.RegisterType((*GroupGauge)(nil), "osmosis.incentives.GroupGauge")
}

func init() {
	proto.RegisterFile("osmosis/incentives/group_gauge.proto", fileDescriptor_693357a19afe66cf)
}

var fileDescriptor_693357a19afe66cf = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xb6, 0xdb, 0xa8, 0xc0, 0x26, 0xcd, 0x8f, 0x03, 0x52, 0x89, 0x84, 0x5d, 0x99, 0x22, 0x2a,
	0x24, 0xbc, 0x4a, 0x40, 0x48, 0xf4, 0x18, 0x81, 0xaa, 0x70, 0x40, 0x95, 0x91, 0xa8, 0x04, 0x87,
	0xc8, 0x3f, 0x9b, 0xcd, 0xaa, 0xeb, 0x5d, 0xcb, 0xbb, 0x0e, 0xe4, 0x0d, 0x38, 0xf2, 0x08, 0x48,
	0xbc, 0x08, 0xc7, 0x1e, 0x7b, 0x44, 0x48, 0x04, 0x94, 0x5c, 0x38, 0xf7, 0x09, 0x90, 0xd7, 0x76,
	0x93, 0xfe, 0x08, 0x7a, 0xf2, 0x8c, 0xe7, 0x9b, 0x6f, 0xbe, 0x6f, 0x34, 0x0b, 0x76, 0xb8, 0x88,
	0xb8, 0x20, 0x02, 0x12, 0x16, 0x20, 0x26, 0xc9, 0x04, 0x09, 0x88, 0x13, 0x9e, 0xc6, 0x43, 0xec,
	0xa5, 0x18, 0x39, 0x71, 0xc2, 0x25, 0x37, 0x8c, 0x02, 0xe5, 0x2c, 0x51, 0x9d, 0xdb, 0x98, 0x63,
	0xae, 0xca, 0x30, 0x8b, 0x72, 0x64, 0xc7, 0xc4, 0x9c, 0x63, 0x8a, 0xa0, 0xca, 0xfc, 0x74, 0x04,
	0xc3, 0x34, 0xf1, 0x24, 0xe1, 0xac, 0xa8, 0x5b, 0x17, 0xeb, 0x92, 0x44, 0x48, 0x48, 0x2f, 0x8a,
	0x4b, 0x82, 0x40, 0xcd, 0x82, 0xbe, 0x27, 0x10, 0x9c, 0x74, 0x7d, 0x24, 0xbd, 0x2e, 0x0c, 0x38,
	0x29, 0x09, 0xee, 0x96, 0x82, 0x29, 0x0f, 0x8e, 0xd2, 0x58, 0x7d, 0xf2, 0x92, 0xfd, 0x4d, 0x07,
	0xad, 0x01, 0x93, 0x28, 0x61, 0x1e, 0xdd, 0xcf, 0xd4, 0x0f, 0xd8, 0x88, 0x1b, 0x87, 0xa0, 0x26,
	0xb9, 0xf4, 0xe8, 0xf0, 0x03, 0x22, 0x78, 0x2c, 0xb7, 0xf4, 0x6d, 0x7d, 0xf7, 0x56, 0xff, 0xe9,
	0xf1, 0xcc, 0xd2, 0x7e, 0xcc, 0xac, 0x3b, 0xf9, 0x38, 0x11, 0x1e, 0x39, 0x84, 0xc3, 0xc8, 0x93,
	0x63, 0x67, 0xc0, 0xe4, 0xe9, 0xcc, 0x6a, 0x4f, 0xbd, 0x88, 0xee, 0xd9, 0xab, 0xad, 0xb6, 0x5b,
	0x55, 0xe9, 0xa1, 0xca, 0x0c, 0x17, 0x6c, 0xaa, 0x1d, 0x0d, 0x13, 0x14, 0xf0, 0x24, 0x14, 0x5b,
	0x6b, 0xdb, 0xeb, 0xbb, 0xd5, 0xde, 0x43, 0xe7, 0xf2, 0xb2, 0x9c, 0x73, 0xb2, 0x5c, 0x85, 0xef,
	0x57, 0x32, 0x09, 0x6e, 0x0d, 0x2f, 0x7f, 0x09, 0xfb, 0xa7, 0x0e, 0xda, 0x57, 0x60, 0x0d, 0x07,
	0xdc, 0xcc, 0x67, 0x91, 0x50, 0x19, 0xa8, 0xf4, 0xdb, 0xa7, 0x33, 0xab, 0x91, 0x6b, 0x2c, 0x2b,
	0xb6, 0x7b, 0x43, 0x85, 0x83, 0xd0, 0x78, 0x01, 0xea, 0x41, 0x9a, 0x24, 0x88, 0xc9, 0xd2, 0xf6,
	0x9a, 0xb2, 0x7d, 0xef, 0x9f, 0xb6, 0xdd, 0xcd, 0xa2, 0xa9, 0x70, 0xf8, 0x0a, 0xb4, 0x82, 0x34,
	0x4a, 0xa9, 0x97, 0x99, 0x28, 0x89, 0xd6, 0xaf, 0x43, 0xd4, 0x5c, 0xf6, 0xe5, 0x5c, 0x7b, 0x95,
	0x3f, 0x5f, 0x2c, 0xdd, 0xfe, 0xa5, 0x03, 0xb0, 0x9f, 0x9d, 0x97, 0x32, 0x67, 0xec, 0x80, 0xfa,
	0xca, 0xb1, 0x9d, 0x99, 0x73, 0x6b, 0xf8, 0x0c, 0x33, 0x08, 0x8d, 0xf7, 0xa0, 0x4d, 0x8a, 0x9d,
	0x94, 0x40, 0x36, 0xe2, 0xca, 0x51, 0xb5, 0xf7, 0xe0, 0xbf, 0xeb, 0xce, 0xae, 0xa0, 0x58, 0x76,
	0x8b, 0x5c, 0x3a, 0x8f, 0xd7, 0xa0, 0x29, 0x62, 0x4a, 0xa4, 0x24, 0x0c, 0x0f, 0x63, 0x4e, 0x49,
	0x30, 0x55, 0x16, 0xeb, 0xbd, 0xfb, 0x57, 0x31, 0xbf, 0x29, 0xb1, 0x07, 0x0a, 0xea, 0x36, 0xc4,
	0xf9, 0x1f, 0x8f, 0x20, 0x68, 0x5c, 0xc0, 0x18, 0x00, 0x6c, 0xbc, 0xe5, 0x34, 0x8d, 0x50, 0x53,
	0xcb, 0xe2, 0x97, 0x13, 0xc4, 0xe8, 0xb4, 0xa9, 0x77, 0x2a, 0x9f, 0xbe, 0x9a, 0x5a, 0xff, 0xe0,
	0x78, 0x6e, 0xea, 0x27, 0x73, 0x53, 0xff, 0x3d, 0x37, 0xf5, 0xcf, 0x0b, 0x53, 0x3b, 0x59, 0x98,
	0xda, 0xf7, 0x85, 0xa9, 0xbd, 0x7b, 0x86, 0x89, 0x1c, 0xa7, 0xbe, 0x13, 0xf0, 0x08, 0x16, 0x52,
	0x1e, 0x53, 0xcf, 0x17, 0x65, 0x02, 0x27, 0xdd, 0xe7, 0xf0, 0xe3, 0xea, 0xcb, 0x95, 0xd3, 0x18,
	0x09, 0x7f, 0x43, 0x3d, 0x87, 0x27, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x22, 0x42, 0x79,
	0xdc, 0x03, 0x00, 0x00,
}

func (this *InternalGaugeRecord) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*InternalGaugeRecord)
	if !ok {
		that2, ok := that.(InternalGaugeRecord)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.GaugeId != that1.GaugeId {
		return false
	}
	if !this.CurrentWeight.Equal(that1.CurrentWeight) {
		return false
	}
	if !this.CumulativeWeight.Equal(that1.CumulativeWeight) {
		return false
	}
	return true
}
func (m *InternalGaugeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InternalGaugeInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InternalGaugeInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GaugeRecords) > 0 {
		for iNdEx := len(m.GaugeRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GaugeRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGroupGauge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size := m.TotalWeight.Size()
		i -= size
		if _, err := m.TotalWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGroupGauge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *InternalGaugeRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InternalGaugeRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InternalGaugeRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CumulativeWeight.Size()
		i -= size
		if _, err := m.CumulativeWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGroupGauge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.CurrentWeight.Size()
		i -= size
		if _, err := m.CurrentWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGroupGauge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.GaugeId != 0 {
		i = encodeVarintGroupGauge(dAtA, i, uint64(m.GaugeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GroupGauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupGauge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupGauge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SplittingPolicy != 0 {
		i = encodeVarintGroupGauge(dAtA, i, uint64(m.SplittingPolicy))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.InternalGaugeInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGroupGauge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.GroupGaugeId != 0 {
		i = encodeVarintGroupGauge(dAtA, i, uint64(m.GroupGaugeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGroupGauge(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroupGauge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InternalGaugeInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TotalWeight.Size()
	n += 1 + l + sovGroupGauge(uint64(l))
	if len(m.GaugeRecords) > 0 {
		for _, e := range m.GaugeRecords {
			l = e.Size()
			n += 1 + l + sovGroupGauge(uint64(l))
		}
	}
	return n
}

func (m *InternalGaugeRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GaugeId != 0 {
		n += 1 + sovGroupGauge(uint64(m.GaugeId))
	}
	l = m.CurrentWeight.Size()
	n += 1 + l + sovGroupGauge(uint64(l))
	l = m.CumulativeWeight.Size()
	n += 1 + l + sovGroupGauge(uint64(l))
	return n
}

func (m *GroupGauge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GroupGaugeId != 0 {
		n += 1 + sovGroupGauge(uint64(m.GroupGaugeId))
	}
	l = m.InternalGaugeInfo.Size()
	n += 1 + l + sovGroupGauge(uint64(l))
	if m.SplittingPolicy != 0 {
		n += 1 + sovGroupGauge(uint64(m.SplittingPolicy))
	}
	return n
}

func sovGroupGauge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroupGauge(x uint64) (n int) {
	return sovGroupGauge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InternalGaugeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupGauge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InternalGaugeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalGaugeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroupGauge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroupGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroupGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroupGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GaugeRecords = append(m.GaugeRecords, InternalGaugeRecord{})
			if err := m.GaugeRecords[len(m.GaugeRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroupGauge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroupGauge
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
func (m *InternalGaugeRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupGauge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InternalGaugeRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalGaugeRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeId", wireType)
			}
			m.GaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroupGauge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroupGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CurrentWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CumulativeWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGroupGauge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGroupGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CumulativeWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroupGauge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroupGauge
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
func (m *GroupGauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroupGauge
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GroupGauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupGauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupGaugeId", wireType)
			}
			m.GroupGaugeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupGaugeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InternalGaugeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroupGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroupGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InternalGaugeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SplittingPolicy", wireType)
			}
			m.SplittingPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroupGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SplittingPolicy |= SplittingPolicy(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGroupGauge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroupGauge
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
func skipGroupGauge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroupGauge
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
					return 0, ErrIntOverflowGroupGauge
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGroupGauge
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
			if length < 0 {
				return 0, ErrInvalidLengthGroupGauge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroupGauge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroupGauge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroupGauge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroupGauge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroupGauge = fmt.Errorf("proto: unexpected end of group")
)
