// Copied from https://github.com/capnproto/go-capnp/blob/v3.0.1-alpha.2/internal/schema/schema.capnp.go
// That seems to be generated; but it didn't change in 3 years, so it's good enough? I guess

package schema

import (
	math "math"
	strconv "strconv"

	capnp "capnproto.org/go/capnp/v3"
)

// Constants defined in schema.capnp.
const (
	Field_noDiscriminant = uint16(65535)
)

type Node capnp.Struct
type Node_structNode Node
type Node_enum Node
type Node_interface Node
type Node_const Node
type Node_annotation Node
type Node_Which uint16

const (
	Node_Which_file       Node_Which = 0
	Node_Which_structNode Node_Which = 1
	Node_Which_enum       Node_Which = 2
	Node_Which_interface  Node_Which = 3
	Node_Which_const      Node_Which = 4
	Node_Which_annotation Node_Which = 5
)

func (w Node_Which) String() string {
	const s = "filestructNodeenuminterfaceconstannotation"
	switch w {
	case Node_Which_file:
		return s[0:4]
	case Node_Which_structNode:
		return s[4:14]
	case Node_Which_enum:
		return s[14:18]
	case Node_Which_interface:
		return s[18:27]
	case Node_Which_const:
		return s[27:32]
	case Node_Which_annotation:
		return s[32:42]

	}
	return "Node_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Node_TypeID is the unique identifier for the type Node.
const Node_TypeID = 0xe682ab4cf923a417

func NewNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 6})
	return Node(st), err
}

func NewRootNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 6})
	return Node(st), err
}

func ReadRootNode(msg *capnp.Message) (Node, error) {
	root, err := msg.Root()
	return Node(root.Struct()), err
}

func (s Node) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Node) DecodeFromPtr(p capnp.Ptr) Node {
	return Node(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Node) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}

func (s Node) Which() Node_Which {
	return Node_Which(capnp.Struct(s).Uint16(12))
}
func (s Node) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Node) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Node) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Node) Id() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Node) SetId(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s Node) DisplayName() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Node) HasDisplayName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Node) DisplayNameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Node) SetDisplayName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Node) DisplayNamePrefixLength() uint32 {
	return capnp.Struct(s).Uint32(8)
}

func (s Node) SetDisplayNamePrefixLength(v uint32) {
	capnp.Struct(s).SetUint32(8, v)
}

func (s Node) ScopeId() uint64 {
	return capnp.Struct(s).Uint64(16)
}

func (s Node) SetScopeId(v uint64) {
	capnp.Struct(s).SetUint64(16, v)
}

func (s Node) Parameters() (Node_Parameter_List, error) {
	p, err := capnp.Struct(s).Ptr(5)
	return Node_Parameter_List(p.List()), err
}

func (s Node) HasParameters() bool {
	return capnp.Struct(s).HasPtr(5)
}

func (s Node) SetParameters(v Node_Parameter_List) error {
	return capnp.Struct(s).SetPtr(5, v.ToPtr())
}

// NewParameters sets the parameters field to a newly
// allocated Node_Parameter_List, preferring placement in s's segment.
func (s Node) NewParameters(n int32) (Node_Parameter_List, error) {
	l, err := NewNode_Parameter_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Node_Parameter_List{}, err
	}
	err = capnp.Struct(s).SetPtr(5, l.ToPtr())
	return l, err
}
func (s Node) IsGeneric() bool {
	return capnp.Struct(s).Bit(288)
}

func (s Node) SetIsGeneric(v bool) {
	capnp.Struct(s).SetBit(288, v)
}

func (s Node) NestedNodes() (Node_NestedNode_List, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return Node_NestedNode_List(p.List()), err
}

func (s Node) HasNestedNodes() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Node) SetNestedNodes(v Node_NestedNode_List) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewNestedNodes sets the nestedNodes field to a newly
// allocated Node_NestedNode_List, preferring placement in s's segment.
func (s Node) NewNestedNodes(n int32) (Node_NestedNode_List, error) {
	l, err := NewNode_NestedNode_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Node_NestedNode_List{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}
func (s Node) Annotations() (Annotation_List, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return Annotation_List(p.List()), err
}

func (s Node) HasAnnotations() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s Node) SetAnnotations(v Annotation_List) error {
	return capnp.Struct(s).SetPtr(2, v.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Node) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = capnp.Struct(s).SetPtr(2, l.ToPtr())
	return l, err
}
func (s Node) SetFile() {
	capnp.Struct(s).SetUint16(12, 0)

}

func (s Node) StructNode() Node_structNode { return Node_structNode(s) }

func (s Node) SetStructNode() {
	capnp.Struct(s).SetUint16(12, 1)
}

func (s Node_structNode) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Node_structNode) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Node_structNode) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Node_structNode) DataWordCount() uint16 {
	return capnp.Struct(s).Uint16(14)
}

func (s Node_structNode) SetDataWordCount(v uint16) {
	capnp.Struct(s).SetUint16(14, v)
}

func (s Node_structNode) PointerCount() uint16 {
	return capnp.Struct(s).Uint16(24)
}

func (s Node_structNode) SetPointerCount(v uint16) {
	capnp.Struct(s).SetUint16(24, v)
}

func (s Node_structNode) PreferredListEncoding() ElementSize {
	return ElementSize(capnp.Struct(s).Uint16(26))
}

func (s Node_structNode) SetPreferredListEncoding(v ElementSize) {
	capnp.Struct(s).SetUint16(26, uint16(v))
}

func (s Node_structNode) IsGroup() bool {
	return capnp.Struct(s).Bit(224)
}

func (s Node_structNode) SetIsGroup(v bool) {
	capnp.Struct(s).SetBit(224, v)
}

func (s Node_structNode) DiscriminantCount() uint16 {
	return capnp.Struct(s).Uint16(30)
}

func (s Node_structNode) SetDiscriminantCount(v uint16) {
	capnp.Struct(s).SetUint16(30, v)
}

func (s Node_structNode) DiscriminantOffset() uint32 {
	return capnp.Struct(s).Uint32(32)
}

func (s Node_structNode) SetDiscriminantOffset(v uint32) {
	capnp.Struct(s).SetUint32(32, v)
}

func (s Node_structNode) Fields() (Field_List, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return Field_List(p.List()), err
}

func (s Node_structNode) HasFields() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s Node_structNode) SetFields(v Field_List) error {
	return capnp.Struct(s).SetPtr(3, v.ToPtr())
}

// NewFields sets the fields field to a newly
// allocated Field_List, preferring placement in s's segment.
func (s Node_structNode) NewFields(n int32) (Field_List, error) {
	l, err := NewField_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Field_List{}, err
	}
	err = capnp.Struct(s).SetPtr(3, l.ToPtr())
	return l, err
}
func (s Node) Enum() Node_enum { return Node_enum(s) }

func (s Node) SetEnum() {
	capnp.Struct(s).SetUint16(12, 2)
}

func (s Node_enum) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Node_enum) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Node_enum) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Node_enum) Enumerants() (Enumerant_List, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return Enumerant_List(p.List()), err
}

func (s Node_enum) HasEnumerants() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s Node_enum) SetEnumerants(v Enumerant_List) error {
	return capnp.Struct(s).SetPtr(3, v.ToPtr())
}

// NewEnumerants sets the enumerants field to a newly
// allocated Enumerant_List, preferring placement in s's segment.
func (s Node_enum) NewEnumerants(n int32) (Enumerant_List, error) {
	l, err := NewEnumerant_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Enumerant_List{}, err
	}
	err = capnp.Struct(s).SetPtr(3, l.ToPtr())
	return l, err
}
func (s Node) Interface() Node_interface { return Node_interface(s) }

func (s Node) SetInterface() {
	capnp.Struct(s).SetUint16(12, 3)
}

func (s Node_interface) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Node_interface) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Node_interface) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Node_interface) Methods() (Method_List, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return Method_List(p.List()), err
}

func (s Node_interface) HasMethods() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s Node_interface) SetMethods(v Method_List) error {
	return capnp.Struct(s).SetPtr(3, v.ToPtr())
}

// NewMethods sets the methods field to a newly
// allocated Method_List, preferring placement in s's segment.
func (s Node_interface) NewMethods(n int32) (Method_List, error) {
	l, err := NewMethod_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Method_List{}, err
	}
	err = capnp.Struct(s).SetPtr(3, l.ToPtr())
	return l, err
}
func (s Node_interface) Superclasses() (Superclass_List, error) {
	p, err := capnp.Struct(s).Ptr(4)
	return Superclass_List(p.List()), err
}

func (s Node_interface) HasSuperclasses() bool {
	return capnp.Struct(s).HasPtr(4)
}

func (s Node_interface) SetSuperclasses(v Superclass_List) error {
	return capnp.Struct(s).SetPtr(4, v.ToPtr())
}

// NewSuperclasses sets the superclasses field to a newly
// allocated Superclass_List, preferring placement in s's segment.
func (s Node_interface) NewSuperclasses(n int32) (Superclass_List, error) {
	l, err := NewSuperclass_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Superclass_List{}, err
	}
	err = capnp.Struct(s).SetPtr(4, l.ToPtr())
	return l, err
}
func (s Node) Const() Node_const { return Node_const(s) }

func (s Node) SetConst() {
	capnp.Struct(s).SetUint16(12, 4)
}

func (s Node_const) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Node_const) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Node_const) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Node_const) Type() (Type, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return Type(p.Struct()), err
}

