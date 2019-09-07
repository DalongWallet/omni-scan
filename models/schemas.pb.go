package models

import "github.com/golang/protobuf/proto"

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Transaction) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *Transaction) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *Transaction) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *Transaction) GetBlockHeight() int64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func init() { proto.RegisterFile("schemas.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x55, 0xfe, 0xbf, 0xdc, 0x7e, 0x85, 0xca, 0x03, 0xf2, 0xc0, 0x10, 0x65, 0x21, 0xea, 0xd0,
	0x01, 0x5e, 0x00, 0xc1, 0x42, 0x85, 0x04, 0x95, 0x15, 0x46, 0x86, 0x34, 0xb6, 0x68, 0xa0, 0x89,
	0xa3, 0xd8, 0x41, 0xe9, 0x83, 0xf2, 0x06, 0x3c, 0x08, 0x8a, 0x93, 0x50, 0xa7, 0x65, 0xe8, 0x76,
	0xcf, 0x89, 0xed, 0x7b, 0xee, 0x39, 0x37, 0x30, 0x15, 0xe9, 0x86, 0xe5, 0x89, 0x58, 0x94, 0x15,
	0x97, 0x1c, 0xb9, 0x39, 0xa7, 0x6c, 0x2b, 0xc2, 0x77, 0xb0, 0x5f, 0x64, 0xc3, 0x11, 0x02, 0x5b,
	0x36, 0x4b, 0x8a, 0x8d, 0xc0, 0x88, 0x7c, 0xa2, 0xea, 0x96, 0xfb, 0xe4, 0xb5, 0xc4, 0x66, 0x60,
	0x44, 0x53, 0xa2, 0x6a, 0x74, 0x01, 0x6e, 0x92, 0xf3, 0xba, 0x90, 0xd8, 0x0a, 0x8c, 0xc8, 0x22,
	0x3d, 0x42, 0x21, 0xfc, 0x17, 0x69, 0x95, 0x95, 0x72, 0x55, 0xaf, 0x1f, 0xd9, 0x0e, 0xdb, 0xea,
	0x9d, 0x11, 0x17, 0x7e, 0x19, 0xe0, 0xdd, 0xf3, 0x42, 0xb2, 0x46, 0xa2, 0x00, 0x26, 0xeb, 0x2d,
	0x4f, 0x3f, 0x1e, 0x58, 0xf6, 0xb6, 0x91, 0xaa, 0xad, 0x45, 0x74, 0x0a, 0x5d, 0x82, 0xdf, 0xc1,
	0x44, 0x6c, 0x94, 0x04, 0x9f, 0xec, 0x89, 0xb6, 0x9f, 0x02, 0x31, 0x97, 0xc9, 0x36, 0x6e, 0x94,
	0x1a, 0x87, 0x8c, 0x38, 0x34, 0x87, 0xd9, 0x5d, 0x8b, 0x57, 0x15, 0x4f, 0x99, 0x10, 0x8c, 0xc6,
	0x8d, 0xd2, 0xe5, 0x90, 0x23, 0x1e, 0x61, 0xf0, 0x64, 0xff, 0x94, 0xa3, 0xb4, 0x0c, 0xb0, 0x55,
	0xda, 0x97, 0x4f, 0x9c, 0x32, 0xec, 0x76, 0x4a, 0x35, 0x2a, 0x7c, 0x05, 0x2f, 0x6e, 0x96, 0x45,
	0x59, 0xcb, 0x93, 0x6d, 0x44, 0x60, 0x27, 0x94, 0x56, 0x4a, 0xb6, 0x4f, 0x54, 0xad, 0x59, 0x6b,
	0xeb, 0xd6, 0x86, 0x05, 0xfc, 0x8b, 0x9b, 0xe7, 0x5a, 0x96, 0xda, 0x3d, 0xe3, 0xcf, 0x7b, 0xe6,
	0x28, 0x92, 0xa1, 0xaf, 0xa5, 0xf5, 0x3d, 0x25, 0xa6, 0x6f, 0x03, 0x26, 0x71, 0x95, 0x14, 0x22,
	0x49, 0x65, 0xc6, 0x0b, 0x74, 0x06, 0x66, 0x46, 0xfb, 0x84, 0xcc, 0x8c, 0xfe, 0xce, 0x68, 0x6a,
	0x33, 0x0e, 0x61, 0xc5, 0x59, 0xce, 0xfa, 0xcd, 0xd8, 0x13, 0xe3, 0x28, 0xed, 0xc3, 0x28, 0x0f,
	0x56, 0xc1, 0x39, 0x5e, 0x85, 0x2b, 0x70, 0xb3, 0xd6, 0x5e, 0x81, 0xdd, 0xc0, 0x8a, 0x26, 0xd7,
	0xe7, 0x8b, 0x6e, 0x7b, 0x17, 0xbd, 0xed, 0xa4, 0xff, 0x8c, 0xe6, 0xe0, 0x71, 0x65, 0x94, 0xc0,
	0x9e, 0x3a, 0x39, 0xdb, 0x9f, 0xec, 0x1c, 0x24, 0xc3, 0x81, 0xf0, 0x16, 0xdc, 0x2e, 0xbf, 0x93,
	0x06, 0x6c, 0xb9, 0x5d, 0xc9, 0x06, 0x33, 0xdb, 0x7a, 0xed, 0xaa, 0x5f, 0xe9, 0xe6, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x6c, 0xa4, 0x3e, 0x1d, 0x5b, 0x03, 0x00, 0x00,
}