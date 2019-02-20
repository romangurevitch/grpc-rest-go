// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type Result_Status int32

const (
	Result_Success Result_Status = 0
	Result_Fail    Result_Status = 1
)

var Result_Status_name = map[int32]string{
	0: "Success",
	1: "Fail",
}

var Result_Status_value = map[string]int32{
	"Success": 0,
	"Fail":    1,
}

func (x Result_Status) String() string {
	return proto.EnumName(Result_Status_name, int32(x))
}

func (Result_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4, 0}
}

type Word struct {
	Word                 string   `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Word) Reset()         { *m = Word{} }
func (m *Word) String() string { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()    {}
func (*Word) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *Word) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Word.Unmarshal(m, b)
}
func (m *Word) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Word.Marshal(b, m, deterministic)
}
func (m *Word) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Word.Merge(m, src)
}
func (m *Word) XXX_Size() int {
	return xxx_messageInfo_Word.Size(m)
}
func (m *Word) XXX_DiscardUnknown() {
	xxx_messageInfo_Word.DiscardUnknown(m)
}

var xxx_messageInfo_Word proto.InternalMessageInfo

func (m *Word) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

type FindWordRes struct {
	Exist                bool     `protobuf:"varint,1,opt,name=exist,proto3" json:"exist,omitempty"`
	Result               *Result  `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindWordRes) Reset()         { *m = FindWordRes{} }
func (m *FindWordRes) String() string { return proto.CompactTextString(m) }
func (*FindWordRes) ProtoMessage()    {}
func (*FindWordRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *FindWordRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWordRes.Unmarshal(m, b)
}
func (m *FindWordRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWordRes.Marshal(b, m, deterministic)
}
func (m *FindWordRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWordRes.Merge(m, src)
}
func (m *FindWordRes) XXX_Size() int {
	return xxx_messageInfo_FindWordRes.Size(m)
}
func (m *FindWordRes) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWordRes.DiscardUnknown(m)
}

var xxx_messageInfo_FindWordRes proto.InternalMessageInfo

func (m *FindWordRes) GetExist() bool {
	if m != nil {
		return m.Exist
	}
	return false
}

func (m *FindWordRes) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type GetMostPopularWordsRes struct {
	Words                map[string]int64 `protobuf:"bytes,1,rep,name=words,proto3" json:"words,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Result               *Result          `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetMostPopularWordsRes) Reset()         { *m = GetMostPopularWordsRes{} }
func (m *GetMostPopularWordsRes) String() string { return proto.CompactTextString(m) }
func (*GetMostPopularWordsRes) ProtoMessage()    {}
func (*GetMostPopularWordsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *GetMostPopularWordsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMostPopularWordsRes.Unmarshal(m, b)
}
func (m *GetMostPopularWordsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMostPopularWordsRes.Marshal(b, m, deterministic)
}
func (m *GetMostPopularWordsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMostPopularWordsRes.Merge(m, src)
}
func (m *GetMostPopularWordsRes) XXX_Size() int {
	return xxx_messageInfo_GetMostPopularWordsRes.Size(m)
}
func (m *GetMostPopularWordsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMostPopularWordsRes.DiscardUnknown(m)
}

var xxx_messageInfo_GetMostPopularWordsRes proto.InternalMessageInfo

func (m *GetMostPopularWordsRes) GetWords() map[string]int64 {
	if m != nil {
		return m.Words
	}
	return nil
}