func (s Node_const) HasType() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s Node_const) SetType(v Type) error {
	return capnp.Struct(s).SetPtr(3, capnp.Struct(v).ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Node_const) NewType() (Type, error) {
	ss, err := NewType(capnp.Struct(s).Segment())
	if err != nil {
		return Type{}, err
	}
	err = capnp.Struct(s).SetPtr(3, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Node_const) Value() (Value, error) {
	p, err := capnp.Struct(s).Ptr(4)
	return Value(p.Struct()), err
}

func (s Node_const) HasValue() bool {
	return capnp.Struct(s).HasPtr(4)
}

func (s Node_const) SetValue(v Value) error {
	return capnp.Struct(s).SetPtr(4, capnp.Struct(v).ToPtr())
}

// NewValue sets the value field to a newly
// allocated Value struct, preferring placement in s's segment.
func (s Node_const) NewValue() (Value, error) {
	ss, err := NewValue(capnp.Struct(s).Segment())
	if err != nil {
		return Value{}, err
	}
	err = capnp.Struct(s).SetPtr(4, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Node) Annotation() Node_annotation { return Node_annotation(s) }

func (s Node) SetAnnotation() {
	capnp.Struct(s).SetUint16(12, 5)
}

func (s Node_annotation) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Node_annotation) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Node_annotation) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Node_annotation) Type() (Type, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return Type(p.Struct()), err
}

func (s Node_annotation) HasType() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s Node_annotation) SetType(v Type) error {
	return capnp.Struct(s).SetPtr(3, capnp.Struct(v).ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Node_annotation) NewType() (Type, error) {
	ss, err := NewType(capnp.Struct(s).Segment())
	if err != nil {
		return Type{}, err
	}
	err = capnp.Struct(s).SetPtr(3, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Node_annotation) TargetsFile() bool {
	return capnp.Struct(s).Bit(112)
}

func (s Node_annotation) SetTargetsFile(v bool) {
	capnp.Struct(s).SetBit(112, v)
}

func (s Node_annotation) TargetsConst() bool {
	return capnp.Struct(s).Bit(113)
}

func (s Node_annotation) SetTargetsConst(v bool) {
	capnp.Struct(s).SetBit(113, v)
}

func (s Node_annotation) TargetsEnum() bool {
	return capnp.Struct(s).Bit(114)
}

func (s Node_annotation) SetTargetsEnum(v bool) {
	capnp.Struct(s).SetBit(114, v)
}

func (s Node_annotation) TargetsEnumerant() bool {
	return capnp.Struct(s).Bit(115)
}

func (s Node_annotation) SetTargetsEnumerant(v bool) {
	capnp.Struct(s).SetBit(115, v)
}

func (s Node_annotation) TargetsStruct() bool {
	return capnp.Struct(s).Bit(116)
}

func (s Node_annotation) SetTargetsStruct(v bool) {
	capnp.Struct(s).SetBit(116, v)
}

func (s Node_annotation) TargetsField() bool {
	return capnp.Struct(s).Bit(117)
}

func (s Node_annotation) SetTargetsField(v bool) {
	capnp.Struct(s).SetBit(117, v)
}

func (s Node_annotation) TargetsUnion() bool {
	return capnp.Struct(s).Bit(118)
}

func (s Node_annotation) SetTargetsUnion(v bool) {
	capnp.Struct(s).SetBit(118, v)
}

func (s Node_annotation) TargetsGroup() bool {
	return capnp.Struct(s).Bit(119)
}

func (s Node_annotation) SetTargetsGroup(v bool) {
	capnp.Struct(s).SetBit(119, v)
}

func (s Node_annotation) TargetsInterface() bool {
	return capnp.Struct(s).Bit(120)
}

func (s Node_annotation) SetTargetsInterface(v bool) {
	capnp.Struct(s).SetBit(120, v)
}

func (s Node_annotation) TargetsMethod() bool {
	return capnp.Struct(s).Bit(121)
}

func (s Node_annotation) SetTargetsMethod(v bool) {
	capnp.Struct(s).SetBit(121, v)
}

func (s Node_annotation) TargetsParam() bool {
	return capnp.Struct(s).Bit(122)
}

func (s Node_annotation) SetTargetsParam(v bool) {
	capnp.Struct(s).SetBit(122, v)
}

func (s Node_annotation) TargetsAnnotation() bool {
	return capnp.Struct(s).Bit(123)
}

func (s Node_annotation) SetTargetsAnnotation(v bool) {
	capnp.Struct(s).SetBit(123, v)
}

// Node_List is a list of Node.
type Node_List = capnp.StructList[Node]

// NewNode creates a new list of Node.
func NewNode_List(s *capnp.Segment, sz int32) (Node_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 40, PointerCount: 6}, sz)
	return capnp.StructList[Node](l), err
}

type Node_Parameter capnp.Struct

// Node_Parameter_TypeID is the unique identifier for the type Node_Parameter.
const Node_Parameter_TypeID = 0xb9521bccf10fa3b1

func NewNode_Parameter(s *capnp.Segment) (Node_Parameter, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Node_Parameter(st), err
}

func NewRootNode_Parameter(s *capnp.Segment) (Node_Parameter, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Node_Parameter(st), err
}

func ReadRootNode_Parameter(msg *capnp.Message) (Node_Parameter, error) {
	root, err := msg.Root()
	return Node_Parameter(root.Struct()), err
}

func (s Node_Parameter) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Node_Parameter) DecodeFromPtr(p capnp.Ptr) Node_Parameter {
	return Node_Parameter(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Node_Parameter) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Node_Parameter) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Node_Parameter) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Node_Parameter) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Node_Parameter) Name() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Node_Parameter) HasName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Node_Parameter) NameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Node_Parameter) SetName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

// Node_Parameter_List is a list of Node_Parameter.
type Node_Parameter_List = capnp.StructList[Node_Parameter]

// NewNode_Parameter creates a new list of Node_Parameter.
func NewNode_Parameter_List(s *capnp.Segment, sz int32) (Node_Parameter_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Node_Parameter](l), err
}

type Node_NestedNode capnp.Struct

// Node_NestedNode_TypeID is the unique identifier for the type Node_NestedNode.
const Node_NestedNode_TypeID = 0xdebf55bbfa0fc242

func NewNode_NestedNode(s *capnp.Segment) (Node_NestedNode, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Node_NestedNode(st), err
}

func NewRootNode_NestedNode(s *capnp.Segment) (Node_NestedNode, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Node_NestedNode(st), err
}

func ReadRootNode_NestedNode(msg *capnp.Message) (Node_NestedNode, error) {
	root, err := msg.Root()
	return Node_NestedNode(root.Struct()), err
}

func (s Node_NestedNode) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Node_NestedNode) DecodeFromPtr(p capnp.Ptr) Node_NestedNode {
	return Node_NestedNode(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Node_NestedNode) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Node_NestedNode) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Node_NestedNode) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Node_NestedNode) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Node_NestedNode) Name() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Node_NestedNode) HasName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Node_NestedNode) NameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Node_NestedNode) SetName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Node_NestedNode) Id() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Node_NestedNode) SetId(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

// Node_NestedNode_List is a list of Node_NestedNode.
type Node_NestedNode_List = capnp.StructList[Node_NestedNode]

// NewNode_NestedNode creates a new list of Node_NestedNode.
func NewNode_NestedNode_List(s *capnp.Segment, sz int32) (Node_NestedNode_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[Node_NestedNode](l), err
}

type Node_SourceInfo capnp.Struct

// Node_SourceInfo_TypeID is the unique identifier for the type Node_SourceInfo.
const Node_SourceInfo_TypeID = 0xf38e1de3041357ae

func NewNode_SourceInfo(s *capnp.Segment) (Node_SourceInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Node_SourceInfo(st), err
}

func NewRootNode_SourceInfo(s *capnp.Segment) (Node_SourceInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Node_SourceInfo(st), err
}

func ReadRootNode_SourceInfo(msg *capnp.Message) (Node_SourceInfo, error) {
	root, err := msg.Root()
	return Node_SourceInfo(root.Struct()), err
}

func (s Node_SourceInfo) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Node_SourceInfo) DecodeFromPtr(p capnp.Ptr) Node_SourceInfo {
	return Node_SourceInfo(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Node_SourceInfo) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Node_SourceInfo) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Node_SourceInfo) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Node_SourceInfo) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Node_SourceInfo) Id() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Node_SourceInfo) SetId(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s Node_SourceInfo) DocComment() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Node_SourceInfo) HasDocComment() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Node_SourceInfo) DocCommentBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Node_SourceInfo) SetDocComment(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Node_SourceInfo) Members() (Node_SourceInfo_Member_List, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return Node_SourceInfo_Member_List(p.List()), err
}

func (s Node_SourceInfo) HasMembers() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Node_SourceInfo) SetMembers(v Node_SourceInfo_Member_List) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewMembers sets the members field to a newly
// allocated Node_SourceInfo_Member_List, preferring placement in s's segment.
func (s Node_SourceInfo) NewMembers(n int32) (Node_SourceInfo_Member_List, error) {
	l, err := NewNode_SourceInfo_Member_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Node_SourceInfo_Member_List{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}

// Node_SourceInfo_List is a list of Node_SourceInfo.
type Node_SourceInfo_List = capnp.StructList[Node_SourceInfo]

// NewNode_SourceInfo creates a new list of Node_SourceInfo.
func NewNode_SourceInfo_List(s *capnp.Segment, sz int32) (Node_SourceInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return capnp.StructList[Node_SourceInfo](l), err
}

type Node_SourceInfo_Member capnp.Struct

// Node_SourceInfo_Member_TypeID is the unique identifier for the type Node_SourceInfo_Member.
const Node_SourceInfo_Member_TypeID = 0xc2ba9038898e1fa2

func NewNode_SourceInfo_Member(s *capnp.Segment) (Node_SourceInfo_Member, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Node_SourceInfo_Member(st), err
}

func NewRootNode_SourceInfo_Member(s *capnp.Segment) (Node_SourceInfo_Member, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Node_SourceInfo_Member(st), err
}

func ReadRootNode_SourceInfo_Member(msg *capnp.Message) (Node_SourceInfo_Member, error) {
	root, err := msg.Root()
	return Node_SourceInfo_Member(root.Struct()), err
}

func (s Node_SourceInfo_Member) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Node_SourceInfo_Member) DecodeFromPtr(p capnp.Ptr) Node_SourceInfo_Member {
	return Node_SourceInfo_Member(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Node_SourceInfo_Member) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Node_SourceInfo_Member) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Node_SourceInfo_Member) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Node_SourceInfo_Member) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Node_SourceInfo_Member) DocComment() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Node_SourceInfo_Member) HasDocComment() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Node_SourceInfo_Member) DocCommentBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Node_SourceInfo_Member) SetDocComment(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

// Node_SourceInfo_Member_List is a list of Node_SourceInfo_Member.
type Node_SourceInfo_Member_List = capnp.StructList[Node_SourceInfo_Member]

// NewNode_SourceInfo_Member creates a new list of Node_SourceInfo_Member.
func NewNode_SourceInfo_Member_List(s *capnp.Segment, sz int32) (Node_SourceInfo_Member_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Node_SourceInfo_Member](l), err
}

