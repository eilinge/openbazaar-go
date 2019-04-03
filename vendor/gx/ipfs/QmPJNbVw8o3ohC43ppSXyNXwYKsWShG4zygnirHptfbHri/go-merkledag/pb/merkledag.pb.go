// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merkledag.proto

package merkledag_pb

import (
	bytes "bytes"
	fmt "fmt"
	io "io"
	math "math"
	reflect "reflect"
	strings "strings"

	_ "gx/ipfs/QmddjPSGZb3ieihSseFeCfVRpZzcqczPNsD2DvarSwnjJB/gogo-protobuf/gogoproto"
	proto "gx/ipfs/QmddjPSGZb3ieihSseFeCfVRpZzcqczPNsD2DvarSwnjJB/gogo-protobuf/proto"
)

// DoNotUpgradeFileEverItWillChangeYourHashes warns users about not breaking
// their file hashes.
const DoNotUpgradeFileEverItWillChangeYourHashes = `
This file does not produce canonical protobufs. Unfortunately, if we change it,
we'll change the hashes of the files we produce.

Do *not regenerate this file.
`

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// An IPFS MerkleDAG Link
type PBLink struct {
	// multihash of the target object
	Hash []byte `protobuf:"bytes,1,opt,name=Hash" json:"Hash,omitempty"`
	// utf string name. should be unique per object
	Name *string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	// cumulative size of target object
	Tsize                *uint64  `protobuf:"varint,3,opt,name=Tsize" json:"Tsize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBLink) Reset()      { *m = PBLink{} }
func (*PBLink) ProtoMessage() {}
func (*PBLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_10837cc3557cec00, []int{0}
}
func (m *PBLink) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PBLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PBLink.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PBLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBLink.Merge(m, src)
}
func (m *PBLink) XXX_Size() int {
	return m.Size()
}
func (m *PBLink) XXX_DiscardUnknown() {
	xxx_messageInfo_PBLink.DiscardUnknown(m)
}

var xxx_messageInfo_PBLink proto.InternalMessageInfo

func (m *PBLink) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PBLink) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *PBLink) GetTsize() uint64 {
	if m != nil && m.Tsize != nil {
		return *m.Tsize
	}
	return 0
}

// An IPFS MerkleDAG Node
type PBNode struct {
	// refs to other objects
	Links []*PBLink `protobuf:"bytes,2,rep,name=Links" json:"Links,omitempty"`
	// opaque user data
	Data                 []byte   `protobuf:"bytes,1,opt,name=Data" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PBNode) Reset()      { *m = PBNode{} }
func (*PBNode) ProtoMessage() {}
func (*PBNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_10837cc3557cec00, []int{1}
}
func (m *PBNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PBNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PBNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PBNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PBNode.Merge(m, src)
}
func (m *PBNode) XXX_Size() int {
	return m.Size()
}
func (m *PBNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PBNode.DiscardUnknown(m)
}

var xxx_messageInfo_PBNode proto.InternalMessageInfo

func (m *PBNode) GetLinks() []*PBLink {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *PBNode) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*PBLink)(nil), "merkledag.pb.PBLink")
	proto.RegisterType((*PBNode)(nil), "merkledag.pb.PBNode")
}

func init() { proto.RegisterFile("merkledag.proto", fileDescriptor_10837cc3557cec00) }

var fileDescriptor_10837cc3557cec00 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4d, 0x2d, 0xca,
	0xce, 0x49, 0x4d, 0x49, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x41, 0x12, 0x48,
	0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f,
	0xcf, 0xd7, 0x07, 0x2b, 0x4a, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x59, 0xc9,
	0x8d, 0x8b, 0x2d, 0xc0, 0xc9, 0x27, 0x33, 0x2f, 0x5b, 0x48, 0x88, 0x8b, 0xc5, 0x23, 0xb1, 0x38,
	0x43, 0x82, 0x51, 0x81, 0x51, 0x83, 0x27, 0x08, 0xcc, 0x06, 0x89, 0xf9, 0x25, 0xe6, 0xa6, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x22, 0x5c, 0xac, 0x21, 0xc5, 0x99, 0x55,
	0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x10, 0x8e, 0x92, 0x07, 0xc8, 0x1c, 0xbf, 0xfc,
	0x94, 0x54, 0x21, 0x2d, 0x2e, 0x56, 0x90, 0x79, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46,
	0x22, 0x7a, 0xc8, 0xce, 0xd3, 0x83, 0x58, 0x16, 0x04, 0x51, 0x02, 0x32, 0xdf, 0x25, 0xb1, 0x24,
	0x11, 0x66, 0x27, 0x88, 0xed, 0xa4, 0x73, 0xe3, 0xa1, 0x1c, 0xc3, 0x83, 0x87, 0x72, 0x8c, 0x1f,
	0x1e, 0xca, 0x31, 0xfe, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6,
	0x1d, 0x8f, 0xe4, 0x18, 0x0f, 0x3c, 0x92, 0x63, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x56, 0xb5, 0x6e, 0x0e,
	0x01, 0x00, 0x00,
}

