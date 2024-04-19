// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: schema.proto

package schema_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScalarType int32

const (
	ScalarType_BOOLEAN ScalarType = 0
	ScalarType_INTEGER ScalarType = 1
	ScalarType_LONG    ScalarType = 3
	ScalarType_FLOAT   ScalarType = 4
	ScalarType_DOUBLE  ScalarType = 5
	ScalarType_BYTES   ScalarType = 6
	ScalarType_STRING  ScalarType = 7
)

// Enum value maps for ScalarType.
var (
	ScalarType_name = map[int32]string{
		0: "BOOLEAN",
		1: "INTEGER",
		3: "LONG",
		4: "FLOAT",
		5: "DOUBLE",
		6: "BYTES",
		7: "STRING",
	}
	ScalarType_value = map[string]int32{
		"BOOLEAN": 0,
		"INTEGER": 1,
		"LONG":    3,
		"FLOAT":   4,
		"DOUBLE":  5,
		"BYTES":   6,
		"STRING":  7,
	}
)

func (x ScalarType) Enum() *ScalarType {
	p := new(ScalarType)
	*p = x
	return p
}

func (x ScalarType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ScalarType) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_proto_enumTypes[0].Descriptor()
}

func (ScalarType) Type() protoreflect.EnumType {
	return &file_schema_proto_enumTypes[0]
}

func (x ScalarType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ScalarType.Descriptor instead.
func (ScalarType) EnumDescriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

type RecordType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields []*Field `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
}

func (x *RecordType) Reset() {
	*x = RecordType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordType) ProtoMessage() {}

func (x *RecordType) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordType.ProtoReflect.Descriptor instead.
func (*RecordType) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

func (x *RecordType) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type       *Type  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Index      int32  `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	IsRepeated bool   `protobuf:"varint,4,opt,name=is_repeated,json=isRepeated,proto3" json:"is_repeated,omitempty"`
	IsRequired bool   `protobuf:"varint,5,opt,name=is_required,json=isRequired,proto3" json:"is_required,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{1}
}

func (x *Field) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Field) GetType() *Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Field) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Field) GetIsRepeated() bool {
	if x != nil {
		return x.IsRepeated
	}
	return false
}

func (x *Field) GetIsRequired() bool {
	if x != nil {
		return x.IsRequired
	}
	return false
}

type Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*Type_ScalarType
	//	*Type_RecordType
	//	*Type_ListType
	Kind isType_Kind `protobuf_oneof:"kind"`
}

func (x *Type) Reset() {
	*x = Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Type) ProtoMessage() {}

func (x *Type) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Type.ProtoReflect.Descriptor instead.
func (*Type) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{2}
}

func (m *Type) GetKind() isType_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Type) GetScalarType() ScalarType {
	if x, ok := x.GetKind().(*Type_ScalarType); ok {
		return x.ScalarType
	}
	return ScalarType_BOOLEAN
}

func (x *Type) GetRecordType() *RecordType {
	if x, ok := x.GetKind().(*Type_RecordType); ok {
		return x.RecordType
	}
	return nil
}

func (x *Type) GetListType() *ListType {
	if x, ok := x.GetKind().(*Type_ListType); ok {
		return x.ListType
	}
	return nil
}

type isType_Kind interface {
	isType_Kind()
}

type Type_ScalarType struct {
	ScalarType ScalarType `protobuf:"varint,1,opt,name=scalar_type,json=scalarType,proto3,enum=schema_pb.ScalarType,oneof"`
}

type Type_RecordType struct {
	RecordType *RecordType `protobuf:"bytes,2,opt,name=record_type,json=recordType,proto3,oneof"`
}

type Type_ListType struct {
	ListType *ListType `protobuf:"bytes,3,opt,name=list_type,json=listType,proto3,oneof"`
}

func (*Type_ScalarType) isType_Kind() {}

func (*Type_RecordType) isType_Kind() {}

func (*Type_ListType) isType_Kind() {}

type ListType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ElementType *Type `protobuf:"bytes,1,opt,name=element_type,json=elementType,proto3" json:"element_type,omitempty"`
}

func (x *ListType) Reset() {
	*x = ListType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListType) ProtoMessage() {}

func (x *ListType) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListType.ProtoReflect.Descriptor instead.
func (*ListType) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{3}
}

func (x *ListType) GetElementType() *Type {
	if x != nil {
		return x.ElementType
	}
	return nil
}

// /////////////////////////
// value definition
// /////////////////////////
type RecordValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fields map[string]*Value `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RecordValue) Reset() {
	*x = RecordValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordValue) ProtoMessage() {}

func (x *RecordValue) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordValue.ProtoReflect.Descriptor instead.
func (*RecordValue) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{4}
}