type Field capnp.Struct
type Field_slot Field
type Field_group Field
type Field_ordinal Field
type Field_Which uint16

const (
	Field_Which_slot  Field_Which = 0
	Field_Which_group Field_Which = 1
)

func (w Field_Which) String() string {
	const s = "slotgroup"
	switch w {
	case Field_Which_slot:
		return s[0:4]
	case Field_Which_group:
		return s[4:9]

	}
	return "Field_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type Field_ordinal_Which uint16

const (
	Field_ordinal_Which_implicit Field_ordinal_Which = 0
	Field_ordinal_Which_explicit Field_ordinal_Which = 1
)

func (w Field_ordinal_Which) String() string {
	const s = "implicitexplicit"
	switch w {
	case Field_ordinal_Which_implicit:
		return s[0:8]
	case Field_ordinal_Which_explicit:
		return s[8:16]

	}
	return "Field_ordinal_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Field_TypeID is the unique identifier for the type Field.
const Field_TypeID = 0x9aad50a41f4af45f

func NewField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4})
	return Field(st), err
}

func NewRootField(s *capnp.Segment) (Field, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4})
	return Field(st), err
}

func ReadRootField(msg *capnp.Message) (Field, error) {
	root, err := msg.Root()
	return Field(root.Struct()), err
}

func (s Field) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Field) DecodeFromPtr(p capnp.Ptr) Field {
	return Field(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Field) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}

func (s Field) Which() Field_Which {
	return Field_Which(capnp.Struct(s).Uint16(8))
}
func (s Field) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Field) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Field) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Field) Name() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Field) HasName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Field) NameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Field) SetName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Field) CodeOrder() uint16 {
	return capnp.Struct(s).Uint16(0)
}

func (s Field) SetCodeOrder(v uint16) {
	capnp.Struct(s).SetUint16(0, v)
}

func (s Field) Annotations() (Annotation_List, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return Annotation_List(p.List()), err
}

func (s Field) HasAnnotations() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Field) SetAnnotations(v Annotation_List) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Field) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}
func (s Field) DiscriminantValue() uint16 {
	return capnp.Struct(s).Uint16(2) ^ 65535
}

func (s Field) SetDiscriminantValue(v uint16) {
	capnp.Struct(s).SetUint16(2, v^65535)
}

func (s Field) Slot() Field_slot { return Field_slot(s) }

func (s Field) SetSlot() {
	capnp.Struct(s).SetUint16(8, 0)
}

func (s Field_slot) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Field_slot) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Field_slot) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Field_slot) Offset() uint32 {
	return capnp.Struct(s).Uint32(4)
}

func (s Field_slot) SetOffset(v uint32) {
	capnp.Struct(s).SetUint32(4, v)
}

func (s Field_slot) Type() (Type, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return Type(p.Struct()), err
}

func (s Field_slot) HasType() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s Field_slot) SetType(v Type) error {
	return capnp.Struct(s).SetPtr(2, capnp.Struct(v).ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Field_slot) NewType() (Type, error) {
	ss, err := NewType(capnp.Struct(s).Segment())
	if err != nil {
		return Type{}, err
	}
	err = capnp.Struct(s).SetPtr(2, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Field_slot) DefaultValue() (Value, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return Value(p.Struct()), err
}

func (s Field_slot) HasDefaultValue() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s Field_slot) SetDefaultValue(v Value) error {
	return capnp.Struct(s).SetPtr(3, capnp.Struct(v).ToPtr())
}

// NewDefaultValue sets the defaultValue field to a newly
// allocated Value struct, preferring placement in s's segment.
func (s Field_slot) NewDefaultValue() (Value, error) {
	ss, err := NewValue(capnp.Struct(s).Segment())
	if err != nil {
		return Value{}, err
	}
	err = capnp.Struct(s).SetPtr(3, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Field_slot) HadExplicitDefault() bool {
	return capnp.Struct(s).Bit(128)
}

func (s Field_slot) SetHadExplicitDefault(v bool) {
	capnp.Struct(s).SetBit(128, v)
}

func (s Field) Group() Field_group { return Field_group(s) }

func (s Field) SetGroup() {
	capnp.Struct(s).SetUint16(8, 1)
}

func (s Field_group) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Field_group) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Field_group) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Field_group) TypeId() uint64 {
	return capnp.Struct(s).Uint64(16)
}

func (s Field_group) SetTypeId(v uint64) {
	capnp.Struct(s).SetUint64(16, v)
}

func (s Field) Ordinal() Field_ordinal { return Field_ordinal(s) }

func (s Field_ordinal) Which() Field_ordinal_Which {
	return Field_ordinal_Which(capnp.Struct(s).Uint16(10))
}
func (s Field_ordinal) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Field_ordinal) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Field_ordinal) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Field_ordinal) SetImplicit() {
	capnp.Struct(s).SetUint16(10, 0)

}

func (s Field_ordinal) Explicit() uint16 {
	if capnp.Struct(s).Uint16(10) != 1 {
		panic("Which() != explicit")
	}
	return capnp.Struct(s).Uint16(12)
}

func (s Field_ordinal) SetExplicit(v uint16) {
	capnp.Struct(s).SetUint16(10, 1)
	capnp.Struct(s).SetUint16(12, v)
}

// Field_List is a list of Field.
type Field_List = capnp.StructList[Field]

// NewField creates a new list of Field.
func NewField_List(s *capnp.Segment, sz int32) (Field_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 4}, sz)
	return capnp.StructList[Field](l), err
}

type Enumerant capnp.Struct

// Enumerant_TypeID is the unique identifier for the type Enumerant.
const Enumerant_TypeID = 0x978a7cebdc549a4d

func NewEnumerant(s *capnp.Segment) (Enumerant, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Enumerant(st), err
}

func NewRootEnumerant(s *capnp.Segment) (Enumerant, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Enumerant(st), err
}

func ReadRootEnumerant(msg *capnp.Message) (Enumerant, error) {
	root, err := msg.Root()
	return Enumerant(root.Struct()), err
}

func (s Enumerant) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Enumerant) DecodeFromPtr(p capnp.Ptr) Enumerant {
	return Enumerant(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Enumerant) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Enumerant) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Enumerant) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Enumerant) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Enumerant) Name() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Enumerant) HasName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Enumerant) NameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Enumerant) SetName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Enumerant) CodeOrder() uint16 {
	return capnp.Struct(s).Uint16(0)
}

func (s Enumerant) SetCodeOrder(v uint16) {
	capnp.Struct(s).SetUint16(0, v)
}

func (s Enumerant) Annotations() (Annotation_List, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return Annotation_List(p.List()), err
}

func (s Enumerant) HasAnnotations() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Enumerant) SetAnnotations(v Annotation_List) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Enumerant) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}

// Enumerant_List is a list of Enumerant.
type Enumerant_List = capnp.StructList[Enumerant]

// NewEnumerant creates a new list of Enumerant.
func NewEnumerant_List(s *capnp.Segment, sz int32) (Enumerant_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return capnp.StructList[Enumerant](l), err
}

type Superclass capnp.Struct

// Superclass_TypeID is the unique identifier for the type Superclass.
const Superclass_TypeID = 0xa9962a9ed0a4d7f8

func NewSuperclass(s *capnp.Segment) (Superclass, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Superclass(st), err
}

func NewRootSuperclass(s *capnp.Segment) (Superclass, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Superclass(st), err
}

func ReadRootSuperclass(msg *capnp.Message) (Superclass, error) {
	root, err := msg.Root()
	return Superclass(root.Struct()), err
}

func (s Superclass) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Superclass) DecodeFromPtr(p capnp.Ptr) Superclass {
	return Superclass(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Superclass) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Superclass) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Superclass) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Superclass) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Superclass) Id() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Superclass) SetId(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s Superclass) Brand() (Brand, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Brand(p.Struct()), err
}

func (s Superclass) HasBrand() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Superclass) SetBrand(v Brand) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Superclass) NewBrand() (Brand, error) {
	ss, err := NewBrand(capnp.Struct(s).Segment())
	if err != nil {
		return Brand{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

// Superclass_List is a list of Superclass.
type Superclass_List = capnp.StructList[Superclass]

// NewSuperclass creates a new list of Superclass.
func NewSuperclass_List(s *capnp.Segment, sz int32) (Superclass_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[Superclass](l), err
}

type Method capnp.Struct

// Method_TypeID is the unique identifier for the type Method.
const Method_TypeID = 0x9500cce23b334d80

func NewMethod(s *capnp.Segment) (Method, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 5})
	return Method(st), err
}

func NewRootMethod(s *capnp.Segment) (Method, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 5})
	return Method(st), err
}

func ReadRootMethod(msg *capnp.Message) (Method, error) {
	root, err := msg.Root()
	return Method(root.Struct()), err
}

func (s Method) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Method) DecodeFromPtr(p capnp.Ptr) Method {
	return Method(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Method) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Method) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Method) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Method) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Method) Name() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Method) HasName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Method) NameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Method) SetName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Method) CodeOrder() uint16 {
	return capnp.Struct(s).Uint16(0)
}

func (s Method) SetCodeOrder(v uint16) {
	capnp.Struct(s).SetUint16(0, v)
}

func (s Method) ImplicitParameters() (Node_Parameter_List, error) {
	p, err := capnp.Struct(s).Ptr(4)
	return Node_Parameter_List(p.List()), err
}

