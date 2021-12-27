// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: thorchain/v1/x/thorchain/types/type_observed_tx.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	common "gitlab.com/thorchain/thornode/common"
	gitlab_com_thorchain_thornode_common "gitlab.com/thorchain/thornode/common"
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

type Status int32

const (
	Status_incomplete Status = 0
	Status_done       Status = 1
	Status_reverted   Status = 2
)

var Status_name = map[int32]string{
	0: "incomplete",
	1: "done",
	2: "reverted",
}

var Status_value = map[string]int32{
	"incomplete": 0,
	"done":       1,
	"reverted":   2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_05a7dd8a7022d1fb, []int{0}
}

type ObservedTx struct {
	Tx             common.Tx                                   `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx"`
	Status         Status                                      `protobuf:"varint,2,opt,name=status,proto3,enum=types.Status" json:"status,omitempty"`
	OutHashes      []string                                    `protobuf:"bytes,3,rep,name=out_hashes,json=outHashes,proto3" json:"out_hashes,omitempty"`
	BlockHeight    int64                                       `protobuf:"varint,4,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Signers        []string                                    `protobuf:"bytes,5,rep,name=signers,proto3" json:"signers,omitempty"`
	ObservedPubKey gitlab_com_thorchain_thornode_common.PubKey `protobuf:"bytes,6,opt,name=observed_pub_key,json=observedPubKey,proto3,casttype=gitlab.com/thorchain/thornode/common.PubKey" json:"observed_pub_key,omitempty"`
	KeysignMs      int64                                       `protobuf:"varint,7,opt,name=keysign_ms,json=keysignMs,proto3" json:"keysign_ms,omitempty"`
	FinaliseHeight int64                                       `protobuf:"varint,8,opt,name=finalise_height,json=finaliseHeight,proto3" json:"finalise_height,omitempty"`
}

func (m *ObservedTx) Reset()      { *m = ObservedTx{} }
func (*ObservedTx) ProtoMessage() {}
func (*ObservedTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a7dd8a7022d1fb, []int{0}
}
func (m *ObservedTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObservedTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObservedTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObservedTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObservedTx.Merge(m, src)
}
func (m *ObservedTx) XXX_Size() int {
	return m.Size()
}
func (m *ObservedTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ObservedTx.DiscardUnknown(m)
}

var xxx_messageInfo_ObservedTx proto.InternalMessageInfo

type ObservedTxVoter struct {
	TxID            gitlab_com_thorchain_thornode_common.TxID `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,casttype=gitlab.com/thorchain/thornode/common.TxID" json:"tx_id,omitempty"`
	Tx              ObservedTx                                `protobuf:"bytes,2,opt,name=tx,proto3" json:"tx"`
	Height          int64                                     `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Txs             ObservedTxs                               `protobuf:"bytes,4,rep,name=txs,proto3,castrepeated=ObservedTxs" json:"txs"`
	Actions         []TxOutItem                               `protobuf:"bytes,5,rep,name=actions,proto3" json:"actions"`
	OutTxs          gitlab_com_thorchain_thornode_common.Txs  `protobuf:"bytes,6,rep,name=out_txs,json=outTxs,proto3,castrepeated=gitlab.com/thorchain/thornode/common.Txs" json:"out_txs"`
	FinalisedHeight int64                                     `protobuf:"varint,7,opt,name=finalised_height,json=finalisedHeight,proto3" json:"finalised_height,omitempty"`
	UpdatedVault    bool                                      `protobuf:"varint,8,opt,name=updated_vault,json=updatedVault,proto3" json:"updated_vault,omitempty"`
	Reverted        bool                                      `protobuf:"varint,9,opt,name=reverted,proto3" json:"reverted,omitempty"`
	OutboundHeight  int64                                     `protobuf:"varint,10,opt,name=outbound_height,json=outboundHeight,proto3" json:"outbound_height,omitempty"`
}

func (m *ObservedTxVoter) Reset()      { *m = ObservedTxVoter{} }
func (*ObservedTxVoter) ProtoMessage() {}
func (*ObservedTxVoter) Descriptor() ([]byte, []int) {
	return fileDescriptor_05a7dd8a7022d1fb, []int{1}
}
func (m *ObservedTxVoter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObservedTxVoter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObservedTxVoter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObservedTxVoter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObservedTxVoter.Merge(m, src)
}
func (m *ObservedTxVoter) XXX_Size() int {
	return m.Size()
}
func (m *ObservedTxVoter) XXX_DiscardUnknown() {
	xxx_messageInfo_ObservedTxVoter.DiscardUnknown(m)
}

var xxx_messageInfo_ObservedTxVoter proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("types.Status", Status_name, Status_value)
	proto.RegisterType((*ObservedTx)(nil), "types.ObservedTx")
	proto.RegisterType((*ObservedTxVoter)(nil), "types.ObservedTxVoter")
}

func init() {
	proto.RegisterFile("thorchain/v1/x/thorchain/types/type_observed_tx.proto", fileDescriptor_05a7dd8a7022d1fb)
}

var fileDescriptor_05a7dd8a7022d1fb = []byte{
	// 635 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0xb5, 0xe3, 0xd4, 0x49, 0x6e, 0xda, 0x34, 0x6f, 0xde, 0xd3, 0x93, 0x55, 0x09, 0xc7, 0x14,
	0xa1, 0xa6, 0x20, 0x25, 0xa1, 0xa8, 0x12, 0xeb, 0x88, 0x45, 0x0b, 0x42, 0x45, 0x26, 0xad, 0x04,
	0x1b, 0xcb, 0x8e, 0x87, 0xd8, 0x6a, 0xe2, 0x89, 0x3c, 0xd7, 0x91, 0xb3, 0xe3, 0x13, 0xf8, 0x0e,
	0xbe, 0xa4, 0xcb, 0x2e, 0x58, 0x74, 0x15, 0x68, 0xfa, 0x0f, 0x2c, 0xba, 0x42, 0x1e, 0x8f, 0x93,
	0x22, 0x10, 0xea, 0x26, 0xf1, 0x3d, 0x73, 0xef, 0x3d, 0x47, 0xe7, 0x8c, 0x0d, 0x87, 0x18, 0xb0,
	0x78, 0x18, 0xb8, 0x61, 0xd4, 0x9d, 0x3d, 0xeb, 0xa6, 0xdd, 0x75, 0x89, 0xf3, 0x29, 0xe5, 0xe2,
	0xd7, 0x61, 0x1e, 0xa7, 0xf1, 0x8c, 0xfa, 0x0e, 0xa6, 0x9d, 0x69, 0xcc, 0x90, 0x91, 0x0d, 0x71,
	0xba, 0x63, 0xfd, 0x32, 0x3d, 0x64, 0x93, 0x09, 0x8b, 0xe4, 0x5f, 0xde, 0xb8, 0xd3, 0xbb, 0xcf,
	0x7e, 0x4c, 0x1d, 0x96, 0xa0, 0x9c, 0xf8, 0x6f, 0xc4, 0x46, 0x4c, 0x3c, 0x76, 0xb3, 0xa7, 0x1c,
	0xdd, 0xfd, 0x5a, 0x02, 0x38, 0x91, 0x32, 0x06, 0x29, 0xb1, 0xa0, 0x84, 0xa9, 0xa1, 0x5a, 0x6a,
	0xbb, 0x7e, 0x00, 0x1d, 0xc9, 0x38, 0x48, 0xfb, 0xe5, 0x8b, 0x45, 0x4b, 0xb1, 0x4b, 0x98, 0x92,
	0xc7, 0xa0, 0x73, 0x74, 0x31, 0xe1, 0x46, 0xc9, 0x52, 0xdb, 0x8d, 0x83, 0xad, 0x8e, 0x20, 0xec,
	0xbc, 0x13, 0xa0, 0x2d, 0x0f, 0xc9, 0x03, 0x00, 0x96, 0xa0, 0x13, 0xb8, 0x3c, 0xa0, 0xdc, 0xd0,
	0x2c, 0xad, 0x5d, 0xb3, 0x6b, 0x2c, 0xc1, 0x23, 0x01, 0x90, 0x87, 0xb0, 0xe9, 0x8d, 0xd9, 0xf0,
	0xdc, 0x09, 0x68, 0x38, 0x0a, 0xd0, 0x28, 0x5b, 0x6a, 0x5b, 0xb3, 0xeb, 0x02, 0x3b, 0x12, 0x10,
	0x31, 0xa0, 0xc2, 0xc3, 0x51, 0x44, 0x63, 0x6e, 0x6c, 0x88, 0xf1, 0xa2, 0x24, 0xef, 0xa1, 0xb9,
	0x72, 0x6e, 0x9a, 0x78, 0xce, 0x39, 0x9d, 0x1b, 0xba, 0xa5, 0xb6, 0x6b, 0xfd, 0xee, 0xed, 0xa2,
	0xf5, 0x74, 0x14, 0xe2, 0xd8, 0xf5, 0x32, 0xf1, 0x77, 0x5d, 0x09, 0x58, 0x1c, 0x31, 0x9f, 0x16,
	0x2e, 0xbe, 0x4d, 0xbc, 0xd7, 0x74, 0x6e, 0x37, 0x8a, 0x45, 0x79, 0x9d, 0xc9, 0x3e, 0xa7, 0xf3,
	0x8c, 0xc8, 0x99, 0x70, 0xa3, 0x22, 0x54, 0xd5, 0x24, 0xf2, 0x86, 0x93, 0x3d, 0xd8, 0xfe, 0x18,
	0x46, 0xee, 0x38, 0xe4, 0xb4, 0x50, 0x5e, 0x15, 0x3d, 0x8d, 0x02, 0xce, 0xc5, 0xef, 0xfe, 0xd0,
	0x60, 0x7b, 0x6d, 0xeb, 0x19, 0x43, 0x1a, 0x93, 0x57, 0xb0, 0x81, 0xa9, 0x13, 0xfa, 0xc2, 0xde,
	0x5a, 0xff, 0x70, 0xb9, 0x68, 0x95, 0x07, 0xe9, 0xf1, 0xcb, 0xdb, 0x45, 0x6b, 0xff, 0x5e, 0x9a,
	0xb3, 0x66, 0xbb, 0x8c, 0xe9, 0xb1, 0x4f, 0xf6, 0x44, 0x4e, 0x25, 0x91, 0xd3, 0x3f, 0x32, 0x81,
	0x35, 0xdf, 0x9d, 0xb8, 0xfe, 0x07, 0x5d, 0x0a, 0xd5, 0x84, 0x50, 0x59, 0x91, 0x17, 0xa0, 0x61,
	0xca, 0x8d, 0xb2, 0xa5, 0xfd, 0x79, 0xc3, 0xbf, 0xd9, 0x86, 0x2f, 0xdf, 0x5a, 0xf5, 0x35, 0xc6,
	0xed, 0x6c, 0x84, 0xf4, 0xa0, 0xe2, 0x0e, 0x31, 0x64, 0x51, 0x9e, 0x4b, 0xfd, 0xa0, 0x29, 0xa7,
	0x07, 0xe9, 0x49, 0x82, 0xc7, 0x48, 0x27, 0x92, 0xbe, 0x68, 0x23, 0xa7, 0x50, 0xc9, 0xee, 0x42,
	0xc6, 0xa7, 0x8b, 0x89, 0xbb, 0x37, 0xab, 0x27, 0x89, 0xda, 0xf7, 0xb4, 0x80, 0xdb, 0x3a, 0x4b,
	0x70, 0x90, 0x72, 0xb2, 0x0f, 0xcd, 0xc2, 0x75, 0xbf, 0x48, 0x23, 0x4f, 0x6c, 0x15, 0x92, 0x2f,
	0xef, 0xd2, 0x23, 0xd8, 0x4a, 0xa6, 0xbe, 0x8b, 0xd4, 0x77, 0x66, 0x6e, 0x32, 0xce, 0x53, 0xab,
	0xda, 0x9b, 0x12, 0x3c, 0xcb, 0x30, 0xb2, 0x03, 0xd5, 0x98, 0xce, 0x68, 0x8c, 0xd4, 0x37, 0x6a,
	0xe2, 0x7c, 0x55, 0x67, 0xc1, 0xb3, 0x04, 0x3d, 0x96, 0x44, 0x2b, 0x2a, 0xc8, 0x83, 0x2f, 0xe0,
	0x9c, 0xe9, 0x49, 0x0f, 0xf4, 0xfc, 0x4d, 0x20, 0x0d, 0x80, 0x30, 0x1a, 0xb2, 0xc9, 0x74, 0x4c,
	0x91, 0x36, 0x15, 0x52, 0x85, 0xb2, 0xcf, 0x22, 0xda, 0x54, 0xc9, 0xe6, 0x9a, 0xa8, 0x59, 0xea,
	0x9f, 0x5e, 0x5c, 0x9b, 0xca, 0xd5, 0xb5, 0xa9, 0x7c, 0x5a, 0x9a, 0xca, 0xc5, 0xd2, 0x54, 0x2f,
	0x97, 0xa6, 0xfa, 0x7d, 0x69, 0xaa, 0x9f, 0x6f, 0x4c, 0xe5, 0xf2, 0xc6, 0x54, 0xae, 0x6e, 0x4c,
	0xe5, 0x43, 0xf7, 0xef, 0xf6, 0xfc, 0xf6, 0x01, 0xf0, 0x74, 0xf1, 0x7e, 0x3f, 0xff, 0x19, 0x00,
	0x00, 0xff, 0xff, 0xcb, 0x72, 0x73, 0xdc, 0x89, 0x04, 0x00, 0x00,
}

func (m *ObservedTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObservedTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObservedTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FinaliseHeight != 0 {
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(m.FinaliseHeight))
		i--
		dAtA[i] = 0x40
	}
	if m.KeysignMs != 0 {
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(m.KeysignMs))
		i--
		dAtA[i] = 0x38
	}
	if len(m.ObservedPubKey) > 0 {
		i -= len(m.ObservedPubKey)
		copy(dAtA[i:], m.ObservedPubKey)
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(len(m.ObservedPubKey)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Signers) > 0 {
		for iNdEx := len(m.Signers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Signers[iNdEx])
			copy(dAtA[i:], m.Signers[iNdEx])
			i = encodeVarintTypeObservedTx(dAtA, i, uint64(len(m.Signers[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.OutHashes) > 0 {
		for iNdEx := len(m.OutHashes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.OutHashes[iNdEx])
			copy(dAtA[i:], m.OutHashes[iNdEx])
			i = encodeVarintTypeObservedTx(dAtA, i, uint64(len(m.OutHashes[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Status != 0 {
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ObservedTxVoter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObservedTxVoter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObservedTxVoter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OutboundHeight != 0 {
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(m.OutboundHeight))
		i--
		dAtA[i] = 0x50
	}
	if m.Reverted {
		i--
		if m.Reverted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.UpdatedVault {
		i--
		if m.UpdatedVault {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.FinalisedHeight != 0 {
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(m.FinalisedHeight))
		i--
		dAtA[i] = 0x38
	}
	if len(m.OutTxs) > 0 {
		for iNdEx := len(m.OutTxs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.OutTxs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypeObservedTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Actions) > 0 {
		for iNdEx := len(m.Actions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Actions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypeObservedTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Txs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypeObservedTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Height != 0 {
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	{
		size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.TxID) > 0 {
		i -= len(m.TxID)
		copy(dAtA[i:], m.TxID)
		i = encodeVarintTypeObservedTx(dAtA, i, uint64(len(m.TxID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypeObservedTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypeObservedTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ObservedTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Tx.Size()
	n += 1 + l + sovTypeObservedTx(uint64(l))
	if m.Status != 0 {
		n += 1 + sovTypeObservedTx(uint64(m.Status))
	}
	if len(m.OutHashes) > 0 {
		for _, s := range m.OutHashes {
			l = len(s)
			n += 1 + l + sovTypeObservedTx(uint64(l))
		}
	}
	if m.BlockHeight != 0 {
		n += 1 + sovTypeObservedTx(uint64(m.BlockHeight))
	}
	if len(m.Signers) > 0 {
		for _, s := range m.Signers {
			l = len(s)
			n += 1 + l + sovTypeObservedTx(uint64(l))
		}
	}
	l = len(m.ObservedPubKey)
	if l > 0 {
		n += 1 + l + sovTypeObservedTx(uint64(l))
	}
	if m.KeysignMs != 0 {
		n += 1 + sovTypeObservedTx(uint64(m.KeysignMs))
	}
	if m.FinaliseHeight != 0 {
		n += 1 + sovTypeObservedTx(uint64(m.FinaliseHeight))
	}
	return n
}

func (m *ObservedTxVoter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TxID)
	if l > 0 {
		n += 1 + l + sovTypeObservedTx(uint64(l))
	}
	l = m.Tx.Size()
	n += 1 + l + sovTypeObservedTx(uint64(l))
	if m.Height != 0 {
		n += 1 + sovTypeObservedTx(uint64(m.Height))
	}
	if len(m.Txs) > 0 {
		for _, e := range m.Txs {
			l = e.Size()
			n += 1 + l + sovTypeObservedTx(uint64(l))
		}
	}
	if len(m.Actions) > 0 {
		for _, e := range m.Actions {
			l = e.Size()
			n += 1 + l + sovTypeObservedTx(uint64(l))
		}
	}
	if len(m.OutTxs) > 0 {
		for _, e := range m.OutTxs {
			l = e.Size()
			n += 1 + l + sovTypeObservedTx(uint64(l))
		}
	}
	if m.FinalisedHeight != 0 {
		n += 1 + sovTypeObservedTx(uint64(m.FinalisedHeight))
	}
	if m.UpdatedVault {
		n += 2
	}
	if m.Reverted {
		n += 2
	}
	if m.OutboundHeight != 0 {
		n += 1 + sovTypeObservedTx(uint64(m.OutboundHeight))
	}
	return n
}

func sovTypeObservedTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypeObservedTx(x uint64) (n int) {
	return sovTypeObservedTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ObservedTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeObservedTx
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
			return fmt.Errorf("proto: ObservedTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObservedTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
				return ErrInvalidLengthTypeObservedTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutHashes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
				return ErrInvalidLengthTypeObservedTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutHashes = append(m.OutHashes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
				return ErrInvalidLengthTypeObservedTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signers = append(m.Signers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservedPubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
				return ErrInvalidLengthTypeObservedTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObservedPubKey = gitlab_com_thorchain_thornode_common.PubKey(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeysignMs", wireType)
			}
			m.KeysignMs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.KeysignMs |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinaliseHeight", wireType)
			}
			m.FinaliseHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinaliseHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypeObservedTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeObservedTx
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
func (m *ObservedTxVoter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypeObservedTx
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
			return fmt.Errorf("proto: ObservedTxVoter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObservedTxVoter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
				return ErrInvalidLengthTypeObservedTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxID = gitlab_com_thorchain_thornode_common.TxID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
				return ErrInvalidLengthTypeObservedTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
				return ErrInvalidLengthTypeObservedTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, ObservedTx{})
			if err := m.Txs[len(m.Txs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Actions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
				return ErrInvalidLengthTypeObservedTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Actions = append(m.Actions, TxOutItem{})
			if err := m.Actions[len(m.Actions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutTxs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
				return ErrInvalidLengthTypeObservedTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutTxs = append(m.OutTxs, common.Tx{})
			if err := m.OutTxs[len(m.OutTxs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalisedHeight", wireType)
			}
			m.FinalisedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FinalisedHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedVault", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
			m.UpdatedVault = bool(v != 0)
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reverted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
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
			m.Reverted = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutboundHeight", wireType)
			}
			m.OutboundHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypeObservedTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OutboundHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypeObservedTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypeObservedTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypeObservedTx
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
func skipTypeObservedTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypeObservedTx
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
					return 0, ErrIntOverflowTypeObservedTx
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
					return 0, ErrIntOverflowTypeObservedTx
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
				return 0, ErrInvalidLengthTypeObservedTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypeObservedTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypeObservedTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypeObservedTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypeObservedTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypeObservedTx = fmt.Errorf("proto: unexpected end of group")
)