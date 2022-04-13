// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gogo_permessage_example.proto

package gogo

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EventType int32

const (
	EventType_EVENT_TYPE_UNDEFINED EventType = 0
	EventType_EVENT_TYPE_ONE       EventType = 1
	EventType_EVENT_TYPE_TWO       EventType = 2
)

var EventType_name = map[int32]string{
	0: "EVENT_TYPE_UNDEFINED",
	1: "EVENT_TYPE_ONE",
	2: "EVENT_TYPE_TWO",
}

var EventType_value = map[string]int32{
	"EVENT_TYPE_UNDEFINED": 0,
	"EVENT_TYPE_ONE":       1,
	"EVENT_TYPE_TWO":       2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_77954bdb6fa16f29, []int{0}
}

type TestEvent struct {
	Name      string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Info      string         `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	IsAwesome bool           `protobuf:"varint,3,opt,name=isAwesome,proto3" json:"isAwesome,omitempty"`
	Labels    []string       `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty"`
	Embedded  *EmbeddedEvent `protobuf:"bytes,5,opt,name=embedded,proto3" json:"embedded,omitempty"`
	// Types that are valid to be assigned to Path:
	//	*TestEvent_Jedi
	//	*TestEvent_Sith
	//	*TestEvent_Other
	Path                 isTestEvent_Path     `protobuf_oneof:"path"`
	Nested               *TestEvent_NestedMsg `protobuf:"bytes,9,opt,name=nested,proto3" json:"nested,omitempty"`
	Ts                   *types.Timestamp     `protobuf:"bytes,10,opt,name=ts,proto3" json:"ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TestEvent) Reset()         { *m = TestEvent{} }
func (m *TestEvent) String() string { return proto.CompactTextString(m) }
func (*TestEvent) ProtoMessage()    {}
func (*TestEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_77954bdb6fa16f29, []int{0}
}
func (m *TestEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestEvent.Unmarshal(m, b)
}
func (m *TestEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestEvent.Marshal(b, m, deterministic)
}
func (m *TestEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestEvent.Merge(m, src)
}
func (m *TestEvent) XXX_Size() int {
	return xxx_messageInfo_TestEvent.Size(m)
}
func (m *TestEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TestEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TestEvent proto.InternalMessageInfo

type isTestEvent_Path interface {
	isTestEvent_Path()
}

type TestEvent_Jedi struct {
	Jedi bool `protobuf:"varint,6,opt,name=jedi,proto3,oneof" json:"jedi,omitempty"`
}
type TestEvent_Sith struct {
	Sith bool `protobuf:"varint,7,opt,name=sith,proto3,oneof" json:"sith,omitempty"`
}
type TestEvent_Other struct {
	Other string `protobuf:"bytes,8,opt,name=other,proto3,oneof" json:"other,omitempty"`
}

func (*TestEvent_Jedi) isTestEvent_Path()  {}
func (*TestEvent_Sith) isTestEvent_Path()  {}
func (*TestEvent_Other) isTestEvent_Path() {}

func (m *TestEvent) GetPath() isTestEvent_Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *TestEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestEvent) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *TestEvent) GetIsAwesome() bool {
	if m != nil {
		return m.IsAwesome
	}
	return false
}

func (m *TestEvent) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TestEvent) GetEmbedded() *EmbeddedEvent {
	if m != nil {
		return m.Embedded
	}
	return nil
}

func (m *TestEvent) GetJedi() bool {
	if x, ok := m.GetPath().(*TestEvent_Jedi); ok {
		return x.Jedi
	}
	return false
}

func (m *TestEvent) GetSith() bool {
	if x, ok := m.GetPath().(*TestEvent_Sith); ok {
		return x.Sith
	}
	return false
}

func (m *TestEvent) GetOther() string {
	if x, ok := m.GetPath().(*TestEvent_Other); ok {
		return x.Other
	}
	return ""
}

func (m *TestEvent) GetNested() *TestEvent_NestedMsg {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *TestEvent) GetTs() *types.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestEvent_Jedi)(nil),
		(*TestEvent_Sith)(nil),
		(*TestEvent_Other)(nil),
	}
}

type TestEvent_NestedMsg struct {
	Details              string   `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestEvent_NestedMsg) Reset()         { *m = TestEvent_NestedMsg{} }
