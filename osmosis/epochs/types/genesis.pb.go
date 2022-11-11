// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/epochs/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// EpochInfo is a struct that describes the data going into
// a timer defined by the x/epochs module.
type EpochInfo struct {
	// identifier is a unique reference to this particular timer.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// start_time is the time at which the timer first ever ticks.
	// If start_time is in the future, the epoch will not begin until the start
	// time.
	StartTime time.Time `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	// duration is the time in between epoch ticks.
	// In order for intended behavior to be met, duration should
	// be greater than the chains expected block time.
	// Duration must be non-zero.
	Duration time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	// current_epoch is the current epoch number, or in other words,
	// how many times has the timer 'ticked'.
	// The first tick (current_epoch=1) is defined as
	// the first block whose blocktime is greater than the EpochInfo start_time.
	CurrentEpoch int64 `protobuf:"varint,4,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
	// current_epoch_start_time describes the start time of the current timer
	// interval. The interval is (current_epoch_start_time,
	// current_epoch_start_time + duration] When the timer ticks, this is set to
	// current_epoch_start_time = last_epoch_start_time + duration only one timer
	// tick for a given identifier can occur per block.
	//
	// NOTE! The current_epoch_start_time may diverge significantly from the
	// wall-clock time the epoch began at. Wall-clock time of epoch start may be
	// >> current_epoch_start_time. Suppose current_epoch_start_time = 10,
	// duration = 5. Suppose the chain goes offline at t=14, and comes back online
	// at t=30, and produces blocks at every successive time. (t=31, 32, etc.)
	// * The t=30 block will start the epoch for (10, 15]
	// * The t=31 block will start the epoch for (15, 20]
	// * The t=32 block will start the epoch for (20, 25]
	// * The t=33 block will start the epoch for (25, 30]
	// * The t=34 block will start the epoch for (30, 35]
	// * The **t=36** block will start the epoch for (35, 40]
	CurrentEpochStartTime time.Time `protobuf:"bytes,5,opt,name=current_epoch_start_time,json=currentEpochStartTime,proto3,stdtime" json:"current_epoch_start_time" yaml:"current_epoch_start_time"`
	// epoch_counting_started is a boolean, that indicates whether this
	// epoch timer has began yet.
	EpochCountingStarted bool `protobuf:"varint,6,opt,name=epoch_counting_started,json=epochCountingStarted,proto3" json:"epoch_counting_started,omitempty"`
	// current_epoch_start_height is the block height at which the current epoch
	// started. (The block height at which the timer last ticked)
	CurrentEpochStartHeight int64 `protobuf:"varint,8,opt,name=current_epoch_start_height,json=currentEpochStartHeight,proto3" json:"current_epoch_start_height,omitempty"`
}

func (m *EpochInfo) Reset()         { *m = EpochInfo{} }
func (m *EpochInfo) String() string { return proto.CompactTextString(m) }
func (*EpochInfo) ProtoMessage()    {}
func (*EpochInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ecf3e4d59074cbd, []int{0}
}
func (m *EpochInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochInfo.Merge(m, src)
}
func (m *EpochInfo) XXX_Size() int {
	return m.Size()
}
func (m *EpochInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EpochInfo proto.InternalMessageInfo

func (m *EpochInfo) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *EpochInfo) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *EpochInfo) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *EpochInfo) GetCurrentEpoch() int64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func (m *EpochInfo) GetCurrentEpochStartTime() time.Time {
	if m != nil {
		return m.CurrentEpochStartTime
	}
	return time.Time{}
}

func (m *EpochInfo) GetEpochCountingStarted() bool {
	if m != nil {
		return m.EpochCountingStarted
	}
	return false
}

func (m *EpochInfo) GetCurrentEpochStartHeight() int64 {
	if m != nil {
		return m.CurrentEpochStartHeight
	}
	return 0
}

