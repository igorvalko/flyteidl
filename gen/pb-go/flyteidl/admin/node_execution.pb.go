// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/node_execution.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A message used to fetch a single node execution entity.
type NodeExecutionGetRequest struct {
	// Uniquely identifies an individual node execution.
	Id                   *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *NodeExecutionGetRequest) Reset()         { *m = NodeExecutionGetRequest{} }
func (m *NodeExecutionGetRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionGetRequest) ProtoMessage()    {}
func (*NodeExecutionGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_execution_e35f342f70a69e87, []int{0}
}
func (m *NodeExecutionGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionGetRequest.Unmarshal(m, b)
}
func (m *NodeExecutionGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionGetRequest.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionGetRequest.Merge(dst, src)
}
func (m *NodeExecutionGetRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionGetRequest.Size(m)
}
func (m *NodeExecutionGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionGetRequest proto.InternalMessageInfo

func (m *NodeExecutionGetRequest) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Represents a request structure to retrieve a list of node execution entities.
type NodeExecutionListRequest struct {
	// Indicates the number of resources to be returned.
	Limit uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// In the case of multiple pages of results, the, server-provided token can be used to fetch the next page
	// in a query.
	// +optional
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// Indicates a list of filters passed as string.
	// More info on constructing filters : <Link>
	// +optional
	Filters string `protobuf:"bytes,3,opt,name=filters,proto3" json:"filters,omitempty"`
	// Sort ordering.
	// +optional
	SortBy               *Sort    `protobuf:"bytes,4,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionListRequest) Reset()         { *m = NodeExecutionListRequest{} }
func (m *NodeExecutionListRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionListRequest) ProtoMessage()    {}
func (*NodeExecutionListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_execution_e35f342f70a69e87, []int{1}
}
func (m *NodeExecutionListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionListRequest.Unmarshal(m, b)
}
func (m *NodeExecutionListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionListRequest.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionListRequest.Merge(dst, src)
}
func (m *NodeExecutionListRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionListRequest.Size(m)
}
func (m *NodeExecutionListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionListRequest proto.InternalMessageInfo

func (m *NodeExecutionListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *NodeExecutionListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *NodeExecutionListRequest) GetFilters() string {
	if m != nil {
		return m.Filters
	}
	return ""
}

func (m *NodeExecutionListRequest) GetSortBy() *Sort {
	if m != nil {
		return m.SortBy
	}
	return nil
}

// Encapsulates all details for a single node execution entity.
// A node represents a component in the overall workflow graph. A node launch a task, multiple tasks, an entire nested
// sub-workflow, or even a separate child-workflow execution.
// The same task can be called repeatedly in a single workflow but each node is unique.
type NodeExecution struct {
	// Uniquely identifies an individual node execution.
	Id *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Path to remote data store where input blob is stored.
	InputUri string `protobuf:"bytes,2,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Computed results associated with this node execution.
	Closure              *NodeExecutionClosure `protobuf:"bytes,3,opt,name=closure,proto3" json:"closure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NodeExecution) Reset()         { *m = NodeExecution{} }
func (m *NodeExecution) String() string { return proto.CompactTextString(m) }
func (*NodeExecution) ProtoMessage()    {}
func (*NodeExecution) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_execution_e35f342f70a69e87, []int{2}
}
func (m *NodeExecution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecution.Unmarshal(m, b)
}
func (m *NodeExecution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecution.Marshal(b, m, deterministic)
}
func (dst *NodeExecution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecution.Merge(dst, src)
}
func (m *NodeExecution) XXX_Size() int {
	return xxx_messageInfo_NodeExecution.Size(m)
}
func (m *NodeExecution) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecution.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecution proto.InternalMessageInfo

func (m *NodeExecution) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *NodeExecution) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

func (m *NodeExecution) GetClosure() *NodeExecutionClosure {
	if m != nil {
		return m.Closure
	}
	return nil
}

// Request structure to retrieve a list of node execution entities.
type NodeExecutionList struct {
	NodeExecutions []*NodeExecution `protobuf:"bytes,1,rep,name=node_executions,json=nodeExecutions,proto3" json:"node_executions,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionList) Reset()         { *m = NodeExecutionList{} }
func (m *NodeExecutionList) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionList) ProtoMessage()    {}
func (*NodeExecutionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_execution_e35f342f70a69e87, []int{3}
}
func (m *NodeExecutionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionList.Unmarshal(m, b)
}
func (m *NodeExecutionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionList.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionList.Merge(dst, src)
}
func (m *NodeExecutionList) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionList.Size(m)
}
func (m *NodeExecutionList) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionList.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionList proto.InternalMessageInfo