func (s Method) HasImplicitParameters() bool {
	return capnp.Struct(s).HasPtr(4)
}

func (s Method) SetImplicitParameters(v Node_Parameter_List) error {
	return capnp.Struct(s).SetPtr(4, v.ToPtr())
}

// NewImplicitParameters sets the implicitParameters field to a newly
// allocated Node_Parameter_List, preferring placement in s's segment.
func (s Method) NewImplicitParameters(n int32) (Node_Parameter_List, error) {
	l, err := NewNode_Parameter_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Node_Parameter_List{}, err
	}
	err = capnp.Struct(s).SetPtr(4, l.ToPtr())
	return l, err
}
func (s Method) ParamStructType() uint64 {
	return capnp.Struct(s).Uint64(8)
}

func (s Method) SetParamStructType(v uint64) {
	capnp.Struct(s).SetUint64(8, v)
}

func (s Method) ParamBrand() (Brand, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return Brand(p.Struct()), err
}

func (s Method) HasParamBrand() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s Method) SetParamBrand(v Brand) error {
	return capnp.Struct(s).SetPtr(2, capnp.Struct(v).ToPtr())
}

// NewParamBrand sets the paramBrand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Method) NewParamBrand() (Brand, error) {
	ss, err := NewBrand(capnp.Struct(s).Segment())
	if err != nil {
		return Brand{}, err
	}
	err = capnp.Struct(s).SetPtr(2, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Method) ResultStructType() uint64 {
	return capnp.Struct(s).Uint64(16)
}

func (s Method) SetResultStructType(v uint64) {
	capnp.Struct(s).SetUint64(16, v)
}

func (s Method) ResultBrand() (Brand, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return Brand(p.Struct()), err
}

func (s Method) HasResultBrand() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s Method) SetResultBrand(v Brand) error {
	return capnp.Struct(s).SetPtr(3, capnp.Struct(v).ToPtr())
}

// NewResultBrand sets the resultBrand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Method) NewResultBrand() (Brand, error) {
	ss, err := NewBrand(capnp.Struct(s).Segment())
	if err != nil {
		return Brand{}, err
	}
	err = capnp.Struct(s).SetPtr(3, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Method) Annotations() (Annotation_List, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return Annotation_List(p.List()), err
}

func (s Method) HasAnnotations() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Method) SetAnnotations(v Annotation_List) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewAnnotations sets the annotations field to a newly
// allocated Annotation_List, preferring placement in s's segment.
func (s Method) NewAnnotations(n int32) (Annotation_List, error) {
	l, err := NewAnnotation_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Annotation_List{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}

// Method_List is a list of Method.
type Method_List = capnp.StructList[Method]

// NewMethod creates a new list of Method.
func NewMethod_List(s *capnp.Segment, sz int32) (Method_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 5}, sz)
	return capnp.StructList[Method](l), err
}

type Type capnp.Struct
type Type_list Type
type Type_enum Type
type Type_structType Type
type Type_interface Type
type Type_anyPointer Type
type Type_anyPointer_unconstrained Type
type Type_anyPointer_parameter Type
type Type_anyPointer_implicitMethodParameter Type
type Type_Which uint16

const (
	Type_Which_void       Type_Which = 0
	Type_Which_bool       Type_Which = 1
	Type_Which_int8       Type_Which = 2
	Type_Which_int16      Type_Which = 3
	Type_Which_int32      Type_Which = 4
	Type_Which_int64      Type_Which = 5
	Type_Which_uint8      Type_Which = 6
	Type_Which_uint16     Type_Which = 7
	Type_Which_uint32     Type_Which = 8
	Type_Which_uint64     Type_Which = 9
	Type_Which_float32    Type_Which = 10
	Type_Which_float64    Type_Which = 11
	Type_Which_text       Type_Which = 12
	Type_Which_data       Type_Which = 13
	Type_Which_list       Type_Which = 14
	Type_Which_enum       Type_Which = 15
	Type_Which_structType Type_Which = 16
	Type_Which_interface  Type_Which = 17
	Type_Which_anyPointer Type_Which = 18
)

func (w Type_Which) String() string {
	const s = "voidboolint8int16int32int64uint8uint16uint32uint64float32float64textdatalistenumstructTypeinterfaceanyPointer"
	switch w {
	case Type_Which_void:
		return s[0:4]
	case Type_Which_bool:
		return s[4:8]
	case Type_Which_int8:
		return s[8:12]
	case Type_Which_int16:
		return s[12:17]
	case Type_Which_int32:
		return s[17:22]
	case Type_Which_int64:
		return s[22:27]
	case Type_Which_uint8:
		return s[27:32]
	case Type_Which_uint16:
		return s[32:38]
	case Type_Which_uint32:
		return s[38:44]
	case Type_Which_uint64:
		return s[44:50]
	case Type_Which_float32:
		return s[50:57]
	case Type_Which_float64:
		return s[57:64]
	case Type_Which_text:
		return s[64:68]
	case Type_Which_data:
		return s[68:72]
	case Type_Which_list:
		return s[72:76]
	case Type_Which_enum:
		return s[76:80]
	case Type_Which_structType:
		return s[80:90]
	case Type_Which_interface:
		return s[90:99]
	case Type_Which_anyPointer:
		return s[99:109]

	}
	return "Type_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type Type_anyPointer_Which uint16

const (
	Type_anyPointer_Which_unconstrained           Type_anyPointer_Which = 0
	Type_anyPointer_Which_parameter               Type_anyPointer_Which = 1
	Type_anyPointer_Which_implicitMethodParameter Type_anyPointer_Which = 2
)

func (w Type_anyPointer_Which) String() string {
	const s = "unconstrainedparameterimplicitMethodParameter"
	switch w {
	case Type_anyPointer_Which_unconstrained:
		return s[0:13]
	case Type_anyPointer_Which_parameter:
		return s[13:22]
	case Type_anyPointer_Which_implicitMethodParameter:
		return s[22:45]

	}
	return "Type_anyPointer_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

type Type_anyPointer_unconstrained_Which uint16

const (
	Type_anyPointer_unconstrained_Which_anyKind    Type_anyPointer_unconstrained_Which = 0
	Type_anyPointer_unconstrained_Which_struct     Type_anyPointer_unconstrained_Which = 1
	Type_anyPointer_unconstrained_Which_list       Type_anyPointer_unconstrained_Which = 2
	Type_anyPointer_unconstrained_Which_capability Type_anyPointer_unconstrained_Which = 3
)

func (w Type_anyPointer_unconstrained_Which) String() string {
	const s = "anyKindstructlistcapability"
	switch w {
	case Type_anyPointer_unconstrained_Which_anyKind:
		return s[0:7]
	case Type_anyPointer_unconstrained_Which_struct:
		return s[7:13]
	case Type_anyPointer_unconstrained_Which_list:
		return s[13:17]
	case Type_anyPointer_unconstrained_Which_capability:
		return s[17:27]

	}
	return "Type_anyPointer_unconstrained_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Type_TypeID is the unique identifier for the type Type.
const Type_TypeID = 0xd07378ede1f9cc60

func NewType(s *capnp.Segment) (Type, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	return Type(st), err
}

func NewRootType(s *capnp.Segment) (Type, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1})
	return Type(st), err
}

func ReadRootType(msg *capnp.Message) (Type, error) {
	root, err := msg.Root()
	return Type(root.Struct()), err
}

func (s Type) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Type) DecodeFromPtr(p capnp.Ptr) Type {
	return Type(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Type) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}

func (s Type) Which() Type_Which {
	return Type_Which(capnp.Struct(s).Uint16(0))
}
func (s Type) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Type) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Type) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Type) SetVoid() {
	capnp.Struct(s).SetUint16(0, 0)

}

func (s Type) SetBool() {
	capnp.Struct(s).SetUint16(0, 1)

}

func (s Type) SetInt8() {
	capnp.Struct(s).SetUint16(0, 2)

}

func (s Type) SetInt16() {
	capnp.Struct(s).SetUint16(0, 3)

}

func (s Type) SetInt32() {
	capnp.Struct(s).SetUint16(0, 4)

}

func (s Type) SetInt64() {
	capnp.Struct(s).SetUint16(0, 5)

}

func (s Type) SetUint8() {
	capnp.Struct(s).SetUint16(0, 6)

}

func (s Type) SetUint16() {
	capnp.Struct(s).SetUint16(0, 7)

}

func (s Type) SetUint32() {
	capnp.Struct(s).SetUint16(0, 8)

}

func (s Type) SetUint64() {
	capnp.Struct(s).SetUint16(0, 9)

}

func (s Type) SetFloat32() {
	capnp.Struct(s).SetUint16(0, 10)

}

func (s Type) SetFloat64() {
	capnp.Struct(s).SetUint16(0, 11)

}

func (s Type) SetText() {
	capnp.Struct(s).SetUint16(0, 12)

}

func (s Type) SetData() {
	capnp.Struct(s).SetUint16(0, 13)

}

func (s Type) List() Type_list { return Type_list(s) }

func (s Type) SetList() {
	capnp.Struct(s).SetUint16(0, 14)
}

func (s Type_list) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Type_list) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Type_list) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Type_list) ElementType() (Type, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Type(p.Struct()), err
}

func (s Type_list) HasElementType() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Type_list) SetElementType(v Type) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewElementType sets the elementType field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Type_list) NewElementType() (Type, error) {
	ss, err := NewType(capnp.Struct(s).Segment())
	if err != nil {
		return Type{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Type) Enum() Type_enum { return Type_enum(s) }

func (s Type) SetEnum() {
	capnp.Struct(s).SetUint16(0, 15)
}

func (s Type_enum) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Type_enum) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Type_enum) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Type_enum) TypeId() uint64 {
	return capnp.Struct(s).Uint64(8)
}

func (s Type_enum) SetTypeId(v uint64) {
	capnp.Struct(s).SetUint64(8, v)
}

func (s Type_enum) Brand() (Brand, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Brand(p.Struct()), err
}