func (this *PBLink) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*PBLink)
	if !ok {
		that2, ok := that.(PBLink)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *PBLink")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *PBLink but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *PBLink but is not nil && this == nil")
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return fmt.Errorf("Hash this(%v) Not Equal that(%v)", this.Hash, that1.Hash)
	}
	if this.Name != nil && that1.Name != nil {
		if *this.Name != *that1.Name {
			return fmt.Errorf("Name this(%v) Not Equal that(%v)", *this.Name, *that1.Name)
		}
	} else if this.Name != nil {
		return fmt.Errorf("this.Name == nil && that.Name != nil")
	} else if that1.Name != nil {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.Tsize != nil && that1.Tsize != nil {
		if *this.Tsize != *that1.Tsize {
			return fmt.Errorf("Tsize this(%v) Not Equal that(%v)", *this.Tsize, *that1.Tsize)
		}
	} else if this.Tsize != nil {
		return fmt.Errorf("this.Tsize == nil && that.Tsize != nil")
	} else if that1.Tsize != nil {
		return fmt.Errorf("Tsize this(%v) Not Equal that(%v)", this.Tsize, that1.Tsize)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *PBLink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PBLink)
	if !ok {
		that2, ok := that.(PBLink)
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
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	if this.Name != nil && that1.Name != nil {
		if *this.Name != *that1.Name {
			return false
		}
	} else if this.Name != nil {
		return false
	} else if that1.Name != nil {
		return false
	}
	if this.Tsize != nil && that1.Tsize != nil {
		if *this.Tsize != *that1.Tsize {
			return false
		}
	} else if this.Tsize != nil {
		return false
	} else if that1.Tsize != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PBNode) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*PBNode)
	if !ok {
		that2, ok := that.(PBNode)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *PBNode")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *PBNode but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *PBNode but is not nil && this == nil")
	}
	if len(this.Links) != len(that1.Links) {
		return fmt.Errorf("Links this(%v) Not Equal that(%v)", len(this.Links), len(that1.Links))
	}
	for i := range this.Links {
		if !this.Links[i].Equal(that1.Links[i]) {
			return fmt.Errorf("Links this[%v](%v) Not Equal that[%v](%v)", i, this.Links[i], i, that1.Links[i])
		}
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return fmt.Errorf("Data this(%v) Not Equal that(%v)", this.Data, that1.Data)
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *PBNode) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PBNode)
	if !ok {
		that2, ok := that.(PBNode)
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
	if len(this.Links) != len(that1.Links) {
		return false
	}
	for i := range this.Links {
		if !this.Links[i].Equal(that1.Links[i]) {
			return false
		}
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PBLink) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&merkledag_pb.PBLink{")
	if this.Hash != nil {
		s = append(s, "Hash: "+valueToGoStringMerkledag(this.Hash, "byte")+",\n")
	}
	if this.Name != nil {
		s = append(s, "Name: "+valueToGoStringMerkledag(this.Name, "string")+",\n")
	}
	if this.Tsize != nil {
		s = append(s, "Tsize: "+valueToGoStringMerkledag(this.Tsize, "uint64")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *PBNode) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&merkledag_pb.PBNode{")
	if this.Links != nil {
		s = append(s, "Links: "+fmt.Sprintf("%#v", this.Links)+",\n")
	}
	if this.Data != nil {
		s = append(s, "Data: "+valueToGoStringMerkledag(this.Data, "byte")+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMerkledag(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *PBLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PBLink) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Hash != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMerkledag(dAtA, i, uint64(len(m.Hash)))
		i += copy(dAtA[i:], m.Hash)
	}
	if m.Name != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintMerkledag(dAtA, i, uint64(len(*m.Name)))
		i += copy(dAtA[i:], *m.Name)
	}
	if m.Tsize != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMerkledag(dAtA, i, uint64(*m.Tsize))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *PBNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PBNode) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Links) > 0 {
		for _, msg := range m.Links {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMerkledag(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMerkledag(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMerkledag(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedPBLink(r randyMerkledag, easy bool) *PBLink {
	this := &PBLink{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(100)
		this.Hash = make([]byte, v1)
		for i := 0; i < v1; i++ {
			this.Hash[i] = byte(r.Intn(256))
		}
	}
	if r.Intn(10) != 0 {
		v2 := string(randStringMerkledag(r))
		this.Name = &v2
	}
	if r.Intn(10) != 0 {
		v3 := uint64(uint64(r.Uint32()))
		this.Tsize = &v3
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMerkledag(r, 4)
	}
	return this
}

func NewPopulatedPBNode(r randyMerkledag, easy bool) *PBNode {
	this := &PBNode{}
	if r.Intn(10) != 0 {
		v4 := r.Intn(100)
		this.Data = make([]byte, v4)
		for i := 0; i < v4; i++ {
			this.Data[i] = byte(r.Intn(256))
		}
	}
	if r.Intn(10) != 0 {
		v5 := r.Intn(5)
		this.Links = make([]*PBLink, v5)
		for i := 0; i < v5; i++ {
			this.Links[i] = NewPopulatedPBLink(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMerkledag(r, 3)
	}
	return this
}

type randyMerkledag interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMerkledag(r randyMerkledag) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMerkledag(r randyMerkledag) string {
	v6 := r.Intn(100)
	tmps := make([]rune, v6)
	for i := 0; i < v6; i++ {
		tmps[i] = randUTF8RuneMerkledag(r)
	}
	return string(tmps)
}
func randUnrecognizedMerkledag(r randyMerkledag, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMerkledag(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMerkledag(dAtA []byte, r randyMerkledag, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMerkledag(dAtA, uint64(key))
		v7 := r.Int63()
		if r.Intn(2) == 0 {
			v7 *= -1
		}
		dAtA = encodeVarintPopulateMerkledag(dAtA, uint64(v7))
	case 1:
		dAtA = encodeVarintPopulateMerkledag(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMerkledag(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMerkledag(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMerkledag(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMerkledag(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *PBLink) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Hash != nil {
		l = len(m.Hash)
		n += 1 + l + sovMerkledag(uint64(l))
	}
	if m.Name != nil {
		l = len(*m.Name)
		n += 1 + l + sovMerkledag(uint64(l))
	}
	if m.Tsize != nil {
		n += 1 + sovMerkledag(uint64(*m.Tsize))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *PBNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovMerkledag(uint64(l))
	}
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovMerkledag(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMerkledag(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMerkledag(x uint64) (n int) {
	return sovMerkledag(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PBLink) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PBLink{`,
		`Hash:` + valueToStringMerkledag(this.Hash) + `,`,
		`Name:` + valueToStringMerkledag(this.Name) + `,`,
		`Tsize:` + valueToStringMerkledag(this.Tsize) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PBNode) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PBNode{`,
		`Data:` + valueToStringMerkledag(this.Data) + `,`,
		`Links:` + strings.Replace(fmt.Sprintf("%v", this.Links), "PBLink", "PBLink", 1) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMerkledag(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PBLink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMerkledag
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
			return fmt.Errorf("proto: PBLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PBLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkledag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMerkledag
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMerkledag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkledag
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
				return ErrInvalidLengthMerkledag
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMerkledag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.Name = &s
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tsize", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkledag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tsize = &v
		default:
			iNdEx = preIndex
			skippy, err := skipMerkledag(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMerkledag
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMerkledag
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *PBNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMerkledag
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
			return fmt.Errorf("proto: PBNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PBNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkledag
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMerkledag
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMerkledag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMerkledag
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
				return ErrInvalidLengthMerkledag
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMerkledag
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, &PBLink{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMerkledag(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMerkledag
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMerkledag
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMerkledag(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMerkledag
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
					return 0, ErrIntOverflowMerkledag
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
					return 0, ErrIntOverflowMerkledag
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
				return 0, ErrInvalidLengthMerkledag
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMerkledag
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMerkledag
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
				next, err := skipMerkledag(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMerkledag
				}
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
	ErrInvalidLengthMerkledag = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMerkledag   = fmt.Errorf("proto: integer overflow")
)