func (m *NodeExecutionList) GetNodeExecutions() []*NodeExecution {
	if m != nil {
		return m.NodeExecutions
	}
	return nil
}

func (m *NodeExecutionList) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Container for node execution details and results.
type NodeExecutionClosure struct {
	// Only a node in a terminal state will have a non-empty output_result.
	//
	// Types that are valid to be assigned to OutputResult:
	//	*NodeExecutionClosure_OutputUri
	//	*NodeExecutionClosure_Error
	OutputResult isNodeExecutionClosure_OutputResult `protobuf_oneof:"output_result"`
	// The last recorded phase for this node execution.
	Phase core.NodeExecutionPhase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.NodeExecutionPhase" json:"phase,omitempty"`
	// Time at which the node execution began running.
	StartedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// The amount of time the node execution spent running.
	Duration *duration.Duration `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	// Time at which the node execution was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Time at which the node execution was last updated.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *NodeExecutionClosure) Reset()         { *m = NodeExecutionClosure{} }
func (m *NodeExecutionClosure) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionClosure) ProtoMessage()    {}
func (*NodeExecutionClosure) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_execution_e35f342f70a69e87, []int{4}
}
func (m *NodeExecutionClosure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionClosure.Unmarshal(m, b)
}
func (m *NodeExecutionClosure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionClosure.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionClosure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionClosure.Merge(dst, src)
}
func (m *NodeExecutionClosure) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionClosure.Size(m)
}
func (m *NodeExecutionClosure) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionClosure.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionClosure proto.InternalMessageInfo

type isNodeExecutionClosure_OutputResult interface {
	isNodeExecutionClosure_OutputResult()
}

type NodeExecutionClosure_OutputUri struct {
	OutputUri string `protobuf:"bytes,1,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type NodeExecutionClosure_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*NodeExecutionClosure_OutputUri) isNodeExecutionClosure_OutputResult() {}

func (*NodeExecutionClosure_Error) isNodeExecutionClosure_OutputResult() {}

func (m *NodeExecutionClosure) GetOutputResult() isNodeExecutionClosure_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *NodeExecutionClosure) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*NodeExecutionClosure_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *NodeExecutionClosure) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*NodeExecutionClosure_Error); ok {
		return x.Error
	}
	return nil
}

func (m *NodeExecutionClosure) GetPhase() core.NodeExecutionPhase {
	if m != nil {
		return m.Phase
	}
	return core.NodeExecutionPhase_NODE_PHASE_UNDEFINED
}

func (m *NodeExecutionClosure) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *NodeExecutionClosure) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *NodeExecutionClosure) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *NodeExecutionClosure) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NodeExecutionClosure) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NodeExecutionClosure_OneofMarshaler, _NodeExecutionClosure_OneofUnmarshaler, _NodeExecutionClosure_OneofSizer, []interface{}{
		(*NodeExecutionClosure_OutputUri)(nil),
		(*NodeExecutionClosure_Error)(nil),
	}
}

func _NodeExecutionClosure_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NodeExecutionClosure)
	// output_result
	switch x := m.OutputResult.(type) {
	case *NodeExecutionClosure_OutputUri:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OutputUri)
	case *NodeExecutionClosure_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NodeExecutionClosure.OutputResult has unexpected type %T", x)
	}
	return nil
}