func (s Type_enum) HasBrand() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Type_enum) SetBrand(v Brand) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Type_enum) NewBrand() (Brand, error) {
	ss, err := NewBrand(capnp.Struct(s).Segment())
	if err != nil {
		return Brand{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Type) StructType() Type_structType { return Type_structType(s) }

func (s Type) SetStructType() {
	capnp.Struct(s).SetUint16(0, 16)
}

func (s Type_structType) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Type_structType) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Type_structType) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Type_structType) TypeId() uint64 {
	return capnp.Struct(s).Uint64(8)
}

func (s Type_structType) SetTypeId(v uint64) {
	capnp.Struct(s).SetUint64(8, v)
}

func (s Type_structType) Brand() (Brand, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Brand(p.Struct()), err
}

func (s Type_structType) HasBrand() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Type_structType) SetBrand(v Brand) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Type_structType) NewBrand() (Brand, error) {
	ss, err := NewBrand(capnp.Struct(s).Segment())
	if err != nil {
		return Brand{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Type) Interface() Type_interface { return Type_interface(s) }

func (s Type) SetInterface() {
	capnp.Struct(s).SetUint16(0, 17)
}

func (s Type_interface) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Type_interface) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Type_interface) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Type_interface) TypeId() uint64 {
	return capnp.Struct(s).Uint64(8)
}

func (s Type_interface) SetTypeId(v uint64) {
	capnp.Struct(s).SetUint64(8, v)
}

func (s Type_interface) Brand() (Brand, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Brand(p.Struct()), err
}

func (s Type_interface) HasBrand() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Type_interface) SetBrand(v Brand) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Type_interface) NewBrand() (Brand, error) {
	ss, err := NewBrand(capnp.Struct(s).Segment())
	if err != nil {
		return Brand{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Type) AnyPointer() Type_anyPointer { return Type_anyPointer(s) }

func (s Type) SetAnyPointer() {
	capnp.Struct(s).SetUint16(0, 18)
}

func (s Type_anyPointer) Which() Type_anyPointer_Which {
	return Type_anyPointer_Which(capnp.Struct(s).Uint16(8))
}
func (s Type_anyPointer) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Type_anyPointer) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Type_anyPointer) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Type_anyPointer) Unconstrained() Type_anyPointer_unconstrained {
	return Type_anyPointer_unconstrained(s)
}

func (s Type_anyPointer) SetUnconstrained() {
	capnp.Struct(s).SetUint16(8, 0)
}

func (s Type_anyPointer_unconstrained) Which() Type_anyPointer_unconstrained_Which {
	return Type_anyPointer_unconstrained_Which(capnp.Struct(s).Uint16(10))
}
func (s Type_anyPointer_unconstrained) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Type_anyPointer_unconstrained) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Type_anyPointer_unconstrained) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Type_anyPointer_unconstrained) SetAnyKind() {
	capnp.Struct(s).SetUint16(10, 0)

}

func (s Type_anyPointer_unconstrained) SetStruct() {
	capnp.Struct(s).SetUint16(10, 1)

}

func (s Type_anyPointer_unconstrained) SetList() {
	capnp.Struct(s).SetUint16(10, 2)

}

func (s Type_anyPointer_unconstrained) SetCapability() {
	capnp.Struct(s).SetUint16(10, 3)

}

func (s Type_anyPointer) Parameter() Type_anyPointer_parameter { return Type_anyPointer_parameter(s) }

func (s Type_anyPointer) SetParameter() {
	capnp.Struct(s).SetUint16(8, 1)
}

func (s Type_anyPointer_parameter) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Type_anyPointer_parameter) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Type_anyPointer_parameter) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Type_anyPointer_parameter) ScopeId() uint64 {
	return capnp.Struct(s).Uint64(16)
}

func (s Type_anyPointer_parameter) SetScopeId(v uint64) {
	capnp.Struct(s).SetUint64(16, v)
}

func (s Type_anyPointer_parameter) ParameterIndex() uint16 {
	return capnp.Struct(s).Uint16(10)
}

func (s Type_anyPointer_parameter) SetParameterIndex(v uint16) {
	capnp.Struct(s).SetUint16(10, v)
}

func (s Type_anyPointer) ImplicitMethodParameter() Type_anyPointer_implicitMethodParameter {
	return Type_anyPointer_implicitMethodParameter(s)
}

func (s Type_anyPointer) SetImplicitMethodParameter() {
	capnp.Struct(s).SetUint16(8, 2)
}

func (s Type_anyPointer_implicitMethodParameter) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Type_anyPointer_implicitMethodParameter) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Type_anyPointer_implicitMethodParameter) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Type_anyPointer_implicitMethodParameter) ParameterIndex() uint16 {
	return capnp.Struct(s).Uint16(10)
}

func (s Type_anyPointer_implicitMethodParameter) SetParameterIndex(v uint16) {
	capnp.Struct(s).SetUint16(10, v)
}

// Type_List is a list of Type.
type Type_List = capnp.StructList[Type]

// NewType creates a new list of Type.
func NewType_List(s *capnp.Segment, sz int32) (Type_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 1}, sz)
	return capnp.StructList[Type](l), err
}

type Brand capnp.Struct

// Brand_TypeID is the unique identifier for the type Brand.
const Brand_TypeID = 0x903455f06065422b

func NewBrand(s *capnp.Segment) (Brand, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Brand(st), err
}

func NewRootBrand(s *capnp.Segment) (Brand, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Brand(st), err
}

func ReadRootBrand(msg *capnp.Message) (Brand, error) {
	root, err := msg.Root()
	return Brand(root.Struct()), err
}

func (s Brand) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Brand) DecodeFromPtr(p capnp.Ptr) Brand {
	return Brand(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Brand) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Brand) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Brand) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Brand) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Brand) Scopes() (Brand_Scope_List, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Brand_Scope_List(p.List()), err
}

func (s Brand) HasScopes() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Brand) SetScopes(v Brand_Scope_List) error {
	return capnp.Struct(s).SetPtr(0, v.ToPtr())
}

// NewScopes sets the scopes field to a newly
// allocated Brand_Scope_List, preferring placement in s's segment.
func (s Brand) NewScopes(n int32) (Brand_Scope_List, error) {
	l, err := NewBrand_Scope_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Brand_Scope_List{}, err
	}
	err = capnp.Struct(s).SetPtr(0, l.ToPtr())
	return l, err
}

// Brand_List is a list of Brand.
type Brand_List = capnp.StructList[Brand]

// NewBrand creates a new list of Brand.
func NewBrand_List(s *capnp.Segment, sz int32) (Brand_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Brand](l), err
}

type Brand_Scope capnp.Struct
type Brand_Scope_Which uint16

const (
	Brand_Scope_Which_bind    Brand_Scope_Which = 0
	Brand_Scope_Which_inherit Brand_Scope_Which = 1
)

func (w Brand_Scope_Which) String() string {
	const s = "bindinherit"
	switch w {
	case Brand_Scope_Which_bind:
		return s[0:4]
	case Brand_Scope_Which_inherit:
		return s[4:11]

	}
	return "Brand_Scope_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Brand_Scope_TypeID is the unique identifier for the type Brand_Scope.
const Brand_Scope_TypeID = 0xabd73485a9636bc9

func NewBrand_Scope(s *capnp.Segment) (Brand_Scope, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Brand_Scope(st), err
}

func NewRootBrand_Scope(s *capnp.Segment) (Brand_Scope, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Brand_Scope(st), err
}

func ReadRootBrand_Scope(msg *capnp.Message) (Brand_Scope, error) {
	root, err := msg.Root()
	return Brand_Scope(root.Struct()), err
}

func (s Brand_Scope) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Brand_Scope) DecodeFromPtr(p capnp.Ptr) Brand_Scope {
	return Brand_Scope(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Brand_Scope) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}

func (s Brand_Scope) Which() Brand_Scope_Which {
	return Brand_Scope_Which(capnp.Struct(s).Uint16(8))
}
func (s Brand_Scope) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Brand_Scope) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Brand_Scope) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Brand_Scope) ScopeId() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Brand_Scope) SetScopeId(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s Brand_Scope) Bind() (Brand_Binding_List, error) {
	if capnp.Struct(s).Uint16(8) != 0 {
		panic("Which() != bind")
	}
	p, err := capnp.Struct(s).Ptr(0)
	return Brand_Binding_List(p.List()), err
}

func (s Brand_Scope) HasBind() bool {
	if capnp.Struct(s).Uint16(8) != 0 {
		return false
	}
	return capnp.Struct(s).HasPtr(0)
}

func (s Brand_Scope) SetBind(v Brand_Binding_List) error {
	capnp.Struct(s).SetUint16(8, 0)
	return capnp.Struct(s).SetPtr(0, v.ToPtr())
}

// NewBind sets the bind field to a newly
// allocated Brand_Binding_List, preferring placement in s's segment.
func (s Brand_Scope) NewBind(n int32) (Brand_Binding_List, error) {
	capnp.Struct(s).SetUint16(8, 0)
	l, err := NewBrand_Binding_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Brand_Binding_List{}, err
	}
	err = capnp.Struct(s).SetPtr(0, l.ToPtr())
	return l, err
}
func (s Brand_Scope) SetInherit() {
	capnp.Struct(s).SetUint16(8, 1)

}

// Brand_Scope_List is a list of Brand_Scope.
type Brand_Scope_List = capnp.StructList[Brand_Scope]

// NewBrand_Scope creates a new list of Brand_Scope.
func NewBrand_Scope_List(s *capnp.Segment, sz int32) (Brand_Scope_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return capnp.StructList[Brand_Scope](l), err
}

type Brand_Binding capnp.Struct
type Brand_Binding_Which uint16

const (
	Brand_Binding_Which_unbound Brand_Binding_Which = 0
	Brand_Binding_Which_type    Brand_Binding_Which = 1
)