// GenesisState defines the epochs module's genesis state.
type GenesisState struct {
	Epochs []EpochInfo `protobuf:"bytes,1,rep,name=epochs,proto3" json:"epochs"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ecf3e4d59074cbd, []int{1}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetEpochs() []EpochInfo {
	if m != nil {
		return m.Epochs
	}
	return nil
}

func init() {
	proto.RegisterType((*EpochInfo)(nil), "osmosis.epochs.v1beta1.EpochInfo")
	proto.RegisterType((*GenesisState)(nil), "osmosis.epochs.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("osmosis/epochs/genesis.proto", fileDescriptor_7ecf3e4d59074cbd) }

var fileDescriptor_7ecf3e4d59074cbd = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x3f, 0x8f, 0xd3, 0x30,
	0x1c, 0xad, 0x69, 0x29, 0xa9, 0xef, 0x10, 0x60, 0x1d, 0x47, 0xa8, 0x20, 0x09, 0x61, 0x89, 0x04,
	0x72, 0xd4, 0x03, 0x16, 0x18, 0x90, 0x0a, 0x88, 0x3f, 0xcb, 0x49, 0x29, 0x03, 0x62, 0xa9, 0x9c,
	0xd6, 0x4d, 0x2c, 0x35, 0x71, 0x88, 0x1d, 0xa4, 0x6e, 0x7c, 0x84, 0x8e, 0x7c, 0xa4, 0x1b, 0x6f,
	0x64, 0x0a, 0xa8, 0xdd, 0x18, 0xef, 0x13, 0xa0, 0xd8, 0x49, 0x29, 0x77, 0x45, 0xb7, 0x39, 0x7e,
	0xef, 0xf7, 0x9e, 0xdf, 0xd3, 0x2f, 0xf0, 0x1e, 0x17, 0x09, 0x17, 0x4c, 0xf8, 0x34, 0xe3, 0x93,
	0x58, 0xf8, 0x11, 0x4d, 0xa9, 0x60, 0x02, 0x67, 0x39, 0x97, 0x1c, 0x1d, 0xd6, 0x28, 0xd6, 0x28,
	0xfe, 0x3a, 0x08, 0xa9, 0x24, 0x83, 0xfe, 0x41, 0xc4, 0x23, 0xae, 0x28, 0x7e, 0x75, 0xd2, 0xec,
	0xbe, 0x15, 0x71, 0x1e, 0xcd, 0xa9, 0xaf, 0xbe, 0xc2, 0x62, 0xe6, 0x4f, 0x8b, 0x9c, 0x48, 0xc6,
	0xd3, 0x1a, 0xb7, 0xcf, 0xe3, 0x92, 0x25, 0x54, 0x48, 0x92, 0x64, 0x9a, 0xe0, 0x2e, 0x3b, 0xb0,
	0xf7, 0xa6, 0x72, 0x7a, 0x9f, 0xce, 0x38, 0xb2, 0x20, 0x64, 0x53, 0x9a, 0x4a, 0x36, 0x63, 0x34,
	0x37, 0x81, 0x03, 0xbc, 0x5e, 0xb0, 0x75, 0x83, 0x3e, 0x41, 0x28, 0x24, 0xc9, 0xe5, 0xb8, 0x92,
	0x31, 0xaf, 0x38, 0xc0, 0xdb, 0x3b, 0xea, 0x63, 0xed, 0x81, 0x1b, 0x0f, 0xfc, 0xb1, 0xf1, 0x18,
	0xde, 0x3f, 0x29, 0xed, 0xd6, 0x59, 0x69, 0xdf, 0x5a, 0x90, 0x64, 0xfe, 0xdc, 0xfd, 0x3b, 0xeb,
	0x2e, 0x7f, 0xda, 0x20, 0xe8, 0xa9, 0x8b, 0x8a, 0x8e, 0x62, 0x68, 0x34, 0x4f, 0x37, 0xdb, 0x4a,
	0xf7, 0xee, 0x05, 0xdd, 0xd7, 0x35, 0x61, 0x38, 0xa8, 0x64, 0x7f, 0x97, 0x36, 0x6a, 0x46, 0x1e,
	0xf3, 0x84, 0x49, 0x9a, 0x64, 0x72, 0x71, 0x56, 0xda, 0x37, 0xb4, 0x59, 0x83, 0xb9, 0xdf, 0x2b,
	0xab, 0x8d, 0x3a, 0x7a, 0x08, 0xaf, 0x4f, 0x8a, 0x3c, 0xa7, 0xa9, 0x1c, 0xab, 0x8a, 0xcd, 0x8e,
	0x03, 0xbc, 0x76, 0xb0, 0x5f, 0x5f, 0xaa, 0x32, 0xd0, 0x37, 0x00, 0xcd, 0x7f, 0x58, 0xe3, 0xad,
	0xdc, 0x57, 0x2f, 0xcd, 0xfd, 0xa8, 0xce, 0x6d, 0xeb, 0xa7, 0xfc, 0x4f, 0x49, 0xb7, 0x70, 0x7b,
	0xdb, 0x79, 0xb4, 0x69, 0xe4, 0x29, 0x3c, 0xd4, 0xfc, 0x09, 0x2f, 0x52, 0xc9, 0xd2, 0x48, 0x0f,
	0xd2, 0xa9, 0xd9, 0x75, 0x80, 0x67, 0x04, 0x07, 0x0a, 0x7d, 0x55, 0x83, 0x23, 0x8d, 0xa1, 0x17,
	0xb0, 0xbf, 0xcb, 0x2d, 0xa6, 0x2c, 0x8a, 0xa5, 0x69, 0xa8, 0xa8, 0x77, 0x2e, 0x18, 0xbe, 0x53,
	0xf0, 0x87, 0x8e, 0x71, 0xed, 0xa6, 0xe1, 0x1e, 0xc3, 0xfd, 0xb7, 0x7a, 0x25, 0x47, 0x92, 0x48,
	0x8a, 0x5e, 0xc2, 0xae, 0xde, 0x45, 0x13, 0x38, 0x6d, 0x6f, 0xef, 0xe8, 0x01, 0xde, 0xbd, 0xa2,
	0x78, 0xb3, 0x47, 0xc3, 0x4e, 0x95, 0x3f, 0xa8, 0xc7, 0x86, 0xc7, 0x27, 0x2b, 0x0b, 0x9c, 0xae,
	0x2c, 0xf0, 0x6b, 0x65, 0x81, 0xe5, 0xda, 0x6a, 0x9d, 0xae, 0xad, 0xd6, 0x8f, 0xb5, 0xd5, 0xfa,
	0xfc, 0x2c, 0x62, 0x32, 0x2e, 0x42, 0x3c, 0xe1, 0x89, 0xff, 0xa5, 0x20, 0x82, 0xe4, 0x73, 0x12,
	0x8a, 0xfa, 0x98, 0xf2, 0x29, 0xf5, 0xcf, 0xfd, 0x2b, 0x72, 0x91, 0x51, 0x11, 0x76, 0x55, 0xe5,
	0x4f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x99, 0x0c, 0x63, 0x4a, 0x03, 0x00, 0x00,
}

func (m *EpochInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpochStartHeight != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CurrentEpochStartHeight))
		i--
		dAtA[i] = 0x40
	}
	if m.EpochCountingStarted {
		i--
		if m.EpochCountingStarted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CurrentEpochStartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.CurrentEpoch != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Epochs) > 0 {
		for iNdEx := len(m.Epochs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Epochs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpochInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovGenesis(uint64(l))
	if m.CurrentEpoch != 0 {
		n += 1 + sovGenesis(uint64(m.CurrentEpoch))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CurrentEpochStartTime)
	n += 1 + l + sovGenesis(uint64(l))
	if m.EpochCountingStarted {
		n += 2
	}
	if m.CurrentEpochStartHeight != 0 {
		n += 1 + sovGenesis(uint64(m.CurrentEpochStartHeight))
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Epochs) > 0 {
		for _, e := range m.Epochs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpochInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: EpochInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CurrentEpochStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochCountingStarted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EpochCountingStarted = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartHeight", wireType)
			}
			m.CurrentEpochStartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpochStartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epochs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Epochs = append(m.Epochs, EpochInfo{})
			if err := m.Epochs[len(m.Epochs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