func (x *RecordValue) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//
	//	*Value_BoolValue
	//	*Value_Int32Value
	//	*Value_Int64Value
	//	*Value_FloatValue
	//	*Value_DoubleValue
	//	*Value_BytesValue
	//	*Value_StringValue
	//	*Value_ListValue
	//	*Value_RecordValue
	Kind isValue_Kind `protobuf_oneof:"kind"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{5}
}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *Value) GetBoolValue() bool {
	if x, ok := x.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *Value) GetInt32Value() int32 {
	if x, ok := x.GetKind().(*Value_Int32Value); ok {
		return x.Int32Value
	}
	return 0
}

func (x *Value) GetInt64Value() int64 {
	if x, ok := x.GetKind().(*Value_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (x *Value) GetFloatValue() float32 {
	if x, ok := x.GetKind().(*Value_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (x *Value) GetDoubleValue() float64 {
	if x, ok := x.GetKind().(*Value_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *Value) GetBytesValue() []byte {
	if x, ok := x.GetKind().(*Value_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (x *Value) GetStringValue() string {
	if x, ok := x.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Value) GetListValue() *ListValue {
	if x, ok := x.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

func (x *Value) GetRecordValue() *RecordValue {
	if x, ok := x.GetKind().(*Value_RecordValue); ok {
		return x.RecordValue
	}
	return nil
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,1,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_Int32Value struct {
	Int32Value int32 `protobuf:"varint,2,opt,name=int32_value,json=int32Value,proto3,oneof"`
}

type Value_Int64Value struct {
	Int64Value int64 `protobuf:"varint,3,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Value_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,4,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,6,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,7,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,14,opt,name=list_value,json=listValue,proto3,oneof"`
}

type Value_RecordValue struct {
	RecordValue *RecordValue `protobuf:"bytes,15,opt,name=record_value,json=recordValue,proto3,oneof"`
}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_Int32Value) isValue_Kind() {}

func (*Value_Int64Value) isValue_Kind() {}

func (*Value_FloatValue) isValue_Kind() {}

func (*Value_DoubleValue) isValue_Kind() {}

func (*Value_BytesValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_ListValue) isValue_Kind() {}

func (*Value_RecordValue) isValue_Kind() {}

type ListValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ListValue) Reset() {
	*x = ListValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListValue) ProtoMessage() {}

func (x *ListValue) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListValue.ProtoReflect.Descriptor instead.
func (*ListValue) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{6}
}

func (x *ListValue) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_schema_proto protoreflect.FileDescriptor

var file_schema_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x22, 0x36, 0x0a, 0x0a, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x5f, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x22, 0x98, 0x01, 0x0a, 0x05, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73,
	0x5f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x69, 0x73, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0xb6, 0x01, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x38, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x3e, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x32, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x5f, 0x70, 0x62, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70,
	0x62, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x1a, 0x4b, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x26, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xfa,
	0x02, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x09,
	0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x21, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x35, 0x0a, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x35, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x2a, 0x5e, 0x0a, 0x0a, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4f, 0x4f, 0x4c, 0x45, 0x41, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45, 0x52, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f,
	0x4e, 0x47, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x04, 0x12,
	0x0a, 0x0a, 0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x42,
	0x59, 0x54, 0x45, 0x53, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47,
	0x10, 0x07, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x65, 0x61, 0x77, 0x65, 0x65, 0x64, 0x66, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x77, 0x65,
	0x65, 0x64, 0x66, 0x73, 0x2f, 0x77, 0x65, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_proto_rawDescOnce sync.Once
	file_schema_proto_rawDescData = file_schema_proto_rawDesc
)

func file_schema_proto_rawDescGZIP() []byte {
	file_schema_proto_rawDescOnce.Do(func() {
		file_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_proto_rawDescData)
	})
	return file_schema_proto_rawDescData
}

var file_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_schema_proto_goTypes = []interface{}{
	(ScalarType)(0),     // 0: schema_pb.ScalarType
	(*RecordType)(nil),  // 1: schema_pb.RecordType
	(*Field)(nil),       // 2: schema_pb.Field
	(*Type)(nil),        // 3: schema_pb.Type
	(*ListType)(nil),    // 4: schema_pb.ListType
	(*RecordValue)(nil), // 5: schema_pb.RecordValue
	(*Value)(nil),       // 6: schema_pb.Value
	(*ListValue)(nil),   // 7: schema_pb.ListValue
	nil,                 // 8: schema_pb.RecordValue.FieldsEntry
}
var file_schema_proto_depIdxs = []int32{
	2,  // 0: schema_pb.RecordType.fields:type_name -> schema_pb.Field
	3,  // 1: schema_pb.Field.type:type_name -> schema_pb.Type
	0,  // 2: schema_pb.Type.scalar_type:type_name -> schema_pb.ScalarType
	1,  // 3: schema_pb.Type.record_type:type_name -> schema_pb.RecordType
	4,  // 4: schema_pb.Type.list_type:type_name -> schema_pb.ListType
	3,  // 5: schema_pb.ListType.element_type:type_name -> schema_pb.Type
	8,  // 6: schema_pb.RecordValue.fields:type_name -> schema_pb.RecordValue.FieldsEntry
	7,  // 7: schema_pb.Value.list_value:type_name -> schema_pb.ListValue
	5,  // 8: schema_pb.Value.record_value:type_name -> schema_pb.RecordValue
	6,  // 9: schema_pb.ListValue.values:type_name -> schema_pb.Value
	6,  // 10: schema_pb.RecordValue.FieldsEntry.value:type_name -> schema_pb.Value
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_schema_proto_init() }
func file_schema_proto_init() {
	if File_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Type); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_schema_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Type_ScalarType)(nil),
		(*Type_RecordType)(nil),
		(*Type_ListType)(nil),
	}
	file_schema_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Value_BoolValue)(nil),
		(*Value_Int32Value)(nil),
		(*Value_Int64Value)(nil),
		(*Value_FloatValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_BytesValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_ListValue)(nil),
		(*Value_RecordValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_proto_goTypes,
		DependencyIndexes: file_schema_proto_depIdxs,
		EnumInfos:         file_schema_proto_enumTypes,
		MessageInfos:      file_schema_proto_msgTypes,
	}.Build()
	File_schema_proto = out.File
	file_schema_proto_rawDesc = nil
	file_schema_proto_goTypes = nil
	file_schema_proto_depIdxs = nil
}