func _NodeExecutionClosure_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NodeExecutionClosure)
	switch tag {
	case 1: // output_result.output_uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OutputResult = &NodeExecutionClosure_OutputUri{x}
		return true, err
	case 2: // output_result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ExecutionError)
		err := b.DecodeMessage(msg)
		m.OutputResult = &NodeExecutionClosure_Error{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NodeExecutionClosure_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NodeExecutionClosure)
	// output_result
	switch x := m.OutputResult.(type) {
	case *NodeExecutionClosure_OutputUri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.OutputUri)))
		n += len(x.OutputUri)
	case *NodeExecutionClosure_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Request structure to fetch inputs and output urls for a node execution.
type NodeExecutionGetDataRequest struct {
	// The identifier of the node execution for which to fetch inputs and outputs.
	Id                   *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *NodeExecutionGetDataRequest) Reset()         { *m = NodeExecutionGetDataRequest{} }
func (m *NodeExecutionGetDataRequest) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionGetDataRequest) ProtoMessage()    {}
func (*NodeExecutionGetDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_execution_e35f342f70a69e87, []int{5}
}
func (m *NodeExecutionGetDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionGetDataRequest.Unmarshal(m, b)
}
func (m *NodeExecutionGetDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionGetDataRequest.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionGetDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionGetDataRequest.Merge(dst, src)
}
func (m *NodeExecutionGetDataRequest) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionGetDataRequest.Size(m)
}
func (m *NodeExecutionGetDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionGetDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionGetDataRequest proto.InternalMessageInfo

func (m *NodeExecutionGetDataRequest) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// Response structure for NodeExecutionGetDataRequest which contains inputs and outputs for a node execution.
type NodeExecutionGetDataResponse struct {
	// Signed url to fetch a core.LiteralMap of node execution inputs.
	Inputs *UrlBlob `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Signed url to fetch a core.LiteralMap of node execution outputs.
	Outputs              *UrlBlob `protobuf:"bytes,2,opt,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionGetDataResponse) Reset()         { *m = NodeExecutionGetDataResponse{} }
func (m *NodeExecutionGetDataResponse) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionGetDataResponse) ProtoMessage()    {}
func (*NodeExecutionGetDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_node_execution_e35f342f70a69e87, []int{6}
}
func (m *NodeExecutionGetDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionGetDataResponse.Unmarshal(m, b)
}
func (m *NodeExecutionGetDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionGetDataResponse.Marshal(b, m, deterministic)
}
func (dst *NodeExecutionGetDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionGetDataResponse.Merge(dst, src)
}
func (m *NodeExecutionGetDataResponse) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionGetDataResponse.Size(m)
}
func (m *NodeExecutionGetDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionGetDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionGetDataResponse proto.InternalMessageInfo

func (m *NodeExecutionGetDataResponse) GetInputs() *UrlBlob {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *NodeExecutionGetDataResponse) GetOutputs() *UrlBlob {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeExecutionGetRequest)(nil), "flyteidl.admin.NodeExecutionGetRequest")
	proto.RegisterType((*NodeExecutionListRequest)(nil), "flyteidl.admin.NodeExecutionListRequest")
	proto.RegisterType((*NodeExecution)(nil), "flyteidl.admin.NodeExecution")
	proto.RegisterType((*NodeExecutionList)(nil), "flyteidl.admin.NodeExecutionList")
	proto.RegisterType((*NodeExecutionClosure)(nil), "flyteidl.admin.NodeExecutionClosure")
	proto.RegisterType((*NodeExecutionGetDataRequest)(nil), "flyteidl.admin.NodeExecutionGetDataRequest")
	proto.RegisterType((*NodeExecutionGetDataResponse)(nil), "flyteidl.admin.NodeExecutionGetDataResponse")
}

func init() {
	proto.RegisterFile("flyteidl/admin/node_execution.proto", fileDescriptor_node_execution_e35f342f70a69e87)
}