func (w Brand_Binding_Which) String() string {
	const s = "unboundtype"
	switch w {
	case Brand_Binding_Which_unbound:
		return s[0:7]
	case Brand_Binding_Which_type:
		return s[7:11]

	}
	return "Brand_Binding_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Brand_Binding_TypeID is the unique identifier for the type Brand_Binding.
const Brand_Binding_TypeID = 0xc863cd16969ee7fc

func NewBrand_Binding(s *capnp.Segment) (Brand_Binding, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Brand_Binding(st), err
}

func NewRootBrand_Binding(s *capnp.Segment) (Brand_Binding, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Brand_Binding(st), err
}

func ReadRootBrand_Binding(msg *capnp.Message) (Brand_Binding, error) {
	root, err := msg.Root()
	return Brand_Binding(root.Struct()), err
}

func (s Brand_Binding) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Brand_Binding) DecodeFromPtr(p capnp.Ptr) Brand_Binding {
	return Brand_Binding(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Brand_Binding) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}

func (s Brand_Binding) Which() Brand_Binding_Which {
	return Brand_Binding_Which(capnp.Struct(s).Uint16(0))
}
func (s Brand_Binding) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Brand_Binding) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Brand_Binding) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Brand_Binding) SetUnbound() {
	capnp.Struct(s).SetUint16(0, 0)

}

func (s Brand_Binding) Type() (Type, error) {
	if capnp.Struct(s).Uint16(0) != 1 {
		panic("Which() != type")
	}
	p, err := capnp.Struct(s).Ptr(0)
	return Type(p.Struct()), err
}

func (s Brand_Binding) HasType() bool {
	if capnp.Struct(s).Uint16(0) != 1 {
		return false
	}
	return capnp.Struct(s).HasPtr(0)
}

func (s Brand_Binding) SetType(v Type) error {
	capnp.Struct(s).SetUint16(0, 1)
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewType sets the type field to a newly
// allocated Type struct, preferring placement in s's segment.
func (s Brand_Binding) NewType() (Type, error) {
	capnp.Struct(s).SetUint16(0, 1)
	ss, err := NewType(capnp.Struct(s).Segment())
	if err != nil {
		return Type{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

// Brand_Binding_List is a list of Brand_Binding.
type Brand_Binding_List = capnp.StructList[Brand_Binding]

// NewBrand_Binding creates a new list of Brand_Binding.
func NewBrand_Binding_List(s *capnp.Segment, sz int32) (Brand_Binding_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[Brand_Binding](l), err
}

type Value capnp.Struct
type Value_Which uint16

const (
	Value_Which_void        Value_Which = 0
	Value_Which_bool        Value_Which = 1
	Value_Which_int8        Value_Which = 2
	Value_Which_int16       Value_Which = 3
	Value_Which_int32       Value_Which = 4
	Value_Which_int64       Value_Which = 5
	Value_Which_uint8       Value_Which = 6
	Value_Which_uint16      Value_Which = 7
	Value_Which_uint32      Value_Which = 8
	Value_Which_uint64      Value_Which = 9
	Value_Which_float32     Value_Which = 10
	Value_Which_float64     Value_Which = 11
	Value_Which_text        Value_Which = 12
	Value_Which_data        Value_Which = 13
	Value_Which_list        Value_Which = 14
	Value_Which_enum        Value_Which = 15
	Value_Which_structValue Value_Which = 16
	Value_Which_interface   Value_Which = 17
	Value_Which_anyPointer  Value_Which = 18
)

func (w Value_Which) String() string {
	const s = "voidboolint8int16int32int64uint8uint16uint32uint64float32float64textdatalistenumstructValueinterfaceanyPointer"
	switch w {
	case Value_Which_void:
		return s[0:4]
	case Value_Which_bool:
		return s[4:8]
	case Value_Which_int8:
		return s[8:12]
	case Value_Which_int16:
		return s[12:17]
	case Value_Which_int32:
		return s[17:22]
	case Value_Which_int64:
		return s[22:27]
	case Value_Which_uint8:
		return s[27:32]
	case Value_Which_uint16:
		return s[32:38]
	case Value_Which_uint32:
		return s[38:44]
	case Value_Which_uint64:
		return s[44:50]
	case Value_Which_float32:
		return s[50:57]
	case Value_Which_float64:
		return s[57:64]
	case Value_Which_text:
		return s[64:68]
	case Value_Which_data:
		return s[68:72]
	case Value_Which_list:
		return s[72:76]
	case Value_Which_enum:
		return s[76:80]
	case Value_Which_structValue:
		return s[80:91]
	case Value_Which_interface:
		return s[91:100]
	case Value_Which_anyPointer:
		return s[100:110]

	}
	return "Value_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Value_TypeID is the unique identifier for the type Value.
const Value_TypeID = 0xce23dcd2d7b00c9b

func NewValue(s *capnp.Segment) (Value, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Value(st), err
}

func NewRootValue(s *capnp.Segment) (Value, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return Value(st), err
}

func ReadRootValue(msg *capnp.Message) (Value, error) {
	root, err := msg.Root()
	return Value(root.Struct()), err
}

func (s Value) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Value) DecodeFromPtr(p capnp.Ptr) Value {
	return Value(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Value) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}

func (s Value) Which() Value_Which {
	return Value_Which(capnp.Struct(s).Uint16(0))
}
func (s Value) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Value) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Value) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Value) SetVoid() {
	capnp.Struct(s).SetUint16(0, 0)

}

func (s Value) Bool() bool {
	if capnp.Struct(s).Uint16(0) != 1 {
		panic("Which() != bool")
	}
	return capnp.Struct(s).Bit(16)
}

func (s Value) SetBool(v bool) {
	capnp.Struct(s).SetUint16(0, 1)
	capnp.Struct(s).SetBit(16, v)
}

func (s Value) Int8() int8 {
	if capnp.Struct(s).Uint16(0) != 2 {
		panic("Which() != int8")
	}
	return int8(capnp.Struct(s).Uint8(2))
}

func (s Value) SetInt8(v int8) {
	capnp.Struct(s).SetUint16(0, 2)
	capnp.Struct(s).SetUint8(2, uint8(v))
}

func (s Value) Int16() int16 {
	if capnp.Struct(s).Uint16(0) != 3 {
		panic("Which() != int16")
	}
	return int16(capnp.Struct(s).Uint16(2))
}

func (s Value) SetInt16(v int16) {
	capnp.Struct(s).SetUint16(0, 3)
	capnp.Struct(s).SetUint16(2, uint16(v))
}

func (s Value) Int32() int32 {
	if capnp.Struct(s).Uint16(0) != 4 {
		panic("Which() != int32")
	}
	return int32(capnp.Struct(s).Uint32(4))
}

func (s Value) SetInt32(v int32) {
	capnp.Struct(s).SetUint16(0, 4)
	capnp.Struct(s).SetUint32(4, uint32(v))
}

func (s Value) Int64() int64 {
	if capnp.Struct(s).Uint16(0) != 5 {
		panic("Which() != int64")
	}
	return int64(capnp.Struct(s).Uint64(8))
}

func (s Value) SetInt64(v int64) {
	capnp.Struct(s).SetUint16(0, 5)
	capnp.Struct(s).SetUint64(8, uint64(v))
}

func (s Value) Uint8() uint8 {
	if capnp.Struct(s).Uint16(0) != 6 {
		panic("Which() != uint8")
	}
	return capnp.Struct(s).Uint8(2)
}

func (s Value) SetUint8(v uint8) {
	capnp.Struct(s).SetUint16(0, 6)
	capnp.Struct(s).SetUint8(2, v)
}

func (s Value) Uint16() uint16 {
	if capnp.Struct(s).Uint16(0) != 7 {
		panic("Which() != uint16")
	}
	return capnp.Struct(s).Uint16(2)
}

func (s Value) SetUint16(v uint16) {
	capnp.Struct(s).SetUint16(0, 7)
	capnp.Struct(s).SetUint16(2, v)
}

func (s Value) Uint32() uint32 {
	if capnp.Struct(s).Uint16(0) != 8 {
		panic("Which() != uint32")
	}
	return capnp.Struct(s).Uint32(4)
}

func (s Value) SetUint32(v uint32) {
	capnp.Struct(s).SetUint16(0, 8)
	capnp.Struct(s).SetUint32(4, v)
}

func (s Value) Uint64() uint64 {
	if capnp.Struct(s).Uint16(0) != 9 {
		panic("Which() != uint64")
	}
	return capnp.Struct(s).Uint64(8)
}

func (s Value) SetUint64(v uint64) {
	capnp.Struct(s).SetUint16(0, 9)
	capnp.Struct(s).SetUint64(8, v)
}

func (s Value) Float32() float32 {
	if capnp.Struct(s).Uint16(0) != 10 {
		panic("Which() != float32")
	}
	return math.Float32frombits(capnp.Struct(s).Uint32(4))
}

func (s Value) SetFloat32(v float32) {
	capnp.Struct(s).SetUint16(0, 10)
	capnp.Struct(s).SetUint32(4, math.Float32bits(v))
}

func (s Value) Float64() float64 {
	if capnp.Struct(s).Uint16(0) != 11 {
		panic("Which() != float64")
	}
	return math.Float64frombits(capnp.Struct(s).Uint64(8))
}

func (s Value) SetFloat64(v float64) {
	capnp.Struct(s).SetUint16(0, 11)
	capnp.Struct(s).SetUint64(8, math.Float64bits(v))
}

func (s Value) Text() (string, error) {
	if capnp.Struct(s).Uint16(0) != 12 {
		panic("Which() != text")
	}
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Value) HasText() bool {
	if capnp.Struct(s).Uint16(0) != 12 {
		return false
	}
	return capnp.Struct(s).HasPtr(0)
}

func (s Value) TextBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Value) SetText(v string) error {
	capnp.Struct(s).SetUint16(0, 12)
	return capnp.Struct(s).SetText(0, v)
}

func (s Value) Data() ([]byte, error) {
	if capnp.Struct(s).Uint16(0) != 13 {
		panic("Which() != data")
	}
	p, err := capnp.Struct(s).Ptr(0)
	return []byte(p.Data()), err
}