func (m *TestEvent_NestedMsg) String() string { return proto.CompactTextString(m) }
func (*TestEvent_NestedMsg) ProtoMessage()    {}
func (*TestEvent_NestedMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_77954bdb6fa16f29, []int{0, 0}
}
func (m *TestEvent_NestedMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestEvent_NestedMsg.Unmarshal(m, b)
}
func (m *TestEvent_NestedMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestEvent_NestedMsg.Marshal(b, m, deterministic)
}
func (m *TestEvent_NestedMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestEvent_NestedMsg.Merge(m, src)
}
func (m *TestEvent_NestedMsg) XXX_Size() int {
	return xxx_messageInfo_TestEvent_NestedMsg.Size(m)
}
func (m *TestEvent_NestedMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TestEvent_NestedMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TestEvent_NestedMsg proto.InternalMessageInfo

func (m *TestEvent_NestedMsg) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type EmbeddedEvent struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Stuff                string   `protobuf:"bytes,2,opt,name=stuff,proto3" json:"stuff,omitempty"`
	FavoriteNumbers      []int32  `protobuf:"varint,3,rep,packed,name=favoriteNumbers,proto3" json:"favoriteNumbers,omitempty"`
	RandomThings         [][]byte `protobuf:"bytes,4,rep,name=randomThings,proto3" json:"randomThings,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmbeddedEvent) Reset()         { *m = EmbeddedEvent{} }
func (m *EmbeddedEvent) String() string { return proto.CompactTextString(m) }
func (*EmbeddedEvent) ProtoMessage()    {}
func (*EmbeddedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_77954bdb6fa16f29, []int{1}
}
func (m *EmbeddedEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmbeddedEvent.Unmarshal(m, b)
}
func (m *EmbeddedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmbeddedEvent.Marshal(b, m, deterministic)
}
func (m *EmbeddedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmbeddedEvent.Merge(m, src)
}
func (m *EmbeddedEvent) XXX_Size() int {
	return xxx_messageInfo_EmbeddedEvent.Size(m)
}
func (m *EmbeddedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EmbeddedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EmbeddedEvent proto.InternalMessageInfo

func (m *EmbeddedEvent) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *EmbeddedEvent) GetStuff() string {
	if m != nil {
		return m.Stuff
	}
	return ""
}

func (m *EmbeddedEvent) GetFavoriteNumbers() []int32 {
	if m != nil {
		return m.FavoriteNumbers
	}
	return nil
}

func (m *EmbeddedEvent) GetRandomThings() [][]byte {
	if m != nil {
		return m.RandomThings
	}
	return nil
}

type AllTheThings struct {
	ID                   int32          `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TheString            string         `protobuf:"bytes,2,opt,name=theString,proto3" json:"theString,omitempty"`
	TheBool              bool           `protobuf:"varint,3,opt,name=theBool,proto3" json:"theBool,omitempty"`
	TheInt32             int32          `protobuf:"varint,4,opt,name=theInt32,proto3" json:"theInt32,omitempty"`
	TheInt64             int64          `protobuf:"varint,5,opt,name=theInt64,proto3" json:"theInt64,omitempty"`
	TheUInt32            uint32         `protobuf:"varint,6,opt,name=theUInt32,proto3" json:"theUInt32,omitempty"`
	TheUInt64            uint64         `protobuf:"varint,7,opt,name=theUInt64,proto3" json:"theUInt64,omitempty"`
	TheSInt32            int32          `protobuf:"zigzag32,8,opt,name=theSInt32,proto3" json:"theSInt32,omitempty"`
	TheSInt64            int64          `protobuf:"zigzag64,9,opt,name=theSInt64,proto3" json:"theSInt64,omitempty"`
	TheFixed32           uint32         `protobuf:"fixed32,10,opt,name=theFixed32,proto3" json:"theFixed32,omitempty"`
	TheFixed64           uint64         `protobuf:"fixed64,11,opt,name=theFixed64,proto3" json:"theFixed64,omitempty"`
	TheSFixed32          int32          `protobuf:"fixed32,12,opt,name=theSFixed32,proto3" json:"theSFixed32,omitempty"`
	TheSFixed64          int64          `protobuf:"fixed64,13,opt,name=theSFixed64,proto3" json:"theSFixed64,omitempty"`
	TheEventType         EventType      `protobuf:"varint,14,opt,name=theEventType,proto3,enum=crowdstrike.csproto.example.permessage.gogo.EventType" json:"theEventType,omitempty"`
	TheBytes             []byte         `protobuf:"bytes,15,opt,name=theBytes,proto3" json:"theBytes,omitempty"`
	TheMessage           *EmbeddedEvent `protobuf:"bytes,16,opt,name=theMessage,proto3" json:"theMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AllTheThings) Reset()         { *m = AllTheThings{} }
func (m *AllTheThings) String() string { return proto.CompactTextString(m) }
func (*AllTheThings) ProtoMessage()    {}
func (*AllTheThings) Descriptor() ([]byte, []int) {
	return fileDescriptor_77954bdb6fa16f29, []int{2}
}
func (m *AllTheThings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllTheThings.Unmarshal(m, b)
}
func (m *AllTheThings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllTheThings.Marshal(b, m, deterministic)
}
func (m *AllTheThings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllTheThings.Merge(m, src)
}
func (m *AllTheThings) XXX_Size() int {
	return xxx_messageInfo_AllTheThings.Size(m)
}
func (m *AllTheThings) XXX_DiscardUnknown() {
	xxx_messageInfo_AllTheThings.DiscardUnknown(m)
}

var xxx_messageInfo_AllTheThings proto.InternalMessageInfo

func (m *AllTheThings) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *AllTheThings) GetTheString() string {
	if m != nil {
		return m.TheString
	}
	return ""
}

func (m *AllTheThings) GetTheBool() bool {
	if m != nil {
		return m.TheBool
	}
	return false
}

func (m *AllTheThings) GetTheInt32() int32 {
	if m != nil {
		return m.TheInt32
	}
	return 0
}

func (m *AllTheThings) GetTheInt64() int64 {
	if m != nil {
		return m.TheInt64
	}
	return 0
}

func (m *AllTheThings) GetTheUInt32() uint32 {
	if m != nil {
		return m.TheUInt32
	}
	return 0
}

func (m *AllTheThings) GetTheUInt64() uint64 {
	if m != nil {
		return m.TheUInt64
	}
	return 0
}

func (m *AllTheThings) GetTheSInt32() int32 {
	if m != nil {
		return m.TheSInt32
	}
	return 0
}

func (m *AllTheThings) GetTheSInt64() int64 {
	if m != nil {
		return m.TheSInt64
	}
	return 0
}

func (m *AllTheThings) GetTheFixed32() uint32 {
	if m != nil {
		return m.TheFixed32
	}
	return 0
}

func (m *AllTheThings) GetTheFixed64() uint64 {
	if m != nil {
		return m.TheFixed64
	}
	return 0
}

func (m *AllTheThings) GetTheSFixed32() int32 {
	if m != nil {
		return m.TheSFixed32
	}
	return 0
}

func (m *AllTheThings) GetTheSFixed64() int64 {
	if m != nil {
		return m.TheSFixed64
	}
	return 0
}

func (m *AllTheThings) GetTheEventType() EventType {
	if m != nil {
		return m.TheEventType
	}
	return EventType_EVENT_TYPE_UNDEFINED
}

func (m *AllTheThings) GetTheBytes() []byte {
	if m != nil {
		return m.TheBytes
	}
	return nil
}

func (m *AllTheThings) GetTheMessage() *EmbeddedEvent {
	if m != nil {
		return m.TheMessage
	}
	return nil
}

type RepeatAllTheThings struct {
	ID                   int32            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TheStrings           []string         `protobuf:"bytes,2,rep,name=theStrings,proto3" json:"theStrings,omitempty"`
	TheBools             []bool           `protobuf:"varint,3,rep,packed,name=theBools,proto3" json:"theBools,omitempty"`
	TheInt32S            []int32          `protobuf:"varint,4,rep,packed,name=theInt32s,proto3" json:"theInt32s,omitempty"`
	TheInt64S            []int64          `protobuf:"varint,5,rep,packed,name=theInt64s,proto3" json:"theInt64s,omitempty"`
	TheUInt32S           []uint32         `protobuf:"varint,6,rep,packed,name=theUInt32s,proto3" json:"theUInt32s,omitempty"`
	TheUInt64S           []uint64         `protobuf:"varint,7,rep,packed,name=theUInt64s,proto3" json:"theUInt64s,omitempty"`
	TheSInt32S           []int32          `protobuf:"zigzag32,8,rep,packed,name=theSInt32s,proto3" json:"theSInt32s,omitempty"`
	TheSInt64S           []int64          `protobuf:"zigzag64,9,rep,packed,name=theSInt64s,proto3" json:"theSInt64s,omitempty"`
	TheFixed32S          []uint32         `protobuf:"fixed32,10,rep,packed,name=theFixed32s,proto3" json:"theFixed32s,omitempty"`
	TheFixed64S          []uint64         `protobuf:"fixed64,11,rep,packed,name=theFixed64s,proto3" json:"theFixed64s,omitempty"`
	TheSFixed32S         []int32          `protobuf:"fixed32,12,rep,packed,name=theSFixed32s,proto3" json:"theSFixed32s,omitempty"`
	TheSFixed64S         []int64          `protobuf:"fixed64,13,rep,packed,name=theSFixed64s,proto3" json:"theSFixed64s,omitempty"`
	TheEventTypes        []EventType      `protobuf:"varint,14,rep,packed,name=theEventTypes,proto3,enum=crowdstrike.csproto.example.permessage.gogo.EventType" json:"theEventTypes,omitempty"`
	TheBytes             [][]byte         `protobuf:"bytes,15,rep,name=theBytes,proto3" json:"theBytes,omitempty"`
	TheMessages          []*EmbeddedEvent `protobuf:"bytes,16,rep,name=theMessages,proto3" json:"theMessages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RepeatAllTheThings) Reset()         { *m = RepeatAllTheThings{} }
func (m *RepeatAllTheThings) String() string { return proto.CompactTextString(m) }
func (*RepeatAllTheThings) ProtoMessage()    {}
func (*RepeatAllTheThings) Descriptor() ([]byte, []int) {
	return fileDescriptor_77954bdb6fa16f29, []int{3}
}
func (m *RepeatAllTheThings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RepeatAllTheThings.Unmarshal(m, b)
}
func (m *RepeatAllTheThings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RepeatAllTheThings.Marshal(b, m, deterministic)
}
func (m *RepeatAllTheThings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RepeatAllTheThings.Merge(m, src)
}
func (m *RepeatAllTheThings) XXX_Size() int {
	return xxx_messageInfo_RepeatAllTheThings.Size(m)
}
func (m *RepeatAllTheThings) XXX_DiscardUnknown() {
	xxx_messageInfo_RepeatAllTheThings.DiscardUnknown(m)
}

var xxx_messageInfo_RepeatAllTheThings proto.InternalMessageInfo

func (m *RepeatAllTheThings) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *RepeatAllTheThings) GetTheStrings() []string {
	if m != nil {
		return m.TheStrings
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheBools() []bool {
	if m != nil {
		return m.TheBools
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheInt32S() []int32 {
	if m != nil {
		return m.TheInt32S
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheInt64S() []int64 {
	if m != nil {
		return m.TheInt64S
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheUInt32S() []uint32 {
	if m != nil {
		return m.TheUInt32S
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheUInt64S() []uint64 {
	if m != nil {
		return m.TheUInt64S
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheSInt32S() []int32 {
	if m != nil {
		return m.TheSInt32S
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheSInt64S() []int64 {
	if m != nil {
		return m.TheSInt64S
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheFixed32S() []uint32 {
	if m != nil {
		return m.TheFixed32S
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheFixed64S() []uint64 {
	if m != nil {
		return m.TheFixed64S
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheSFixed32S() []int32 {
	if m != nil {
		return m.TheSFixed32S
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheSFixed64S() []int64 {
	if m != nil {
		return m.TheSFixed64S
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheEventTypes() []EventType {
	if m != nil {
		return m.TheEventTypes
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheBytes() [][]byte {
	if m != nil {
		return m.TheBytes
	}
	return nil
}

func (m *RepeatAllTheThings) GetTheMessages() []*EmbeddedEvent {
	if m != nil {
		return m.TheMessages
	}
	return nil
}

type AllTheMaps struct {
	ToInt32              map[string]int32          `protobuf:"bytes,1,rep,name=toInt32,proto3" json:"toInt32,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ToInt64              map[string]int64          `protobuf:"bytes,2,rep,name=toInt64,proto3" json:"toInt64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ToUInt32             map[string]uint32         `protobuf:"bytes,3,rep,name=toUInt32,proto3" json:"toUInt32,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ToUInt64             map[string]uint64         `protobuf:"bytes,4,rep,name=toUInt64,proto3" json:"toUInt64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	ToString             map[string]string         `protobuf:"bytes,5,rep,name=toString,proto3" json:"toString,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ToBytes              map[string][]byte         `protobuf:"bytes,6,rep,name=toBytes,proto3" json:"toBytes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ToSInt32             map[string]int32          `protobuf:"bytes,7,rep,name=toSInt32,proto3" json:"toSInt32,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"zigzag32,2,opt,name=value,proto3"`
	ToSInt64             map[string]int64          `protobuf:"bytes,8,rep,name=toSInt64,proto3" json:"toSInt64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"zigzag64,2,opt,name=value,proto3"`
	ToFixed32            map[string]uint32         `protobuf:"bytes,9,rep,name=toFixed32,proto3" json:"toFixed32,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	ToSFixed32           map[string]int32          `protobuf:"bytes,10,rep,name=toSFixed32,proto3" json:"toSFixed32,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	ToFixed64            map[string]uint64         `protobuf:"bytes,11,rep,name=toFixed64,proto3" json:"toFixed64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	ToSFixed64           map[string]int64          `protobuf:"bytes,12,rep,name=toSFixed64,proto3" json:"toSFixed64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	ToFloat              map[string]float32        `protobuf:"bytes,13,rep,name=toFloat,proto3" json:"toFloat,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	ToDouble             map[string]float64        `protobuf:"bytes,14,rep,name=toDouble,proto3" json:"toDouble,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	ToMessage            map[string]*EmbeddedEvent `protobuf:"bytes,15,rep,name=toMessage,proto3" json:"toMessage,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ToEnum               map[string]EventType      `protobuf:"bytes,16,rep,name=toEnum,proto3" json:"toEnum,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=crowdstrike.csproto.example.permessage.gogo.EventType"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AllTheMaps) Reset()         { *m = AllTheMaps{} }
func (m *AllTheMaps) String() string { return proto.CompactTextString(m) }
func (*AllTheMaps) ProtoMessage()    {}
func (*AllTheMaps) Descriptor() ([]byte, []int) {
	return fileDescriptor_77954bdb6fa16f29, []int{4}
}
func (m *AllTheMaps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllTheMaps.Unmarshal(m, b)
}
func (m *AllTheMaps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllTheMaps.Marshal(b, m, deterministic)
}
func (m *AllTheMaps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllTheMaps.Merge(m, src)
}
func (m *AllTheMaps) XXX_Size() int {
	return xxx_messageInfo_AllTheMaps.Size(m)
}
func (m *AllTheMaps) XXX_DiscardUnknown() {
	xxx_messageInfo_AllTheMaps.DiscardUnknown(m)
}

var xxx_messageInfo_AllTheMaps proto.InternalMessageInfo

func (m *AllTheMaps) GetToInt32() map[string]int32 {
	if m != nil {
		return m.ToInt32
	}
	return nil
}

func (m *AllTheMaps) GetToInt64() map[string]int64 {
	if m != nil {
		return m.ToInt64
	}
	return nil
}

func (m *AllTheMaps) GetToUInt32() map[string]uint32 {
	if m != nil {
		return m.ToUInt32
	}
	return nil
}

func (m *AllTheMaps) GetToUInt64() map[string]uint64 {
	if m != nil {
		return m.ToUInt64
	}
	return nil
}

func (m *AllTheMaps) GetToString() map[string]string {
	if m != nil {
		return m.ToString
	}
	return nil
}

func (m *AllTheMaps) GetToBytes() map[string][]byte {
	if m != nil {
		return m.ToBytes
	}
	return nil
}

func (m *AllTheMaps) GetToSInt32() map[string]int32 {
	if m != nil {
		return m.ToSInt32
	}
	return nil
}

func (m *AllTheMaps) GetToSInt64() map[string]int64 {
	if m != nil {
		return m.ToSInt64
	}
	return nil
}

func (m *AllTheMaps) GetToFixed32() map[string]uint32 {
	if m != nil {
		return m.ToFixed32
	}
	return nil
}

func (m *AllTheMaps) GetToSFixed32() map[string]int32 {
	if m != nil {
		return m.ToSFixed32
	}
	return nil
}

func (m *AllTheMaps) GetToFixed64() map[string]uint64 {
	if m != nil {
		return m.ToFixed64
	}
	return nil
}

func (m *AllTheMaps) GetToSFixed64() map[string]int64 {
	if m != nil {
		return m.ToSFixed64
	}
	return nil
}

func (m *AllTheMaps) GetToFloat() map[string]float32 {
	if m != nil {
		return m.ToFloat
	}
	return nil
}

func (m *AllTheMaps) GetToDouble() map[string]float64 {
	if m != nil {
		return m.ToDouble
	}
	return nil
}

func (m *AllTheMaps) GetToMessage() map[string]*EmbeddedEvent {
	if m != nil {
		return m.ToMessage
	}
	return nil
}

func (m *AllTheMaps) GetToEnum() map[string]EventType {
	if m != nil {
		return m.ToEnum
	}
	return nil
}

func init() {
	proto.RegisterEnum("crowdstrike.csproto.example.permessage.gogo.EventType", EventType_name, EventType_value)
	proto.RegisterType((*TestEvent)(nil), "crowdstrike.csproto.example.permessage.gogo.TestEvent")
	proto.RegisterType((*TestEvent_NestedMsg)(nil), "crowdstrike.csproto.example.permessage.gogo.TestEvent.NestedMsg")
	proto.RegisterType((*EmbeddedEvent)(nil), "crowdstrike.csproto.example.permessage.gogo.EmbeddedEvent")
	proto.RegisterType((*AllTheThings)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheThings")
	proto.RegisterType((*RepeatAllTheThings)(nil), "crowdstrike.csproto.example.permessage.gogo.RepeatAllTheThings")
	proto.RegisterType((*AllTheMaps)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps")
	proto.RegisterMapType((map[string][]byte)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToBytesEntry")
	proto.RegisterMapType((map[string]float64)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToDoubleEntry")
	proto.RegisterMapType((map[string]EventType)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToEnumEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToFixed32Entry")
	proto.RegisterMapType((map[string]uint64)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToFixed64Entry")
	proto.RegisterMapType((map[string]float32)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToFloatEntry")
	proto.RegisterMapType((map[string]int32)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToInt32Entry")
	proto.RegisterMapType((map[string]int64)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToInt64Entry")
	proto.RegisterMapType((map[string]*EmbeddedEvent)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToMessageEntry")
	proto.RegisterMapType((map[string]int32)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToSFixed32Entry")
	proto.RegisterMapType((map[string]int64)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToSFixed64Entry")
	proto.RegisterMapType((map[string]int32)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToSInt32Entry")
	proto.RegisterMapType((map[string]int64)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToSInt64Entry")
	proto.RegisterMapType((map[string]string)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToStringEntry")
	proto.RegisterMapType((map[string]uint32)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToUInt32Entry")
	proto.RegisterMapType((map[string]uint64)(nil), "crowdstrike.csproto.example.permessage.gogo.AllTheMaps.ToUInt64Entry")
}

func init() { proto.RegisterFile("gogo_permessage_example.proto", fileDescriptor_77954bdb6fa16f29) }

var fileDescriptor_77954bdb6fa16f29 = []byte{
	// 1281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0x6d, 0x6f, 0xdb, 0x36,
	0x17, 0xad, 0x2c, 0xbf, 0x89, 0xb1, 0x63, 0x97, 0xe8, 0x53, 0x10, 0x41, 0x9f, 0x4d, 0x30, 0x30,
	0x40, 0xe8, 0x00, 0x19, 0x48, 0x0d, 0xa1, 0xe8, 0x36, 0x60, 0x4d, 0xed, 0xac, 0x01, 0x16, 0xb7,
	0x60, 0x9c, 0x6e, 0xeb, 0x8a, 0x05, 0x72, 0xcd, 0xd8, 0x5a, 0x65, 0xd3, 0x33, 0xe9, 0xb6, 0xf9,
	0x36, 0x60, 0xfb, 0x55, 0xfb, 0x31, 0xfb, 0x2d, 0x03, 0x5f, 0x24, 0x51, 0x71, 0x36, 0x27, 0x51,
	0xbf, 0xf9, 0x1e, 0xea, 0x9c, 0x7b, 0xaf, 0xee, 0x21, 0x29, 0x83, 0xff, 0x4f, 0xe9, 0x94, 0x9e,
	0x2d, 0xc9, 0x6a, 0x4e, 0x18, 0x0b, 0xa7, 0xe4, 0x8c, 0x7c, 0x0c, 0xe7, 0xcb, 0x98, 0xf8, 0xcb,
	0x15, 0xe5, 0x14, 0x7e, 0xf9, 0x76, 0x45, 0x3f, 0x4c, 0x18, 0x5f, 0x45, 0xef, 0x88, 0xff, 0x96,
	0x49, 0xd0, 0x4f, 0x1f, 0x49, 0x59, 0xbe, 0x50, 0xd9, 0xfb, 0x7c, 0x4a, 0xe9, 0x34, 0x26, 0x5d,
	0xf9, 0xd4, 0x78, 0x7d, 0xde, 0xe5, 0xd1, 0x9c, 0x30, 0x1e, 0xce, 0x97, 0x4a, 0xad, 0xf3, 0x97,
	0x0d, 0x9c, 0x11, 0x61, 0x7c, 0xf0, 0x9e, 0x2c, 0x38, 0x84, 0xa0, 0xbc, 0x08, 0xe7, 0x04, 0x59,
	0xae, 0xe5, 0x39, 0x58, 0xfe, 0x16, 0x58, 0xb4, 0x38, 0xa7, 0xa8, 0xa4, 0x30, 0xf1, 0x1b, 0x3e,
	0x00, 0x4e, 0xc4, 0x9e, 0x7e, 0x20, 0x8c, 0xce, 0x09, 0xb2, 0x5d, 0xcb, 0xab, 0xe3, 0x0c, 0x80,
	0xf7, 0x41, 0x35, 0x0e, 0xc7, 0x24, 0x66, 0xa8, 0xec, 0xda, 0x9e, 0x83, 0x75, 0x04, 0x5f, 0x81,
	0x3a, 0x99, 0x8f, 0xc9, 0x64, 0x42, 0x26, 0xa8, 0xe2, 0x5a, 0xde, 0xce, 0xfe, 0x13, 0xff, 0x06,
	0xcd, 0xf8, 0x03, 0x4d, 0x96, 0xb5, 0xe2, 0x54, 0x0b, 0xde, 0x03, 0xe5, 0x5f, 0xc9, 0x24, 0x42,
	0x55, 0x51, 0xc8, 0xf3, 0x3b, 0x58, 0x46, 0x02, 0x65, 0x11, 0x9f, 0xa1, 0x5a, 0x82, 0x8a, 0x08,
	0xde, 0x07, 0x15, 0xca, 0x67, 0x64, 0x85, 0xea, 0xa2, 0x9d, 0xe7, 0x77, 0xb0, 0x0a, 0xe1, 0x8f,
	0xa0, 0xba, 0x20, 0x8c, 0x93, 0x09, 0x72, 0x64, 0x65, 0xdf, 0xde, 0xa8, 0xb2, 0xf4, 0x0d, 0xfa,
	0x43, 0x29, 0x72, 0xcc, 0xa6, 0x58, 0xeb, 0xc1, 0x87, 0xa0, 0xc4, 0x19, 0x02, 0x52, 0x75, 0xcf,
	0x57, 0xf3, 0xf0, 0x93, 0x79, 0xf8, 0xa3, 0x64, 0x1e, 0xb8, 0xc4, 0xd9, 0xde, 0x17, 0xc0, 0x49,
	0x05, 0x20, 0x02, 0xb5, 0x09, 0xe1, 0x61, 0x14, 0x33, 0x3d, 0x8f, 0x24, 0x3c, 0xa8, 0x82, 0xf2,
	0x32, 0xe4, 0xb3, 0xce, 0x1f, 0x16, 0x68, 0xe6, 0x5e, 0x0a, 0xdc, 0x05, 0xa5, 0xa3, 0xbe, 0x7c,
	0xbc, 0x82, 0x4b, 0x47, 0x7d, 0x78, 0x0f, 0x54, 0x18, 0x5f, 0x9f, 0x9f, 0xeb, 0xe9, 0xa9, 0x00,
	0x7a, 0xa0, 0x75, 0x1e, 0xbe, 0xa7, 0xab, 0x88, 0x93, 0xe1, 0x7a, 0x3e, 0x26, 0x2b, 0x86, 0x6c,
	0xd7, 0xf6, 0x2a, 0xf8, 0x32, 0x0c, 0x3b, 0xa0, 0xb1, 0x0a, 0x17, 0x13, 0x3a, 0x1f, 0xcd, 0xa2,
	0xc5, 0x54, 0x0d, 0xb4, 0x81, 0x73, 0x58, 0xe7, 0xef, 0x32, 0x68, 0x3c, 0x8d, 0xe3, 0xd1, 0x8c,
	0x28, 0x60, 0xa3, 0x88, 0x07, 0xc0, 0xe1, 0x33, 0x72, 0xc2, 0x57, 0xd1, 0x62, 0xaa, 0x0b, 0xc9,
	0x00, 0xd1, 0x26, 0x9f, 0x91, 0x03, 0x4a, 0x63, 0xed, 0xa4, 0x24, 0x84, 0x7b, 0xa0, 0xce, 0x67,
	0xe4, 0x68, 0xc1, 0x1f, 0xed, 0xa3, 0xb2, 0x54, 0x4b, 0xe3, 0x6c, 0x2d, 0xe8, 0x49, 0x2f, 0xd9,
	0x38, 0x8d, 0x75, 0xbe, 0x53, 0x45, 0x14, 0xa6, 0x68, 0xe2, 0x0c, 0x30, 0x56, 0x83, 0x9e, 0x34,
	0x47, 0x19, 0x67, 0x40, 0x52, 0xab, 0xe2, 0x0a, 0x8f, 0xdc, 0xc5, 0x19, 0x60, 0xac, 0x06, 0x3d,
	0x69, 0x14, 0x88, 0x33, 0x00, 0x7e, 0x06, 0x00, 0x9f, 0x91, 0xc3, 0xe8, 0x23, 0x99, 0x3c, 0xda,
	0x97, 0x13, 0xaf, 0x61, 0x03, 0x31, 0xd7, 0x83, 0x1e, 0xda, 0x71, 0x2d, 0xaf, 0x8a, 0x0d, 0x04,
	0xba, 0x60, 0x47, 0x88, 0x25, 0x02, 0x0d, 0xd7, 0xf2, 0x5a, 0xd8, 0x84, 0x72, 0x4f, 0x04, 0x3d,
	0xd4, 0x74, 0x2d, 0xaf, 0x8d, 0x4d, 0x08, 0xbe, 0x06, 0x0d, 0x3e, 0x23, 0xd2, 0x0c, 0xa3, 0x8b,
	0x25, 0x41, 0xbb, 0xae, 0xe5, 0xed, 0xee, 0x07, 0x37, 0xdb, 0x67, 0x09, 0x1b, 0xe7, 0xb4, 0xf4,
	0x3b, 0x3f, 0xb8, 0xe0, 0x84, 0xa1, 0x96, 0x6b, 0x79, 0x0d, 0x9c, 0xc6, 0xf0, 0xb5, 0xec, 0xed,
	0x58, 0xc9, 0xa0, 0x76, 0xe1, 0xdd, 0x6d, 0xa8, 0x75, 0x7e, 0xaf, 0x00, 0x88, 0xc9, 0x92, 0x84,
	0xfc, 0x3f, 0x6d, 0xa6, 0x5e, 0xaf, 0x72, 0x15, 0x43, 0x25, 0x79, 0xf4, 0x18, 0x48, 0x52, 0x3e,
	0xa5, 0xb1, 0xb2, 0x7b, 0x1d, 0xa7, 0xb1, 0x1e, 0xac, 0x1c, 0xb2, 0x32, 0x79, 0x05, 0x67, 0x40,
	0xb6, 0x1a, 0xf4, 0x18, 0xaa, 0xb8, 0xb6, 0x67, 0xe3, 0x0c, 0xd0, 0x79, 0x4f, 0x35, 0xb9, 0xea,
	0xda, 0x5e, 0x13, 0x1b, 0x88, 0xb1, 0x2e, 0xe8, 0x35, 0xd7, 0xf6, 0xca, 0xd8, 0x40, 0x92, 0xba,
	0x35, 0xbf, 0xee, 0xda, 0xde, 0x5d, 0x6c, 0x20, 0xc6, 0xba, 0xe0, 0x3b, 0xae, 0xed, 0x41, 0x6c,
	0x20, 0xda, 0x14, 0xda, 0x22, 0xe2, 0xa4, 0xb1, 0xbd, 0x1a, 0x36, 0x21, 0xf3, 0x09, 0x21, 0xb1,
	0xe3, 0xda, 0x5e, 0x15, 0x9b, 0x90, 0xd8, 0xe7, 0x86, 0xcf, 0x18, 0x6a, 0xb8, 0xb6, 0xd7, 0xc2,
	0x39, 0x2c, 0xf7, 0x8c, 0x90, 0x69, 0xba, 0xb6, 0xd7, 0xc6, 0x39, 0x0c, 0xbe, 0x01, 0x4d, 0xd3,
	0x32, 0x0c, 0xed, 0xba, 0x76, 0x01, 0xff, 0xe5, 0xc5, 0x2e, 0x19, 0xd0, 0xce, 0x19, 0xf0, 0x8d,
	0xec, 0x51, 0x5b, 0x86, 0xa1, 0xb6, 0x6b, 0x17, 0x74, 0xa0, 0x29, 0xd7, 0xf9, 0xf3, 0x7f, 0x00,
	0x28, 0xf3, 0x1d, 0x87, 0x4b, 0x06, 0x7f, 0x01, 0x35, 0x4e, 0xd5, 0x19, 0x61, 0xc9, 0x44, 0xfd,
	0x1b, 0x25, 0xca, 0x94, 0xfc, 0x91, 0x92, 0x19, 0x2c, 0xf8, 0xea, 0x02, 0x27, 0xa2, 0xa9, 0x7e,
	0xd0, 0x93, 0x3e, 0x2e, 0xaa, 0x1f, 0xf4, 0x4c, 0xfd, 0xa0, 0x07, 0x43, 0x50, 0xe7, 0x54, 0x1f,
	0x90, 0xb6, 0x4c, 0x30, 0xb8, 0x7d, 0x82, 0x53, 0xa3, 0x83, 0x54, 0x36, 0x4b, 0x11, 0xf4, 0xe4,
	0x86, 0x2a, 0x9c, 0x22, 0x69, 0x22, 0x95, 0x55, 0x29, 0xf4, 0xb5, 0x52, 0x29, 0x9a, 0x42, 0xe9,
	0xa4, 0x29, 0xf4, 0xe5, 0x24, 0x07, 0xa1, 0x0c, 0x57, 0x2d, 0x3a, 0x08, 0x29, 0x93, 0x0e, 0x42,
	0xb9, 0x56, 0xb5, 0xa0, 0x06, 0x51, 0x2b, 0xdc, 0x42, 0x7e, 0x10, 0x27, 0xc6, 0x20, 0xf4, 0x95,
	0x55, 0xff, 0x14, 0x29, 0x8c, 0x41, 0xe8, 0x8b, 0x6f, 0x02, 0x1c, 0x4e, 0x93, 0x6b, 0xcb, 0x91,
	0x39, 0x0e, 0x6f, 0x9f, 0x43, 0x0b, 0xa9, 0x24, 0x99, 0x30, 0x9c, 0x02, 0xc0, 0xe9, 0x49, 0x76,
	0xbd, 0x8a, 0x34, 0xdf, 0x15, 0x68, 0x25, 0x97, 0xc7, 0x90, 0x36, 0xda, 0x91, 0xd7, 0xf4, 0xa7,
	0x68, 0x27, 0x79, 0x67, 0x99, 0xb0, 0xd9, 0x4e, 0xd0, 0x93, 0x07, 0xee, 0x27, 0x68, 0x27, 0xc9,
	0x63, 0x48, 0x2b, 0x0f, 0x1f, 0xc6, 0x34, 0xe4, 0xf2, 0xc8, 0x2e, 0xe4, 0x61, 0x29, 0x93, 0x7a,
	0x58, 0x46, 0xca, 0x60, 0x7d, 0xba, 0x1e, 0xc7, 0x44, 0x1e, 0xf7, 0x85, 0x0c, 0xa6, 0x74, 0x52,
	0x83, 0xa9, 0x50, 0x4d, 0x24, 0xf9, 0xb8, 0x68, 0x15, 0x9d, 0x88, 0x16, 0x4a, 0x27, 0xa2, 0x63,
	0xf8, 0x33, 0xa8, 0x72, 0x3a, 0x58, 0xac, 0xe7, 0xfa, 0xf6, 0x78, 0x76, 0xfb, 0x14, 0x42, 0x45,
	0xe9, 0x6b, 0xc9, 0xbd, 0x27, 0xa0, 0x61, 0x9e, 0xf5, 0xb0, 0x0d, 0xec, 0x77, 0xe4, 0x42, 0x7f,
	0xd9, 0x8b, 0x9f, 0xe2, 0x5b, 0xfd, 0x7d, 0x18, 0xaf, 0x89, 0xfc, 0x44, 0xae, 0x60, 0x15, 0x3c,
	0x29, 0x3d, 0xb6, 0x52, 0xae, 0x9e, 0xee, 0x36, 0xae, 0x6d, 0x72, 0xbf, 0x02, 0xcd, 0xdc, 0x11,
	0xbd, 0x8d, 0xdc, 0xbc, 0x92, 0x7c, 0xcd, 0xcc, 0xe5, 0x0d, 0xb2, 0x71, 0xac, 0x6e, 0x23, 0x3b,
	0x1b, 0x2d, 0x67, 0x27, 0xe6, 0x36, 0x6e, 0x63, 0x33, 0xf1, 0xf5, 0x5b, 0xbe, 0x7b, 0x25, 0xf9,
	0x9a, 0x2d, 0x43, 0x93, 0xfc, 0x35, 0xd8, 0xcd, 0x9f, 0x5f, 0xdb, 0xd8, 0x35, 0x93, 0xfd, 0x0d,
	0x68, 0x5d, 0x3a, 0x96, 0xb6, 0xd1, 0x5b, 0x57, 0x27, 0xbf, 0x66, 0xe9, 0xd5, 0x7f, 0x49, 0x7e,
	0x4d, 0x7a, 0x7b, 0x63, 0x5e, 0xd9, 0xe9, 0xb0, 0x8d, 0x5b, 0xda, 0x78, 0xe5, 0xc6, 0xc6, 0xdf,
	0x46, 0xb6, 0x4c, 0xf2, 0x47, 0xd1, 0xb5, 0xb9, 0xa3, 0xaf, 0x60, 0xbf, 0x34, 0xd9, 0xc5, 0xbe,
	0x0a, 0x8d, 0xcc, 0xbf, 0x81, 0x1d, 0x63, 0xa3, 0x5f, 0x91, 0xf6, 0x7b, 0x33, 0xed, 0xed, 0x3f,
	0x82, 0xb3, 0x94, 0x0f, 0x8f, 0x81, 0x93, 0xfd, 0x1d, 0x43, 0xe0, 0xde, 0xe0, 0xd5, 0x60, 0x38,
	0x3a, 0x1b, 0xfd, 0xf4, 0x72, 0x70, 0x76, 0x3a, 0xec, 0x0f, 0x0e, 0x8f, 0x86, 0x83, 0x7e, 0xfb,
	0x0e, 0x84, 0x60, 0xd7, 0x58, 0x79, 0x31, 0x1c, 0xb4, 0xad, 0x4b, 0xd8, 0xe8, 0x87, 0x17, 0xed,
	0xd2, 0xc1, 0xe3, 0xd7, 0xc1, 0x34, 0xe2, 0xb3, 0xf5, 0xd8, 0x7f, 0x4b, 0xe7, 0xdd, 0x67, 0xa2,
	0xba, 0x13, 0x59, 0x5d, 0x57, 0x57, 0xd7, 0xd5, 0xd5, 0x75, 0xb3, 0xea, 0xba, 0xa2, 0xba, 0x71,
	0x55, 0x2e, 0x3f, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0x90, 0xbd, 0x74, 0xf9, 0xac, 0x12, 0x00,
	0x00,
}