var fileDescriptor_node_execution_e35f342f70a69e87 = []byte{
	// 589 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x4f, 0xd4, 0x40,
	0x14, 0xa5, 0xe0, 0xee, 0xb2, 0x97, 0x00, 0x71, 0x42, 0x42, 0x5d, 0x44, 0x70, 0x35, 0x86, 0x17,
	0xda, 0xb0, 0x04, 0x8d, 0x2f, 0x26, 0xac, 0xa0, 0x98, 0x18, 0xa3, 0xa3, 0xbc, 0xf8, 0xb2, 0xe9,
	0xc7, 0xec, 0x32, 0xb1, 0xed, 0x94, 0x99, 0xdb, 0xc4, 0x7d, 0xf5, 0x0f, 0xf8, 0x17, 0x7c, 0xf0,
	0x87, 0x9a, 0x4e, 0x67, 0xaa, 0x2d, 0x08, 0x89, 0x3e, 0xde, 0x99, 0x73, 0xce, 0x3d, 0xf7, 0xf4,
	0x76, 0xe0, 0xd1, 0x34, 0x99, 0x23, 0xe3, 0x71, 0xe2, 0x07, 0x71, 0xca, 0x33, 0x3f, 0x13, 0x31,
	0x9b, 0xb0, 0xaf, 0x2c, 0x2a, 0x90, 0x8b, 0xcc, 0xcb, 0xa5, 0x40, 0x41, 0xd6, 0x2c, 0xc8, 0xd3,
	0xa0, 0xc1, 0x56, 0x8b, 0x14, 0x89, 0x34, 0xb5, 0xe0, 0xc1, 0x76, 0x7d, 0x19, 0x09, 0xc9, 0xfc,
	0x96, 0xd6, 0xe0, 0x41, 0xf3, 0x9a, 0xc7, 0x2c, 0x43, 0x3e, 0xe5, 0x4c, 0x9a, 0xfb, 0x9d, 0x99,
	0x10, 0xb3, 0x84, 0xf9, 0xba, 0x0a, 0x8b, 0xa9, 0x8f, 0x3c, 0x65, 0x0a, 0x83, 0x34, 0xb7, 0x02,
	0x6d, 0x40, 0x5c, 0xc8, 0xe0, 0x77, 0x83, 0xe1, 0x07, 0xd8, 0x7c, 0x27, 0x62, 0x76, 0x6a, 0xfb,
	0xbe, 0x66, 0x48, 0xd9, 0x65, 0xc1, 0x14, 0x92, 0xa7, 0xb0, 0xc8, 0x63, 0xd7, 0xd9, 0x75, 0xf6,
	0x56, 0x46, 0x4f, 0xbc, 0x7a, 0xa8, 0xd2, 0x88, 0xd7, 0xe0, 0xbc, 0xa9, 0x5d, 0xd1, 0x45, 0x1e,
	0x0f, 0xbf, 0x3b, 0xe0, 0x36, 0xee, 0xdf, 0x72, 0x55, 0x8b, 0x6e, 0x40, 0x27, 0xe1, 0x29, 0x47,
	0xad, 0xbb, 0x4a, 0xab, 0xa2, 0x3c, 0x45, 0xf1, 0x85, 0x65, 0xee, 0xe2, 0xae, 0xb3, 0xd7, 0xa7,
	0x55, 0x41, 0x5c, 0xe8, 0x4d, 0x79, 0x82, 0x4c, 0x2a, 0x77, 0x49, 0x9f, 0xdb, 0x92, 0xec, 0x43,
	0x4f, 0x09, 0x89, 0x93, 0x70, 0xee, 0xde, 0xd1, 0xfe, 0x36, 0xbc, 0x66, 0xe8, 0xde, 0x47, 0x21,
	0x91, 0x76, 0x4b, 0xd0, 0x78, 0x3e, 0xfc, 0xe9, 0xc0, 0x6a, 0xc3, 0xd1, 0xbf, 0xce, 0x46, 0xb6,
	0xa0, 0xcf, 0xb3, 0xbc, 0xc0, 0x49, 0x21, 0xb9, 0x31, 0xbb, 0xac, 0x0f, 0xce, 0x25, 0x27, 0x2f,
	0xa0, 0x17, 0x25, 0x42, 0x15, 0x92, 0x69, 0xbf, 0x2b, 0xa3, 0xc7, 0x6d, 0x57, 0x0d, 0xe9, 0x97,
	0x15, 0x96, 0x5a, 0xd2, 0xf0, 0x12, 0xee, 0x5e, 0xc9, 0x8d, 0xbc, 0x82, 0xf5, 0xe6, 0x96, 0x29,
	0xd7, 0xd9, 0x5d, 0xda, 0x5b, 0x19, 0x6d, 0xdf, 0x28, 0x4e, 0xd7, 0xb2, 0x3f, 0x4b, 0x75, 0x7d,
	0xc4, 0xc3, 0x1f, 0x4b, 0xb0, 0x71, 0x9d, 0x29, 0xb2, 0x03, 0x20, 0x0a, 0xb4, 0x93, 0x96, 0x41,
	0xf5, 0xcf, 0x16, 0x68, 0xbf, 0x3a, 0x2b, 0x87, 0x3d, 0x82, 0x0e, 0x93, 0x52, 0x48, 0xad, 0xd7,
	0x70, 0xa3, 0x43, 0xac, 0x05, 0x4f, 0x4b, 0xd0, 0xd9, 0x02, 0xad, 0xd0, 0xe4, 0x19, 0x74, 0xf2,
	0x8b, 0x40, 0x55, 0x09, 0xad, 0x8d, 0x1e, 0xde, 0x94, 0xfd, 0xfb, 0x12, 0x48, 0x2b, 0x3c, 0x79,
	0x0e, 0xa0, 0x30, 0x90, 0xc8, 0xe2, 0x49, 0x80, 0xe6, 0xab, 0x0f, 0xbc, 0x6a, 0xbb, 0x3d, 0xbb,
	0xdd, 0xde, 0x27, 0xbb, 0xfe, 0xb4, 0x6f, 0xd0, 0xc7, 0x48, 0x8e, 0x60, 0xd9, 0x6e, 0xbd, 0xdb,
	0xd1, 0xc4, 0x7b, 0x57, 0x88, 0x27, 0x06, 0x40, 0x6b, 0x68, 0xd9, 0x31, 0x92, 0x2c, 0x30, 0x1d,
	0xbb, 0xb7, 0x77, 0x34, 0xe8, 0x63, 0x2c, 0xa9, 0x45, 0x1e, 0x5b, 0x6a, 0xef, 0x76, 0xaa, 0x41,
	0x1f, 0xe3, 0x78, 0x1d, 0x56, 0x4d, 0xf0, 0x92, 0xa9, 0x22, 0xc1, 0xe1, 0x39, 0x6c, 0xb5, 0xff,
	0xd0, 0x93, 0x00, 0x83, 0xff, 0xfd, 0x4b, 0xbf, 0x39, 0x70, 0xff, 0x7a, 0x5d, 0x95, 0x8b, 0x4c,
	0x31, 0xe2, 0x43, 0x57, 0x6f, 0xb6, 0x32, 0xe2, 0x9b, 0xed, 0x7d, 0x3b, 0x97, 0xc9, 0x38, 0x11,
	0x21, 0x35, 0x30, 0x72, 0x00, 0xbd, 0xca, 0xb9, 0x32, 0x3b, 0xf1, 0x57, 0x86, 0xc5, 0x8d, 0x0f,
	0x3f, 0x1f, 0xcc, 0x38, 0x5e, 0x14, 0xa1, 0x17, 0x89, 0xd4, 0x4f, 0xe6, 0x53, 0xf4, 0xeb, 0x07,
	0x6f, 0xc6, 0x32, 0x3f, 0x0f, 0xf7, 0x67, 0xc2, 0x6f, 0xbe, 0x9f, 0x61, 0x57, 0x07, 0x78, 0xf8,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x3f, 0xbd, 0x55, 0x8d, 0x05, 0x00, 0x00,
}