func (m *GetMostPopularWordsRes) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Result struct {
	Status               Result_Status `protobuf:"varint,1,opt,name=status,proto3,enum=api.Result_Status" json:"status,omitempty"`
	Error                string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetStatus() Result_Status {
	if m != nil {
		return m.Status
	}
	return Result_Success
}

func (m *Result) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.Result_Status", Result_Status_name, Result_Status_value)
	proto.RegisterType((*Word)(nil), "api.Word")
	proto.RegisterType((*FindWordRes)(nil), "api.FindWordRes")
	proto.RegisterType((*GetMostPopularWordsRes)(nil), "api.GetMostPopularWordsRes")
	proto.RegisterMapType((map[string]int64)(nil), "api.GetMostPopularWordsRes.WordsEntry")
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*Result)(nil), "api.Result")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x49, 0xba, 0x75, 0xeb, 0x29, 0x42, 0xc1, 0x20, 0x34, 0x85, 0xc1, 0x2a, 0x23, 0xa6,
	0xaa, 0xa2, 0x09, 0x0b, 0x37, 0x53, 0xc5, 0x05, 0x43, 0xdb, 0xd8, 0x0d, 0x30, 0xa5, 0x17, 0x5c,
	0xbb, 0x89, 0x49, 0xcd, 0xd2, 0xd8, 0xb2, 0x9d, 0x6e, 0x15, 0xe2, 0x86, 0x17, 0x40, 0x62, 0xcf,
	0xc2, 0x93, 0xf0, 0x06, 0x88, 0x07, 0x41, 0x39, 0xe9, 0x46, 0x51, 0xab, 0x89, 0xab, 0x9c, 0x63,
	0xff, 0xe7, 0xfb, 0x63, 0x9f, 0x63, 0x68, 0x31, 0x25, 0x02, 0xa5, 0xa5, 0x95, 0xa4, 0xc1, 0x94,
	0xf0, 0xb7, 0x33, 0x29, 0xb3, 0x9c, 0x87, 0x4c, 0x89, 0x90, 0x15, 0x85, 0xb4, 0xcc, 0x0a, 0x59,
	0x98, 0x5a, 0xe2, 0x3f, 0xc3, 0x4f, 0xd2, 0xcf, 0x78, 0xd1, 0x37, 0xe7, 0x2c, 0xcb, 0xb8, 0x0e,
	0xa5, 0x42, 0xc5, 0xb2, 0x9a, 0xfa, 0xb0, 0xf6, 0x41, 0xea, 0x94, 0x10, 0x58, 0x3b, 0x97, 0x3a,
	0xdd, 0x72, 0x3a, 0x4e, 0xb7, 0x15, 0x63, 0x4c, 0x4f, 0xa0, 0x7d, 0x2c, 0x8a, 0xb4, 0xda, 0x8f,
	0xb9, 0x21, 0xf7, 0x61, 0x9d, 0x5f, 0x08, 0x63, 0x51, 0xb3, 0x19, 0xd7, 0x09, 0x79, 0x02, 0x4d,
	0xcd, 0x4d, 0x99, 0xdb, 0x2d, 0xb7, 0xe3, 0x74, 0xdb, 0x51, 0x3b, 0xa8, 0xfe, 0x36, 0xc6, 0xa5,
	0x78, 0xbe, 0x45, 0x7f, 0x38, 0xf0, 0xe0, 0x0d, 0xb7, 0x6f, 0xa5, 0xb1, 0xa7, 0x52, 0x95, 0x39,
	0xd3, 0x15, 0xd4, 0x54, 0xd4, 0x97, 0xb0, 0x5e, 0x99, 0x99, 0x2d, 0xa7, 0xd3, 0xe8, 0xb6, 0xa3,
	0x5d, 0x2c, 0x5f, 0xad, 0x0d, 0x30, 0x38, 0x2a, 0xac, 0x9e, 0xc5, 0x75, 0xd1, 0x7f, 0xb9, 0xfb,
	0xfb, 0x00, 0x7f, 0x2b, 0x89, 0x07, 0x8d, 0x33, 0x3e, 0x9b, 0x1f, 0xb4, 0x0a, 0xab, 0x83, 0x4d,
	0x59, 0x5e, 0x72, 0x64, 0x34, 0xe2, 0x3a, 0x19, 0xb8, 0xfb, 0x0e, 0xdd, 0x80, 0xf5, 0xa3, 0x89,
	0xb2, 0x33, 0x7a, 0x06, 0xcd, 0x1a, 0x4a, 0x7a, 0xd0, 0x34, 0x96, 0xd9, 0xd2, 0x20, 0xe1, 0x4e,
	0x44, 0x16, 0x1c, 0x83, 0x21, 0xee, 0xc4, 0x73, 0x05, 0xde, 0x98, 0xd6, 0x52, 0x23, 0xb8, 0x15,
	0xd7, 0x09, 0xdd, 0x81, 0x66, 0xad, 0x23, 0x6d, 0xd8, 0x18, 0x96, 0x49, 0xc2, 0x8d, 0xf1, 0x6e,
	0x91, 0x4d, 0x58, 0x3b, 0x66, 0x22, 0xf7, 0x9c, 0xe8, 0x97, 0x0b, 0x77, 0x0f, 0x45, 0x52, 0xb5,
	0x89, 0xe9, 0xd9, 0x90, 0xeb, 0xa9, 0x48, 0x38, 0x79, 0x0f, 0x9b, 0x57, 0xdd, 0x20, 0x2d, 0x34,
	0xad, 0x42, 0xdf, 0xc3, 0x70, 0xa1, 0x4f, 0x74, 0xf7, 0xeb, 0xcf, 0xdf, 0x97, 0x6e, 0x87, 0x3c,
	0xc6, 0x01, 0x99, 0xee, 0x85, 0xe9, 0x35, 0x2b, 0xfc, 0x28, 0x8a, 0x34, 0xfc, 0x5c, 0xdd, 0xdd,
	0x17, 0x72, 0x02, 0x1b, 0x07, 0xe9, 0x12, 0x6f, 0xf1, 0x06, 0xe9, 0x53, 0x44, 0xed, 0xf8, 0x8f,
	0x56, 0xa0, 0x58, 0x7a, 0x4d, 0x7a, 0x07, 0x70, 0xc8, 0x73, 0x6e, 0xf9, 0x8d, 0xb0, 0x2e, 0xc2,
	0x68, 0xaf, 0xb3, 0x02, 0x96, 0x62, 0xf9, 0x15, 0x6f, 0x04, 0xf7, 0x56, 0x4c, 0x00, 0x01, 0xa4,
	0x61, 0x43, 0xfc, 0x87, 0x37, 0xcc, 0x09, 0xa5, 0xe8, 0xb4, 0x4d, 0xfc, 0x15, 0x4e, 0xaa, 0xd6,
	0xbe, 0xbe, 0x74, 0xbe, 0x1f, 0x7c, 0x73, 0x08, 0x87, 0xdb, 0x99, 0x56, 0x49, 0x5f, 0x73, 0x63,
	0xfb, 0x99, 0xa4, 0xa7, 0xff, 0xe6, 0x24, 0x18, 0x5b, 0xab, 0xcc, 0x20, 0x0c, 0x33, 0x61, 0xc7,
	0xe5, 0x28, 0x48, 0xe4, 0x24, 0xd4, 0x72, 0xc2, 0x8a, 0xac, 0xd4, 0x7c, 0x2a, 0x6c, 0x32, 0x0e,
	0x17, 0xf5, 0xbe, 0x57, 0xc8, 0x82, 0xbf, 0xe2, 0x17, 0x6c, 0xa2, 0x72, 0x5e, 0xa9, 0xa3, 0xc6,
	0x5e, 0xf0, 0xbc, 0xe7, 0x3a, 0x6e, 0xe4, 0x31, 0xa5, 0x72, 0x91, 0xe0, 0xf3, 0x0b, 0x3f, 0x19,
	0x59, 0x0c, 0x96, 0x56, 0x46, 0x4d, 0x7c, 0x95, 0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfb,
	0x86, 0x5b, 0x12, 0xf3, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DictionaryServiceClient is the client API for DictionaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DictionaryServiceClient interface {
	// Find a word in the dictionary
	FindWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*FindWordRes, error)
	// Add a word to the dictionary
	AddWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Result, error)
	// Delete a word from the dictionary
	DeleteWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Result, error)
	// Get most popular words in the dictionary
	GetMostPopularWords(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetMostPopularWordsRes, error)
}

type dictionaryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDictionaryServiceClient(cc *grpc.ClientConn) DictionaryServiceClient {
	return &dictionaryServiceClient{cc}
}

func (c *dictionaryServiceClient) FindWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*FindWordRes, error) {
	out := new(FindWordRes)
	err := c.cc.Invoke(ctx, "/api.DictionaryService/FindWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryServiceClient) AddWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/api.DictionaryService/AddWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryServiceClient) DeleteWord(ctx context.Context, in *Word, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/api.DictionaryService/DeleteWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictionaryServiceClient) GetMostPopularWords(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetMostPopularWordsRes, error) {
	out := new(GetMostPopularWordsRes)
	err := c.cc.Invoke(ctx, "/api.DictionaryService/GetMostPopularWords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictionaryServiceServer is the server API for DictionaryService service.
type DictionaryServiceServer interface {
	// Find a word in the dictionary
	FindWord(context.Context, *Word) (*FindWordRes, error)
	// Add a word to the dictionary
	AddWord(context.Context, *Word) (*Result, error)
	// Delete a word from the dictionary
	DeleteWord(context.Context, *Word) (*Result, error)
	// Get most popular words in the dictionary
	GetMostPopularWords(context.Context, *Empty) (*GetMostPopularWordsRes, error)
}

func RegisterDictionaryServiceServer(s *grpc.Server, srv DictionaryServiceServer) {
	s.RegisterService(&_DictionaryService_serviceDesc, srv)
}

func _DictionaryService_FindWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServiceServer).FindWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DictionaryService/FindWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServiceServer).FindWord(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictionaryService_AddWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServiceServer).AddWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DictionaryService/AddWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServiceServer).AddWord(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictionaryService_DeleteWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Word)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServiceServer).DeleteWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DictionaryService/DeleteWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServiceServer).DeleteWord(ctx, req.(*Word))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictionaryService_GetMostPopularWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictionaryServiceServer).GetMostPopularWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DictionaryService/GetMostPopularWords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictionaryServiceServer).GetMostPopularWords(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DictionaryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DictionaryService",
	HandlerType: (*DictionaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindWord",
			Handler:    _DictionaryService_FindWord_Handler,
		},
		{
			MethodName: "AddWord",
			Handler:    _DictionaryService_AddWord_Handler,
		},
		{
			MethodName: "DeleteWord",
			Handler:    _DictionaryService_DeleteWord_Handler,
		},
		{
			MethodName: "GetMostPopularWords",
			Handler:    _DictionaryService_GetMostPopularWords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
