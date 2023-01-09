// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/expiration/v1/events.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// EventTxCompleted is an event message indicating that a TX has completed.
type EventTxCompleted struct {
	// module is the module the TX belongs to.
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// endpoint is the rpc endpoint that was just completed.
	Endpoint string `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// signers are the bech32 address strings of the signers of this TX.
	Signers []string `protobuf:"bytes,3,rep,name=signers,proto3" json:"signers,omitempty"`
}

func (m *EventTxCompleted) Reset()         { *m = EventTxCompleted{} }
func (m *EventTxCompleted) String() string { return proto.CompactTextString(m) }
func (*EventTxCompleted) ProtoMessage()    {}
func (*EventTxCompleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_db63118c179972b3, []int{0}
}
func (m *EventTxCompleted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventTxCompleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventTxCompleted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventTxCompleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTxCompleted.Merge(m, src)
}
func (m *EventTxCompleted) XXX_Size() int {
	return m.Size()
}
func (m *EventTxCompleted) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTxCompleted.DiscardUnknown(m)
}

var xxx_messageInfo_EventTxCompleted proto.InternalMessageInfo

func (m *EventTxCompleted) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *EventTxCompleted) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *EventTxCompleted) GetSigners() []string {
	if m != nil {
		return m.Signers
	}
	return nil
}

// EventExpirationAdd is an event message indicating an expiration has been created
type EventExpirationAdd struct {
	// module_asset_id is the bech32 address string of the expiration id that was created.
	ModuleAssetId string `protobuf:"bytes,1,opt,name=module_asset_id,json=moduleAssetId,proto3" json:"module_asset_id,omitempty"`
}

func (m *EventExpirationAdd) Reset()         { *m = EventExpirationAdd{} }
func (m *EventExpirationAdd) String() string { return proto.CompactTextString(m) }
func (*EventExpirationAdd) ProtoMessage()    {}
func (*EventExpirationAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_db63118c179972b3, []int{1}
}
func (m *EventExpirationAdd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventExpirationAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventExpirationAdd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventExpirationAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventExpirationAdd.Merge(m, src)
}
func (m *EventExpirationAdd) XXX_Size() int {
	return m.Size()
}
func (m *EventExpirationAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_EventExpirationAdd.DiscardUnknown(m)
}

var xxx_messageInfo_EventExpirationAdd proto.InternalMessageInfo

func (m *EventExpirationAdd) GetModuleAssetId() string {
	if m != nil {
		return m.ModuleAssetId
	}
	return ""
}

// EventExpirationDeposit is an event message indicating a deposit was collected for expiration
type EventExpirationDeposit struct {
	// module_asset_id is the bech32 address string of the expiration id that was created.
	ModuleAssetId string `protobuf:"bytes,1,opt,name=module_asset_id,json=moduleAssetId,proto3" json:"module_asset_id,omitempty"`
	// bech32 address string of the account that provided the expiration deposit.
	Depositor string `protobuf:"bytes,2,opt,name=depositor,proto3" json:"depositor,omitempty"`
	// deposit amount held until asset is expired
	Deposit types.Coin `protobuf:"bytes,3,opt,name=deposit,proto3" json:"deposit"`
}

func (m *EventExpirationDeposit) Reset()         { *m = EventExpirationDeposit{} }
func (m *EventExpirationDeposit) String() string { return proto.CompactTextString(m) }
func (*EventExpirationDeposit) ProtoMessage()    {}
func (*EventExpirationDeposit) Descriptor() ([]byte, []int) {
	return fileDescriptor_db63118c179972b3, []int{2}
}
func (m *EventExpirationDeposit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventExpirationDeposit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventExpirationDeposit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventExpirationDeposit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventExpirationDeposit.Merge(m, src)
}
func (m *EventExpirationDeposit) XXX_Size() int {
	return m.Size()
}
func (m *EventExpirationDeposit) XXX_DiscardUnknown() {
	xxx_messageInfo_EventExpirationDeposit.DiscardUnknown(m)
}

var xxx_messageInfo_EventExpirationDeposit proto.InternalMessageInfo

func (m *EventExpirationDeposit) GetModuleAssetId() string {
	if m != nil {
		return m.ModuleAssetId
	}
	return ""
}

func (m *EventExpirationDeposit) GetDepositor() string {
	if m != nil {
		return m.Depositor
	}
	return ""
}

func (m *EventExpirationDeposit) GetDeposit() types.Coin {
	if m != nil {
		return m.Deposit
	}
	return types.Coin{}
}

// EventExpirationExtend is an event message indicating an expiration has been extended
type EventExpirationExtend struct {
	// module_asset_id is the bech32 address string of the expiration id that was created.
	ModuleAssetId string `protobuf:"bytes,1,opt,name=module_asset_id,json=moduleAssetId,proto3" json:"module_asset_id,omitempty"`
}

func (m *EventExpirationExtend) Reset()         { *m = EventExpirationExtend{} }
func (m *EventExpirationExtend) String() string { return proto.CompactTextString(m) }
func (*EventExpirationExtend) ProtoMessage()    {}
func (*EventExpirationExtend) Descriptor() ([]byte, []int) {
	return fileDescriptor_db63118c179972b3, []int{3}
}
func (m *EventExpirationExtend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventExpirationExtend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventExpirationExtend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventExpirationExtend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventExpirationExtend.Merge(m, src)
}
func (m *EventExpirationExtend) XXX_Size() int {
	return m.Size()
}
func (m *EventExpirationExtend) XXX_DiscardUnknown() {
	xxx_messageInfo_EventExpirationExtend.DiscardUnknown(m)
}

var xxx_messageInfo_EventExpirationExtend proto.InternalMessageInfo

func (m *EventExpirationExtend) GetModuleAssetId() string {
	if m != nil {
		return m.ModuleAssetId
	}
	return ""
}

// EventExpirationInvoke is an event message indicating an expiration has been invoked
type EventExpirationInvoke struct {
	// module_asset_id is the bech32 address string of the expiration id that was created.
	ModuleAssetId string `protobuf:"bytes,1,opt,name=module_asset_id,json=moduleAssetId,proto3" json:"module_asset_id,omitempty"`
}

func (m *EventExpirationInvoke) Reset()         { *m = EventExpirationInvoke{} }
func (m *EventExpirationInvoke) String() string { return proto.CompactTextString(m) }
func (*EventExpirationInvoke) ProtoMessage()    {}
func (*EventExpirationInvoke) Descriptor() ([]byte, []int) {
	return fileDescriptor_db63118c179972b3, []int{4}
}
func (m *EventExpirationInvoke) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventExpirationInvoke) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventExpirationInvoke.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventExpirationInvoke) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventExpirationInvoke.Merge(m, src)
}
func (m *EventExpirationInvoke) XXX_Size() int {
	return m.Size()
}
func (m *EventExpirationInvoke) XXX_DiscardUnknown() {
	xxx_messageInfo_EventExpirationInvoke.DiscardUnknown(m)
}

var xxx_messageInfo_EventExpirationInvoke proto.InternalMessageInfo

func (m *EventExpirationInvoke) GetModuleAssetId() string {
	if m != nil {
		return m.ModuleAssetId
	}
	return ""
}

// EventExpirationRemove is an event message indicating an expiration has been removed
type EventExpirationRemove struct {
	// module_asset_id is the bech32 address string of the expiration id that was created.
	ModuleAssetId string `protobuf:"bytes,1,opt,name=module_asset_id,json=moduleAssetId,proto3" json:"module_asset_id,omitempty"`
}

func (m *EventExpirationRemove) Reset()         { *m = EventExpirationRemove{} }
func (m *EventExpirationRemove) String() string { return proto.CompactTextString(m) }
func (*EventExpirationRemove) ProtoMessage()    {}
func (*EventExpirationRemove) Descriptor() ([]byte, []int) {
	return fileDescriptor_db63118c179972b3, []int{5}
}
func (m *EventExpirationRemove) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventExpirationRemove) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventExpirationRemove.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventExpirationRemove) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventExpirationRemove.Merge(m, src)
}
func (m *EventExpirationRemove) XXX_Size() int {
	return m.Size()
}
func (m *EventExpirationRemove) XXX_DiscardUnknown() {
	xxx_messageInfo_EventExpirationRemove.DiscardUnknown(m)
}

var xxx_messageInfo_EventExpirationRemove proto.InternalMessageInfo

func (m *EventExpirationRemove) GetModuleAssetId() string {
	if m != nil {
		return m.ModuleAssetId
	}
	return ""
}

func init() {
	proto.RegisterType((*EventTxCompleted)(nil), "provenance.expiration.v1.EventTxCompleted")
	proto.RegisterType((*EventExpirationAdd)(nil), "provenance.expiration.v1.EventExpirationAdd")
	proto.RegisterType((*EventExpirationDeposit)(nil), "provenance.expiration.v1.EventExpirationDeposit")
	proto.RegisterType((*EventExpirationExtend)(nil), "provenance.expiration.v1.EventExpirationExtend")
	proto.RegisterType((*EventExpirationInvoke)(nil), "provenance.expiration.v1.EventExpirationInvoke")
	proto.RegisterType((*EventExpirationRemove)(nil), "provenance.expiration.v1.EventExpirationRemove")
}

func init() {
	proto.RegisterFile("provenance/expiration/v1/events.proto", fileDescriptor_db63118c179972b3)
}

var fileDescriptor_db63118c179972b3 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0xcd, 0x58, 0x79, 0xcf, 0x8e, 0x88, 0x32, 0xe8, 0x23, 0x56, 0x89, 0x25, 0xa0, 0x74, 0xe3,
	0x0c, 0x51, 0x10, 0x04, 0x41, 0xde, 0x7b, 0x76, 0xd1, 0x9d, 0x04, 0x57, 0x6e, 0x6a, 0x92, 0xb9,
	0xc4, 0xc1, 0x66, 0x6e, 0xc8, 0x4c, 0x43, 0xfc, 0x17, 0x6e, 0xfd, 0x47, 0x5d, 0x76, 0xe9, 0x4a,
	0xa4, 0xfd, 0x23, 0x92, 0x8f, 0x36, 0xa5, 0xba, 0x28, 0xbc, 0xdd, 0x3d, 0x67, 0xce, 0x3d, 0xf7,
	0xce, 0xe5, 0xd0, 0xe7, 0x79, 0x81, 0x25, 0xe8, 0x48, 0x27, 0x20, 0xa0, 0xca, 0x55, 0x11, 0x59,
	0x85, 0x5a, 0x94, 0x81, 0x80, 0x12, 0xb4, 0x35, 0x3c, 0x2f, 0xd0, 0x22, 0x73, 0x7b, 0x19, 0xef,
	0x65, 0xbc, 0x0c, 0x46, 0x5e, 0x82, 0x26, 0x43, 0x23, 0xe2, 0xc8, 0x80, 0x28, 0x83, 0x18, 0x6c,
	0x14, 0x88, 0x04, 0x95, 0x6e, 0x3b, 0x47, 0x0f, 0x53, 0x4c, 0xb1, 0x29, 0x45, 0x5d, 0xb5, 0xac,
	0xff, 0x85, 0x3e, 0x98, 0xd6, 0xfe, 0x9f, 0xaa, 0x6b, 0xcc, 0xf2, 0x05, 0x58, 0x90, 0xec, 0x82,
	0x9e, 0x65, 0x28, 0x97, 0x0b, 0x70, 0xc9, 0x98, 0x4c, 0x86, 0x61, 0x87, 0xd8, 0x88, 0xde, 0x01,
	0x2d, 0x73, 0x54, 0xda, 0xba, 0xb7, 0x9a, 0x97, 0x3d, 0x66, 0x2e, 0x3d, 0x37, 0x2a, 0xd5, 0x50,
	0x18, 0x77, 0x30, 0x1e, 0x4c, 0x86, 0xe1, 0x0e, 0xfa, 0xef, 0x28, 0x6b, 0x26, 0x4c, 0xf7, 0xdb,
	0x5e, 0x4a, 0xc9, 0x5e, 0xd0, 0xfb, 0xad, 0xeb, 0x3c, 0x32, 0x06, 0xec, 0x5c, 0xc9, 0x6e, 0xd8,
	0xbd, 0x96, 0xbe, 0xac, 0xd9, 0x99, 0xf4, 0x7f, 0x12, 0x7a, 0x71, 0xd4, 0xfe, 0x01, 0x72, 0x34,
	0xca, 0x9e, 0x6a, 0xc1, 0x9e, 0xd2, 0xa1, 0x6c, 0x5b, 0xb0, 0xe8, 0xf6, 0xee, 0x09, 0xf6, 0x96,
	0x9e, 0x77, 0xc0, 0x1d, 0x8c, 0xc9, 0xe4, 0xee, 0xab, 0xc7, 0xbc, 0x3d, 0x24, 0xaf, 0x0f, 0xc9,
	0xbb, 0x43, 0xf2, 0x6b, 0x54, 0xfa, 0xea, 0xf6, 0xea, 0xf7, 0x33, 0x27, 0xdc, 0xe9, 0xfd, 0xf7,
	0xf4, 0xd1, 0xd1, 0x6a, 0xd3, 0xca, 0x82, 0x3e, 0xfd, 0x73, 0xff, 0x1a, 0xcc, 0x74, 0x89, 0xdf,
	0xe0, 0x06, 0x06, 0x21, 0x64, 0x58, 0x9e, 0x6c, 0x70, 0x85, 0xab, 0x8d, 0x47, 0xd6, 0x1b, 0x8f,
	0xfc, 0xd9, 0x78, 0xe4, 0xc7, 0xd6, 0x73, 0xd6, 0x5b, 0xcf, 0xf9, 0xb5, 0xf5, 0x1c, 0xfa, 0x44,
	0x35, 0x11, 0xf9, 0x6f, 0xd6, 0x3e, 0x92, 0xcf, 0x6f, 0x52, 0x65, 0xbf, 0x2e, 0x63, 0x9e, 0x60,
	0x26, 0x7a, 0xd9, 0x4b, 0x85, 0x07, 0x48, 0x54, 0x87, 0x49, 0xb6, 0xdf, 0x73, 0x30, 0xf1, 0x59,
	0x13, 0xbb, 0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x62, 0x8f, 0xdb, 0x33, 0xef, 0x02, 0x00,
	0x00,
}

func (m *EventTxCompleted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventTxCompleted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventTxCompleted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Endpoint) > 0 {
		i -= len(m.Endpoint)
		copy(dAtA[i:], m.Endpoint)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Endpoint)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Module) > 0 {
		i -= len(m.Module)
		copy(dAtA[i:], m.Module)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Module)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventExpirationAdd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventExpirationAdd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventExpirationAdd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModuleAssetId) > 0 {
		i -= len(m.ModuleAssetId)
		copy(dAtA[i:], m.ModuleAssetId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ModuleAssetId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventExpirationDeposit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventExpirationDeposit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventExpirationDeposit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Deposit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Depositor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ModuleAssetId) > 0 {
		i -= len(m.ModuleAssetId)
		copy(dAtA[i:], m.ModuleAssetId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ModuleAssetId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventExpirationExtend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventExpirationExtend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventExpirationExtend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModuleAssetId) > 0 {
		i -= len(m.ModuleAssetId)
		copy(dAtA[i:], m.ModuleAssetId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ModuleAssetId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventExpirationInvoke) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventExpirationInvoke) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventExpirationInvoke) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModuleAssetId) > 0 {
		i -= len(m.ModuleAssetId)
		copy(dAtA[i:], m.ModuleAssetId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ModuleAssetId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventExpirationRemove) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventExpirationRemove) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventExpirationRemove) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ModuleAssetId) > 0 {
		i -= len(m.ModuleAssetId)
		copy(dAtA[i:], m.ModuleAssetId)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.ModuleAssetId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventTxCompleted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Module)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Endpoint)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventExpirationAdd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleAssetId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventExpirationDeposit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleAssetId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = m.Deposit.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventExpirationExtend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleAssetId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventExpirationInvoke) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleAssetId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventExpirationRemove) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ModuleAssetId)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventTxCompleted) Unmarshal(dAtA []byte) error {
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
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventTxCompleted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventTxCompleted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Module", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Module = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoint", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *EventExpirationAdd) Unmarshal(dAtA []byte) error {
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
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventExpirationAdd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventExpirationAdd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAssetId", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleAssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *EventExpirationDeposit) Unmarshal(dAtA []byte) error {
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
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventExpirationDeposit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventExpirationDeposit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAssetId", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleAssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
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
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Deposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *EventExpirationExtend) Unmarshal(dAtA []byte) error {
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
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventExpirationExtend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventExpirationExtend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAssetId", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleAssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *EventExpirationInvoke) Unmarshal(dAtA []byte) error {
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
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventExpirationInvoke: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventExpirationInvoke: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAssetId", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleAssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *EventExpirationRemove) Unmarshal(dAtA []byte) error {
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
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EventExpirationRemove: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventExpirationRemove: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModuleAssetId", wireType)
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
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ModuleAssetId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
	depth := 0
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
		case 1:
			iNdEx += 8
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
			if length < 0 {
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