func (s Value) HasData() bool {
	if capnp.Struct(s).Uint16(0) != 13 {
		return false
	}
	return capnp.Struct(s).HasPtr(0)
}

func (s Value) SetData(v []byte) error {
	capnp.Struct(s).SetUint16(0, 13)
	return capnp.Struct(s).SetData(0, v)
}

func (s Value) List() (capnp.Ptr, error) {
	if capnp.Struct(s).Uint16(0) != 14 {
		panic("Which() != list")
	}
	return capnp.Struct(s).Ptr(0)
}

func (s Value) HasList() bool {
	if capnp.Struct(s).Uint16(0) != 14 {
		return false
	}
	return capnp.Struct(s).HasPtr(0)
}

func (s Value) SetList(v capnp.Ptr) error {
	capnp.Struct(s).SetUint16(0, 14)
	return capnp.Struct(s).SetPtr(0, v)
}
func (s Value) Enum() uint16 {
	if capnp.Struct(s).Uint16(0) != 15 {
		panic("Which() != enum")
	}
	return capnp.Struct(s).Uint16(2)
}

func (s Value) SetEnum(v uint16) {
	capnp.Struct(s).SetUint16(0, 15)
	capnp.Struct(s).SetUint16(2, v)
}

func (s Value) StructValue() (capnp.Ptr, error) {
	if capnp.Struct(s).Uint16(0) != 16 {
		panic("Which() != structValue")
	}
	return capnp.Struct(s).Ptr(0)
}

func (s Value) HasStructValue() bool {
	if capnp.Struct(s).Uint16(0) != 16 {
		return false
	}
	return capnp.Struct(s).HasPtr(0)
}

func (s Value) SetStructValue(v capnp.Ptr) error {
	capnp.Struct(s).SetUint16(0, 16)
	return capnp.Struct(s).SetPtr(0, v)
}
func (s Value) SetInterface() {
	capnp.Struct(s).SetUint16(0, 17)

}

func (s Value) AnyPointer() (capnp.Ptr, error) {
	if capnp.Struct(s).Uint16(0) != 18 {
		panic("Which() != anyPointer")
	}
	return capnp.Struct(s).Ptr(0)
}

func (s Value) HasAnyPointer() bool {
	if capnp.Struct(s).Uint16(0) != 18 {
		return false
	}
	return capnp.Struct(s).HasPtr(0)
}

func (s Value) SetAnyPointer(v capnp.Ptr) error {
	capnp.Struct(s).SetUint16(0, 18)
	return capnp.Struct(s).SetPtr(0, v)
}

// Value_List is a list of Value.
type Value_List = capnp.StructList[Value]

// NewValue creates a new list of Value.
func NewValue_List(s *capnp.Segment, sz int32) (Value_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return capnp.StructList[Value](l), err
}

type Annotation capnp.Struct

// Annotation_TypeID is the unique identifier for the type Annotation.
const Annotation_TypeID = 0xf1c8950dab257542

func NewAnnotation(s *capnp.Segment) (Annotation, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Annotation(st), err
}

func NewRootAnnotation(s *capnp.Segment) (Annotation, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Annotation(st), err
}

func ReadRootAnnotation(msg *capnp.Message) (Annotation, error) {
	root, err := msg.Root()
	return Annotation(root.Struct()), err
}

func (s Annotation) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Annotation) DecodeFromPtr(p capnp.Ptr) Annotation {
	return Annotation(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Annotation) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Annotation) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Annotation) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Annotation) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Annotation) Id() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s Annotation) SetId(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s Annotation) Brand() (Brand, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return Brand(p.Struct()), err
}

func (s Annotation) HasBrand() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s Annotation) SetBrand(v Brand) error {
	return capnp.Struct(s).SetPtr(1, capnp.Struct(v).ToPtr())
}

// NewBrand sets the brand field to a newly
// allocated Brand struct, preferring placement in s's segment.
func (s Annotation) NewBrand() (Brand, error) {
	ss, err := NewBrand(capnp.Struct(s).Segment())
	if err != nil {
		return Brand{}, err
	}
	err = capnp.Struct(s).SetPtr(1, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s Annotation) Value() (Value, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Value(p.Struct()), err
}

func (s Annotation) HasValue() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Annotation) SetValue(v Value) error {
	return capnp.Struct(s).SetPtr(0, capnp.Struct(v).ToPtr())
}

// NewValue sets the value field to a newly
// allocated Value struct, preferring placement in s's segment.
func (s Annotation) NewValue() (Value, error) {
	ss, err := NewValue(capnp.Struct(s).Segment())
	if err != nil {
		return Value{}, err
	}
	err = capnp.Struct(s).SetPtr(0, capnp.Struct(ss).ToPtr())
	return ss, err
}

// Annotation_List is a list of Annotation.
type Annotation_List = capnp.StructList[Annotation]

// NewAnnotation creates a new list of Annotation.
func NewAnnotation_List(s *capnp.Segment, sz int32) (Annotation_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return capnp.StructList[Annotation](l), err
}

type ElementSize uint16

// ElementSize_TypeID is the unique identifier for the type ElementSize.
const ElementSize_TypeID = 0xd1958f7dba521926

// Values of ElementSize.
const (
	ElementSize_empty           ElementSize = 0
	ElementSize_bit             ElementSize = 1
	ElementSize_byte            ElementSize = 2
	ElementSize_twoBytes        ElementSize = 3
	ElementSize_fourBytes       ElementSize = 4
	ElementSize_eightBytes      ElementSize = 5
	ElementSize_pointer         ElementSize = 6
	ElementSize_inlineComposite ElementSize = 7
)

// String returns the enum's constant name.
func (c ElementSize) String() string {
	switch c {
	case ElementSize_empty:
		return "empty"
	case ElementSize_bit:
		return "bit"
	case ElementSize_byte:
		return "byte"
	case ElementSize_twoBytes:
		return "twoBytes"
	case ElementSize_fourBytes:
		return "fourBytes"
	case ElementSize_eightBytes:
		return "eightBytes"
	case ElementSize_pointer:
		return "pointer"
	case ElementSize_inlineComposite:
		return "inlineComposite"

	default:
		return ""
	}
}

// ElementSizeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func ElementSizeFromString(c string) ElementSize {
	switch c {
	case "empty":
		return ElementSize_empty
	case "bit":
		return ElementSize_bit
	case "byte":
		return ElementSize_byte
	case "twoBytes":
		return ElementSize_twoBytes
	case "fourBytes":
		return ElementSize_fourBytes
	case "eightBytes":
		return ElementSize_eightBytes
	case "pointer":
		return ElementSize_pointer
	case "inlineComposite":
		return ElementSize_inlineComposite

	default:
		return 0
	}
}

type ElementSize_List = capnp.EnumList[ElementSize]

func NewElementSize_List(s *capnp.Segment, sz int32) (ElementSize_List, error) {
	return capnp.NewEnumList[ElementSize](s, sz)
}

type CapnpVersion capnp.Struct

// CapnpVersion_TypeID is the unique identifier for the type CapnpVersion.
const CapnpVersion_TypeID = 0xd85d305b7d839963

func NewCapnpVersion(s *capnp.Segment) (CapnpVersion, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return CapnpVersion(st), err
}

func NewRootCapnpVersion(s *capnp.Segment) (CapnpVersion, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return CapnpVersion(st), err
}

func ReadRootCapnpVersion(msg *capnp.Message) (CapnpVersion, error) {
	root, err := msg.Root()
	return CapnpVersion(root.Struct()), err
}

func (s CapnpVersion) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CapnpVersion) DecodeFromPtr(p capnp.Ptr) CapnpVersion {
	return CapnpVersion(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CapnpVersion) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CapnpVersion) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CapnpVersion) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CapnpVersion) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s CapnpVersion) Major() uint16 {
	return capnp.Struct(s).Uint16(0)
}

func (s CapnpVersion) SetMajor(v uint16) {
	capnp.Struct(s).SetUint16(0, v)
}

func (s CapnpVersion) Minor() uint8 {
	return capnp.Struct(s).Uint8(2)
}

func (s CapnpVersion) SetMinor(v uint8) {
	capnp.Struct(s).SetUint8(2, v)
}

func (s CapnpVersion) Micro() uint8 {
	return capnp.Struct(s).Uint8(3)
}

func (s CapnpVersion) SetMicro(v uint8) {
	capnp.Struct(s).SetUint8(3, v)
}

// CapnpVersion_List is a list of CapnpVersion.
type CapnpVersion_List = capnp.StructList[CapnpVersion]

// NewCapnpVersion creates a new list of CapnpVersion.
func NewCapnpVersion_List(s *capnp.Segment, sz int32) (CapnpVersion_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return capnp.StructList[CapnpVersion](l), err
}

type CodeGeneratorRequest capnp.Struct

// CodeGeneratorRequest_TypeID is the unique identifier for the type CodeGeneratorRequest.
const CodeGeneratorRequest_TypeID = 0xbfc546f6210ad7ce

func NewCodeGeneratorRequest(s *capnp.Segment) (CodeGeneratorRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return CodeGeneratorRequest(st), err
}

func NewRootCodeGeneratorRequest(s *capnp.Segment) (CodeGeneratorRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4})
	return CodeGeneratorRequest(st), err
}

func ReadRootCodeGeneratorRequest(msg *capnp.Message) (CodeGeneratorRequest, error) {
	root, err := msg.Root()
	return CodeGeneratorRequest(root.Struct()), err
}

func (s CodeGeneratorRequest) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CodeGeneratorRequest) DecodeFromPtr(p capnp.Ptr) CodeGeneratorRequest {
	return CodeGeneratorRequest(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CodeGeneratorRequest) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CodeGeneratorRequest) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CodeGeneratorRequest) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CodeGeneratorRequest) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s CodeGeneratorRequest) CapnpVersion() (CapnpVersion, error) {
	p, err := capnp.Struct(s).Ptr(2)
	return CapnpVersion(p.Struct()), err
}

