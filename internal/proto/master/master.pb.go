// Code generated by protoc-gen-go. DO NOT EDIT.
// source: master.proto

package master

import (
	context "context"
	fmt "fmt"
	message "github.com/zilliztech/milvus-distributed/internal/proto/message"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SegmentStatus int32

const (
	SegmentStatus_OPENED   SegmentStatus = 0
	SegmentStatus_CLOSED   SegmentStatus = 1
	SegmentStatus_INDEXING SegmentStatus = 2
	SegmentStatus_INDEXED  SegmentStatus = 3
)

var SegmentStatus_name = map[int32]string{
	0: "OPENED",
	1: "CLOSED",
	2: "INDEXING",
	3: "INDEXED",
}

var SegmentStatus_value = map[string]int32{
	"OPENED":   0,
	"CLOSED":   1,
	"INDEXING": 2,
	"INDEXED":  3,
}

func (x SegmentStatus) String() string {
	return proto.EnumName(SegmentStatus_name, int32(x))
}

func (SegmentStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{0}
}

type Collection struct {
	Id                   uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Schema               *message.Schema       `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	CreateTime           uint64                `protobuf:"varint,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	SegmentIds           []uint64              `protobuf:"varint,5,rep,packed,name=segment_ids,json=segmentIds,proto3" json:"segment_ids,omitempty"`
	PartitionTags        []string              `protobuf:"bytes,6,rep,name=partition_tags,json=partitionTags,proto3" json:"partition_tags,omitempty"`
	Indexes              []*message.IndexParam `protobuf:"bytes,7,rep,name=indexes,proto3" json:"indexes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{0}
}

func (m *Collection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collection.Unmarshal(m, b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return xxx_messageInfo_Collection.Size(m)
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Collection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Collection) GetSchema() *message.Schema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *Collection) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Collection) GetSegmentIds() []uint64 {
	if m != nil {
		return m.SegmentIds
	}
	return nil
}

func (m *Collection) GetPartitionTags() []string {
	if m != nil {
		return m.PartitionTags
	}
	return nil
}

func (m *Collection) GetIndexes() []*message.IndexParam {
	if m != nil {
		return m.Indexes
	}
	return nil
}

type Segment struct {
	SegmentId            uint64        `protobuf:"varint,1,opt,name=segment_id,json=segmentId,proto3" json:"segment_id,omitempty"`
	CollectionId         uint64        `protobuf:"varint,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	PartitionTag         string        `protobuf:"bytes,3,opt,name=partition_tag,json=partitionTag,proto3" json:"partition_tag,omitempty"`
	ChannelStart         int32         `protobuf:"varint,4,opt,name=channel_start,json=channelStart,proto3" json:"channel_start,omitempty"`
	ChannelEnd           int32         `protobuf:"varint,5,opt,name=channel_end,json=channelEnd,proto3" json:"channel_end,omitempty"`
	OpenTimestamp        uint64        `protobuf:"varint,6,opt,name=open_timestamp,json=openTimestamp,proto3" json:"open_timestamp,omitempty"`
	CloseTimestamp       uint64        `protobuf:"varint,7,opt,name=close_timestamp,json=closeTimestamp,proto3" json:"close_timestamp,omitempty"`
	CollectionName       string        `protobuf:"bytes,8,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	Status               SegmentStatus `protobuf:"varint,9,opt,name=status,proto3,enum=masterpb.SegmentStatus" json:"status,omitempty"`
	Rows                 int64         `protobuf:"varint,10,opt,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Segment) Reset()         { *m = Segment{} }
func (m *Segment) String() string { return proto.CompactTextString(m) }
func (*Segment) ProtoMessage()    {}
func (*Segment) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{1}
}

func (m *Segment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Segment.Unmarshal(m, b)
}
func (m *Segment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Segment.Marshal(b, m, deterministic)
}
func (m *Segment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Segment.Merge(m, src)
}
func (m *Segment) XXX_Size() int {
	return xxx_messageInfo_Segment.Size(m)
}
func (m *Segment) XXX_DiscardUnknown() {
	xxx_messageInfo_Segment.DiscardUnknown(m)
}

var xxx_messageInfo_Segment proto.InternalMessageInfo

func (m *Segment) GetSegmentId() uint64 {
	if m != nil {
		return m.SegmentId
	}
	return 0
}

func (m *Segment) GetCollectionId() uint64 {
	if m != nil {
		return m.CollectionId
	}
	return 0
}

func (m *Segment) GetPartitionTag() string {
	if m != nil {
		return m.PartitionTag
	}
	return ""
}

func (m *Segment) GetChannelStart() int32 {
	if m != nil {
		return m.ChannelStart
	}
	return 0
}

func (m *Segment) GetChannelEnd() int32 {
	if m != nil {
		return m.ChannelEnd
	}
	return 0
}

func (m *Segment) GetOpenTimestamp() uint64 {
	if m != nil {
		return m.OpenTimestamp
	}
	return 0
}

func (m *Segment) GetCloseTimestamp() uint64 {
	if m != nil {
		return m.CloseTimestamp
	}
	return 0
}

func (m *Segment) GetCollectionName() string {
	if m != nil {
		return m.CollectionName
	}
	return ""
}

func (m *Segment) GetStatus() SegmentStatus {
	if m != nil {
		return m.Status
	}
	return SegmentStatus_OPENED
}

func (m *Segment) GetRows() int64 {
	if m != nil {
		return m.Rows
	}
	return 0
}

type SegmentStat struct {
	SegmentId            uint64        `protobuf:"varint,1,opt,name=segment_id,json=segmentId,proto3" json:"segment_id,omitempty"`
	MemorySize           uint64        `protobuf:"varint,2,opt,name=memory_size,json=memorySize,proto3" json:"memory_size,omitempty"`
	MemoryRate           float32       `protobuf:"fixed32,3,opt,name=memory_rate,json=memoryRate,proto3" json:"memory_rate,omitempty"`
	Status               SegmentStatus `protobuf:"varint,4,opt,name=status,proto3,enum=masterpb.SegmentStatus" json:"status,omitempty"`
	Rows                 int64         `protobuf:"varint,5,opt,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SegmentStat) Reset()         { *m = SegmentStat{} }
func (m *SegmentStat) String() string { return proto.CompactTextString(m) }
func (*SegmentStat) ProtoMessage()    {}
func (*SegmentStat) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9c348dec43a6705, []int{2}
}

func (m *SegmentStat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SegmentStat.Unmarshal(m, b)
}
func (m *SegmentStat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SegmentStat.Marshal(b, m, deterministic)
}
func (m *SegmentStat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SegmentStat.Merge(m, src)
}
func (m *SegmentStat) XXX_Size() int {
	return xxx_messageInfo_SegmentStat.Size(m)
}
func (m *SegmentStat) XXX_DiscardUnknown() {
	xxx_messageInfo_SegmentStat.DiscardUnknown(m)
}

var xxx_messageInfo_SegmentStat proto.InternalMessageInfo

func (m *SegmentStat) GetSegmentId() uint64 {
	if m != nil {
		return m.SegmentId
	}
	return 0
}

func (m *SegmentStat) GetMemorySize() uint64 {
	if m != nil {
		return m.MemorySize
	}
	return 0
}

func (m *SegmentStat) GetMemoryRate() float32 {
	if m != nil {
		return m.MemoryRate
	}
	return 0
}

func (m *SegmentStat) GetStatus() SegmentStatus {
	if m != nil {
		return m.Status
	}
	return SegmentStatus_OPENED
}

func (m *SegmentStat) GetRows() int64 {
	if m != nil {
		return m.Rows
	}
	return 0
}

func init() {
	proto.RegisterEnum("masterpb.SegmentStatus", SegmentStatus_name, SegmentStatus_value)
	proto.RegisterType((*Collection)(nil), "masterpb.Collection")
	proto.RegisterType((*Segment)(nil), "masterpb.Segment")
	proto.RegisterType((*SegmentStat)(nil), "masterpb.SegmentStat")
}

func init() { proto.RegisterFile("master.proto", fileDescriptor_f9c348dec43a6705) }

var fileDescriptor_f9c348dec43a6705 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x86, 0x6b, 0x27, 0x71, 0x9a, 0xe3, 0x24, 0x0b, 0xda, 0x60, 0xa6, 0x30, 0x6a, 0x32, 0xc6,
	0xcc, 0x06, 0x4e, 0xd7, 0x5d, 0xec, 0x6a, 0x0c, 0xda, 0x84, 0x11, 0x58, 0xd3, 0x62, 0xf7, 0x62,
	0xec, 0x26, 0xa8, 0xb6, 0x70, 0x05, 0x96, 0x6c, 0x2c, 0xa5, 0xdb, 0xfa, 0x04, 0x7b, 0x9f, 0xbd,
	0xdb, 0x60, 0x77, 0x43, 0xc7, 0x6e, 0xe3, 0x0c, 0xca, 0xd8, 0x9d, 0xf4, 0x9f, 0x2f, 0xd2, 0x9f,
	0x5f, 0xe7, 0x18, 0x86, 0x82, 0x2a, 0xcd, 0xaa, 0xb0, 0xac, 0x0a, 0x5d, 0x90, 0xfd, 0x7a, 0x57,
	0x5e, 0x1d, 0x8c, 0x04, 0x53, 0x8a, 0x66, 0xac, 0x2e, 0x4c, 0x7f, 0x5b, 0x00, 0xa7, 0x45, 0x9e,
	0xb3, 0x44, 0xf3, 0x42, 0x92, 0x31, 0xd8, 0x3c, 0xf5, 0x2c, 0xdf, 0x0a, 0xba, 0x91, 0xcd, 0x53,
	0x42, 0xa0, 0x2b, 0xa9, 0x60, 0x9e, 0xed, 0x5b, 0xc1, 0x20, 0xc2, 0x35, 0x79, 0x0d, 0x8e, 0x4a,
	0xae, 0x99, 0xa0, 0x5e, 0xc7, 0xb7, 0x02, 0xf7, 0xf8, 0x71, 0x28, 0x78, 0x7e, 0xb3, 0x51, 0x61,
	0x56, 0x95, 0x49, 0x18, 0x63, 0x29, 0x6a, 0x10, 0x72, 0x08, 0x6e, 0x52, 0x31, 0xaa, 0xd9, 0x5a,
	0x73, 0xc1, 0xbc, 0x2e, 0x9e, 0x0c, 0xb5, 0x74, 0xc9, 0x05, 0x33, 0x80, 0x62, 0x99, 0x60, 0x52,
	0xaf, 0x79, 0xaa, 0xbc, 0x9e, 0xdf, 0x31, 0x40, 0x23, 0x2d, 0x53, 0x45, 0x5e, 0xc0, 0xb8, 0xa4,
	0x95, 0xe6, 0xc6, 0xdf, 0x5a, 0xd3, 0x4c, 0x79, 0x8e, 0xdf, 0x09, 0x06, 0xd1, 0xe8, 0x5e, 0xbd,
	0xa4, 0x99, 0x22, 0x6f, 0xa0, 0xcf, 0x65, 0xca, 0xbe, 0x31, 0xe5, 0xf5, 0xfd, 0x4e, 0xe0, 0x1e,
	0x3f, 0xdd, 0xb1, 0xb5, 0x34, 0xb5, 0x0b, 0x5a, 0x51, 0x11, 0xdd, 0x71, 0xd3, 0x5f, 0x36, 0xf4,
	0xe3, 0xfa, 0x22, 0xf2, 0x0c, 0x60, 0x6b, 0xa3, 0x09, 0x60, 0x70, 0xef, 0x82, 0x3c, 0x87, 0x51,
	0x72, 0x9f, 0x92, 0x21, 0x6c, 0x24, 0x86, 0x5b, 0xb1, 0x86, 0x76, 0x9c, 0x62, 0x3e, 0x83, 0x68,
	0xd8, 0x36, 0x8a, 0x27, 0x5d, 0x53, 0x29, 0x59, 0xbe, 0x56, 0x9a, 0x56, 0x1a, 0x23, 0xe9, 0x45,
	0xc3, 0x46, 0x8c, 0x8d, 0x86, 0xa9, 0x35, 0x10, 0x93, 0xa9, 0xd7, 0x43, 0x04, 0x1a, 0x69, 0x21,
	0x53, 0x13, 0x4a, 0x51, 0x32, 0x89, 0xa1, 0x2a, 0x4d, 0x45, 0xe9, 0x39, 0x68, 0x68, 0x64, 0xd4,
	0xcb, 0x3b, 0x91, 0xbc, 0x84, 0x47, 0x49, 0x5e, 0x28, 0xd6, 0xe2, 0xfa, 0xc8, 0x8d, 0x51, 0xde,
	0x05, 0xb7, 0xff, 0x0f, 0x9f, 0x7c, 0x1f, 0xcd, 0x8f, 0xb7, 0xf2, 0xca, 0x3c, 0xfe, 0x0c, 0x1c,
	0xa5, 0xa9, 0xde, 0x28, 0x6f, 0xe0, 0x5b, 0xc1, 0xd8, 0xa4, 0xdc, 0x74, 0x56, 0xd8, 0x44, 0x19,
	0x63, 0x39, 0x6a, 0x30, 0xd3, 0x41, 0x55, 0xf1, 0x55, 0x79, 0xe0, 0x5b, 0x41, 0x27, 0xc2, 0xf5,
	0xf4, 0xa7, 0x05, 0x6e, 0x8b, 0xfe, 0x57, 0xf8, 0x87, 0xe0, 0x0a, 0x26, 0x8a, 0xea, 0xfb, 0x5a,
	0xf1, 0x5b, 0xd6, 0x44, 0x0f, 0xb5, 0x14, 0xf3, 0x5b, 0xd6, 0x02, 0x2a, 0xaa, 0x19, 0xc6, 0x6e,
	0xdf, 0x01, 0x11, 0xd5, 0x6d, 0xd7, 0xdd, 0xff, 0x73, 0xdd, 0xdb, 0xba, 0x7e, 0x75, 0x02, 0xa3,
	0x1d, 0x98, 0x00, 0x38, 0xe7, 0x17, 0x8b, 0xd5, 0x62, 0x3e, 0xd9, 0x33, 0xeb, 0xd3, 0x4f, 0xe7,
	0xf1, 0x62, 0x3e, 0xb1, 0xc8, 0x10, 0xf6, 0x97, 0xab, 0xf9, 0xe2, 0xf3, 0x72, 0xf5, 0x71, 0x62,
	0x13, 0x17, 0xfa, 0xb8, 0x5b, 0xcc, 0x27, 0x9d, 0xe3, 0x1f, 0x16, 0x38, 0x67, 0x78, 0x35, 0xf9,
	0x00, 0x93, 0x53, 0x1c, 0x83, 0xd6, 0xf8, 0x3d, 0xd9, 0xe9, 0xd9, 0x33, 0x5a, 0x96, 0x5c, 0x66,
	0x07, 0x7f, 0x0d, 0x18, 0x5e, 0x3e, 0xdd, 0x23, 0xef, 0xc1, 0xad, 0x0f, 0xc0, 0xde, 0x26, 0x0f,
	0xf5, 0xfb, 0x03, 0x3f, 0x3f, 0x39, 0xfa, 0x12, 0x66, 0x5c, 0x5f, 0x6f, 0xae, 0xc2, 0xa4, 0x10,
	0xb3, 0xe4, 0x56, 0x1d, 0x1d, 0xbd, 0x9b, 0xa9, 0xcd, 0x4d, 0xce, 0xc5, 0x8c, 0x4b, 0xcd, 0x2a,
	0x49, 0xf3, 0x19, 0x7e, 0x24, 0x66, 0x75, 0x58, 0x57, 0x0e, 0xee, 0xde, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x6c, 0xc2, 0x7d, 0xd6, 0x5b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MasterClient interface {
	CreateCollection(ctx context.Context, in *message.Mapping, opts ...grpc.CallOption) (*message.Status, error)
	CreateIndex(ctx context.Context, in *message.IndexParam, opts ...grpc.CallOption) (*message.Status, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) CreateCollection(ctx context.Context, in *message.Mapping, opts ...grpc.CallOption) (*message.Status, error) {
	out := new(message.Status)
	err := c.cc.Invoke(ctx, "/masterpb.Master/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) CreateIndex(ctx context.Context, in *message.IndexParam, opts ...grpc.CallOption) (*message.Status, error) {
	out := new(message.Status)
	err := c.cc.Invoke(ctx, "/masterpb.Master/CreateIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
type MasterServer interface {
	CreateCollection(context.Context, *message.Mapping) (*message.Status, error)
	CreateIndex(context.Context, *message.IndexParam) (*message.Status, error)
}

// UnimplementedMasterServer can be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (*UnimplementedMasterServer) CreateCollection(ctx context.Context, req *message.Mapping) (*message.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (*UnimplementedMasterServer) CreateIndex(ctx context.Context, req *message.IndexParam) (*message.Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIndex not implemented")
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.Mapping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/masterpb.Master/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).CreateCollection(ctx, req.(*message.Mapping))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_CreateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(message.IndexParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).CreateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/masterpb.Master/CreateIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).CreateIndex(ctx, req.(*message.IndexParam))
	}
	return interceptor(ctx, in, info, handler)
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "masterpb.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollection",
			Handler:    _Master_CreateCollection_Handler,
		},
		{
			MethodName: "CreateIndex",
			Handler:    _Master_CreateIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}