func (s CodeGeneratorRequest) HasCapnpVersion() bool {
	return capnp.Struct(s).HasPtr(2)
}

func (s CodeGeneratorRequest) SetCapnpVersion(v CapnpVersion) error {
	return capnp.Struct(s).SetPtr(2, capnp.Struct(v).ToPtr())
}

// NewCapnpVersion sets the capnpVersion field to a newly
// allocated CapnpVersion struct, preferring placement in s's segment.
func (s CodeGeneratorRequest) NewCapnpVersion() (CapnpVersion, error) {
	ss, err := NewCapnpVersion(capnp.Struct(s).Segment())
	if err != nil {
		return CapnpVersion{}, err
	}
	err = capnp.Struct(s).SetPtr(2, capnp.Struct(ss).ToPtr())
	return ss, err
}

func (s CodeGeneratorRequest) Nodes() (Node_List, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return Node_List(p.List()), err
}

func (s CodeGeneratorRequest) HasNodes() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s CodeGeneratorRequest) SetNodes(v Node_List) error {
	return capnp.Struct(s).SetPtr(0, v.ToPtr())
}

// NewNodes sets the nodes field to a newly
// allocated Node_List, preferring placement in s's segment.
func (s CodeGeneratorRequest) NewNodes(n int32) (Node_List, error) {
	l, err := NewNode_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Node_List{}, err
	}
	err = capnp.Struct(s).SetPtr(0, l.ToPtr())
	return l, err
}
func (s CodeGeneratorRequest) SourceInfo() (Node_SourceInfo_List, error) {
	p, err := capnp.Struct(s).Ptr(3)
	return Node_SourceInfo_List(p.List()), err
}

func (s CodeGeneratorRequest) HasSourceInfo() bool {
	return capnp.Struct(s).HasPtr(3)
}

func (s CodeGeneratorRequest) SetSourceInfo(v Node_SourceInfo_List) error {
	return capnp.Struct(s).SetPtr(3, v.ToPtr())
}

// NewSourceInfo sets the sourceInfo field to a newly
// allocated Node_SourceInfo_List, preferring placement in s's segment.
func (s CodeGeneratorRequest) NewSourceInfo(n int32) (Node_SourceInfo_List, error) {
	l, err := NewNode_SourceInfo_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return Node_SourceInfo_List{}, err
	}
	err = capnp.Struct(s).SetPtr(3, l.ToPtr())
	return l, err
}
func (s CodeGeneratorRequest) RequestedFiles() (CodeGeneratorRequest_RequestedFile_List, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return CodeGeneratorRequest_RequestedFile_List(p.List()), err
}

func (s CodeGeneratorRequest) HasRequestedFiles() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s CodeGeneratorRequest) SetRequestedFiles(v CodeGeneratorRequest_RequestedFile_List) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewRequestedFiles sets the requestedFiles field to a newly
// allocated CodeGeneratorRequest_RequestedFile_List, preferring placement in s's segment.
func (s CodeGeneratorRequest) NewRequestedFiles(n int32) (CodeGeneratorRequest_RequestedFile_List, error) {
	l, err := NewCodeGeneratorRequest_RequestedFile_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return CodeGeneratorRequest_RequestedFile_List{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}

// CodeGeneratorRequest_List is a list of CodeGeneratorRequest.
type CodeGeneratorRequest_List = capnp.StructList[CodeGeneratorRequest]

// NewCodeGeneratorRequest creates a new list of CodeGeneratorRequest.
func NewCodeGeneratorRequest_List(s *capnp.Segment, sz int32) (CodeGeneratorRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 4}, sz)
	return capnp.StructList[CodeGeneratorRequest](l), err
}

type CodeGeneratorRequest_RequestedFile capnp.Struct

// CodeGeneratorRequest_RequestedFile_TypeID is the unique identifier for the type CodeGeneratorRequest_RequestedFile.
const CodeGeneratorRequest_RequestedFile_TypeID = 0xcfea0eb02e810062

func NewCodeGeneratorRequest_RequestedFile(s *capnp.Segment) (CodeGeneratorRequest_RequestedFile, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return CodeGeneratorRequest_RequestedFile(st), err
}

func NewRootCodeGeneratorRequest_RequestedFile(s *capnp.Segment) (CodeGeneratorRequest_RequestedFile, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return CodeGeneratorRequest_RequestedFile(st), err
}

func ReadRootCodeGeneratorRequest_RequestedFile(msg *capnp.Message) (CodeGeneratorRequest_RequestedFile, error) {
	root, err := msg.Root()
	return CodeGeneratorRequest_RequestedFile(root.Struct()), err
}

func (s CodeGeneratorRequest_RequestedFile) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CodeGeneratorRequest_RequestedFile) DecodeFromPtr(p capnp.Ptr) CodeGeneratorRequest_RequestedFile {
	return CodeGeneratorRequest_RequestedFile(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CodeGeneratorRequest_RequestedFile) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CodeGeneratorRequest_RequestedFile) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CodeGeneratorRequest_RequestedFile) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CodeGeneratorRequest_RequestedFile) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s CodeGeneratorRequest_RequestedFile) Id() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s CodeGeneratorRequest_RequestedFile) SetId(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s CodeGeneratorRequest_RequestedFile) Filename() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s CodeGeneratorRequest_RequestedFile) HasFilename() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s CodeGeneratorRequest_RequestedFile) FilenameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s CodeGeneratorRequest_RequestedFile) SetFilename(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s CodeGeneratorRequest_RequestedFile) Imports() (CodeGeneratorRequest_RequestedFile_Import_List, error) {
	p, err := capnp.Struct(s).Ptr(1)
	return CodeGeneratorRequest_RequestedFile_Import_List(p.List()), err
}

func (s CodeGeneratorRequest_RequestedFile) HasImports() bool {
	return capnp.Struct(s).HasPtr(1)
}

func (s CodeGeneratorRequest_RequestedFile) SetImports(v CodeGeneratorRequest_RequestedFile_Import_List) error {
	return capnp.Struct(s).SetPtr(1, v.ToPtr())
}

// NewImports sets the imports field to a newly
// allocated CodeGeneratorRequest_RequestedFile_Import_List, preferring placement in s's segment.
func (s CodeGeneratorRequest_RequestedFile) NewImports(n int32) (CodeGeneratorRequest_RequestedFile_Import_List, error) {
	l, err := NewCodeGeneratorRequest_RequestedFile_Import_List(capnp.Struct(s).Segment(), n)
	if err != nil {
		return CodeGeneratorRequest_RequestedFile_Import_List{}, err
	}
	err = capnp.Struct(s).SetPtr(1, l.ToPtr())
	return l, err
}

// CodeGeneratorRequest_RequestedFile_List is a list of CodeGeneratorRequest_RequestedFile.
type CodeGeneratorRequest_RequestedFile_List = capnp.StructList[CodeGeneratorRequest_RequestedFile]

// NewCodeGeneratorRequest_RequestedFile creates a new list of CodeGeneratorRequest_RequestedFile.
func NewCodeGeneratorRequest_RequestedFile_List(s *capnp.Segment, sz int32) (CodeGeneratorRequest_RequestedFile_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return capnp.StructList[CodeGeneratorRequest_RequestedFile](l), err
}

type CodeGeneratorRequest_RequestedFile_Import capnp.Struct

// CodeGeneratorRequest_RequestedFile_Import_TypeID is the unique identifier for the type CodeGeneratorRequest_RequestedFile_Import.
const CodeGeneratorRequest_RequestedFile_Import_TypeID = 0xae504193122357e5

func NewCodeGeneratorRequest_RequestedFile_Import(s *capnp.Segment) (CodeGeneratorRequest_RequestedFile_Import, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CodeGeneratorRequest_RequestedFile_Import(st), err
}

func NewRootCodeGeneratorRequest_RequestedFile_Import(s *capnp.Segment) (CodeGeneratorRequest_RequestedFile_Import, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CodeGeneratorRequest_RequestedFile_Import(st), err
}

func ReadRootCodeGeneratorRequest_RequestedFile_Import(msg *capnp.Message) (CodeGeneratorRequest_RequestedFile_Import, error) {
	root, err := msg.Root()
	return CodeGeneratorRequest_RequestedFile_Import(root.Struct()), err
}

func (s CodeGeneratorRequest_RequestedFile_Import) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (CodeGeneratorRequest_RequestedFile_Import) DecodeFromPtr(p capnp.Ptr) CodeGeneratorRequest_RequestedFile_Import {
	return CodeGeneratorRequest_RequestedFile_Import(capnp.Struct{}.DecodeFromPtr(p))
}

func (s CodeGeneratorRequest_RequestedFile_Import) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s CodeGeneratorRequest_RequestedFile_Import) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s CodeGeneratorRequest_RequestedFile_Import) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s CodeGeneratorRequest_RequestedFile_Import) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s CodeGeneratorRequest_RequestedFile_Import) Id() uint64 {
	return capnp.Struct(s).Uint64(0)
}

func (s CodeGeneratorRequest_RequestedFile_Import) SetId(v uint64) {
	capnp.Struct(s).SetUint64(0, v)
}

func (s CodeGeneratorRequest_RequestedFile_Import) Name() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s CodeGeneratorRequest_RequestedFile_Import) HasName() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s CodeGeneratorRequest_RequestedFile_Import) NameBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s CodeGeneratorRequest_RequestedFile_Import) SetName(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

// CodeGeneratorRequest_RequestedFile_Import_List is a list of CodeGeneratorRequest_RequestedFile_Import.
type CodeGeneratorRequest_RequestedFile_Import_List = capnp.StructList[CodeGeneratorRequest_RequestedFile_Import]

// NewCodeGeneratorRequest_RequestedFile_Import creates a new list of CodeGeneratorRequest_RequestedFile_Import.
func NewCodeGeneratorRequest_RequestedFile_Import_List(s *capnp.Segment, sz int32) (CodeGeneratorRequest_RequestedFile_Import_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[CodeGeneratorRequest_RequestedFile_Import](l), err
}
